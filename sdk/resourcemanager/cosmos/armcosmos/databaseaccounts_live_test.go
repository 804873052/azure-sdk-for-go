//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/internal/testutil"
	"github.com/stretchr/testify/suite"
)

type DatabaseaccountsTestSuite struct {
	suite.Suite

	ctx               context.Context
	cred              azcore.TokenCredential
	options           *arm.ClientOptions
	location          string
	resourceGroupName string
	subscriptionId    string
}

func (testsuite *DatabaseaccountsTestSuite) SetupSuite() {
	testsuite.ctx = context.Background()
	testsuite.cred, testsuite.options = testutil.GetCredAndClientOptions(testsuite.T())
	testsuite.location = testutil.GetEnv("LOCATION", "westus")
	testsuite.resourceGroupName = testutil.GetEnv("RESOURCE_GROUP_NAME", "scenarioTestTempGroup")
	testsuite.subscriptionId = testutil.GetEnv("AZURE_SUBSCRIPTION_ID", "")

	testutil.StartRecording(testsuite.T(), "sdk/resourcemanager/cosmos/armcosmos/testdata")
	resourceGroup, _, err := testutil.CreateResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.location)
	testsuite.Require().NoError(err)
	testsuite.resourceGroupName = *resourceGroup.Name
}

func (testsuite *DatabaseaccountsTestSuite) TearDownSuite() {
	_, err := testutil.DeleteResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.resourceGroupName)
	testsuite.Require().NoError(err)
	testutil.StopRecording(testsuite.T())
}

func TestDatabaseaccountsTestSuite(t *testing.T) {
	suite.Run(t, new(DatabaseaccountsTestSuite))
}

// Microsoft.DocumentDB/databaseAccounts
func (testsuite *DatabaseaccountsTestSuite) TestDatabaseaccount() {
	accountName := "databaseaccountcosmos"
	var err error
	// From step DatabaseAccount_Create
	databaseAccountsClient, err := armcosmos.NewDatabaseAccountsClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	databaseAccountsClientCreateOrUpdateResponsePoller, err := databaseAccountsClient.BeginCreateOrUpdate(testsuite.ctx, testsuite.resourceGroupName, accountName, armcosmos.DatabaseAccountCreateUpdateParameters{
		Location: to.Ptr(testsuite.location),
		Properties: &armcosmos.DatabaseAccountCreateUpdateProperties{
			CreateMode:               to.Ptr(armcosmos.CreateModeDefault),
			DatabaseAccountOfferType: to.Ptr("Standard"),
			Locations: []*armcosmos.Location{
				{
					FailoverPriority: to.Ptr[int32](2),
					LocationName:     to.Ptr("southcentralus"),
				},
				{
					FailoverPriority: to.Ptr[int32](1),
					LocationName:     to.Ptr("eastus"),
				},
				{
					FailoverPriority: to.Ptr[int32](0),
					LocationName:     to.Ptr("westus"),
				}},
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, databaseAccountsClientCreateOrUpdateResponsePoller)
	testsuite.Require().NoError(err)

	// From step DatabaseAccount_CheckNameExists
	_, err = databaseAccountsClient.CheckNameExists(testsuite.ctx, accountName, nil)
	testsuite.Require().NoError(err)

	// From step DatabaseAccount_Get
	_, err = databaseAccountsClient.Get(testsuite.ctx, testsuite.resourceGroupName, accountName, nil)
	testsuite.Require().NoError(err)

	// From step DatabaseAccount_List
	databaseAccountsClientNewListPager := databaseAccountsClient.NewListPager(nil)
	for databaseAccountsClientNewListPager.More() {
		_, err := databaseAccountsClientNewListPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step DatabaseAccount_ListByResourceGroup
	databaseAccountsClientNewListByResourceGroupPager := databaseAccountsClient.NewListByResourceGroupPager(testsuite.resourceGroupName, nil)
	for databaseAccountsClientNewListByResourceGroupPager.More() {
		_, err := databaseAccountsClientNewListByResourceGroupPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step DatabaseAccount_ListKeys
	_, err = databaseAccountsClient.ListKeys(testsuite.ctx, testsuite.resourceGroupName, accountName, nil)
	testsuite.Require().NoError(err)

	// From step DatabaseAccount_ListConnectionStrings
	_, err = databaseAccountsClient.ListConnectionStrings(testsuite.ctx, testsuite.resourceGroupName, accountName, nil)
	testsuite.Require().NoError(err)

	// From step DatabaseAccount_GetReadOnlyKeys
	_, err = databaseAccountsClient.GetReadOnlyKeys(testsuite.ctx, testsuite.resourceGroupName, accountName, nil)
	testsuite.Require().NoError(err)

	// From step DatabaseAccount_ListReadOnlyKeys
	_, err = databaseAccountsClient.ListReadOnlyKeys(testsuite.ctx, testsuite.resourceGroupName, accountName, nil)
	testsuite.Require().NoError(err)

	// From step DatabaseAccount_ListMetricDefinitions
	databaseAccountsClientNewListMetricDefinitionsPager := databaseAccountsClient.NewListMetricDefinitionsPager(testsuite.resourceGroupName, accountName, nil)
	for databaseAccountsClientNewListMetricDefinitionsPager.More() {
		_, err := databaseAccountsClientNewListMetricDefinitionsPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step DatabaseAccount_FailoverPriorityChange
	databaseAccountsClientFailoverPriorityChangeResponsePoller, err := databaseAccountsClient.BeginFailoverPriorityChange(testsuite.ctx, testsuite.resourceGroupName, accountName, armcosmos.FailoverPolicies{
		FailoverPolicies: []*armcosmos.FailoverPolicy{
			{
				FailoverPriority: to.Ptr[int32](0),
				LocationName:     to.Ptr("eastus"),
			},
			{
				FailoverPriority: to.Ptr[int32](2),
				LocationName:     to.Ptr("southcentralus"),
			},
			{
				FailoverPriority: to.Ptr[int32](1),
				LocationName:     to.Ptr("westus"),
			}},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, databaseAccountsClientFailoverPriorityChangeResponsePoller)
	testsuite.Require().NoError(err)

	// From step DatabaseAccount_RegenerateKey
	databaseAccountsClientRegenerateKeyResponsePoller, err := databaseAccountsClient.BeginRegenerateKey(testsuite.ctx, testsuite.resourceGroupName, accountName, armcosmos.DatabaseAccountRegenerateKeyParameters{
		KeyKind: to.Ptr(armcosmos.KeyKindPrimary),
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, databaseAccountsClientRegenerateKeyResponsePoller)
	testsuite.Require().NoError(err)

	// From step DatabaseAccount_Update
	databaseAccountsClientUpdateResponsePoller, err := databaseAccountsClient.BeginUpdate(testsuite.ctx, testsuite.resourceGroupName, accountName, armcosmos.DatabaseAccountUpdateParameters{
		Tags: map[string]*string{
			"dept": to.Ptr("finance"),
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, databaseAccountsClientUpdateResponsePoller)
	testsuite.Require().NoError(err)

	// From step DatabaseAccount_Delete
	databaseAccountsClientDeleteResponsePoller, err := databaseAccountsClient.BeginDelete(testsuite.ctx, testsuite.resourceGroupName, accountName, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, databaseAccountsClientDeleteResponsePoller)
	testsuite.Require().NoError(err)
}
