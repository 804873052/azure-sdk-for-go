//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armservicebus_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus/v2"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-01-01-preview/examples/Subscriptions/SBSubscriptionListByTopic.json
func ExampleSubscriptionsClient_NewListByTopicPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSubscriptionsClient().NewListByTopicPager("ResourceGroup", "sdk-Namespace-1349", "sdk-Topics-8740", &armservicebus.SubscriptionsClientListByTopicOptions{Skip: nil,
		Top: nil,
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
		// page.SBSubscriptionListResult = armservicebus.SBSubscriptionListResult{
		// 	Value: []*armservicebus.SBSubscription{
		// 		{
		// 			Name: to.Ptr("sdk-Subscriptions-2178"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces/Topics/Subscriptions"),
		// 			ID: to.Ptr("/subscriptions/Subscriptionid/resourceGroups/ResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-1349/topics/sdk-Topics-8740/subscriptions/sdk-Subscriptions-2178"),
		// 			Properties: &armservicebus.SBSubscriptionProperties{
		// 				AccessedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-04T18:02:20.5992764Z"); return t}()),
		// 				AutoDeleteOnIdle: to.Ptr("P10675199DT2H48M5.4775807S"),
		// 				CountDetails: &armservicebus.MessageCountDetails{
		// 					ActiveMessageCount: to.Ptr[int64](0),
		// 					DeadLetterMessageCount: to.Ptr[int64](0),
		// 					ScheduledMessageCount: to.Ptr[int64](0),
		// 					TransferDeadLetterMessageCount: to.Ptr[int64](0),
		// 					TransferMessageCount: to.Ptr[int64](0),
		// 				},
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-04T18:02:20.5992764Z"); return t}()),
		// 				DeadLetteringOnFilterEvaluationExceptions: to.Ptr(true),
		// 				DeadLetteringOnMessageExpiration: to.Ptr(true),
		// 				DefaultMessageTimeToLive: to.Ptr("P10675199DT2H48M5.4775807S"),
		// 				EnableBatchedOperations: to.Ptr(true),
		// 				ForwardDeadLetteredMessagesTo: to.Ptr("sdk-Topics-3065"),
		// 				ForwardTo: to.Ptr("sdk-Topics-3065"),
		// 				LockDuration: to.Ptr("PT1M"),
		// 				MaxDeliveryCount: to.Ptr[int32](10),
		// 				MessageCount: to.Ptr[int64](0),
		// 				RequiresSession: to.Ptr(false),
		// 				Status: to.Ptr(armservicebus.EntityStatusActive),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-04T18:02:20.5992764Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-01-01-preview/examples/Subscriptions/SBSubscriptionCreate.json
func ExampleSubscriptionsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSubscriptionsClient().CreateOrUpdate(ctx, "ResourceGroup", "sdk-Namespace-1349", "sdk-Topics-8740", "sdk-Subscriptions-2178", armservicebus.SBSubscription{
		Properties: &armservicebus.SBSubscriptionProperties{
			EnableBatchedOperations: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SBSubscription = armservicebus.SBSubscription{
	// 	Name: to.Ptr("sdk-Subscriptions-2178"),
	// 	Type: to.Ptr("Microsoft.ServiceBus/Namespaces/Topics/Subscriptions"),
	// 	ID: to.Ptr("/subscriptions/Subscriptionid/resourceGroups/ResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-1349/topics/sdk-Topics-8740/subscriptions/sdk-Subscriptions-2178"),
	// 	Properties: &armservicebus.SBSubscriptionProperties{
	// 		AccessedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-04T18:02:20.5992764Z"); return t}()),
	// 		AutoDeleteOnIdle: to.Ptr("P10675199DT2H48M5.4775807S"),
	// 		CountDetails: &armservicebus.MessageCountDetails{
	// 			ActiveMessageCount: to.Ptr[int64](0),
	// 			DeadLetterMessageCount: to.Ptr[int64](0),
	// 			ScheduledMessageCount: to.Ptr[int64](0),
	// 			TransferDeadLetterMessageCount: to.Ptr[int64](0),
	// 			TransferMessageCount: to.Ptr[int64](0),
	// 		},
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-04T18:02:20.5992764Z"); return t}()),
	// 		DeadLetteringOnFilterEvaluationExceptions: to.Ptr(true),
	// 		DeadLetteringOnMessageExpiration: to.Ptr(true),
	// 		DefaultMessageTimeToLive: to.Ptr("P10675199DT2H48M5.4775807S"),
	// 		EnableBatchedOperations: to.Ptr(true),
	// 		ForwardDeadLetteredMessagesTo: to.Ptr("sdk-Topics-3065"),
	// 		ForwardTo: to.Ptr("sdk-Topics-3065"),
	// 		LockDuration: to.Ptr("PT1M"),
	// 		MaxDeliveryCount: to.Ptr[int32](10),
	// 		MessageCount: to.Ptr[int64](0),
	// 		RequiresSession: to.Ptr(false),
	// 		Status: to.Ptr(armservicebus.EntityStatusActive),
	// 		UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-04T18:02:20.5992764Z"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-01-01-preview/examples/Subscriptions/SBSubscriptionDelete.json
func ExampleSubscriptionsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewSubscriptionsClient().Delete(ctx, "ResourceGroup", "sdk-Namespace-5882", "sdk-Topics-1804", "sdk-Subscriptions-3670", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-01-01-preview/examples/Subscriptions/SBSubscriptionGet.json
func ExampleSubscriptionsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSubscriptionsClient().Get(ctx, "ResourceGroup", "sdk-Namespace-1349", "sdk-Topics-8740", "sdk-Subscriptions-2178", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SBSubscription = armservicebus.SBSubscription{
	// 	Name: to.Ptr("sdk-Subscriptions-2178"),
	// 	Type: to.Ptr("Microsoft.ServiceBus/Namespaces/Topics/Subscriptions"),
	// 	ID: to.Ptr("/subscriptions/Subscriptionid/resourceGroups/ResourceGroup/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-1349/topics/sdk-Topics-8740/subscriptions/sdk-Subscriptions-2178"),
	// 	Properties: &armservicebus.SBSubscriptionProperties{
	// 		AccessedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-04T18:02:20.5992764Z"); return t}()),
	// 		AutoDeleteOnIdle: to.Ptr("P10675199DT2H48M5.4775807S"),
	// 		CountDetails: &armservicebus.MessageCountDetails{
	// 			ActiveMessageCount: to.Ptr[int64](0),
	// 			DeadLetterMessageCount: to.Ptr[int64](0),
	// 			ScheduledMessageCount: to.Ptr[int64](0),
	// 			TransferDeadLetterMessageCount: to.Ptr[int64](0),
	// 			TransferMessageCount: to.Ptr[int64](0),
	// 		},
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-04T18:02:20.5992764Z"); return t}()),
	// 		DeadLetteringOnFilterEvaluationExceptions: to.Ptr(true),
	// 		DeadLetteringOnMessageExpiration: to.Ptr(true),
	// 		DefaultMessageTimeToLive: to.Ptr("P10675199DT2H48M5.4775807S"),
	// 		EnableBatchedOperations: to.Ptr(true),
	// 		ForwardDeadLetteredMessagesTo: to.Ptr("sdk-Topics-3065"),
	// 		ForwardTo: to.Ptr("sdk-Topics-3065"),
	// 		LockDuration: to.Ptr("PT1M"),
	// 		MaxDeliveryCount: to.Ptr[int32](10),
	// 		MessageCount: to.Ptr[int64](0),
	// 		RequiresSession: to.Ptr(false),
	// 		Status: to.Ptr(armservicebus.EntityStatusActive),
	// 		UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-01-04T18:02:20.5992764Z"); return t}()),
	// 	},
	// }
}
