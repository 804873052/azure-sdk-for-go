//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armservicenetworking_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicenetworking/armservicenetworking"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/71121282e39bccae590462648e77bca283df6d2b/specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/AssociationsGet.json
func ExampleAssociationsInterfaceClient_NewListByTrafficControllerPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicenetworking.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewAssociationsInterfaceClient().NewListByTrafficControllerPager("rg1", "TC1", nil)
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
		// page.AssociationListResult = armservicenetworking.AssociationListResult{
		// 	Value: []*armservicenetworking.Association{
		// 		{
		// 			Name: to.Ptr("associatedvnet-2"),
		// 			Type: to.Ptr("Microsoft.ServiceNetworking/trafficControllers/associations"),
		// 			ID: to.Ptr("resourceUriAsString"),
		// 			Location: to.Ptr("West US"),
		// 			Properties: &armservicenetworking.AssociationProperties{
		// 				AssociationType: to.Ptr("subnets"),
		// 				ProvisioningState: to.Ptr(armservicenetworking.ProvisioningStateSucceeded),
		// 				Subnet: &armservicenetworking.AssociationSubnet{
		// 					ID: to.Ptr("subnetFullRef"),
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/71121282e39bccae590462648e77bca283df6d2b/specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/AssociationGet.json
func ExampleAssociationsInterfaceClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicenetworking.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssociationsInterfaceClient().Get(ctx, "rg1", "TC1", "associatedvnet-2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Association = armservicenetworking.Association{
	// 	Name: to.Ptr("associatedvnet-2"),
	// 	Type: to.Ptr("Microsoft.ServiceNetworking/trafficControllers/associations"),
	// 	ID: to.Ptr("resourceUriAsString"),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armservicenetworking.AssociationProperties{
	// 		AssociationType: to.Ptr("subnets"),
	// 		ProvisioningState: to.Ptr(armservicenetworking.ProvisioningStateSucceeded),
	// 		Subnet: &armservicenetworking.AssociationSubnet{
	// 			ID: to.Ptr("subnetFullRef"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/71121282e39bccae590462648e77bca283df6d2b/specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/AssociationPut.json
func ExampleAssociationsInterfaceClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicenetworking.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAssociationsInterfaceClient().BeginCreateOrUpdate(ctx, "rg1", "TC1", "associatedvnet-1", armservicenetworking.Association{
		Location: to.Ptr("West US"),
		Properties: &armservicenetworking.AssociationProperties{
			AssociationType: to.Ptr("subnets"),
			Subnet: &armservicenetworking.AssociationSubnet{
				ID: to.Ptr("subnetFullRef"),
			},
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
	// res.Association = armservicenetworking.Association{
	// 	Name: to.Ptr("associatedvnet-1"),
	// 	Type: to.Ptr("Microsoft.ServiceNetworking/trafficControllers/associations"),
	// 	ID: to.Ptr("resourceUriAsString  "),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armservicenetworking.AssociationProperties{
	// 		AssociationType: to.Ptr("subnets"),
	// 		ProvisioningState: to.Ptr(armservicenetworking.ProvisioningStateSucceeded),
	// 		Subnet: &armservicenetworking.AssociationSubnet{
	// 			ID: to.Ptr("subnetFullRef"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/71121282e39bccae590462648e77bca283df6d2b/specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/AssociationPatch.json
func ExampleAssociationsInterfaceClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicenetworking.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewAssociationsInterfaceClient().Update(ctx, "rg1", "TC1", "associatedvnet-1", armservicenetworking.AssociationUpdate{
		Properties: &armservicenetworking.AssociationUpdateProperties{
			AssociationType: to.Ptr("subnets"),
			Subnet: &armservicenetworking.AssociationSubnet{
				ID: to.Ptr("subnetFullRef"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Association = armservicenetworking.Association{
	// 	Name: to.Ptr("associatedvnet-1"),
	// 	Type: to.Ptr("Microsoft.ServiceNetworking/trafficControllers/associations"),
	// 	ID: to.Ptr("resourceUriAsString  "),
	// 	Location: to.Ptr("West US"),
	// 	Properties: &armservicenetworking.AssociationProperties{
	// 		AssociationType: to.Ptr("subnets"),
	// 		Subnet: &armservicenetworking.AssociationSubnet{
	// 			ID: to.Ptr("subnetFullRef"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/71121282e39bccae590462648e77bca283df6d2b/specification/servicenetworking/resource-manager/Microsoft.ServiceNetworking/cadl/examples/AssociationDelete.json
func ExampleAssociationsInterfaceClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicenetworking.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewAssociationsInterfaceClient().BeginDelete(ctx, "rg1", "TC1", "associatedvnet-2", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
