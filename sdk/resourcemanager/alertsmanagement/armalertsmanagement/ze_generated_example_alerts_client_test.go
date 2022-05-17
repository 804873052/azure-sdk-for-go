//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armalertsmanagement_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/AlertsMetaData_MonitorService.json
func ExampleAlertsClient_MetaData() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armalertsmanagement.NewAlertsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.MetaData(ctx,
		armalertsmanagement.IdentifierMonitorServiceList,
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/Alerts_List.json
func ExampleAlertsClient_NewGetAllPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armalertsmanagement.NewAlertsClient("1e3ff1c0-771a-4119-a03b-be82a51e232d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewGetAllPager(&armalertsmanagement.AlertsClientGetAllOptions{TargetResource: nil,
		TargetResourceType:  nil,
		TargetResourceGroup: nil,
		MonitorService:      nil,
		MonitorCondition:    nil,
		Severity:            nil,
		AlertState:          nil,
		AlertRule:           nil,
		SmartGroupID:        nil,
		IncludeContext:      nil,
		IncludeEgressConfig: nil,
		PageCount:           nil,
		SortBy:              nil,
		SortOrder:           nil,
		Select:              nil,
		TimeRange:           nil,
		CustomTimeRange:     nil,
	})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/Alerts_ChangeState.json
func ExampleAlertsClient_ChangeState() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armalertsmanagement.NewAlertsClient("9e261de7-c804-4b9d-9ebf-6f50fe350a9a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.ChangeState(ctx,
		"66114d64-d9d9-478b-95c9-b789d6502100",
		armalertsmanagement.AlertStateAcknowledged,
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/Alerts_History.json
func ExampleAlertsClient_GetHistory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armalertsmanagement.NewAlertsClient("9e261de7-c804-4b9d-9ebf-6f50fe350a9a", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetHistory(ctx,
		"66114d64-d9d9-478b-95c9-b789d6502100",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/alertsmanagement/resource-manager/Microsoft.AlertsManagement/preview/2019-05-05-preview/examples/Alerts_Summary.json
func ExampleAlertsClient_GetSummary() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armalertsmanagement.NewAlertsClient("1e3ff1c0-771a-4119-a03b-be82a51e232d", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.GetSummary(ctx,
		armalertsmanagement.AlertsSummaryGroupByFields("severity,alertState"),
		&armalertsmanagement.AlertsClientGetSummaryOptions{IncludeSmartGroupsCount: nil,
			TargetResource:      nil,
			TargetResourceType:  nil,
			TargetResourceGroup: nil,
			MonitorService:      nil,
			MonitorCondition:    nil,
			Severity:            nil,
			AlertState:          nil,
			AlertRule:           nil,
			TimeRange:           nil,
			CustomTimeRange:     nil,
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
