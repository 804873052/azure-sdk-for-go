//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhealthcareapis_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis"
)

// x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2021-06-01-preview/examples/dicomservices/DicomServices_List.json
func ExampleDicomServicesClient_ListByWorkspace() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhealthcareapis.NewDicomServicesClient("<subscription-id>", cred, nil)
	pager := client.ListByWorkspace("<resource-group-name>",
		"<workspace-name>",
		nil)
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("DicomService.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2021-06-01-preview/examples/dicomservices/DicomServices_Get.json
func ExampleDicomServicesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhealthcareapis.NewDicomServicesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<dicom-service-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DicomService.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2021-06-01-preview/examples/dicomservices/DicomServices_Create.json
func ExampleDicomServicesClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhealthcareapis.NewDicomServicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<workspace-name>",
		"<dicom-service-name>",
		armhealthcareapis.DicomService{
			TaggedResource: armhealthcareapis.TaggedResource{
				LocationBasedResource: armhealthcareapis.LocationBasedResource{
					Location: to.StringPtr("<location>"),
				},
			},
			Properties: &armhealthcareapis.DicomServiceProperties{},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DicomService.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2021-06-01-preview/examples/dicomservices/DicomServices_Patch.json
func ExampleDicomServicesClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhealthcareapis.NewDicomServicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<dicom-service-name>",
		"<workspace-name>",
		armhealthcareapis.DicomServicePatchResource{
			ResourceTags: armhealthcareapis.ResourceTags{
				Tags: map[string]*string{
					"tagKey": to.StringPtr("tagValue"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("DicomService.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/preview/2021-06-01-preview/examples/dicomservices/DicomServices_Delete.json
func ExampleDicomServicesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armhealthcareapis.NewDicomServicesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<dicom-service-name>",
		"<workspace-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}
