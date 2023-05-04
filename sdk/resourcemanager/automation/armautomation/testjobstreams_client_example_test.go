//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armautomation_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/getTestJobStream.json
func ExampleTestJobStreamsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTestJobStreamsClient().Get(ctx, "mygroup", "ContoseAutomationAccount", "Get-AzureVMTutorial", "851b2101-686f-40e2-8a4b-5b8df08afbd1_00636535684910693884_00000000000000000001", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.JobStream = armautomation.JobStream{
	// 	ID: to.Ptr("/subscriptions/51766542-3ed7-4a72-a187-0c8ab644ddab/resourcegroups/mygroup/providers/Microsoft.Automation/automationAccounts/ContoseAutomationAccount/runbooks/foo/draft/testJob/streams/851b2101-686f-40e2-8a4b-5b8df08afbd1_00636535684910693884_00000000000000000001"),
	// 	Properties: &armautomation.JobStreamProperties{
	// 		JobStreamID: to.Ptr("851b2101-686f-40e2-8a4b-5b8df08afbd1:00636535684910693884:00000000000000000001"),
	// 		StreamText: to.Ptr(""),
	// 		StreamType: to.Ptr(armautomation.JobStreamTypeOutput),
	// 		Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-07T02:48:11.0693884+00:00"); return t}()),
	// 		Value: map[string]any{
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/main/specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/listTestJobStreamsByJob.json
func ExampleTestJobStreamsClient_NewListByTestJobPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armautomation.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTestJobStreamsClient().NewListByTestJobPager("mygroup", "ContoseAutomationAccount", "Get-AzureVMTutorial", &armautomation.TestJobStreamsClientListByTestJobOptions{Filter: nil})
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
		// page.JobStreamListResult = armautomation.JobStreamListResult{
		// 	Value: []*armautomation.JobStream{
		// 		{
		// 			ID: to.Ptr("/subscriptions/51766542-3ed7-4a72-a187-0c8ab644ddab/resourcegroups/mygroup/providers/Microsoft.Automation/automationAccounts/ContoseAutomationAccount/jobs/jobName/streams/24456a8a-2857-4af6-932c-3455f38bd05e_00636535675981232703_00000000000000000001"),
		// 			Properties: &armautomation.JobStreamProperties{
		// 				JobStreamID: to.Ptr("24456a8a-2857-4af6-932c-3455f38bd05e_00636535675981232703_00000000000000000001"),
		// 				StreamType: to.Ptr(armautomation.JobStreamTypeOutput),
		// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-07T02:33:18.1232703+00:00"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			ID: to.Ptr("/subscriptions/51766542-3ed7-4a72-a187-0c8ab644ddab/resourcegroups/mygroup/providers/Microsoft.Automation/automationAccounts/ContoseAutomationAccount/jobs/jobName/streams/24456a8a-2857-4af6-932c-3455f38bd05e_00636535675984691350_00000000000000000002"),
		// 			Properties: &armautomation.JobStreamProperties{
		// 				JobStreamID: to.Ptr("24456a8a-2857-4af6-932c-3455f38bd05e_00636535675984691350_00000000000000000002"),
		// 				StreamType: to.Ptr(armautomation.JobStreamTypeOutput),
		// 				Time: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-02-07T02:33:18.469135+00:00"); return t}()),
		// 			},
		// 	}},
		// }
	}
}
