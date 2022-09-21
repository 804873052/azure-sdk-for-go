//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/eng/tools/profileBuilder

package operationalinsights

import (
	"context"

	original "github.com/Azure/dev/azure-sdk-for-go/services/operationalinsights/mgmt/2021-06-01/operationalinsights"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BillingType = original.BillingType

const (
	BillingTypeCluster    BillingType = original.BillingTypeCluster
	BillingTypeWorkspaces BillingType = original.BillingTypeWorkspaces
)

type ClusterEntityStatus = original.ClusterEntityStatus

const (
	Canceled            ClusterEntityStatus = original.Canceled
	Creating            ClusterEntityStatus = original.Creating
	Deleting            ClusterEntityStatus = original.Deleting
	Failed              ClusterEntityStatus = original.Failed
	ProvisioningAccount ClusterEntityStatus = original.ProvisioningAccount
	Succeeded           ClusterEntityStatus = original.Succeeded
	Updating            ClusterEntityStatus = original.Updating
)

type ClusterSkuNameEnum = original.ClusterSkuNameEnum

const (
	CapacityReservation ClusterSkuNameEnum = original.CapacityReservation
)

type DataIngestionStatus = original.DataIngestionStatus

const (
	ApproachingQuota      DataIngestionStatus = original.ApproachingQuota
	ForceOff              DataIngestionStatus = original.ForceOff
	ForceOn               DataIngestionStatus = original.ForceOn
	OverQuota             DataIngestionStatus = original.OverQuota
	RespectQuota          DataIngestionStatus = original.RespectQuota
	SubscriptionSuspended DataIngestionStatus = original.SubscriptionSuspended
)

type DataSourceKind = original.DataSourceKind

const (
	ApplicationInsights                                  DataSourceKind = original.ApplicationInsights
	AzureActivityLog                                     DataSourceKind = original.AzureActivityLog
	AzureAuditLog                                        DataSourceKind = original.AzureAuditLog
	ChangeTrackingContentLocation                        DataSourceKind = original.ChangeTrackingContentLocation
	ChangeTrackingCustomPath                             DataSourceKind = original.ChangeTrackingCustomPath
	ChangeTrackingDataTypeConfiguration                  DataSourceKind = original.ChangeTrackingDataTypeConfiguration
	ChangeTrackingDefaultRegistry                        DataSourceKind = original.ChangeTrackingDefaultRegistry
	ChangeTrackingLinuxPath                              DataSourceKind = original.ChangeTrackingLinuxPath
	ChangeTrackingPath                                   DataSourceKind = original.ChangeTrackingPath
	ChangeTrackingRegistry                               DataSourceKind = original.ChangeTrackingRegistry
	ChangeTrackingServices                               DataSourceKind = original.ChangeTrackingServices
	CustomLog                                            DataSourceKind = original.CustomLog
	CustomLogCollection                                  DataSourceKind = original.CustomLogCollection
	DNSAnalytics                                         DataSourceKind = original.DNSAnalytics
	GenericDataSource                                    DataSourceKind = original.GenericDataSource
	IISLogs                                              DataSourceKind = original.IISLogs
	ImportComputerGroup                                  DataSourceKind = original.ImportComputerGroup
	Itsm                                                 DataSourceKind = original.Itsm
	LinuxChangeTrackingPath                              DataSourceKind = original.LinuxChangeTrackingPath
	LinuxPerformanceCollection                           DataSourceKind = original.LinuxPerformanceCollection
	LinuxPerformanceObject                               DataSourceKind = original.LinuxPerformanceObject
	LinuxSyslog                                          DataSourceKind = original.LinuxSyslog
	LinuxSyslogCollection                                DataSourceKind = original.LinuxSyslogCollection
	NetworkMonitoring                                    DataSourceKind = original.NetworkMonitoring
	Office365                                            DataSourceKind = original.Office365
	SecurityCenterSecurityWindowsBaselineConfiguration   DataSourceKind = original.SecurityCenterSecurityWindowsBaselineConfiguration
	SecurityEventCollectionConfiguration                 DataSourceKind = original.SecurityEventCollectionConfiguration
	SecurityInsightsSecurityEventCollectionConfiguration DataSourceKind = original.SecurityInsightsSecurityEventCollectionConfiguration
	SecurityWindowsBaselineConfiguration                 DataSourceKind = original.SecurityWindowsBaselineConfiguration
	SQLDataClassification                                DataSourceKind = original.SQLDataClassification
	WindowsEvent                                         DataSourceKind = original.WindowsEvent
	WindowsPerformanceCounter                            DataSourceKind = original.WindowsPerformanceCounter
	WindowsTelemetry                                     DataSourceKind = original.WindowsTelemetry
)

type DataSourceType = original.DataSourceType

const (
	Alerts      DataSourceType = original.Alerts
	AzureWatson DataSourceType = original.AzureWatson
	CustomLogs  DataSourceType = original.CustomLogs
	Ingestion   DataSourceType = original.Ingestion
	Query       DataSourceType = original.Query
)

type IdentityTypeForCluster = original.IdentityTypeForCluster

const (
	None           IdentityTypeForCluster = original.None
	SystemAssigned IdentityTypeForCluster = original.SystemAssigned
	UserAssigned   IdentityTypeForCluster = original.UserAssigned
)

type LinkedServiceEntityStatus = original.LinkedServiceEntityStatus

const (
	LinkedServiceEntityStatusDeleting            LinkedServiceEntityStatus = original.LinkedServiceEntityStatusDeleting
	LinkedServiceEntityStatusProvisioningAccount LinkedServiceEntityStatus = original.LinkedServiceEntityStatusProvisioningAccount
	LinkedServiceEntityStatusSucceeded           LinkedServiceEntityStatus = original.LinkedServiceEntityStatusSucceeded
	LinkedServiceEntityStatusUpdating            LinkedServiceEntityStatus = original.LinkedServiceEntityStatusUpdating
)

type PublicNetworkAccessType = original.PublicNetworkAccessType

const (
	Disabled PublicNetworkAccessType = original.Disabled
	Enabled  PublicNetworkAccessType = original.Enabled
)

type PurgeState = original.PurgeState

const (
	Completed PurgeState = original.Completed
	Pending   PurgeState = original.Pending
)

type SearchSortEnum = original.SearchSortEnum

const (
	Asc  SearchSortEnum = original.Asc
	Desc SearchSortEnum = original.Desc
)

type SkuNameEnum = original.SkuNameEnum

const (
	SkuNameEnumCapacityReservation SkuNameEnum = original.SkuNameEnumCapacityReservation
	SkuNameEnumFree                SkuNameEnum = original.SkuNameEnumFree
	SkuNameEnumPerGB2018           SkuNameEnum = original.SkuNameEnumPerGB2018
	SkuNameEnumPerNode             SkuNameEnum = original.SkuNameEnumPerNode
	SkuNameEnumPremium             SkuNameEnum = original.SkuNameEnumPremium
	SkuNameEnumStandalone          SkuNameEnum = original.SkuNameEnumStandalone
	SkuNameEnumStandard            SkuNameEnum = original.SkuNameEnumStandard
)

type StorageInsightState = original.StorageInsightState

const (
	ERROR StorageInsightState = original.ERROR
	OK    StorageInsightState = original.OK
)

type Type = original.Type

const (
	TypeEventHub       Type = original.TypeEventHub
	TypeStorageAccount Type = original.TypeStorageAccount
)

type WorkspaceEntityStatus = original.WorkspaceEntityStatus

const (
	WorkspaceEntityStatusCanceled            WorkspaceEntityStatus = original.WorkspaceEntityStatusCanceled
	WorkspaceEntityStatusCreating            WorkspaceEntityStatus = original.WorkspaceEntityStatusCreating
	WorkspaceEntityStatusDeleting            WorkspaceEntityStatus = original.WorkspaceEntityStatusDeleting
	WorkspaceEntityStatusFailed              WorkspaceEntityStatus = original.WorkspaceEntityStatusFailed
	WorkspaceEntityStatusProvisioningAccount WorkspaceEntityStatus = original.WorkspaceEntityStatusProvisioningAccount
	WorkspaceEntityStatusSucceeded           WorkspaceEntityStatus = original.WorkspaceEntityStatusSucceeded
	WorkspaceEntityStatusUpdating            WorkspaceEntityStatus = original.WorkspaceEntityStatusUpdating
)

type WorkspaceSkuNameEnum = original.WorkspaceSkuNameEnum

const (
	WorkspaceSkuNameEnumCapacityReservation WorkspaceSkuNameEnum = original.WorkspaceSkuNameEnumCapacityReservation
	WorkspaceSkuNameEnumFree                WorkspaceSkuNameEnum = original.WorkspaceSkuNameEnumFree
	WorkspaceSkuNameEnumLACluster           WorkspaceSkuNameEnum = original.WorkspaceSkuNameEnumLACluster
	WorkspaceSkuNameEnumPerGB2018           WorkspaceSkuNameEnum = original.WorkspaceSkuNameEnumPerGB2018
	WorkspaceSkuNameEnumPerNode             WorkspaceSkuNameEnum = original.WorkspaceSkuNameEnumPerNode
	WorkspaceSkuNameEnumPremium             WorkspaceSkuNameEnum = original.WorkspaceSkuNameEnumPremium
	WorkspaceSkuNameEnumStandalone          WorkspaceSkuNameEnum = original.WorkspaceSkuNameEnumStandalone
	WorkspaceSkuNameEnumStandard            WorkspaceSkuNameEnum = original.WorkspaceSkuNameEnumStandard
)

type AssociatedWorkspace = original.AssociatedWorkspace
type AvailableServiceTier = original.AvailableServiceTier
type AvailableServiceTiersClient = original.AvailableServiceTiersClient
type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type CapacityReservationProperties = original.CapacityReservationProperties
type Cluster = original.Cluster
type ClusterListResult = original.ClusterListResult
type ClusterListResultIterator = original.ClusterListResultIterator
type ClusterListResultPage = original.ClusterListResultPage
type ClusterPatch = original.ClusterPatch
type ClusterPatchProperties = original.ClusterPatchProperties
type ClusterProperties = original.ClusterProperties
type ClusterSku = original.ClusterSku
type ClustersClient = original.ClustersClient
type ClustersCreateOrUpdateFuture = original.ClustersCreateOrUpdateFuture
type ClustersDeleteFuture = original.ClustersDeleteFuture
type ClustersUpdateFuture = original.ClustersUpdateFuture
type CoreSummary = original.CoreSummary
type DataExport = original.DataExport
type DataExportListResult = original.DataExportListResult
type DataExportProperties = original.DataExportProperties
type DataExportsClient = original.DataExportsClient
type DataSource = original.DataSource
type DataSourceFilter = original.DataSourceFilter
type DataSourceListResult = original.DataSourceListResult
type DataSourceListResultIterator = original.DataSourceListResultIterator
type DataSourceListResultPage = original.DataSourceListResultPage
type DataSourcesClient = original.DataSourcesClient
type DeletedWorkspacesClient = original.DeletedWorkspacesClient
type Destination = original.Destination
type DestinationMetaData = original.DestinationMetaData
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type GatewaysClient = original.GatewaysClient
type Identity = original.Identity
type IntelligencePack = original.IntelligencePack
type IntelligencePacksClient = original.IntelligencePacksClient
type KeyVaultProperties = original.KeyVaultProperties
type LinkedService = original.LinkedService
type LinkedServiceListResult = original.LinkedServiceListResult
type LinkedServiceProperties = original.LinkedServiceProperties
type LinkedServicesClient = original.LinkedServicesClient
type LinkedServicesCreateOrUpdateFuture = original.LinkedServicesCreateOrUpdateFuture
type LinkedServicesDeleteFuture = original.LinkedServicesDeleteFuture
type LinkedStorageAccountsClient = original.LinkedStorageAccountsClient
type LinkedStorageAccountsListResult = original.LinkedStorageAccountsListResult
type LinkedStorageAccountsProperties = original.LinkedStorageAccountsProperties
type LinkedStorageAccountsResource = original.LinkedStorageAccountsResource
type ListAvailableServiceTier = original.ListAvailableServiceTier
type ListIntelligencePack = original.ListIntelligencePack
type ManagementGroup = original.ManagementGroup
type ManagementGroupProperties = original.ManagementGroupProperties
type ManagementGroupsClient = original.ManagementGroupsClient
type MetricName = original.MetricName
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationStatus = original.OperationStatus
type OperationStatusesClient = original.OperationStatusesClient
type OperationsClient = original.OperationsClient
type PrivateLinkScopedResource = original.PrivateLinkScopedResource
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type SavedSearch = original.SavedSearch
type SavedSearchProperties = original.SavedSearchProperties
type SavedSearchesClient = original.SavedSearchesClient
type SavedSearchesListResult = original.SavedSearchesListResult
type SchemaClient = original.SchemaClient
type SearchGetSchemaResponse = original.SearchGetSchemaResponse
type SearchMetadata = original.SearchMetadata
type SearchMetadataSchema = original.SearchMetadataSchema
type SearchSchemaValue = original.SearchSchemaValue
type SearchSort = original.SearchSort
type SharedKeys = original.SharedKeys
type SharedKeysClient = original.SharedKeysClient
type StorageAccount = original.StorageAccount
type StorageInsight = original.StorageInsight
type StorageInsightConfigsClient = original.StorageInsightConfigsClient
type StorageInsightListResult = original.StorageInsightListResult
type StorageInsightListResultIterator = original.StorageInsightListResultIterator
type StorageInsightListResultPage = original.StorageInsightListResultPage
type StorageInsightProperties = original.StorageInsightProperties
type StorageInsightStatus = original.StorageInsightStatus
type Table = original.Table
type TableProperties = original.TableProperties
type TablesClient = original.TablesClient
type TablesListResult = original.TablesListResult
type Tag = original.Tag
type TrackedResource = original.TrackedResource
type UsageMetric = original.UsageMetric
type UsagesClient = original.UsagesClient
type UserIdentityProperties = original.UserIdentityProperties
type Workspace = original.Workspace
type WorkspaceCapping = original.WorkspaceCapping
type WorkspaceFeatures = original.WorkspaceFeatures
type WorkspaceListManagementGroupsResult = original.WorkspaceListManagementGroupsResult
type WorkspaceListResult = original.WorkspaceListResult
type WorkspaceListUsagesResult = original.WorkspaceListUsagesResult
type WorkspacePatch = original.WorkspacePatch
type WorkspaceProperties = original.WorkspaceProperties
type WorkspacePurgeBody = original.WorkspacePurgeBody
type WorkspacePurgeBodyFilters = original.WorkspacePurgeBodyFilters
type WorkspacePurgeClient = original.WorkspacePurgeClient
type WorkspacePurgeResponse = original.WorkspacePurgeResponse
type WorkspacePurgeStatusResponse = original.WorkspacePurgeStatusResponse
type WorkspaceSku = original.WorkspaceSku
type WorkspacesClient = original.WorkspacesClient
type WorkspacesCreateOrUpdateFuture = original.WorkspacesCreateOrUpdateFuture
type WorkspacesDeleteFuture = original.WorkspacesDeleteFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAvailableServiceTiersClient(subscriptionID string) AvailableServiceTiersClient {
	return original.NewAvailableServiceTiersClient(subscriptionID)
}
func NewAvailableServiceTiersClientWithBaseURI(baseURI string, subscriptionID string) AvailableServiceTiersClient {
	return original.NewAvailableServiceTiersClientWithBaseURI(baseURI, subscriptionID)
}
func NewClusterListResultIterator(page ClusterListResultPage) ClusterListResultIterator {
	return original.NewClusterListResultIterator(page)
}
func NewClusterListResultPage(cur ClusterListResult, getNextPage func(context.Context, ClusterListResult) (ClusterListResult, error)) ClusterListResultPage {
	return original.NewClusterListResultPage(cur, getNextPage)
}
func NewClustersClient(subscriptionID string) ClustersClient {
	return original.NewClustersClient(subscriptionID)
}
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return original.NewClustersClientWithBaseURI(baseURI, subscriptionID)
}
func NewDataExportsClient(subscriptionID string) DataExportsClient {
	return original.NewDataExportsClient(subscriptionID)
}
func NewDataExportsClientWithBaseURI(baseURI string, subscriptionID string) DataExportsClient {
	return original.NewDataExportsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDataSourceListResultIterator(page DataSourceListResultPage) DataSourceListResultIterator {
	return original.NewDataSourceListResultIterator(page)
}
func NewDataSourceListResultPage(cur DataSourceListResult, getNextPage func(context.Context, DataSourceListResult) (DataSourceListResult, error)) DataSourceListResultPage {
	return original.NewDataSourceListResultPage(cur, getNextPage)
}
func NewDataSourcesClient(subscriptionID string) DataSourcesClient {
	return original.NewDataSourcesClient(subscriptionID)
}
func NewDataSourcesClientWithBaseURI(baseURI string, subscriptionID string) DataSourcesClient {
	return original.NewDataSourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewDeletedWorkspacesClient(subscriptionID string) DeletedWorkspacesClient {
	return original.NewDeletedWorkspacesClient(subscriptionID)
}
func NewDeletedWorkspacesClientWithBaseURI(baseURI string, subscriptionID string) DeletedWorkspacesClient {
	return original.NewDeletedWorkspacesClientWithBaseURI(baseURI, subscriptionID)
}
func NewGatewaysClient(subscriptionID string) GatewaysClient {
	return original.NewGatewaysClient(subscriptionID)
}
func NewGatewaysClientWithBaseURI(baseURI string, subscriptionID string) GatewaysClient {
	return original.NewGatewaysClientWithBaseURI(baseURI, subscriptionID)
}
func NewIntelligencePacksClient(subscriptionID string) IntelligencePacksClient {
	return original.NewIntelligencePacksClient(subscriptionID)
}
func NewIntelligencePacksClientWithBaseURI(baseURI string, subscriptionID string) IntelligencePacksClient {
	return original.NewIntelligencePacksClientWithBaseURI(baseURI, subscriptionID)
}
func NewLinkedServicesClient(subscriptionID string) LinkedServicesClient {
	return original.NewLinkedServicesClient(subscriptionID)
}
func NewLinkedServicesClientWithBaseURI(baseURI string, subscriptionID string) LinkedServicesClient {
	return original.NewLinkedServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewLinkedStorageAccountsClient(subscriptionID string) LinkedStorageAccountsClient {
	return original.NewLinkedStorageAccountsClient(subscriptionID)
}
func NewLinkedStorageAccountsClientWithBaseURI(baseURI string, subscriptionID string) LinkedStorageAccountsClient {
	return original.NewLinkedStorageAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewManagementGroupsClient(subscriptionID string) ManagementGroupsClient {
	return original.NewManagementGroupsClient(subscriptionID)
}
func NewManagementGroupsClientWithBaseURI(baseURI string, subscriptionID string) ManagementGroupsClient {
	return original.NewManagementGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationStatusesClient(subscriptionID string) OperationStatusesClient {
	return original.NewOperationStatusesClient(subscriptionID)
}
func NewOperationStatusesClientWithBaseURI(baseURI string, subscriptionID string) OperationStatusesClient {
	return original.NewOperationStatusesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSavedSearchesClient(subscriptionID string) SavedSearchesClient {
	return original.NewSavedSearchesClient(subscriptionID)
}
func NewSavedSearchesClientWithBaseURI(baseURI string, subscriptionID string) SavedSearchesClient {
	return original.NewSavedSearchesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSchemaClient(subscriptionID string) SchemaClient {
	return original.NewSchemaClient(subscriptionID)
}
func NewSchemaClientWithBaseURI(baseURI string, subscriptionID string) SchemaClient {
	return original.NewSchemaClientWithBaseURI(baseURI, subscriptionID)
}
func NewSharedKeysClient(subscriptionID string) SharedKeysClient {
	return original.NewSharedKeysClient(subscriptionID)
}
func NewSharedKeysClientWithBaseURI(baseURI string, subscriptionID string) SharedKeysClient {
	return original.NewSharedKeysClientWithBaseURI(baseURI, subscriptionID)
}
func NewStorageInsightConfigsClient(subscriptionID string) StorageInsightConfigsClient {
	return original.NewStorageInsightConfigsClient(subscriptionID)
}
func NewStorageInsightConfigsClientWithBaseURI(baseURI string, subscriptionID string) StorageInsightConfigsClient {
	return original.NewStorageInsightConfigsClientWithBaseURI(baseURI, subscriptionID)
}
func NewStorageInsightListResultIterator(page StorageInsightListResultPage) StorageInsightListResultIterator {
	return original.NewStorageInsightListResultIterator(page)
}
func NewStorageInsightListResultPage(cur StorageInsightListResult, getNextPage func(context.Context, StorageInsightListResult) (StorageInsightListResult, error)) StorageInsightListResultPage {
	return original.NewStorageInsightListResultPage(cur, getNextPage)
}
func NewTablesClient(subscriptionID string) TablesClient {
	return original.NewTablesClient(subscriptionID)
}
func NewTablesClientWithBaseURI(baseURI string, subscriptionID string) TablesClient {
	return original.NewTablesClientWithBaseURI(baseURI, subscriptionID)
}
func NewUsagesClient(subscriptionID string) UsagesClient {
	return original.NewUsagesClient(subscriptionID)
}
func NewUsagesClientWithBaseURI(baseURI string, subscriptionID string) UsagesClient {
	return original.NewUsagesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewWorkspacePurgeClient(subscriptionID string) WorkspacePurgeClient {
	return original.NewWorkspacePurgeClient(subscriptionID)
}
func NewWorkspacePurgeClientWithBaseURI(baseURI string, subscriptionID string) WorkspacePurgeClient {
	return original.NewWorkspacePurgeClientWithBaseURI(baseURI, subscriptionID)
}
func NewWorkspacesClient(subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClient(subscriptionID)
}
func NewWorkspacesClientWithBaseURI(baseURI string, subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleBillingTypeValues() []BillingType {
	return original.PossibleBillingTypeValues()
}
func PossibleClusterEntityStatusValues() []ClusterEntityStatus {
	return original.PossibleClusterEntityStatusValues()
}
func PossibleClusterSkuNameEnumValues() []ClusterSkuNameEnum {
	return original.PossibleClusterSkuNameEnumValues()
}
func PossibleDataIngestionStatusValues() []DataIngestionStatus {
	return original.PossibleDataIngestionStatusValues()
}
func PossibleDataSourceKindValues() []DataSourceKind {
	return original.PossibleDataSourceKindValues()
}
func PossibleDataSourceTypeValues() []DataSourceType {
	return original.PossibleDataSourceTypeValues()
}
func PossibleIdentityTypeForClusterValues() []IdentityTypeForCluster {
	return original.PossibleIdentityTypeForClusterValues()
}
func PossibleLinkedServiceEntityStatusValues() []LinkedServiceEntityStatus {
	return original.PossibleLinkedServiceEntityStatusValues()
}
func PossiblePublicNetworkAccessTypeValues() []PublicNetworkAccessType {
	return original.PossiblePublicNetworkAccessTypeValues()
}
func PossiblePurgeStateValues() []PurgeState {
	return original.PossiblePurgeStateValues()
}
func PossibleSearchSortEnumValues() []SearchSortEnum {
	return original.PossibleSearchSortEnumValues()
}
func PossibleSkuNameEnumValues() []SkuNameEnum {
	return original.PossibleSkuNameEnumValues()
}
func PossibleStorageInsightStateValues() []StorageInsightState {
	return original.PossibleStorageInsightStateValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func PossibleWorkspaceEntityStatusValues() []WorkspaceEntityStatus {
	return original.PossibleWorkspaceEntityStatusValues()
}
func PossibleWorkspaceSkuNameEnumValues() []WorkspaceSkuNameEnum {
	return original.PossibleWorkspaceSkuNameEnumValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
