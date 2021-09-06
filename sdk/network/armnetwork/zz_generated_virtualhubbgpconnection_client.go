//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// VirtualHubBgpConnectionClient contains the methods for the VirtualHubBgpConnection group.
// Don't use this type directly, use NewVirtualHubBgpConnectionClient() instead.
type VirtualHubBgpConnectionClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewVirtualHubBgpConnectionClient creates a new instance of VirtualHubBgpConnectionClient with the specified values.
func NewVirtualHubBgpConnectionClient(con *arm.Connection, subscriptionID string) *VirtualHubBgpConnectionClient {
	return &VirtualHubBgpConnectionClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates a VirtualHubBgpConnection resource if it doesn't exist else updates the existing VirtualHubBgpConnection.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualHubBgpConnectionClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, parameters BgpConnection, options *VirtualHubBgpConnectionBeginCreateOrUpdateOptions) (VirtualHubBgpConnectionCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualHubName, connectionName, parameters, options)
	if err != nil {
		return VirtualHubBgpConnectionCreateOrUpdatePollerResponse{}, err
	}
	result := VirtualHubBgpConnectionCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualHubBgpConnectionClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return VirtualHubBgpConnectionCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VirtualHubBgpConnectionCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a VirtualHubBgpConnection resource if it doesn't exist else updates the existing VirtualHubBgpConnection.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualHubBgpConnectionClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, parameters BgpConnection, options *VirtualHubBgpConnectionBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualHubName, connectionName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VirtualHubBgpConnectionClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, parameters BgpConnection, options *VirtualHubBgpConnectionBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/bgpConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VirtualHubBgpConnectionClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes a VirtualHubBgpConnection.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualHubBgpConnectionClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *VirtualHubBgpConnectionBeginDeleteOptions) (VirtualHubBgpConnectionDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, virtualHubName, connectionName, options)
	if err != nil {
		return VirtualHubBgpConnectionDeletePollerResponse{}, err
	}
	result := VirtualHubBgpConnectionDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VirtualHubBgpConnectionClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return VirtualHubBgpConnectionDeletePollerResponse{}, err
	}
	result.Poller = &VirtualHubBgpConnectionDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a VirtualHubBgpConnection.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualHubBgpConnectionClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *VirtualHubBgpConnectionBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualHubName, connectionName, options)
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
func (client *VirtualHubBgpConnectionClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *VirtualHubBgpConnectionBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/bgpConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *VirtualHubBgpConnectionClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Retrieves the details of a Virtual Hub Bgp Connection.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualHubBgpConnectionClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *VirtualHubBgpConnectionGetOptions) (VirtualHubBgpConnectionGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualHubName, connectionName, options)
	if err != nil {
		return VirtualHubBgpConnectionGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualHubBgpConnectionGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualHubBgpConnectionGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualHubBgpConnectionClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *VirtualHubBgpConnectionGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/bgpConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualHubBgpConnectionClient) getHandleResponse(resp *http.Response) (VirtualHubBgpConnectionGetResponse, error) {
	result := VirtualHubBgpConnectionGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.BgpConnection); err != nil {
		return VirtualHubBgpConnectionGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VirtualHubBgpConnectionClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
