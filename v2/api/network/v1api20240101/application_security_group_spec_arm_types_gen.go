// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type ApplicationSecurityGroup_Spec_ARM struct {
	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &ApplicationSecurityGroup_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2024-01-01"
func (group ApplicationSecurityGroup_Spec_ARM) GetAPIVersion() string {
	return "2024-01-01"
}

// GetName returns the Name of the resource
func (group *ApplicationSecurityGroup_Spec_ARM) GetName() string {
	return group.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/applicationSecurityGroups"
func (group *ApplicationSecurityGroup_Spec_ARM) GetType() string {
	return "Microsoft.Network/applicationSecurityGroups"
}
