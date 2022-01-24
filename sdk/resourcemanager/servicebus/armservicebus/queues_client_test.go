//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armservicebus_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueuesClient_CreateOrUpdate(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	ctx := context.Background()
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	location := "westus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// create namespaces
	namespacesClient := armservicebus.NewNamespacesClient(subscriptionID, cred, opt)
	namespaces := createNamespaces(t, ctx, namespacesClient, rgName, location)
	namespaceName := *namespaces.Name

	// create queue
	queuesClient := armservicebus.NewQueuesClient(subscriptionID, cred, opt)
	createQueue(t, ctx, queuesClient, rgName, namespaceName)
}

func TestQueuesClient_Delete(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	ctx := context.Background()
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	location := "westus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// create namespaces
	namespacesClient := armservicebus.NewNamespacesClient(subscriptionID, cred, opt)
	namespaces := createNamespaces(t, ctx, namespacesClient, rgName, location)
	namespaceName := *namespaces.Name

	// create queue
	queuesClient := armservicebus.NewQueuesClient(subscriptionID, cred, opt)
	queue := createQueue(t, ctx, queuesClient, rgName, namespaceName)
	queueName := *queue.Name

	// delete queue
	delResp, err := queuesClient.Delete(ctx, rgName, namespaceName, queueName, nil)
	require.NoError(t, err)
	require.Equal(t, 200, delResp.RawResponse.StatusCode)
}

func TestQueuesClient_Get(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	ctx := context.Background()
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	location := "westus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// create namespaces
	namespacesClient := armservicebus.NewNamespacesClient(subscriptionID, cred, opt)
	namespaces := createNamespaces(t, ctx, namespacesClient, rgName, location)
	namespaceName := *namespaces.Name

	// create queue
	queuesClient := armservicebus.NewQueuesClient(subscriptionID, cred, opt)
	queue := createQueue(t, ctx, queuesClient, rgName, namespaceName)
	queueName := *queue.Name

	// get queue
	getResp, err := queuesClient.Get(ctx, rgName, namespaceName, queueName, nil)
	require.NoError(t, err)
	require.Equal(t, queueName, *getResp.Name)
}

func TestQueuesClient_ListByNamespace(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	ctx := context.Background()
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	location := "westus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// create namespaces
	namespacesClient := armservicebus.NewNamespacesClient(subscriptionID, cred, opt)
	namespaces := createNamespaces(t, ctx, namespacesClient, rgName, location)
	namespaceName := *namespaces.Name

	// create queue
	queuesClient := armservicebus.NewQueuesClient(subscriptionID, cred, opt)
	createQueue(t, ctx, queuesClient, rgName, namespaceName)

	// list queue
	list := queuesClient.ListByNamespace(rgName, namespaceName, nil)
	require.NoError(t, list.Err())
	require.Equal(t, 1, len(list.PageResponse().Value))
}

func createQueue(t *testing.T, ctx context.Context, queuesClient *armservicebus.QueuesClient, rgName, namespaceName string) *armservicebus.SBQueue {
	queueName, _ := createRandomName(t, "queue")
	queueResp, err := queuesClient.CreateOrUpdate(
		ctx,
		rgName,
		namespaceName,
		queueName,
		armservicebus.SBQueue{
			Properties: &armservicebus.SBQueueProperties{
				EnablePartitioning: to.BoolPtr(true),
			},
		},
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, queueName, *queueResp.Name)
	return &queueResp.SBQueue
}
