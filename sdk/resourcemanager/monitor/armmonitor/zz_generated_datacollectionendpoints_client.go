//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

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

// DataCollectionEndpointsClient contains the methods for the DataCollectionEndpoints group.
// Don't use this type directly, use NewDataCollectionEndpointsClient() instead.
type DataCollectionEndpointsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDataCollectionEndpointsClient creates a new instance of DataCollectionEndpointsClient with the specified values.
func NewDataCollectionEndpointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DataCollectionEndpointsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &DataCollectionEndpointsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Create - Creates or updates a data collection endpoint.
// If the operation fails it returns the *ErrorResponseCommonV2 error type.
func (client *DataCollectionEndpointsClient) Create(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *DataCollectionEndpointsCreateOptions) (DataCollectionEndpointsCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, dataCollectionEndpointName, options)
	if err != nil {
		return DataCollectionEndpointsCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataCollectionEndpointsCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return DataCollectionEndpointsCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *DataCollectionEndpointsClient) createCreateRequest(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *DataCollectionEndpointsCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionEndpoints/{dataCollectionEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataCollectionEndpointName == "" {
		return nil, errors.New("parameter dataCollectionEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCollectionEndpointName}", url.PathEscape(dataCollectionEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *DataCollectionEndpointsClient) createHandleResponse(resp *http.Response) (DataCollectionEndpointsCreateResponse, error) {
	result := DataCollectionEndpointsCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCollectionEndpointResource); err != nil {
		return DataCollectionEndpointsCreateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *DataCollectionEndpointsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponseCommonV2{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes a data collection endpoint.
// If the operation fails it returns the *ErrorResponseCommonV2 error type.
func (client *DataCollectionEndpointsClient) Delete(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *DataCollectionEndpointsDeleteOptions) (DataCollectionEndpointsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, dataCollectionEndpointName, options)
	if err != nil {
		return DataCollectionEndpointsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataCollectionEndpointsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DataCollectionEndpointsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return DataCollectionEndpointsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DataCollectionEndpointsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *DataCollectionEndpointsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionEndpoints/{dataCollectionEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataCollectionEndpointName == "" {
		return nil, errors.New("parameter dataCollectionEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCollectionEndpointName}", url.PathEscape(dataCollectionEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *DataCollectionEndpointsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponseCommonV2{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Returns the specified data collection endpoint.
// If the operation fails it returns the *ErrorResponseCommonV2 error type.
func (client *DataCollectionEndpointsClient) Get(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *DataCollectionEndpointsGetOptions) (DataCollectionEndpointsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, dataCollectionEndpointName, options)
	if err != nil {
		return DataCollectionEndpointsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataCollectionEndpointsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataCollectionEndpointsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DataCollectionEndpointsClient) getCreateRequest(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *DataCollectionEndpointsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionEndpoints/{dataCollectionEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataCollectionEndpointName == "" {
		return nil, errors.New("parameter dataCollectionEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCollectionEndpointName}", url.PathEscape(dataCollectionEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DataCollectionEndpointsClient) getHandleResponse(resp *http.Response) (DataCollectionEndpointsGetResponse, error) {
	result := DataCollectionEndpointsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCollectionEndpointResource); err != nil {
		return DataCollectionEndpointsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *DataCollectionEndpointsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponseCommonV2{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Lists all data collection endpoints in the specified resource group.
// If the operation fails it returns the *ErrorResponseCommonV2 error type.
func (client *DataCollectionEndpointsClient) ListByResourceGroup(resourceGroupName string, options *DataCollectionEndpointsListByResourceGroupOptions) *DataCollectionEndpointsListByResourceGroupPager {
	return &DataCollectionEndpointsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp DataCollectionEndpointsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DataCollectionEndpointResourceListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DataCollectionEndpointsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DataCollectionEndpointsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionEndpoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DataCollectionEndpointsClient) listByResourceGroupHandleResponse(resp *http.Response) (DataCollectionEndpointsListByResourceGroupResponse, error) {
	result := DataCollectionEndpointsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCollectionEndpointResourceListResult); err != nil {
		return DataCollectionEndpointsListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *DataCollectionEndpointsClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponseCommonV2{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - Lists all data collection endpoints in the specified subscription
// If the operation fails it returns the *ErrorResponseCommonV2 error type.
func (client *DataCollectionEndpointsClient) ListBySubscription(options *DataCollectionEndpointsListBySubscriptionOptions) *DataCollectionEndpointsListBySubscriptionPager {
	return &DataCollectionEndpointsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp DataCollectionEndpointsListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DataCollectionEndpointResourceListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DataCollectionEndpointsClient) listBySubscriptionCreateRequest(ctx context.Context, options *DataCollectionEndpointsListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/dataCollectionEndpoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DataCollectionEndpointsClient) listBySubscriptionHandleResponse(resp *http.Response) (DataCollectionEndpointsListBySubscriptionResponse, error) {
	result := DataCollectionEndpointsListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCollectionEndpointResourceListResult); err != nil {
		return DataCollectionEndpointsListBySubscriptionResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *DataCollectionEndpointsClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponseCommonV2{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Updates part of a data collection endpoint.
// If the operation fails it returns the *ErrorResponseCommonV2 error type.
func (client *DataCollectionEndpointsClient) Update(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *DataCollectionEndpointsUpdateOptions) (DataCollectionEndpointsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, dataCollectionEndpointName, options)
	if err != nil {
		return DataCollectionEndpointsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataCollectionEndpointsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataCollectionEndpointsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *DataCollectionEndpointsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, dataCollectionEndpointName string, options *DataCollectionEndpointsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/dataCollectionEndpoints/{dataCollectionEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataCollectionEndpointName == "" {
		return nil, errors.New("parameter dataCollectionEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataCollectionEndpointName}", url.PathEscape(dataCollectionEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *DataCollectionEndpointsClient) updateHandleResponse(resp *http.Response) (DataCollectionEndpointsUpdateResponse, error) {
	result := DataCollectionEndpointsUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DataCollectionEndpointResource); err != nil {
		return DataCollectionEndpointsUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *DataCollectionEndpointsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponseCommonV2{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
