//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a436672b07fb1fe276c203b086b3f0e0d0c4aa24/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/SmartGroups_List.json
func ExampleSmartGroupsClient_NewGetAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSmartGroupsClient().NewGetAllPager(&armalertsmanagement.SmartGroupsClientGetAllOptions{TargetResource: nil,
		TargetResourceGroup: nil,
		TargetResourceType:  nil,
		MonitorService:      nil,
		MonitorCondition:    nil,
		Severity:            nil,
		SmartGroupState:     nil,
		TimeRange:           nil,
		PageCount:           nil,
		SortBy:              nil,
		SortOrder:           nil,
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
		// page.SmartGroupsList = armalertsmanagement.SmartGroupsList{
		// 	Value: []*armalertsmanagement.SmartGroup{
		// 		{
		// 			Name: to.Ptr("cpu alert"),
		// 			Type: to.Ptr("Microsoft.AlertsManagement/smartGroups"),
		// 			ID: to.Ptr("/subscriptions/dd91de05-d791-4ceb-b6dc-988682dc7d72/providers/Microsoft.AlertsManagement/smartGroups/a808445e-bb38-4751-85c2-1b109ccc1059"),
		// 			Properties: &armalertsmanagement.SmartGroupProperties{
		// 				AlertSeverities: []*armalertsmanagement.SmartGroupAggregatedProperty{
		// 					{
		// 						Name: to.Ptr("Sev3"),
		// 						Count: to.Ptr[int64](1942),
		// 				}},
		// 				AlertStates: []*armalertsmanagement.SmartGroupAggregatedProperty{
		// 					{
		// 						Name: to.Ptr("New"),
		// 						Count: to.Ptr[int64](1941),
		// 					},
		// 					{
		// 						Name: to.Ptr("Acknowledged"),
		// 						Count: to.Ptr[int64](1),
		// 				}},
		// 				AlertsCount: to.Ptr[int64](1942),
		// 				LastModifiedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-13T06:30:09Z"); return t}()),
		// 				LastModifiedUserName: to.Ptr("System"),
		// 				MonitorConditions: []*armalertsmanagement.SmartGroupAggregatedProperty{
		// 					{
		// 						Name: to.Ptr("Fired"),
		// 						Count: to.Ptr[int64](1942),
		// 				}},
		// 				MonitorServices: []*armalertsmanagement.SmartGroupAggregatedProperty{
		// 					{
		// 						Name: to.Ptr("Application Insights"),
		// 						Count: to.Ptr[int64](1942),
		// 				}},
		// 				ResourceGroups: []*armalertsmanagement.SmartGroupAggregatedProperty{
		// 					{
		// 						Name: to.Ptr("alertscorrelationrg"),
		// 						Count: to.Ptr[int64](1942),
		// 				}},
		// 				ResourceTypes: []*armalertsmanagement.SmartGroupAggregatedProperty{
		// 					{
		// 						Name: to.Ptr("components"),
		// 						Count: to.Ptr[int64](1942),
		// 				}},
		// 				Resources: []*armalertsmanagement.SmartGroupAggregatedProperty{
		// 					{
		// 						Name: to.Ptr("/subscriptions/dd91de05-d791-4ceb-b6dc-988682dc7d72/resourcegroups/alertscorrelationrg/providers/microsoft.insights/components/alertscorrelationworkerrole_int"),
		// 						Count: to.Ptr[int64](1942),
		// 				}},
		// 				Severity: to.Ptr(armalertsmanagement.SeveritySev3),
		// 				SmartGroupState: to.Ptr(armalertsmanagement.StateNew),
		// 				StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-06T12:35:09Z"); return t}()),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("CPU Alert"),
		// 			Type: to.Ptr("Microsoft.AlertsManagement/smartGroups"),
		// 			ID: to.Ptr("/subscriptions/dd91de05-d791-4ceb-b6dc-988682dc7d72/providers/Microsoft.AlertsManagement/smartGroups/01114c7c-769f-4fd4-b6fa-ab77693b83cd"),
		// 			Properties: &armalertsmanagement.SmartGroupProperties{
		// 				AlertSeverities: []*armalertsmanagement.SmartGroupAggregatedProperty{
		// 					{
		// 						Name: to.Ptr("Sev0"),
		// 						Count: to.Ptr[int64](6984),
		// 					},
		// 					{
		// 						Name: to.Ptr("Sev1"),
		// 						Count: to.Ptr[int64](6927),
		// 				}},
		// 				AlertStates: []*armalertsmanagement.SmartGroupAggregatedProperty{
		// 					{
		// 						Name: to.Ptr("New"),
		// 						Count: to.Ptr[int64](15358),
		// 					},
		// 					{
		// 						Name: to.Ptr("Acknowledged"),
		// 						Count: to.Ptr[int64](12),
		// 					},
		// 					{
		// 						Name: to.Ptr("Closed"),
		// 						Count: to.Ptr[int64](4),
		// 				}},
		// 				AlertsCount: to.Ptr[int64](15374),
		// 				LastModifiedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-13T06:29:01Z"); return t}()),
		// 				LastModifiedUserName: to.Ptr("System"),
		// 				MonitorConditions: []*armalertsmanagement.SmartGroupAggregatedProperty{
		// 					{
		// 						Name: to.Ptr("Fired"),
		// 						Count: to.Ptr[int64](15374),
		// 				}},
		// 				MonitorServices: []*armalertsmanagement.SmartGroupAggregatedProperty{
		// 					{
		// 						Name: to.Ptr("Log Analytics"),
		// 						Count: to.Ptr[int64](13911),
		// 				}},
		// 				ResourceGroups: []*armalertsmanagement.SmartGroupAggregatedProperty{
		// 					{
		// 						Name: to.Ptr("mms-eus"),
		// 						Count: to.Ptr[int64](15374),
		// 				}},
		// 				ResourceTypes: []*armalertsmanagement.SmartGroupAggregatedProperty{
		// 					{
		// 						Name: to.Ptr("microsoft.operationalinsights/workspaces"),
		// 						Count: to.Ptr[int64](6912),
		// 					},
		// 					{
		// 						Name: to.Ptr("workspaces"),
		// 						Count: to.Ptr[int64](8462),
		// 				}},
		// 				Resources: []*armalertsmanagement.SmartGroupAggregatedProperty{
		// 					{
		// 						Name: to.Ptr("/subscriptions/dd91de05-d791-4ceb-b6dc-988682dc7d72/resourcegroups/mms-eus/providers/microsoft.operationalinsights/workspaces/alertsint"),
		// 						Count: to.Ptr[int64](15374),
		// 				}},
		// 				Severity: to.Ptr(armalertsmanagement.SeveritySev0),
		// 				SmartGroupState: to.Ptr(armalertsmanagement.StateAcknowledged),
		// 				StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-05-17T10:18:44.2020997Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a436672b07fb1fe276c203b086b3f0e0d0c4aa24/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/SmartGroups_GetById.json
func ExampleSmartGroupsClient_GetByID() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSmartGroupsClient().GetByID(ctx, "603675da-9851-4b26-854a-49fc53d32715", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SmartGroup = armalertsmanagement.SmartGroup{
	// 	Name: to.Ptr("cpu alert"),
	// 	Type: to.Ptr("Microsoft.AlertsManagement/smartGroups"),
	// 	ID: to.Ptr("/subscriptions/dd91de05-d791-4ceb-b6dc-988682dc7d72/providers/Microsoft.AlertsManagement/smartGroups/a808445e-bb38-4751-85c2-1b109ccc1059"),
	// 	Properties: &armalertsmanagement.SmartGroupProperties{
	// 		AlertSeverities: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("Sev3"),
	// 				Count: to.Ptr[int64](1942),
	// 		}},
	// 		AlertStates: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("New"),
	// 				Count: to.Ptr[int64](1941),
	// 			},
	// 			{
	// 				Name: to.Ptr("Acknowledged"),
	// 				Count: to.Ptr[int64](1),
	// 		}},
	// 		AlertsCount: to.Ptr[int64](1942),
	// 		LastModifiedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-13T06:30:09Z"); return t}()),
	// 		LastModifiedUserName: to.Ptr("System"),
	// 		MonitorConditions: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("Fired"),
	// 				Count: to.Ptr[int64](1942),
	// 		}},
	// 		MonitorServices: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("Application Insights"),
	// 				Count: to.Ptr[int64](1942),
	// 		}},
	// 		ResourceGroups: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("alertscorrelationrg"),
	// 				Count: to.Ptr[int64](1942),
	// 		}},
	// 		ResourceTypes: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("components"),
	// 				Count: to.Ptr[int64](1942),
	// 		}},
	// 		Resources: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("/subscriptions/dd91de05-d791-4ceb-b6dc-988682dc7d72/resourcegroups/alertscorrelationrg/providers/microsoft.insights/components/alertscorrelationworkerrole_int"),
	// 				Count: to.Ptr[int64](1942),
	// 		}},
	// 		Severity: to.Ptr(armalertsmanagement.SeveritySev3),
	// 		SmartGroupState: to.Ptr(armalertsmanagement.StateNew),
	// 		StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-06T12:35:09Z"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a436672b07fb1fe276c203b086b3f0e0d0c4aa24/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/SmartGroups_ChangeState.json
func ExampleSmartGroupsClient_ChangeState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSmartGroupsClient().ChangeState(ctx, "a808445e-bb38-4751-85c2-1b109ccc1059", armalertsmanagement.AlertStateAcknowledged, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SmartGroup = armalertsmanagement.SmartGroup{
	// 	Name: to.Ptr("cpu alert"),
	// 	Type: to.Ptr("Microsoft.AlertsManagement/smartGroups"),
	// 	ID: to.Ptr("/subscriptions/dd91de05-d791-4ceb-b6dc-988682dc7d72/providers/Microsoft.AlertsManagement/smartGroups/a808445e-bb38-4751-85c2-1b109ccc1059"),
	// 	Properties: &armalertsmanagement.SmartGroupProperties{
	// 		AlertSeverities: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("Sev3"),
	// 				Count: to.Ptr[int64](1942),
	// 		}},
	// 		AlertStates: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("New"),
	// 				Count: to.Ptr[int64](1941),
	// 			},
	// 			{
	// 				Name: to.Ptr("Acknowledged"),
	// 				Count: to.Ptr[int64](1),
	// 		}},
	// 		AlertsCount: to.Ptr[int64](1942),
	// 		LastModifiedDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-13T06:30:09Z"); return t}()),
	// 		LastModifiedUserName: to.Ptr("System"),
	// 		MonitorConditions: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("Fired"),
	// 				Count: to.Ptr[int64](1942),
	// 		}},
	// 		MonitorServices: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("Application Insights"),
	// 				Count: to.Ptr[int64](1942),
	// 		}},
	// 		ResourceGroups: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("alertscorrelationrg"),
	// 				Count: to.Ptr[int64](1942),
	// 		}},
	// 		ResourceTypes: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("components"),
	// 				Count: to.Ptr[int64](1942),
	// 		}},
	// 		Resources: []*armalertsmanagement.SmartGroupAggregatedProperty{
	// 			{
	// 				Name: to.Ptr("/subscriptions/dd91de05-d791-4ceb-b6dc-988682dc7d72/resourcegroups/alertscorrelationrg/providers/microsoft.insights/components/alertscorrelationworkerrole_int"),
	// 				Count: to.Ptr[int64](1942),
	// 		}},
	// 		Severity: to.Ptr(armalertsmanagement.SeveritySev3),
	// 		SmartGroupState: to.Ptr(armalertsmanagement.StateAcknowledged),
	// 		StartDateTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-06T12:35:09Z"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/a436672b07fb1fe276c203b086b3f0e0d0c4aa24/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/SmartGroups_History.json
func ExampleSmartGroupsClient_GetHistory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armalertsmanagement.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSmartGroupsClient().GetHistory(ctx, "a808445e-bb38-4751-85c2-1b109ccc1059", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SmartGroupModification = armalertsmanagement.SmartGroupModification{
	// 	Name: to.Ptr("cpu alert"),
	// 	Type: to.Ptr("Microsoft.AlertsManagement/alerts"),
	// 	ID: to.Ptr("/subscriptions/9e261de7-c804-4b9d-9ebf-6f50fe350a9a/providers/Microsoft.AlertsManagement/smartGroups/a808445e-bb38-4751-85c2-1b109ccc1059/history/default"),
	// 	Properties: &armalertsmanagement.SmartGroupModificationProperties{
	// 		Modifications: []*armalertsmanagement.SmartGroupModificationItem{
	// 			{
	// 				Description: to.Ptr("New Smart Group is created"),
	// 				Comments: to.Ptr(""),
	// 				ModificationEvent: to.Ptr(armalertsmanagement.SmartGroupModificationEventSmartGroupCreated),
	// 				ModifiedAt: to.Ptr("2018-06-06T12:35:09Z"),
	// 				ModifiedBy: to.Ptr("System"),
	// 				NewValue: to.Ptr(""),
	// 				OldValue: to.Ptr(""),
	// 		}},
	// 		SmartGroupID: to.Ptr("a808445e-bb38-4751-85c2-1b109ccc1059"),
	// 	},
	// }
}
