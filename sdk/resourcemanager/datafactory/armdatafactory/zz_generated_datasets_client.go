//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// DatasetsClient contains the methods for the Datasets group.
// Don't use this type directly, use NewDatasetsClient() instead.
type DatasetsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDatasetsClient creates a new instance of DatasetsClient with the specified values.
func NewDatasetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DatasetsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &DatasetsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Creates or updates a dataset.
// If the operation fails it returns the *CloudError error type.
func (client *DatasetsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, datasetName string, dataset DatasetResource, options *DatasetsCreateOrUpdateOptions) (DatasetsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, factoryName, datasetName, dataset, options)
	if err != nil {
		return DatasetsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatasetsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatasetsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DatasetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, datasetName string, dataset DatasetResource, options *DatasetsCreateOrUpdateOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, dataset)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DatasetsClient) createOrUpdateHandleResponse(resp *http.Response) (DatasetsCreateOrUpdateResponse, error) {
	result := DatasetsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatasetResource); err != nil {
		return DatasetsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *DatasetsClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes a dataset.
// If the operation fails it returns the *CloudError error type.
func (client *DatasetsClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, datasetName string, options *DatasetsDeleteOptions) (DatasetsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, factoryName, datasetName, options)
	if err != nil {
		return DatasetsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatasetsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DatasetsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return DatasetsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DatasetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, datasetName string, options *DatasetsDeleteOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DatasetsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets a dataset.
// If the operation fails it returns the *CloudError error type.
func (client *DatasetsClient) Get(ctx context.Context, resourceGroupName string, factoryName string, datasetName string, options *DatasetsGetOptions) (DatasetsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, factoryName, datasetName, options)
	if err != nil {
		return DatasetsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatasetsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return DatasetsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DatasetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, datasetName string, options *DatasetsGetOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DatasetsClient) getHandleResponse(resp *http.Response) (DatasetsGetResponse, error) {
	result := DatasetsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatasetResource); err != nil {
		return DatasetsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DatasetsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByFactory - Lists datasets.
// If the operation fails it returns the *CloudError error type.
func (client *DatasetsClient) ListByFactory(resourceGroupName string, factoryName string, options *DatasetsListByFactoryOptions) *DatasetsListByFactoryPager {
	return &DatasetsListByFactoryPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByFactoryCreateRequest(ctx, resourceGroupName, factoryName, options)
		},
		advancer: func(ctx context.Context, resp DatasetsListByFactoryResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DatasetListResponse.NextLink)
		},
	}
}

// listByFactoryCreateRequest creates the ListByFactory request.
func (client *DatasetsClient) listByFactoryCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, options *DatasetsListByFactoryOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByFactoryHandleResponse handles the ListByFactory response.
func (client *DatasetsClient) listByFactoryHandleResponse(resp *http.Response) (DatasetsListByFactoryResponse, error) {
	result := DatasetsListByFactoryResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatasetListResponse); err != nil {
		return DatasetsListByFactoryResponse{}, err
	}
	return result, nil
}

// listByFactoryHandleError handles the ListByFactory error response.
func (client *DatasetsClient) listByFactoryHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
