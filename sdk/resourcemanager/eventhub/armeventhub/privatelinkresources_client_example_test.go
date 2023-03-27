//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/2f2b633d940230cbbd5bcf1339a2e1c48674e4a2/specification/eventhub/resource-manager/Microsoft.EventHub/preview/2022-01-01-preview/examples/NameSpaces/PrivateLinkResourcesGet.json
func ExamplePrivateLinkResourcesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateLinkResourcesClient().Get(ctx, "ArunMonocle", "sdk-Namespace-2924", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateLinkResourcesListResult = armeventhub.PrivateLinkResourcesListResult{
	// 	Value: []*armeventhub.PrivateLinkResource{
	// 		{
	// 			Name: to.Ptr("namespace"),
	// 			Type: to.Ptr("Microsoft.EventHub/namespaces/privateLinkResources"),
	// 			ID: to.Ptr("subscriptions/subID/resourceGroups/SDK-EventHub-4794/providers/Microsoft.EventHub/namespaces/sdk-Namespace-5828/privateLinkResources/namespace"),
	// 			Properties: &armeventhub.PrivateLinkResourceProperties{
	// 				GroupID: to.Ptr("namespace"),
	// 				RequiredMembers: []*string{
	// 					to.Ptr("namespace")},
	// 					RequiredZoneNames: []*string{
	// 						to.Ptr("privatelink.EventHub.windows.net")},
	// 					},
	// 			}},
	// 		}
}
