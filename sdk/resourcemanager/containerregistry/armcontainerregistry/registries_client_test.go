//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armcontainerregistry_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerregistry/armcontainerregistry"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRegistriesClient_BeginCreate(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	ctx := context.Background()
	location := "westus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// container registry
	registriesClient := armcontainerregistry.NewRegistriesClient(subscriptionID, cred, opt)
	registryName, _ := createRandomName(t, "registry")
	registryPollerResp, err := registriesClient.BeginCreate(
		ctx,
		rgName,
		registryName,
		armcontainerregistry.Registry{
			Location: to.StringPtr(location),
			Tags: map[string]*string{
				"key": to.StringPtr("value"),
			},
			SKU: &armcontainerregistry.SKU{
				Name: armcontainerregistry.SKUNamePremium.ToPtr(),
			},
			Properties: &armcontainerregistry.RegistryProperties{
				AdminUserEnabled: to.BoolPtr(false),
			},
		},
		nil,
	)
	require.NoError(t, err)
	registryResp, err := registryPollerResp.PollUntilDone(ctx, 10*time.Second)
	require.NoError(t, err)
	require.Equal(t, registryName, *registryResp.Name)
}

func TestRegistriesClient_BeginDelete(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	ctx := context.Background()
	location := "westus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// container registry
	registriesClient := armcontainerregistry.NewRegistriesClient(subscriptionID, cred, opt)
	registryName, _ := createRandomName(t, "registry")
	registryPollerResp, err := registriesClient.BeginCreate(
		ctx,
		rgName,
		registryName,
		armcontainerregistry.Registry{
			Location: to.StringPtr(location),
			Tags: map[string]*string{
				"key": to.StringPtr("value"),
			},
			SKU: &armcontainerregistry.SKU{
				Name: armcontainerregistry.SKUNamePremium.ToPtr(),
			},
			Properties: &armcontainerregistry.RegistryProperties{
				AdminUserEnabled: to.BoolPtr(false),
			},
		},
		nil,
	)
	require.NoError(t, err)
	registryResp, err := registryPollerResp.PollUntilDone(ctx, 10*time.Second)
	require.NoError(t, err)
	require.Equal(t, registryName, *registryResp.Name)

	// delete container registry
	delPollerResp, err := registriesClient.BeginDelete(ctx, rgName, registryName, nil)
	require.NoError(t, err)
	delResp, err := delPollerResp.PollUntilDone(ctx, 10*time.Second)
	require.NoError(t, err)
	require.Equal(t, 200, delResp.RawResponse.StatusCode)
}

func TestRegistriesClient_BeginUpdate(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	ctx := context.Background()
	location := "westus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// container registry
	registriesClient := armcontainerregistry.NewRegistriesClient(subscriptionID, cred, opt)
	registryName, _ := createRandomName(t, "registry")
	registryPollerResp, err := registriesClient.BeginCreate(
		ctx,
		rgName,
		registryName,
		armcontainerregistry.Registry{
			Location: to.StringPtr(location),
			Tags: map[string]*string{
				"key": to.StringPtr("value"),
			},
			SKU: &armcontainerregistry.SKU{
				Name: armcontainerregistry.SKUNamePremium.ToPtr(),
			},
			Properties: &armcontainerregistry.RegistryProperties{
				AdminUserEnabled: to.BoolPtr(false),
			},
		},
		nil,
	)
	require.NoError(t, err)
	registryResp, err := registryPollerResp.PollUntilDone(ctx, 10*time.Second)
	require.NoError(t, err)
	require.Equal(t, registryName, *registryResp.Name)

	// update container registry
	updatePollerResp, err := registriesClient.BeginUpdate(
		ctx,
		rgName,
		registryName,
		armcontainerregistry.RegistryUpdateParameters{
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

func TestRegistriesClient_Get(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	ctx := context.Background()
	location := "westus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// container registry
	registriesClient := armcontainerregistry.NewRegistriesClient(subscriptionID, cred, opt)
	registryName, _ := createRandomName(t, "registry")
	registryPollerResp, err := registriesClient.BeginCreate(
		ctx,
		rgName,
		registryName,
		armcontainerregistry.Registry{
			Location: to.StringPtr(location),
			Tags: map[string]*string{
				"key": to.StringPtr("value"),
			},
			SKU: &armcontainerregistry.SKU{
				Name: armcontainerregistry.SKUNamePremium.ToPtr(),
			},
			Properties: &armcontainerregistry.RegistryProperties{
				AdminUserEnabled: to.BoolPtr(false),
			},
		},
		nil,
	)
	require.NoError(t, err)
	registryResp, err := registryPollerResp.PollUntilDone(ctx, 10*time.Second)
	require.NoError(t, err)
	require.Equal(t, registryName, *registryResp.Name)

	// get container registry
	getResp, err := registriesClient.Get(ctx, rgName, registryName, nil)
	require.NoError(t, err)
	require.Equal(t, registryName, *getResp.Name)
}

func TestRegistriesClient_GetBuildSourceUploadURL(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")
	ctx := context.Background()
	location := "westus"

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "rg", location)
	defer clean()
	rgName := *rg.Name

	// container registry
	registriesClient := armcontainerregistry.NewRegistriesClient(subscriptionID, cred, opt)
	registryName, _ := createRandomName(t, "registry")
	registryPollerResp, err := registriesClient.BeginCreate(
		ctx,
		rgName,
		registryName,
		armcontainerregistry.Registry{
			Location: to.StringPtr(location),
			Tags: map[string]*string{
				"key": to.StringPtr("value"),
			},
			SKU: &armcontainerregistry.SKU{
				Name: armcontainerregistry.SKUNamePremium.ToPtr(),
			},
			Properties: &armcontainerregistry.RegistryProperties{
				AdminUserEnabled: to.BoolPtr(false),
			},
		},
		nil,
	)
	require.NoError(t, err)
	registryResp, err := registryPollerResp.PollUntilDone(ctx, 10*time.Second)
	require.NoError(t, err)
	require.Equal(t, registryName, *registryResp.Name)

	// get build source upload url
	getResp, err := registriesClient.GetBuildSourceUploadURL(ctx, rgName, registryName, nil)
	require.NoError(t, err)
	require.NotNil(t, getResp)
}

func createContainerRegistry(t *testing.T, ctx context.Context, registriesClient *armcontainerregistry.RegistriesClient, rgName, location string) *armcontainerregistry.Registry {
	registryName, _ := createRandomName(t, "registry")
	registryPollerResp, err := registriesClient.BeginCreate(
		ctx,
		rgName,
		registryName,
		armcontainerregistry.Registry{
			Location: to.StringPtr(location),
			Tags: map[string]*string{
				"key": to.StringPtr("value"),
			},
			SKU: &armcontainerregistry.SKU{
				Name: armcontainerregistry.SKUNamePremium.ToPtr(),
			},
			Properties: &armcontainerregistry.RegistryProperties{
				AdminUserEnabled: to.BoolPtr(false),
			},
		},
		nil,
	)
	require.NoError(t, err)
	registryResp, err := registryPollerResp.PollUntilDone(ctx, 10*time.Second)
	require.NoError(t, err)
	require.Equal(t, registryName, *registryResp.Name)
	return &registryResp.Registry
}
