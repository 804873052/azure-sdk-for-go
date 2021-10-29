//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

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
	"strconv"
	"strings"
)

// WorkflowRunsClient contains the methods for the WorkflowRuns group.
// Don't use this type directly, use NewWorkflowRunsClient() instead.
type WorkflowRunsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewWorkflowRunsClient creates a new instance of WorkflowRunsClient with the specified values.
func NewWorkflowRunsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WorkflowRunsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &WorkflowRunsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Cancel - Cancels a workflow run.
// If the operation fails it returns the *ErrorResponse error type.
func (client *WorkflowRunsClient) Cancel(ctx context.Context, resourceGroupName string, workflowName string, runName string, options *WorkflowRunsCancelOptions) (WorkflowRunsCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, workflowName, runName, options)
	if err != nil {
		return WorkflowRunsCancelResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowRunsCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowRunsCancelResponse{}, client.cancelHandleError(resp)
	}
	return WorkflowRunsCancelResponse{RawResponse: resp}, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *WorkflowRunsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, runName string, options *WorkflowRunsCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}/cancel"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if runName == "" {
		return nil, errors.New("parameter runName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runName}", url.PathEscape(runName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// cancelHandleError handles the Cancel error response.
func (client *WorkflowRunsClient) cancelHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets a workflow run.
// If the operation fails it returns the *ErrorResponse error type.
func (client *WorkflowRunsClient) Get(ctx context.Context, resourceGroupName string, workflowName string, runName string, options *WorkflowRunsGetOptions) (WorkflowRunsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workflowName, runName, options)
	if err != nil {
		return WorkflowRunsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkflowRunsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkflowRunsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkflowRunsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, runName string, options *WorkflowRunsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs/{runName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	if runName == "" {
		return nil, errors.New("parameter runName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runName}", url.PathEscape(runName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkflowRunsClient) getHandleResponse(resp *http.Response) (WorkflowRunsGetResponse, error) {
	result := WorkflowRunsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowRun); err != nil {
		return WorkflowRunsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *WorkflowRunsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets a list of workflow runs.
// If the operation fails it returns the *ErrorResponse error type.
func (client *WorkflowRunsClient) List(resourceGroupName string, workflowName string, options *WorkflowRunsListOptions) *WorkflowRunsListPager {
	return &WorkflowRunsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, workflowName, options)
		},
		advancer: func(ctx context.Context, resp WorkflowRunsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WorkflowRunListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *WorkflowRunsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workflowName string, options *WorkflowRunsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Logic/workflows/{workflowName}/runs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workflowName == "" {
		return nil, errors.New("parameter workflowName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workflowName}", url.PathEscape(workflowName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkflowRunsClient) listHandleResponse(resp *http.Response) (WorkflowRunsListResponse, error) {
	result := WorkflowRunsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkflowRunListResult); err != nil {
		return WorkflowRunsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *WorkflowRunsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
