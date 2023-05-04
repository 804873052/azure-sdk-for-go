//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armbillingbenefits

const (
	moduleName    = "armbillingbenefits"
	moduleVersion = "v2.0.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// AppliedScopeType - Type of the Applied Scope.
type AppliedScopeType string

const (
	AppliedScopeTypeManagementGroup AppliedScopeType = "ManagementGroup"
	AppliedScopeTypeShared          AppliedScopeType = "Shared"
	AppliedScopeTypeSingle          AppliedScopeType = "Single"
)

// PossibleAppliedScopeTypeValues returns the possible values for the AppliedScopeType const type.
func PossibleAppliedScopeTypeValues() []AppliedScopeType {
	return []AppliedScopeType{
		AppliedScopeTypeManagementGroup,
		AppliedScopeTypeShared,
		AppliedScopeTypeSingle,
	}
}

// BillingPlan - Represents the billing plan in ISO 8601 format. Required only for monthly billing plans.
type BillingPlan string

const (
	BillingPlanP1M BillingPlan = "P1M"
)

// PossibleBillingPlanValues returns the possible values for the BillingPlan const type.
func PossibleBillingPlanValues() []BillingPlan {
	return []BillingPlan{
		BillingPlanP1M,
	}
}

// CommitmentGrain - Commitment grain.
type CommitmentGrain string

const (
	CommitmentGrainHourly CommitmentGrain = "Hourly"
)

// PossibleCommitmentGrainValues returns the possible values for the CommitmentGrain const type.
func PossibleCommitmentGrainValues() []CommitmentGrain {
	return []CommitmentGrain{
		CommitmentGrainHourly,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// InstanceFlexibility - Turning this on will apply the reservation discount to other VMs in the same VM size group.
type InstanceFlexibility string

const (
	InstanceFlexibilityOff InstanceFlexibility = "Off"
	InstanceFlexibilityOn  InstanceFlexibility = "On"
)

// PossibleInstanceFlexibilityValues returns the possible values for the InstanceFlexibility const type.
func PossibleInstanceFlexibilityValues() []InstanceFlexibility {
	return []InstanceFlexibility{
		InstanceFlexibilityOff,
		InstanceFlexibilityOn,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// PaymentStatus - Describes whether the payment is completed, failed, cancelled or scheduled in the future.
type PaymentStatus string

const (
	PaymentStatusCancelled PaymentStatus = "Cancelled"
	PaymentStatusFailed    PaymentStatus = "Failed"
	PaymentStatusScheduled PaymentStatus = "Scheduled"
	PaymentStatusSucceeded PaymentStatus = "Succeeded"
)

// PossiblePaymentStatusValues returns the possible values for the PaymentStatus const type.
func PossiblePaymentStatusValues() []PaymentStatus {
	return []PaymentStatus{
		PaymentStatusCancelled,
		PaymentStatusFailed,
		PaymentStatusScheduled,
		PaymentStatusSucceeded,
	}
}

// ProvisioningState - Provisioning state
type ProvisioningState string

const (
	ProvisioningStateCancelled        ProvisioningState = "Cancelled"
	ProvisioningStateConfirmedBilling ProvisioningState = "ConfirmedBilling"
	ProvisioningStateCreated          ProvisioningState = "Created"
	ProvisioningStateCreating         ProvisioningState = "Creating"
	ProvisioningStateExpired          ProvisioningState = "Expired"
	ProvisioningStateFailed           ProvisioningState = "Failed"
	ProvisioningStatePendingBilling   ProvisioningState = "PendingBilling"
	ProvisioningStateSucceeded        ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCancelled,
		ProvisioningStateConfirmedBilling,
		ProvisioningStateCreated,
		ProvisioningStateCreating,
		ProvisioningStateExpired,
		ProvisioningStateFailed,
		ProvisioningStatePendingBilling,
		ProvisioningStateSucceeded,
	}
}

// ReservedResourceType - The type of the resource that is being reserved.
type ReservedResourceType string

const (
	ReservedResourceTypeAVS                    ReservedResourceType = "AVS"
	ReservedResourceTypeAppService             ReservedResourceType = "AppService"
	ReservedResourceTypeAzureDataExplorer      ReservedResourceType = "AzureDataExplorer"
	ReservedResourceTypeAzureFiles             ReservedResourceType = "AzureFiles"
	ReservedResourceTypeBlockBlob              ReservedResourceType = "BlockBlob"
	ReservedResourceTypeCosmosDb               ReservedResourceType = "CosmosDb"
	ReservedResourceTypeDataFactory            ReservedResourceType = "DataFactory"
	ReservedResourceTypeDatabricks             ReservedResourceType = "Databricks"
	ReservedResourceTypeDedicatedHost          ReservedResourceType = "DedicatedHost"
	ReservedResourceTypeManagedDisk            ReservedResourceType = "ManagedDisk"
	ReservedResourceTypeMariaDb                ReservedResourceType = "MariaDb"
	ReservedResourceTypeMySQL                  ReservedResourceType = "MySql"
	ReservedResourceTypeNetAppStorage          ReservedResourceType = "NetAppStorage"
	ReservedResourceTypePostgreSQL             ReservedResourceType = "PostgreSql"
	ReservedResourceTypeRedHat                 ReservedResourceType = "RedHat"
	ReservedResourceTypeRedHatOsa              ReservedResourceType = "RedHatOsa"
	ReservedResourceTypeRedisCache             ReservedResourceType = "RedisCache"
	ReservedResourceTypeSQLAzureHybridBenefit  ReservedResourceType = "SqlAzureHybridBenefit"
	ReservedResourceTypeSQLDataWarehouse       ReservedResourceType = "SqlDataWarehouse"
	ReservedResourceTypeSQLDatabases           ReservedResourceType = "SqlDatabases"
	ReservedResourceTypeSQLEdge                ReservedResourceType = "SqlEdge"
	ReservedResourceTypeSapHana                ReservedResourceType = "SapHana"
	ReservedResourceTypeSuseLinux              ReservedResourceType = "SuseLinux"
	ReservedResourceTypeVMwareCloudSimple      ReservedResourceType = "VMwareCloudSimple"
	ReservedResourceTypeVirtualMachineSoftware ReservedResourceType = "VirtualMachineSoftware"
	ReservedResourceTypeVirtualMachines        ReservedResourceType = "VirtualMachines"
)

// PossibleReservedResourceTypeValues returns the possible values for the ReservedResourceType const type.
func PossibleReservedResourceTypeValues() []ReservedResourceType {
	return []ReservedResourceType{
		ReservedResourceTypeAVS,
		ReservedResourceTypeAppService,
		ReservedResourceTypeAzureDataExplorer,
		ReservedResourceTypeAzureFiles,
		ReservedResourceTypeBlockBlob,
		ReservedResourceTypeCosmosDb,
		ReservedResourceTypeDataFactory,
		ReservedResourceTypeDatabricks,
		ReservedResourceTypeDedicatedHost,
		ReservedResourceTypeManagedDisk,
		ReservedResourceTypeMariaDb,
		ReservedResourceTypeMySQL,
		ReservedResourceTypeNetAppStorage,
		ReservedResourceTypePostgreSQL,
		ReservedResourceTypeRedHat,
		ReservedResourceTypeRedHatOsa,
		ReservedResourceTypeRedisCache,
		ReservedResourceTypeSQLAzureHybridBenefit,
		ReservedResourceTypeSQLDataWarehouse,
		ReservedResourceTypeSQLDatabases,
		ReservedResourceTypeSQLEdge,
		ReservedResourceTypeSapHana,
		ReservedResourceTypeSuseLinux,
		ReservedResourceTypeVMwareCloudSimple,
		ReservedResourceTypeVirtualMachineSoftware,
		ReservedResourceTypeVirtualMachines,
	}
}

// Term - Represent benefit term in ISO 8601 format.
type Term string

const (
	TermP1Y Term = "P1Y"
	TermP3Y Term = "P3Y"
	TermP5Y Term = "P5Y"
)

// PossibleTermValues returns the possible values for the Term const type.
func PossibleTermValues() []Term {
	return []Term{
		TermP1Y,
		TermP3Y,
		TermP5Y,
	}
}
