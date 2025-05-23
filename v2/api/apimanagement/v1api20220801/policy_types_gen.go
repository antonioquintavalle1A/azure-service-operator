// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801

import (
	"fmt"
	arm "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/arm"
	storage "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimpolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies/{policyId}
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Policy_Spec   `json:"spec,omitempty"`
	Status            Policy_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Policy{}

// GetConditions returns the conditions of the resource
func (policy *Policy) GetConditions() conditions.Conditions {
	return policy.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (policy *Policy) SetConditions(conditions conditions.Conditions) {
	policy.Status.Conditions = conditions
}

var _ conversion.Convertible = &Policy{}

// ConvertFrom populates our Policy from the provided hub Policy
func (policy *Policy) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.Policy)
	if !ok {
		return fmt.Errorf("expected apimanagement/v1api20220801/storage/Policy but received %T instead", hub)
	}

	return policy.AssignProperties_From_Policy(source)
}

// ConvertTo populates the provided hub Policy from our Policy
func (policy *Policy) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.Policy)
	if !ok {
		return fmt.Errorf("expected apimanagement/v1api20220801/storage/Policy but received %T instead", hub)
	}

	return policy.AssignProperties_To_Policy(destination)
}

var _ configmaps.Exporter = &Policy{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (policy *Policy) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if policy.Spec.OperatorSpec == nil {
		return nil
	}
	return policy.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &Policy{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (policy *Policy) SecretDestinationExpressions() []*core.DestinationExpression {
	if policy.Spec.OperatorSpec == nil {
		return nil
	}
	return policy.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.ImportableResource = &Policy{}

// InitializeSpec initializes the spec for this resource from the given status
func (policy *Policy) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*Policy_STATUS); ok {
		return policy.Spec.Initialize_From_Policy_STATUS(s)
	}

	return fmt.Errorf("expected Status of type Policy_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &Policy{}

// AzureName returns the Azure name of the resource (always "policy")
func (policy *Policy) AzureName() string {
	return "policy"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-08-01"
func (policy Policy) GetAPIVersion() string {
	return "2022-08-01"
}

// GetResourceScope returns the scope of the resource
func (policy *Policy) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (policy *Policy) GetSpec() genruntime.ConvertibleSpec {
	return &policy.Spec
}

// GetStatus returns the status of this resource
func (policy *Policy) GetStatus() genruntime.ConvertibleStatus {
	return &policy.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (policy *Policy) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationHead,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/policies"
func (policy *Policy) GetType() string {
	return "Microsoft.ApiManagement/service/policies"
}

// NewEmptyStatus returns a new empty (blank) status
func (policy *Policy) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Policy_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (policy *Policy) Owner() *genruntime.ResourceReference {
	if policy.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(policy.Spec)
	return policy.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (policy *Policy) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Policy_STATUS); ok {
		policy.Status = *st
		return nil
	}

	// Convert status to required version
	var st Policy_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	policy.Status = st
	return nil
}

// AssignProperties_From_Policy populates our Policy from the provided source Policy
func (policy *Policy) AssignProperties_From_Policy(source *storage.Policy) error {

	// ObjectMeta
	policy.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Policy_Spec
	err := spec.AssignProperties_From_Policy_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_Policy_Spec() to populate field Spec")
	}
	policy.Spec = spec

	// Status
	var status Policy_STATUS
	err = status.AssignProperties_From_Policy_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_Policy_STATUS() to populate field Status")
	}
	policy.Status = status

	// No error
	return nil
}

// AssignProperties_To_Policy populates the provided destination Policy from our Policy
func (policy *Policy) AssignProperties_To_Policy(destination *storage.Policy) error {

	// ObjectMeta
	destination.ObjectMeta = *policy.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.Policy_Spec
	err := policy.Spec.AssignProperties_To_Policy_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_Policy_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.Policy_STATUS
	err = policy.Status.AssignProperties_To_Policy_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_Policy_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (policy *Policy) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: policy.Spec.OriginalVersion(),
		Kind:    "Policy",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimpolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies/{policyId}
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

type Policy_Spec struct {
	// Format: Format of the policyContent.
	Format *PolicyContractProperties_Format `json:"format,omitempty"`

	// OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
	// passed directly to Azure
	OperatorSpec *PolicyOperatorSpec `json:"operatorSpec,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a apimanagement.azure.com/Service resource
	Owner *genruntime.KnownResourceReference `group:"apimanagement.azure.com" json:"owner,omitempty" kind:"Service"`

	// +kubebuilder:validation:Required
	// Value: Contents of the Policy as defined by the format.
	Value *string `json:"value,omitempty"`
}

var _ genruntime.ARMTransformer = &Policy_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (policy *Policy_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if policy == nil {
		return nil, nil
	}
	result := &arm.Policy_Spec{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if policy.Format != nil || policy.Value != nil {
		result.Properties = &arm.PolicyContractProperties{}
	}
	if policy.Format != nil {
		var temp string
		temp = string(*policy.Format)
		format := arm.PolicyContractProperties_Format(temp)
		result.Properties.Format = &format
	}
	if policy.Value != nil {
		value := *policy.Value
		result.Properties.Value = &value
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (policy *Policy_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.Policy_Spec{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (policy *Policy_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.Policy_Spec)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.Policy_Spec, got %T", armInput)
	}

	// Set property "Format":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Format != nil {
			var temp string
			temp = string(*typedInput.Properties.Format)
			format := PolicyContractProperties_Format(temp)
			policy.Format = &format
		}
	}

	// no assignment for property "OperatorSpec"

	// Set property "Owner":
	policy.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "Value":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Value != nil {
			value := *typedInput.Properties.Value
			policy.Value = &value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Policy_Spec{}

// ConvertSpecFrom populates our Policy_Spec from the provided source
func (policy *Policy_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.Policy_Spec)
	if ok {
		// Populate our instance from source
		return policy.AssignProperties_From_Policy_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.Policy_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = policy.AssignProperties_From_Policy_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Policy_Spec
func (policy *Policy_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.Policy_Spec)
	if ok {
		// Populate destination from our instance
		return policy.AssignProperties_To_Policy_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Policy_Spec{}
	err := policy.AssignProperties_To_Policy_Spec(dst)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_Policy_Spec populates our Policy_Spec from the provided source Policy_Spec
func (policy *Policy_Spec) AssignProperties_From_Policy_Spec(source *storage.Policy_Spec) error {

	// Format
	if source.Format != nil {
		format := *source.Format
		formatTemp := genruntime.ToEnum(format, policyContractProperties_Format_Values)
		policy.Format = &formatTemp
	} else {
		policy.Format = nil
	}

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec PolicyOperatorSpec
		err := operatorSpec.AssignProperties_From_PolicyOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_PolicyOperatorSpec() to populate field OperatorSpec")
		}
		policy.OperatorSpec = &operatorSpec
	} else {
		policy.OperatorSpec = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		policy.Owner = &owner
	} else {
		policy.Owner = nil
	}

	// Value
	policy.Value = genruntime.ClonePointerToString(source.Value)

	// No error
	return nil
}

// AssignProperties_To_Policy_Spec populates the provided destination Policy_Spec from our Policy_Spec
func (policy *Policy_Spec) AssignProperties_To_Policy_Spec(destination *storage.Policy_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Format
	if policy.Format != nil {
		format := string(*policy.Format)
		destination.Format = &format
	} else {
		destination.Format = nil
	}

	// OperatorSpec
	if policy.OperatorSpec != nil {
		var operatorSpec storage.PolicyOperatorSpec
		err := policy.OperatorSpec.AssignProperties_To_PolicyOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_PolicyOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = policy.OriginalVersion()

	// Owner
	if policy.Owner != nil {
		owner := policy.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Value
	destination.Value = genruntime.ClonePointerToString(policy.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_Policy_STATUS populates our Policy_Spec from the provided source Policy_STATUS
func (policy *Policy_Spec) Initialize_From_Policy_STATUS(source *Policy_STATUS) error {

	// Format
	if source.Format != nil {
		format := genruntime.ToEnum(string(*source.Format), policyContractProperties_Format_Values)
		policy.Format = &format
	} else {
		policy.Format = nil
	}

	// Value
	policy.Value = genruntime.ClonePointerToString(source.Value)

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (policy *Policy_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

type Policy_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Format: Format of the policyContent.
	Format *PolicyContractProperties_Format_STATUS `json:"format,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`

	// Value: Contents of the Policy as defined by the format.
	Value *string `json:"value,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Policy_STATUS{}

// ConvertStatusFrom populates our Policy_STATUS from the provided source
func (policy *Policy_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.Policy_STATUS)
	if ok {
		// Populate our instance from source
		return policy.AssignProperties_From_Policy_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.Policy_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = policy.AssignProperties_From_Policy_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Policy_STATUS
func (policy *Policy_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.Policy_STATUS)
	if ok {
		// Populate destination from our instance
		return policy.AssignProperties_To_Policy_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Policy_STATUS{}
	err := policy.AssignProperties_To_Policy_STATUS(dst)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &Policy_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (policy *Policy_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.Policy_STATUS{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (policy *Policy_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.Policy_STATUS)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.Policy_STATUS, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "Format":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Format != nil {
			var temp string
			temp = string(*typedInput.Properties.Format)
			format := PolicyContractProperties_Format_STATUS(temp)
			policy.Format = &format
		}
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		policy.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		policy.Name = &name
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		policy.Type = &typeVar
	}

	// Set property "Value":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Value != nil {
			value := *typedInput.Properties.Value
			policy.Value = &value
		}
	}

	// No error
	return nil
}

// AssignProperties_From_Policy_STATUS populates our Policy_STATUS from the provided source Policy_STATUS
func (policy *Policy_STATUS) AssignProperties_From_Policy_STATUS(source *storage.Policy_STATUS) error {

	// Conditions
	policy.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Format
	if source.Format != nil {
		format := *source.Format
		formatTemp := genruntime.ToEnum(format, policyContractProperties_Format_STATUS_Values)
		policy.Format = &formatTemp
	} else {
		policy.Format = nil
	}

	// Id
	policy.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	policy.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	policy.Type = genruntime.ClonePointerToString(source.Type)

	// Value
	policy.Value = genruntime.ClonePointerToString(source.Value)

	// No error
	return nil
}

// AssignProperties_To_Policy_STATUS populates the provided destination Policy_STATUS from our Policy_STATUS
func (policy *Policy_STATUS) AssignProperties_To_Policy_STATUS(destination *storage.Policy_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(policy.Conditions)

	// Format
	if policy.Format != nil {
		format := string(*policy.Format)
		destination.Format = &format
	} else {
		destination.Format = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(policy.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(policy.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(policy.Type)

	// Value
	destination.Value = genruntime.ClonePointerToString(policy.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"rawxml","rawxml-link","xml","xml-link"}
type PolicyContractProperties_Format string

const (
	PolicyContractProperties_Format_Rawxml     = PolicyContractProperties_Format("rawxml")
	PolicyContractProperties_Format_RawxmlLink = PolicyContractProperties_Format("rawxml-link")
	PolicyContractProperties_Format_Xml        = PolicyContractProperties_Format("xml")
	PolicyContractProperties_Format_XmlLink    = PolicyContractProperties_Format("xml-link")
)

// Mapping from string to PolicyContractProperties_Format
var policyContractProperties_Format_Values = map[string]PolicyContractProperties_Format{
	"rawxml":      PolicyContractProperties_Format_Rawxml,
	"rawxml-link": PolicyContractProperties_Format_RawxmlLink,
	"xml":         PolicyContractProperties_Format_Xml,
	"xml-link":    PolicyContractProperties_Format_XmlLink,
}

type PolicyContractProperties_Format_STATUS string

const (
	PolicyContractProperties_Format_STATUS_Rawxml     = PolicyContractProperties_Format_STATUS("rawxml")
	PolicyContractProperties_Format_STATUS_RawxmlLink = PolicyContractProperties_Format_STATUS("rawxml-link")
	PolicyContractProperties_Format_STATUS_Xml        = PolicyContractProperties_Format_STATUS("xml")
	PolicyContractProperties_Format_STATUS_XmlLink    = PolicyContractProperties_Format_STATUS("xml-link")
)

// Mapping from string to PolicyContractProperties_Format_STATUS
var policyContractProperties_Format_STATUS_Values = map[string]PolicyContractProperties_Format_STATUS{
	"rawxml":      PolicyContractProperties_Format_STATUS_Rawxml,
	"rawxml-link": PolicyContractProperties_Format_STATUS_RawxmlLink,
	"xml":         PolicyContractProperties_Format_STATUS_Xml,
	"xml-link":    PolicyContractProperties_Format_STATUS_XmlLink,
}

// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type PolicyOperatorSpec struct {
	// ConfigMapExpressions: configures where to place operator written dynamic ConfigMaps (created with CEL expressions).
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`

	// SecretExpressions: configures where to place operator written dynamic secrets (created with CEL expressions).
	SecretExpressions []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_PolicyOperatorSpec populates our PolicyOperatorSpec from the provided source PolicyOperatorSpec
func (operator *PolicyOperatorSpec) AssignProperties_From_PolicyOperatorSpec(source *storage.PolicyOperatorSpec) error {

	// ConfigMapExpressions
	if source.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(source.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range source.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		operator.ConfigMapExpressions = configMapExpressionList
	} else {
		operator.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if source.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(source.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range source.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		operator.SecretExpressions = secretExpressionList
	} else {
		operator.SecretExpressions = nil
	}

	// No error
	return nil
}

// AssignProperties_To_PolicyOperatorSpec populates the provided destination PolicyOperatorSpec from our PolicyOperatorSpec
func (operator *PolicyOperatorSpec) AssignProperties_To_PolicyOperatorSpec(destination *storage.PolicyOperatorSpec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// ConfigMapExpressions
	if operator.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(operator.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range operator.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		destination.ConfigMapExpressions = configMapExpressionList
	} else {
		destination.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if operator.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(operator.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range operator.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		destination.SecretExpressions = secretExpressionList
	} else {
		destination.SecretExpressions = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}
