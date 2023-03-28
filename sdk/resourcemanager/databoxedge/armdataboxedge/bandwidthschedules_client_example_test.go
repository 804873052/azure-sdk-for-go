//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdataboxedge_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/databoxedge/armdataboxedge/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/BandwidthScheduleGetAllInDevice.json
func ExampleBandwidthSchedulesClient_NewListByDataBoxEdgeDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewBandwidthSchedulesClient().NewListByDataBoxEdgeDevicePager("testedgedevice", "GroupForEdgeAutomation", nil)
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
		// page.BandwidthSchedulesList = armdataboxedge.BandwidthSchedulesList{
		// 	Value: []*armdataboxedge.BandwidthSchedule{
		// 		{
		// 			Name: to.Ptr("bandwidth-1"),
		// 			Type: to.Ptr("dataBoxEdgeDevices/bandwidthSchedules"),
		// 			ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/bandwidthSchedules/bandwidth-1"),
		// 			Properties: &armdataboxedge.BandwidthScheduleProperties{
		// 				Days: []*armdataboxedge.DayOfWeek{
		// 					to.Ptr(armdataboxedge.DayOfWeekSunday),
		// 					to.Ptr(armdataboxedge.DayOfWeekMonday)},
		// 					RateInMbps: to.Ptr[int32](100),
		// 					Start: to.Ptr("00:00:00"),
		// 					Stop: to.Ptr("13:59:00"),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/BandwidthScheduleGet.json
func ExampleBandwidthSchedulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewBandwidthSchedulesClient().Get(ctx, "testedgedevice", "bandwidth-1", "GroupForEdgeAutomation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.BandwidthSchedule = armdataboxedge.BandwidthSchedule{
	// 	Name: to.Ptr("bandwidth-1"),
	// 	Type: to.Ptr("dataBoxEdgeDevices/bandwidthSchedules"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/bandwidthSchedules/bandwidth-1"),
	// 	Properties: &armdataboxedge.BandwidthScheduleProperties{
	// 		Days: []*armdataboxedge.DayOfWeek{
	// 			to.Ptr(armdataboxedge.DayOfWeekSunday),
	// 			to.Ptr(armdataboxedge.DayOfWeekMonday)},
	// 			RateInMbps: to.Ptr[int32](100),
	// 			Start: to.Ptr("00:00:00"),
	// 			Stop: to.Ptr("13:59:00"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/BandwidthSchedulePut.json
func ExampleBandwidthSchedulesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBandwidthSchedulesClient().BeginCreateOrUpdate(ctx, "testedgedevice", "bandwidth-1", "GroupForEdgeAutomation", armdataboxedge.BandwidthSchedule{
		Properties: &armdataboxedge.BandwidthScheduleProperties{
			Days: []*armdataboxedge.DayOfWeek{
				to.Ptr(armdataboxedge.DayOfWeekSunday),
				to.Ptr(armdataboxedge.DayOfWeekMonday)},
			RateInMbps: to.Ptr[int32](100),
			Start:      to.Ptr("0:0:0"),
			Stop:       to.Ptr("13:59:0"),
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
	// res.BandwidthSchedule = armdataboxedge.BandwidthSchedule{
	// 	Name: to.Ptr("bandwidth-1"),
	// 	Type: to.Ptr("dataBoxEdgeDevices/bandwidthSchedules"),
	// 	ID: to.Ptr("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/bandwidthSchedules/bandwidth-1"),
	// 	Properties: &armdataboxedge.BandwidthScheduleProperties{
	// 		Days: []*armdataboxedge.DayOfWeek{
	// 			to.Ptr(armdataboxedge.DayOfWeekSunday),
	// 			to.Ptr(armdataboxedge.DayOfWeekMonday)},
	// 			RateInMbps: to.Ptr[int32](100),
	// 			Start: to.Ptr("00:00:00"),
	// 			Stop: to.Ptr("13:59:00"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2022-03-01/examples/BandwidthScheduleDelete.json
func ExampleBandwidthSchedulesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdataboxedge.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewBandwidthSchedulesClient().BeginDelete(ctx, "testedgedevice", "bandwidth-1", "GroupForEdgeAutomation", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
