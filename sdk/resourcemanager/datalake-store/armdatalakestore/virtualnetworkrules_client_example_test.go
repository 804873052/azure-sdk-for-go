//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatalakestore_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datalake-store/armdatalakestore"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/VirtualNetworkRules_ListByAccount.json
func ExampleVirtualNetworkRulesClient_NewListByAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakestore.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewVirtualNetworkRulesClient().NewListByAccountPager("contosorg", "contosoadla", nil)
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
		// page.VirtualNetworkRuleListResult = armdatalakestore.VirtualNetworkRuleListResult{
		// 	Value: []*armdatalakestore.VirtualNetworkRule{
		// 		{
		// 			Name: to.Ptr("test_virtual_network_rules_name"),
		// 			Type: to.Ptr("test_type"),
		// 			ID: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
		// 			Properties: &armdatalakestore.VirtualNetworkRuleProperties{
		// 				SubnetID: to.Ptr("test_subnetId"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/VirtualNetworkRules_CreateOrUpdate.json
func ExampleVirtualNetworkRulesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakestore.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworkRulesClient().CreateOrUpdate(ctx, "contosorg", "contosoadla", "test_virtual_network_rules_name", armdatalakestore.CreateOrUpdateVirtualNetworkRuleParameters{
		Properties: &armdatalakestore.CreateOrUpdateVirtualNetworkRuleProperties{
			SubnetID: to.Ptr("test_subnetId"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetworkRule = armdatalakestore.VirtualNetworkRule{
	// 	Name: to.Ptr("test_virtual_network_rules_name"),
	// 	Type: to.Ptr("test_type"),
	// 	ID: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
	// 	Properties: &armdatalakestore.VirtualNetworkRuleProperties{
	// 		SubnetID: to.Ptr("test_subnetId"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/VirtualNetworkRules_Get.json
func ExampleVirtualNetworkRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakestore.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworkRulesClient().Get(ctx, "contosorg", "contosoadla", "test_virtual_network_rules_name", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetworkRule = armdatalakestore.VirtualNetworkRule{
	// 	Name: to.Ptr("test_virtual_network_rules_name"),
	// 	Type: to.Ptr("test_type"),
	// 	ID: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
	// 	Properties: &armdatalakestore.VirtualNetworkRuleProperties{
	// 		SubnetID: to.Ptr("test_subnetId"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/VirtualNetworkRules_Update.json
func ExampleVirtualNetworkRulesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakestore.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewVirtualNetworkRulesClient().Update(ctx, "contosorg", "contosoadla", "test_virtual_network_rules_name", &armdatalakestore.VirtualNetworkRulesClientUpdateOptions{Parameters: &armdatalakestore.UpdateVirtualNetworkRuleParameters{
		Properties: &armdatalakestore.UpdateVirtualNetworkRuleProperties{
			SubnetID: to.Ptr("test_subnetId"),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VirtualNetworkRule = armdatalakestore.VirtualNetworkRule{
	// 	Name: to.Ptr("test_virtual_network_rules_name"),
	// 	Type: to.Ptr("test_type"),
	// 	ID: to.Ptr("34adfa4f-cedf-4dc0-ba29-b6d1a69ab345"),
	// 	Properties: &armdatalakestore.VirtualNetworkRuleProperties{
	// 		SubnetID: to.Ptr("test_subnetId"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/datalake-store/resource-manager/Microsoft.DataLakeStore/stable/2016-11-01/examples/VirtualNetworkRules_Delete.json
func ExampleVirtualNetworkRulesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdatalakestore.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewVirtualNetworkRulesClient().Delete(ctx, "contosorg", "contosoadla", "test_virtual_network_rules_name", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}