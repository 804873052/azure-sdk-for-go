//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetLogAnalyticsMetrics.json
func ExampleLogAnalyticsClient_GetLogAnalyticsMetrics() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcdn.NewLogAnalyticsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetLogAnalyticsMetrics(ctx,
		"<resource-group-name>",
		"<profile-name>",
		[]armcdn.LogMetric{
			armcdn.LogMetricClientRequestCount},
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T04:30:00.000Z"); return t }(),
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T05:00:00.000Z"); return t }(),
		armcdn.LogMetricsGranularityPT5M,
		[]string{
			"customdomain1.azurecdn.net",
			"customdomain2.azurecdn.net"},
		[]string{
			"https"},
		&armcdn.LogAnalyticsClientGetLogAnalyticsMetricsOptions{GroupBy: []armcdn.LogMetricsGroupBy{
			armcdn.LogMetricsGroupByProtocol},
			Continents:       []string{},
			CountryOrRegions: []string{},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetLogAnalyticsRankings.json
func ExampleLogAnalyticsClient_GetLogAnalyticsRankings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcdn.NewLogAnalyticsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetLogAnalyticsRankings(ctx,
		"<resource-group-name>",
		"<profile-name>",
		[]armcdn.LogRanking{
			armcdn.LogRankingURL},
		[]armcdn.LogRankingMetric{
			armcdn.LogRankingMetricClientRequestCount},
		5,
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T06:49:27.554Z"); return t }(),
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T09:49:27.554Z"); return t }(),
		&armcdn.LogAnalyticsClientGetLogAnalyticsRankingsOptions{CustomDomains: []string{}})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetLogAnalyticsLocations.json
func ExampleLogAnalyticsClient_GetLogAnalyticsLocations() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcdn.NewLogAnalyticsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetLogAnalyticsLocations(ctx,
		"<resource-group-name>",
		"<profile-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetLogAnalyticsResources.json
func ExampleLogAnalyticsClient_GetLogAnalyticsResources() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcdn.NewLogAnalyticsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetLogAnalyticsResources(ctx,
		"<resource-group-name>",
		"<profile-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetWafLogAnalyticsMetrics.json
func ExampleLogAnalyticsClient_GetWafLogAnalyticsMetrics() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcdn.NewLogAnalyticsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetWafLogAnalyticsMetrics(ctx,
		"<resource-group-name>",
		"<profile-name>",
		[]armcdn.WafMetric{
			armcdn.WafMetricClientRequestCount},
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T06:49:27.554Z"); return t }(),
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T09:49:27.554Z"); return t }(),
		armcdn.WafGranularityPT5M,
		&armcdn.LogAnalyticsClientGetWafLogAnalyticsMetricsOptions{Actions: []armcdn.WafAction{
			armcdn.WafActionBlock,
			armcdn.WafActionLog},
			GroupBy:   []armcdn.WafRankingGroupBy{},
			RuleTypes: []armcdn.WafRuleType{},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/LogAnalytics_GetWafLogAnalyticsRankings.json
func ExampleLogAnalyticsClient_GetWafLogAnalyticsRankings() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armcdn.NewLogAnalyticsClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.GetWafLogAnalyticsRankings(ctx,
		"<resource-group-name>",
		"<profile-name>",
		[]armcdn.WafMetric{
			armcdn.WafMetricClientRequestCount},
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T06:49:27.554Z"); return t }(),
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-11-04T09:49:27.554Z"); return t }(),
		5,
		[]armcdn.WafRankingType{
			armcdn.WafRankingTypeRuleID},
		&armcdn.LogAnalyticsClientGetWafLogAnalyticsRankingsOptions{Actions: []armcdn.WafAction{},
			RuleTypes: []armcdn.WafRuleType{},
		})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}
