//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstoragemover_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagemover/armstoragemover"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/241e964afe675a7be98aa6a2e171a3c5f830816c/specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-03-01/examples/Operations_List.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstoragemover.NewOperationsClient(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(nil)
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
		// page.OperationListResult = armstoragemover.OperationListResult{
		// 	Value: []*armstoragemover.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.StorageMover/storageMovers/read"),
		// 			Display: &armstoragemover.OperationDisplay{
		// 				Description: to.Ptr("Gets or Lists existing StorageMover resource(s)."),
		// 				Operation: to.Ptr("Get or List StorageMover resource(s)."),
		// 				Provider: to.Ptr("Microsoft StorageMover"),
		// 				Resource: to.Ptr("StorageMovers"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorageMover/storageMovers/write"),
		// 			Display: &armstoragemover.OperationDisplay{
		// 				Description: to.Ptr("Creates or Updates StorageMover resource."),
		// 				Operation: to.Ptr("Create or Update StorageMover resource."),
		// 				Provider: to.Ptr("Microsoft StorageMover"),
		// 				Resource: to.Ptr("StorageMovers"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.StorageMover/storageMovers/delete"),
		// 			Display: &armstoragemover.OperationDisplay{
		// 				Description: to.Ptr("Deletes StorageMover resource."),
		// 				Operation: to.Ptr("Delete StorageMover resource."),
		// 				Provider: to.Ptr("Microsoft StorageMover"),
		// 				Resource: to.Ptr("StorageMovers"),
		// 			},
		// 			IsDataAction: to.Ptr(false),
		// 	}},
		// }
	}
}
