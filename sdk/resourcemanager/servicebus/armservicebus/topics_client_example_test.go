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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-01-01-preview/examples/Topics/SBTopicAuthorizationRuleListAll.json
func ExampleTopicsClient_NewListAuthorizationRulesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTopicsClient().NewListAuthorizationRulesPager("ArunMonocle", "sdk-Namespace-6261", "sdk-Topics-1984", nil)
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
		// page.SBAuthorizationRuleListResult = armservicebus.SBAuthorizationRuleListResult{
		// 	Value: []*armservicebus.SBAuthorizationRule{
		// 		{
		// 			Name: to.Ptr("sdk-AuthRules-4310"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces/Topics/AuthorizationRules"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-6261/topics/sdk-Topics-1984/authorizationRules/sdk-AuthRules-4310"),
		// 			Properties: &armservicebus.SBAuthorizationRuleProperties{
		// 				Rights: []*armservicebus.AccessRights{
		// 					to.Ptr(armservicebus.AccessRightsListen),
		// 					to.Ptr(armservicebus.AccessRightsSend)},
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-01-01-preview/examples/Topics/SBTopicAuthorizationRuleCreate.json
func ExampleTopicsClient_CreateOrUpdateAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTopicsClient().CreateOrUpdateAuthorizationRule(ctx, "ArunMonocle", "sdk-Namespace-6261", "sdk-Topics-1984", "sdk-AuthRules-4310", armservicebus.SBAuthorizationRule{
		Properties: &armservicebus.SBAuthorizationRuleProperties{
			Rights: []*armservicebus.AccessRights{
				to.Ptr(armservicebus.AccessRightsListen),
				to.Ptr(armservicebus.AccessRightsSend)},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SBAuthorizationRule = armservicebus.SBAuthorizationRule{
	// 	Name: to.Ptr("sdk-AuthRules-4310"),
	// 	Type: to.Ptr("Microsoft.ServiceBus/Namespaces/Topics/AuthorizationRules"),
	// 	ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-6261/topics/sdk-Topics-1984/authorizationRules/sdk-AuthRules-4310"),
	// 	Properties: &armservicebus.SBAuthorizationRuleProperties{
	// 		Rights: []*armservicebus.AccessRights{
	// 			to.Ptr(armservicebus.AccessRightsListen),
	// 			to.Ptr(armservicebus.AccessRightsSend)},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-01-01-preview/examples/Topics/SBTopicAuthorizationRuleGet.json
func ExampleTopicsClient_GetAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTopicsClient().GetAuthorizationRule(ctx, "ArunMonocle", "sdk-Namespace-6261", "sdk-Topics-1984", "sdk-AuthRules-4310", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SBAuthorizationRule = armservicebus.SBAuthorizationRule{
	// 	Name: to.Ptr("sdk-AuthRules-4310"),
	// 	Type: to.Ptr("Microsoft.ServiceBus/Namespaces/Topics/AuthorizationRules"),
	// 	ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-6261/topics/sdk-Topics-1984/authorizationRules/sdk-AuthRules-4310"),
	// 	Properties: &armservicebus.SBAuthorizationRuleProperties{
	// 		Rights: []*armservicebus.AccessRights{
	// 			to.Ptr(armservicebus.AccessRightsListen),
	// 			to.Ptr(armservicebus.AccessRightsSend)},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-01-01-preview/examples/Topics/SBTopicAuthorizationRuleDelete.json
func ExampleTopicsClient_DeleteAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTopicsClient().DeleteAuthorizationRule(ctx, "ArunMonocle", "sdk-Namespace-6261", "sdk-Topics-1984", "sdk-AuthRules-4310", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-01-01-preview/examples/Topics/SBTopicAuthorizationRuleListKey.json
func ExampleTopicsClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTopicsClient().ListKeys(ctx, "Default-ServiceBus-WestUS", "sdk-Namespace8408", "sdk-Topics2075", "sdk-Authrules5067", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessKeys = armservicebus.AccessKeys{
	// 	KeyName: to.Ptr("sdk-AuthRules-4310"),
	// 	PrimaryConnectionString: to.Ptr("Endpoint=sb://sdk-namespace-6261.servicebus.windows-int.net/;SharedAccessKeyName=sdk-AuthRules-4310;SharedAccessKey=#############################################;EntityPath=sdk-Topics-1984"),
	// 	PrimaryKey: to.Ptr("#############################################"),
	// 	SecondaryConnectionString: to.Ptr("Endpoint=sb://sdk-namespace-6261.servicebus.windows-int.net/;SharedAccessKeyName=sdk-AuthRules-4310;SharedAccessKey=#############################################;EntityPath=sdk-Topics-1984"),
	// 	SecondaryKey: to.Ptr("#############################################"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-01-01-preview/examples/Topics/SBTopicAuthorizationRuleRegenerateKey.json
func ExampleTopicsClient_RegenerateKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTopicsClient().RegenerateKeys(ctx, "Default-ServiceBus-WestUS", "sdk-Namespace8408", "sdk-Topics2075", "sdk-Authrules5067", armservicebus.RegenerateAccessKeyParameters{
		KeyType: to.Ptr(armservicebus.KeyTypePrimaryKey),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessKeys = armservicebus.AccessKeys{
	// 	KeyName: to.Ptr("sdk-AuthRules-4310"),
	// 	PrimaryConnectionString: to.Ptr("Endpoint=sb://sdk-namespace-6261.servicebus.windows-int.net/;SharedAccessKeyName=sdk-AuthRules-4310;SharedAccessKey=#############################################;EntityPath=sdk-Topics-1984"),
	// 	PrimaryKey: to.Ptr("#############################################"),
	// 	SecondaryConnectionString: to.Ptr("Endpoint=sb://sdk-namespace-6261.servicebus.windows-int.net/;SharedAccessKeyName=sdk-AuthRules-4310;SharedAccessKey=#############################################;EntityPath=sdk-Topics-1984"),
	// 	SecondaryKey: to.Ptr("#############################################"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-01-01-preview/examples/Topics/SBTopicListByNameSpace.json
func ExampleTopicsClient_NewListByNamespacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewTopicsClient().NewListByNamespacePager("Default-ServiceBus-WestUS", "sdk-Namespace-1617", &armservicebus.TopicsClientListByNamespaceOptions{Skip: nil,
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
		// page.SBTopicListResult = armservicebus.SBTopicListResult{
		// 	Value: []*armservicebus.SBTopic{
		// 		{
		// 			Name: to.Ptr("sdk-topics-5488"),
		// 			Type: to.Ptr("Microsoft.ServiceBus/Namespaces/Topics"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-1617/topics/sdk-topics-5488"),
		// 			Properties: &armservicebus.SBTopicProperties{
		// 				AccessedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T00:00:00Z"); return t}()),
		// 				AutoDeleteOnIdle: to.Ptr("P10675199DT2H48M5.4775807S"),
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-26T20:50:31.4442694Z"); return t}()),
		// 				DefaultMessageTimeToLive: to.Ptr("P10675199DT2H48M5.4775807S"),
		// 				DuplicateDetectionHistoryTimeWindow: to.Ptr("PT10M"),
		// 				EnableBatchedOperations: to.Ptr(true),
		// 				EnableExpress: to.Ptr(true),
		// 				EnablePartitioning: to.Ptr(false),
		// 				MaxMessageSizeInKilobytes: to.Ptr[int64](10240),
		// 				MaxSizeInMegabytes: to.Ptr[int32](10240),
		// 				RequiresDuplicateDetection: to.Ptr(false),
		// 				SizeInBytes: to.Ptr[int64](0),
		// 				Status: to.Ptr(armservicebus.EntityStatusActive),
		// 				SubscriptionCount: to.Ptr[int32](0),
		// 				SupportOrdering: to.Ptr(true),
		// 				UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-26T20:52:32.2092264Z"); return t}()),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-01-01-preview/examples/Topics/SBTopicCreate.json
func ExampleTopicsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTopicsClient().CreateOrUpdate(ctx, "ArunMonocle", "sdk-Namespace-1617", "sdk-Topics-5488", armservicebus.SBTopic{
		Properties: &armservicebus.SBTopicProperties{
			EnableExpress: to.Ptr(true),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SBTopic = armservicebus.SBTopic{
	// 	Name: to.Ptr("sdk-Topics-5488"),
	// 	Type: to.Ptr("Microsoft.ServiceBus/Namespaces/Topics"),
	// 	ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-1617/topics/sdk-Topics-5488"),
	// 	Properties: &armservicebus.SBTopicProperties{
	// 		AccessedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-26T20:50:34.32Z"); return t}()),
	// 		AutoDeleteOnIdle: to.Ptr("P10675199DT2H48M5.4775807S"),
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-26T20:50:34.1Z"); return t}()),
	// 		DefaultMessageTimeToLive: to.Ptr("P10675199DT2H48M5.4775807S"),
	// 		DuplicateDetectionHistoryTimeWindow: to.Ptr("PT10M"),
	// 		EnableBatchedOperations: to.Ptr(true),
	// 		EnableExpress: to.Ptr(true),
	// 		EnablePartitioning: to.Ptr(false),
	// 		MaxMessageSizeInKilobytes: to.Ptr[int64](10240),
	// 		MaxSizeInMegabytes: to.Ptr[int32](10240),
	// 		RequiresDuplicateDetection: to.Ptr(false),
	// 		SizeInBytes: to.Ptr[int64](0),
	// 		Status: to.Ptr(armservicebus.EntityStatusActive),
	// 		SubscriptionCount: to.Ptr[int32](0),
	// 		SupportOrdering: to.Ptr(true),
	// 		UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-26T20:50:34.32Z"); return t}()),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-01-01-preview/examples/Topics/SBTopicDelete.json
func ExampleTopicsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewTopicsClient().Delete(ctx, "ArunMonocle", "sdk-Namespace-1617", "sdk-Topics-5488", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/7a2ac91de424f271cf91cc8009f3fe9ee8249086/specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-01-01-preview/examples/Topics/SBTopicGet.json
func ExampleTopicsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicebus.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewTopicsClient().Get(ctx, "ArunMonocle", "sdk-Namespace-1617", "sdk-Topics-5488", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SBTopic = armservicebus.SBTopic{
	// 	Name: to.Ptr("sdk-Topics-5488"),
	// 	Type: to.Ptr("Microsoft.ServiceBus/Namespaces/Topics"),
	// 	ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.ServiceBus/namespaces/sdk-Namespace-1617/topics/sdk-Topics-5488"),
	// 	Properties: &armservicebus.SBTopicProperties{
	// 		AccessedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "0001-01-01T00:00:00Z"); return t}()),
	// 		AutoDeleteOnIdle: to.Ptr("P10675199DT2H48M5.4775807S"),
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-26T20:50:31.4442694Z"); return t}()),
	// 		DefaultMessageTimeToLive: to.Ptr("P10675199DT2H48M5.4775807S"),
	// 		DuplicateDetectionHistoryTimeWindow: to.Ptr("PT10M"),
	// 		EnableBatchedOperations: to.Ptr(true),
	// 		EnableExpress: to.Ptr(true),
	// 		EnablePartitioning: to.Ptr(false),
	// 		MaxMessageSizeInKilobytes: to.Ptr[int64](10240),
	// 		MaxSizeInMegabytes: to.Ptr[int32](10240),
	// 		RequiresDuplicateDetection: to.Ptr(false),
	// 		SizeInBytes: to.Ptr[int64](0),
	// 		Status: to.Ptr(armservicebus.EntityStatusActive),
	// 		SubscriptionCount: to.Ptr[int32](0),
	// 		SupportOrdering: to.Ptr(true),
	// 		UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-05-26T20:52:32.2092264Z"); return t}()),
	// 	},
	// }
}
