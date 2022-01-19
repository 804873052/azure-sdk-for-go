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

func TestAgentPoolsClient_BeginCreateOrUpdate(t *testing.T) {
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
	managedCLuster := createManagedCluster(t, ctx, managedClustersClient, rgName, location, clientId, clientSecret)
	clusterName := *managedCLuster.Name

	// create agent pool
	agentPoolsClient := armcontainerservice.NewAgentPoolsClient(subscriptionID, cred, opt)
	createAgentPool(t, ctx, agentPoolsClient, rgName, clusterName)
}

func TestAgentPoolsClient_BeginDelete(t *testing.T) {
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
	managedCLuster := createManagedCluster(t, ctx, managedClustersClient, rgName, location, clientId, clientSecret)
	clusterName := *managedCLuster.Name

	// create agent pool
	agentPoolsClient := armcontainerservice.NewAgentPoolsClient(subscriptionID, cred, opt)
	agentPool := createAgentPool(t, ctx, agentPoolsClient, rgName, clusterName)
	agentPoolName := *agentPool.Name

	// delete agent pool
	delPollerResp, err := agentPoolsClient.BeginDelete(ctx, rgName, clusterName, agentPoolName, nil)
	require.NoError(t, err)
	delResp, err := delPollerResp.PollUntilDone(ctx, 10*time.Second)
	require.NoError(t, err)
	require.Equal(t, 200, delResp.RawResponse.StatusCode)
}

func TestAgentPoolsClient_Get(t *testing.T) {
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
	managedCLuster := createManagedCluster(t, ctx, managedClustersClient, rgName, location, clientId, clientSecret)
	clusterName := *managedCLuster.Name

	// create agent pool
	agentPoolsClient := armcontainerservice.NewAgentPoolsClient(subscriptionID, cred, opt)
	agentPool := createAgentPool(t, ctx, agentPoolsClient, rgName, clusterName)
	agentPoolName := *agentPool.Name

	// get agent pool
	getResp, err := agentPoolsClient.Get(ctx, rgName, clusterName, agentPoolName, nil)
	require.NoError(t, err)
	require.Equal(t, agentPoolName, *getResp.Name)
}

func TestAgentPoolsClient_GetAvailableAgentPoolVersions(t *testing.T) {
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
	managedCLuster := createManagedCluster(t, ctx, managedClustersClient, rgName, location, clientId, clientSecret)
	clusterName := *managedCLuster.Name

	// create agent pool
	agentPoolsClient := armcontainerservice.NewAgentPoolsClient(subscriptionID, cred, opt)
	agentPool := createAgentPool(t, ctx, agentPoolsClient, rgName, clusterName)
	agentPoolName := *agentPool.Name

	// get available agent pool
	getResp, err := agentPoolsClient.GetAvailableAgentPoolVersions(ctx, rgName, clusterName, nil)
	require.NoError(t, err)
	require.Equal(t, agentPoolName, *getResp.Name)
}

func TestAgentPoolsClient_GetUpgradeProfile(t *testing.T) {
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
	managedCLuster := createManagedCluster(t, ctx, managedClustersClient, rgName, location, clientId, clientSecret)
	clusterName := *managedCLuster.Name

	// create agent pool
	agentPoolsClient := armcontainerservice.NewAgentPoolsClient(subscriptionID, cred, opt)
	agentPool := createAgentPool(t, ctx, agentPoolsClient, rgName, clusterName)
	agentPoolName := *agentPool.Name

	// get upgrade Profile
	getResp, err := agentPoolsClient.GetUpgradeProfile(ctx, rgName, clusterName, agentPoolName, nil)
	require.NoError(t, err)
	require.Equal(t, agentPoolName, *getResp.Name)
}

func TestAgentPoolsClient_List(t *testing.T) {
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
	managedCLuster := createManagedCluster(t, ctx, managedClustersClient, rgName, location, clientId, clientSecret)
	clusterName := *managedCLuster.Name

	// create agent pool
	agentPoolsClient := armcontainerservice.NewAgentPoolsClient(subscriptionID, cred, opt)
	createAgentPool(t, ctx, agentPoolsClient, rgName, clusterName)

	// list agent pool
	pager := agentPoolsClient.List(rgName, clusterName, nil)
	require.NoError(t, pager.Err())
	require.Equal(t, 1, len(pager.PageResponse().Value))
}

func createAgentPool(t *testing.T, ctx context.Context, agentPoolsClient *armcontainerservice.AgentPoolsClient, rgName, clusterName string) *armcontainerservice.AgentPool {
	agentPoolName, _ := createRandomName(t, "agent")
	agentPoolPollerResp, err := agentPoolsClient.BeginCreateOrUpdate(
		ctx,
		rgName,
		clusterName,
		agentPoolName,
		armcontainerservice.AgentPool{
			Properties: &armcontainerservice.ManagedClusterAgentPoolProfileProperties{
				OrchestratorVersion: to.StringPtr(""),
				Count:               to.Int32Ptr(3),
				VMSize:              to.StringPtr("Standard_DS2_v2"),
				OSType:              armcontainerservice.OSTypeLinux.ToPtr(),
				Mode:                armcontainerservice.AgentPoolModeSystem.ToPtr(),
				AvailabilityZones: []*string{
					to.StringPtr("1"),
					to.StringPtr("2"),
					to.StringPtr("3"),
				},
				NodeTaints: []*string{},
			},
		},
		nil,
	)
	require.NoError(t, err)
	agentPoolResp, err := agentPoolPollerResp.PollUntilDone(ctx, 10*time.Second)
	require.NoError(t, err)
	require.Equal(t, agentPoolName, *agentPoolResp.Name)
	return &agentPoolResp.AgentPool
}
