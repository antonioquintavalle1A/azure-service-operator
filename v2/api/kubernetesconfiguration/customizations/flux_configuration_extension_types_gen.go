// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	v20230501 "github.com/Azure/azure-service-operator/v2/api/kubernetesconfiguration/v1api20230501"
	v20230501s "github.com/Azure/azure-service-operator/v2/api/kubernetesconfiguration/v1api20230501/storage"
	v20241101 "github.com/Azure/azure-service-operator/v2/api/kubernetesconfiguration/v1api20241101"
	v20241101s "github.com/Azure/azure-service-operator/v2/api/kubernetesconfiguration/v1api20241101/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type FluxConfigurationExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *FluxConfigurationExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&v20230501.FluxConfiguration{},
		&v20230501s.FluxConfiguration{},
		&v20241101.FluxConfiguration{},
		&v20241101s.FluxConfiguration{}}
}
