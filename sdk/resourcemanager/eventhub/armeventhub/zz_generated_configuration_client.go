//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

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

// ConfigurationClient contains the methods for the Configuration group.
// Don't use this type directly, use NewConfigurationClient() instead.
type ConfigurationClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewConfigurationClient creates a new instance of ConfigurationClient with the specified values.
func NewConfigurationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ConfigurationClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ConfigurationClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Get all Event Hubs Cluster settings - a collection of key/value pairs which represent the quotas and settings imposed on the cluster.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConfigurationClient) Get(ctx context.Context, resourceGroupName string, clusterName string, options *ConfigurationGetOptions) (ConfigurationGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, options)
	if err != nil {
		return ConfigurationGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConfigurationGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConfigurationClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *ConfigurationGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/clusters/{clusterName}/quotaConfiguration/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConfigurationClient) getHandleResponse(resp *http.Response) (ConfigurationGetResponse, error) {
	result := ConfigurationGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterQuotaConfigurationProperties); err != nil {
		return ConfigurationGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ConfigurationClient) getHandleError(resp *http.Response) error {
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

// Patch - Replace all specified Event Hubs Cluster settings with those contained in the request body. Leaves the settings not specified in the request
// body unmodified.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConfigurationClient) Patch(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterQuotaConfigurationProperties, options *ConfigurationPatchOptions) (ConfigurationPatchResponse, error) {
	req, err := client.patchCreateRequest(ctx, resourceGroupName, clusterName, parameters, options)
	if err != nil {
		return ConfigurationPatchResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigurationPatchResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return ConfigurationPatchResponse{}, client.patchHandleError(resp)
	}
	return client.patchHandleResponse(resp)
}

// patchCreateRequest creates the Patch request.
func (client *ConfigurationClient) patchCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, parameters ClusterQuotaConfigurationProperties, options *ConfigurationPatchOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/clusters/{clusterName}/quotaConfiguration/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// patchHandleResponse handles the Patch response.
func (client *ConfigurationClient) patchHandleResponse(resp *http.Response) (ConfigurationPatchResponse, error) {
	result := ConfigurationPatchResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterQuotaConfigurationProperties); err != nil {
		return ConfigurationPatchResponse{}, err
	}
	return result, nil
}

// patchHandleError handles the Patch error response.
func (client *ConfigurationClient) patchHandleError(resp *http.Response) error {
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
