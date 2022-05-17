//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

const (
	moduleName    = "armrecoveryservicessiterecovery"
	moduleVersion = "v1.0.0"
)

// A2ARecoveryAvailabilityType - The recovery availability type of the virtual machine.
type A2ARecoveryAvailabilityType string

const (
	A2ARecoveryAvailabilityTypeAvailabilitySet  A2ARecoveryAvailabilityType = "AvailabilitySet"
	A2ARecoveryAvailabilityTypeAvailabilityZone A2ARecoveryAvailabilityType = "AvailabilityZone"
	A2ARecoveryAvailabilityTypeSingle           A2ARecoveryAvailabilityType = "Single"
)

// PossibleA2ARecoveryAvailabilityTypeValues returns the possible values for the A2ARecoveryAvailabilityType const type.
func PossibleA2ARecoveryAvailabilityTypeValues() []A2ARecoveryAvailabilityType {
	return []A2ARecoveryAvailabilityType{
		A2ARecoveryAvailabilityTypeAvailabilitySet,
		A2ARecoveryAvailabilityTypeAvailabilityZone,
		A2ARecoveryAvailabilityTypeSingle,
	}
}

// A2ARpRecoveryPointType - The recovery point type.
type A2ARpRecoveryPointType string

const (
	A2ARpRecoveryPointTypeLatest                      A2ARpRecoveryPointType = "Latest"
	A2ARpRecoveryPointTypeLatestApplicationConsistent A2ARpRecoveryPointType = "LatestApplicationConsistent"
	A2ARpRecoveryPointTypeLatestCrashConsistent       A2ARpRecoveryPointType = "LatestCrashConsistent"
	A2ARpRecoveryPointTypeLatestProcessed             A2ARpRecoveryPointType = "LatestProcessed"
)

// PossibleA2ARpRecoveryPointTypeValues returns the possible values for the A2ARpRecoveryPointType const type.
func PossibleA2ARpRecoveryPointTypeValues() []A2ARpRecoveryPointType {
	return []A2ARpRecoveryPointType{
		A2ARpRecoveryPointTypeLatest,
		A2ARpRecoveryPointTypeLatestApplicationConsistent,
		A2ARpRecoveryPointTypeLatestCrashConsistent,
		A2ARpRecoveryPointTypeLatestProcessed,
	}
}

// AgentAutoUpdateStatus - A value indicating whether the auto update is enabled.
type AgentAutoUpdateStatus string

const (
	AgentAutoUpdateStatusDisabled AgentAutoUpdateStatus = "Disabled"
	AgentAutoUpdateStatusEnabled  AgentAutoUpdateStatus = "Enabled"
)

// PossibleAgentAutoUpdateStatusValues returns the possible values for the AgentAutoUpdateStatus const type.
func PossibleAgentAutoUpdateStatusValues() []AgentAutoUpdateStatus {
	return []AgentAutoUpdateStatus{
		AgentAutoUpdateStatusDisabled,
		AgentAutoUpdateStatusEnabled,
	}
}

type AgentUpgradeBlockedReason string

const (
	AgentUpgradeBlockedReasonAgentNoHeartbeat              AgentUpgradeBlockedReason = "AgentNoHeartbeat"
	AgentUpgradeBlockedReasonAlreadyOnLatestVersion        AgentUpgradeBlockedReason = "AlreadyOnLatestVersion"
	AgentUpgradeBlockedReasonDistroIsNotReported           AgentUpgradeBlockedReason = "DistroIsNotReported"
	AgentUpgradeBlockedReasonDistroNotSupportedForUpgrade  AgentUpgradeBlockedReason = "DistroNotSupportedForUpgrade"
	AgentUpgradeBlockedReasonIncompatibleApplianceVersion  AgentUpgradeBlockedReason = "IncompatibleApplianceVersion"
	AgentUpgradeBlockedReasonInvalidAgentVersion           AgentUpgradeBlockedReason = "InvalidAgentVersion"
	AgentUpgradeBlockedReasonInvalidDriverVersion          AgentUpgradeBlockedReason = "InvalidDriverVersion"
	AgentUpgradeBlockedReasonMissingUpgradePath            AgentUpgradeBlockedReason = "MissingUpgradePath"
	AgentUpgradeBlockedReasonNotProtected                  AgentUpgradeBlockedReason = "NotProtected"
	AgentUpgradeBlockedReasonProcessServerNoHeartbeat      AgentUpgradeBlockedReason = "ProcessServerNoHeartbeat"
	AgentUpgradeBlockedReasonRcmProxyNoHeartbeat           AgentUpgradeBlockedReason = "RcmProxyNoHeartbeat"
	AgentUpgradeBlockedReasonRebootRequired                AgentUpgradeBlockedReason = "RebootRequired"
	AgentUpgradeBlockedReasonUnknown                       AgentUpgradeBlockedReason = "Unknown"
	AgentUpgradeBlockedReasonUnsupportedProtectionScenario AgentUpgradeBlockedReason = "UnsupportedProtectionScenario"
)

// PossibleAgentUpgradeBlockedReasonValues returns the possible values for the AgentUpgradeBlockedReason const type.
func PossibleAgentUpgradeBlockedReasonValues() []AgentUpgradeBlockedReason {
	return []AgentUpgradeBlockedReason{
		AgentUpgradeBlockedReasonAgentNoHeartbeat,
		AgentUpgradeBlockedReasonAlreadyOnLatestVersion,
		AgentUpgradeBlockedReasonDistroIsNotReported,
		AgentUpgradeBlockedReasonDistroNotSupportedForUpgrade,
		AgentUpgradeBlockedReasonIncompatibleApplianceVersion,
		AgentUpgradeBlockedReasonInvalidAgentVersion,
		AgentUpgradeBlockedReasonInvalidDriverVersion,
		AgentUpgradeBlockedReasonMissingUpgradePath,
		AgentUpgradeBlockedReasonNotProtected,
		AgentUpgradeBlockedReasonProcessServerNoHeartbeat,
		AgentUpgradeBlockedReasonRcmProxyNoHeartbeat,
		AgentUpgradeBlockedReasonRebootRequired,
		AgentUpgradeBlockedReasonUnknown,
		AgentUpgradeBlockedReasonUnsupportedProtectionScenario,
	}
}

// AgentVersionStatus - A value indicating whether security update required.
type AgentVersionStatus string

const (
	AgentVersionStatusDeprecated             AgentVersionStatus = "Deprecated"
	AgentVersionStatusNotSupported           AgentVersionStatus = "NotSupported"
	AgentVersionStatusSecurityUpdateRequired AgentVersionStatus = "SecurityUpdateRequired"
	AgentVersionStatusSupported              AgentVersionStatus = "Supported"
	AgentVersionStatusUpdateRequired         AgentVersionStatus = "UpdateRequired"
)

// PossibleAgentVersionStatusValues returns the possible values for the AgentVersionStatus const type.
func PossibleAgentVersionStatusValues() []AgentVersionStatus {
	return []AgentVersionStatus{
		AgentVersionStatusDeprecated,
		AgentVersionStatusNotSupported,
		AgentVersionStatusSecurityUpdateRequired,
		AgentVersionStatusSupported,
		AgentVersionStatusUpdateRequired,
	}
}

// AlternateLocationRecoveryOption - The ALR option.
type AlternateLocationRecoveryOption string

const (
	AlternateLocationRecoveryOptionCreateVMIfNotFound AlternateLocationRecoveryOption = "CreateVmIfNotFound"
	AlternateLocationRecoveryOptionNoAction           AlternateLocationRecoveryOption = "NoAction"
)

// PossibleAlternateLocationRecoveryOptionValues returns the possible values for the AlternateLocationRecoveryOption const type.
func PossibleAlternateLocationRecoveryOptionValues() []AlternateLocationRecoveryOption {
	return []AlternateLocationRecoveryOption{
		AlternateLocationRecoveryOptionCreateVMIfNotFound,
		AlternateLocationRecoveryOptionNoAction,
	}
}

// AutoProtectionOfDataDisk - A value indicating whether the auto protection is enabled.
type AutoProtectionOfDataDisk string

const (
	AutoProtectionOfDataDiskDisabled AutoProtectionOfDataDisk = "Disabled"
	AutoProtectionOfDataDiskEnabled  AutoProtectionOfDataDisk = "Enabled"
)

// PossibleAutoProtectionOfDataDiskValues returns the possible values for the AutoProtectionOfDataDisk const type.
func PossibleAutoProtectionOfDataDiskValues() []AutoProtectionOfDataDisk {
	return []AutoProtectionOfDataDisk{
		AutoProtectionOfDataDiskDisabled,
		AutoProtectionOfDataDiskEnabled,
	}
}

// AutomationAccountAuthenticationType - A value indicating the type authentication to use for automation Account.
type AutomationAccountAuthenticationType string

const (
	AutomationAccountAuthenticationTypeRunAsAccount           AutomationAccountAuthenticationType = "RunAsAccount"
	AutomationAccountAuthenticationTypeSystemAssignedIdentity AutomationAccountAuthenticationType = "SystemAssignedIdentity"
)

// PossibleAutomationAccountAuthenticationTypeValues returns the possible values for the AutomationAccountAuthenticationType const type.
func PossibleAutomationAccountAuthenticationTypeValues() []AutomationAccountAuthenticationType {
	return []AutomationAccountAuthenticationType{
		AutomationAccountAuthenticationTypeRunAsAccount,
		AutomationAccountAuthenticationTypeSystemAssignedIdentity,
	}
}

// DataSyncStatus - The data sync option.
type DataSyncStatus string

const (
	DataSyncStatusForDownTime        DataSyncStatus = "ForDownTime"
	DataSyncStatusForSynchronization DataSyncStatus = "ForSynchronization"
)

// PossibleDataSyncStatusValues returns the possible values for the DataSyncStatus const type.
func PossibleDataSyncStatusValues() []DataSyncStatus {
	return []DataSyncStatus{
		DataSyncStatusForDownTime,
		DataSyncStatusForSynchronization,
	}
}

// DisableProtectionReason - Disable protection reason. It can have values NotSpecified/MigrationComplete.
type DisableProtectionReason string

const (
	DisableProtectionReasonMigrationComplete DisableProtectionReason = "MigrationComplete"
	DisableProtectionReasonNotSpecified      DisableProtectionReason = "NotSpecified"
)

// PossibleDisableProtectionReasonValues returns the possible values for the DisableProtectionReason const type.
func PossibleDisableProtectionReasonValues() []DisableProtectionReason {
	return []DisableProtectionReason{
		DisableProtectionReasonMigrationComplete,
		DisableProtectionReasonNotSpecified,
	}
}

// DiskAccountType - The DiskType.
type DiskAccountType string

const (
	DiskAccountTypePremiumLRS     DiskAccountType = "Premium_LRS"
	DiskAccountTypeStandardLRS    DiskAccountType = "Standard_LRS"
	DiskAccountTypeStandardSSDLRS DiskAccountType = "StandardSSD_LRS"
)

// PossibleDiskAccountTypeValues returns the possible values for the DiskAccountType const type.
func PossibleDiskAccountTypeValues() []DiskAccountType {
	return []DiskAccountType{
		DiskAccountTypePremiumLRS,
		DiskAccountTypeStandardLRS,
		DiskAccountTypeStandardSSDLRS,
	}
}

// DiskReplicationProgressHealth - The progress health.
type DiskReplicationProgressHealth string

const (
	DiskReplicationProgressHealthInProgress   DiskReplicationProgressHealth = "InProgress"
	DiskReplicationProgressHealthNoProgress   DiskReplicationProgressHealth = "NoProgress"
	DiskReplicationProgressHealthNone         DiskReplicationProgressHealth = "None"
	DiskReplicationProgressHealthQueued       DiskReplicationProgressHealth = "Queued"
	DiskReplicationProgressHealthSlowProgress DiskReplicationProgressHealth = "SlowProgress"
)

// PossibleDiskReplicationProgressHealthValues returns the possible values for the DiskReplicationProgressHealth const type.
func PossibleDiskReplicationProgressHealthValues() []DiskReplicationProgressHealth {
	return []DiskReplicationProgressHealth{
		DiskReplicationProgressHealthInProgress,
		DiskReplicationProgressHealthNoProgress,
		DiskReplicationProgressHealthNone,
		DiskReplicationProgressHealthQueued,
		DiskReplicationProgressHealthSlowProgress,
	}
}

// EthernetAddressType - The source IP address type.
type EthernetAddressType string

const (
	EthernetAddressTypeDynamic EthernetAddressType = "Dynamic"
	EthernetAddressTypeStatic  EthernetAddressType = "Static"
)

// PossibleEthernetAddressTypeValues returns the possible values for the EthernetAddressType const type.
func PossibleEthernetAddressTypeValues() []EthernetAddressType {
	return []EthernetAddressType{
		EthernetAddressTypeDynamic,
		EthernetAddressTypeStatic,
	}
}

// ExportJobOutputSerializationType - The output type of the jobs.
type ExportJobOutputSerializationType string

const (
	ExportJobOutputSerializationTypeExcel ExportJobOutputSerializationType = "Excel"
	ExportJobOutputSerializationTypeJSON  ExportJobOutputSerializationType = "Json"
	ExportJobOutputSerializationTypeXML   ExportJobOutputSerializationType = "Xml"
)

// PossibleExportJobOutputSerializationTypeValues returns the possible values for the ExportJobOutputSerializationType const type.
func PossibleExportJobOutputSerializationTypeValues() []ExportJobOutputSerializationType {
	return []ExportJobOutputSerializationType{
		ExportJobOutputSerializationTypeExcel,
		ExportJobOutputSerializationTypeJSON,
		ExportJobOutputSerializationTypeXML,
	}
}

// ExtendedLocationType - The extended location type.
type ExtendedLocationType string

const (
	ExtendedLocationTypeEdgeZone ExtendedLocationType = "EdgeZone"
)

// PossibleExtendedLocationTypeValues returns the possible values for the ExtendedLocationType const type.
func PossibleExtendedLocationTypeValues() []ExtendedLocationType {
	return []ExtendedLocationType{
		ExtendedLocationTypeEdgeZone,
	}
}

// FailoverDeploymentModel - The failover deployment model.
type FailoverDeploymentModel string

const (
	FailoverDeploymentModelClassic         FailoverDeploymentModel = "Classic"
	FailoverDeploymentModelNotApplicable   FailoverDeploymentModel = "NotApplicable"
	FailoverDeploymentModelResourceManager FailoverDeploymentModel = "ResourceManager"
)

// PossibleFailoverDeploymentModelValues returns the possible values for the FailoverDeploymentModel const type.
func PossibleFailoverDeploymentModelValues() []FailoverDeploymentModel {
	return []FailoverDeploymentModel{
		FailoverDeploymentModelClassic,
		FailoverDeploymentModelNotApplicable,
		FailoverDeploymentModelResourceManager,
	}
}

// HealthErrorCategory - The category of the health error.
type HealthErrorCategory string

const (
	HealthErrorCategoryAgentAutoUpdateArtifactDeleted     HealthErrorCategory = "AgentAutoUpdateArtifactDeleted"
	HealthErrorCategoryAgentAutoUpdateInfra               HealthErrorCategory = "AgentAutoUpdateInfra"
	HealthErrorCategoryAgentAutoUpdateRunAsAccount        HealthErrorCategory = "AgentAutoUpdateRunAsAccount"
	HealthErrorCategoryAgentAutoUpdateRunAsAccountExpired HealthErrorCategory = "AgentAutoUpdateRunAsAccountExpired"
	HealthErrorCategoryAgentAutoUpdateRunAsAccountExpiry  HealthErrorCategory = "AgentAutoUpdateRunAsAccountExpiry"
	HealthErrorCategoryConfiguration                      HealthErrorCategory = "Configuration"
	HealthErrorCategoryFabricInfrastructure               HealthErrorCategory = "FabricInfrastructure"
	HealthErrorCategoryNone                               HealthErrorCategory = "None"
	HealthErrorCategoryReplication                        HealthErrorCategory = "Replication"
	HealthErrorCategoryTestFailover                       HealthErrorCategory = "TestFailover"
	HealthErrorCategoryVersionExpiry                      HealthErrorCategory = "VersionExpiry"
)

// PossibleHealthErrorCategoryValues returns the possible values for the HealthErrorCategory const type.
func PossibleHealthErrorCategoryValues() []HealthErrorCategory {
	return []HealthErrorCategory{
		HealthErrorCategoryAgentAutoUpdateArtifactDeleted,
		HealthErrorCategoryAgentAutoUpdateInfra,
		HealthErrorCategoryAgentAutoUpdateRunAsAccount,
		HealthErrorCategoryAgentAutoUpdateRunAsAccountExpired,
		HealthErrorCategoryAgentAutoUpdateRunAsAccountExpiry,
		HealthErrorCategoryConfiguration,
		HealthErrorCategoryFabricInfrastructure,
		HealthErrorCategoryNone,
		HealthErrorCategoryReplication,
		HealthErrorCategoryTestFailover,
		HealthErrorCategoryVersionExpiry,
	}
}

// HealthErrorCustomerResolvability - Value indicating whether the health error is customer resolvable.
type HealthErrorCustomerResolvability string

const (
	HealthErrorCustomerResolvabilityAllowed    HealthErrorCustomerResolvability = "Allowed"
	HealthErrorCustomerResolvabilityNotAllowed HealthErrorCustomerResolvability = "NotAllowed"
)

// PossibleHealthErrorCustomerResolvabilityValues returns the possible values for the HealthErrorCustomerResolvability const type.
func PossibleHealthErrorCustomerResolvabilityValues() []HealthErrorCustomerResolvability {
	return []HealthErrorCustomerResolvability{
		HealthErrorCustomerResolvabilityAllowed,
		HealthErrorCustomerResolvabilityNotAllowed,
	}
}

// HyperVReplicaAzureRpRecoveryPointType - The recovery point type.
type HyperVReplicaAzureRpRecoveryPointType string

const (
	HyperVReplicaAzureRpRecoveryPointTypeLatest                      HyperVReplicaAzureRpRecoveryPointType = "Latest"
	HyperVReplicaAzureRpRecoveryPointTypeLatestApplicationConsistent HyperVReplicaAzureRpRecoveryPointType = "LatestApplicationConsistent"
	HyperVReplicaAzureRpRecoveryPointTypeLatestProcessed             HyperVReplicaAzureRpRecoveryPointType = "LatestProcessed"
)

// PossibleHyperVReplicaAzureRpRecoveryPointTypeValues returns the possible values for the HyperVReplicaAzureRpRecoveryPointType const type.
func PossibleHyperVReplicaAzureRpRecoveryPointTypeValues() []HyperVReplicaAzureRpRecoveryPointType {
	return []HyperVReplicaAzureRpRecoveryPointType{
		HyperVReplicaAzureRpRecoveryPointTypeLatest,
		HyperVReplicaAzureRpRecoveryPointTypeLatestApplicationConsistent,
		HyperVReplicaAzureRpRecoveryPointTypeLatestProcessed,
	}
}

// InMageRcmFailbackRecoveryPointType - The recovery point type.
type InMageRcmFailbackRecoveryPointType string

const (
	InMageRcmFailbackRecoveryPointTypeApplicationConsistent InMageRcmFailbackRecoveryPointType = "ApplicationConsistent"
	InMageRcmFailbackRecoveryPointTypeCrashConsistent       InMageRcmFailbackRecoveryPointType = "CrashConsistent"
)

// PossibleInMageRcmFailbackRecoveryPointTypeValues returns the possible values for the InMageRcmFailbackRecoveryPointType const type.
func PossibleInMageRcmFailbackRecoveryPointTypeValues() []InMageRcmFailbackRecoveryPointType {
	return []InMageRcmFailbackRecoveryPointType{
		InMageRcmFailbackRecoveryPointTypeApplicationConsistent,
		InMageRcmFailbackRecoveryPointTypeCrashConsistent,
	}
}

// InMageV2RpRecoveryPointType - The recovery point type.
type InMageV2RpRecoveryPointType string

const (
	InMageV2RpRecoveryPointTypeLatest                      InMageV2RpRecoveryPointType = "Latest"
	InMageV2RpRecoveryPointTypeLatestApplicationConsistent InMageV2RpRecoveryPointType = "LatestApplicationConsistent"
	InMageV2RpRecoveryPointTypeLatestCrashConsistent       InMageV2RpRecoveryPointType = "LatestCrashConsistent"
	InMageV2RpRecoveryPointTypeLatestProcessed             InMageV2RpRecoveryPointType = "LatestProcessed"
)

// PossibleInMageV2RpRecoveryPointTypeValues returns the possible values for the InMageV2RpRecoveryPointType const type.
func PossibleInMageV2RpRecoveryPointTypeValues() []InMageV2RpRecoveryPointType {
	return []InMageV2RpRecoveryPointType{
		InMageV2RpRecoveryPointTypeLatest,
		InMageV2RpRecoveryPointTypeLatestApplicationConsistent,
		InMageV2RpRecoveryPointTypeLatestCrashConsistent,
		InMageV2RpRecoveryPointTypeLatestProcessed,
	}
}

// LicenseType - License type.
type LicenseType string

const (
	LicenseTypeNoLicenseType LicenseType = "NoLicenseType"
	LicenseTypeNotSpecified  LicenseType = "NotSpecified"
	LicenseTypeWindowsServer LicenseType = "WindowsServer"
)

// PossibleLicenseTypeValues returns the possible values for the LicenseType const type.
func PossibleLicenseTypeValues() []LicenseType {
	return []LicenseType{
		LicenseTypeNoLicenseType,
		LicenseTypeNotSpecified,
		LicenseTypeWindowsServer,
	}
}

type MigrationItemOperation string

const (
	MigrationItemOperationDisableMigration   MigrationItemOperation = "DisableMigration"
	MigrationItemOperationMigrate            MigrationItemOperation = "Migrate"
	MigrationItemOperationStartResync        MigrationItemOperation = "StartResync"
	MigrationItemOperationTestMigrate        MigrationItemOperation = "TestMigrate"
	MigrationItemOperationTestMigrateCleanup MigrationItemOperation = "TestMigrateCleanup"
)

// PossibleMigrationItemOperationValues returns the possible values for the MigrationItemOperation const type.
func PossibleMigrationItemOperationValues() []MigrationItemOperation {
	return []MigrationItemOperation{
		MigrationItemOperationDisableMigration,
		MigrationItemOperationMigrate,
		MigrationItemOperationStartResync,
		MigrationItemOperationTestMigrate,
		MigrationItemOperationTestMigrateCleanup,
	}
}

// MigrationRecoveryPointType - The recovery point type.
type MigrationRecoveryPointType string

const (
	MigrationRecoveryPointTypeApplicationConsistent MigrationRecoveryPointType = "ApplicationConsistent"
	MigrationRecoveryPointTypeCrashConsistent       MigrationRecoveryPointType = "CrashConsistent"
	MigrationRecoveryPointTypeNotSpecified          MigrationRecoveryPointType = "NotSpecified"
)

// PossibleMigrationRecoveryPointTypeValues returns the possible values for the MigrationRecoveryPointType const type.
func PossibleMigrationRecoveryPointTypeValues() []MigrationRecoveryPointType {
	return []MigrationRecoveryPointType{
		MigrationRecoveryPointTypeApplicationConsistent,
		MigrationRecoveryPointTypeCrashConsistent,
		MigrationRecoveryPointTypeNotSpecified,
	}
}

// MigrationState - The migration status.
type MigrationState string

const (
	MigrationStateDisableMigrationFailed     MigrationState = "DisableMigrationFailed"
	MigrationStateDisableMigrationInProgress MigrationState = "DisableMigrationInProgress"
	MigrationStateEnableMigrationFailed      MigrationState = "EnableMigrationFailed"
	MigrationStateEnableMigrationInProgress  MigrationState = "EnableMigrationInProgress"
	MigrationStateInitialSeedingFailed       MigrationState = "InitialSeedingFailed"
	MigrationStateInitialSeedingInProgress   MigrationState = "InitialSeedingInProgress"
	MigrationStateMigrationFailed            MigrationState = "MigrationFailed"
	MigrationStateMigrationInProgress        MigrationState = "MigrationInProgress"
	MigrationStateMigrationSucceeded         MigrationState = "MigrationSucceeded"
	MigrationStateNone                       MigrationState = "None"
	MigrationStateReplicating                MigrationState = "Replicating"
)

// PossibleMigrationStateValues returns the possible values for the MigrationState const type.
func PossibleMigrationStateValues() []MigrationState {
	return []MigrationState{
		MigrationStateDisableMigrationFailed,
		MigrationStateDisableMigrationInProgress,
		MigrationStateEnableMigrationFailed,
		MigrationStateEnableMigrationInProgress,
		MigrationStateInitialSeedingFailed,
		MigrationStateInitialSeedingInProgress,
		MigrationStateMigrationFailed,
		MigrationStateMigrationInProgress,
		MigrationStateMigrationSucceeded,
		MigrationStateNone,
		MigrationStateReplicating,
	}
}

// MobilityAgentUpgradeState - The agent auto upgrade state.
type MobilityAgentUpgradeState string

const (
	MobilityAgentUpgradeStateCommit    MobilityAgentUpgradeState = "Commit"
	MobilityAgentUpgradeStateCompleted MobilityAgentUpgradeState = "Completed"
	MobilityAgentUpgradeStateNone      MobilityAgentUpgradeState = "None"
	MobilityAgentUpgradeStateStarted   MobilityAgentUpgradeState = "Started"
)

// PossibleMobilityAgentUpgradeStateValues returns the possible values for the MobilityAgentUpgradeState const type.
func PossibleMobilityAgentUpgradeStateValues() []MobilityAgentUpgradeState {
	return []MobilityAgentUpgradeState{
		MobilityAgentUpgradeStateCommit,
		MobilityAgentUpgradeStateCompleted,
		MobilityAgentUpgradeStateNone,
		MobilityAgentUpgradeStateStarted,
	}
}

// MultiVMGroupCreateOption - Whether Multi VM group is auto created or specified by user.
type MultiVMGroupCreateOption string

const (
	MultiVMGroupCreateOptionAutoCreated   MultiVMGroupCreateOption = "AutoCreated"
	MultiVMGroupCreateOptionUserSpecified MultiVMGroupCreateOption = "UserSpecified"
)

// PossibleMultiVMGroupCreateOptionValues returns the possible values for the MultiVMGroupCreateOption const type.
func PossibleMultiVMGroupCreateOptionValues() []MultiVMGroupCreateOption {
	return []MultiVMGroupCreateOption{
		MultiVMGroupCreateOptionAutoCreated,
		MultiVMGroupCreateOptionUserSpecified,
	}
}

// MultiVMSyncPointOption - A value indicating whether multi VM sync enabled VMs should use multi VM sync points for failover.
type MultiVMSyncPointOption string

const (
	MultiVMSyncPointOptionUseMultiVMSyncRecoveryPoint MultiVMSyncPointOption = "UseMultiVmSyncRecoveryPoint"
	MultiVMSyncPointOptionUsePerVMRecoveryPoint       MultiVMSyncPointOption = "UsePerVmRecoveryPoint"
)

// PossibleMultiVMSyncPointOptionValues returns the possible values for the MultiVMSyncPointOption const type.
func PossibleMultiVMSyncPointOptionValues() []MultiVMSyncPointOption {
	return []MultiVMSyncPointOption{
		MultiVMSyncPointOptionUseMultiVMSyncRecoveryPoint,
		MultiVMSyncPointOptionUsePerVMRecoveryPoint,
	}
}

// PlannedFailoverStatus - The last planned failover status.
type PlannedFailoverStatus string

const (
	PlannedFailoverStatusCancelled PlannedFailoverStatus = "Cancelled"
	PlannedFailoverStatusFailed    PlannedFailoverStatus = "Failed"
	PlannedFailoverStatusSucceeded PlannedFailoverStatus = "Succeeded"
	PlannedFailoverStatusUnknown   PlannedFailoverStatus = "Unknown"
)

// PossiblePlannedFailoverStatusValues returns the possible values for the PlannedFailoverStatus const type.
func PossiblePlannedFailoverStatusValues() []PlannedFailoverStatus {
	return []PlannedFailoverStatus{
		PlannedFailoverStatusCancelled,
		PlannedFailoverStatusFailed,
		PlannedFailoverStatusSucceeded,
		PlannedFailoverStatusUnknown,
	}
}

type PossibleOperationsDirections string

const (
	PossibleOperationsDirectionsPrimaryToRecovery PossibleOperationsDirections = "PrimaryToRecovery"
	PossibleOperationsDirectionsRecoveryToPrimary PossibleOperationsDirections = "RecoveryToPrimary"
)

// PossiblePossibleOperationsDirectionsValues returns the possible values for the PossibleOperationsDirections const type.
func PossiblePossibleOperationsDirectionsValues() []PossibleOperationsDirections {
	return []PossibleOperationsDirections{
		PossibleOperationsDirectionsPrimaryToRecovery,
		PossibleOperationsDirectionsRecoveryToPrimary,
	}
}

// PresenceStatus - A value indicating whether the VM has a physical disk attached. String value of SrsDataContract.PresenceStatus
// enum.
type PresenceStatus string

const (
	PresenceStatusNotPresent PresenceStatus = "NotPresent"
	PresenceStatusPresent    PresenceStatus = "Present"
	PresenceStatusUnknown    PresenceStatus = "Unknown"
)

// PossiblePresenceStatusValues returns the possible values for the PresenceStatus const type.
func PossiblePresenceStatusValues() []PresenceStatus {
	return []PresenceStatus{
		PresenceStatusNotPresent,
		PresenceStatusPresent,
		PresenceStatusUnknown,
	}
}

// ProtectionHealth - The health.
type ProtectionHealth string

const (
	ProtectionHealthCritical ProtectionHealth = "Critical"
	ProtectionHealthNone     ProtectionHealth = "None"
	ProtectionHealthNormal   ProtectionHealth = "Normal"
	ProtectionHealthWarning  ProtectionHealth = "Warning"
)

// PossibleProtectionHealthValues returns the possible values for the ProtectionHealth const type.
func PossibleProtectionHealthValues() []ProtectionHealth {
	return []ProtectionHealth{
		ProtectionHealthCritical,
		ProtectionHealthNone,
		ProtectionHealthNormal,
		ProtectionHealthWarning,
	}
}

// RcmComponentStatus - The throughput status.
type RcmComponentStatus string

const (
	RcmComponentStatusCritical RcmComponentStatus = "Critical"
	RcmComponentStatusHealthy  RcmComponentStatus = "Healthy"
	RcmComponentStatusUnknown  RcmComponentStatus = "Unknown"
	RcmComponentStatusWarning  RcmComponentStatus = "Warning"
)

// PossibleRcmComponentStatusValues returns the possible values for the RcmComponentStatus const type.
func PossibleRcmComponentStatusValues() []RcmComponentStatus {
	return []RcmComponentStatus{
		RcmComponentStatusCritical,
		RcmComponentStatusHealthy,
		RcmComponentStatusUnknown,
		RcmComponentStatusWarning,
	}
}

// RecoveryPlanActionLocation - The fabric location.
type RecoveryPlanActionLocation string

const (
	RecoveryPlanActionLocationPrimary  RecoveryPlanActionLocation = "Primary"
	RecoveryPlanActionLocationRecovery RecoveryPlanActionLocation = "Recovery"
)

// PossibleRecoveryPlanActionLocationValues returns the possible values for the RecoveryPlanActionLocation const type.
func PossibleRecoveryPlanActionLocationValues() []RecoveryPlanActionLocation {
	return []RecoveryPlanActionLocation{
		RecoveryPlanActionLocationPrimary,
		RecoveryPlanActionLocationRecovery,
	}
}

// RecoveryPlanGroupType - The group type.
type RecoveryPlanGroupType string

const (
	RecoveryPlanGroupTypeBoot     RecoveryPlanGroupType = "Boot"
	RecoveryPlanGroupTypeFailover RecoveryPlanGroupType = "Failover"
	RecoveryPlanGroupTypeShutdown RecoveryPlanGroupType = "Shutdown"
)

// PossibleRecoveryPlanGroupTypeValues returns the possible values for the RecoveryPlanGroupType const type.
func PossibleRecoveryPlanGroupTypeValues() []RecoveryPlanGroupType {
	return []RecoveryPlanGroupType{
		RecoveryPlanGroupTypeBoot,
		RecoveryPlanGroupTypeFailover,
		RecoveryPlanGroupTypeShutdown,
	}
}

// RecoveryPlanPointType - The recovery point type.
type RecoveryPlanPointType string

const (
	RecoveryPlanPointTypeLatest                      RecoveryPlanPointType = "Latest"
	RecoveryPlanPointTypeLatestApplicationConsistent RecoveryPlanPointType = "LatestApplicationConsistent"
	RecoveryPlanPointTypeLatestCrashConsistent       RecoveryPlanPointType = "LatestCrashConsistent"
	RecoveryPlanPointTypeLatestProcessed             RecoveryPlanPointType = "LatestProcessed"
)

// PossibleRecoveryPlanPointTypeValues returns the possible values for the RecoveryPlanPointType const type.
func PossibleRecoveryPlanPointTypeValues() []RecoveryPlanPointType {
	return []RecoveryPlanPointType{
		RecoveryPlanPointTypeLatest,
		RecoveryPlanPointTypeLatestApplicationConsistent,
		RecoveryPlanPointTypeLatestCrashConsistent,
		RecoveryPlanPointTypeLatestProcessed,
	}
}

// RecoveryPointSyncType - A value indicating whether the recovery point is multi VM consistent.
type RecoveryPointSyncType string

const (
	RecoveryPointSyncTypeMultiVMSyncRecoveryPoint RecoveryPointSyncType = "MultiVmSyncRecoveryPoint"
	RecoveryPointSyncTypePerVMRecoveryPoint       RecoveryPointSyncType = "PerVmRecoveryPoint"
)

// PossibleRecoveryPointSyncTypeValues returns the possible values for the RecoveryPointSyncType const type.
func PossibleRecoveryPointSyncTypeValues() []RecoveryPointSyncType {
	return []RecoveryPointSyncType{
		RecoveryPointSyncTypeMultiVMSyncRecoveryPoint,
		RecoveryPointSyncTypePerVMRecoveryPoint,
	}
}

// RecoveryPointType - The recovery point type. Values from LatestTime, LatestTag or Custom. In the case of custom, the recovery
// point provided by RecoveryPointId will be used. In the other two cases, recovery point id will
// be ignored.
type RecoveryPointType string

const (
	RecoveryPointTypeCustom     RecoveryPointType = "Custom"
	RecoveryPointTypeLatestTag  RecoveryPointType = "LatestTag"
	RecoveryPointTypeLatestTime RecoveryPointType = "LatestTime"
)

// PossibleRecoveryPointTypeValues returns the possible values for the RecoveryPointType const type.
func PossibleRecoveryPointTypeValues() []RecoveryPointType {
	return []RecoveryPointType{
		RecoveryPointTypeCustom,
		RecoveryPointTypeLatestTag,
		RecoveryPointTypeLatestTime,
	}
}

type ReplicationProtectedItemOperation string

const (
	ReplicationProtectedItemOperationCancelFailover      ReplicationProtectedItemOperation = "CancelFailover"
	ReplicationProtectedItemOperationChangePit           ReplicationProtectedItemOperation = "ChangePit"
	ReplicationProtectedItemOperationCommit              ReplicationProtectedItemOperation = "Commit"
	ReplicationProtectedItemOperationCompleteMigration   ReplicationProtectedItemOperation = "CompleteMigration"
	ReplicationProtectedItemOperationDisableProtection   ReplicationProtectedItemOperation = "DisableProtection"
	ReplicationProtectedItemOperationFailback            ReplicationProtectedItemOperation = "Failback"
	ReplicationProtectedItemOperationFinalizeFailback    ReplicationProtectedItemOperation = "FinalizeFailback"
	ReplicationProtectedItemOperationPlannedFailover     ReplicationProtectedItemOperation = "PlannedFailover"
	ReplicationProtectedItemOperationRepairReplication   ReplicationProtectedItemOperation = "RepairReplication"
	ReplicationProtectedItemOperationReverseReplicate    ReplicationProtectedItemOperation = "ReverseReplicate"
	ReplicationProtectedItemOperationSwitchProtection    ReplicationProtectedItemOperation = "SwitchProtection"
	ReplicationProtectedItemOperationTestFailover        ReplicationProtectedItemOperation = "TestFailover"
	ReplicationProtectedItemOperationTestFailoverCleanup ReplicationProtectedItemOperation = "TestFailoverCleanup"
	ReplicationProtectedItemOperationUnplannedFailover   ReplicationProtectedItemOperation = "UnplannedFailover"
)

// PossibleReplicationProtectedItemOperationValues returns the possible values for the ReplicationProtectedItemOperation const type.
func PossibleReplicationProtectedItemOperationValues() []ReplicationProtectedItemOperation {
	return []ReplicationProtectedItemOperation{
		ReplicationProtectedItemOperationCancelFailover,
		ReplicationProtectedItemOperationChangePit,
		ReplicationProtectedItemOperationCommit,
		ReplicationProtectedItemOperationCompleteMigration,
		ReplicationProtectedItemOperationDisableProtection,
		ReplicationProtectedItemOperationFailback,
		ReplicationProtectedItemOperationFinalizeFailback,
		ReplicationProtectedItemOperationPlannedFailover,
		ReplicationProtectedItemOperationRepairReplication,
		ReplicationProtectedItemOperationReverseReplicate,
		ReplicationProtectedItemOperationSwitchProtection,
		ReplicationProtectedItemOperationTestFailover,
		ReplicationProtectedItemOperationTestFailoverCleanup,
		ReplicationProtectedItemOperationUnplannedFailover,
	}
}

// ResyncState - The resync state.
type ResyncState string

const (
	ResyncStateNone                         ResyncState = "None"
	ResyncStatePreparedForResynchronization ResyncState = "PreparedForResynchronization"
	ResyncStateStartedResynchronization     ResyncState = "StartedResynchronization"
)

// PossibleResyncStateValues returns the possible values for the ResyncState const type.
func PossibleResyncStateValues() []ResyncState {
	return []ResyncState{
		ResyncStateNone,
		ResyncStatePreparedForResynchronization,
		ResyncStateStartedResynchronization,
	}
}

// RpInMageRecoveryPointType - The recovery point type.
type RpInMageRecoveryPointType string

const (
	RpInMageRecoveryPointTypeCustom     RpInMageRecoveryPointType = "Custom"
	RpInMageRecoveryPointTypeLatestTag  RpInMageRecoveryPointType = "LatestTag"
	RpInMageRecoveryPointTypeLatestTime RpInMageRecoveryPointType = "LatestTime"
)

// PossibleRpInMageRecoveryPointTypeValues returns the possible values for the RpInMageRecoveryPointType const type.
func PossibleRpInMageRecoveryPointTypeValues() []RpInMageRecoveryPointType {
	return []RpInMageRecoveryPointType{
		RpInMageRecoveryPointTypeCustom,
		RpInMageRecoveryPointTypeLatestTag,
		RpInMageRecoveryPointTypeLatestTime,
	}
}

// SQLServerLicenseType - The SQL Server license type.
type SQLServerLicenseType string

const (
	SQLServerLicenseTypeAHUB          SQLServerLicenseType = "AHUB"
	SQLServerLicenseTypeNoLicenseType SQLServerLicenseType = "NoLicenseType"
	SQLServerLicenseTypeNotSpecified  SQLServerLicenseType = "NotSpecified"
	SQLServerLicenseTypePAYG          SQLServerLicenseType = "PAYG"
)

// PossibleSQLServerLicenseTypeValues returns the possible values for the SQLServerLicenseType const type.
func PossibleSQLServerLicenseTypeValues() []SQLServerLicenseType {
	return []SQLServerLicenseType{
		SQLServerLicenseTypeAHUB,
		SQLServerLicenseTypeNoLicenseType,
		SQLServerLicenseTypeNotSpecified,
		SQLServerLicenseTypePAYG,
	}
}

// SetMultiVMSyncStatus - A value indicating whether multi-VM sync has to be enabled. Value should be 'Enabled' or 'Disabled'.
type SetMultiVMSyncStatus string

const (
	SetMultiVMSyncStatusDisable SetMultiVMSyncStatus = "Disable"
	SetMultiVMSyncStatusEnable  SetMultiVMSyncStatus = "Enable"
)

// PossibleSetMultiVMSyncStatusValues returns the possible values for the SetMultiVMSyncStatus const type.
func PossibleSetMultiVMSyncStatusValues() []SetMultiVMSyncStatus {
	return []SetMultiVMSyncStatus{
		SetMultiVMSyncStatusDisable,
		SetMultiVMSyncStatusEnable,
	}
}

// Severity - Severity of error.
type Severity string

const (
	SeverityError   Severity = "Error"
	SeverityInfo    Severity = "Info"
	SeverityNONE    Severity = "NONE"
	SeverityWarning Severity = "Warning"
)

// PossibleSeverityValues returns the possible values for the Severity const type.
func PossibleSeverityValues() []Severity {
	return []Severity{
		SeverityError,
		SeverityInfo,
		SeverityNONE,
		SeverityWarning,
	}
}

// SourceSiteOperations - A value indicating whether source site operations are required.
type SourceSiteOperations string

const (
	SourceSiteOperationsNotRequired SourceSiteOperations = "NotRequired"
	SourceSiteOperationsRequired    SourceSiteOperations = "Required"
)

// PossibleSourceSiteOperationsValues returns the possible values for the SourceSiteOperations const type.
func PossibleSourceSiteOperationsValues() []SourceSiteOperations {
	return []SourceSiteOperations{
		SourceSiteOperationsNotRequired,
		SourceSiteOperationsRequired,
	}
}

// TestMigrationState - The test migrate state.
type TestMigrationState string

const (
	TestMigrationStateNone                           TestMigrationState = "None"
	TestMigrationStateTestMigrationCleanupInProgress TestMigrationState = "TestMigrationCleanupInProgress"
	TestMigrationStateTestMigrationFailed            TestMigrationState = "TestMigrationFailed"
	TestMigrationStateTestMigrationInProgress        TestMigrationState = "TestMigrationInProgress"
	TestMigrationStateTestMigrationSucceeded         TestMigrationState = "TestMigrationSucceeded"
)

// PossibleTestMigrationStateValues returns the possible values for the TestMigrationState const type.
func PossibleTestMigrationStateValues() []TestMigrationState {
	return []TestMigrationState{
		TestMigrationStateNone,
		TestMigrationStateTestMigrationCleanupInProgress,
		TestMigrationStateTestMigrationFailed,
		TestMigrationStateTestMigrationInProgress,
		TestMigrationStateTestMigrationSucceeded,
	}
}

// VMEncryptionType - The encryption type of the VM.
type VMEncryptionType string

const (
	VMEncryptionTypeNotEncrypted     VMEncryptionType = "NotEncrypted"
	VMEncryptionTypeOnePassEncrypted VMEncryptionType = "OnePassEncrypted"
	VMEncryptionTypeTwoPassEncrypted VMEncryptionType = "TwoPassEncrypted"
)

// PossibleVMEncryptionTypeValues returns the possible values for the VMEncryptionType const type.
func PossibleVMEncryptionTypeValues() []VMEncryptionType {
	return []VMEncryptionType{
		VMEncryptionTypeNotEncrypted,
		VMEncryptionTypeOnePassEncrypted,
		VMEncryptionTypeTwoPassEncrypted,
	}
}

// VMReplicationProgressHealth - The initial replication progress health.
type VMReplicationProgressHealth string

const (
	VMReplicationProgressHealthInProgress   VMReplicationProgressHealth = "InProgress"
	VMReplicationProgressHealthNoProgress   VMReplicationProgressHealth = "NoProgress"
	VMReplicationProgressHealthNone         VMReplicationProgressHealth = "None"
	VMReplicationProgressHealthSlowProgress VMReplicationProgressHealth = "SlowProgress"
)

// PossibleVMReplicationProgressHealthValues returns the possible values for the VMReplicationProgressHealth const type.
func PossibleVMReplicationProgressHealthValues() []VMReplicationProgressHealth {
	return []VMReplicationProgressHealth{
		VMReplicationProgressHealthInProgress,
		VMReplicationProgressHealthNoProgress,
		VMReplicationProgressHealthNone,
		VMReplicationProgressHealthSlowProgress,
	}
}
