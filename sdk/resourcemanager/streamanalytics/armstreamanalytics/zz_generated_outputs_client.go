//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstreamanalytics

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// OutputsClient contains the methods for the Outputs group.
// Don't use this type directly, use NewOutputsClient() instead.
type OutputsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewOutputsClient creates a new instance of OutputsClient with the specified values.
func NewOutputsClient(con *arm.Connection, subscriptionID string) *OutputsClient {
	return &OutputsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrReplace - Creates an output or replaces an already existing output under an existing streaming job.
// If the operation fails it returns the *Error error type.
func (client *OutputsClient) CreateOrReplace(ctx context.Context, resourceGroupName string, jobName string, outputName string, output Output, options *OutputsCreateOrReplaceOptions) (OutputsCreateOrReplaceResponse, error) {
	req, err := client.createOrReplaceCreateRequest(ctx, resourceGroupName, jobName, outputName, output, options)
	if err != nil {
		return OutputsCreateOrReplaceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OutputsCreateOrReplaceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return OutputsCreateOrReplaceResponse{}, client.createOrReplaceHandleError(resp)
	}
	return client.createOrReplaceHandleResponse(resp)
}

// createOrReplaceCreateRequest creates the CreateOrReplace request.
func (client *OutputsClient) createOrReplaceCreateRequest(ctx context.Context, resourceGroupName string, jobName string, outputName string, output Output, options *OutputsCreateOrReplaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if outputName == "" {
		return nil, errors.New("parameter outputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{outputName}", url.PathEscape(outputName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, output)
}

// createOrReplaceHandleResponse handles the CreateOrReplace response.
func (client *OutputsClient) createOrReplaceHandleResponse(resp *http.Response) (OutputsCreateOrReplaceResponse, error) {
	result := OutputsCreateOrReplaceResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Output); err != nil {
		return OutputsCreateOrReplaceResponse{}, err
	}
	return result, nil
}

// createOrReplaceHandleError handles the CreateOrReplace error response.
func (client *OutputsClient) createOrReplaceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes an output from the streaming job.
// If the operation fails it returns the *Error error type.
func (client *OutputsClient) Delete(ctx context.Context, resourceGroupName string, jobName string, outputName string, options *OutputsDeleteOptions) (OutputsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, jobName, outputName, options)
	if err != nil {
		return OutputsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OutputsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return OutputsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return OutputsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *OutputsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, jobName string, outputName string, options *OutputsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if outputName == "" {
		return nil, errors.New("parameter outputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{outputName}", url.PathEscape(outputName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *OutputsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets details about the specified output.
// If the operation fails it returns the *Error error type.
func (client *OutputsClient) Get(ctx context.Context, resourceGroupName string, jobName string, outputName string, options *OutputsGetOptions) (OutputsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, jobName, outputName, options)
	if err != nil {
		return OutputsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OutputsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OutputsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OutputsClient) getCreateRequest(ctx context.Context, resourceGroupName string, jobName string, outputName string, options *OutputsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if outputName == "" {
		return nil, errors.New("parameter outputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{outputName}", url.PathEscape(outputName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OutputsClient) getHandleResponse(resp *http.Response) (OutputsGetResponse, error) {
	result := OutputsGetResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Output); err != nil {
		return OutputsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *OutputsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByStreamingJob - Lists all of the outputs under the specified streaming job.
// If the operation fails it returns the *Error error type.
func (client *OutputsClient) ListByStreamingJob(resourceGroupName string, jobName string, options *OutputsListByStreamingJobOptions) *OutputsListByStreamingJobPager {
	return &OutputsListByStreamingJobPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByStreamingJobCreateRequest(ctx, resourceGroupName, jobName, options)
		},
		advancer: func(ctx context.Context, resp OutputsListByStreamingJobResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.OutputListResult.NextLink)
		},
	}
}

// listByStreamingJobCreateRequest creates the ListByStreamingJob request.
func (client *OutputsClient) listByStreamingJobCreateRequest(ctx context.Context, resourceGroupName string, jobName string, options *OutputsListByStreamingJobOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	reqQP.Set("api-version", "2017-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByStreamingJobHandleResponse handles the ListByStreamingJob response.
func (client *OutputsClient) listByStreamingJobHandleResponse(resp *http.Response) (OutputsListByStreamingJobResponse, error) {
	result := OutputsListByStreamingJobResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OutputListResult); err != nil {
		return OutputsListByStreamingJobResponse{}, err
	}
	return result, nil
}

// listByStreamingJobHandleError handles the ListByStreamingJob error response.
func (client *OutputsClient) listByStreamingJobHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginTest - Tests whether an output’s datasource is reachable and usable by the Azure Stream Analytics service.
// If the operation fails it returns the *Error error type.
func (client *OutputsClient) BeginTest(ctx context.Context, resourceGroupName string, jobName string, outputName string, options *OutputsBeginTestOptions) (OutputsTestPollerResponse, error) {
	resp, err := client.test(ctx, resourceGroupName, jobName, outputName, options)
	if err != nil {
		return OutputsTestPollerResponse{}, err
	}
	result := OutputsTestPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("OutputsClient.Test", "", resp, client.pl, client.testHandleError)
	if err != nil {
		return OutputsTestPollerResponse{}, err
	}
	result.Poller = &OutputsTestPoller{
		pt: pt,
	}
	return result, nil
}

// Test - Tests whether an output’s datasource is reachable and usable by the Azure Stream Analytics service.
// If the operation fails it returns the *Error error type.
func (client *OutputsClient) test(ctx context.Context, resourceGroupName string, jobName string, outputName string, options *OutputsBeginTestOptions) (*http.Response, error) {
	req, err := client.testCreateRequest(ctx, resourceGroupName, jobName, outputName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.testHandleError(resp)
	}
	return resp, nil
}

// testCreateRequest creates the Test request.
func (client *OutputsClient) testCreateRequest(ctx context.Context, resourceGroupName string, jobName string, outputName string, options *OutputsBeginTestOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}/test"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if outputName == "" {
		return nil, errors.New("parameter outputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{outputName}", url.PathEscape(outputName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Output != nil {
		return req, runtime.MarshalAsJSON(req, *options.Output)
	}
	return req, nil
}

// testHandleError handles the Test error response.
func (client *OutputsClient) testHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Updates an existing output under an existing streaming job. This can be used to partially update (ie. update one or two properties) an output
// without affecting the rest the job or output definition.
// If the operation fails it returns the *Error error type.
func (client *OutputsClient) Update(ctx context.Context, resourceGroupName string, jobName string, outputName string, output Output, options *OutputsUpdateOptions) (OutputsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, jobName, outputName, output, options)
	if err != nil {
		return OutputsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OutputsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OutputsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *OutputsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, jobName string, outputName string, output Output, options *OutputsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StreamAnalytics/streamingjobs/{jobName}/outputs/{outputName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	if outputName == "" {
		return nil, errors.New("parameter outputName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{outputName}", url.PathEscape(outputName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, output)
}

// updateHandleResponse handles the Update response.
func (client *OutputsClient) updateHandleResponse(resp *http.Response) (OutputsUpdateResponse, error) {
	result := OutputsUpdateResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Output); err != nil {
		return OutputsUpdateResponse{}, err
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *OutputsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := Error{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
