//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// MigrationRecoveryPointsClient contains the methods for the MigrationRecoveryPoints group.
// Don't use this type directly, use NewMigrationRecoveryPointsClient() instead.
type MigrationRecoveryPointsClient struct {
	host              string
	resourceName      string
	resourceGroupName string
	subscriptionID    string
	pl                runtime.Pipeline
}

// NewMigrationRecoveryPointsClient creates a new instance of MigrationRecoveryPointsClient with the specified values.
// resourceName - The name of the recovery services vault.
// resourceGroupName - The name of the resource group where the recovery services vault is present.
// subscriptionID - The subscription Id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMigrationRecoveryPointsClient(resourceName string, resourceGroupName string, subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MigrationRecoveryPointsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &MigrationRecoveryPointsClient{
		resourceName:      resourceName,
		resourceGroupName: resourceGroupName,
		subscriptionID:    subscriptionID,
		host:              ep,
		pl:                pl,
	}
	return client, nil
}

// Get - Gets a recovery point for a migration item.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01
// fabricName - Fabric unique name.
// protectionContainerName - Protection container name.
// migrationItemName - Migration item name.
// migrationRecoveryPointName - The migration recovery point name.
// options - MigrationRecoveryPointsClientGetOptions contains the optional parameters for the MigrationRecoveryPointsClient.Get
// method.
func (client *MigrationRecoveryPointsClient) Get(ctx context.Context, fabricName string, protectionContainerName string, migrationItemName string, migrationRecoveryPointName string, options *MigrationRecoveryPointsClientGetOptions) (MigrationRecoveryPointsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, fabricName, protectionContainerName, migrationItemName, migrationRecoveryPointName, options)
	if err != nil {
		return MigrationRecoveryPointsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MigrationRecoveryPointsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MigrationRecoveryPointsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MigrationRecoveryPointsClient) getCreateRequest(ctx context.Context, fabricName string, protectionContainerName string, migrationItemName string, migrationRecoveryPointName string, options *MigrationRecoveryPointsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}/migrationRecoveryPoints/{migrationRecoveryPointName}"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if protectionContainerName == "" {
		return nil, errors.New("parameter protectionContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectionContainerName}", url.PathEscape(protectionContainerName))
	if migrationItemName == "" {
		return nil, errors.New("parameter migrationItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrationItemName}", url.PathEscape(migrationItemName))
	if migrationRecoveryPointName == "" {
		return nil, errors.New("parameter migrationRecoveryPointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrationRecoveryPointName}", url.PathEscape(migrationRecoveryPointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MigrationRecoveryPointsClient) getHandleResponse(resp *http.Response) (MigrationRecoveryPointsClientGetResponse, error) {
	result := MigrationRecoveryPointsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MigrationRecoveryPoint); err != nil {
		return MigrationRecoveryPointsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByReplicationMigrationItemsPager - Gets the recovery points for a migration item.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01
// fabricName - Fabric unique name.
// protectionContainerName - Protection container name.
// migrationItemName - Migration item name.
// options - MigrationRecoveryPointsClientListByReplicationMigrationItemsOptions contains the optional parameters for the
// MigrationRecoveryPointsClient.ListByReplicationMigrationItems method.
func (client *MigrationRecoveryPointsClient) NewListByReplicationMigrationItemsPager(fabricName string, protectionContainerName string, migrationItemName string, options *MigrationRecoveryPointsClientListByReplicationMigrationItemsOptions) *runtime.Pager[MigrationRecoveryPointsClientListByReplicationMigrationItemsResponse] {
	return runtime.NewPager(runtime.PagingHandler[MigrationRecoveryPointsClientListByReplicationMigrationItemsResponse]{
		More: func(page MigrationRecoveryPointsClientListByReplicationMigrationItemsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MigrationRecoveryPointsClientListByReplicationMigrationItemsResponse) (MigrationRecoveryPointsClientListByReplicationMigrationItemsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByReplicationMigrationItemsCreateRequest(ctx, fabricName, protectionContainerName, migrationItemName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return MigrationRecoveryPointsClientListByReplicationMigrationItemsResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return MigrationRecoveryPointsClientListByReplicationMigrationItemsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MigrationRecoveryPointsClientListByReplicationMigrationItemsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByReplicationMigrationItemsHandleResponse(resp)
		},
	})
}

// listByReplicationMigrationItemsCreateRequest creates the ListByReplicationMigrationItems request.
func (client *MigrationRecoveryPointsClient) listByReplicationMigrationItemsCreateRequest(ctx context.Context, fabricName string, protectionContainerName string, migrationItemName string, options *MigrationRecoveryPointsClientListByReplicationMigrationItemsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationMigrationItems/{migrationItemName}/migrationRecoveryPoints"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if protectionContainerName == "" {
		return nil, errors.New("parameter protectionContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectionContainerName}", url.PathEscape(protectionContainerName))
	if migrationItemName == "" {
		return nil, errors.New("parameter migrationItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{migrationItemName}", url.PathEscape(migrationItemName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByReplicationMigrationItemsHandleResponse handles the ListByReplicationMigrationItems response.
func (client *MigrationRecoveryPointsClient) listByReplicationMigrationItemsHandleResponse(resp *http.Response) (MigrationRecoveryPointsClientListByReplicationMigrationItemsResponse, error) {
	result := MigrationRecoveryPointsClientListByReplicationMigrationItemsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MigrationRecoveryPointCollection); err != nil {
		return MigrationRecoveryPointsClientListByReplicationMigrationItemsResponse{}, err
	}
	return result, nil
}
