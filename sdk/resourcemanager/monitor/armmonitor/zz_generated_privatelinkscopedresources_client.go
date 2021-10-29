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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// PrivateLinkScopedResourcesClient contains the methods for the PrivateLinkScopedResources group.
// Don't use this type directly, use NewPrivateLinkScopedResourcesClient() instead.
type PrivateLinkScopedResourcesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPrivateLinkScopedResourcesClient creates a new instance of PrivateLinkScopedResourcesClient with the specified values.
func NewPrivateLinkScopedResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PrivateLinkScopedResourcesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &PrivateLinkScopedResourcesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreateOrUpdate - Approve or reject a private endpoint connection with a given name.
// If the operation fails it returns a generic error.
func (client *PrivateLinkScopedResourcesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, scopeName string, name string, parameters ScopedResource, options *PrivateLinkScopedResourcesBeginCreateOrUpdateOptions) (PrivateLinkScopedResourcesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, scopeName, name, parameters, options)
	if err != nil {
		return PrivateLinkScopedResourcesCreateOrUpdatePollerResponse{}, err
	}
	result := PrivateLinkScopedResourcesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PrivateLinkScopedResourcesClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return PrivateLinkScopedResourcesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &PrivateLinkScopedResourcesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Approve or reject a private endpoint connection with a given name.
// If the operation fails it returns a generic error.
func (client *PrivateLinkScopedResourcesClient) createOrUpdate(ctx context.Context, resourceGroupName string, scopeName string, name string, parameters ScopedResource, options *PrivateLinkScopedResourcesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, scopeName, name, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PrivateLinkScopedResourcesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, name string, parameters ScopedResource, options *PrivateLinkScopedResourcesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/privateLinkScopes/{scopeName}/scopedResources/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-17-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *PrivateLinkScopedResourcesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// BeginDelete - Deletes a private endpoint connection with a given name.
// If the operation fails it returns a generic error.
func (client *PrivateLinkScopedResourcesClient) BeginDelete(ctx context.Context, resourceGroupName string, scopeName string, name string, options *PrivateLinkScopedResourcesBeginDeleteOptions) (PrivateLinkScopedResourcesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, scopeName, name, options)
	if err != nil {
		return PrivateLinkScopedResourcesDeletePollerResponse{}, err
	}
	result := PrivateLinkScopedResourcesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PrivateLinkScopedResourcesClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return PrivateLinkScopedResourcesDeletePollerResponse{}, err
	}
	result.Poller = &PrivateLinkScopedResourcesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a private endpoint connection with a given name.
// If the operation fails it returns a generic error.
func (client *PrivateLinkScopedResourcesClient) deleteOperation(ctx context.Context, resourceGroupName string, scopeName string, name string, options *PrivateLinkScopedResourcesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, scopeName, name, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateLinkScopedResourcesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, name string, options *PrivateLinkScopedResourcesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/privateLinkScopes/{scopeName}/scopedResources/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-17-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *PrivateLinkScopedResourcesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets a scoped resource in a private link scope.
// If the operation fails it returns a generic error.
func (client *PrivateLinkScopedResourcesClient) Get(ctx context.Context, resourceGroupName string, scopeName string, name string, options *PrivateLinkScopedResourcesGetOptions) (PrivateLinkScopedResourcesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, scopeName, name, options)
	if err != nil {
		return PrivateLinkScopedResourcesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PrivateLinkScopedResourcesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateLinkScopedResourcesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkScopedResourcesClient) getCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, name string, options *PrivateLinkScopedResourcesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/privateLinkScopes/{scopeName}/scopedResources/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-17-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateLinkScopedResourcesClient) getHandleResponse(resp *http.Response) (PrivateLinkScopedResourcesGetResponse, error) {
	result := PrivateLinkScopedResourcesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScopedResource); err != nil {
		return PrivateLinkScopedResourcesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *PrivateLinkScopedResourcesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByPrivateLinkScope - Gets all private endpoint connections on a private link scope.
// If the operation fails it returns a generic error.
func (client *PrivateLinkScopedResourcesClient) ListByPrivateLinkScope(resourceGroupName string, scopeName string, options *PrivateLinkScopedResourcesListByPrivateLinkScopeOptions) *PrivateLinkScopedResourcesListByPrivateLinkScopePager {
	return &PrivateLinkScopedResourcesListByPrivateLinkScopePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByPrivateLinkScopeCreateRequest(ctx, resourceGroupName, scopeName, options)
		},
		advancer: func(ctx context.Context, resp PrivateLinkScopedResourcesListByPrivateLinkScopeResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ScopedResourceListResult.NextLink)
		},
	}
}

// listByPrivateLinkScopeCreateRequest creates the ListByPrivateLinkScope request.
func (client *PrivateLinkScopedResourcesClient) listByPrivateLinkScopeCreateRequest(ctx context.Context, resourceGroupName string, scopeName string, options *PrivateLinkScopedResourcesListByPrivateLinkScopeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/privateLinkScopes/{scopeName}/scopedResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if scopeName == "" {
		return nil, errors.New("parameter scopeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopeName}", url.PathEscape(scopeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-17-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByPrivateLinkScopeHandleResponse handles the ListByPrivateLinkScope response.
func (client *PrivateLinkScopedResourcesClient) listByPrivateLinkScopeHandleResponse(resp *http.Response) (PrivateLinkScopedResourcesListByPrivateLinkScopeResponse, error) {
	result := PrivateLinkScopedResourcesListByPrivateLinkScopeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScopedResourceListResult); err != nil {
		return PrivateLinkScopedResourcesListByPrivateLinkScopeResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByPrivateLinkScopeHandleError handles the ListByPrivateLinkScope error response.
func (client *PrivateLinkScopedResourcesClient) listByPrivateLinkScopeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
