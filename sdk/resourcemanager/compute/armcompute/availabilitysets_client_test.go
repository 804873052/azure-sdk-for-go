//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armcompute_test

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/internal/recording"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAvailabilitySetsClient(t *testing.T) {
	stop := startTest(t)
	defer stop()

	cred, opt := authenticateTest(t)
	subscriptionID := recording.GetEnvVariable("AZURE_SUBSCRIPTION_ID", "00000000-0000-0000-0000-000000000000")

	// create resource group
	rg, clean := createResourceGroup(t, cred, opt, subscriptionID, "createAS", "westus")
	rgName := *rg.Name
	defer clean()

	// create availability sets
	client := armcompute.NewAvailabilitySetsClient(subscriptionID, cred, opt)
	name, err := createRandomName(t, "set")
	require.NoError(t, err)
	resp, err := client.CreateOrUpdate(
		context.Background(),
		rgName,
		name,
		armcompute.AvailabilitySet{
			Location: to.StringPtr("westus"),
			SKU: &armcompute.SKU{
				Name: to.StringPtr(string(armcompute.AvailabilitySetSKUTypesAligned)),
			},
			Properties: &armcompute.AvailabilitySetProperties{
				PlatformFaultDomainCount:  to.Int32Ptr(1),
				PlatformUpdateDomainCount: to.Int32Ptr(1),
			},
		},
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, *resp.Name, name)

	// get
	getResp, err := client.Get(context.Background(), rgName, name, nil)
	require.NoError(t, err)
	require.Equal(t, *getResp.Name, name)

	// list
	listPager := client.List(rgName, nil)
	require.Equal(t, listPager.Err(), nil)

	// list available size
	listResp, err := client.ListAvailableSizes(context.Background(), rgName, name, nil)
	require.NoError(t, err)
	require.Equal(t, listResp.RawResponse.StatusCode, 200)

	// list by subscription
	listBySubscription := client.ListBySubscription(nil)
	require.NoError(t, listBySubscription.Err())
	require.Equal(t, listBySubscription.NextPage(context.Background()), true)

	// update
	updateResp, err := client.Update(
		context.Background(),
		rgName,
		name,
		armcompute.AvailabilitySetUpdate{
			Tags: map[string]*string{
				"tag": to.StringPtr("value"),
			},
		},
		nil,
	)
	require.NoError(t, err)
	require.Equal(t, *updateResp.Name, name)

	// delete
	delResp, err := client.Delete(context.Background(), rgName, name, nil)
	require.NoError(t, err)
	require.Equal(t, delResp.RawResponse.StatusCode, 200)
}
