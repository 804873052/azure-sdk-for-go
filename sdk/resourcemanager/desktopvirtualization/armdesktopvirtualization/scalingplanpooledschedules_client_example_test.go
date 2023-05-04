//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdesktopvirtualization_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/desktopvirtualization/armdesktopvirtualization/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f339d469c0fe83466edfe295a7960c82ebecf4f/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2022-09-09/examples/ScalingPlanPooledSchedule_Get.json
func ExampleScalingPlanPooledSchedulesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScalingPlanPooledSchedulesClient().Get(ctx, "resourceGroup1", "scalingPlan1", "scalingPlanScheduleWeekdays1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ScalingPlanPooledSchedule = armdesktopvirtualization.ScalingPlanPooledSchedule{
	// 	Name: to.Ptr("scalingPlanScheduleWeekdays1"),
	// 	Type: to.Ptr("Microsoft.DesktopVirtualization/scalingPlans/pooledSchedules"),
	// 	ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/scalingPlans/scalingPlan1/pooledSchedules/scalingPlanScheduleWeekdays1"),
	// 	Properties: &armdesktopvirtualization.ScalingPlanPooledScheduleProperties{
	// 		DaysOfWeek: []*armdesktopvirtualization.DayOfWeek{
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekMonday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekTuesday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekWednesday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekThursday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekFriday)},
	// 			OffPeakLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
	// 			OffPeakStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](20),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 			PeakLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmBreadthFirst),
	// 			PeakStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](8),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 			RampDownCapacityThresholdPct: to.Ptr[int32](50),
	// 			RampDownForceLogoffUsers: to.Ptr(true),
	// 			RampDownLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
	// 			RampDownMinimumHostsPct: to.Ptr[int32](20),
	// 			RampDownNotificationMessage: to.Ptr("message"),
	// 			RampDownStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](18),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 			RampDownWaitTimeMinutes: to.Ptr[int32](30),
	// 			RampUpCapacityThresholdPct: to.Ptr[int32](80),
	// 			RampUpLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
	// 			RampUpMinimumHostsPct: to.Ptr[int32](20),
	// 			RampUpStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](6),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 		},
	// 		SystemData: &armdesktopvirtualization.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f339d469c0fe83466edfe295a7960c82ebecf4f/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2022-09-09/examples/ScalingPlanPooledSchedule_Create.json
func ExampleScalingPlanPooledSchedulesClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScalingPlanPooledSchedulesClient().Create(ctx, "resourceGroup1", "scalingPlan1", "scalingPlanScheduleWeekdays1", armdesktopvirtualization.ScalingPlanPooledSchedule{
		Properties: &armdesktopvirtualization.ScalingPlanPooledScheduleProperties{
			DaysOfWeek: []*armdesktopvirtualization.DayOfWeek{
				to.Ptr(armdesktopvirtualization.DayOfWeekMonday),
				to.Ptr(armdesktopvirtualization.DayOfWeekTuesday),
				to.Ptr(armdesktopvirtualization.DayOfWeekWednesday),
				to.Ptr(armdesktopvirtualization.DayOfWeekThursday),
				to.Ptr(armdesktopvirtualization.DayOfWeekFriday)},
			OffPeakLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
			OffPeakStartTime: &armdesktopvirtualization.Time{
				Hour:   to.Ptr[int32](20),
				Minute: to.Ptr[int32](0),
			},
			PeakLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmBreadthFirst),
			PeakStartTime: &armdesktopvirtualization.Time{
				Hour:   to.Ptr[int32](8),
				Minute: to.Ptr[int32](0),
			},
			RampDownCapacityThresholdPct:   to.Ptr[int32](50),
			RampDownForceLogoffUsers:       to.Ptr(true),
			RampDownLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
			RampDownMinimumHostsPct:        to.Ptr[int32](20),
			RampDownNotificationMessage:    to.Ptr("message"),
			RampDownStartTime: &armdesktopvirtualization.Time{
				Hour:   to.Ptr[int32](18),
				Minute: to.Ptr[int32](0),
			},
			RampDownWaitTimeMinutes:      to.Ptr[int32](30),
			RampUpCapacityThresholdPct:   to.Ptr[int32](80),
			RampUpLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
			RampUpMinimumHostsPct:        to.Ptr[int32](20),
			RampUpStartTime: &armdesktopvirtualization.Time{
				Hour:   to.Ptr[int32](6),
				Minute: to.Ptr[int32](0),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ScalingPlanPooledSchedule = armdesktopvirtualization.ScalingPlanPooledSchedule{
	// 	Name: to.Ptr("scalingPlanScheduleWeekdays1"),
	// 	Type: to.Ptr("Microsoft.DesktopVirtualization/scalingPlans/pooledSchedules"),
	// 	ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/scalingPlans/scalingPlan1/pooledSchedules/scalingPlanScheduleWeekdays1"),
	// 	Properties: &armdesktopvirtualization.ScalingPlanPooledScheduleProperties{
	// 		DaysOfWeek: []*armdesktopvirtualization.DayOfWeek{
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekMonday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekTuesday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekWednesday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekThursday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekFriday)},
	// 			OffPeakLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
	// 			OffPeakStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](20),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 			PeakLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmBreadthFirst),
	// 			PeakStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](8),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 			RampDownCapacityThresholdPct: to.Ptr[int32](50),
	// 			RampDownForceLogoffUsers: to.Ptr(true),
	// 			RampDownLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
	// 			RampDownMinimumHostsPct: to.Ptr[int32](20),
	// 			RampDownNotificationMessage: to.Ptr("message"),
	// 			RampDownStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](18),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 			RampDownWaitTimeMinutes: to.Ptr[int32](30),
	// 			RampUpCapacityThresholdPct: to.Ptr[int32](80),
	// 			RampUpLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
	// 			RampUpMinimumHostsPct: to.Ptr[int32](20),
	// 			RampUpStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](6),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 		},
	// 		SystemData: &armdesktopvirtualization.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f339d469c0fe83466edfe295a7960c82ebecf4f/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2022-09-09/examples/ScalingPlanPooledSchedule_Delete.json
func ExampleScalingPlanPooledSchedulesClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewScalingPlanPooledSchedulesClient().Delete(ctx, "resourceGroup1", "scalingPlan1", "scalingPlanScheduleWeekdays1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f339d469c0fe83466edfe295a7960c82ebecf4f/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2022-09-09/examples/ScalingPlanPooledSchedule_Update.json
func ExampleScalingPlanPooledSchedulesClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewScalingPlanPooledSchedulesClient().Update(ctx, "resourceGroup1", "scalingPlan1", "scalingPlanScheduleWeekdays1", &armdesktopvirtualization.ScalingPlanPooledSchedulesClientUpdateOptions{ScalingPlanSchedule: &armdesktopvirtualization.ScalingPlanPooledSchedulePatch{
		Properties: &armdesktopvirtualization.ScalingPlanPooledScheduleProperties{
			DaysOfWeek: []*armdesktopvirtualization.DayOfWeek{
				to.Ptr(armdesktopvirtualization.DayOfWeekMonday),
				to.Ptr(armdesktopvirtualization.DayOfWeekTuesday),
				to.Ptr(armdesktopvirtualization.DayOfWeekWednesday),
				to.Ptr(armdesktopvirtualization.DayOfWeekThursday),
				to.Ptr(armdesktopvirtualization.DayOfWeekFriday)},
			PeakStartTime: &armdesktopvirtualization.Time{
				Hour:   to.Ptr[int32](8),
				Minute: to.Ptr[int32](0),
			},
			RampDownLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
			RampDownMinimumHostsPct:        to.Ptr[int32](20),
			RampDownWaitTimeMinutes:        to.Ptr[int32](30),
			RampUpCapacityThresholdPct:     to.Ptr[int32](80),
			RampUpLoadBalancingAlgorithm:   to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
		},
	},
	})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ScalingPlanPooledSchedule = armdesktopvirtualization.ScalingPlanPooledSchedule{
	// 	Name: to.Ptr("scalingPlanScheduleWeekdays1"),
	// 	Type: to.Ptr("Microsoft.DesktopVirtualization/scalingPlans/pooledSchedules"),
	// 	ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/scalingPlans/scalingPlan1/pooledSchedules/scalingPlanScheduleWeekdays1"),
	// 	Properties: &armdesktopvirtualization.ScalingPlanPooledScheduleProperties{
	// 		DaysOfWeek: []*armdesktopvirtualization.DayOfWeek{
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekMonday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekTuesday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekWednesday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekThursday),
	// 			to.Ptr(armdesktopvirtualization.DayOfWeekFriday)},
	// 			OffPeakLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
	// 			OffPeakStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](20),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 			PeakLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmBreadthFirst),
	// 			PeakStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](8),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 			RampDownCapacityThresholdPct: to.Ptr[int32](50),
	// 			RampDownForceLogoffUsers: to.Ptr(true),
	// 			RampDownLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
	// 			RampDownMinimumHostsPct: to.Ptr[int32](20),
	// 			RampDownNotificationMessage: to.Ptr("message"),
	// 			RampDownStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](18),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 			RampDownWaitTimeMinutes: to.Ptr[int32](30),
	// 			RampUpCapacityThresholdPct: to.Ptr[int32](80),
	// 			RampUpLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
	// 			RampUpMinimumHostsPct: to.Ptr[int32](20),
	// 			RampUpStartTime: &armdesktopvirtualization.Time{
	// 				Hour: to.Ptr[int32](6),
	// 				Minute: to.Ptr[int32](0),
	// 			},
	// 		},
	// 		SystemData: &armdesktopvirtualization.SystemData{
	// 			CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
	// 			CreatedBy: to.Ptr("user1"),
	// 			CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 			LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
	// 			LastModifiedBy: to.Ptr("user2"),
	// 			LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/9f339d469c0fe83466edfe295a7960c82ebecf4f/specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2022-09-09/examples/ScalingPlanPooledSchedule_List.json
func ExampleScalingPlanPooledSchedulesClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armdesktopvirtualization.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewScalingPlanPooledSchedulesClient().NewListPager("resourceGroup1", "scalingPlan1", &armdesktopvirtualization.ScalingPlanPooledSchedulesClientListOptions{PageSize: to.Ptr[int32](10),
		IsDescending: to.Ptr(true),
		InitialSkip:  to.Ptr[int32](0),
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
		// page.ScalingPlanPooledScheduleList = armdesktopvirtualization.ScalingPlanPooledScheduleList{
		// 	Value: []*armdesktopvirtualization.ScalingPlanPooledSchedule{
		// 		{
		// 			Name: to.Ptr("scalingPlanScheduleWeekdays1"),
		// 			Type: to.Ptr("Microsoft.DesktopVirtualization/scalingPlans/pooledSchedules"),
		// 			ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/scalingPlans/scalingPlan1/pooledSchedules/scalingPlanScheduleWeekdays1"),
		// 			Properties: &armdesktopvirtualization.ScalingPlanPooledScheduleProperties{
		// 				DaysOfWeek: []*armdesktopvirtualization.DayOfWeek{
		// 					to.Ptr(armdesktopvirtualization.DayOfWeekMonday),
		// 					to.Ptr(armdesktopvirtualization.DayOfWeekTuesday),
		// 					to.Ptr(armdesktopvirtualization.DayOfWeekWednesday),
		// 					to.Ptr(armdesktopvirtualization.DayOfWeekThursday),
		// 					to.Ptr(armdesktopvirtualization.DayOfWeekFriday)},
		// 					OffPeakLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
		// 					OffPeakStartTime: &armdesktopvirtualization.Time{
		// 						Hour: to.Ptr[int32](20),
		// 						Minute: to.Ptr[int32](0),
		// 					},
		// 					PeakLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmBreadthFirst),
		// 					PeakStartTime: &armdesktopvirtualization.Time{
		// 						Hour: to.Ptr[int32](8),
		// 						Minute: to.Ptr[int32](0),
		// 					},
		// 					RampDownCapacityThresholdPct: to.Ptr[int32](50),
		// 					RampDownForceLogoffUsers: to.Ptr(true),
		// 					RampDownLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
		// 					RampDownMinimumHostsPct: to.Ptr[int32](20),
		// 					RampDownNotificationMessage: to.Ptr("message"),
		// 					RampDownStartTime: &armdesktopvirtualization.Time{
		// 						Hour: to.Ptr[int32](18),
		// 						Minute: to.Ptr[int32](0),
		// 					},
		// 					RampDownWaitTimeMinutes: to.Ptr[int32](30),
		// 					RampUpCapacityThresholdPct: to.Ptr[int32](80),
		// 					RampUpLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
		// 					RampUpMinimumHostsPct: to.Ptr[int32](20),
		// 					RampUpStartTime: &armdesktopvirtualization.Time{
		// 						Hour: to.Ptr[int32](6),
		// 						Minute: to.Ptr[int32](0),
		// 					},
		// 				},
		// 				SystemData: &armdesktopvirtualization.SystemData{
		// 					CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 					CreatedBy: to.Ptr("user1"),
		// 					CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 					LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 					LastModifiedBy: to.Ptr("user2"),
		// 					LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("scalingPlanScheduleWeekends1"),
		// 				Type: to.Ptr("Microsoft.DesktopVirtualization/scalingPlans/pooledSchedules"),
		// 				ID: to.Ptr("/subscriptions/daefabc0-95b4-48b3-b645-8a753a63c4fa/resourceGroups/resourceGroup1/providers/Microsoft.DesktopVirtualization/scalingPlans/scalingPlan1/pooledSchedules/scalingPlanScheduleWeekends1"),
		// 				Properties: &armdesktopvirtualization.ScalingPlanPooledScheduleProperties{
		// 					DaysOfWeek: []*armdesktopvirtualization.DayOfWeek{
		// 						to.Ptr(armdesktopvirtualization.DayOfWeekSaturday),
		// 						to.Ptr(armdesktopvirtualization.DayOfWeekSunday)},
		// 						OffPeakLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
		// 						OffPeakStartTime: &armdesktopvirtualization.Time{
		// 							Hour: to.Ptr[int32](20),
		// 							Minute: to.Ptr[int32](0),
		// 						},
		// 						PeakLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmBreadthFirst),
		// 						PeakStartTime: &armdesktopvirtualization.Time{
		// 							Hour: to.Ptr[int32](8),
		// 							Minute: to.Ptr[int32](0),
		// 						},
		// 						RampDownCapacityThresholdPct: to.Ptr[int32](100),
		// 						RampDownForceLogoffUsers: to.Ptr(true),
		// 						RampDownLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
		// 						RampDownMinimumHostsPct: to.Ptr[int32](0),
		// 						RampDownNotificationMessage: to.Ptr("message"),
		// 						RampDownStartTime: &armdesktopvirtualization.Time{
		// 							Hour: to.Ptr[int32](18),
		// 							Minute: to.Ptr[int32](0),
		// 						},
		// 						RampDownWaitTimeMinutes: to.Ptr[int32](30),
		// 						RampUpCapacityThresholdPct: to.Ptr[int32](90),
		// 						RampUpLoadBalancingAlgorithm: to.Ptr(armdesktopvirtualization.SessionHostLoadBalancingAlgorithmDepthFirst),
		// 						RampUpMinimumHostsPct: to.Ptr[int32](10),
		// 						RampUpStartTime: &armdesktopvirtualization.Time{
		// 							Hour: to.Ptr[int32](6),
		// 							Minute: to.Ptr[int32](0),
		// 						},
		// 					},
		// 					SystemData: &armdesktopvirtualization.SystemData{
		// 						CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.1234567Z"); return t}()),
		// 						CreatedBy: to.Ptr("user1"),
		// 						CreatedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 						LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.1234567Z"); return t}()),
		// 						LastModifiedBy: to.Ptr("user2"),
		// 						LastModifiedByType: to.Ptr(armdesktopvirtualization.CreatedByTypeUser),
		// 					},
		// 			}},
		// 		}
	}
}
