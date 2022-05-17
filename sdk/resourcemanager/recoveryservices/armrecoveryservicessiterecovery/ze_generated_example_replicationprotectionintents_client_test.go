//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/recoveryservices/armrecoveryservicessiterecovery"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectionIntents_List.json
func ExampleReplicationProtectionIntentsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectionIntentsClient("2007vttp",
		"resourceGroupPS1",
		"509099b2-9d2c-4636-b43e-bd5cafb6be69", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager(&armrecoveryservicessiterecovery.ReplicationProtectionIntentsClientListOptions{SkipToken: nil,
		TakeToken: nil,
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectionIntents_Get.json
func ExampleReplicationProtectionIntentsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectionIntentsClient("vault1",
		"resourceGroupPS1",
		"509099b2-9d2c-4636-b43e-bd5cafb6be69", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"vm1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/ReplicationProtectionIntents_Create.json
func ExampleReplicationProtectionIntentsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armrecoveryservicessiterecovery.NewReplicationProtectionIntentsClient("vault1",
		"resourceGroupPS1",
		"509099b2-9d2c-4636-b43e-bd5cafb6be69", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Create(ctx,
		"vm1",
		armrecoveryservicessiterecovery.CreateProtectionIntentInput{
			Properties: &armrecoveryservicessiterecovery.CreateProtectionIntentProperties{
				ProviderSpecificDetails: &armrecoveryservicessiterecovery.A2ACreateProtectionIntentInput{
					InstanceType:             to.Ptr("A2A"),
					FabricObjectID:           to.Ptr("/subscriptions/509099b2-9d2c-4636-b43e-bd5cafb6be69/resourceGroups/removeOne/providers/Microsoft.Compute/virtualMachines/vmPpgAv5"),
					PrimaryLocation:          to.Ptr("eastUs2"),
					RecoveryAvailabilityType: to.Ptr(armrecoveryservicessiterecovery.A2ARecoveryAvailabilityTypeSingle),
					RecoveryLocation:         to.Ptr("westus2"),
					RecoveryResourceGroupID:  to.Ptr("/subscriptions/509099b2-9d2c-4636-b43e-bd5cafb6be69/resourceGroups/removeOne-asr"),
					RecoverySubscriptionID:   to.Ptr("ed5bcdf6-d61e-47bd-8ea9-f2bd379a2640"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
