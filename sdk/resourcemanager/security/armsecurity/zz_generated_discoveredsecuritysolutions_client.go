//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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

// DiscoveredSecuritySolutionsClient contains the methods for the DiscoveredSecuritySolutions group.
// Don't use this type directly, use NewDiscoveredSecuritySolutionsClient() instead.
type DiscoveredSecuritySolutionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDiscoveredSecuritySolutionsClient creates a new instance of DiscoveredSecuritySolutionsClient with the specified values.
// subscriptionID - Azure subscription ID
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDiscoveredSecuritySolutionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DiscoveredSecuritySolutionsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &DiscoveredSecuritySolutionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets a specific discovered Security Solution.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
// ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
// discoveredSecuritySolutionName - Name of a discovered security solution.
// options - DiscoveredSecuritySolutionsClientGetOptions contains the optional parameters for the DiscoveredSecuritySolutionsClient.Get
// method.
func (client *DiscoveredSecuritySolutionsClient) Get(ctx context.Context, resourceGroupName string, ascLocation string, discoveredSecuritySolutionName string, options *DiscoveredSecuritySolutionsClientGetOptions) (DiscoveredSecuritySolutionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ascLocation, discoveredSecuritySolutionName, options)
	if err != nil {
		return DiscoveredSecuritySolutionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DiscoveredSecuritySolutionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DiscoveredSecuritySolutionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DiscoveredSecuritySolutionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, ascLocation string, discoveredSecuritySolutionName string, options *DiscoveredSecuritySolutionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/discoveredSecuritySolutions/{discoveredSecuritySolutionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ascLocation == "" {
		return nil, errors.New("parameter ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(ascLocation))
	if discoveredSecuritySolutionName == "" {
		return nil, errors.New("parameter discoveredSecuritySolutionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{discoveredSecuritySolutionName}", url.PathEscape(discoveredSecuritySolutionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DiscoveredSecuritySolutionsClient) getHandleResponse(resp *http.Response) (DiscoveredSecuritySolutionsClientGetResponse, error) {
	result := DiscoveredSecuritySolutionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiscoveredSecuritySolution); err != nil {
		return DiscoveredSecuritySolutionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of discovered Security Solutions for the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - DiscoveredSecuritySolutionsClientListOptions contains the optional parameters for the DiscoveredSecuritySolutionsClient.List
// method.
func (client *DiscoveredSecuritySolutionsClient) NewListPager(options *DiscoveredSecuritySolutionsClientListOptions) *runtime.Pager[DiscoveredSecuritySolutionsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[DiscoveredSecuritySolutionsClientListResponse]{
		More: func(page DiscoveredSecuritySolutionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DiscoveredSecuritySolutionsClientListResponse) (DiscoveredSecuritySolutionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DiscoveredSecuritySolutionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DiscoveredSecuritySolutionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DiscoveredSecuritySolutionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DiscoveredSecuritySolutionsClient) listCreateRequest(ctx context.Context, options *DiscoveredSecuritySolutionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/discoveredSecuritySolutions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DiscoveredSecuritySolutionsClient) listHandleResponse(resp *http.Response) (DiscoveredSecuritySolutionsClientListResponse, error) {
	result := DiscoveredSecuritySolutionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiscoveredSecuritySolutionList); err != nil {
		return DiscoveredSecuritySolutionsClientListResponse{}, err
	}
	return result, nil
}

// NewListByHomeRegionPager - Gets a list of discovered Security Solutions for the subscription and location.
// If the operation fails it returns an *azcore.ResponseError type.
// ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
// options - DiscoveredSecuritySolutionsClientListByHomeRegionOptions contains the optional parameters for the DiscoveredSecuritySolutionsClient.ListByHomeRegion
// method.
func (client *DiscoveredSecuritySolutionsClient) NewListByHomeRegionPager(ascLocation string, options *DiscoveredSecuritySolutionsClientListByHomeRegionOptions) *runtime.Pager[DiscoveredSecuritySolutionsClientListByHomeRegionResponse] {
	return runtime.NewPager(runtime.PageProcessor[DiscoveredSecuritySolutionsClientListByHomeRegionResponse]{
		More: func(page DiscoveredSecuritySolutionsClientListByHomeRegionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DiscoveredSecuritySolutionsClientListByHomeRegionResponse) (DiscoveredSecuritySolutionsClientListByHomeRegionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByHomeRegionCreateRequest(ctx, ascLocation, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DiscoveredSecuritySolutionsClientListByHomeRegionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DiscoveredSecuritySolutionsClientListByHomeRegionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DiscoveredSecuritySolutionsClientListByHomeRegionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByHomeRegionHandleResponse(resp)
		},
	})
}

// listByHomeRegionCreateRequest creates the ListByHomeRegion request.
func (client *DiscoveredSecuritySolutionsClient) listByHomeRegionCreateRequest(ctx context.Context, ascLocation string, options *DiscoveredSecuritySolutionsClientListByHomeRegionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/discoveredSecuritySolutions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if ascLocation == "" {
		return nil, errors.New("parameter ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(ascLocation))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByHomeRegionHandleResponse handles the ListByHomeRegion response.
func (client *DiscoveredSecuritySolutionsClient) listByHomeRegionHandleResponse(resp *http.Response) (DiscoveredSecuritySolutionsClientListByHomeRegionResponse, error) {
	result := DiscoveredSecuritySolutionsClientListByHomeRegionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiscoveredSecuritySolutionList); err != nil {
		return DiscoveredSecuritySolutionsClientListByHomeRegionResponse{}, err
	}
	return result, nil
}
