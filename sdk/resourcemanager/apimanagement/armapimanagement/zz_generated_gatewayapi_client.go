//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// GatewayAPIClient contains the methods for the GatewayAPI group.
// Don't use this type directly, use NewGatewayAPIClient() instead.
type GatewayAPIClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewGatewayAPIClient creates a new instance of GatewayAPIClient with the specified values.
func NewGatewayAPIClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *GatewayAPIClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &GatewayAPIClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Adds an API to the specified Gateway.
// If the operation fails it returns the *ErrorResponse error type.
func (client *GatewayAPIClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, apiID string, options *GatewayAPICreateOrUpdateOptions) (GatewayAPICreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, gatewayID, apiID, options)
	if err != nil {
		return GatewayAPICreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GatewayAPICreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return GatewayAPICreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *GatewayAPIClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, apiID string, options *GatewayAPICreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/apis/{apiId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayID == "" {
		return nil, errors.New("parameter gatewayID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayId}", url.PathEscape(gatewayID))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *GatewayAPIClient) createOrUpdateHandleResponse(resp *http.Response) (GatewayAPICreateOrUpdateResponse, error) {
	result := GatewayAPICreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.APIContract); err != nil {
		return GatewayAPICreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *GatewayAPIClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes the specified API from the specified Gateway.
// If the operation fails it returns the *ErrorResponse error type.
func (client *GatewayAPIClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, apiID string, options *GatewayAPIDeleteOptions) (GatewayAPIDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, gatewayID, apiID, options)
	if err != nil {
		return GatewayAPIDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GatewayAPIDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return GatewayAPIDeleteResponse{}, client.deleteHandleError(resp)
	}
	return GatewayAPIDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *GatewayAPIClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, apiID string, options *GatewayAPIDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/apis/{apiId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayID == "" {
		return nil, errors.New("parameter gatewayID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayId}", url.PathEscape(gatewayID))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *GatewayAPIClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetEntityTag - Checks that API entity specified by identifier is associated with the Gateway entity.
// If the operation fails it returns the *ErrorResponse error type.
func (client *GatewayAPIClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, apiID string, options *GatewayAPIGetEntityTagOptions) (GatewayAPIGetEntityTagResponse, error) {
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, gatewayID, apiID, options)
	if err != nil {
		return GatewayAPIGetEntityTagResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GatewayAPIGetEntityTagResponse{}, err
	}
	return client.getEntityTagHandleResponse(resp)
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *GatewayAPIClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, apiID string, options *GatewayAPIGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/apis/{apiId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayID == "" {
		return nil, errors.New("parameter gatewayID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayId}", url.PathEscape(gatewayID))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *GatewayAPIClient) getEntityTagHandleResponse(resp *http.Response) (GatewayAPIGetEntityTagResponse, error) {
	result := GatewayAPIGetEntityTagResponse{RawResponse: resp}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// ListByService - Lists a collection of the APIs associated with a gateway.
// If the operation fails it returns the *ErrorResponse error type.
func (client *GatewayAPIClient) ListByService(resourceGroupName string, serviceName string, gatewayID string, options *GatewayAPIListByServiceOptions) *GatewayAPIListByServicePager {
	return &GatewayAPIListByServicePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, gatewayID, options)
		},
		advancer: func(ctx context.Context, resp GatewayAPIListByServiceResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.APICollection.NextLink)
		},
	}
}

// listByServiceCreateRequest creates the ListByService request.
func (client *GatewayAPIClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, gatewayID string, options *GatewayAPIListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/gateways/{gatewayId}/apis"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if gatewayID == "" {
		return nil, errors.New("parameter gatewayID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayId}", url.PathEscape(gatewayID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *GatewayAPIClient) listByServiceHandleResponse(resp *http.Response) (GatewayAPIListByServiceResponse, error) {
	result := GatewayAPIListByServiceResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.APICollection); err != nil {
		return GatewayAPIListByServiceResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByServiceHandleError handles the ListByService error response.
func (client *GatewayAPIClient) listByServiceHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
