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
	"time"
)

func TestNamespacesClient_BeginCreateOrUpdate(t *testing.T) {
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
	createNamespaces(t, ctx, namespacesClient, rgName, location)
}

func TestNamespacesClient_BeginDelete(t *testing.T) {
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
	sbNamespaces := createNamespaces(t, ctx, namespacesClient, rgName, location)
	namespaceName := *sbNamespaces.Name

	// delete namespaces
	delPollerResp, err := namespacesClient.BeginDelete(ctx, rgName, namespaceName, nil)
	require.NoError(t, err)
	delResp, err := delPollerResp.PollUntilDone(ctx, 30*time.Second)
	require.NoError(t, err)
	require.Equal(t, 200, delResp.RawResponse.StatusCode)
}

func TestNamespacesClient_Update(t *testing.T) {
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
	sbNamespaces := createNamespaces(t, ctx, namespacesClient, rgName, location)
	namespaceName := *sbNamespaces.Name

	// update namespaces
	updateResp, err := namespacesClient.Update(
		ctx,
		rgName,
		namespaceName,
		armservicebus.SBNamespaceUpdateParameters{
			Tags: map[string]*string{
				"test": to.StringPtr("recording"),
			},
		},
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, "recording", *updateResp.Tags["test"])
}

func TestNamespacesClient_Get(t *testing.T) {
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
	sbNamespaces := createNamespaces(t, ctx, namespacesClient, rgName, location)
	namespaceName := *sbNamespaces.Name

	// get namespaces
	getResp, err := namespacesClient.Get(ctx, rgName, namespaceName, nil)
	require.NoError(t, err)
	require.Equal(t, namespaceName, *getResp.Name)
}

func TestNamespacesClient_List(t *testing.T) {
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
	createNamespaces(t, ctx, namespacesClient, rgName, location)

	// list namespaces
	list := namespacesClient.List(nil)
	require.NoError(t, list.Err())
	require.Equal(t, 1, len(list.PageResponse().Value))
}

func TestNamespacesClient_CheckNameAvailability(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	ctx := context.Background()
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")

	// create namespaces
	namespacesClient := armservicebus.NewNamespacesClient(subscriptionID, cred, opt)
	namespaceName, _ := createRandomName(t, "namespace")
	checkName, err := namespacesClient.CheckNameAvailability(
		ctx,
		armservicebus.CheckNameAvailability{
			Name: to.StringPtr(namespaceName),
		},
		nil,
	)
	require.NoError(t, err)
	require.True(t, *checkName.NameAvailable)
}

func createNamespaces(t *testing.T, ctx context.Context, namespacesClient *armservicebus.NamespacesClient, rgName, location string) *armservicebus.SBNamespace {
	namespaceName, _ := createRandomName(t, "namespace")
	npPollerResp, err := namespacesClient.BeginCreateOrUpdate(
		ctx,
		rgName,
		namespaceName,
		armservicebus.SBNamespace{
			Location: to.StringPtr(location),
			SKU: &armservicebus.SBSKU{
				Name: armservicebus.SKUNamePremium.ToPtr(),
				Tier: armservicebus.SKUTierPremium.ToPtr(),
			},
		},
		nil,
	)
	require.NoError(t, err)
	npResp, err := npPollerResp.PollUntilDone(ctx, 30*time.Second)
	require.NoError(t, err)
	require.Equal(t, namespaceName, *npResp.Name)
	return &npResp.SBNamespace
}
