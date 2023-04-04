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

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridnetwork/armhybridnetwork/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/NetworkFunctionDelete.json
func ExampleNetworkFunctionsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkFunctionsClient().BeginDelete(ctx, "rg", "testNf", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/NetworkFunctionGet.json
func ExampleNetworkFunctionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkFunctionsClient().Get(ctx, "rg", "testNf", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkFunction = armhybridnetwork.NetworkFunction{
	// 	Name: to.Ptr("testNf"),
	// 	Type: to.Ptr("Microsoft.HybridNetwork/networkFunctions"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/networkFunctions/testNf"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armhybridnetwork.NetworkFunctionPropertiesFormat{
	// 		Device: &armhybridnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/devices/testDevice"),
	// 		},
	// 		ManagedApplication: &armhybridnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/testServiceKey"),
	// 		},
	// 		ManagedApplicationParameters: map[string]any{
	// 		},
	// 		NetworkFunctionUserConfigurations: []*armhybridnetwork.NetworkFunctionUserConfiguration{
	// 			{
	// 				NetworkInterfaces: []*armhybridnetwork.NetworkInterface{
	// 					{
	// 						IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
	// 							{
	// 								Gateway: to.Ptr(""),
	// 								IPAddress: to.Ptr(""),
	// 								IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
	// 								IPVersion: to.Ptr(armhybridnetwork.IPVersionIPv4),
	// 								Subnet: to.Ptr(""),
	// 						}},
	// 						MacAddress: to.Ptr(""),
	// 						NetworkInterfaceName: to.Ptr("nic1"),
	// 						VMSwitchType: to.Ptr(armhybridnetwork.VMSwitchTypeManagement),
	// 					},
	// 					{
	// 						IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
	// 							{
	// 								Gateway: to.Ptr(""),
	// 								IPAddress: to.Ptr(""),
	// 								IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
	// 								IPVersion: to.Ptr(armhybridnetwork.IPVersionIPv4),
	// 								Subnet: to.Ptr(""),
	// 						}},
	// 						MacAddress: to.Ptr("DC-97-F8-79-16-7D"),
	// 						NetworkInterfaceName: to.Ptr("nic2"),
	// 						VMSwitchType: to.Ptr(armhybridnetwork.VMSwitchTypeWan),
	// 				}},
	// 				RoleName: to.Ptr("testRole"),
	// 				UserDataParameters: map[string]any{
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armhybridnetwork.ProvisioningStateSucceeded),
	// 		ServiceKey: to.Ptr("testServiceKey"),
	// 		SKUName: to.Ptr("testSku"),
	// 		SKUType: to.Ptr(armhybridnetwork.SKUTypeSDWAN),
	// 		VendorName: to.Ptr("testVendor"),
	// 		VendorProvisioningState: to.Ptr(armhybridnetwork.VendorProvisioningStateNotProvisioned),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/NetworkFunctionCreate.json
func ExampleNetworkFunctionsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkFunctionsClient().BeginCreateOrUpdate(ctx, "rg", "testNf", armhybridnetwork.NetworkFunction{
		Location: to.Ptr("eastus"),
		Properties: &armhybridnetwork.NetworkFunctionPropertiesFormat{
			Device: &armhybridnetwork.SubResource{
				ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/devices/testDevice"),
			},
			ManagedApplicationParameters: map[string]any{},
			NetworkFunctionUserConfigurations: []*armhybridnetwork.NetworkFunctionUserConfiguration{
				{
					NetworkInterfaces: []*armhybridnetwork.NetworkInterface{
						{
							IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
								{
									Gateway:            to.Ptr(""),
									IPAddress:          to.Ptr(""),
									IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
									IPVersion:          to.Ptr(armhybridnetwork.IPVersionIPv4),
									Subnet:             to.Ptr(""),
								}},
							MacAddress:           to.Ptr(""),
							NetworkInterfaceName: to.Ptr("nic1"),
							VMSwitchType:         to.Ptr(armhybridnetwork.VMSwitchTypeManagement),
						},
						{
							IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
								{
									Gateway:            to.Ptr(""),
									IPAddress:          to.Ptr(""),
									IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
									IPVersion:          to.Ptr(armhybridnetwork.IPVersionIPv4),
									Subnet:             to.Ptr(""),
								}},
							MacAddress:           to.Ptr("DC-97-F8-79-16-7D"),
							NetworkInterfaceName: to.Ptr("nic2"),
							VMSwitchType:         to.Ptr(armhybridnetwork.VMSwitchTypeWan),
						}},
					RoleName:           to.Ptr("testRole"),
					UserDataParameters: map[string]any{},
				}},
			SKUName:    to.Ptr("testSku"),
			SKUType:    to.Ptr(armhybridnetwork.SKUTypeSDWAN),
			VendorName: to.Ptr("testVendor"),
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
	// res.NetworkFunction = armhybridnetwork.NetworkFunction{
	// 	Name: to.Ptr("testNf"),
	// 	Type: to.Ptr("Microsoft.HybridNetwork/networkFunctions"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/networkFunctions/testNf"),
	// 	Location: to.Ptr("eastus"),
	// 	Properties: &armhybridnetwork.NetworkFunctionPropertiesFormat{
	// 		Device: &armhybridnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/devices/testDevice"),
	// 		},
	// 		ManagedApplication: &armhybridnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/testServiceKey"),
	// 		},
	// 		ManagedApplicationParameters: map[string]any{
	// 		},
	// 		NetworkFunctionUserConfigurations: []*armhybridnetwork.NetworkFunctionUserConfiguration{
	// 			{
	// 				NetworkInterfaces: []*armhybridnetwork.NetworkInterface{
	// 					{
	// 						IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
	// 							{
	// 								Gateway: to.Ptr(""),
	// 								IPAddress: to.Ptr(""),
	// 								IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
	// 								IPVersion: to.Ptr(armhybridnetwork.IPVersionIPv4),
	// 								Subnet: to.Ptr(""),
	// 						}},
	// 						MacAddress: to.Ptr(""),
	// 						NetworkInterfaceName: to.Ptr("nic1"),
	// 						VMSwitchType: to.Ptr(armhybridnetwork.VMSwitchTypeManagement),
	// 					},
	// 					{
	// 						IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
	// 							{
	// 								Gateway: to.Ptr(""),
	// 								IPAddress: to.Ptr(""),
	// 								IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
	// 								IPVersion: to.Ptr(armhybridnetwork.IPVersionIPv4),
	// 								Subnet: to.Ptr(""),
	// 						}},
	// 						MacAddress: to.Ptr("DC-97-F8-79-16-7D"),
	// 						NetworkInterfaceName: to.Ptr("nic2"),
	// 						VMSwitchType: to.Ptr(armhybridnetwork.VMSwitchTypeWan),
	// 				}},
	// 				RoleName: to.Ptr("testRole"),
	// 				UserDataParameters: map[string]any{
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armhybridnetwork.ProvisioningStateSucceeded),
	// 		ServiceKey: to.Ptr("testServiceKey"),
	// 		SKUName: to.Ptr("testSku"),
	// 		SKUType: to.Ptr(armhybridnetwork.SKUTypeSDWAN),
	// 		VendorName: to.Ptr("testVendor"),
	// 		VendorProvisioningState: to.Ptr(armhybridnetwork.VendorProvisioningStateNotProvisioned),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/NetworkFunctionUpdateTags.json
func ExampleNetworkFunctionsClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewNetworkFunctionsClient().UpdateTags(ctx, "rg", "testNf", armhybridnetwork.TagsObject{
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.NetworkFunction = armhybridnetwork.NetworkFunction{
	// 	Name: to.Ptr("testNf"),
	// 	Type: to.Ptr("Microsoft.HybridNetwork/networkFunctions"),
	// 	ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/networkFunctions/testNf"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Properties: &armhybridnetwork.NetworkFunctionPropertiesFormat{
	// 		Device: &armhybridnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/devices/testDevice"),
	// 		},
	// 		ManagedApplication: &armhybridnetwork.SubResource{
	// 			ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/testServiceKey"),
	// 		},
	// 		ManagedApplicationParameters: map[string]any{
	// 		},
	// 		NetworkFunctionUserConfigurations: []*armhybridnetwork.NetworkFunctionUserConfiguration{
	// 			{
	// 				NetworkInterfaces: []*armhybridnetwork.NetworkInterface{
	// 					{
	// 						IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
	// 							{
	// 								Gateway: to.Ptr(""),
	// 								IPAddress: to.Ptr(""),
	// 								IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
	// 								IPVersion: to.Ptr(armhybridnetwork.IPVersionIPv4),
	// 								Subnet: to.Ptr(""),
	// 						}},
	// 						MacAddress: to.Ptr(""),
	// 						NetworkInterfaceName: to.Ptr("nic1"),
	// 						VMSwitchType: to.Ptr(armhybridnetwork.VMSwitchTypeManagement),
	// 					},
	// 					{
	// 						IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
	// 							{
	// 								Gateway: to.Ptr(""),
	// 								IPAddress: to.Ptr(""),
	// 								IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
	// 								IPVersion: to.Ptr(armhybridnetwork.IPVersionIPv4),
	// 								Subnet: to.Ptr(""),
	// 						}},
	// 						MacAddress: to.Ptr("DC-97-F8-79-16-7D"),
	// 						NetworkInterfaceName: to.Ptr("nic2"),
	// 						VMSwitchType: to.Ptr(armhybridnetwork.VMSwitchTypeWan),
	// 				}},
	// 				RoleName: to.Ptr("testRole"),
	// 				UserDataParameters: map[string]any{
	// 				},
	// 		}},
	// 		ProvisioningState: to.Ptr(armhybridnetwork.ProvisioningStateSucceeded),
	// 		ServiceKey: to.Ptr("testServiceKey"),
	// 		SKUName: to.Ptr("testSku"),
	// 		SKUType: to.Ptr(armhybridnetwork.SKUTypeSDWAN),
	// 		VendorName: to.Ptr("testVendor"),
	// 		VendorProvisioningState: to.Ptr(armhybridnetwork.VendorProvisioningStateNotProvisioned),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/NetworkFunctionListBySubscription.json
func ExampleNetworkFunctionsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkFunctionsClient().NewListBySubscriptionPager(nil)
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
		// page.NetworkFunctionListResult = armhybridnetwork.NetworkFunctionListResult{
		// 	Value: []*armhybridnetwork.NetworkFunction{
		// 		{
		// 			Name: to.Ptr("testNf"),
		// 			Type: to.Ptr("Microsoft.HybridNetwork/networkFunctions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/networkFunctions/testNf"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armhybridnetwork.NetworkFunctionPropertiesFormat{
		// 				Device: &armhybridnetwork.SubResource{
		// 					ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/devices/testDevice"),
		// 				},
		// 				ManagedApplication: &armhybridnetwork.SubResource{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/testServiceKey"),
		// 				},
		// 				ManagedApplicationParameters: map[string]any{
		// 				},
		// 				NetworkFunctionUserConfigurations: []*armhybridnetwork.NetworkFunctionUserConfiguration{
		// 					{
		// 						NetworkInterfaces: []*armhybridnetwork.NetworkInterface{
		// 							{
		// 								IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
		// 									{
		// 										Gateway: to.Ptr(""),
		// 										IPAddress: to.Ptr(""),
		// 										IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
		// 										IPVersion: to.Ptr(armhybridnetwork.IPVersionIPv4),
		// 										Subnet: to.Ptr(""),
		// 								}},
		// 								MacAddress: to.Ptr(""),
		// 								NetworkInterfaceName: to.Ptr("nic1"),
		// 								VMSwitchType: to.Ptr(armhybridnetwork.VMSwitchTypeManagement),
		// 							},
		// 							{
		// 								IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
		// 									{
		// 										Gateway: to.Ptr(""),
		// 										IPAddress: to.Ptr(""),
		// 										IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
		// 										IPVersion: to.Ptr(armhybridnetwork.IPVersionIPv4),
		// 										Subnet: to.Ptr(""),
		// 								}},
		// 								MacAddress: to.Ptr("DC-97-F8-79-16-7D"),
		// 								NetworkInterfaceName: to.Ptr("nic2"),
		// 								VMSwitchType: to.Ptr(armhybridnetwork.VMSwitchTypeWan),
		// 						}},
		// 						RoleName: to.Ptr("testRole"),
		// 						UserDataParameters: map[string]any{
		// 						},
		// 				}},
		// 				ProvisioningState: to.Ptr(armhybridnetwork.ProvisioningStateSucceeded),
		// 				ServiceKey: to.Ptr("testServiceKey"),
		// 				SKUName: to.Ptr("testSku"),
		// 				SKUType: to.Ptr(armhybridnetwork.SKUTypeSDWAN),
		// 				VendorName: to.Ptr("testVendor"),
		// 				VendorProvisioningState: to.Ptr(armhybridnetwork.VendorProvisioningStateNotProvisioned),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/NetworkFunctionListByResourceGroup.json
func ExampleNetworkFunctionsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewNetworkFunctionsClient().NewListByResourceGroupPager("rg", nil)
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
		// page.NetworkFunctionListResult = armhybridnetwork.NetworkFunctionListResult{
		// 	Value: []*armhybridnetwork.NetworkFunction{
		// 		{
		// 			Name: to.Ptr("testNf"),
		// 			Type: to.Ptr("Microsoft.HybridNetwork/networkFunctions"),
		// 			ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/networkFunctions/testNf"),
		// 			Location: to.Ptr("eastus"),
		// 			Properties: &armhybridnetwork.NetworkFunctionPropertiesFormat{
		// 				Device: &armhybridnetwork.SubResource{
		// 					ID: to.Ptr("/subscriptions/subid/resourcegroups/rg/providers/Microsoft.HybridNetwork/devices/testDevice"),
		// 				},
		// 				ManagedApplication: &armhybridnetwork.SubResource{
		// 					ID: to.Ptr("/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/testServiceKey"),
		// 				},
		// 				ManagedApplicationParameters: map[string]any{
		// 				},
		// 				NetworkFunctionUserConfigurations: []*armhybridnetwork.NetworkFunctionUserConfiguration{
		// 					{
		// 						NetworkInterfaces: []*armhybridnetwork.NetworkInterface{
		// 							{
		// 								IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
		// 									{
		// 										Gateway: to.Ptr(""),
		// 										IPAddress: to.Ptr(""),
		// 										IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
		// 										IPVersion: to.Ptr(armhybridnetwork.IPVersionIPv4),
		// 										Subnet: to.Ptr(""),
		// 								}},
		// 								MacAddress: to.Ptr(""),
		// 								NetworkInterfaceName: to.Ptr("nic1"),
		// 								VMSwitchType: to.Ptr(armhybridnetwork.VMSwitchTypeManagement),
		// 							},
		// 							{
		// 								IPConfigurations: []*armhybridnetwork.NetworkInterfaceIPConfiguration{
		// 									{
		// 										Gateway: to.Ptr(""),
		// 										IPAddress: to.Ptr(""),
		// 										IPAllocationMethod: to.Ptr(armhybridnetwork.IPAllocationMethodDynamic),
		// 										IPVersion: to.Ptr(armhybridnetwork.IPVersionIPv4),
		// 										Subnet: to.Ptr(""),
		// 								}},
		// 								MacAddress: to.Ptr("DC-97-F8-79-16-7D"),
		// 								NetworkInterfaceName: to.Ptr("nic2"),
		// 								VMSwitchType: to.Ptr(armhybridnetwork.VMSwitchTypeWan),
		// 						}},
		// 						RoleName: to.Ptr("testRole"),
		// 						UserDataParameters: map[string]any{
		// 						},
		// 				}},
		// 				ProvisioningState: to.Ptr(armhybridnetwork.ProvisioningStateSucceeded),
		// 				ServiceKey: to.Ptr("testServiceKey"),
		// 				SKUName: to.Ptr("testSku"),
		// 				SKUType: to.Ptr(armhybridnetwork.SKUTypeSDWAN),
		// 				VendorName: to.Ptr("testVendor"),
		// 				VendorProvisioningState: to.Ptr(armhybridnetwork.VendorProvisioningStateNotProvisioned),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/preview/2022-01-01-preview/examples/NetworkFunctionsExecuteRequest.json
func ExampleNetworkFunctionsClient_BeginExecuteRequest() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armhybridnetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewNetworkFunctionsClient().BeginExecuteRequest(ctx, "rg", "testNetworkfunction", armhybridnetwork.ExecuteRequestParameters{
		RequestMetadata: &armhybridnetwork.RequestMetadata{
			APIVersion:     to.Ptr("apiVersionQueryString"),
			HTTPMethod:     to.Ptr(armhybridnetwork.HTTPMethodPost),
			RelativePath:   to.Ptr("/simProfiles/testSimProfile"),
			SerializedBody: to.Ptr("{\"subscriptionProfile\":\"ChantestSubscription15\",\"permanentKey\":\"00112233445566778899AABBCCDDEEFF\",\"opcOperatorCode\":\"63bfa50ee6523365ff14c1f45f88737d\",\"staticIpAddresses\":{\"internet\":{\"ipv4Addr\":\"198.51.100.1\",\"ipv6Prefix\":\"2001:db8:abcd:12::0/64\"},\"another_network\":{\"ipv6Prefix\":\"2001:111:cdef:22::0/64\"}}}"),
		},
		ServiceEndpoint: to.Ptr("serviceEndpoint"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
