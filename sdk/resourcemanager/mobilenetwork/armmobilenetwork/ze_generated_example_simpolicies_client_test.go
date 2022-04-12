//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmobilenetwork_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimPolicyDelete.json
func ExampleSimPoliciesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<mobile-network-name>",
		"<sim-policy-name>",
		&armmobilenetwork.SimPoliciesClientBeginDeleteOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimPolicyGet.json
func ExampleSimPoliciesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<mobile-network-name>",
		"<sim-policy-name>",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimPolicyCreate.json
func ExampleSimPoliciesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<mobile-network-name>",
		"<sim-policy-name>",
		armmobilenetwork.SimPolicy{
			Location: to.Ptr("<location>"),
			Properties: &armmobilenetwork.SimPolicyPropertiesFormat{
				DefaultSlice: &armmobilenetwork.SliceResourceID{
					ID: to.Ptr("<id>"),
				},
				RegistrationTimer: to.Ptr[int32](3240),
				SliceConfigurations: []*armmobilenetwork.SliceConfiguration{
					{
						DataNetworkConfigurations: []*armmobilenetwork.DataNetworkConfiguration{
							{
								FiveQi:                              to.Ptr[int32](9),
								AdditionalAllowedSessionTypes:       []*armmobilenetwork.PduSessionType{},
								AllocationAndRetentionPriorityLevel: to.Ptr[int32](9),
								AllowedServices: []*armmobilenetwork.ServiceResourceID{
									{
										ID: to.Ptr("<id>"),
									}},
								DataNetwork: &armmobilenetwork.DataNetworkResourceID{
									ID: to.Ptr("<id>"),
								},
								DefaultSessionType:      to.Ptr(armmobilenetwork.PduSessionTypeIPv4),
								PreemptionCapability:    to.Ptr(armmobilenetwork.PreemptionCapabilityNotPreempt),
								PreemptionVulnerability: to.Ptr(armmobilenetwork.PreemptionVulnerabilityPreemptable),
								SessionAmbr: &armmobilenetwork.Ambr{
									Downlink: to.Ptr("<downlink>"),
									Uplink:   to.Ptr("<uplink>"),
								},
							}},
						DefaultDataNetwork: &armmobilenetwork.DataNetworkResourceID{
							ID: to.Ptr("<id>"),
						},
						Slice: &armmobilenetwork.SliceResourceID{
							ID: to.Ptr("<id>"),
						},
					}},
				UeAmbr: &armmobilenetwork.Ambr{
					Downlink: to.Ptr("<downlink>"),
					Uplink:   to.Ptr("<uplink>"),
				},
			},
		},
		&armmobilenetwork.SimPoliciesClientBeginCreateOrUpdateOptions{ResumeToken: ""})
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimPolicyUpdateTags.json
func ExampleSimPoliciesClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	res, err := client.UpdateTags(ctx,
		"<resource-group-name>",
		"<mobile-network-name>",
		"<sim-policy-name>",
		armmobilenetwork.TagsObject{
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
		return
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-03-01-preview/examples/SimPolicyListByMobileNetwork.json
func ExampleSimPoliciesClient_ListByMobileNetwork() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
		return
	}
	ctx := context.Background()
	client, err := armmobilenetwork.NewSimPoliciesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
		return
	}
	pager := client.ListByMobileNetwork("<resource-group-name>",
		"<mobile-network-name>",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
