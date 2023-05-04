//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmigrate_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/Machines_ListByProject.json
func ExampleMachinesClient_NewListByProjectPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewMachinesClient().NewListByProjectPager("abgoyal-westEurope", "abgoyalWEselfhostb72bproject", nil)
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
		// page.MachineResultList = armmigrate.MachineResultList{
		// 	Value: []*armmigrate.Machine{
		// 		{
		// 			Name: to.Ptr("269ef295-a38d-4f8f-9779-77ce79088311"),
		// 			Type: to.Ptr("Microsoft.Migrate/assessmentprojects/machines"),
		// 			ETag: to.Ptr("\"04006052-0000-0d00-0000-5cd4065a0000\""),
		// 			ID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westeurope/providers/Microsoft.Migrate/assessmentprojects/abgoyalWEselfhostb72bproject/machines/269ef295-a38d-4f8f-9779-77ce79088311"),
		// 			Properties: &armmigrate.MachineProperties{
		// 				Description: to.Ptr("Microsoft Azure Migration Image on Windows Server 2016"),
		// 				BootType: to.Ptr(armmigrate.MachineBootTypeBIOS),
		// 				CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T09:58:22.1734149Z"); return t}()),
		// 				DatacenterManagementServerArmID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westEurope/providers/Microsoft.OffAzure/VMwareSites/PortalvCenterbc2fsite/vcenters/idclab-a360-fareast-corp-micros-86617dcf-effe-59ad-8c3a-cdd3ea7300d3"),
		// 				DatacenterManagementServerName: to.Ptr("IDCLAB-A360.fareast.corp.microsoft.com"),
		// 				DiscoveryMachineArmID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourcegroups/abgoyal-westeurope/providers/microsoft.offazure/vmwaresites/portalvcenterbc2fsite/machines/idclab-a360-fareast-corp-micros-86617dcf-effe-59ad-8c3a-cdd3ea7300d3_52bd4eeb-faf4-7d95-4dd5-5524350ce2bb"),
		// 				Disks: map[string]*armmigrate.Disk{
		// 					"6000C29f-9065-8fe0-ab83-7e58ff6ba442": &armmigrate.Disk{
		// 						DisplayName: to.Ptr("scsi0:0"),
		// 						GigabytesAllocated: to.Ptr[float64](80),
		// 					},
		// 				},
		// 				DisplayName: to.Ptr("ShubhamFirstAndThird"),
		// 				Groups: []*string{
		// 					to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westeurope/providers/Microsoft.Migrate/assessmentprojects/abgoyalWEselfhostb72bproject/groups/test1")},
		// 					MegabytesOfMemory: to.Ptr[float32](16384),
		// 					NetworkAdapters: map[string]*armmigrate.NetworkAdapter{
		// 						"4000": &armmigrate.NetworkAdapter{
		// 							DisplayName: to.Ptr("VM Network"),
		// 							IPAddresses: []*string{
		// 							},
		// 							MacAddress: to.Ptr("00:0c:29:ac:e3:6d"),
		// 						},
		// 					},
		// 					NumberOfCores: to.Ptr[int32](8),
		// 					OperatingSystemName: to.Ptr("Microsoft Windows Server 2016 (64-bit)"),
		// 					OperatingSystemType: to.Ptr("windowsGuest"),
		// 					UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T09:58:22.2990529Z"); return t}()),
		// 				},
		// 			},
		// 			{
		// 				Name: to.Ptr("3ad6c8b7-08d6-45dc-87f5-cd533501f553"),
		// 				Type: to.Ptr("Microsoft.Migrate/assessmentprojects/machines"),
		// 				ETag: to.Ptr("\"04005652-0000-0d00-0000-5cd4065a0000\""),
		// 				ID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westeurope/providers/Microsoft.Migrate/assessmentprojects/abgoyalWEselfhostb72bproject/machines/3ad6c8b7-08d6-45dc-87f5-cd533501f553"),
		// 				Properties: &armmigrate.MachineProperties{
		// 					Description: to.Ptr("Microsoft Azure Migration Image on Windows Server 2016"),
		// 					BootType: to.Ptr(armmigrate.MachineBootTypeBIOS),
		// 					CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T09:58:22.1734149Z"); return t}()),
		// 					DatacenterManagementServerArmID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westEurope/providers/Microsoft.OffAzure/VMwareSites/PortalvCenterbc2fsite/vcenters/idclab-a360-fareast-corp-micros-86617dcf-effe-59ad-8c3a-cdd3ea7300d3"),
		// 					DatacenterManagementServerName: to.Ptr("IDCLAB-A360.fareast.corp.microsoft.com"),
		// 					DiscoveryMachineArmID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourcegroups/abgoyal-westeurope/providers/microsoft.offazure/vmwaresites/portalvcenterbc2fsite/machines/idclab-a360-fareast-corp-micros-86617dcf-effe-59ad-8c3a-cdd3ea7300d3_52e5be9c-2758-a483-c252-eeef7919095c"),
		// 					Disks: map[string]*armmigrate.Disk{
		// 						"6000C296-54c9-e29c-51be-125e76850958": &armmigrate.Disk{
		// 							DisplayName: to.Ptr("scsi0:0"),
		// 							GigabytesAllocated: to.Ptr[float64](80),
		// 						},
		// 					},
		// 					DisplayName: to.Ptr("shsinglaVM5"),
		// 					Groups: []*string{
		// 						to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westeurope/providers/Microsoft.Migrate/assessmentprojects/abgoyalWEselfhostb72bproject/groups/test1")},
		// 						MegabytesOfMemory: to.Ptr[float32](16384),
		// 						NetworkAdapters: map[string]*armmigrate.NetworkAdapter{
		// 							"4000": &armmigrate.NetworkAdapter{
		// 								DisplayName: to.Ptr("VM Network"),
		// 								IPAddresses: []*string{
		// 								},
		// 								MacAddress: to.Ptr("00:0c:29:1f:ac:0a"),
		// 							},
		// 						},
		// 						NumberOfCores: to.Ptr[int32](8),
		// 						OperatingSystemName: to.Ptr("Microsoft Windows Server 2016 (64-bit)"),
		// 						OperatingSystemType: to.Ptr("windowsGuest"),
		// 						UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T09:58:22.3142054Z"); return t}()),
		// 					},
		// 			}},
		// 		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/163e27c0ca7570bc39e00a46f255740d9b3ba3cb/specification/migrate/resource-manager/Microsoft.Migrate/stable/2019-10-01/examples/Machines_Get.json
func ExampleMachinesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmigrate.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewMachinesClient().Get(ctx, "abgoyal-westEurope", "abgoyalWEselfhostb72bproject", "269ef295-a38d-4f8f-9779-77ce79088311", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Machine = armmigrate.Machine{
	// 	Name: to.Ptr("269ef295-a38d-4f8f-9779-77ce79088311"),
	// 	Type: to.Ptr("Microsoft.Migrate/assessmentprojects/machines"),
	// 	ETag: to.Ptr("\"04006052-0000-0d00-0000-5cd4065a0000\""),
	// 	ID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westeurope/providers/Microsoft.Migrate/assessmentprojects/abgoyalWEselfhostb72bproject/machines/269ef295-a38d-4f8f-9779-77ce79088311"),
	// 	Properties: &armmigrate.MachineProperties{
	// 		Description: to.Ptr("Microsoft Azure Migration Image on Windows Server 2016"),
	// 		BootType: to.Ptr(armmigrate.MachineBootTypeBIOS),
	// 		CreatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T09:58:22.1734149Z"); return t}()),
	// 		DatacenterManagementServerArmID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westEurope/providers/Microsoft.OffAzure/VMwareSites/PortalvCenterbc2fsite/vcenters/idclab-a360-fareast-corp-micros-86617dcf-effe-59ad-8c3a-cdd3ea7300d3"),
	// 		DatacenterManagementServerName: to.Ptr("IDCLAB-A360.fareast.corp.microsoft.com"),
	// 		DiscoveryMachineArmID: to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourcegroups/abgoyal-westeurope/providers/microsoft.offazure/vmwaresites/portalvcenterbc2fsite/machines/idclab-a360-fareast-corp-micros-86617dcf-effe-59ad-8c3a-cdd3ea7300d3_52bd4eeb-faf4-7d95-4dd5-5524350ce2bb"),
	// 		Disks: map[string]*armmigrate.Disk{
	// 			"6000C29f-9065-8fe0-ab83-7e58ff6ba442": &armmigrate.Disk{
	// 				DisplayName: to.Ptr("scsi0:0"),
	// 				GigabytesAllocated: to.Ptr[float64](80),
	// 			},
	// 		},
	// 		DisplayName: to.Ptr("ShubhamFirstAndThird"),
	// 		Groups: []*string{
	// 			to.Ptr("/subscriptions/6393a73f-8d55-47ef-b6dd-179b3e0c7910/resourceGroups/abgoyal-westeurope/providers/Microsoft.Migrate/assessmentprojects/abgoyalWEselfhostb72bproject/groups/test1")},
	// 			MegabytesOfMemory: to.Ptr[float32](16384),
	// 			NetworkAdapters: map[string]*armmigrate.NetworkAdapter{
	// 				"4000": &armmigrate.NetworkAdapter{
	// 					DisplayName: to.Ptr("VM Network"),
	// 					IPAddresses: []*string{
	// 					},
	// 					MacAddress: to.Ptr("00:0c:29:ac:e3:6d"),
	// 				},
	// 			},
	// 			NumberOfCores: to.Ptr[int32](8),
	// 			OperatingSystemName: to.Ptr("Microsoft Windows Server 2016 (64-bit)"),
	// 			OperatingSystemType: to.Ptr("windowsGuest"),
	// 			UpdatedTimestamp: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2019-05-09T09:58:22.2990529Z"); return t}()),
	// 		},
	// 	}
}
