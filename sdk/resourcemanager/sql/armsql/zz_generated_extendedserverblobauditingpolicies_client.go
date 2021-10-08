//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ExtendedServerBlobAuditingPoliciesClient contains the methods for the ExtendedServerBlobAuditingPolicies group.
// Don't use this type directly, use NewExtendedServerBlobAuditingPoliciesClient() instead.
type ExtendedServerBlobAuditingPoliciesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewExtendedServerBlobAuditingPoliciesClient creates a new instance of ExtendedServerBlobAuditingPoliciesClient with the specified values.
func NewExtendedServerBlobAuditingPoliciesClient(con *arm.Connection, subscriptionID string) *ExtendedServerBlobAuditingPoliciesClient {
	return &ExtendedServerBlobAuditingPoliciesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates an extended server's blob auditing policy.
// If the operation fails it returns a generic error.
func (client *ExtendedServerBlobAuditingPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, blobAuditingPolicyName Enum5, parameters ExtendedServerBlobAuditingPolicy, options *ExtendedServerBlobAuditingPoliciesBeginCreateOrUpdateOptions) (ExtendedServerBlobAuditingPoliciesCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, blobAuditingPolicyName, parameters, options)
	if err != nil {
		return ExtendedServerBlobAuditingPoliciesCreateOrUpdatePollerResponse{}, err
	}
	result := ExtendedServerBlobAuditingPoliciesCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExtendedServerBlobAuditingPoliciesClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return ExtendedServerBlobAuditingPoliciesCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ExtendedServerBlobAuditingPoliciesCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates an extended server's blob auditing policy.
// If the operation fails it returns a generic error.
func (client *ExtendedServerBlobAuditingPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, blobAuditingPolicyName Enum5, parameters ExtendedServerBlobAuditingPolicy, options *ExtendedServerBlobAuditingPoliciesBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, blobAuditingPolicyName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExtendedServerBlobAuditingPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, blobAuditingPolicyName Enum5, parameters ExtendedServerBlobAuditingPolicy, options *ExtendedServerBlobAuditingPoliciesBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/extendedAuditingSettings/{blobAuditingPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if blobAuditingPolicyName == "" {
		return nil, errors.New("parameter blobAuditingPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blobAuditingPolicyName}", url.PathEscape(string(blobAuditingPolicyName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ExtendedServerBlobAuditingPoliciesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets an extended server's blob auditing policy.
// If the operation fails it returns a generic error.
func (client *ExtendedServerBlobAuditingPoliciesClient) Get(ctx context.Context, resourceGroupName string, serverName string, blobAuditingPolicyName Enum5, options *ExtendedServerBlobAuditingPoliciesGetOptions) (ExtendedServerBlobAuditingPoliciesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, blobAuditingPolicyName, options)
	if err != nil {
		return ExtendedServerBlobAuditingPoliciesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExtendedServerBlobAuditingPoliciesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExtendedServerBlobAuditingPoliciesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExtendedServerBlobAuditingPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, blobAuditingPolicyName Enum5, options *ExtendedServerBlobAuditingPoliciesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/extendedAuditingSettings/{blobAuditingPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if blobAuditingPolicyName == "" {
		return nil, errors.New("parameter blobAuditingPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blobAuditingPolicyName}", url.PathEscape(string(blobAuditingPolicyName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExtendedServerBlobAuditingPoliciesClient) getHandleResponse(resp *http.Response) (ExtendedServerBlobAuditingPoliciesGetResponse, error) {
	result := ExtendedServerBlobAuditingPoliciesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtendedServerBlobAuditingPolicy); err != nil {
		return ExtendedServerBlobAuditingPoliciesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ExtendedServerBlobAuditingPoliciesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByServer - Lists extended auditing settings of a server.
// If the operation fails it returns a generic error.
func (client *ExtendedServerBlobAuditingPoliciesClient) ListByServer(resourceGroupName string, serverName string, options *ExtendedServerBlobAuditingPoliciesListByServerOptions) *ExtendedServerBlobAuditingPoliciesListByServerPager {
	return &ExtendedServerBlobAuditingPoliciesListByServerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
		},
		advancer: func(ctx context.Context, resp ExtendedServerBlobAuditingPoliciesListByServerResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ExtendedServerBlobAuditingPolicyListResult.NextLink)
		},
	}
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ExtendedServerBlobAuditingPoliciesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ExtendedServerBlobAuditingPoliciesListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/extendedAuditingSettings"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ExtendedServerBlobAuditingPoliciesClient) listByServerHandleResponse(resp *http.Response) (ExtendedServerBlobAuditingPoliciesListByServerResponse, error) {
	result := ExtendedServerBlobAuditingPoliciesListByServerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtendedServerBlobAuditingPolicyListResult); err != nil {
		return ExtendedServerBlobAuditingPoliciesListByServerResponse{}, err
	}
	return result, nil
}

// listByServerHandleError handles the ListByServer error response.
func (client *ExtendedServerBlobAuditingPoliciesClient) listByServerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
