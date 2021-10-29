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

// DataFlowDebugSessionClient contains the methods for the DataFlowDebugSession group.
// Don't use this type directly, use NewDataFlowDebugSessionClient() instead.
type DataFlowDebugSessionClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewDataFlowDebugSessionClient creates a new instance of DataFlowDebugSessionClient with the specified values.
func NewDataFlowDebugSessionClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DataFlowDebugSessionClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &DataFlowDebugSessionClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// AddDataFlow - Add a data flow into debug session.
// If the operation fails it returns the *CloudError error type.
func (client *DataFlowDebugSessionClient) AddDataFlow(ctx context.Context, resourceGroupName string, factoryName string, request DataFlowDebugPackage, options *DataFlowDebugSessionAddDataFlowOptions) (DataFlowDebugSessionAddDataFlowResponse, error) {
	req, err := client.addDataFlowCreateRequest(ctx, resourceGroupName, factoryName, request, options)
	if err != nil {
		return DataFlowDebugSessionAddDataFlowResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataFlowDebugSessionAddDataFlowResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataFlowDebugSessionAddDataFlowResponse{}, client.addDataFlowHandleError(resp)
	}
	return client.addDataFlowHandleResponse(resp)
}

// addDataFlowCreateRequest creates the AddDataFlow request.
func (client *DataFlowDebugSessionClient) addDataFlowCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, request DataFlowDebugPackage, options *DataFlowDebugSessionAddDataFlowOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/addDataFlowToDebugSession"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// addDataFlowHandleResponse handles the AddDataFlow response.
func (client *DataFlowDebugSessionClient) addDataFlowHandleResponse(resp *http.Response) (DataFlowDebugSessionAddDataFlowResponse, error) {
	result := DataFlowDebugSessionAddDataFlowResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.AddDataFlowToDebugSessionResponse); err != nil {
		return DataFlowDebugSessionAddDataFlowResponse{}, err
	}
	return result, nil
}

// addDataFlowHandleError handles the AddDataFlow error response.
func (client *DataFlowDebugSessionClient) addDataFlowHandleError(resp *http.Response) error {
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

// BeginCreate - Creates a data flow debug session.
// If the operation fails it returns the *CloudError error type.
func (client *DataFlowDebugSessionClient) BeginCreate(ctx context.Context, resourceGroupName string, factoryName string, request CreateDataFlowDebugSessionRequest, options *DataFlowDebugSessionBeginCreateOptions) (DataFlowDebugSessionCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, factoryName, request, options)
	if err != nil {
		return DataFlowDebugSessionCreatePollerResponse{}, err
	}
	result := DataFlowDebugSessionCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DataFlowDebugSessionClient.Create", "", resp, client.pl, client.createHandleError)
	if err != nil {
		return DataFlowDebugSessionCreatePollerResponse{}, err
	}
	result.Poller = &DataFlowDebugSessionCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates a data flow debug session.
// If the operation fails it returns the *CloudError error type.
func (client *DataFlowDebugSessionClient) create(ctx context.Context, resourceGroupName string, factoryName string, request CreateDataFlowDebugSessionRequest, options *DataFlowDebugSessionBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, factoryName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *DataFlowDebugSessionClient) createCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, request CreateDataFlowDebugSessionRequest, options *DataFlowDebugSessionBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/createDataFlowDebugSession"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// createHandleError handles the Create error response.
func (client *DataFlowDebugSessionClient) createHandleError(resp *http.Response) error {
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

// Delete - Deletes a data flow debug session.
// If the operation fails it returns the *CloudError error type.
func (client *DataFlowDebugSessionClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, request DeleteDataFlowDebugSessionRequest, options *DataFlowDebugSessionDeleteOptions) (DataFlowDebugSessionDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, factoryName, request, options)
	if err != nil {
		return DataFlowDebugSessionDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DataFlowDebugSessionDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DataFlowDebugSessionDeleteResponse{}, client.deleteHandleError(resp)
	}
	return DataFlowDebugSessionDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DataFlowDebugSessionClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, request DeleteDataFlowDebugSessionRequest, options *DataFlowDebugSessionDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/deleteDataFlowDebugSession"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// deleteHandleError handles the Delete error response.
func (client *DataFlowDebugSessionClient) deleteHandleError(resp *http.Response) error {
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

// BeginExecuteCommand - Execute a data flow debug command.
// If the operation fails it returns the *CloudError error type.
func (client *DataFlowDebugSessionClient) BeginExecuteCommand(ctx context.Context, resourceGroupName string, factoryName string, request DataFlowDebugCommandRequest, options *DataFlowDebugSessionBeginExecuteCommandOptions) (DataFlowDebugSessionExecuteCommandPollerResponse, error) {
	resp, err := client.executeCommand(ctx, resourceGroupName, factoryName, request, options)
	if err != nil {
		return DataFlowDebugSessionExecuteCommandPollerResponse{}, err
	}
	result := DataFlowDebugSessionExecuteCommandPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DataFlowDebugSessionClient.ExecuteCommand", "", resp, client.pl, client.executeCommandHandleError)
	if err != nil {
		return DataFlowDebugSessionExecuteCommandPollerResponse{}, err
	}
	result.Poller = &DataFlowDebugSessionExecuteCommandPoller{
		pt: pt,
	}
	return result, nil
}

// ExecuteCommand - Execute a data flow debug command.
// If the operation fails it returns the *CloudError error type.
func (client *DataFlowDebugSessionClient) executeCommand(ctx context.Context, resourceGroupName string, factoryName string, request DataFlowDebugCommandRequest, options *DataFlowDebugSessionBeginExecuteCommandOptions) (*http.Response, error) {
	req, err := client.executeCommandCreateRequest(ctx, resourceGroupName, factoryName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.executeCommandHandleError(resp)
	}
	return resp, nil
}

// executeCommandCreateRequest creates the ExecuteCommand request.
func (client *DataFlowDebugSessionClient) executeCommandCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, request DataFlowDebugCommandRequest, options *DataFlowDebugSessionBeginExecuteCommandOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/executeDataFlowDebugCommand"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, request)
}

// executeCommandHandleError handles the ExecuteCommand error response.
func (client *DataFlowDebugSessionClient) executeCommandHandleError(resp *http.Response) error {
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

// QueryByFactory - Query all active data flow debug sessions.
// If the operation fails it returns the *CloudError error type.
func (client *DataFlowDebugSessionClient) QueryByFactory(resourceGroupName string, factoryName string, options *DataFlowDebugSessionQueryByFactoryOptions) *DataFlowDebugSessionQueryByFactoryPager {
	return &DataFlowDebugSessionQueryByFactoryPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.queryByFactoryCreateRequest(ctx, resourceGroupName, factoryName, options)
		},
		advancer: func(ctx context.Context, resp DataFlowDebugSessionQueryByFactoryResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.QueryDataFlowDebugSessionsResponse.NextLink)
		},
	}
}

// queryByFactoryCreateRequest creates the QueryByFactory request.
func (client *DataFlowDebugSessionClient) queryByFactoryCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, options *DataFlowDebugSessionQueryByFactoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/queryDataFlowDebugSessions"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// queryByFactoryHandleResponse handles the QueryByFactory response.
func (client *DataFlowDebugSessionClient) queryByFactoryHandleResponse(resp *http.Response) (DataFlowDebugSessionQueryByFactoryResponse, error) {
	result := DataFlowDebugSessionQueryByFactoryResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueryDataFlowDebugSessionsResponse); err != nil {
		return DataFlowDebugSessionQueryByFactoryResponse{}, err
	}
	return result, nil
}

// queryByFactoryHandleError handles the QueryByFactory error response.
func (client *DataFlowDebugSessionClient) queryByFactoryHandleError(resp *http.Response) error {
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
