//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdataboxedge

import (
	"github.com/Azure/azure-sdk-for-go/profile/p20200901/internal"
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
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(internal.ModuleName+"/armdataboxedge.ClientFactory", internal.ModuleVersion, credential, options)
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

func (c *ClientFactory) NewDevicesClient() *DevicesClient {
	subClient, _ := NewDevicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewAlertsClient() *AlertsClient {
	subClient, _ := NewAlertsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewBandwidthSchedulesClient() *BandwidthSchedulesClient {
	subClient, _ := NewBandwidthSchedulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewJobsClient() *JobsClient {
	subClient, _ := NewJobsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewNodesClient() *NodesClient {
	subClient, _ := NewNodesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOperationsStatusClient() *OperationsStatusClient {
	subClient, _ := NewOperationsStatusClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewOrdersClient() *OrdersClient {
	subClient, _ := NewOrdersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewRolesClient() *RolesClient {
	subClient, _ := NewRolesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSharesClient() *SharesClient {
	subClient, _ := NewSharesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewStorageAccountCredentialsClient() *StorageAccountCredentialsClient {
	subClient, _ := NewStorageAccountCredentialsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewStorageAccountsClient() *StorageAccountsClient {
	subClient, _ := NewStorageAccountsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewContainersClient() *ContainersClient {
	subClient, _ := NewContainersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewTriggersClient() *TriggersClient {
	subClient, _ := NewTriggersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewUsersClient() *UsersClient {
	subClient, _ := NewUsersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewSKUsClient() *SKUsClient {
	subClient, _ := NewSKUsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
