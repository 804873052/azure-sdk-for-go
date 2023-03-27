//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcontainers_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcontainers/armappcontainers/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/212686c8383679e034b19143e13cbeb5a40ab454/specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/ConnectedEnvironmentsStorages_List.json
func ExampleConnectedEnvironmentsStoragesClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectedEnvironmentsStoragesClient().List(ctx, "examplerg", "managedEnv", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectedEnvironmentStoragesCollection = armappcontainers.ConnectedEnvironmentStoragesCollection{
	// 	Value: []*armappcontainers.ConnectedEnvironmentStorage{
	// 		{
	// 			Name: to.Ptr("jlaw-demo1"),
	// 			Type: to.Ptr("Microsoft.App/connectedEnvironments/storages"),
	// 			ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/connectedEnvironments/managedEnv/storages/jlaw-demo1"),
	// 			Properties: &armappcontainers.ConnectedEnvironmentStorageProperties{
	// 				AzureFile: &armappcontainers.AzureFileProperties{
	// 					AccessMode: to.Ptr(armappcontainers.AccessModeReadOnly),
	// 					AccountName: to.Ptr("account1"),
	// 					ShareName: to.Ptr("share1"),
	// 				},
	// 			},
	// 	}},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/212686c8383679e034b19143e13cbeb5a40ab454/specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/ConnectedEnvironmentsStorages_Get.json
func ExampleConnectedEnvironmentsStoragesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectedEnvironmentsStoragesClient().Get(ctx, "examplerg", "env", "jlaw-demo1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectedEnvironmentStorage = armappcontainers.ConnectedEnvironmentStorage{
	// 	Name: to.Ptr("jlaw-demo1"),
	// 	Type: to.Ptr("Microsoft.App/connectedEnvironments/storages"),
	// 	ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/connectedEnvironments/env/storages/jlaw-demo1"),
	// 	Properties: &armappcontainers.ConnectedEnvironmentStorageProperties{
	// 		AzureFile: &armappcontainers.AzureFileProperties{
	// 			AccessMode: to.Ptr(armappcontainers.AccessModeReadOnly),
	// 			AccountName: to.Ptr("account1"),
	// 			ShareName: to.Ptr("share1"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/212686c8383679e034b19143e13cbeb5a40ab454/specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/ConnectedEnvironmentsStorages_CreateOrUpdate.json
func ExampleConnectedEnvironmentsStoragesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewConnectedEnvironmentsStoragesClient().CreateOrUpdate(ctx, "examplerg", "env", "jlaw-demo1", armappcontainers.ConnectedEnvironmentStorage{
		Properties: &armappcontainers.ConnectedEnvironmentStorageProperties{
			AzureFile: &armappcontainers.AzureFileProperties{
				AccessMode:  to.Ptr(armappcontainers.AccessModeReadOnly),
				AccountKey:  to.Ptr("key"),
				AccountName: to.Ptr("account1"),
				ShareName:   to.Ptr("share1"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ConnectedEnvironmentStorage = armappcontainers.ConnectedEnvironmentStorage{
	// 	Name: to.Ptr("jlaw-demo1"),
	// 	Type: to.Ptr("Microsoft.App/connectedEnvironments/storages"),
	// 	ID: to.Ptr("/subscriptions/8efdecc5-919e-44eb-b179-915dca89ebf9/resourceGroups/examplerg/providers/Microsoft.App/connectedEnvironments/env/storages/jlaw-demo1"),
	// 	Properties: &armappcontainers.ConnectedEnvironmentStorageProperties{
	// 		AzureFile: &armappcontainers.AzureFileProperties{
	// 			AccessMode: to.Ptr(armappcontainers.AccessModeReadOnly),
	// 			AccountName: to.Ptr("account1"),
	// 			ShareName: to.Ptr("share1"),
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/212686c8383679e034b19143e13cbeb5a40ab454/specification/app/resource-manager/Microsoft.App/preview/2022-06-01-preview/examples/ConnectedEnvironmentsStorages_Delete.json
func ExampleConnectedEnvironmentsStoragesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armappcontainers.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewConnectedEnvironmentsStoragesClient().Delete(ctx, "examplerg", "env", "jlaw-demo1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}
