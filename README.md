# Azure Service Operator (for Kubernetes)
[![Go Report Card](https://goreportcard.com/badge/github.com/Azure/azure-service-operator)](https://goreportcard.com/report/github.com/Azure/azure-service-operator)
<!-- ![v2 Status](https://github.com/azure/azure-service-operator/actions/workflows/live-validation.yml/badge.svg?branch=main) -->

## What is it?
**Azure Service Operator** (ASO) helps you provision Azure resources and connect your applications to them from within Kubernetes.

If you want to use Azure resources but would prefer to manage those resources using Kubernetes tooling and primitives (for example `kubectl apply`), then Azure Service Operator might be for you.

## Overview

The Azure Service Operator consists of:

- The Custom Resource Definitions (CRDs) for each of the Azure services a Kubernetes user can provision.
- The Kubernetes controller that manages the Azure resources represented by the user specified Custom Resources. The controller attempts to synchronize the desired state in the user specified Custom Resource with the actual state of that resource in Azure, creating it if it doesn't exist, updating it if it has been changed, or deleting it.

## Versions of Azure Service Operator
There are two major versions of Azure Service Operator: v1 and v2. Consult the below table and descriptions to learn more about which you should use.

> Note: ASO v1 and v2 are two totally independent operators. Each has its own unique set of CRDs and controllers. They can be deployed side by side in the same cluster.

| ASO Version | Lifecycle stage | Development status        | Installation options                                                                                                                                                                                    |
| ----------- | --------------- |---------------------------| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| v2          | Stable          | Under active development. | [Helm chart](/v2/charts), [GitHub release 2.x](https://github.com/Azure/azure-service-operator/releases). See [installation](https://azure.github.io/azure-service-operator/#installation) for example. |
| v1          | Decprecated     | Halted                    | [Helm chart](/charts), [OperatorHub](https://operatorhub.io/operator/azure-service-operator) or [GitHub release 1.x](https://github.com/Azure/azure-service-operator/releases)                          |

### ASO v2
Azure Service Operator v2 was built based on the lessons learned from ASO v1, with the following improvements:

* Supports code-generated CRDs based on [Azure OpenAPI specifications](https://github.com/Azure/azure-rest-api-specs). This enables us to quickly add new resources as they are requested.
* More powerful `Status`. You can view the actual state of the resource in Azure through ASO v2, which enables you to see server-side applied defaults and more easily debug issues.
* Dedicated storage versions. This enables faster (and less error prone) support for new Azure API versions, even if there were significant changes in resource shape.
* Uniformity. ASO v2 resources are very uniform due to their code-generated nature.
* Clearer resource states. The state a resource is in is exposed via a [Ready condition](https://azure.github.io/azure-service-operator/design/resource-states/).

[Learn more about Azure Service Operator v2](https://azure.github.io/azure-service-operator/)

### ASO v1

> **ASO v1 has been deprecated. New users should use [ASO v2](https://azure.github.io/azure-service-operator/).**

Azure Service Operator v1 is no longer being maintained - this means no bug fixes and no security updates.

See the [ASOv1 to ASOv2 migration guide](https://azure.github.io/azure-service-operator/guide/asov1-asov2-migration/) for migrating from ASOv1 to ASOv2.

[Learn more about Azure Service Operator v1](/docs/v1/README.md)

## Contributing

The [contribution guide](CONTRIBUTING.md) covers everything you need to know about how you can contribute to Azure Service Operators.

## Support and feedback

For help, please use the following resources:

1. Review the [documentation](https://azure.github.io/azure-service-operator/)
2. Search [open issues](https://github.com/Azure/azure-service-operator/issues). If your issue is not represented there already, please [open a new one](https://github.com/Azure/azure-service-operator/issues/new/choose).
3. Chat with us on the `azure-service-operator` channel of the [Kubernetes Slack](https://kubernetes.slack.com/). If you are not a member you can get an invitation from the [community inviter](https://communityinviter.com/apps/kubernetes/community).

For more information, see [SUPPORT.md](SUPPORT.md).

## Code of conduct

This project has adopted the [Microsoft Open Source Code of Conduct](https://opensource.microsoft.com/codeofconduct/). For more information, see the [Code of Conduct FAQ](https://opensource.microsoft.com/codeofconduct/faq) or contact [opencode@microsoft.com](mailto:opencode@microsoft.com) with any additional questions or comments.
