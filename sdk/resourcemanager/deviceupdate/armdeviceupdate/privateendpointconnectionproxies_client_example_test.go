//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdeviceupdate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/deviceupdate/armdeviceupdate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80065490402157d0df0dd37ab347c651b22eb576/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_ListByAccount.json
func ExamplePrivateEndpointConnectionProxiesClient_NewListByAccountPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceupdate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewPrivateEndpointConnectionProxiesClient().NewListByAccountPager("test-rg", "contoso", nil)
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
		// page.PrivateEndpointConnectionProxyListResult = armdeviceupdate.PrivateEndpointConnectionProxyListResult{
		// 	Value: []*armdeviceupdate.PrivateEndpointConnectionProxy{
		// 		{
		// 			RemotePrivateEndpoint: &armdeviceupdate.RemotePrivateEndpoint{
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}"),
		// 				ImmutableResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}"),
		// 				ImmutableSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		// 				Location: to.Ptr("westus2"),
		// 				ManualPrivateLinkServiceConnections: []*armdeviceupdate.PrivateLinkServiceConnection{
		// 					{
		// 						Name: to.Ptr("{plsConnectionName}"),
		// 						GroupIDs: []*string{
		// 							to.Ptr("DeviceUpdate")},
		// 							RequestMessage: to.Ptr("Please approve my connection, thanks."),
		// 					}},
		// 					PrivateLinkServiceProxies: []*armdeviceupdate.PrivateLinkServiceProxy{
		// 						{
		// 							GroupConnectivityInformation: []*armdeviceupdate.GroupConnectivityInformation{
		// 								{
		// 									GroupID: to.Ptr("DeviceUpdate"),
		// 									MemberName: to.Ptr("adu"),
		// 							}},
		// 							ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{privateEndpointConnectionProxyId}/privateLinkServiceProxies/{privateEndpointConnectionProxyId}"),
		// 					}},
		// 				},
		// 				Name: to.Ptr("peexample01"),
		// 				Type: to.Ptr("Microsoft.DeviceUpdate/accounts/privateEndpointConnectionProxies"),
		// 				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DeviceUpdate/accounts/contoso/privateEndpointConnectionProxies/peexample01"),
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80065490402157d0df0dd37ab347c651b22eb576/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_Validate.json
func ExamplePrivateEndpointConnectionProxiesClient_Validate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceupdate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPrivateEndpointConnectionProxiesClient().Validate(ctx, "test-rg", "contoso", "peexample01", armdeviceupdate.PrivateEndpointConnectionProxy{
		RemotePrivateEndpoint: &armdeviceupdate.RemotePrivateEndpoint{
			ID:                      to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{privateEndpointConnectionProxyId}"),
			ImmutableResourceID:     to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}"),
			ImmutableSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
			Location:                to.Ptr("westus2"),
			ManualPrivateLinkServiceConnections: []*armdeviceupdate.PrivateLinkServiceConnection{
				{
					Name: to.Ptr("{privateEndpointConnectionProxyId}"),
					GroupIDs: []*string{
						to.Ptr("DeviceUpdate")},
					RequestMessage: to.Ptr("Please approve my connection, thanks."),
				}},
			PrivateLinkServiceProxies: []*armdeviceupdate.PrivateLinkServiceProxy{
				{
					GroupConnectivityInformation: []*armdeviceupdate.GroupConnectivityInformation{},
					ID:                           to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{privateEndpointConnectionProxyId}/privateLinkServiceProxies/{privateEndpointConnectionProxyId}"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80065490402157d0df0dd37ab347c651b22eb576/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_PrivateEndpointUpdate.json
func ExamplePrivateEndpointConnectionProxiesClient_UpdatePrivateEndpointProperties() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceupdate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewPrivateEndpointConnectionProxiesClient().UpdatePrivateEndpointProperties(ctx, "test-rg", "contoso", "peexample01", armdeviceupdate.PrivateEndpointUpdate{
		ID:                      to.Ptr("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}"),
		ImmutableResourceID:     to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}"),
		ImmutableSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
		Location:                to.Ptr("westus2"),
		VnetTrafficTag:          to.Ptr("12345678"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80065490402157d0df0dd37ab347c651b22eb576/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_Get.json
func ExamplePrivateEndpointConnectionProxiesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceupdate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewPrivateEndpointConnectionProxiesClient().Get(ctx, "test-rg", "contoso", "peexample01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.PrivateEndpointConnectionProxy = armdeviceupdate.PrivateEndpointConnectionProxy{
	// 	RemotePrivateEndpoint: &armdeviceupdate.RemotePrivateEndpoint{
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}"),
	// 		ImmutableResourceID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}"),
	// 		ImmutableSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
	// 		Location: to.Ptr("westus2"),
	// 		ManualPrivateLinkServiceConnections: []*armdeviceupdate.PrivateLinkServiceConnection{
	// 			{
	// 				Name: to.Ptr("{plsConnectionName}"),
	// 				GroupIDs: []*string{
	// 					to.Ptr("DeviceUpdate")},
	// 					RequestMessage: to.Ptr("Please approve my connection, thanks."),
	// 			}},
	// 			PrivateLinkServiceProxies: []*armdeviceupdate.PrivateLinkServiceProxy{
	// 				{
	// 					GroupConnectivityInformation: []*armdeviceupdate.GroupConnectivityInformation{
	// 						{
	// 							GroupID: to.Ptr("DeviceUpdate"),
	// 							MemberName: to.Ptr("adu"),
	// 					}},
	// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{privateEndpointConnectionProxyId}/privateLinkServiceProxies/{privateEndpointConnectionProxyId}"),
	// 			}},
	// 		},
	// 		Name: to.Ptr("peexample01"),
	// 		Type: to.Ptr("Microsoft.DeviceUpdate/accounts/privateEndpointConnectionProxies"),
	// 		ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.DeviceUpdate/accounts/contoso/privateEndpointConnectionProxies/peexample01"),
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80065490402157d0df0dd37ab347c651b22eb576/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_CreateOrUpdate.json
func ExamplePrivateEndpointConnectionProxiesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceupdate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionProxiesClient().BeginCreateOrUpdate(ctx, "test-rg", "contoso", "peexample01", armdeviceupdate.PrivateEndpointConnectionProxy{
		RemotePrivateEndpoint: &armdeviceupdate.RemotePrivateEndpoint{
			ID:                      to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}"),
			ImmutableResourceID:     to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{peName}"),
			ImmutableSubscriptionID: to.Ptr("00000000-0000-0000-0000-000000000000"),
			Location:                to.Ptr("westus2"),
			ManualPrivateLinkServiceConnections: []*armdeviceupdate.PrivateLinkServiceConnection{
				{
					Name: to.Ptr("{privateEndpointConnectionProxyId}"),
					GroupIDs: []*string{
						to.Ptr("DeviceUpdate")},
					RequestMessage: to.Ptr("Please approve my connection, thanks."),
				}},
			PrivateLinkServiceProxies: []*armdeviceupdate.PrivateLinkServiceProxy{
				{
					GroupConnectivityInformation: []*armdeviceupdate.GroupConnectivityInformation{},
					ID:                           to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/test-rg/providers/Microsoft.Network/privateEndpoints/{privateEndpointConnectionProxyId}/privateLinkServiceProxies/{privateEndpointConnectionProxyId}"),
				}},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/80065490402157d0df0dd37ab347c651b22eb576/specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/PrivateEndpointConnectionProxies/PrivateEndpointConnectionProxy_Delete.json
func ExamplePrivateEndpointConnectionProxiesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdeviceupdate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewPrivateEndpointConnectionProxiesClient().BeginDelete(ctx, "test-rg", "contoso", "peexample01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
