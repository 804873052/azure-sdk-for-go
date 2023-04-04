//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/VendorSkuPreviewListBySku.json
func ExampleVendorSKUPreviewClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVendorSKUPreviewClient().NewListPager("TestVendor", "TestSku", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.PreviewSubscriptionsList = armhybridnetwork.PreviewSubscriptionsList{
		// 	Value: []*armhybridnetwork.PreviewSubscription{
		// 		{
		// 			Name: to.Ptr("previewSub1"),
		// 			Type: to.Ptr("Microsoft.HybridNetwork/vendors/vendorskus/previewsubscriptions"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.HybridNetwork/vendors/TestVendor/vendorskus/TestSku/previewsubscriptions/previewSub1"),
		// 		},
		// 		{
		// 			Name: to.Ptr("previewSub2"),
		// 			Type: to.Ptr("Microsoft.HybridNetwork/vendors/vendorskus/previewsubscriptions"),
		// 			ID: to.Ptr("/subscriptions/subid/providers/Microsoft.HybridNetwork/vendors/TestVendor/vendorskus/TestSku/previewsubscriptions/previewSub2"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/VendorSkuPreviewCreate.json
func ExampleVendorSKUPreviewClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVendorSKUPreviewClient().BeginCreateOrUpdate(ctx, "TestVendor", "TestSku", "previewSub", armhybridnetwork.PreviewSubscription{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PreviewSubscription = armhybridnetwork.PreviewSubscription{
	// 	Name: to.Ptr("previewSub"),
	// 	Type: to.Ptr("Microsoft.HybridNetwork/vendors/vendorskus/previewsubscriptions"),
	// 	ID: to.Ptr("/subscriptions/subid/providers/Microsoft.HybridNetwork/vendors/TestVendor/vendorskus/TestSku/previewsubscriptions/previewSub"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/VendorSkuPreviewGet.json
func ExampleVendorSKUPreviewClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVendorSKUPreviewClient().Get(ctx, "TestVendor", "TestSku", "previewSub", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PreviewSubscription = armhybridnetwork.PreviewSubscription{
	// 	Name: to.Ptr("previewSub"),
	// 	Type: to.Ptr("Microsoft.HybridNetwork/vendors/vendorskus/previewsubscriptions"),
	// 	ID: to.Ptr("/subscriptions/subid/providers/Microsoft.HybridNetwork/vendors/TestVendor/vendorskus/TestSku/previewsubscriptions/previewSub"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/VendorSkuPreviewDelete.json
func ExampleVendorSKUPreviewClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewVendorSKUPreviewClient().BeginDelete(ctx, "TestVendor", "TestSku", "previewSub", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
