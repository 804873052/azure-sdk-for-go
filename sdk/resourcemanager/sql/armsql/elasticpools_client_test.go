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

func TestElasticPoolsClient_BeginCreateOrUpdate(t *testing.T) {
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

	// create elastic pool
	elasticPoolsClient := armsql.NewElasticPoolsClient(subscriptionID, cred, opt)
	elasticPoolName, _ := createRandomName(t, "elastic")
	epPollerResp, err := elasticPoolsClient.BeginCreateOrUpdate(ctx, rgName, serverName, elasticPoolName, armsql.ElasticPool{Location: to.StringPtr(location)}, nil)
	require.NoError(t, err)
	epResp, err := epPollerResp.PollUntilDone(ctx, 30*time.Second)
	require.NoError(t, err)
	require.Equal(t, elasticPoolName, *epResp.Name)
}

func TestElasticPoolsClient_BeginUpdate(t *testing.T) {
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

	// create elastic pool
	elasticPoolsClient := armsql.NewElasticPoolsClient(subscriptionID, cred, opt)
	elasticPoolName, _ := createRandomName(t, "elastic")
	epPollerResp, err := elasticPoolsClient.BeginCreateOrUpdate(ctx, rgName, serverName, elasticPoolName, armsql.ElasticPool{Location: to.StringPtr(location)}, nil)
	require.NoError(t, err)
	epResp, err := epPollerResp.PollUntilDone(ctx, 30*time.Second)
	require.NoError(t, err)
	require.Equal(t, elasticPoolName, *epResp.Name)

	// update elastic pool
	updatePollerResp, err := elasticPoolsClient.BeginUpdate(
		ctx,
		rgName,
		serverName,
		elasticPoolName,
		armsql.ElasticPoolUpdate{Tags: map[string]*string{
			"test": to.StringPtr("recording"),
		}},
		nil,
	)
	require.NoError(t, err)
	updateResp, err := updatePollerResp.PollUntilDone(ctx, 30*time.Second)
	require.NoError(t, err)
	require.Equal(t, "recording", *updateResp.Tags["test"])
}

func TestElasticPoolsClient_BeginDelete(t *testing.T) {
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

	// create elastic pool
	elasticPoolsClient := armsql.NewElasticPoolsClient(subscriptionID, cred, opt)
	elasticPoolName, _ := createRandomName(t, "elastic")
	epPollerResp, err := elasticPoolsClient.BeginCreateOrUpdate(ctx, rgName, serverName, elasticPoolName, armsql.ElasticPool{Location: to.StringPtr(location)}, nil)
	require.NoError(t, err)
	epResp, err := epPollerResp.PollUntilDone(ctx, 30*time.Second)
	require.NoError(t, err)
	require.Equal(t, elasticPoolName, *epResp.Name)

	// delete elastic pool
	deletePollerResp, err := elasticPoolsClient.BeginDelete(ctx, rgName, serverName, elasticPoolName, nil)
	require.NoError(t, err)
	deleteResp, err := deletePollerResp.PollUntilDone(ctx, 30*time.Second)
	require.NoError(t, err)
	require.Equal(t, 200, deleteResp.RawResponse.StatusCode)
}

func TestElasticPoolsClient_Get(t *testing.T) {
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

	// create elastic pool
	elasticPoolsClient := armsql.NewElasticPoolsClient(subscriptionID, cred, opt)
	elasticPoolName, _ := createRandomName(t, "elastic")
	epPollerResp, err := elasticPoolsClient.BeginCreateOrUpdate(ctx, rgName, serverName, elasticPoolName, armsql.ElasticPool{Location: to.StringPtr(location)}, nil)
	require.NoError(t, err)
	epResp, err := epPollerResp.PollUntilDone(ctx, 30*time.Second)
	require.NoError(t, err)
	require.Equal(t, elasticPoolName, *epResp.Name)

	// get elastic pool
	getResp, err := elasticPoolsClient.Get(ctx, rgName, serverName, elasticPoolName, nil)
	require.NoError(t, err)
	require.Equal(t, elasticPoolName, *getResp.Name)
}

func TestElasticPoolsClient_ListByServer(t *testing.T) {
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

	// create elastic pool
	elasticPoolsClient := armsql.NewElasticPoolsClient(subscriptionID, cred, opt)
	elasticPoolName, _ := createRandomName(t, "elastic")
	epPollerResp, err := elasticPoolsClient.BeginCreateOrUpdate(ctx, rgName, serverName, elasticPoolName, armsql.ElasticPool{Location: to.StringPtr(location)}, nil)
	require.NoError(t, err)
	epResp, err := epPollerResp.PollUntilDone(ctx, 30*time.Second)
	require.NoError(t, err)
	require.Equal(t, elasticPoolName, *epResp.Name)

	// list elastic pool
	pager := elasticPoolsClient.ListByServer(rgName, serverName, nil)
	require.NoError(t, pager.Err())
	require.Equal(t, 1, len(pager.PageResponse().Value))
}
