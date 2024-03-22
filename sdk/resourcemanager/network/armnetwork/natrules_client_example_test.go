//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v5"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/network/resource-manager/Microsoft.Network/stable/2023-09-01/examples/NatRuleGet.json
func ExampleNatRulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNatRulesClient().Get(ctx, "rg1", "gateway1", "natRule1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.VPNGatewayNatRule = armnetwork.VPNGatewayNatRule{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/natRules/natRule1"),
	// 	Name: to.Ptr("natRule1"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.VPNGatewayNatRuleProperties{
	// 		Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
	// 		EgressVPNSiteLinkConnections: []*armnetwork.SubResource{
	// 		},
	// 		ExternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 		},
	// 		IngressVPNSiteLinkConnections: []*armnetwork.SubResource{
	// 		},
	// 		InternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 			{
	// 				AddressSpace: to.Ptr("10.4.0.0/24"),
	// 		}},
	// 		Mode: to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/network/resource-manager/Microsoft.Network/stable/2023-09-01/examples/NatRulePut.json
func ExampleNatRulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNatRulesClient().BeginCreateOrUpdate(ctx, "rg1", "gateway1", "natRule1", armnetwork.VPNGatewayNatRule{
		Properties: &armnetwork.VPNGatewayNatRuleProperties{
			Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
			ExternalMappings: []*armnetwork.VPNNatRuleMapping{
				{
					AddressSpace: to.Ptr("192.168.21.0/24"),
				}},
			InternalMappings: []*armnetwork.VPNNatRuleMapping{
				{
					AddressSpace: to.Ptr("10.4.0.0/24"),
				}},
			IPConfigurationID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/cloudnet1-VNG/ipConfigurations/default"),
			Mode:              to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
		},
	}, nil)
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
	// res.VPNGatewayNatRule = armnetwork.VPNGatewayNatRule{
	// 	ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/natRules/natRule1"),
	// 	Name: to.Ptr("natRule1"),
	// 	Etag: to.Ptr("w/\\00000000-0000-0000-0000-000000000000\\"),
	// 	Properties: &armnetwork.VPNGatewayNatRuleProperties{
	// 		Type: to.Ptr(armnetwork.VPNNatRuleTypeStatic),
	// 		EgressVPNSiteLinkConnections: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/vpnConnections/vpnConnection1/vpnLinkConnections/vpnLinkConnection1"),
	// 		}},
	// 		ExternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 			{
	// 				AddressSpace: to.Ptr("192.168.21.0/24"),
	// 		}},
	// 		IngressVPNSiteLinkConnections: []*armnetwork.SubResource{
	// 			{
	// 				ID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/vpnConnections/vpnConnection1/vpnLinkConnections/vpnLinkConnection2"),
	// 		}},
	// 		InternalMappings: []*armnetwork.VPNNatRuleMapping{
	// 			{
	// 				AddressSpace: to.Ptr("10.4.0.0/24"),
	// 		}},
	// 		IPConfigurationID: to.Ptr("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/cloudnet1-VNG/ipConfigurations/default"),
	// 		Mode: to.Ptr(armnetwork.VPNNatRuleModeEgressSnat),
	// 		ProvisioningState: to.Ptr(armnetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/network/resource-manager/Microsoft.Network/stable/2023-09-01/examples/NatRuleDelete.json
func ExampleNatRulesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNatRulesClient().BeginDelete(ctx, "rg1", "gateway1", "natRule1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d4205894880b989ede35d62d97c8e901ed14fb5a/specification/network/resource-manager/Microsoft.Network/stable/2023-09-01/examples/NatRuleList.json
func ExampleNatRulesClient_NewListByVPNGatewayPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNatRulesClient().NewListByVPNGatewayPager("rg1", "gateway1", nil)
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
		// page.ListVPNGatewayNatRulesResult = armnetwork.ListVPNGatewayNatRulesResult{
		// }
	}
}
