//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomerinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// HubsClient contains the methods for the Hubs group.
// Don't use this type directly, use NewHubsClient() instead.
type HubsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewHubsClient creates a new instance of HubsClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewHubsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HubsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &HubsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates a hub, or updates an existing hub.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the Hub.
// parameters - Parameters supplied to the CreateOrUpdate Hub operation.
// options - HubsClientCreateOrUpdateOptions contains the optional parameters for the HubsClient.CreateOrUpdate method.
func (client *HubsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, hubName string, parameters Hub, options *HubsClientCreateOrUpdateOptions) (HubsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, hubName, parameters, options)
	if err != nil {
		return HubsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HubsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return HubsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *HubsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, hubName string, parameters Hub, options *HubsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *HubsClient) createOrUpdateHandleResponse(resp *http.Response) (HubsClientCreateOrUpdateResponse, error) {
	result := HubsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Hub); err != nil {
		return HubsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes the specified hub.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// options - HubsClientBeginDeleteOptions contains the optional parameters for the HubsClient.BeginDelete method.
func (client *HubsClient) BeginDelete(ctx context.Context, resourceGroupName string, hubName string, options *HubsClientBeginDeleteOptions) (*runtime.Poller[HubsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, hubName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[HubsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[HubsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified hub.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
func (client *HubsClient) deleteOperation(ctx context.Context, resourceGroupName string, hubName string, options *HubsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, hubName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HubsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, hubName string, options *HubsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets information about the specified hub.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the hub.
// options - HubsClientGetOptions contains the optional parameters for the HubsClient.Get method.
func (client *HubsClient) Get(ctx context.Context, resourceGroupName string, hubName string, options *HubsClientGetOptions) (HubsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, hubName, options)
	if err != nil {
		return HubsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HubsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HubsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HubsClient) getCreateRequest(ctx context.Context, resourceGroupName string, hubName string, options *HubsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HubsClient) getHandleResponse(resp *http.Response) (HubsClientGetResponse, error) {
	result := HubsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Hub); err != nil {
		return HubsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all hubs in the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// options - HubsClientListOptions contains the optional parameters for the HubsClient.List method.
func (client *HubsClient) NewListPager(options *HubsClientListOptions) *runtime.Pager[HubsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[HubsClientListResponse]{
		More: func(page HubsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HubsClientListResponse) (HubsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return HubsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return HubsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HubsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *HubsClient) listCreateRequest(ctx context.Context, options *HubsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.CustomerInsights/hubs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *HubsClient) listHandleResponse(resp *http.Response) (HubsClientListResponse, error) {
	result := HubsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HubListResult); err != nil {
		return HubsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets all the hubs in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// options - HubsClientListByResourceGroupOptions contains the optional parameters for the HubsClient.ListByResourceGroup
// method.
func (client *HubsClient) NewListByResourceGroupPager(resourceGroupName string, options *HubsClientListByResourceGroupOptions) *runtime.Pager[HubsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[HubsClientListByResourceGroupResponse]{
		More: func(page HubsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HubsClientListByResourceGroupResponse) (HubsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return HubsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return HubsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HubsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *HubsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *HubsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *HubsClient) listByResourceGroupHandleResponse(resp *http.Response) (HubsClientListByResourceGroupResponse, error) {
	result := HubsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HubListResult); err != nil {
		return HubsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// Update - Updates a Hub.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2017-04-26
// resourceGroupName - The name of the resource group.
// hubName - The name of the Hub.
// parameters - Parameters supplied to the Update Hub operation.
// options - HubsClientUpdateOptions contains the optional parameters for the HubsClient.Update method.
func (client *HubsClient) Update(ctx context.Context, resourceGroupName string, hubName string, parameters Hub, options *HubsClientUpdateOptions) (HubsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, hubName, parameters, options)
	if err != nil {
		return HubsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HubsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HubsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *HubsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, hubName string, parameters Hub, options *HubsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.CustomerInsights/hubs/{hubName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-04-26")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *HubsClient) updateHandleResponse(resp *http.Response) (HubsClientUpdateResponse, error) {
	result := HubsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Hub); err != nil {
		return HubsClientUpdateResponse{}, err
	}
	return result, nil
}
