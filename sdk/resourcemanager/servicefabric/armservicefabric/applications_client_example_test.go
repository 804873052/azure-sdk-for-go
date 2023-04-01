//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armservicefabric_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicefabric/armservicefabric"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationGetOperation_example.json
func ExampleApplicationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationsClient().Get(ctx, "resRg", "myCluster", "myApp", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplicationResource = armservicefabric.ApplicationResource{
	// 	Name: to.Ptr("myCluster"),
	// 	Type: to.Ptr("applications"),
	// 	Etag: to.Ptr("W/\"636462502180261859\""),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApp"),
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Properties: &armservicefabric.ApplicationResourceProperties{
	// 		MaximumNodes: to.Ptr[int64](3),
	// 		Metrics: []*armservicefabric.ApplicationMetricDescription{
	// 			{
	// 				Name: to.Ptr("metric1"),
	// 				MaximumCapacity: to.Ptr[int64](3),
	// 				ReservationCapacity: to.Ptr[int64](1),
	// 				TotalApplicationCapacity: to.Ptr[int64](5),
	// 		}},
	// 		MinimumNodes: to.Ptr[int64](1),
	// 		Parameters: map[string]*string{
	// 			"param1": to.Ptr("value1"),
	// 		},
	// 		RemoveApplicationCapacity: to.Ptr(false),
	// 		TypeVersion: to.Ptr("1.0"),
	// 		UpgradePolicy: &armservicefabric.ApplicationUpgradePolicy{
	// 			ApplicationHealthPolicy: &armservicefabric.ArmApplicationHealthPolicy{
	// 				ConsiderWarningAsError: to.Ptr(true),
	// 				DefaultServiceTypeHealthPolicy: &armservicefabric.ArmServiceTypeHealthPolicy{
	// 					MaxPercentUnhealthyPartitionsPerService: to.Ptr[int32](0),
	// 					MaxPercentUnhealthyReplicasPerPartition: to.Ptr[int32](0),
	// 					MaxPercentUnhealthyServices: to.Ptr[int32](0),
	// 				},
	// 				MaxPercentUnhealthyDeployedApplications: to.Ptr[int32](0),
	// 			},
	// 			ForceRestart: to.Ptr(false),
	// 			RollingUpgradeMonitoringPolicy: &armservicefabric.ArmRollingUpgradeMonitoringPolicy{
	// 				FailureAction: to.Ptr(armservicefabric.ArmUpgradeFailureActionRollback),
	// 				HealthCheckRetryTimeout: to.Ptr("00:10:00"),
	// 				HealthCheckStableDuration: to.Ptr("00:05:00"),
	// 				HealthCheckWaitDuration: to.Ptr("00:02:00"),
	// 				UpgradeDomainTimeout: to.Ptr("1.06:00:00"),
	// 				UpgradeTimeout: to.Ptr("01:00:00"),
	// 			},
	// 			UpgradeMode: to.Ptr(armservicefabric.RollingUpgradeModeMonitored),
	// 			UpgradeReplicaSetCheckTimeout: to.Ptr("01:00:00"),
	// 		},
	// 		ProvisioningState: to.Ptr("Succeeded"),
	// 		TypeName: to.Ptr("myAppType"),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationPutOperation_example_max.json
func ExampleApplicationsClient_BeginCreateOrUpdate_putAnApplicationWithMaximumParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationsClient().BeginCreateOrUpdate(ctx, "resRg", "myCluster", "myApp", armservicefabric.ApplicationResource{
		Tags: map[string]*string{},
		Properties: &armservicefabric.ApplicationResourceProperties{
			MaximumNodes: to.Ptr[int64](3),
			Metrics: []*armservicefabric.ApplicationMetricDescription{
				{
					Name:                     to.Ptr("metric1"),
					MaximumCapacity:          to.Ptr[int64](3),
					ReservationCapacity:      to.Ptr[int64](1),
					TotalApplicationCapacity: to.Ptr[int64](5),
				}},
			MinimumNodes: to.Ptr[int64](1),
			Parameters: map[string]*string{
				"param1": to.Ptr("value1"),
			},
			RemoveApplicationCapacity: to.Ptr(false),
			TypeVersion:               to.Ptr("1.0"),
			UpgradePolicy: &armservicefabric.ApplicationUpgradePolicy{
				ApplicationHealthPolicy: &armservicefabric.ArmApplicationHealthPolicy{
					ConsiderWarningAsError: to.Ptr(true),
					DefaultServiceTypeHealthPolicy: &armservicefabric.ArmServiceTypeHealthPolicy{
						MaxPercentUnhealthyPartitionsPerService: to.Ptr[int32](0),
						MaxPercentUnhealthyReplicasPerPartition: to.Ptr[int32](0),
						MaxPercentUnhealthyServices:             to.Ptr[int32](0),
					},
					MaxPercentUnhealthyDeployedApplications: to.Ptr[int32](0),
				},
				ForceRestart: to.Ptr(false),
				RollingUpgradeMonitoringPolicy: &armservicefabric.ArmRollingUpgradeMonitoringPolicy{
					FailureAction:             to.Ptr(armservicefabric.ArmUpgradeFailureActionRollback),
					HealthCheckRetryTimeout:   to.Ptr("00:10:00"),
					HealthCheckStableDuration: to.Ptr("00:05:00"),
					HealthCheckWaitDuration:   to.Ptr("00:02:00"),
					UpgradeDomainTimeout:      to.Ptr("1.06:00:00"),
					UpgradeTimeout:            to.Ptr("01:00:00"),
				},
				UpgradeMode:                   to.Ptr(armservicefabric.RollingUpgradeModeMonitored),
				UpgradeReplicaSetCheckTimeout: to.Ptr("01:00:00"),
			},
			TypeName: to.Ptr("myAppType"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationPutOperation_example_min.json
func ExampleApplicationsClient_BeginCreateOrUpdate_putAnApplicationWithMinimumParameters() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationsClient().BeginCreateOrUpdate(ctx, "resRg", "myCluster", "myApp", armservicefabric.ApplicationResource{
		Location: to.Ptr("eastus"),
		Tags:     map[string]*string{},
		Properties: &armservicefabric.ApplicationResourceProperties{
			RemoveApplicationCapacity: to.Ptr(false),
			TypeVersion:               to.Ptr("1.0"),
			TypeName:                  to.Ptr("myAppType"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationPutOperation_recreate_example.json
func ExampleApplicationsClient_BeginCreateOrUpdate_putAnApplicationWithRecreateOption() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationsClient().BeginCreateOrUpdate(ctx, "resRg", "myCluster", "myApp", armservicefabric.ApplicationResource{
		Tags: map[string]*string{},
		Properties: &armservicefabric.ApplicationResourceProperties{
			Parameters: map[string]*string{
				"param1": to.Ptr("value1"),
			},
			TypeVersion: to.Ptr("1.0"),
			UpgradePolicy: &armservicefabric.ApplicationUpgradePolicy{
				RecreateApplication: to.Ptr(true),
			},
			TypeName: to.Ptr("myAppType"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationPatchOperation_example.json
func ExampleApplicationsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationsClient().BeginUpdate(ctx, "resRg", "myCluster", "myApp", armservicefabric.ApplicationResourceUpdate{
		Location: to.Ptr("eastus"),
		Tags:     map[string]*string{},
		Properties: &armservicefabric.ApplicationResourceUpdateProperties{
			Metrics: []*armservicefabric.ApplicationMetricDescription{
				{
					Name:                     to.Ptr("metric1"),
					MaximumCapacity:          to.Ptr[int64](3),
					ReservationCapacity:      to.Ptr[int64](1),
					TotalApplicationCapacity: to.Ptr[int64](5),
				}},
			RemoveApplicationCapacity: to.Ptr(false),
			TypeVersion:               to.Ptr("1.0"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationDeleteOperation_example.json
func ExampleApplicationsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewApplicationsClient().BeginDelete(ctx, "resRg", "myCluster", "myApp", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/servicefabric/resource-manager/Microsoft.ServiceFabric/stable/2021-06-01/examples/ApplicationListOperation_example.json
func ExampleApplicationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armservicefabric.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewApplicationsClient().List(ctx, "resRg", "myCluster", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.ApplicationResourceList = armservicefabric.ApplicationResourceList{
	// 	Value: []*armservicefabric.ApplicationResource{
	// 		{
	// 			Name: to.Ptr("myCluster"),
	// 			Type: to.Ptr("applications"),
	// 			Etag: to.Ptr("W/\"636462502180261858\""),
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/resRg/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApp"),
	// 			Location: to.Ptr("eastus"),
	// 			Tags: map[string]*string{
	// 			},
	// 			Properties: &armservicefabric.ApplicationResourceProperties{
	// 				Metrics: []*armservicefabric.ApplicationMetricDescription{
	// 					{
	// 						Name: to.Ptr("metric1"),
	// 						MaximumCapacity: to.Ptr[int64](3),
	// 						ReservationCapacity: to.Ptr[int64](1),
	// 						TotalApplicationCapacity: to.Ptr[int64](5),
	// 				}},
	// 				RemoveApplicationCapacity: to.Ptr(false),
	// 				TypeVersion: to.Ptr("1.0"),
	// 				ProvisioningState: to.Ptr("Succeeded"),
	// 				TypeName: to.Ptr("myAppType"),
	// 			},
	// 	}},
	// }
}