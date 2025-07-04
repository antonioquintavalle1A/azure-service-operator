/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package sql

import (
	"context"

	"github.com/rotisserie/eris"
	ctrlconversion "sigs.k8s.io/controller-runtime/pkg/conversion"

	asosql "github.com/Azure/azure-service-operator/v2/api/sql/v1"
	sql "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101/storage"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

type Connector interface {
	CreateOrUpdate(ctx context.Context) error
	Delete(ctx context.Context) error
	Exists(ctx context.Context) (bool, error)
}

type ownerDetails struct {
	fqdn     string
	database string
}

func getOwnerDetails(ctx context.Context, resourceResolver *resolver.Resolver, user *asosql.User) (ownerDetails, error) {
	// Get the owner - at this point it must exist
	owner, err := resourceResolver.ResolveOwner(ctx, user)
	if err != nil {
		return ownerDetails{}, err
	}

	// Note that this is not actually possible for this type because we don't allow ARMID references for these owners,
	// but protecting against it here anyway.
	if !owner.FoundKubernetesOwner() {
		return ownerDetails{}, eris.Errorf("user owner must exist in Kubernetes for user %s", user.Name)
	}

	hierarchy, err := resourceResolver.ResolveResourceHierarchy(ctx, owner.Owner)
	if err != nil {
		return ownerDetails{}, err
	}

	var server *sql.Server
	var database *sql.ServersDatabase
	for _, object := range hierarchy {
		switch object.(type) {
		case *sql.Server:
			server = object.(*sql.Server)
		case *sql.ServersDatabase:
			database = object.(*sql.ServersDatabase)
		default:
		}
	}

	if server == nil {
		return ownerDetails{}, eris.New("could not find a value of type Server in the hierarchy")
	}

	if database == nil {
		return ownerDetails{}, eris.New("could not find a value of type ServersDatabase in the hierarchy")
	}

	// Assertion to ensure that this is still the storage type
	// If this doesn't compile, update the version being imported to the new Hub version
	var _ ctrlconversion.Hub = &sql.ServersDatabase{}
	// Assertion to ensure that this is still the storage type
	// If this doesn't compile, update the version being imported to the new Hub version
	var _ ctrlconversion.Hub = &sql.Server{}

	// TODO: Possibly we want to source this from a configmap rather than reading it directly off of our parent?
	// TODO: That would work better with https://github.com/Azure/azure-service-operator/issues/2357 as well probably?
	// TODO: Not doing this now as the other SQL users use the parent reference directly and we want to be consistent with them.
	// TODO: If we do support this we can do it in a non-breaking way where we add support for users and move to allow ARM-ID based
	// TODO: owners.
	if server.Status.FullyQualifiedDomainName == nil {
		// This possibly means that the server hasn't finished deploying yet
		err = eris.Errorf("owning Server %q '.status.fullyQualifiedDomainName' not set. Has the server been provisioned successfully?", server.Name)
		return ownerDetails{}, conditions.NewReadyConditionImpactingError(err, conditions.ConditionSeverityWarning, conditions.ReasonWaitingForOwner)
	}
	serverFQDN := *server.Status.FullyQualifiedDomainName

	return ownerDetails{
		fqdn:     serverFQDN,
		database: database.AzureName(),
	}, nil
}
