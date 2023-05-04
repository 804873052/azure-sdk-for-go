//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdatafactory

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DatasetsClient contains the methods for the Datasets group.
// Don't use this type directly, use NewDatasetsClient() instead.
type DatasetsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDatasetsClient creates a new instance of DatasetsClient with the specified values.
//   - subscriptionID - The subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDatasetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DatasetsClient, error) {
	cl, err := arm.NewClient(moduleName+".DatasetsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DatasetsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a dataset.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - datasetName - The dataset name.
//   - dataset - Dataset resource definition.
//   - options - DatasetsClientCreateOrUpdateOptions contains the optional parameters for the DatasetsClient.CreateOrUpdate method.
func (client *DatasetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, datasetName string, dataset DatasetResource, options *DatasetsClientCreateOrUpdateOptions) (DatasetsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, factoryName, datasetName, dataset, options)
	if err != nil {
		return DatasetsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatasetsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatasetsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DatasetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, datasetName string, dataset DatasetResource, options *DatasetsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/datasets/{datasetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if datasetName == "" {
		return nil, errors.New("parameter datasetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, dataset)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DatasetsClient) createOrUpdateHandleResponse(resp *http.Response) (DatasetsClientCreateOrUpdateResponse, error) {
	result := DatasetsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatasetResource); err != nil {
		return DatasetsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a dataset.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - datasetName - The dataset name.
//   - options - DatasetsClientDeleteOptions contains the optional parameters for the DatasetsClient.Delete method.
func (client *DatasetsClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, datasetName string, options *DatasetsClientDeleteOptions) (DatasetsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, factoryName, datasetName, options)
	if err != nil {
		return DatasetsClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatasetsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DatasetsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DatasetsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DatasetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, datasetName string, options *DatasetsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/datasets/{datasetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if datasetName == "" {
		return nil, errors.New("parameter datasetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a dataset.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - datasetName - The dataset name.
//   - options - DatasetsClientGetOptions contains the optional parameters for the DatasetsClient.Get method.
func (client *DatasetsClient) Get(ctx context.Context, resourceGroupName string, factoryName string, datasetName string, options *DatasetsClientGetOptions) (DatasetsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, factoryName, datasetName, options)
	if err != nil {
		return DatasetsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatasetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return DatasetsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DatasetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, datasetName string, options *DatasetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/datasets/{datasetName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if datasetName == "" {
		return nil, errors.New("parameter datasetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DatasetsClient) getHandleResponse(resp *http.Response) (DatasetsClientGetResponse, error) {
	result := DatasetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatasetResource); err != nil {
		return DatasetsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByFactoryPager - Lists datasets.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The resource group name.
//   - factoryName - The factory name.
//   - options - DatasetsClientListByFactoryOptions contains the optional parameters for the DatasetsClient.NewListByFactoryPager
//     method.
func (client *DatasetsClient) NewListByFactoryPager(resourceGroupName string, factoryName string, options *DatasetsClientListByFactoryOptions) *runtime.Pager[DatasetsClientListByFactoryResponse] {
	return runtime.NewPager(runtime.PagingHandler[DatasetsClientListByFactoryResponse]{
		More: func(page DatasetsClientListByFactoryResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DatasetsClientListByFactoryResponse) (DatasetsClientListByFactoryResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByFactoryCreateRequest(ctx, resourceGroupName, factoryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DatasetsClientListByFactoryResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DatasetsClientListByFactoryResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DatasetsClientListByFactoryResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByFactoryHandleResponse(resp)
		},
	})
}

// listByFactoryCreateRequest creates the ListByFactory request.
func (client *DatasetsClient) listByFactoryCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, options *DatasetsClientListByFactoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/datasets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByFactoryHandleResponse handles the ListByFactory response.
func (client *DatasetsClient) listByFactoryHandleResponse(resp *http.Response) (DatasetsClientListByFactoryResponse, error) {
	result := DatasetsClientListByFactoryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatasetListResponse); err != nil {
		return DatasetsClientListByFactoryResponse{}, err
	}
	return result, nil
}
