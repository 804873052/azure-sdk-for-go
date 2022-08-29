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

type CassandraresourcesTestSuite struct {
	suite.Suite

	ctx               context.Context
	cred              azcore.TokenCredential
	options           *arm.ClientOptions
	accountName       string
	keyspaceName      string
	location          string
	resourceGroupName string
	subscriptionId    string
}

func (testsuite *CassandraresourcesTestSuite) SetupSuite() {
	testsuite.ctx = context.Background()
	testsuite.cred, testsuite.options = testutil.GetCredAndClientOptions(testsuite.T())
	testsuite.accountName = "databaseaccountcassandra"
	testsuite.keyspaceName = testutil.GenerateAlphaNumericID(testsuite.T(), "cassandrakeyspace", 6)
	testsuite.location = testutil.GetEnv("LOCATION", "westus")
	testsuite.resourceGroupName = testutil.GetEnv("RESOURCE_GROUP_NAME", "scenarioTestTempGroup")
	testsuite.subscriptionId = testutil.GetEnv("AZURE_SUBSCRIPTION_ID", "")

	testutil.StartRecording(testsuite.T(), "sdk/resourcemanager/cosmos/armcosmos/testdata")
	resourceGroup, _, err := testutil.CreateResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.location)
	testsuite.Require().NoError(err)
	testsuite.resourceGroupName = *resourceGroup.Name
	testsuite.Prepare()
}

func (testsuite *CassandraresourcesTestSuite) TearDownSuite() {
	_, err := testutil.DeleteResourceGroup(testsuite.ctx, testsuite.subscriptionId, testsuite.cred, testsuite.options, testsuite.resourceGroupName)
	testsuite.Require().NoError(err)
	testutil.StopRecording(testsuite.T())
}

func TestCassandraresourcesTestSuite(t *testing.T) {
	suite.Run(t, new(CassandraresourcesTestSuite))
}

func (testsuite *CassandraresourcesTestSuite) Prepare() {
	var err error
	// From step DatabaseAccount_Create
	databaseAccountsClient, err := armcosmos.NewDatabaseAccountsClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	databaseAccountsClientCreateOrUpdateResponsePoller, err := databaseAccountsClient.BeginCreateOrUpdate(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, armcosmos.DatabaseAccountCreateUpdateParameters{
		Location: to.Ptr(testsuite.location),
		Properties: &armcosmos.DatabaseAccountCreateUpdateProperties{
			Capabilities: []*armcosmos.Capability{
				{
					Name: to.Ptr("EnableCassandra"),
				}},
			CreateMode:               to.Ptr(armcosmos.CreateModeDefault),
			DatabaseAccountOfferType: to.Ptr("Standard"),
			Locations: []*armcosmos.Location{
				{
					FailoverPriority: to.Ptr[int32](0),
					IsZoneRedundant:  to.Ptr(false),
					LocationName:     to.Ptr("eastus"),
				}},
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, databaseAccountsClientCreateOrUpdateResponsePoller)
	testsuite.Require().NoError(err)
}

// Microsoft.DocumentDB/databaseAccounts/cassandraKeyspaces
func (testsuite *CassandraresourcesTestSuite) TestCassandraresources() {
	var err error
	// From step CassandraResources_CreateCassandraKeyspace
	cassandraResourcesClient, err := armcosmos.NewCassandraResourcesClient(testsuite.subscriptionId, testsuite.cred, testsuite.options)
	testsuite.Require().NoError(err)
	cassandraResourcesClientCreateUpdateCassandraKeyspaceResponsePoller, err := cassandraResourcesClient.BeginCreateUpdateCassandraKeyspace(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, testsuite.keyspaceName, armcosmos.CassandraKeyspaceCreateUpdateParameters{
		Location: to.Ptr(testsuite.location),
		Tags:     map[string]*string{},
		Properties: &armcosmos.CassandraKeyspaceCreateUpdateProperties{
			Options: &armcosmos.CreateUpdateOptions{
				Throughput: to.Ptr[int32](2000),
			},
			Resource: &armcosmos.CassandraKeyspaceResource{
				ID: to.Ptr(testsuite.keyspaceName),
			},
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, cassandraResourcesClientCreateUpdateCassandraKeyspaceResponsePoller)
	testsuite.Require().NoError(err)

	// From step CassandraResources_GetCassandraKeyspace
	_, err = cassandraResourcesClient.GetCassandraKeyspace(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, testsuite.keyspaceName, nil)
	testsuite.Require().NoError(err)

	// From step CassandraResources_ListCassandraKeyspaces
	cassandraResourcesClientNewListCassandraKeyspacesPager := cassandraResourcesClient.NewListCassandraKeyspacesPager(testsuite.resourceGroupName, testsuite.accountName, nil)
	for cassandraResourcesClientNewListCassandraKeyspacesPager.More() {
		_, err := cassandraResourcesClientNewListCassandraKeyspacesPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step CassandraResources_UpdateCassandraKeyspaceThroughput
	cassandraResourcesClientUpdateCassandraKeyspaceThroughputResponsePoller, err := cassandraResourcesClient.BeginUpdateCassandraKeyspaceThroughput(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, testsuite.keyspaceName, armcosmos.ThroughputSettingsUpdateParameters{
		Location: to.Ptr(testsuite.location),
		Tags:     map[string]*string{},
		Properties: &armcosmos.ThroughputSettingsUpdateProperties{
			Resource: &armcosmos.ThroughputSettingsResource{
				Throughput: to.Ptr[int32](400),
			},
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, cassandraResourcesClientUpdateCassandraKeyspaceThroughputResponsePoller)
	testsuite.Require().NoError(err)

	// From step CassandraResources_GetCassandraKeyspaceThroughput
	_, err = cassandraResourcesClient.GetCassandraKeyspaceThroughput(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, testsuite.keyspaceName, nil)
	testsuite.Require().NoError(err)

	// From step CassandraResources_MigrateCassandraKeyspaceToAutoscale
	cassandraResourcesClientMigrateCassandraKeyspaceToAutoscaleResponsePoller, err := cassandraResourcesClient.BeginMigrateCassandraKeyspaceToAutoscale(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, testsuite.keyspaceName, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, cassandraResourcesClientMigrateCassandraKeyspaceToAutoscaleResponsePoller)
	testsuite.Require().NoError(err)

	// From step CassandraResources_MigrateCassandraKeyspaceToManualThroughput
	cassandraResourcesClientMigrateCassandraKeyspaceToManualThroughputResponsePoller, err := cassandraResourcesClient.BeginMigrateCassandraKeyspaceToManualThroughput(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, testsuite.keyspaceName, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, cassandraResourcesClientMigrateCassandraKeyspaceToManualThroughputResponsePoller)
	testsuite.Require().NoError(err)

	// From step CassandraResources_CreateCassandraTable
	cassandraResourcesClientCreateUpdateCassandraTableResponsePoller, err := cassandraResourcesClient.BeginCreateUpdateCassandraTable(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, testsuite.keyspaceName, "tableName", armcosmos.CassandraTableCreateUpdateParameters{
		Location: to.Ptr(testsuite.location),
		Tags:     map[string]*string{},
		Properties: &armcosmos.CassandraTableCreateUpdateProperties{
			Options: &armcosmos.CreateUpdateOptions{
				Throughput: to.Ptr[int32](2000),
			},
			Resource: &armcosmos.CassandraTableResource{
				Schema: &armcosmos.CassandraSchema{
					Columns: []*armcosmos.Column{
						{
							Name: to.Ptr("columnA"),
							Type: to.Ptr("Ascii"),
						}},
					PartitionKeys: []*armcosmos.CassandraPartitionKey{
						{
							Name: to.Ptr("columnA"),
						}},
				},
				DefaultTTL: to.Ptr[int32](100),
				ID:         to.Ptr("tableName"),
			},
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, cassandraResourcesClientCreateUpdateCassandraTableResponsePoller)
	testsuite.Require().NoError(err)

	// From step CassandraResources_GetCassandraTable
	_, err = cassandraResourcesClient.GetCassandraTable(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, testsuite.keyspaceName, "tableName", nil)
	testsuite.Require().NoError(err)

	// From step CassandraResources_ListCassandraTables
	cassandraResourcesClientNewListCassandraTablesPager := cassandraResourcesClient.NewListCassandraTablesPager(testsuite.resourceGroupName, testsuite.accountName, testsuite.keyspaceName, nil)
	for cassandraResourcesClientNewListCassandraTablesPager.More() {
		_, err := cassandraResourcesClientNewListCassandraTablesPager.NextPage(testsuite.ctx)
		testsuite.Require().NoError(err)
		break
	}

	// From step CassandraResources_UpdateCassandraTableThroughput
	cassandraResourcesClientUpdateCassandraTableThroughputResponsePoller, err := cassandraResourcesClient.BeginUpdateCassandraTableThroughput(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, testsuite.keyspaceName, "tableName", armcosmos.ThroughputSettingsUpdateParameters{
		Location: to.Ptr(testsuite.location),
		Tags:     map[string]*string{},
		Properties: &armcosmos.ThroughputSettingsUpdateProperties{
			Resource: &armcosmos.ThroughputSettingsResource{
				Throughput: to.Ptr[int32](400),
			},
		},
	}, nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, cassandraResourcesClientUpdateCassandraTableThroughputResponsePoller)
	testsuite.Require().NoError(err)

	// From step CassandraResources_GetCassandraTableThroughput
	_, err = cassandraResourcesClient.GetCassandraTableThroughput(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, testsuite.keyspaceName, "tableName", nil)
	testsuite.Require().NoError(err)

	// From step CassandraResources_MigrateCassandraTableToAutoscale
	cassandraResourcesClientMigrateCassandraTableToAutoscaleResponsePoller, err := cassandraResourcesClient.BeginMigrateCassandraTableToAutoscale(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, testsuite.keyspaceName, "tableName", nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, cassandraResourcesClientMigrateCassandraTableToAutoscaleResponsePoller)
	testsuite.Require().NoError(err)

	// From step CassandraResources_MigrateCassandraTableToManualThroughput
	cassandraResourcesClientMigrateCassandraTableToManualThroughputResponsePoller, err := cassandraResourcesClient.BeginMigrateCassandraTableToManualThroughput(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, testsuite.keyspaceName, "tableName", nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, cassandraResourcesClientMigrateCassandraTableToManualThroughputResponsePoller)
	testsuite.Require().NoError(err)

	// From step CassandraResources_DeleteCassandraTable
	cassandraResourcesClientDeleteCassandraTableResponsePoller, err := cassandraResourcesClient.BeginDeleteCassandraTable(testsuite.ctx, testsuite.resourceGroupName, testsuite.accountName, testsuite.keyspaceName, "tableName", nil)
	testsuite.Require().NoError(err)
	_, err = testutil.PollForTest(testsuite.ctx, cassandraResourcesClientDeleteCassandraTableResponsePoller)
	testsuite.Require().NoError(err)
}
