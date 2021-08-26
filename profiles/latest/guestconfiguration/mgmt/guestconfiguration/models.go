//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package guestconfiguration

import original "github.com/Azure/azure-sdk-for-go/services/guestconfiguration/mgmt/2021-01-25/guestconfiguration"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ActionAfterReboot = original.ActionAfterReboot

const (
	ContinueConfiguration ActionAfterReboot = original.ContinueConfiguration
	StopConfiguration     ActionAfterReboot = original.StopConfiguration
)

type ComplianceStatus = original.ComplianceStatus

const (
	Compliant    ComplianceStatus = original.Compliant
	NonCompliant ComplianceStatus = original.NonCompliant
	Pending      ComplianceStatus = original.Pending
)

type ConfigurationMode = original.ConfigurationMode

const (
	ApplyAndAutoCorrect ConfigurationMode = original.ApplyAndAutoCorrect
	ApplyAndMonitor     ConfigurationMode = original.ApplyAndMonitor
	ApplyOnly           ConfigurationMode = original.ApplyOnly
)

type CreatedByType = original.CreatedByType

const (
	CreatedByTypeApplication     CreatedByType = original.CreatedByTypeApplication
	CreatedByTypeKey             CreatedByType = original.CreatedByTypeKey
	CreatedByTypeManagedIdentity CreatedByType = original.CreatedByTypeManagedIdentity
	CreatedByTypeUser            CreatedByType = original.CreatedByTypeUser
)

type Kind = original.Kind

const (
	DSC Kind = original.DSC
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled  ProvisioningState = original.Canceled
	Created   ProvisioningState = original.Created
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
)

type Type = original.Type

const (
	Consistency Type = original.Consistency
	Initial     Type = original.Initial
)

type Assignment = original.Assignment
type AssignmentInfo = original.AssignmentInfo
type AssignmentList = original.AssignmentList
type AssignmentProperties = original.AssignmentProperties
type AssignmentReport = original.AssignmentReport
type AssignmentReportDetails = original.AssignmentReportDetails
type AssignmentReportList = original.AssignmentReportList
type AssignmentReportProperties = original.AssignmentReportProperties
type AssignmentReportResource = original.AssignmentReportResource
type AssignmentReportResourceComplianceReason = original.AssignmentReportResourceComplianceReason
type AssignmentReportType = original.AssignmentReportType
type AssignmentReportsClient = original.AssignmentReportsClient
type AssignmentReportsVMSSClient = original.AssignmentReportsVMSSClient
type AssignmentsClient = original.AssignmentsClient
type AssignmentsVMSSClient = original.AssignmentsVMSSClient
type BaseClient = original.BaseClient
type ConfigurationInfo = original.ConfigurationInfo
type ConfigurationParameter = original.ConfigurationParameter
type ConfigurationSetting = original.ConfigurationSetting
type ErrorResponse = original.ErrorResponse
type ErrorResponseError = original.ErrorResponseError
type HCRPAssignmentReportsClient = original.HCRPAssignmentReportsClient
type HCRPAssignmentsClient = original.HCRPAssignmentsClient
type Navigation = original.Navigation
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationProperties = original.OperationProperties
type OperationsClient = original.OperationsClient
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type SystemData = original.SystemData
type TrackedResource = original.TrackedResource
type VMInfo = original.VMInfo
type VMSSVMInfo = original.VMSSVMInfo

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAssignmentReportsClient(subscriptionID string) AssignmentReportsClient {
	return original.NewAssignmentReportsClient(subscriptionID)
}
func NewAssignmentReportsClientWithBaseURI(baseURI string, subscriptionID string) AssignmentReportsClient {
	return original.NewAssignmentReportsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAssignmentReportsVMSSClient(subscriptionID string) AssignmentReportsVMSSClient {
	return original.NewAssignmentReportsVMSSClient(subscriptionID)
}
func NewAssignmentReportsVMSSClientWithBaseURI(baseURI string, subscriptionID string) AssignmentReportsVMSSClient {
	return original.NewAssignmentReportsVMSSClientWithBaseURI(baseURI, subscriptionID)
}
func NewAssignmentsClient(subscriptionID string) AssignmentsClient {
	return original.NewAssignmentsClient(subscriptionID)
}
func NewAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) AssignmentsClient {
	return original.NewAssignmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAssignmentsVMSSClient(subscriptionID string) AssignmentsVMSSClient {
	return original.NewAssignmentsVMSSClient(subscriptionID)
}
func NewAssignmentsVMSSClientWithBaseURI(baseURI string, subscriptionID string) AssignmentsVMSSClient {
	return original.NewAssignmentsVMSSClientWithBaseURI(baseURI, subscriptionID)
}
func NewHCRPAssignmentReportsClient(subscriptionID string) HCRPAssignmentReportsClient {
	return original.NewHCRPAssignmentReportsClient(subscriptionID)
}
func NewHCRPAssignmentReportsClientWithBaseURI(baseURI string, subscriptionID string) HCRPAssignmentReportsClient {
	return original.NewHCRPAssignmentReportsClientWithBaseURI(baseURI, subscriptionID)
}
func NewHCRPAssignmentsClient(subscriptionID string) HCRPAssignmentsClient {
	return original.NewHCRPAssignmentsClient(subscriptionID)
}
func NewHCRPAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) HCRPAssignmentsClient {
	return original.NewHCRPAssignmentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleActionAfterRebootValues() []ActionAfterReboot {
	return original.PossibleActionAfterRebootValues()
}
func PossibleComplianceStatusValues() []ComplianceStatus {
	return original.PossibleComplianceStatusValues()
}
func PossibleConfigurationModeValues() []ConfigurationMode {
	return original.PossibleConfigurationModeValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
