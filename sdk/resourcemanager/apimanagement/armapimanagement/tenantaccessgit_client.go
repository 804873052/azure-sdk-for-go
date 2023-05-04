//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armapimanagement

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// TenantAccessGitClient contains the methods for the TenantAccessGit group.
// Don't use this type directly, use NewTenantAccessGitClient() instead.
type TenantAccessGitClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTenantAccessGitClient creates a new instance of TenantAccessGitClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTenantAccessGitClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TenantAccessGitClient, error) {
	cl, err := arm.NewClient(moduleName+".TenantAccessGitClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TenantAccessGitClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// RegeneratePrimaryKey - Regenerate primary access key for GIT.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group.
//   - serviceName - The name of the API Management service.
//   - accessName - The identifier of the Access configuration.
//   - options - TenantAccessGitClientRegeneratePrimaryKeyOptions contains the optional parameters for the TenantAccessGitClient.RegeneratePrimaryKey
//     method.
func (client *TenantAccessGitClient) RegeneratePrimaryKey(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessGitClientRegeneratePrimaryKeyOptions) (TenantAccessGitClientRegeneratePrimaryKeyResponse, error) {
	req, err := client.regeneratePrimaryKeyCreateRequest(ctx, resourceGroupName, serviceName, accessName, options)
	if err != nil {
		return TenantAccessGitClientRegeneratePrimaryKeyResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TenantAccessGitClientRegeneratePrimaryKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return TenantAccessGitClientRegeneratePrimaryKeyResponse{}, runtime.NewResponseError(resp)
	}
	return TenantAccessGitClientRegeneratePrimaryKeyResponse{}, nil
}

// regeneratePrimaryKeyCreateRequest creates the RegeneratePrimaryKey request.
func (client *TenantAccessGitClient) regeneratePrimaryKeyCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessGitClientRegeneratePrimaryKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}/git/regeneratePrimaryKey"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accessName == "" {
		return nil, errors.New("parameter accessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessName}", url.PathEscape(string(accessName)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// RegenerateSecondaryKey - Regenerate secondary access key for GIT.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-08-01
//   - resourceGroupName - The name of the resource group.
//   - serviceName - The name of the API Management service.
//   - accessName - The identifier of the Access configuration.
//   - options - TenantAccessGitClientRegenerateSecondaryKeyOptions contains the optional parameters for the TenantAccessGitClient.RegenerateSecondaryKey
//     method.
func (client *TenantAccessGitClient) RegenerateSecondaryKey(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessGitClientRegenerateSecondaryKeyOptions) (TenantAccessGitClientRegenerateSecondaryKeyResponse, error) {
	req, err := client.regenerateSecondaryKeyCreateRequest(ctx, resourceGroupName, serviceName, accessName, options)
	if err != nil {
		return TenantAccessGitClientRegenerateSecondaryKeyResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TenantAccessGitClientRegenerateSecondaryKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return TenantAccessGitClientRegenerateSecondaryKeyResponse{}, runtime.NewResponseError(resp)
	}
	return TenantAccessGitClientRegenerateSecondaryKeyResponse{}, nil
}

// regenerateSecondaryKeyCreateRequest creates the RegenerateSecondaryKey request.
func (client *TenantAccessGitClient) regenerateSecondaryKeyCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessGitClientRegenerateSecondaryKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}/git/regenerateSecondaryKey"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accessName == "" {
		return nil, errors.New("parameter accessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessName}", url.PathEscape(string(accessName)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
