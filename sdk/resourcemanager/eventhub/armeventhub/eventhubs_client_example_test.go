//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeventhub_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/eventhub/armeventhub"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bd2d190bc2aad1e8b30c1ffa8aea94f2d4715b76/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/EventHubs/EHEventHubAuthorizationRuleListAll.json
func ExampleEventHubsClient_NewListAuthorizationRulesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEventHubsClient().NewListAuthorizationRulesPager("ArunMonocle", "sdk-Namespace-960", "sdk-EventHub-532", nil)
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
		// page.AuthorizationRuleListResult = armeventhub.AuthorizationRuleListResult{
		// 	Value: []*armeventhub.AuthorizationRule{
		// 		{
		// 			Name: to.Ptr("sdk-Authrules-2513"),
		// 			Type: to.Ptr("Microsoft.EventHub/Namespaces/EventHubs/AuthorizationRules"),
		// 			ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.EventHub/namespaces/sdk-Namespace-960/eventhubs/sdk-EventHub-532/authorizationRules/sdk-Authrules-2513"),
		// 			Properties: &armeventhub.AuthorizationRuleProperties{
		// 				Rights: []*armeventhub.AccessRights{
		// 					to.Ptr(armeventhub.AccessRightsListen),
		// 					to.Ptr(armeventhub.AccessRightsSend)},
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bd2d190bc2aad1e8b30c1ffa8aea94f2d4715b76/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/EventHubs/EHEventHubAuthorizationRuleCreate.json
func ExampleEventHubsClient_CreateOrUpdateAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEventHubsClient().CreateOrUpdateAuthorizationRule(ctx, "ArunMonocle", "sdk-Namespace-960", "sdk-EventHub-532", "sdk-Authrules-2513", armeventhub.AuthorizationRule{
		Properties: &armeventhub.AuthorizationRuleProperties{
			Rights: []*armeventhub.AccessRights{
				to.Ptr(armeventhub.AccessRightsListen),
				to.Ptr(armeventhub.AccessRightsSend)},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizationRule = armeventhub.AuthorizationRule{
	// 	Name: to.Ptr("sdk-Authrules-2513"),
	// 	Type: to.Ptr("Microsoft.EventHub/Namespaces/EventHubs/AuthorizationRules"),
	// 	ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.EventHub/namespaces/sdk-Namespace-960/eventhubs/sdk-EventHub-532/authorizationRules/sdk-Authrules-2513"),
	// 	Properties: &armeventhub.AuthorizationRuleProperties{
	// 		Rights: []*armeventhub.AccessRights{
	// 			to.Ptr(armeventhub.AccessRightsListen),
	// 			to.Ptr(armeventhub.AccessRightsSend)},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bd2d190bc2aad1e8b30c1ffa8aea94f2d4715b76/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/EventHubs/EHEventHubAuthorizationRuleGet.json
func ExampleEventHubsClient_GetAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEventHubsClient().GetAuthorizationRule(ctx, "ArunMonocle", "sdk-Namespace-960", "sdk-EventHub-532", "sdk-Authrules-2513", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AuthorizationRule = armeventhub.AuthorizationRule{
	// 	Name: to.Ptr("sdk-Authrules-2513"),
	// 	Type: to.Ptr("Microsoft.EventHub/Namespaces/EventHubs/AuthorizationRules"),
	// 	ID: to.Ptr("/subscriptions/5f750a97-50d9-4e36-8081-c9ee4c0210d4/resourceGroups/ArunMonocle/providers/Microsoft.EventHub/namespaces/sdk-Namespace-960/eventhubs/sdk-EventHub-532/authorizationRules/sdk-Authrules-2513"),
	// 	Properties: &armeventhub.AuthorizationRuleProperties{
	// 		Rights: []*armeventhub.AccessRights{
	// 			to.Ptr(armeventhub.AccessRightsListen),
	// 			to.Ptr(armeventhub.AccessRightsSend)},
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bd2d190bc2aad1e8b30c1ffa8aea94f2d4715b76/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/EventHubs/EHEventHubAuthorizationRuleDelete.json
func ExampleEventHubsClient_DeleteAuthorizationRule() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewEventHubsClient().DeleteAuthorizationRule(ctx, "ArunMonocle", "sdk-Namespace-960", "sdk-EventHub-532", "sdk-Authrules-2513", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bd2d190bc2aad1e8b30c1ffa8aea94f2d4715b76/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/EventHubs/EHEventHubAuthorizationRuleListKey.json
func ExampleEventHubsClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEventHubsClient().ListKeys(ctx, "ArunMonocle", "sdk-namespace-960", "sdk-EventHub-532", "sdk-Authrules-2513", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessKeys = armeventhub.AccessKeys{
	// 	KeyName: to.Ptr("sdk-Authrules-2513"),
	// 	PrimaryConnectionString: to.Ptr("Endpoint=sb://sdk-namespace-960.servicebus.windows-int.net/;SharedAccessKeyName=sdk-Authrules-2513;SharedAccessKey=############################################;EntityPath=sdk-EventHub-532"),
	// 	PrimaryKey: to.Ptr("############################################"),
	// 	SecondaryConnectionString: to.Ptr("Endpoint=sb://sdk-namespace-960.servicebus.windows-int.net/;SharedAccessKeyName=sdk-Authrules-2513;SharedAccessKey=############################################;EntityPath=sdk-EventHub-532"),
	// 	SecondaryKey: to.Ptr("############################################"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bd2d190bc2aad1e8b30c1ffa8aea94f2d4715b76/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/EventHubs/EHEventHubAuthorizationRuleRegenerateKey.json
func ExampleEventHubsClient_RegenerateKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEventHubsClient().RegenerateKeys(ctx, "ArunMonocle", "sdk-namespace-960", "sdk-EventHub-532", "sdk-Authrules-1534", armeventhub.RegenerateAccessKeyParameters{
		KeyType: to.Ptr(armeventhub.KeyTypePrimaryKey),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccessKeys = armeventhub.AccessKeys{
	// 	KeyName: to.Ptr("sdk-Authrules-1534"),
	// 	PrimaryConnectionString: to.Ptr("Endpoint=sb://sdk-namespace-9027.servicebus.windows-int.net/;SharedAccessKeyName=sdk-Authrules-1534;SharedAccessKey=#############################################;EntityPath=sdk-EventHub-1647"),
	// 	PrimaryKey: to.Ptr("#############################################"),
	// 	SecondaryConnectionString: to.Ptr("Endpoint=sb://sdk-namespace-9027.servicebus.windows-int.net/;SharedAccessKeyName=sdk-Authrules-1534;SharedAccessKey=#############################################;EntityPath=sdk-EventHub-1647"),
	// 	SecondaryKey: to.Ptr("#############################################"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bd2d190bc2aad1e8b30c1ffa8aea94f2d4715b76/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/EventHubs/EHEventHubListByNameSpace.json
func ExampleEventHubsClient_NewListByNamespacePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewEventHubsClient().NewListByNamespacePager("Default-NotificationHubs-AustraliaEast", "sdk-Namespace-5357", &armeventhub.EventHubsClientListByNamespaceOptions{Skip: nil,
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
		// page.ListResult = armeventhub.ListResult{
		// 	Value: []*armeventhub.Eventhub{
		// 		{
		// 			Name: to.Ptr("sdk-eventhub-10"),
		// 			Type: to.Ptr("Microsoft.EventHub/Namespaces/EventHubs"),
		// 			ID: to.Ptr("/subscriptions/e2f361f0-3b27-4503-a9cc-21cfba380093/resourceGroups/Default-NotificationHubs-AustraliaEast/providers/Microsoft.EventHub/namespaces/sdk-Namespace-716/eventhubs/sdk-eventhub-10"),
		// 			Properties: &armeventhub.Properties{
		// 				CaptureDescription: &armeventhub.CaptureDescription{
		// 					Destination: &armeventhub.Destination{
		// 						Name: to.Ptr("EventHubArchive.AzureBlockBlob"),
		// 						Properties: &armeventhub.DestinationProperties{
		// 							ArchiveNameFormat: to.Ptr("{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}"),
		// 							BlobContainer: to.Ptr("container"),
		// 							StorageAccountResourceID: to.Ptr("/subscriptions/e2f361f0-3b27-4503-a9cc-21cfba380093/resourceGroups/Default-Storage-SouthCentralUS/providers/Microsoft.ClassicStorage/storageAccounts/arjunteststorage"),
		// 						},
		// 					},
		// 					Enabled: to.Ptr(true),
		// 					Encoding: to.Ptr(armeventhub.EncodingCaptureDescriptionAvro),
		// 					IntervalInSeconds: to.Ptr[int32](120),
		// 					SizeLimitInBytes: to.Ptr[int32](10485763),
		// 				},
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-28T02:45:55.877Z"); return t}()),
		// 				MessageRetentionInDays: to.Ptr[int64](4),
		// 				PartitionCount: to.Ptr[int64](4),
		// 				PartitionIDs: []*string{
		// 					to.Ptr("0"),
		// 					to.Ptr("1"),
		// 					to.Ptr("2"),
		// 					to.Ptr("3")},
		// 					RetentionDescription: &armeventhub.RetentionDescription{
		// 						CleanupPolicy: to.Ptr(armeventhub.CleanupPolicyRetentionDescriptionDelete),
		// 						RetentionTimeInHours: to.Ptr[int64](96),
		// 						TombstoneRetentionTimeInHours: to.Ptr[int32](1),
		// 					},
		// 					Status: to.Ptr(armeventhub.EntityStatusActive),
		// 					UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-28T02:46:05.877Z"); return t}()),
		// 				},
		// 		}},
		// 	}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bd2d190bc2aad1e8b30c1ffa8aea94f2d4715b76/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/EventHubs/EHEventHubCreate.json
func ExampleEventHubsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEventHubsClient().CreateOrUpdate(ctx, "Default-NotificationHubs-AustraliaEast", "sdk-Namespace-5357", "sdk-EventHub-6547", armeventhub.Eventhub{
		Properties: &armeventhub.Properties{
			CaptureDescription: &armeventhub.CaptureDescription{
				Destination: &armeventhub.Destination{
					Name: to.Ptr("EventHubArchive.AzureBlockBlob"),
					Identity: &armeventhub.CaptureIdentity{
						Type:                 to.Ptr(armeventhub.CaptureIdentityTypeUserAssigned),
						UserAssignedIdentity: to.Ptr("/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud2"),
					},
					Properties: &armeventhub.DestinationProperties{
						ArchiveNameFormat:        to.Ptr("{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}"),
						BlobContainer:            to.Ptr("container"),
						StorageAccountResourceID: to.Ptr("/subscriptions/e2f361f0-3b27-4503-a9cc-21cfba380093/resourceGroups/Default-Storage-SouthCentralUS/providers/Microsoft.ClassicStorage/storageAccounts/arjunteststorage"),
					},
				},
				Enabled:           to.Ptr(true),
				Encoding:          to.Ptr(armeventhub.EncodingCaptureDescriptionAvro),
				IntervalInSeconds: to.Ptr[int32](120),
				SizeLimitInBytes:  to.Ptr[int32](10485763),
			},
			MessageRetentionInDays: to.Ptr[int64](4),
			PartitionCount:         to.Ptr[int64](4),
			RetentionDescription: &armeventhub.RetentionDescription{
				CleanupPolicy:                 to.Ptr(armeventhub.CleanupPolicyRetentionDescriptionCompact),
				RetentionTimeInHours:          to.Ptr[int64](96),
				TombstoneRetentionTimeInHours: to.Ptr[int32](1),
			},
			Status:       to.Ptr(armeventhub.EntityStatusActive),
			UserMetadata: to.Ptr("key"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Eventhub = armeventhub.Eventhub{
	// 	Name: to.Ptr("sdk-EventHub-10"),
	// 	Type: to.Ptr("Microsoft.EventHub/Namespaces/EventHubs"),
	// 	ID: to.Ptr("/subscriptions/e2f361f0-3b27-4503-a9cc-21cfba380093/resourceGroups/Default-NotificationHubs-AustraliaEast/providers/Microsoft.EventHub/namespaces/sdk-Namespace-716/eventhubs/sdk-EventHub-10"),
	// 	Properties: &armeventhub.Properties{
	// 		CaptureDescription: &armeventhub.CaptureDescription{
	// 			Destination: &armeventhub.Destination{
	// 				Name: to.Ptr("EventHubArchive.AzureBlockBlob"),
	// 				Identity: &armeventhub.CaptureIdentity{
	// 					Type: to.Ptr(armeventhub.CaptureIdentityTypeUserAssigned),
	// 					UserAssignedIdentity: to.Ptr("/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud2"),
	// 				},
	// 				Properties: &armeventhub.DestinationProperties{
	// 					ArchiveNameFormat: to.Ptr("{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}"),
	// 					BlobContainer: to.Ptr("container"),
	// 					StorageAccountResourceID: to.Ptr("/subscriptions/e2f361f0-3b27-4503-a9cc-21cfba380093/resourceGroups/Default-Storage-SouthCentralUS/providers/Microsoft.ClassicStorage/storageAccounts/arjunteststorage"),
	// 				},
	// 			},
	// 			Enabled: to.Ptr(true),
	// 			Encoding: to.Ptr(armeventhub.EncodingCaptureDescriptionAvro),
	// 			IntervalInSeconds: to.Ptr[int32](120),
	// 			SizeLimitInBytes: to.Ptr[int32](10485763),
	// 		},
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-28T02:45:55.877Z"); return t}()),
	// 		MessageRetentionInDays: to.Ptr[int64](4),
	// 		PartitionCount: to.Ptr[int64](4),
	// 		PartitionIDs: []*string{
	// 			to.Ptr("0"),
	// 			to.Ptr("1"),
	// 			to.Ptr("2"),
	// 			to.Ptr("3")},
	// 			RetentionDescription: &armeventhub.RetentionDescription{
	// 				CleanupPolicy: to.Ptr(armeventhub.CleanupPolicyRetentionDescriptionCompact),
	// 				RetentionTimeInHours: to.Ptr[int64](96),
	// 				TombstoneRetentionTimeInHours: to.Ptr[int32](1),
	// 			},
	// 			Status: to.Ptr(armeventhub.EntityStatusActive),
	// 			UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-28T02:46:05.877Z"); return t}()),
	// 			UserMetadata: to.Ptr("key"),
	// 		},
	// 	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bd2d190bc2aad1e8b30c1ffa8aea94f2d4715b76/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/EventHubs/EHEventHubDelete.json
func ExampleEventHubsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewEventHubsClient().Delete(ctx, "ArunMonocle", "sdk-Namespace-5357", "sdk-EventHub-6547", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/bd2d190bc2aad1e8b30c1ffa8aea94f2d4715b76/specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/EventHubs/EHEventHubGet.json
func ExampleEventHubsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armeventhub.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewEventHubsClient().Get(ctx, "Default-NotificationHubs-AustraliaEast", "sdk-Namespace-716", "sdk-EventHub-10", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Eventhub = armeventhub.Eventhub{
	// 	Name: to.Ptr("sdk-EventHub-10"),
	// 	Type: to.Ptr("Microsoft.EventHub/Namespaces/EventHubs"),
	// 	ID: to.Ptr("/subscriptions/e2f361f0-3b27-4503-a9cc-21cfba380093/resourceGroups/Default-NotificationHubs-AustraliaEast/providers/Microsoft.EventHub/namespaces/sdk-Namespace-716/eventhubs/sdk-EventHub-10"),
	// 	Properties: &armeventhub.Properties{
	// 		CaptureDescription: &armeventhub.CaptureDescription{
	// 			Destination: &armeventhub.Destination{
	// 				Name: to.Ptr("EventHubArchive.AzureBlockBlob"),
	// 				Properties: &armeventhub.DestinationProperties{
	// 					ArchiveNameFormat: to.Ptr("{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}"),
	// 					BlobContainer: to.Ptr("container"),
	// 					StorageAccountResourceID: to.Ptr("/subscriptions/e2f361f0-3b27-4503-a9cc-21cfba380093/resourceGroups/Default-Storage-SouthCentralUS/providers/Microsoft.ClassicStorage/storageAccounts/arjunteststorage"),
	// 				},
	// 			},
	// 			Enabled: to.Ptr(true),
	// 			Encoding: to.Ptr(armeventhub.EncodingCaptureDescriptionAvro),
	// 			IntervalInSeconds: to.Ptr[int32](120),
	// 			SizeLimitInBytes: to.Ptr[int32](10485763),
	// 		},
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-28T02:45:55.877Z"); return t}()),
	// 		MessageRetentionInDays: to.Ptr[int64](4),
	// 		PartitionCount: to.Ptr[int64](4),
	// 		PartitionIDs: []*string{
	// 			to.Ptr("0"),
	// 			to.Ptr("1"),
	// 			to.Ptr("2"),
	// 			to.Ptr("3")},
	// 			RetentionDescription: &armeventhub.RetentionDescription{
	// 				CleanupPolicy: to.Ptr(armeventhub.CleanupPolicyRetentionDescriptionCompact),
	// 				RetentionTimeInHours: to.Ptr[int64](96),
	// 				TombstoneRetentionTimeInHours: to.Ptr[int32](1),
	// 			},
	// 			Status: to.Ptr(armeventhub.EntityStatusActive),
	// 			UpdatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-06-28T02:46:05.877Z"); return t}()),
	// 		},
	// 	}
}
