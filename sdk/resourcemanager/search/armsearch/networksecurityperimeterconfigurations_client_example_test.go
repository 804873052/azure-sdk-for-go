//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsearch_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/search/armsearch"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/search/resource-manager/Microsoft.Search/preview/2024-03-01-preview/examples/NetworkSecurityPerimeterConfigurationsListByService.json
func ExampleNetworkSecurityPerimeterConfigurationsClient_NewListByServicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsearch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkSecurityPerimeterConfigurationsClient().NewListByServicePager("rg1", "mysearchservice", nil)
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
		// page.NetworkSecurityPerimeterConfigurationListResult = armsearch.NetworkSecurityPerimeterConfigurationListResult{
		// 	Value: []*armsearch.NetworkSecurityPerimeterConfiguration{
		// 		{
		// 			Name: to.Ptr("00000001-2222-3333-4444-111144444444.assoc1"),
		// 			Type: to.Ptr("Microsoft.Search/searchServices/networkSecurityPerimeterConfigurations"),
		// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Search/searchServices/mysearchservice/networkSecurityPerimeterConfigurations/00000001-2222-3333-4444-111144444444.assoc1"),
		// 			Properties: &armsearch.NetworkSecurityPerimeterConfigurationProperties{
		// 				NetworkSecurityPerimeter: &armsearch.NSPConfigPerimeter{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/networkRG/providers/Microsoft.Network/networkSecurityPerimeters/perimeter1"),
		// 					Location: to.Ptr("westus"),
		// 				},
		// 				Profile: &armsearch.NSPConfigProfile{
		// 					Name: to.Ptr("profile1"),
		// 					AccessRules: []*armsearch.NSPConfigAccessRule{
		// 						{
		// 							Name: to.Ptr("rule1"),
		// 							Properties: &armsearch.NSPConfigAccessRuleProperties{
		// 								AddressPrefixes: []*string{
		// 									to.Ptr("148.0.0.0/8"),
		// 									to.Ptr("152.4.6.0/24")},
		// 									Direction: to.Ptr("Inbound"),
		// 								},
		// 						}},
		// 						AccessRulesVersion: to.Ptr("0"),
		// 					},
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 					ResourceAssociation: &armsearch.NSPConfigAssociation{
		// 						Name: to.Ptr("assoc1"),
		// 						AccessMode: to.Ptr("Enforced"),
		// 					},
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/search/resource-manager/Microsoft.Search/preview/2024-03-01-preview/examples/NetworkSecurityPerimeterConfigurationsGet.json
func ExampleNetworkSecurityPerimeterConfigurationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsearch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkSecurityPerimeterConfigurationsClient().Get(ctx, "rg1", "mysearchservice", "00000001-2222-3333-4444-111144444444.assoc1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkSecurityPerimeterConfiguration = armsearch.NetworkSecurityPerimeterConfiguration{
	// 	Name: to.Ptr("00000001-2222-3333-4444-111144444444.assoc1"),
	// 	Type: to.Ptr("Microsoft.Search/searchServices/networkSecurityPerimeterConfigurations"),
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Search/searchServices/mysearchservice/networkSecurityPerimeterConfigurations/00000001-2222-3333-4444-111144444444.assoc1"),
	// 	Properties: &armsearch.NetworkSecurityPerimeterConfigurationProperties{
	// 		NetworkSecurityPerimeter: &armsearch.NSPConfigPerimeter{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/networkRG/providers/Microsoft.Network/networkSecurityPerimeters/perimeter1"),
	// 			Location: to.Ptr("westus"),
	// 		},
	// 		Profile: &armsearch.NSPConfigProfile{
	// 			Name: to.Ptr("profile1"),
	// 			AccessRules: []*armsearch.NSPConfigAccessRule{
	// 				{
	// 					Name: to.Ptr("rule1"),
	// 					Properties: &armsearch.NSPConfigAccessRuleProperties{
	// 						AddressPrefixes: []*string{
	// 							to.Ptr("148.0.0.0/8"),
	// 							to.Ptr("152.4.6.0/24")},
	// 							Direction: to.Ptr("Inbound"),
	// 						},
	// 				}},
	// 				AccessRulesVersion: to.Ptr("0"),
	// 			},
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 			ResourceAssociation: &armsearch.NSPConfigAssociation{
	// 				Name: to.Ptr("assoc1"),
	// 				AccessMode: to.Ptr("Enforced"),
	// 			},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/search/resource-manager/Microsoft.Search/preview/2024-03-01-preview/examples/NetworkSecurityPerimeterConfigurationsReconcile.json
func ExampleNetworkSecurityPerimeterConfigurationsClient_BeginReconcile() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsearch.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkSecurityPerimeterConfigurationsClient().BeginReconcile(ctx, "rg1", "mysearchservice", "00000001-2222-3333-4444-111144444444.assoc1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
