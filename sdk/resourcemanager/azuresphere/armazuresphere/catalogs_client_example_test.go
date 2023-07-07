//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// DO NOT EDIT.

package armazuresphere_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azuresphere/armazuresphere"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/GetCatalogsSub.json
func ExampleCatalogsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazuresphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCatalogsClient().NewListBySubscriptionPager(nil)
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
		// page.CatalogListResult = armazuresphere.CatalogListResult{
		// 	Value: []*armazuresphere.Catalog{
		// 		{
		// 			Name: to.Ptr("MyCatalog1"),
		// 			Type: to.Ptr("Microsoft.AzureSphere/catalogs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup1/providers/Microsoft.AzureSphere/catalogs/MyCatalog1"),
		// 			Location: to.Ptr("global"),
		// 		},
		// 		{
		// 			Name: to.Ptr("MyCatalog2"),
		// 			Type: to.Ptr("Microsoft.AzureSphere/catalogs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup2/providers/Microsoft.AzureSphere/catalogs/MyCatalog2"),
		// 			Location: to.Ptr("global"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/GetCatalogsRG.json
func ExampleCatalogsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazuresphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCatalogsClient().NewListByResourceGroupPager("MyResourceGroup1", nil)
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
		// page.CatalogListResult = armazuresphere.CatalogListResult{
		// 	Value: []*armazuresphere.Catalog{
		// 		{
		// 			Name: to.Ptr("MyCatalog1"),
		// 			Type: to.Ptr("Microsoft.AzureSphere/catalogs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup1/providers/Microsoft.AzureSphere/catalogs/MyCatalog1"),
		// 			Location: to.Ptr("global"),
		// 		},
		// 		{
		// 			Name: to.Ptr("MyCatalog2"),
		// 			Type: to.Ptr("Microsoft.AzureSphere/catalogs"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup1/providers/Microsoft.AzureSphere/catalogs/MyCatalog2"),
		// 			Location: to.Ptr("global"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/GetCatalog.json
func ExampleCatalogsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazuresphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCatalogsClient().Get(ctx, "MyResourceGroup1", "MyCatalog1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Catalog = armazuresphere.Catalog{
	// 	Name: to.Ptr("MyCatalog1"),
	// 	Type: to.Ptr("Microsoft.AzureSphere/catalogs"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup1/providers/Microsoft.AzureSphere/catalogs/MyCatalog1"),
	// 	Location: to.Ptr("global"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PutCatalog.json
func ExampleCatalogsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazuresphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCatalogsClient().BeginCreateOrUpdate(ctx, "MyResourceGroup1", "MyCatalog1", armazuresphere.Catalog{
		Location: to.Ptr("global"),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Catalog = armazuresphere.Catalog{
	// 	Name: to.Ptr("MyCatalog1"),
	// 	Type: to.Ptr("Microsoft.AzureSphere/catalogs"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup1/providers/Microsoft.AzureSphere/catalogs/MyCatalog1"),
	// 	Location: to.Ptr("global"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PatchCatalog.json
func ExampleCatalogsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazuresphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCatalogsClient().Update(ctx, "MyResourceGroup1", "MyCatalog1", armazuresphere.CatalogUpdate{}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.Catalog = armazuresphere.Catalog{
	// 	Name: to.Ptr("MyCatalog1"),
	// 	Type: to.Ptr("Microsoft.AzureSphere/catalogs"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup1/providers/Microsoft.AzureSphere/catalogs/MyCatalog1"),
	// 	Location: to.Ptr("global"),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/DeleteCatalog.json
func ExampleCatalogsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazuresphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewCatalogsClient().BeginDelete(ctx, "MyResourceGroup1", "MyCatalog1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PostCountDevicesCatalog.json
func ExampleCatalogsClient_CountDevices() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazuresphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewCatalogsClient().CountDevices(ctx, "MyResourceGroup1", "MyCatalog1", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.CountDeviceResponse = armazuresphere.CountDeviceResponse{
	// 	Value: to.Ptr[int32](3),
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PostListDeploymentsByCatalog.json
func ExampleCatalogsClient_NewListDeploymentsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazuresphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCatalogsClient().NewListDeploymentsPager("MyResourceGroup1", "MyCatalog1", &armazuresphere.CatalogsClientListDeploymentsOptions{Filter: nil,
		Top:         nil,
		Skip:        nil,
		Maxpagesize: nil,
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
		// page.DeploymentListResult = armazuresphere.DeploymentListResult{
		// 	Value: []*armazuresphere.Deployment{
		// 		{
		// 			Name: to.Ptr("DeploymentName1111"),
		// 			Properties: &armazuresphere.DeploymentProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("DeploymentName1121"),
		// 			Properties: &armazuresphere.DeploymentProperties{
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PostListDeviceGroupsCatalog.json
func ExampleCatalogsClient_NewListDeviceGroupsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazuresphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCatalogsClient().NewListDeviceGroupsPager("MyResourceGroup1", "MyCatalog1", armazuresphere.ListDeviceGroupsRequest{
		DeviceGroupName: to.Ptr("MyDeviceGroup1"),
	}, &armazuresphere.CatalogsClientListDeviceGroupsOptions{Filter: nil,
		Top:         nil,
		Skip:        nil,
		Maxpagesize: nil,
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
		// page.DeviceGroupListResult = armazuresphere.DeviceGroupListResult{
		// 	Value: []*armazuresphere.DeviceGroup{
		// 		{
		// 			Name: to.Ptr("MyDeviceGroup1"),
		// 			Type: to.Ptr("microsoft.azureSphere/catalogs/products/devicegroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup1/providers/Microsoft.AzureSphere/catalogs/MyCatalog1/products/MyProduct1/devicegroups/MyDeviceGroup1"),
		// 		},
		// 		{
		// 			Name: to.Ptr("MyDeviceGroup2"),
		// 			Type: to.Ptr("microsoft.azureSphere/catalogs/products/devicegroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/MyResourceGroup1/providers/Microsoft.AzureSphere/catalogs/MyCatalog1/products/MyProduct2/devicegroups/MyDeviceGroup2"),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PostListDeviceInsightsCatalog.json
func ExampleCatalogsClient_NewListDeviceInsightsPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazuresphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCatalogsClient().NewListDeviceInsightsPager("MyResourceGroup1", "MyCatalog1", &armazuresphere.CatalogsClientListDeviceInsightsOptions{Filter: nil,
		Top:         to.Ptr[int32](10),
		Skip:        nil,
		Maxpagesize: nil,
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
		// page.PagedDeviceInsight = armazuresphere.PagedDeviceInsight{
		// 	Value: []*armazuresphere.DeviceInsight{
		// 		{
		// 			Description: to.Ptr("eventDescription1"),
		// 			DeviceID: to.Ptr("eventIdentifier1"),
		// 			EndTimestampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-30T23:54:21.96Z"); return t}()),
		// 			EventCategory: to.Ptr("eventCategory1"),
		// 			EventClass: to.Ptr("eventClass1"),
		// 			EventCount: to.Ptr[int32](1),
		// 			EventType: to.Ptr("eventType1"),
		// 			StartTimestampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-09-30T21:51:39.26Z"); return t}()),
		// 		},
		// 		{
		// 			Description: to.Ptr("eventDescription2"),
		// 			DeviceID: to.Ptr("eventIdentifier2"),
		// 			EndTimestampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-07T17:34:12.50Z"); return t}()),
		// 			EventCategory: to.Ptr("eventCategory2"),
		// 			EventClass: to.Ptr("eventClass2"),
		// 			EventCount: to.Ptr[int32](1),
		// 			EventType: to.Ptr("eventType2"),
		// 			StartTimestampUTC: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-12-06T12:41:39.26Z"); return t}()),
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/310a0100f5b020c1900c527a6aa70d21992f078a/specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PostListDevicesByCatalog.json
func ExampleCatalogsClient_NewListDevicesPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armazuresphere.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewCatalogsClient().NewListDevicesPager("MyResourceGroup1", "MyCatalog1", &armazuresphere.CatalogsClientListDevicesOptions{Filter: nil,
		Top:         nil,
		Skip:        nil,
		Maxpagesize: nil,
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
		// page.DeviceListResult = armazuresphere.DeviceListResult{
		// 	Value: []*armazuresphere.Device{
		// 		{
		// 			Name: to.Ptr("00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
		// 			Properties: &armazuresphere.DeviceProperties{
		// 			},
		// 		},
		// 		{
		// 			Name: to.Ptr("00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"),
		// 			Properties: &armazuresphere.DeviceProperties{
		// 			},
		// 	}},
		// }
	}
}