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

func TestTopicsClient_CreateOrUpdate(t *testing.T) {
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

	// create topic
	topicsClient := armservicebus.NewTopicsClient(subscriptionID, cred, opt)
	createTopic(t, ctx, topicsClient, rgName, namespaceName)
}

func TestTopicsClient_Delete(t *testing.T) {
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

	// create topic
	topicsClient := armservicebus.NewTopicsClient(subscriptionID, cred, opt)
	topic := createTopic(t, ctx, topicsClient, rgName, namespaceName)
	topicName := *topic.Name

	// delete topic
	delResp, err := topicsClient.Delete(ctx, rgName, namespaceName, topicName, nil)
	require.NoError(t, err)
	require.Equal(t, 200, delResp.RawResponse.StatusCode)
}

func TestTopicsClient_Get(t *testing.T) {
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

	// create topic
	topicsClient := armservicebus.NewTopicsClient(subscriptionID, cred, opt)
	topic := createTopic(t, ctx, topicsClient, rgName, namespaceName)
	topicName := *topic.Name

	// get topic
	getResp, err := topicsClient.Get(ctx, rgName, namespaceName, topicName, nil)
	require.NoError(t, err)
	require.Equal(t, topicName, *getResp.Name)
}

func TestTopicsClient_ListByNamespace(t *testing.T) {
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

	// create topic
	topicsClient := armservicebus.NewTopicsClient(subscriptionID, cred, opt)
	createTopic(t, ctx, topicsClient, rgName, namespaceName)

	// list topic
	list := topicsClient.ListByNamespace(rgName, namespaceName, nil)
	require.NoError(t, list.Err())
	require.Equal(t, 1, len(list.PageResponse().Value))
}

func createTopic(t *testing.T, ctx context.Context, topicsClient *armservicebus.TopicsClient, rgName, namespaceName string) *armservicebus.SBTopic {
	topicName, _ := createRandomName(t, "topic")
	topicResp, err := topicsClient.CreateOrUpdate(
		ctx,
		rgName,
		namespaceName,
		topicName,
		armservicebus.SBTopic{
			Properties: &armservicebus.SBTopicProperties{
				EnableExpress: to.BoolPtr(true),
			},
		},
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, topicName, *topicResp.Name)
	return &topicResp.SBTopic
}
