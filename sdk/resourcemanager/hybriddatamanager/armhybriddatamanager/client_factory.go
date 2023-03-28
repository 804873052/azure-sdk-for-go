//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybriddatamanager

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
//   - subscriptionID - The Subscription Id
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

func (c *ClientFactory) NewDataManagersClient() *DataManagersClient {
	subClient, _ := NewDataManagersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDataServicesClient() *DataServicesClient {
	subClient, _ := NewDataServicesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewJobDefinitionsClient() *JobDefinitionsClient {
	subClient, _ := NewJobDefinitionsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewJobsClient() *JobsClient {
	subClient, _ := NewJobsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDataStoresClient() *DataStoresClient {
	subClient, _ := NewDataStoresClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewDataStoreTypesClient() *DataStoreTypesClient {
	subClient, _ := NewDataStoreTypesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

func (c *ClientFactory) NewPublicKeysClient() *PublicKeysClient {
	subClient, _ := NewPublicKeysClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
