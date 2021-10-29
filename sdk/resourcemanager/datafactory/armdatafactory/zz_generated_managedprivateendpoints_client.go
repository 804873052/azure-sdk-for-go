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

// ManagedPrivateEndpointsClient contains the methods for the ManagedPrivateEndpoints group.
// Don't use this type directly, use NewManagedPrivateEndpointsClient() instead.
type ManagedPrivateEndpointsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewManagedPrivateEndpointsClient creates a new instance of ManagedPrivateEndpointsClient with the specified values.
func NewManagedPrivateEndpointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ManagedPrivateEndpointsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ManagedPrivateEndpointsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Creates or updates a managed private endpoint.
// If the operation fails it returns the *CloudError error type.
func (client *ManagedPrivateEndpointsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string, managedPrivateEndpoint ManagedPrivateEndpointResource, options *ManagedPrivateEndpointsCreateOrUpdateOptions) (ManagedPrivateEndpointsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, factoryName, managedVirtualNetworkName, managedPrivateEndpointName, managedPrivateEndpoint, options)
	if err != nil {
		return ManagedPrivateEndpointsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedPrivateEndpointsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedPrivateEndpointsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ManagedPrivateEndpointsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string, managedPrivateEndpoint ManagedPrivateEndpointResource, options *ManagedPrivateEndpointsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/managedVirtualNetworks/{managedVirtualNetworkName}/managedPrivateEndpoints/{managedPrivateEndpointName}"
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
	if managedVirtualNetworkName == "" {
		return nil, errors.New("parameter managedVirtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedVirtualNetworkName}", url.PathEscape(managedVirtualNetworkName))
	if managedPrivateEndpointName == "" {
		return nil, errors.New("parameter managedPrivateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedPrivateEndpointName}", url.PathEscape(managedPrivateEndpointName))
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
	return req, runtime.MarshalAsJSON(req, managedPrivateEndpoint)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ManagedPrivateEndpointsClient) createOrUpdateHandleResponse(resp *http.Response) (ManagedPrivateEndpointsCreateOrUpdateResponse, error) {
	result := ManagedPrivateEndpointsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedPrivateEndpointResource); err != nil {
		return ManagedPrivateEndpointsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ManagedPrivateEndpointsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Delete - Deletes a managed private endpoint.
// If the operation fails it returns the *CloudError error type.
func (client *ManagedPrivateEndpointsClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string, options *ManagedPrivateEndpointsDeleteOptions) (ManagedPrivateEndpointsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, factoryName, managedVirtualNetworkName, managedPrivateEndpointName, options)
	if err != nil {
		return ManagedPrivateEndpointsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedPrivateEndpointsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ManagedPrivateEndpointsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return ManagedPrivateEndpointsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ManagedPrivateEndpointsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string, options *ManagedPrivateEndpointsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/managedVirtualNetworks/{managedVirtualNetworkName}/managedPrivateEndpoints/{managedPrivateEndpointName}"
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
	if managedVirtualNetworkName == "" {
		return nil, errors.New("parameter managedVirtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedVirtualNetworkName}", url.PathEscape(managedVirtualNetworkName))
	if managedPrivateEndpointName == "" {
		return nil, errors.New("parameter managedPrivateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedPrivateEndpointName}", url.PathEscape(managedPrivateEndpointName))
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
func (client *ManagedPrivateEndpointsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Gets a managed private endpoint.
// If the operation fails it returns the *CloudError error type.
func (client *ManagedPrivateEndpointsClient) Get(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string, options *ManagedPrivateEndpointsGetOptions) (ManagedPrivateEndpointsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, factoryName, managedVirtualNetworkName, managedPrivateEndpointName, options)
	if err != nil {
		return ManagedPrivateEndpointsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ManagedPrivateEndpointsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ManagedPrivateEndpointsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ManagedPrivateEndpointsClient) getCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, managedPrivateEndpointName string, options *ManagedPrivateEndpointsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/managedVirtualNetworks/{managedVirtualNetworkName}/managedPrivateEndpoints/{managedPrivateEndpointName}"
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
	if managedVirtualNetworkName == "" {
		return nil, errors.New("parameter managedVirtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedVirtualNetworkName}", url.PathEscape(managedVirtualNetworkName))
	if managedPrivateEndpointName == "" {
		return nil, errors.New("parameter managedPrivateEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedPrivateEndpointName}", url.PathEscape(managedPrivateEndpointName))
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
func (client *ManagedPrivateEndpointsClient) getHandleResponse(resp *http.Response) (ManagedPrivateEndpointsGetResponse, error) {
	result := ManagedPrivateEndpointsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedPrivateEndpointResource); err != nil {
		return ManagedPrivateEndpointsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ManagedPrivateEndpointsClient) getHandleError(resp *http.Response) error {
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

// ListByFactory - Lists managed private endpoints.
// If the operation fails it returns the *CloudError error type.
func (client *ManagedPrivateEndpointsClient) ListByFactory(resourceGroupName string, factoryName string, managedVirtualNetworkName string, options *ManagedPrivateEndpointsListByFactoryOptions) *ManagedPrivateEndpointsListByFactoryPager {
	return &ManagedPrivateEndpointsListByFactoryPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByFactoryCreateRequest(ctx, resourceGroupName, factoryName, managedVirtualNetworkName, options)
		},
		advancer: func(ctx context.Context, resp ManagedPrivateEndpointsListByFactoryResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ManagedPrivateEndpointListResponse.NextLink)
		},
	}
}

// listByFactoryCreateRequest creates the ListByFactory request.
func (client *ManagedPrivateEndpointsClient) listByFactoryCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, managedVirtualNetworkName string, options *ManagedPrivateEndpointsListByFactoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/managedVirtualNetworks/{managedVirtualNetworkName}/managedPrivateEndpoints"
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
	if managedVirtualNetworkName == "" {
		return nil, errors.New("parameter managedVirtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedVirtualNetworkName}", url.PathEscape(managedVirtualNetworkName))
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
func (client *ManagedPrivateEndpointsClient) listByFactoryHandleResponse(resp *http.Response) (ManagedPrivateEndpointsListByFactoryResponse, error) {
	result := ManagedPrivateEndpointsListByFactoryResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedPrivateEndpointListResponse); err != nil {
		return ManagedPrivateEndpointsListByFactoryResponse{}, err
	}
	return result, nil
}

// listByFactoryHandleError handles the ListByFactory error response.
func (client *ManagedPrivateEndpointsClient) listByFactoryHandleError(resp *http.Response) error {
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
