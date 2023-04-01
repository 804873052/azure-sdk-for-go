//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstorsimple8000series

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
//   - subscriptionID - The subscription id
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

func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewManagersClient() *ManagersClient {
	subClient, _ := NewManagersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAccessControlRecordsClient() *AccessControlRecordsClient {
	subClient, _ := NewAccessControlRecordsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAlertsClient() *AlertsClient {
	subClient, _ := NewAlertsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBandwidthSettingsClient() *BandwidthSettingsClient {
	subClient, _ := NewBandwidthSettingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewCloudAppliancesClient() *CloudAppliancesClient {
	subClient, _ := NewCloudAppliancesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDevicesClient() *DevicesClient {
	subClient, _ := NewDevicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDeviceSettingsClient() *DeviceSettingsClient {
	subClient, _ := NewDeviceSettingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBackupPoliciesClient() *BackupPoliciesClient {
	subClient, _ := NewBackupPoliciesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBackupSchedulesClient() *BackupSchedulesClient {
	subClient, _ := NewBackupSchedulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBackupsClient() *BackupsClient {
	subClient, _ := NewBackupsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewHardwareComponentGroupsClient() *HardwareComponentGroupsClient {
	subClient, _ := NewHardwareComponentGroupsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewJobsClient() *JobsClient {
	subClient, _ := NewJobsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVolumeContainersClient() *VolumeContainersClient {
	subClient, _ := NewVolumeContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewVolumesClient() *VolumesClient {
	subClient, _ := NewVolumesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewStorageAccountCredentialsClient() *StorageAccountCredentialsClient {
	subClient, _ := NewStorageAccountCredentialsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}