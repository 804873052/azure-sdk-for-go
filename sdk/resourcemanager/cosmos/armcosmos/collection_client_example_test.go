//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcosmos_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b065afacb99e7ec65787383550ee233e0e02a6e6/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-03-15/examples/CosmosDBCollectionGetMetrics.json
func ExampleCollectionClient_NewListMetricsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCollectionClient().NewListMetricsPager("rg1", "ddb1", "databaseRid", "collectionRid", "$filter=(name.value eq 'Total Requests') and timeGrain eq duration'PT5M' and startTime eq '2017-11-19T23:53:55.2780000Z' and endTime eq '2017-11-20T00:13:55.2780000Z", nil)
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
		// page.MetricListResult = armcosmos.MetricListResult{
		// 	Value: []*armcosmos.Metric{
		// 		{
		// 			Name: &armcosmos.MetricName{
		// 				LocalizedValue: to.Ptr("Total Requests"),
		// 				Value: to.Ptr("Total Requests"),
		// 			},
		// 			EndTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-20T00:13:55.2780000Z"); return t}()),
		// 			MetricValues: []*armcosmos.MetricValue{
		// 				{
		// 					Count: to.Ptr[int32](0),
		// 					Average: to.Ptr[float64](0),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-19T23:53:55.2780000Z"); return t}()),
		// 					Total: to.Ptr[float64](0),
		// 				},
		// 				{
		// 					Count: to.Ptr[int32](0),
		// 					Average: to.Ptr[float64](0),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-19T23:58:55.2780000Z"); return t}()),
		// 					Total: to.Ptr[float64](0),
		// 				},
		// 				{
		// 					Count: to.Ptr[int32](0),
		// 					Average: to.Ptr[float64](0),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-20T00:03:55.2780000Z"); return t}()),
		// 					Total: to.Ptr[float64](0),
		// 				},
		// 				{
		// 					Count: to.Ptr[int32](0),
		// 					Average: to.Ptr[float64](0),
		// 					Timestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-20T00:08:55.2780000Z"); return t}()),
		// 					Total: to.Ptr[float64](0),
		// 			}},
		// 			StartTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-11-19T23:53:55.2780000Z"); return t}()),
		// 			TimeGrain: to.Ptr("PT5M"),
		// 			Unit: to.Ptr(armcosmos.UnitTypeCount),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b065afacb99e7ec65787383550ee233e0e02a6e6/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-03-15/examples/CosmosDBCollectionGetUsages.json
func ExampleCollectionClient_NewListUsagesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCollectionClient().NewListUsagesPager("rg1", "ddb1", "databaseRid", "collectionRid", &armcosmos.CollectionClientListUsagesOptions{Filter: to.Ptr("$filter=name.value eq 'Storage'")})
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
		// page.UsagesResult = armcosmos.UsagesResult{
		// 	Value: []*armcosmos.Usage{
		// 		{
		// 			Name: &armcosmos.MetricName{
		// 				LocalizedValue: to.Ptr("Storage"),
		// 				Value: to.Ptr("Storage"),
		// 			},
		// 			CurrentValue: to.Ptr[int64](0),
		// 			Limit: to.Ptr[int64](10737418240),
		// 			QuotaPeriod: to.Ptr("P1D"),
		// 			Unit: to.Ptr(armcosmos.UnitTypeBytes),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/b065afacb99e7ec65787383550ee233e0e02a6e6/specification/cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2023-03-15/examples/CosmosDBCollectionGetMetricDefinitions.json
func ExampleCollectionClient_NewListMetricDefinitionsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armcosmos.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCollectionClient().NewListMetricDefinitionsPager("rg1", "ddb1", "databaseRid", "collectionRid", nil)
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
		// page.MetricDefinitionsListResult = armcosmos.MetricDefinitionsListResult{
		// 	Value: []*armcosmos.MetricDefinition{
		// 		{
		// 			Name: &armcosmos.MetricName{
		// 				LocalizedValue: to.Ptr("Total Requests"),
		// 				Value: to.Ptr("Total Requests"),
		// 			},
		// 			MetricAvailabilities: []*armcosmos.MetricAvailability{
		// 				{
		// 					Retention: to.Ptr("P2D"),
		// 					TimeGrain: to.Ptr("PT5M"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P14D"),
		// 					TimeGrain: to.Ptr("PT1H"),
		// 				},
		// 				{
		// 					Retention: to.Ptr("P60D"),
		// 					TimeGrain: to.Ptr("P1D"),
		// 			}},
		// 			PrimaryAggregationType: to.Ptr(armcosmos.PrimaryAggregationTypeTotal),
		// 			ResourceURI: to.Ptr("/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1"),
		// 			Unit: to.Ptr(armcosmos.UnitTypeCount),
		// 	}},
		// }
	}
}
