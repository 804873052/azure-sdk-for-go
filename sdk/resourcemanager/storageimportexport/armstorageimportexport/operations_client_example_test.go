//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armstorageimportexport_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storageimportexport/armstorageimportexport"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/storageimportexport/resource-manager/Microsoft.ImportExport/preview/2021-01-01/examples/ListOperations.json
func ExampleOperationsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armstorageimportexport.NewClientFactory("<subscription-id>", nil, cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewOperationsClient().NewListPager(nil)
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
		// page.ListOperationsResponse = armstorageimportexport.ListOperationsResponse{
		// 	Value: []*armstorageimportexport.Operation{
		// 		{
		// 			Name: to.Ptr("Microsoft.ImportExport/locations/read"),
		// 			Display: &armstorageimportexport.OperationDisplay{
		// 				Description: to.Ptr("Gets the properties for the specified location or returns the list of locations."),
		// 				Operation: to.Ptr("Get or List Locations"),
		// 				Provider: to.Ptr("Microsoft Import/Export"),
		// 				Resource: to.Ptr("Locations"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ImportExport/jobs/write"),
		// 			Display: &armstorageimportexport.OperationDisplay{
		// 				Description: to.Ptr("Creates a job with the specified parameters or update the properties or tags for the specified job."),
		// 				Operation: to.Ptr("Create or Update Job"),
		// 				Provider: to.Ptr("Microsoft Import/Export"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ImportExport/jobs/read"),
		// 			Display: &armstorageimportexport.OperationDisplay{
		// 				Description: to.Ptr("Gets the properties for the specified job or returns the list of jobs."),
		// 				Operation: to.Ptr("Get or List Jobs"),
		// 				Provider: to.Ptr("Microsoft Import/Export"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ImportExport/jobs/listBitLockerKeys/action"),
		// 			Display: &armstorageimportexport.OperationDisplay{
		// 				Description: to.Ptr("Gets the BitLocker keys for the specified job."),
		// 				Operation: to.Ptr("List BitLocker Keys"),
		// 				Provider: to.Ptr("Microsoft Import/Export"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ImportExport/jobs/delete"),
		// 			Display: &armstorageimportexport.OperationDisplay{
		// 				Description: to.Ptr("Deletes an existing job."),
		// 				Operation: to.Ptr("Delete Job"),
		// 				Provider: to.Ptr("Microsoft Import/Export"),
		// 				Resource: to.Ptr("Jobs"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("Microsoft.ImportExport/register/action"),
		// 			Display: &armstorageimportexport.OperationDisplay{
		// 				Description: to.Ptr("Registers the subscription for the import/export resource provider and enables the creation of import/export jobs."),
		// 				Operation: to.Ptr("Registers the Import/Export Resource Provider"),
		// 				Provider: to.Ptr("Microsoft Import/Export"),
		// 				Resource: to.Ptr("Import/Export Resource Provider"),
		// 			},
		// 	}},
		// }
	}
}
