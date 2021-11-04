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

// ServerConnectionPoliciesClient contains the methods for the ServerConnectionPolicies group.
// Don't use this type directly, use NewServerConnectionPoliciesClient() instead.
type ServerConnectionPoliciesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewServerConnectionPoliciesClient creates a new instance of ServerConnectionPoliciesClient with the specified values.
func NewServerConnectionPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ServerConnectionPoliciesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ServerConnectionPoliciesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Creates or updates the server's connection policy.
// If the operation fails it returns a generic error.
func (client *ServerConnectionPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, connectionPolicyName ConnectionPolicyName, parameters ServerConnectionPolicy, options *ServerConnectionPoliciesCreateOrUpdateOptions) (ServerConnectionPoliciesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, connectionPolicyName, parameters, options)
	if err != nil {
		return ServerConnectionPoliciesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerConnectionPoliciesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ServerConnectionPoliciesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerConnectionPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, connectionPolicyName ConnectionPolicyName, parameters ServerConnectionPolicy, options *ServerConnectionPoliciesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/connectionPolicies/{connectionPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if connectionPolicyName == "" {
		return nil, errors.New("parameter connectionPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionPolicyName}", url.PathEscape(string(connectionPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ServerConnectionPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (ServerConnectionPoliciesCreateOrUpdateResponse, error) {
	result := ServerConnectionPoliciesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerConnectionPolicy); err != nil {
		return ServerConnectionPoliciesCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ServerConnectionPoliciesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Gets the server's secure connection policy.
// If the operation fails it returns a generic error.
func (client *ServerConnectionPoliciesClient) Get(ctx context.Context, resourceGroupName string, serverName string, connectionPolicyName ConnectionPolicyName, options *ServerConnectionPoliciesGetOptions) (ServerConnectionPoliciesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, connectionPolicyName, options)
	if err != nil {
		return ServerConnectionPoliciesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerConnectionPoliciesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerConnectionPoliciesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerConnectionPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, connectionPolicyName ConnectionPolicyName, options *ServerConnectionPoliciesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/connectionPolicies/{connectionPolicyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if connectionPolicyName == "" {
		return nil, errors.New("parameter connectionPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionPolicyName}", url.PathEscape(string(connectionPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerConnectionPoliciesClient) getHandleResponse(resp *http.Response) (ServerConnectionPoliciesGetResponse, error) {
	result := ServerConnectionPoliciesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerConnectionPolicy); err != nil {
		return ServerConnectionPoliciesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ServerConnectionPoliciesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
