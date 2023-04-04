//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armextendedlocation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/extendedlocation/armextendedlocation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsListOperations.json
func ExampleCustomLocationsClient_NewListOperationsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextendedlocation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomLocationsClient().NewListOperationsPager(nil)
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
		// page.CustomLocationOperationsList = armextendedlocation.CustomLocationOperationsList{
		// 	Value: []*armextendedlocation.CustomLocationOperation{
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/operations/read"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Gets list of Available Operations for Custom Locations"),
		// 				Operation: to.Ptr("List Available Operations for Custom Locations"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Operations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/register/action"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Registers the subscription for Custom Location resource provider and enables the creation of Custom Location."),
		// 				Operation: to.Ptr("Registers the Custom Location Resource Provider"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Custom Locations Resource Provider"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/unregister/action"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("UnRegisters the subscription for Custom Location resource provider and disables the creation of Custom Location."),
		// 				Operation: to.Ptr("UnRegisters the Custom Location Resource Provider"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Custom Locations Resource Provider"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/customLocations/read"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Gets an Custom Location resource"),
		// 				Operation: to.Ptr("Get Custom Location"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Custom Locations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/customLocations/write"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Creates or Updates Custom Location resource"),
		// 				Operation: to.Ptr("Create or Update Custom Location"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Custom Locations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/customLocations/deploy/action"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Deploy permissions to a Custom Location resource"),
		// 				Operation: to.Ptr("Deploy permissions to Custom Location"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Custom Locations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/customLocations/delete"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Deletes Custom Location resource"),
		// 				Operation: to.Ptr("Delete Custom Location"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Custom Locations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/customLocations/enabledresourcetypes/read"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Gets EnabledResourceTypes for a Custom Location resource"),
		// 				Operation: to.Ptr("Get EnabledResourceTypes for Custom Location"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Custom Locations"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/locations/operationsstatus/read"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Get result of Custom Location operation"),
		// 				Operation: to.Ptr("Get status of Custom Location operation"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Custom Location Operation Status"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ExtendedLocation/locations/operationresults/read"),
		// 			Display: &armextendedlocation.CustomLocationOperationValueDisplay{
		// 				Description: to.Ptr("Get result of Custom Location operation"),
		// 				Operation: to.Ptr("Get the status of Custom Location operation"),
		// 				Provider: to.Ptr("Microsoft.ExtendedLocation"),
		// 				Resource: to.Ptr("Custom Location Operation Result"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 			Origin: to.Ptr("user,system"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsListBySubscription.json
func ExampleCustomLocationsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextendedlocation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomLocationsClient().NewListBySubscriptionPager(nil)
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
		// page.CustomLocationListResult = armextendedlocation.CustomLocationListResult{
		// 	Value: []*armextendedlocation.CustomLocation{
		// 		{
		// 			Name: to.Ptr("customLocation01"),
		// 			Type: to.Ptr("Microsoft.ExtendedLocation/customLocations"),
		// 			ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ExtendedLocation/"),
		// 			Location: to.Ptr("West US"),
		// 			Identity: &armextendedlocation.Identity{
		// 				Type: to.Ptr(armextendedlocation.ResourceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 				TenantID: to.Ptr("111111-1111-1111-1111-111111111111"),
		// 			},
		// 			Properties: &armextendedlocation.CustomLocationProperties{
		// 				Authentication: &armextendedlocation.CustomLocationPropertiesAuthentication{
		// 					Type: to.Ptr("KubeConfig"),
		// 				},
		// 				ClusterExtensionIDs: []*string{
		// 					to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01/Microsoft.KubernetesConfiguration/clusterExtensions/fooExtension")},
		// 					DisplayName: to.Ptr("customLocationLocation01"),
		// 					HostResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01"),
		// 					Namespace: to.Ptr("namespace01"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 				},
		// 				SystemData: &armextendedlocation.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("customLocation02"),
		// 				Type: to.Ptr("Microsoft.ExtendedLocation/customLocations"),
		// 				ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ExtendedLocation/"),
		// 				Location: to.Ptr("West US"),
		// 				Identity: &armextendedlocation.Identity{
		// 					Type: to.Ptr(armextendedlocation.ResourceIdentityTypeSystemAssigned),
		// 					PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 					TenantID: to.Ptr("111111-1111-1111-1111-111111111111"),
		// 				},
		// 				Properties: &armextendedlocation.CustomLocationProperties{
		// 					Authentication: &armextendedlocation.CustomLocationPropertiesAuthentication{
		// 						Type: to.Ptr("KubeConfig"),
		// 					},
		// 					ClusterExtensionIDs: []*string{
		// 						to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster02/Microsoft.KubernetesConfiguration/clusterExtensions/fooExtension")},
		// 						DisplayName: to.Ptr("customLocationLocation02"),
		// 						HostResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster02"),
		// 						Namespace: to.Ptr("namespace02"),
		// 						ProvisioningState: to.Ptr("Succeeded"),
		// 					},
		// 					SystemData: &armextendedlocation.SystemData{
		// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-24T18:53:29.0928001Z"); return t}()),
		// 						CreatedBy: to.Ptr("string"),
		// 						CreatedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
		// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-24T18:53:29.0928001Z"); return t}()),
		// 						LastModifiedBy: to.Ptr("string"),
		// 						LastModifiedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsListByResourceGroup.json
func ExampleCustomLocationsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextendedlocation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomLocationsClient().NewListByResourceGroupPager("testresourcegroup", nil)
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
		// page.CustomLocationListResult = armextendedlocation.CustomLocationListResult{
		// 	Value: []*armextendedlocation.CustomLocation{
		// 		{
		// 			Name: to.Ptr("customLocation01"),
		// 			Type: to.Ptr("Microsoft.ExtendedLocation/customLocations"),
		// 			ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ExtendedLocation/"),
		// 			Location: to.Ptr("West US"),
		// 			Identity: &armextendedlocation.Identity{
		// 				Type: to.Ptr(armextendedlocation.ResourceIdentityTypeSystemAssigned),
		// 				PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 				TenantID: to.Ptr("111111-1111-1111-1111-111111111111"),
		// 			},
		// 			Properties: &armextendedlocation.CustomLocationProperties{
		// 				Authentication: &armextendedlocation.CustomLocationPropertiesAuthentication{
		// 					Type: to.Ptr("KubeConfig"),
		// 				},
		// 				ClusterExtensionIDs: []*string{
		// 					to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01/Microsoft.KubernetesConfiguration/clusterExtensions/fooExtension")},
		// 					DisplayName: to.Ptr("customLocationLocation01"),
		// 					HostResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01"),
		// 					Namespace: to.Ptr("namespace01"),
		// 					ProvisioningState: to.Ptr("Succeeded"),
		// 				},
		// 				SystemData: &armextendedlocation.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 					CreatedBy: to.Ptr("string"),
		// 					CreatedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("string"),
		// 					LastModifiedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("customLocation02"),
		// 				Type: to.Ptr("Microsoft.ExtendedLocation/customLocations"),
		// 				ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ExtendedLocation/"),
		// 				Location: to.Ptr("West US"),
		// 				Identity: &armextendedlocation.Identity{
		// 					Type: to.Ptr(armextendedlocation.ResourceIdentityTypeSystemAssigned),
		// 					PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
		// 					TenantID: to.Ptr("111111-1111-1111-1111-111111111111"),
		// 				},
		// 				Properties: &armextendedlocation.CustomLocationProperties{
		// 					Authentication: &armextendedlocation.CustomLocationPropertiesAuthentication{
		// 						Type: to.Ptr("KubeConfig"),
		// 					},
		// 					ClusterExtensionIDs: []*string{
		// 						to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster02/Microsoft.KubernetesConfiguration/clusterExtensions/fooExtension")},
		// 						DisplayName: to.Ptr("customLocationLocation02"),
		// 						HostResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster02"),
		// 						Namespace: to.Ptr("namespace02"),
		// 						ProvisioningState: to.Ptr("Succeeded"),
		// 					},
		// 					SystemData: &armextendedlocation.SystemData{
		// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-24T18:53:29.0928001Z"); return t}()),
		// 						CreatedBy: to.Ptr("string"),
		// 						CreatedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
		// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-24T18:53:29.0928001Z"); return t}()),
		// 						LastModifiedBy: to.Ptr("string"),
		// 						LastModifiedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsGet.json
func ExampleCustomLocationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextendedlocation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomLocationsClient().Get(ctx, "testresourcegroup", "customLocation01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomLocation = armextendedlocation.CustomLocation{
	// 	Name: to.Ptr("customLocation01"),
	// 	Type: to.Ptr("Microsoft.ExtendedLocation/customLocations"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/customLocation01"),
	// 	Location: to.Ptr("West US"),
	// 	Identity: &armextendedlocation.Identity{
	// 		Type: to.Ptr(armextendedlocation.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		TenantID: to.Ptr("111111-1111-1111-1111-111111111111"),
	// 	},
	// 	Properties: &armextendedlocation.CustomLocationProperties{
	// 		Authentication: &armextendedlocation.CustomLocationPropertiesAuthentication{
	// 			Type: to.Ptr("KubeConfig"),
	// 		},
	// 		ClusterExtensionIDs: []*string{
	// 			to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01/Microsoft.KubernetesConfiguration/clusterExtensions/fooExtension")},
	// 			DisplayName: to.Ptr("customLocationLocation01"),
	// 			HostResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01"),
	// 			Namespace: to.Ptr("namespace01"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 		SystemData: &armextendedlocation.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-24T18:53:29.0928001Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-24T18:53:29.0928001Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsCreate_Update.json
func ExampleCustomLocationsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextendedlocation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCustomLocationsClient().BeginCreateOrUpdate(ctx, "testresourcegroup", "customLocation01", armextendedlocation.CustomLocation{
		Location: to.Ptr("West US"),
		Identity: &armextendedlocation.Identity{
			Type: to.Ptr(armextendedlocation.ResourceIdentityTypeSystemAssigned),
		},
		Properties: &armextendedlocation.CustomLocationProperties{
			Authentication: &armextendedlocation.CustomLocationPropertiesAuthentication{
				Type:  to.Ptr("KubeConfig"),
				Value: to.Ptr("<base64 KubeConfig>"),
			},
			ClusterExtensionIDs: []*string{
				to.Ptr("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kubernetes/connectedCluster/someCluster/Microsoft.KubernetesConfiguration/clusterExtensions/fooExtension")},
			DisplayName:    to.Ptr("customLocationLocation01"),
			HostResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01"),
			Namespace:      to.Ptr("namespace01"),
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
	// res.CustomLocation = armextendedlocation.CustomLocation{
	// 	Name: to.Ptr("customLocation01"),
	// 	Type: to.Ptr("Microsoft.ExtendedLocation/customLocations"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/customLocation01"),
	// 	Location: to.Ptr("West US"),
	// 	Identity: &armextendedlocation.Identity{
	// 		Type: to.Ptr(armextendedlocation.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		TenantID: to.Ptr("111111-1111-1111-1111-111111111111"),
	// 	},
	// 	Properties: &armextendedlocation.CustomLocationProperties{
	// 		Authentication: &armextendedlocation.CustomLocationPropertiesAuthentication{
	// 			Type: to.Ptr("KubeConfig"),
	// 		},
	// 		ClusterExtensionIDs: []*string{
	// 			to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01/Microsoft.KubernetesConfiguration/clusterExtensions/fooExtension")},
	// 			DisplayName: to.Ptr("customLocationLocation01"),
	// 			HostResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01"),
	// 			Namespace: to.Ptr("namespace01"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 		SystemData: &armextendedlocation.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsDelete.json
func ExampleCustomLocationsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextendedlocation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCustomLocationsClient().BeginDelete(ctx, "testresourcegroup", "customLocation01", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsPatch.json
func ExampleCustomLocationsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextendedlocation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCustomLocationsClient().Update(ctx, "testresourcegroup", "customLocation01", armextendedlocation.PatchableCustomLocations{
		Identity: &armextendedlocation.Identity{
			Type: to.Ptr(armextendedlocation.ResourceIdentityTypeSystemAssigned),
		},
		Properties: &armextendedlocation.CustomLocationProperties{
			ClusterExtensionIDs: []*string{
				to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01/Microsoft.KubernetesConfiguration/clusterExtensions/fooExtension"),
				to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01/Microsoft.KubernetesConfiguration/clusterExtensions/barExtension")},
		},
		Tags: map[string]*string{
			"archv3": to.Ptr(""),
			"tier":   to.Ptr("testing"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CustomLocation = armextendedlocation.CustomLocation{
	// 	Name: to.Ptr("customLocation01"),
	// 	Type: to.Ptr("Microsoft.ExtendedLocation/customLocations"),
	// 	ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/customLocation01"),
	// 	Location: to.Ptr("West US"),
	// 	Tags: map[string]*string{
	// 		"archv3": to.Ptr(""),
	// 		"tier": to.Ptr("testing"),
	// 	},
	// 	Identity: &armextendedlocation.Identity{
	// 		Type: to.Ptr(armextendedlocation.ResourceIdentityTypeSystemAssigned),
	// 		PrincipalID: to.Ptr("11111111-1111-1111-1111-111111111111"),
	// 		TenantID: to.Ptr("111111-1111-1111-1111-111111111111"),
	// 	},
	// 	Properties: &armextendedlocation.CustomLocationProperties{
	// 		Authentication: &armextendedlocation.CustomLocationPropertiesAuthentication{
	// 			Type: to.Ptr("KubeConfig"),
	// 		},
	// 		ClusterExtensionIDs: []*string{
	// 			to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01/Microsoft.KubernetesConfiguration/clusterExtensions/fooExtension"),
	// 			to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01/Microsoft.KubernetesConfiguration/clusterExtensions/barExtension")},
	// 			DisplayName: to.Ptr("customLocationLocation01"),
	// 			HostResourceID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testresourcegroup/providers/Microsoft.ContainerService/managedClusters/cluster01"),
	// 			Namespace: to.Ptr("namespace01"),
	// 			ProvisioningState: to.Ptr("Succeeded"),
	// 		},
	// 		SystemData: &armextendedlocation.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-24T18:53:29.0928001Z"); return t}()),
	// 			CreatedBy: to.Ptr("string"),
	// 			CreatedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-24T18:53:29.0928001Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("string"),
	// 			LastModifiedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsListEnabledResourceTypes.json
func ExampleCustomLocationsClient_NewListEnabledResourceTypesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armextendedlocation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCustomLocationsClient().NewListEnabledResourceTypesPager("testresourcegroup", "customLocation01", nil)
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
		// page.EnabledResourceTypesListResult = armextendedlocation.EnabledResourceTypesListResult{
		// 	Value: []*armextendedlocation.EnabledResourceType{
		// 		{
		// 			Name: to.Ptr("d016ecf26dae90594806aca3c1a6326c668357037f68103587edf2e657824737"),
		// 			Type: to.Ptr("Microsoft.ExtendedLocation/customLocations/enabledResourceTypes"),
		// 			ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/customLocation01/enabledResourceTypes/d016ecf26dae90594806aca3c1a6326c668357037f68103587edf2e657824737"),
		// 			Properties: &armextendedlocation.EnabledResourceTypeProperties{
		// 				ClusterExtensionID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/cldfe2econnectedcluster/providers/Microsoft.KubernetesConfiguration/extensions/vmware-extension"),
		// 				ExtensionType: to.Ptr("arc-vmware"),
		// 				TypesMetadata: []*armextendedlocation.EnabledResourceTypePropertiesTypesMetadataItem{
		// 					{
		// 						APIVersion: to.Ptr("2020-01-01-preview"),
		// 						ResourceProviderNamespace: to.Ptr("Microsoft.VMware"),
		// 						ResourceType: to.Ptr("virtualMachines"),
		// 					},
		// 					{
		// 						APIVersion: to.Ptr("2020-01-22-preview"),
		// 						ResourceProviderNamespace: to.Ptr("Microsoft.VMware"),
		// 						ResourceType: to.Ptr("virtualmachines"),
		// 				}},
		// 			},
		// 			SystemData: &armextendedlocation.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-04-24T18:53:29.0928001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("266e9d31e5be6be1e919574e25780d5783586d502f0b2cc422e0a228a34e00a6"),
		// 			Type: to.Ptr("Microsoft.ExtendedLocation/customLocations/enabledResourceTypes"),
		// 			ID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testrg/providers/Microsoft.ExtendedLocation/customLocations/customLocation01/enabledResourceTypes/266e9d31e5be6be1e919574e25780d5783586d502f0b2cc422e0a228a34e00a6"),
		// 			Properties: &armextendedlocation.EnabledResourceTypeProperties{
		// 				ClusterExtensionID: to.Ptr("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/testrg/providers/Microsoft.Kubernetes/connectedClusters/cldfe2econnectedcluster/providers/Microsoft.KubernetesConfiguration/extensions/cassandra-extension"),
		// 				ExtensionType: to.Ptr("cassandradatacentersoperator"),
		// 				TypesMetadata: []*armextendedlocation.EnabledResourceTypePropertiesTypesMetadataItem{
		// 					{
		// 						APIVersion: to.Ptr("2020-01-01-preview"),
		// 						ResourceProviderNamespace: to.Ptr("Microsoft.Cassandra"),
		// 						ResourceType: to.Ptr("cassandraDataCenters"),
		// 					},
		// 					{
		// 						APIVersion: to.Ptr("2020-01-22-preview"),
		// 						ResourceProviderNamespace: to.Ptr("Microsoft.Cassandra"),
		// 						ResourceType: to.Ptr("cassandrabackups"),
		// 				}},
		// 			},
		// 			SystemData: &armextendedlocation.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-24T18:53:29.0928001Z"); return t}()),
		// 				CreatedBy: to.Ptr("string"),
		// 				CreatedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-08-24T18:53:29.0928001Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("string"),
		// 				LastModifiedByType: to.Ptr(armextendedlocation.CreatedByTypeApplication),
		// 			},
		// 	}},
		// }
	}
}
