//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicesbackup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName+".ClientFactory", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

func (c *ClientFactory) NewBackupResourceStorageConfigsNonCRRClient() *BackupResourceStorageConfigsNonCRRClient {
	subClient, _ := NewBackupResourceStorageConfigsNonCRRClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProtectionIntentClient() *ProtectionIntentClient {
	subClient, _ := NewProtectionIntentClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBackupStatusClient() *BackupStatusClient {
	subClient, _ := NewBackupStatusClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewFeatureSupportClient() *FeatureSupportClient {
	subClient, _ := NewFeatureSupportClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBackupProtectionIntentClient() *BackupProtectionIntentClient {
	subClient, _ := NewBackupProtectionIntentClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBackupUsageSummariesClient() *BackupUsageSummariesClient {
	subClient, _ := NewBackupUsageSummariesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBackupResourceVaultConfigsClient() *BackupResourceVaultConfigsClient {
	subClient, _ := NewBackupResourceVaultConfigsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBackupResourceEncryptionConfigsClient() *BackupResourceEncryptionConfigsClient {
	subClient, _ := NewBackupResourceEncryptionConfigsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateEndpointConnectionClient() *PrivateEndpointConnectionClient {
	subClient, _ := NewPrivateEndpointConnectionClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPrivateEndpointClient() *PrivateEndpointClient {
	subClient, _ := NewPrivateEndpointClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewClient() *Client {
	subClient, _ := NewClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBMSPrepareDataMoveOperationResultClient() *BMSPrepareDataMoveOperationResultClient {
	subClient, _ := NewBMSPrepareDataMoveOperationResultClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProtectedItemsClient() *ProtectedItemsClient {
	subClient, _ := NewProtectedItemsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProtectedItemOperationResultsClient() *ProtectedItemOperationResultsClient {
	subClient, _ := NewProtectedItemOperationResultsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRecoveryPointsClient() *RecoveryPointsClient {
	subClient, _ := NewRecoveryPointsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRestoresClient() *RestoresClient {
	subClient, _ := NewRestoresClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBackupPoliciesClient() *BackupPoliciesClient {
	subClient, _ := NewBackupPoliciesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProtectionPoliciesClient() *ProtectionPoliciesClient {
	subClient, _ := NewProtectionPoliciesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProtectionPolicyOperationResultsClient() *ProtectionPolicyOperationResultsClient {
	subClient, _ := NewProtectionPolicyOperationResultsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBackupJobsClient() *BackupJobsClient {
	subClient, _ := NewBackupJobsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewJobDetailsClient() *JobDetailsClient {
	subClient, _ := NewJobDetailsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewJobCancellationsClient() *JobCancellationsClient {
	subClient, _ := NewJobCancellationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewJobOperationResultsClient() *JobOperationResultsClient {
	subClient, _ := NewJobOperationResultsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewExportJobsOperationResultsClient() *ExportJobsOperationResultsClient {
	subClient, _ := NewExportJobsOperationResultsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewJobsClient() *JobsClient {
	subClient, _ := NewJobsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBackupProtectedItemsClient() *BackupProtectedItemsClient {
	subClient, _ := NewBackupProtectedItemsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationClient() *OperationClient {
	subClient, _ := NewOperationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewValidateOperationClient() *ValidateOperationClient {
	subClient, _ := NewValidateOperationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewValidateOperationResultsClient() *ValidateOperationResultsClient {
	subClient, _ := NewValidateOperationResultsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewValidateOperationStatusesClient() *ValidateOperationStatusesClient {
	subClient, _ := NewValidateOperationStatusesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBackupEnginesClient() *BackupEnginesClient {
	subClient, _ := NewBackupEnginesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProtectionContainerRefreshOperationResultsClient() *ProtectionContainerRefreshOperationResultsClient {
	subClient, _ := NewProtectionContainerRefreshOperationResultsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProtectableContainersClient() *ProtectableContainersClient {
	subClient, _ := NewProtectableContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProtectionContainersClient() *ProtectionContainersClient {
	subClient, _ := NewProtectionContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBackupWorkloadItemsClient() *BackupWorkloadItemsClient {
	subClient, _ := NewBackupWorkloadItemsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProtectionContainerOperationResultsClient() *ProtectionContainerOperationResultsClient {
	subClient, _ := NewProtectionContainerOperationResultsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBackupsClient() *BackupsClient {
	subClient, _ := NewBackupsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProtectedItemOperationStatusesClient() *ProtectedItemOperationStatusesClient {
	subClient, _ := NewProtectedItemOperationStatusesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewItemLevelRecoveryConnectionsClient() *ItemLevelRecoveryConnectionsClient {
	subClient, _ := NewItemLevelRecoveryConnectionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBackupOperationResultsClient() *BackupOperationResultsClient {
	subClient, _ := NewBackupOperationResultsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBackupOperationStatusesClient() *BackupOperationStatusesClient {
	subClient, _ := NewBackupOperationStatusesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewProtectionPolicyOperationStatusesClient() *ProtectionPolicyOperationStatusesClient {
	subClient, _ := NewProtectionPolicyOperationStatusesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBackupProtectableItemsClient() *BackupProtectableItemsClient {
	subClient, _ := NewBackupProtectableItemsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBackupProtectionContainersClient() *BackupProtectionContainersClient {
	subClient, _ := NewBackupProtectionContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDeletedProtectionContainersClient() *DeletedProtectionContainersClient {
	subClient, _ := NewDeletedProtectionContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSecurityPINsClient() *SecurityPINsClient {
	subClient, _ := NewSecurityPINsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRecoveryPointsRecommendedForMoveClient() *RecoveryPointsRecommendedForMoveClient {
	subClient, _ := NewRecoveryPointsRecommendedForMoveClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewResourceGuardProxiesClient() *ResourceGuardProxiesClient {
	subClient, _ := NewResourceGuardProxiesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewResourceGuardProxyClient() *ResourceGuardProxyClient {
	subClient, _ := NewResourceGuardProxyClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
