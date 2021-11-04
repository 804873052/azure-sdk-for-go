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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ManagedInstanceOperationsClient contains the methods for the ManagedInstanceOperations group.
// Don't use this type directly, use NewManagedInstanceOperationsClient() instead.
type ManagedInstanceOperationsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewManagedInstanceOperationsClient creates a new instance of ManagedInstanceOperationsClient with the specified values.
func NewManagedInstanceOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ManagedInstanceOperationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ManagedInstanceOperationsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Cancel - Cancels the asynchronous operation on the managed instance.
// If the operation fails it returns a generic error.
func (client *ManagedInstanceOperationsClient) Cancel(ctx context.Context, resourceGroupName string, managedInstanceName string, operationID string, options *ManagedInstanceOperationsCancelOptions) (ManagedInstanceOperationsCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, managedInstanceName, operationID, options)
	if err != nil {
		return ManagedInstanceOperationsCancelResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedInstanceOperationsCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedInstanceOperationsCancelResponse{}, client.cancelHandleError(resp)
	}
	return ManagedInstanceOperationsCancelResponse{RawResponse: resp}, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *ManagedInstanceOperationsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, operationID string, options *ManagedInstanceOperationsCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/operations/{operationId}/cancel"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// cancelHandleError handles the Cancel error response.
func (client *ManagedInstanceOperationsClient) cancelHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets a management operation on a managed instance.
// If the operation fails it returns a generic error.
func (client *ManagedInstanceOperationsClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, operationID string, options *ManagedInstanceOperationsGetOptions) (ManagedInstanceOperationsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, operationID, options)
	if err != nil {
		return ManagedInstanceOperationsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedInstanceOperationsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedInstanceOperationsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedInstanceOperationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, operationID string, options *ManagedInstanceOperationsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/operations/{operationId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
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
func (client *ManagedInstanceOperationsClient) getHandleResponse(resp *http.Response) (ManagedInstanceOperationsGetResponse, error) {
	result := ManagedInstanceOperationsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceOperation); err != nil {
		return ManagedInstanceOperationsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ManagedInstanceOperationsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByManagedInstance - Gets a list of operations performed on the managed instance.
// If the operation fails it returns a generic error.
func (client *ManagedInstanceOperationsClient) ListByManagedInstance(resourceGroupName string, managedInstanceName string, options *ManagedInstanceOperationsListByManagedInstanceOptions) *ManagedInstanceOperationsListByManagedInstancePager {
	return &ManagedInstanceOperationsListByManagedInstancePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByManagedInstanceCreateRequest(ctx, resourceGroupName, managedInstanceName, options)
		},
		advancer: func(ctx context.Context, resp ManagedInstanceOperationsListByManagedInstanceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ManagedInstanceOperationListResult.NextLink)
		},
	}
}

// listByManagedInstanceCreateRequest creates the ListByManagedInstance request.
func (client *ManagedInstanceOperationsClient) listByManagedInstanceCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, options *ManagedInstanceOperationsListByManagedInstanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/operations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
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

// listByManagedInstanceHandleResponse handles the ListByManagedInstance response.
func (client *ManagedInstanceOperationsClient) listByManagedInstanceHandleResponse(resp *http.Response) (ManagedInstanceOperationsListByManagedInstanceResponse, error) {
	result := ManagedInstanceOperationsListByManagedInstanceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceOperationListResult); err != nil {
		return ManagedInstanceOperationsListByManagedInstanceResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByManagedInstanceHandleError handles the ListByManagedInstance error response.
func (client *ManagedInstanceOperationsClient) listByManagedInstanceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
