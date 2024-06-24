//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armselfhelp_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/selfhelp/armselfhelp/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/ec7ee8842bf615c2f0354bf8b5b8725fdac9454a/specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/ListDiscoverySolutionsAtTenantScope.json
func ExampleDiscoverySolutionClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armselfhelp.NewClientFactory(cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewDiscoverySolutionClient().NewListPager(&armselfhelp.DiscoverySolutionClientListOptions{Filter: to.Ptr("ProblemClassificationId eq 'SampleProblemClassificationId1'"),
		Skiptoken: nil,
	})
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
		// page.DiscoveryResponse = armselfhelp.DiscoveryResponse{
		// 	Value: []*armselfhelp.SolutionMetadataResource{
		// 		{
		// 			Name: to.Ptr("SampleProblemClassificationId1"),
		// 			Type: to.Ptr("Microsoft.Help/discoverySolutions"),
		// 			ID: to.Ptr("/providers/microsoft.help/discoverySolutions/SampleProblemClassificationId1"),
		// 			Properties: &armselfhelp.Solutions{
		// 				Solutions: []*armselfhelp.SolutionMetadataProperties{
		// 					{
		// 						Description: to.Ptr("This is an azure solution to troubleshoot subscription issues."),
		// 						RequiredInputs: []*string{
		// 							to.Ptr("SubscriptionId")},
		// 							SolutionID: to.Ptr("SampleSolutionId1"),
		// 							SolutionType: to.Ptr(armselfhelp.SolutionTypeDiagnostics),
		// 					}},
		// 				},
		// 		}},
		// 	}
	}
}
