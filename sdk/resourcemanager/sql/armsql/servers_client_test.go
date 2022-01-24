//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armsql_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/sql/armsql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestServersClient_BeginCreateOrUpdate(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	ctx := context.Background()
	location := "eastus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// create sql server
	serversClient := armsql.NewServersClient(subscriptionID, cred, opt)
	serverName, _ := createRandomName(t, "sql")
	serverPollerResp, err := serversClient.BeginCreateOrUpdate(ctx, rgName, serverName, armsql.Server{
		Location: to.StringPtr(location),
		Properties: &armsql.ServerProperties{
			AdministratorLogin:         to.StringPtr("testlogin"),
			AdministratorLoginPassword: to.StringPtr("QWE123!@#"),
		},
	}, nil)
	require.NoError(t, err)
	serverResp, err := serverPollerResp.PollUntilDone(ctx, 30*time.Second)
	require.NoError(t, err)
	require.Equal(t, serverName, *serverResp.Name)
}

func TestServersClient_CheckNameAvailability(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	ctx := context.Background()

	// create sql server
	serversClient := armsql.NewServersClient(subscriptionID, cred, opt)
	serverName, _ := createRandomName(t, "sql")
	checkName, err := serversClient.CheckNameAvailability(ctx, armsql.CheckNameAvailabilityRequest{
		Name: to.StringPtr(serverName),
		Type: to.StringPtr("Microsoft.Sql/servers"),
	}, nil)
	require.NoError(t, err)
	require.True(t, *checkName.Available)
}

func TestServersClient_BeginUpdate(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	ctx := context.Background()
	location := "eastus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// create sql server
	serversClient := armsql.NewServersClient(subscriptionID, cred, opt)
	server := createSqlServer(t, ctx, serversClient, rgName, location)
	serverName := *server.Name

	// update sql server
	updatePollerResp, err := serversClient.BeginUpdate(ctx, rgName, serverName, armsql.ServerUpdate{}, nil)
	require.NoError(t, err)
	updateResp, err := updatePollerResp.PollUntilDone(ctx, 30*time.Second)
	require.NoError(t, err)
	require.Equal(t, serverName, *updateResp.Name)
}

func TestServersClient_BeginDelete(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	ctx := context.Background()
	location := "eastus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// create sql server
	serversClient := armsql.NewServersClient(subscriptionID, cred, opt)
	server := createSqlServer(t, ctx, serversClient, rgName, location)
	serverName := *server.Name

	// delete sql server
	delPollerResp, err := serversClient.BeginDelete(ctx, rgName, serverName, nil)
	require.NoError(t, err)
	delResp, err := delPollerResp.PollUntilDone(ctx, 30*time.Second)
	require.NoError(t, err)
	require.Equal(t, 200, delResp.RawResponse.StatusCode)
}

func TestServersClient_Get(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	ctx := context.Background()
	location := "eastus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// create sql server
	serversClient := armsql.NewServersClient(subscriptionID, cred, opt)
	server := createSqlServer(t, ctx, serversClient, rgName, location)
	serverName := *server.Name

	// get sql server
	getResp, err := serversClient.Get(ctx, rgName, serverName, nil)
	require.NoError(t, err)
	require.Equal(t, serverName, *getResp.Name)
}

func TestServersClient_List(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	ctx := context.Background()
	location := "eastus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// create sql server
	serversClient := armsql.NewServersClient(subscriptionID, cred, opt)
	createSqlServer(t, ctx, serversClient, rgName, location)

	// list sql server
	pager := serversClient.List(nil)
	require.NoError(t, pager.Err())
	require.Equal(t, 1, len(pager.PageResponse().Value))
}

func TestServersClient_ListByResourceGroup(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	ctx := context.Background()
	location := "eastus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// create sql server
	serversClient := armsql.NewServersClient(subscriptionID, cred, opt)
	createSqlServer(t, ctx, serversClient, rgName, location)

	// list sql server
	pager := serversClient.ListByResourceGroup(rgName, nil)
	require.NoError(t, pager.Err())
	require.Equal(t, 1, len(pager.PageResponse().Value))
}

func createSqlServer(t *testing.T, ctx context.Context, serversClient *armsql.ServersClient, rgName, location string) *armsql.Server {
	serverName, _ := createRandomName(t, "sql")
	serverPollerResp, err := serversClient.BeginCreateOrUpdate(ctx, rgName, serverName, armsql.Server{
		Location: to.StringPtr(location),
		Properties: &armsql.ServerProperties{
			AdministratorLogin:         to.StringPtr("testlogin"),
			AdministratorLoginPassword: to.StringPtr("QWE123!@#"),
		},
	}, nil)
	require.NoError(t, err)
	serverResp, err := serverPollerResp.PollUntilDone(ctx, 30*time.Second)
	require.NoError(t, err)
	require.Equal(t, serverName, *serverResp.Name)
	return &serverResp.Server
}
