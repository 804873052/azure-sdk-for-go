//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armselfhelp

import "time"

// CheckNameAvailabilityRequest - The check availability request body.
type CheckNameAvailabilityRequest struct {
	// The name of the resource for which availability needs to be checked.
	Name *string

	// The resource type.
	Type *string
}

// CheckNameAvailabilityResponse - Response for whether the requested resource name is available or not.
type CheckNameAvailabilityResponse struct {
	// Gets an error message explaining the 'reason' value with more details. This field is returned iif nameAvailable is false.
	Message *string

	// Returns true or false depending on the availability of the name
	NameAvailable *bool

	// Reason for why value is not available. This field is returned if nameAvailable is false.
	Reason *string
}

// Diagnostic - Properties returned with in an insight.
type Diagnostic struct {
	// Error definition.
	Error *Error

	// The problems (if any) detected by this insight.
	Insights []*Insight

	// Solution Id
	SolutionID *string

	// Denotes the status of the diagnostic resource.
	Status *Status
}

// DiagnosticInvocation - Solution Invocation with additional params needed for invocation.
type DiagnosticInvocation struct {
	// Additional parameters required to invoke the solutionId.
	AdditionalParameters map[string]*string

	// Solution Id to invoke.
	SolutionID *string
}

// DiagnosticResource - Diagnostic resource
type DiagnosticResource struct {
	// Diagnostic Resource properties.
	Properties *DiagnosticResourceProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// DiagnosticResourceProperties - Diagnostic resource properties.
type DiagnosticResourceProperties struct {
	// Global parameters that can be passed to all solutionIds.
	GlobalParameters map[string]*string

	// SolutionIds that are needed to be invoked.
	Insights []*DiagnosticInvocation

	// READ-ONLY; Diagnostic Request Accepted time.
	AcceptedAt *string

	// READ-ONLY; Array of Diagnostics.
	Diagnostics []*Diagnostic

	// READ-ONLY; Status of diagnostic provisioning.
	ProvisioningState *ProvisioningState
}

// DiscoveryResponse - Discovery response.
type DiscoveryResponse struct {
	// The link used to get the next page of solution metadata.
	NextLink *string

	// The list of solution metadata.
	Value []*SolutionMetadataResource
}

// Error definition.
type Error struct {
	// An array of additional nested error response info objects, as described by this contract.
	Details []*Error

	// READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *string

	// READ-ONLY; Description of the error.
	Message *string

	// READ-ONLY; Service specific error type which serves as additional context for the error herein.
	Type *string
}

// Insight - Detailed insights(s) obtained via the invocation of an insight diagnostic troubleshooter.
type Insight struct {
	// Article id.
	ID *string

	// Importance level of the insight.
	ImportanceLevel *ImportanceLevel

	// Detailed result content.
	Results *string

	// This insight's title.
	Title *string
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// SolutionMetadataProperties - Diagnostic solution metadata.
type SolutionMetadataProperties struct {
	// A detailed description of solution.
	Description *string

	// Required parameters for invoking this particular solution.
	RequiredParameterSets [][]*string

	// Solution Id.
	SolutionID *string

	// Solution Type.
	SolutionType *string
}

// SolutionMetadataResource - Solution Metadata resource
type SolutionMetadataResource struct {
	// Solution metadata Resource properties.
	Properties *SolutionMetadataProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}
