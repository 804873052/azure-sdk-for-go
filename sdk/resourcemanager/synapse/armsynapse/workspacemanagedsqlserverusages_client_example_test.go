//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsynapse_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/synapse/armsynapse"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/630ec444f8dd7c09b9cdd5fa99951f8a0d1ad41f/specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListWorkspaceManagedSqlServerUsages.json
func ExampleWorkspaceManagedSQLServerUsagesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armsynapse.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewWorkspaceManagedSQLServerUsagesClient().NewListPager("wsg-7398", "testWorkspace", nil)
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
		// page.ServerUsageListResult = armsynapse.ServerUsageListResult{
		// 	Value: []*armsynapse.ServerUsage{
		// 		{
		// 			Name: to.Ptr("server_dtu_quota"),
		// 			CurrentValue: to.Ptr[float64](0),
		// 			DisplayName: to.Ptr("Database Throughput Unit Quota"),
		// 			Limit: to.Ptr[float64](45000),
		// 			ResourceName: to.Ptr("testWorkspace"),
		// 			Unit: to.Ptr("DTUs"),
		// 		},
		// 		{
		// 			Name: to.Ptr("server_dtu_quota_current"),
		// 			CurrentValue: to.Ptr[float64](0),
		// 			DisplayName: to.Ptr("Database Throughput Unit Quota"),
		// 			Limit: to.Ptr[float64](45000),
		// 			ResourceName: to.Ptr("testWorkspace"),
		// 			Unit: to.Ptr("DTUs"),
		// 	}},
		// }
	}
}
