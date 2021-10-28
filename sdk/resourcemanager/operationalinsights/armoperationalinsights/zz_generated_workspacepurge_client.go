//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

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

// WorkspacePurgeClient contains the methods for the WorkspacePurge group.
// Don't use this type directly, use NewWorkspacePurgeClient() instead.
type WorkspacePurgeClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewWorkspacePurgeClient creates a new instance of WorkspacePurgeClient with the specified values.
func NewWorkspacePurgeClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WorkspacePurgeClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &WorkspacePurgeClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// GetPurgeStatus - Gets status of an ongoing purge operation.
// If the operation fails it returns a generic error.
func (client *WorkspacePurgeClient) GetPurgeStatus(ctx context.Context, resourceGroupName string, workspaceName string, purgeID string, options *WorkspacePurgeGetPurgeStatusOptions) (WorkspacePurgeGetPurgeStatusResponse, error) {
	req, err := client.getPurgeStatusCreateRequest(ctx, resourceGroupName, workspaceName, purgeID, options)
	if err != nil {
		return WorkspacePurgeGetPurgeStatusResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspacePurgeGetPurgeStatusResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspacePurgeGetPurgeStatusResponse{}, client.getPurgeStatusHandleError(resp)
	}
	return client.getPurgeStatusHandleResponse(resp)
}

// getPurgeStatusCreateRequest creates the GetPurgeStatus request.
func (client *WorkspacePurgeClient) getPurgeStatusCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, purgeID string, options *WorkspacePurgeGetPurgeStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/operations/{purgeId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if purgeID == "" {
		return nil, errors.New("parameter purgeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{purgeId}", url.PathEscape(purgeID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getPurgeStatusHandleResponse handles the GetPurgeStatus response.
func (client *WorkspacePurgeClient) getPurgeStatusHandleResponse(resp *http.Response) (WorkspacePurgeGetPurgeStatusResponse, error) {
	result := WorkspacePurgeGetPurgeStatusResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspacePurgeStatusResponse); err != nil {
		return WorkspacePurgeGetPurgeStatusResponse{}, err
	}
	return result, nil
}

// getPurgeStatusHandleError handles the GetPurgeStatus error response.
func (client *WorkspacePurgeClient) getPurgeStatusHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Purge - Purges data in an Log Analytics workspace by a set of user-defined filters.
// In order to manage system resources, purge requests are throttled at 50 requests per hour. You should batch the execution of purge requests by sending
// a single command whose predicate includes all
// user identities that require purging. Use the in operator to specify multiple identities. You should run the query prior to using for a purge request
// to verify that the results are expected.
// If the operation fails it returns a generic error.
func (client *WorkspacePurgeClient) Purge(ctx context.Context, resourceGroupName string, workspaceName string, body WorkspacePurgeBody, options *WorkspacePurgePurgeOptions) (WorkspacePurgePurgeResponse, error) {
	req, err := client.purgeCreateRequest(ctx, resourceGroupName, workspaceName, body, options)
	if err != nil {
		return WorkspacePurgePurgeResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspacePurgePurgeResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return WorkspacePurgePurgeResponse{}, client.purgeHandleError(resp)
	}
	return client.purgeHandleResponse(resp)
}

// purgeCreateRequest creates the Purge request.
func (client *WorkspacePurgeClient) purgeCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, body WorkspacePurgeBody, options *WorkspacePurgePurgeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/purge"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// purgeHandleResponse handles the Purge response.
func (client *WorkspacePurgeClient) purgeHandleResponse(resp *http.Response) (WorkspacePurgePurgeResponse, error) {
	result := WorkspacePurgePurgeResponse{RawResponse: resp}
	if val := resp.Header.Get("x-ms-status-location"); val != "" {
		result.XMSStatusLocation = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspacePurgeResponse); err != nil {
		return WorkspacePurgePurgeResponse{}, err
	}
	return result, nil
}

// purgeHandleError handles the Purge error response.
func (client *WorkspacePurgeClient) purgeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
