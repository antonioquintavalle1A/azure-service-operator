/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package identity

import (
	"context"
	"reflect"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/msi-dataplane/pkg/dataplane"
	"github.com/rotisserie/eris"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	asocloud "github.com/Azure/azure-service-operator/v2/pkg/common/cloud"
	"github.com/Azure/azure-service-operator/v2/pkg/common/config"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

const (
	// #nosec
	globalCredentialSecretName = "aso-controller-settings"
	// #nosec
	NamespacedSecretName = "aso-credential"
	// #nosec
	FederatedTokenFilePath = "/var/run/secrets/tokens/azure-identity"
)

// Credential describes a credential used to connect to Azure
type Credential struct {
	tokenCredential   azcore.TokenCredential
	credentialFrom    types.NamespacedName
	subscriptionID    string
	additionalTenants []string

	// secretData contains the secret
	secretData map[string][]byte
}

func (c *Credential) SecretsEqual(other *Credential) bool {
	return reflect.DeepEqual(c.secretData, other.secretData)
}

func (c *Credential) CredentialFrom() types.NamespacedName {
	return c.credentialFrom
}

func (c *Credential) SubscriptionID() string {
	return c.subscriptionID
}

func (c *Credential) TokenCredential() azcore.TokenCredential {
	return c.tokenCredential
}

func (c *Credential) AdditionalTenants() []string {
	return c.additionalTenants
}

func NewDefaultCredential(
	tokenCred azcore.TokenCredential,
	namespace string,
	subscriptionID string,
	additionalTenants []string,
) *Credential {
	return &Credential{
		tokenCredential:   tokenCred,
		subscriptionID:    subscriptionID,
		credentialFrom:    types.NamespacedName{Namespace: namespace, Name: globalCredentialSecretName},
		additionalTenants: additionalTenants,
	}
}

type CredentialProvider interface {
	GetCredential(ctx context.Context, obj genruntime.MetaObject) (*Credential, error)
}

type CredentialProviderOptions struct {
	TokenProvider TokenCredentialProvider
	Cloud         *cloud.Configuration
}

type credentialProvider struct {
	globalCredential        *Credential
	kubeClient              kubeclient.Client
	tokenCredentialProvider TokenCredentialProvider
	cloud                   cloud.Configuration
}

func NewCredentialProvider(
	globalCredential *Credential,
	kubeClient kubeclient.Client,
	opts *CredentialProviderOptions,
) CredentialProvider {
	if opts == nil {
		opts = &CredentialProviderOptions{}
	}

	if opts.TokenProvider == nil {
		opts.TokenProvider = DefaultTokenCredentialProvider()
	}

	var cloud cloud.Configuration
	if opts.Cloud == nil {
		cloud = asocloud.Configuration{}.Cloud()
	} else {
		cloud = *opts.Cloud
	}

	return &credentialProvider{
		kubeClient:              kubeClient,
		globalCredential:        globalCredential,
		tokenCredentialProvider: opts.TokenProvider,
		cloud:                   cloud,
	}
}

// GetCredential returns the credential used to manage the obj. It performs credential lookup according to the following
// procedure:
//  1. Per-Resource credential specified as an annotation ("serviceoperator.azure.com/credential-from") directly on
//     the resource.
//  2. Per-Namespace credential provided at the namespace scope (a secret named "aso-credential" in the resource's
//     namespace).
//  3. Global credential for the operator.
//
// If no matching credential can be found, an error is returned.
func (c *credentialProvider) GetCredential(ctx context.Context, obj genruntime.MetaObject) (*Credential, error) {
	// Resource annotation
	cred, err := c.getCredentialFromAnnotation(ctx, obj, annotations.PerResourceSecret)
	if err != nil {
		return nil, err
	}
	if cred != nil {
		return cred, nil
	}

	// Namespaced secret
	cred, err = c.getCredentialFromNamespaceSecret(ctx, obj.GetNamespace())
	if err != nil {
		return nil, err
	}
	if cred != nil {
		return cred, nil
	}

	// Global credential
	if c.globalCredential == nil {
		return nil, eris.New("global credential not configured, you must use either namespaced or per-resource credentials")
	}

	// If not found, default is the global credential
	return c.globalCredential, nil
}

// getCredentialFromNamespaceSecret creates a Credential from the namespace scoped secret (a secret named
// "aso-credential" in the same namespace as the resource).
// If the aso-credential secret is not found, an error is returned.
func (c *credentialProvider) getCredentialFromNamespaceSecret(ctx context.Context, namespace string) (*Credential, error) {
	secretNamespacedName := types.NamespacedName{Namespace: namespace, Name: NamespacedSecretName}
	cred, err := c.getCredentialFromSecret(ctx, secretNamespacedName)
	if err != nil {
		var secretNotFound *core.SecretNotFound
		if eris.As(err, &secretNotFound) {
			return nil, nil // Not finding this secret is allowed, allow caller to proceed to higher scope secret
		}
		return nil, err
	}

	return cred, nil
}

// getCredentialFromAnnotation creates a Credential from the secret referenced in the specified annotation.
// The secret must be in the same namespace as the obj parameter.
// If the annotation doesn't exist, a nil credential is returned
// If the annotation exists but refers to a secret that does not exist, an error is returned.
func (c *credentialProvider) getCredentialFromAnnotation(ctx context.Context, obj genruntime.MetaObject, annotation string) (*Credential, error) {
	credentialFrom, ok := obj.GetAnnotations()[annotation]
	if !ok {
		return nil, nil
	}

	// Disallow credentials with `/` in their credentialFrom
	if strings.Contains(credentialFrom, "/") {
		err := eris.Errorf("%s cannot contain '/'. Secret must be in same namespace as resource.", annotation)
		namespacedName := types.NamespacedName{
			Namespace: obj.GetNamespace(),
			Name:      credentialFrom,
		}
		return nil, core.NewSecretNotFoundError(namespacedName, err)
	}

	// annotation exists, use specified secret
	secretNamespacedName := getSecretNameFromAnnotation(credentialFrom, obj.GetNamespace())
	return c.getCredentialFromSecret(ctx, secretNamespacedName)
}

func (c *credentialProvider) getCredentialFromSecret(ctx context.Context, secretNamespacedName types.NamespacedName) (*Credential, error) {
	secret, err := c.getSecret(ctx, secretNamespacedName.Namespace, secretNamespacedName.Name)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, core.NewSecretNotFoundError(secretNamespacedName, eris.Wrapf(err, "credential secret not found"))
		}
		return nil, err
	}

	//nolint:contextcheck
	cred, err := c.newCredentialFromSecret(secret)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to get credential %q", secretNamespacedName)
	}
	return cred, nil
}

func (c *credentialProvider) newCredentialFromSecret(secret *v1.Secret) (*Credential, error) {
	var errs []error

	nsName := types.NamespacedName{Namespace: secret.GetNamespace(), Name: secret.GetName()}

	subscriptionIDBytes, ok := secret.Data[config.AzureSubscriptionID]
	if !ok {
		err := core.NewSecretNotFoundError(
			nsName,
			eris.Errorf(
				"credential Secret %q does not contain key %q",
				nsName,
				config.AzureSubscriptionID),
		)
		errs = append(errs, err)
	}

	tenantIDBytes, ok := secret.Data[config.AzureTenantID]
	if !ok {
		err := core.NewSecretNotFoundError(
			nsName,
			eris.Errorf(
				"credential Secret %q does not contain key %q",
				nsName,
				config.AzureTenantID),
		)
		errs = append(errs, err)
	}

	clientIDBytes, ok := secret.Data[config.AzureClientID]
	if !ok {
		err := core.NewSecretNotFoundError(
			nsName,
			eris.Errorf(
				"credential Secret %q does not contain key %q",
				nsName,
				config.AzureClientID),
		)
		errs = append(errs, err)
	}

	subscriptionID := string(subscriptionIDBytes)
	tenantID := string(tenantIDBytes)
	clientID := string(clientIDBytes)

	// Read optional fields
	var additionalTenants []string
	additionalTenantsBytes, ok := secret.Data[config.AzureAdditionalTenants]
	if ok {
		additionalTenants = config.ParseCommaCollection(string(additionalTenantsBytes))
	}

	// TODO: We could read AzureAuthorityHost and other "cloud-defining"
	// TODO: variables here to allow per-cloud credentials.

	// Missing required properties, fail fast
	if len(errs) > 0 {
		return nil, kerrors.NewAggregate(errs)
	}

	if clientSecret, hasClientSecret := secret.Data[config.AzureClientSecret]; hasClientSecret {
		tokenCredential, err := c.tokenCredentialProvider.NewClientSecretCredential(
			tenantID,
			clientID,
			string(clientSecret),
			&azidentity.ClientSecretCredentialOptions{
				ClientOptions: azcore.ClientOptions{
					Cloud: c.cloud,
				},
				AdditionallyAllowedTenants: additionalTenants,
			})
		if err != nil {
			return nil, eris.Wrap(err, eris.Errorf("invalid Client Secret Credential for %q encountered", nsName).Error())
		}

		return &Credential{
			tokenCredential:   tokenCredential,
			subscriptionID:    subscriptionID,
			credentialFrom:    nsName,
			additionalTenants: additionalTenants,
			secretData:        secret.Data,
		}, nil
	}

	if clientCert, hasClientCert := secret.Data[config.AzureClientCertificate]; hasClientCert {
		var clientCertPassword []byte
		if p, hasClientCertPassword := secret.Data[config.AzureClientCertificatePassword]; hasClientCertPassword {
			clientCertPassword = p
		}

		tokenCredential, err := c.tokenCredentialProvider.NewClientCertificateCredential(
			tenantID,
			clientID,
			clientCert,
			clientCertPassword,
			&azidentity.ClientCertificateCredentialOptions{
				ClientOptions: azcore.ClientOptions{
					Cloud: c.cloud,
				},
				AdditionallyAllowedTenants: additionalTenants,
			})
		if err != nil {
			return nil, eris.Wrap(err, eris.Errorf("invalid Client Certificate Credential for %q encountered", nsName).Error())
		}

		return &Credential{
			tokenCredential:   tokenCredential,
			subscriptionID:    subscriptionID,
			credentialFrom:    nsName,
			additionalTenants: additionalTenants,
			secretData:        secret.Data,
		}, nil
	}

	// This authentication type is similar to user assigned managed identity authentication combined with client certificate
	// authentication. As a 1st party Microsoft application, one has access to pull a user assigned managed identity's backing
	// certificate information from the MSI data plane. Using this data, a user can authenticate to Azure Cloud.
	if userAssignedCredentialsPath, hasUserAssignedCredentials := secret.Data[config.AzureUserAssignedIdentityCredentials]; hasUserAssignedCredentials {
		// Default to AzurePublic
		options := azcore.ClientOptions{
			Cloud: c.cloud,
		}

		tokenCredential, err := c.tokenCredentialProvider.NewUserAssignedIdentityCredentials(context.Background(), string(userAssignedCredentialsPath), dataplane.WithClientOpts(options))
		if err != nil {
			return nil, eris.Wrap(err, eris.Errorf("invalid User Assigned Identity Credential for %q encountered", nsName).Error())
		}

		return &Credential{
			tokenCredential: tokenCredential,
			subscriptionID:  string(subscriptionID),
			credentialFrom:  nsName,
			secretData:      secret.Data,
		}, nil
	}

	if value, hasAuthMode := secret.Data[config.AuthMode]; hasAuthMode {
		authMode, err := authModeOrDefault(string(value))
		if err != nil {
			return nil, eris.Wrap(err, eris.Errorf("invalid identity auth mode for %q encountered", nsName).Error())
		}

		if authMode == config.PodIdentityAuthMode {
			// Note: ManagedIdentityCredentials don't support additionalTenants because Managed Identities
			// are inherently single-tenant, so we don't pass
			tokenCredential, err := c.tokenCredentialProvider.NewManagedIdentityCredential(&azidentity.ManagedIdentityCredentialOptions{
				ClientOptions: azcore.ClientOptions{
					Cloud: c.cloud,
				},
				ID: azidentity.ClientID(clientID),
			})
			if err != nil {
				return nil, eris.Wrap(err, eris.Errorf("invalid Managed Identity for %q encountered", nsName).Error())
			}

			return &Credential{
				tokenCredential: tokenCredential,
				subscriptionID:  subscriptionID,
				credentialFrom:  nsName,
				secretData:      secret.Data,
			}, nil
		}
	}

	// Default to Workload Identity
	tokenCredential, err := c.tokenCredentialProvider.NewWorkloadIdentityCredential(
		&azidentity.WorkloadIdentityCredentialOptions{
			ClientOptions: azcore.ClientOptions{
				Cloud: c.cloud,
			},
			ClientID:                   clientID,
			TenantID:                   tenantID,
			TokenFilePath:              FederatedTokenFilePath,
			AdditionallyAllowedTenants: additionalTenants,
		})
	if err != nil {
		err = eris.Wrapf(
			err,
			"credential secret %q does not contain key %q and failed to get workload identity credential for clientID %q from %q ",
			nsName,
			config.AzureClientSecret,
			clientID,
			FederatedTokenFilePath)

		return nil, err
	}

	return &Credential{
		tokenCredential:   tokenCredential,
		subscriptionID:    subscriptionID,
		credentialFrom:    nsName,
		additionalTenants: additionalTenants,
		secretData:        secret.Data,
	}, nil
}

func (c *credentialProvider) getSecret(ctx context.Context, namespace string, secretName string) (*v1.Secret, error) {
	secret := &v1.Secret{}

	err := c.kubeClient.Get(
		ctx,
		types.NamespacedName{Namespace: namespace, Name: secretName},
		secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

func getSecretNameFromAnnotation(credentialFrom string, resourceNamespace string) types.NamespacedName {
	return types.NamespacedName{Namespace: resourceNamespace, Name: credentialFrom}
}

func authModeOrDefault(mode string) (config.AuthModeOption, error) {
	if mode == string(config.WorkloadIdentityAuthMode) || mode == "" {
		return config.WorkloadIdentityAuthMode, nil
	}

	if mode == string(config.PodIdentityAuthMode) {
		return config.PodIdentityAuthMode, nil
	}

	return "", eris.Errorf("authorization mode %q not valid", mode)
}
