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

func TestDatabasesClient_BeginCreateOrUpdate(t *testing.T) {
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

	// create sql database
	databaseClient := armsql.NewDatabasesClient(subscriptionID, cred, opt)
	createDatabase(t, ctx, databaseClient, rgName, serverName, location)
}

func TestDatabasesClient_BeginUpdate(t *testing.T) {
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

	// create database
	databaseClient := armsql.NewDatabasesClient(subscriptionID, cred, opt)
	database := createDatabase(t, ctx, databaseClient, rgName, serverName, location)

	// update database
	updatePollerResp, err := databaseClient.BeginUpdate(ctx, rgName, serverName, *database.Name, armsql.DatabaseUpdate{
		Tags: map[string]*string{
			"test": to.StringPtr("recording"),
		},
	}, nil)
	require.NoError(t, err)
	updateResp, err := updatePollerResp.PollUntilDone(ctx, 30*time.Second)
	require.NoError(t, err)
	require.Equal(t, "recording", *updateResp.Tags["test"])
}

func TestDatabasesClient_Get(t *testing.T) {
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

	// create database
	databaseClient := armsql.NewDatabasesClient(subscriptionID, cred, opt)
	database := createDatabase(t, ctx, databaseClient, rgName, serverName, location)

	// get database
	getResp, err := databaseClient.Get(ctx, rgName, serverName, *database.Name, nil)
	require.NoError(t, err)
	require.Equal(t, *database.Name, *getResp.Name)
}

func TestDatabasesClient_BeginDelete(t *testing.T) {
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

	// create database
	databaseClient := armsql.NewDatabasesClient(subscriptionID, cred, opt)
	database := createDatabase(t, ctx, databaseClient, rgName, serverName, location)

	// delete database
	deletePollerResp, err := databaseClient.BeginDelete(ctx, rgName, serverName, *database.Name, nil)
	require.NoError(t, err)
	deleteResp, err := deletePollerResp.PollUntilDone(ctx, 30*time.Second)
	require.NoError(t, err)
	require.Equal(t, 200, deleteResp.RawResponse.StatusCode)
}

func TestDatabasesClient_ListByServer(t *testing.T) {
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

	// create database
	databaseClient := armsql.NewDatabasesClient(subscriptionID, cred, opt)
	createDatabase(t, ctx, databaseClient, rgName, serverName, location)

	// update database
	pager := databaseClient.ListByServer(rgName, serverName, nil)
	require.NoError(t, pager.Err())
	require.Equal(t, 1, len(pager.PageResponse().Value))
}

func createDatabase(t *testing.T, ctx context.Context, databaseClient *armsql.DatabasesClient, rgName, serverName, location string) *armsql.Database {
	databaseName, _ := createRandomName(t, "database")
	dbPollerResp, err := databaseClient.BeginCreateOrUpdate(
		ctx,
		rgName,
		serverName,
		databaseName,
		armsql.Database{
			Location: to.StringPtr(location),
			Properties: &armsql.DatabaseProperties{
				ReadScale: armsql.DatabaseReadScaleDisabled.ToPtr(),
			},
		},
		nil,
	)
	require.NoError(t, err)
	dbResp, err := dbPollerResp.PollUntilDone(ctx, 30*time.Second)
	require.NoError(t, err)
	require.Equal(t, databaseName, *dbResp.Name)
	return &dbResp.Database
}
