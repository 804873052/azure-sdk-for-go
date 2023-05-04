//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmixedreality_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mixedreality/armmixedreality"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/GetBySubscription.json
func ExampleRemoteRenderingAccountsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmixedreality.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRemoteRenderingAccountsClient().NewListBySubscriptionPager(nil)
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
		// page.RemoteRenderingAccountPage = armmixedreality.RemoteRenderingAccountPage{
		// 	Value: []*armmixedreality.RemoteRenderingAccount{
		// 		{
		// 			Name: to.Ptr("alpha"),
		// 			Type: to.Ptr("Microsoft.MixedReality/remoteRenderingAccounts"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/MyResourceGroup/providers/Microsoft.MixedReality/remoteRenderingAccounts/alpha"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armmixedreality.Identity{
		// 				Type: to.Ptr("SystemAssigned"),
		// 			},
		// 			Properties: &armmixedreality.AccountProperties{
		// 				AccountDomain: to.Ptr("mixedreality.azure.com"),
		// 				AccountID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("omega"),
		// 			Type: to.Ptr("Microsoft.MixedReality/remoteRenderingAccounts"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/MyResourceGroup/providers/Microsoft.MixedReality/remoteRenderingAccounts/omega"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armmixedreality.Identity{
		// 				Type: to.Ptr("SystemAssigned"),
		// 			},
		// 			Properties: &armmixedreality.AccountProperties{
		// 				AccountDomain: to.Ptr("mixedreality.azure.com"),
		// 				AccountID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/GetByResourceGroup.json
func ExampleRemoteRenderingAccountsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmixedreality.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewRemoteRenderingAccountsClient().NewListByResourceGroupPager("MyResourceGroup", nil)
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
		// page.RemoteRenderingAccountPage = armmixedreality.RemoteRenderingAccountPage{
		// 	Value: []*armmixedreality.RemoteRenderingAccount{
		// 		{
		// 			Name: to.Ptr("alpha"),
		// 			Type: to.Ptr("Microsoft.MixedReality/remoteRenderingAccounts"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/MyResourceGroup/providers/Microsoft.MixedReality/remoteRenderingAccounts/alpha"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armmixedreality.Identity{
		// 				Type: to.Ptr("SystemAssigned"),
		// 			},
		// 			Properties: &armmixedreality.AccountProperties{
		// 				AccountDomain: to.Ptr("mixedreality.azure.com"),
		// 				AccountID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("omega"),
		// 			Type: to.Ptr("Microsoft.MixedReality/remoteRenderingAccounts"),
		// 			ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/MyResourceGroup/providers/Microsoft.MixedReality/remoteRenderingAccounts/omega"),
		// 			Location: to.Ptr("eastus2euap"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armmixedreality.Identity{
		// 				Type: to.Ptr("SystemAssigned"),
		// 			},
		// 			Properties: &armmixedreality.AccountProperties{
		// 				AccountDomain: to.Ptr("mixedreality.azure.com"),
		// 				AccountID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/Delete.json
func ExampleRemoteRenderingAccountsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmixedreality.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewRemoteRenderingAccountsClient().Delete(ctx, "MyResourceGroup", "MyAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/Get.json
func ExampleRemoteRenderingAccountsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmixedreality.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRemoteRenderingAccountsClient().Get(ctx, "MyResourceGroup", "MyAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RemoteRenderingAccount = armmixedreality.RemoteRenderingAccount{
	// 	Name: to.Ptr("MyAccount"),
	// 	Type: to.Ptr("Microsoft.MixedReality/remoteRenderingAccounts"),
	// 	ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/MyResourceGroup/providers/Microsoft.MixedReality/remoteRenderingAccounts/MyAccount"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Identity: &armmixedreality.Identity{
	// 		Type: to.Ptr("SystemAssigned"),
	// 	},
	// 	Properties: &armmixedreality.AccountProperties{
	// 		AccountDomain: to.Ptr("mixedreality.azure.com"),
	// 		AccountID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/Patch.json
func ExampleRemoteRenderingAccountsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmixedreality.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRemoteRenderingAccountsClient().Update(ctx, "MyResourceGroup", "MyAccount", armmixedreality.RemoteRenderingAccount{
		Location: to.Ptr("eastus2euap"),
		Tags: map[string]*string{
			"hero":    to.Ptr("romeo"),
			"heroine": to.Ptr("juliet"),
		},
		Identity: &armmixedreality.Identity{
			Type: to.Ptr("SystemAssigned"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RemoteRenderingAccount = armmixedreality.RemoteRenderingAccount{
	// 	Name: to.Ptr("MyAccount"),
	// 	Type: to.Ptr("Microsoft.MixedReality/remoteRenderingAccounts"),
	// 	ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/MyResourceGroup/providers/Microsoft.MixedReality/remoteRenderingAccounts/MyAccount"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Tags: map[string]*string{
	// 		"hero": to.Ptr("romeo"),
	// 		"heroine": to.Ptr("juliet"),
	// 	},
	// 	Identity: &armmixedreality.Identity{
	// 		Type: to.Ptr("SystemAssigned"),
	// 	},
	// 	Properties: &armmixedreality.AccountProperties{
	// 		AccountDomain: to.Ptr("mixedreality.azure.com"),
	// 		AccountID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/Put.json
func ExampleRemoteRenderingAccountsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmixedreality.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRemoteRenderingAccountsClient().Create(ctx, "MyResourceGroup", "MyAccount", armmixedreality.RemoteRenderingAccount{
		Location: to.Ptr("eastus2euap"),
		Identity: &armmixedreality.Identity{
			Type: to.Ptr("SystemAssigned"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.RemoteRenderingAccount = armmixedreality.RemoteRenderingAccount{
	// 	Name: to.Ptr("MyAccount"),
	// 	Type: to.Ptr("Microsoft.MixedReality/remoteRenderingAccounts"),
	// 	ID: to.Ptr("/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/MyResourceGroup/providers/Microsoft.MixedReality/remoteRenderingAccounts/MyAccount"),
	// 	Location: to.Ptr("eastus2euap"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Identity: &armmixedreality.Identity{
	// 		Type: to.Ptr("SystemAssigned"),
	// 	},
	// 	Properties: &armmixedreality.AccountProperties{
	// 		AccountDomain: to.Ptr("mixedreality.azure.com"),
	// 		AccountID: to.Ptr("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/ListKeys.json
func ExampleRemoteRenderingAccountsClient_ListKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmixedreality.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRemoteRenderingAccountsClient().ListKeys(ctx, "MyResourceGroup", "MyAccount", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccountKeys = armmixedreality.AccountKeys{
	// 	PrimaryKey: to.Ptr("<primaryKey>"),
	// 	SecondaryKey: to.Ptr("<secondaryKey>"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/mixedreality/resource-manager/Microsoft.MixedReality/preview/2021-03-01-preview/examples/remote-rendering/RegenerateKey.json
func ExampleRemoteRenderingAccountsClient_RegenerateKeys() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmixedreality.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewRemoteRenderingAccountsClient().RegenerateKeys(ctx, "MyResourceGroup", "MyAccount", armmixedreality.AccountKeyRegenerateRequest{
		Serial: to.Ptr(armmixedreality.SerialPrimary),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.AccountKeys = armmixedreality.AccountKeys{
	// 	PrimaryKey: to.Ptr("************"),
	// 	SecondaryKey: to.Ptr("************"),
	// }
}
