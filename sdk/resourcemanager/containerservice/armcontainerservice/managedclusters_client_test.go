//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armcontainerservice_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestManagedClustersClient_BeginCreateOrUpdate(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	clientId := lookupEnvVar("AZURE_CLIENT_ID")
	clientSecret := lookupEnvVar("AZURE_CLIENT_SECRET")
	location := "eastus"
	ctx := context.Background()

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	rgName := *rg.Name
	defer clean()

	// create container service
	managedClustersClient := armcontainerservice.NewManagedClustersClient(subscriptionID, cred, opt)
	createManagedCluster(t, ctx, managedClustersClient, rgName, location, clientId, clientSecret)
}

func TestManagedClustersClient_BeginUpdateTags(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	clientId := lookupEnvVar("AZURE_CLIENT_ID")
	clientSecret := lookupEnvVar("AZURE_CLIENT_SECRET")
	location := "eastus"
	ctx := context.Background()

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	rgName := *rg.Name
	defer clean()

	// create container service
	managedClustersClient := armcontainerservice.NewManagedClustersClient(subscriptionID, cred, opt)
	managedCluster := createManagedCluster(t, ctx, managedClustersClient, rgName, location, clientId, clientSecret)
	managedClusterName := *managedCluster.Name

	// update tags
	updatePollerResp, err := managedClustersClient.BeginUpdateTags(
		ctx,
		rgName,
		managedClusterName,
		armcontainerservice.TagsObject{
			Tags: map[string]*string{
				"test": to.StringPtr("recording"),
			},
		},
		nil,
	)
	require.NoError(t, err)
	updateResp, err := updatePollerResp.PollUntilDone(ctx, 10*time.Second)
	require.NoError(t, err)
	require.Equal(t, "recording", updateResp.Tags["test"])
}

func TestManagedClustersClient_Get(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	clientId := lookupEnvVar("AZURE_CLIENT_ID")
	clientSecret := lookupEnvVar("AZURE_CLIENT_SECRET")
	location := "eastus"
	ctx := context.Background()

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	rgName := *rg.Name
	defer clean()

	// create cluster service
	managedClustersClient := armcontainerservice.NewManagedClustersClient(subscriptionID, cred, opt)
	managedCluster := createManagedCluster(t, ctx, managedClustersClient, rgName, location, clientId, clientSecret)
	managedClusterName := *managedCluster.Name

	// get cluster service
	getResp, err := managedClustersClient.Get(ctx, rgName, managedClusterName, nil)
	require.NoError(t, err)
	require.Equal(t, managedClusterName, *getResp.Name)
}

func TestManagedClustersClient_GetUpgradeProfile(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	clientId := lookupEnvVar("AZURE_CLIENT_ID")
	clientSecret := lookupEnvVar("AZURE_CLIENT_SECRET")
	location := "eastus"
	ctx := context.Background()

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	rgName := *rg.Name
	defer clean()

	// create cluster service
	managedClustersClient := armcontainerservice.NewManagedClustersClient(subscriptionID, cred, opt)
	managedCluster := createManagedCluster(t, ctx, managedClustersClient, rgName, location, clientId, clientSecret)
	managedClusterName := *managedCluster.Name

	// get upgrade profile
	getResp, err := managedClustersClient.GetUpgradeProfile(ctx, rgName, managedClusterName, nil)
	require.NoError(t, err)
	require.Equal(t, managedClusterName, *getResp.Name)
}

func TestManagedClustersClient_List(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	clientId := lookupEnvVar("AZURE_CLIENT_ID")
	clientSecret := lookupEnvVar("AZURE_CLIENT_SECRET")
	location := "eastus"
	ctx := context.Background()

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	rgName := *rg.Name
	defer clean()

	// create cluster service
	managedClustersClient := armcontainerservice.NewManagedClustersClient(subscriptionID, cred, opt)
	createManagedCluster(t, ctx, managedClustersClient, rgName, location, clientId, clientSecret)

	// list cluster service
	pager := managedClustersClient.List(nil)
	require.NoError(t, pager.Err())
	require.Equal(t, 1, len(pager.PageResponse().Value))
}

func TestManagedClustersClient_BeginDelete(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	clientId := lookupEnvVar("AZURE_CLIENT_ID")
	clientSecret := lookupEnvVar("AZURE_CLIENT_SECRET")
	location := "eastus"
	ctx := context.Background()

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	rgName := *rg.Name
	defer clean()

	// create cluster service
	managedClustersClient := armcontainerservice.NewManagedClustersClient(subscriptionID, cred, opt)
	managedCluster := createManagedCluster(t, ctx, managedClustersClient, rgName, location, clientId, clientSecret)
	managedClusterName := *managedCluster.Name

	// delete cluster service
	delPollerResp, err := managedClustersClient.BeginDelete(ctx, rgName, managedClusterName, nil)
	require.NoError(t, err)
	delResp, err := delPollerResp.PollUntilDone(ctx, 10*time.Second)
	require.Equal(t, 200, delResp.RawResponse.StatusCode)
}

func TestManagedClustersClient_GetOSOptions(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	location := "eastus"
	ctx := context.Background()

	// create cluster service
	managedClustersClient := armcontainerservice.NewManagedClustersClient(subscriptionID, cred, opt)

	// get os option
	osOptions, err := managedClustersClient.GetOSOptions(ctx, location, nil)
	require.NoError(t, err)
	require.Equal(t, "default", *osOptions.Name)
}

func createManagedCluster(t *testing.T, ctx context.Context, managedClustersClient *armcontainerservice.ManagedClustersClient, rgName, location, clientId, clientSecret string) *armcontainerservice.ManagedCluster {
	managedClusterName, _ := createRandomName(t, "cluster")
	mcPollerResp, err := managedClustersClient.BeginCreateOrUpdate(
		ctx,
		rgName,
		managedClusterName,
		armcontainerservice.ManagedCluster{
			Location: to.StringPtr(location),
			Properties: &armcontainerservice.ManagedClusterProperties{
				DNSPrefix: to.StringPtr("aksgosdk"),
				AgentPoolProfiles: []*armcontainerservice.ManagedClusterAgentPoolProfile{
					{
						Name:              to.StringPtr("askagent"),
						Count:             to.Int32Ptr(1),
						VMSize:            to.StringPtr("Standard_DS2_v2"),
						MaxPods:           to.Int32Ptr(110),
						MinCount:          to.Int32Ptr(1),
						MaxCount:          to.Int32Ptr(100),
						OSType:            armcontainerservice.OSTypeLinux.ToPtr(),
						Type:              armcontainerservice.AgentPoolTypeVirtualMachineScaleSets.ToPtr(),
						EnableAutoScaling: to.BoolPtr(true),
						Mode:              armcontainerservice.AgentPoolModeSystem.ToPtr(),
					},
				},
				ServicePrincipalProfile: &armcontainerservice.ManagedClusterServicePrincipalProfile{
					ClientID: to.StringPtr(clientId),
					Secret:   to.StringPtr(clientSecret),
				},
			},
		},
		nil,
	)
	require.NoError(t, err)
	mcResp, err := mcPollerResp.PollUntilDone(ctx, 10*time.Second)
	require.NoError(t, err)
	require.Equal(t, managedClusterName, *mcResp.Name)
	return &mcResp.ManagedCluster
}
