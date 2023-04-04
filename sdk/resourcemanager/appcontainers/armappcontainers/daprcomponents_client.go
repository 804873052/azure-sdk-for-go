//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappcontainers

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

// DaprComponentsClient contains the methods for the DaprComponents group.
// Don't use this type directly, use NewDaprComponentsClient() instead.
type DaprComponentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDaprComponentsClient creates a new instance of DaprComponentsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDaprComponentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DaprComponentsClient, error) {
	cl, err := arm.NewClient(moduleName+".DaprComponentsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DaprComponentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a Dapr Component in a Managed Environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - componentName - Name of the Dapr Component.
//   - daprComponentEnvelope - Configuration details of the Dapr Component.
//   - options - DaprComponentsClientCreateOrUpdateOptions contains the optional parameters for the DaprComponentsClient.CreateOrUpdate
//     method.
func (client *DaprComponentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, environmentName string, componentName string, daprComponentEnvelope DaprComponent, options *DaprComponentsClientCreateOrUpdateOptions) (DaprComponentsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, environmentName, componentName, daprComponentEnvelope, options)
	if err != nil {
		return DaprComponentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DaprComponentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DaprComponentsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DaprComponentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, componentName string, daprComponentEnvelope DaprComponent, options *DaprComponentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/daprComponents/{componentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if componentName == "" {
		return nil, errors.New("parameter componentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentName}", url.PathEscape(componentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, daprComponentEnvelope)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DaprComponentsClient) createOrUpdateHandleResponse(resp *http.Response) (DaprComponentsClientCreateOrUpdateResponse, error) {
	result := DaprComponentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprComponent); err != nil {
		return DaprComponentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Dapr Component from a Managed Environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - componentName - Name of the Dapr Component.
//   - options - DaprComponentsClientDeleteOptions contains the optional parameters for the DaprComponentsClient.Delete method.
func (client *DaprComponentsClient) Delete(ctx context.Context, resourceGroupName string, environmentName string, componentName string, options *DaprComponentsClientDeleteOptions) (DaprComponentsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, environmentName, componentName, options)
	if err != nil {
		return DaprComponentsClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DaprComponentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return DaprComponentsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return DaprComponentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DaprComponentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, componentName string, options *DaprComponentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/daprComponents/{componentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if componentName == "" {
		return nil, errors.New("parameter componentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentName}", url.PathEscape(componentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a dapr component.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - componentName - Name of the Dapr Component.
//   - options - DaprComponentsClientGetOptions contains the optional parameters for the DaprComponentsClient.Get method.
func (client *DaprComponentsClient) Get(ctx context.Context, resourceGroupName string, environmentName string, componentName string, options *DaprComponentsClientGetOptions) (DaprComponentsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, environmentName, componentName, options)
	if err != nil {
		return DaprComponentsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DaprComponentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DaprComponentsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DaprComponentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, componentName string, options *DaprComponentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/daprComponents/{componentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if componentName == "" {
		return nil, errors.New("parameter componentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentName}", url.PathEscape(componentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DaprComponentsClient) getHandleResponse(resp *http.Response) (DaprComponentsClientGetResponse, error) {
	result := DaprComponentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprComponent); err != nil {
		return DaprComponentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get the Dapr Components for a managed environment.
//
// Generated from API version 2022-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - options - DaprComponentsClientListOptions contains the optional parameters for the DaprComponentsClient.NewListPager method.
func (client *DaprComponentsClient) NewListPager(resourceGroupName string, environmentName string, options *DaprComponentsClientListOptions) *runtime.Pager[DaprComponentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DaprComponentsClientListResponse]{
		More: func(page DaprComponentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DaprComponentsClientListResponse) (DaprComponentsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, environmentName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DaprComponentsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DaprComponentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DaprComponentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DaprComponentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, options *DaprComponentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/daprComponents"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DaprComponentsClient) listHandleResponse(resp *http.Response) (DaprComponentsClientListResponse, error) {
	result := DaprComponentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprComponentsCollection); err != nil {
		return DaprComponentsClientListResponse{}, err
	}
	return result, nil
}

// ListSecrets - List secrets for a dapr component
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - componentName - Name of the Dapr Component.
//   - options - DaprComponentsClientListSecretsOptions contains the optional parameters for the DaprComponentsClient.ListSecrets
//     method.
func (client *DaprComponentsClient) ListSecrets(ctx context.Context, resourceGroupName string, environmentName string, componentName string, options *DaprComponentsClientListSecretsOptions) (DaprComponentsClientListSecretsResponse, error) {
	req, err := client.listSecretsCreateRequest(ctx, resourceGroupName, environmentName, componentName, options)
	if err != nil {
		return DaprComponentsClientListSecretsResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DaprComponentsClientListSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DaprComponentsClientListSecretsResponse{}, runtime.NewResponseError(resp)
	}
	return client.listSecretsHandleResponse(resp)
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *DaprComponentsClient) listSecretsCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, componentName string, options *DaprComponentsClientListSecretsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/daprComponents/{componentName}/listSecrets"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if componentName == "" {
		return nil, errors.New("parameter componentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentName}", url.PathEscape(componentName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *DaprComponentsClient) listSecretsHandleResponse(resp *http.Response) (DaprComponentsClientListSecretsResponse, error) {
	result := DaprComponentsClientListSecretsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DaprSecretsCollection); err != nil {
		return DaprComponentsClientListSecretsResponse{}, err
	}
	return result, nil
}
