//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmanagedservices

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

// MarketplaceRegistrationDefinitionsWithoutScopeClient contains the methods for the MarketplaceRegistrationDefinitionsWithoutScope group.
// Don't use this type directly, use NewMarketplaceRegistrationDefinitionsWithoutScopeClient() instead.
type MarketplaceRegistrationDefinitionsWithoutScopeClient struct {
	internal *arm.Client
}

// NewMarketplaceRegistrationDefinitionsWithoutScopeClient creates a new instance of MarketplaceRegistrationDefinitionsWithoutScopeClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMarketplaceRegistrationDefinitionsWithoutScopeClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*MarketplaceRegistrationDefinitionsWithoutScopeClient, error) {
	cl, err := arm.NewClient(moduleName+".MarketplaceRegistrationDefinitionsWithoutScopeClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MarketplaceRegistrationDefinitionsWithoutScopeClient{
		internal: cl,
	}
	return client, nil
}

// Get - Get the marketplace registration definition for the marketplace identifier.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-01-preview
//   - marketplaceIdentifier - The Azure Marketplace identifier. Expected formats: {publisher}.{product[-preview]}.{planName}.{version}
//     or {publisher}.{product[-preview]}.{planName} or {publisher}.{product[-preview]} or
//     {publisher}).
//   - options - MarketplaceRegistrationDefinitionsWithoutScopeClientGetOptions contains the optional parameters for the MarketplaceRegistrationDefinitionsWithoutScopeClient.Get
//     method.
func (client *MarketplaceRegistrationDefinitionsWithoutScopeClient) Get(ctx context.Context, marketplaceIdentifier string, options *MarketplaceRegistrationDefinitionsWithoutScopeClientGetOptions) (MarketplaceRegistrationDefinitionsWithoutScopeClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, marketplaceIdentifier, options)
	if err != nil {
		return MarketplaceRegistrationDefinitionsWithoutScopeClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MarketplaceRegistrationDefinitionsWithoutScopeClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MarketplaceRegistrationDefinitionsWithoutScopeClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MarketplaceRegistrationDefinitionsWithoutScopeClient) getCreateRequest(ctx context.Context, marketplaceIdentifier string, options *MarketplaceRegistrationDefinitionsWithoutScopeClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions/{marketplaceIdentifier}"
	if marketplaceIdentifier == "" {
		return nil, errors.New("parameter marketplaceIdentifier cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{marketplaceIdentifier}", url.PathEscape(marketplaceIdentifier))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MarketplaceRegistrationDefinitionsWithoutScopeClient) getHandleResponse(resp *http.Response) (MarketplaceRegistrationDefinitionsWithoutScopeClientGetResponse, error) {
	result := MarketplaceRegistrationDefinitionsWithoutScopeClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MarketplaceRegistrationDefinition); err != nil {
		return MarketplaceRegistrationDefinitionsWithoutScopeClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of the marketplace registration definitions for the marketplace identifier.
//
// Generated from API version 2022-01-01-preview
//   - options - MarketplaceRegistrationDefinitionsWithoutScopeClientListOptions contains the optional parameters for the MarketplaceRegistrationDefinitionsWithoutScopeClient.NewListPager
//     method.
func (client *MarketplaceRegistrationDefinitionsWithoutScopeClient) NewListPager(options *MarketplaceRegistrationDefinitionsWithoutScopeClientListOptions) *runtime.Pager[MarketplaceRegistrationDefinitionsWithoutScopeClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[MarketplaceRegistrationDefinitionsWithoutScopeClientListResponse]{
		More: func(page MarketplaceRegistrationDefinitionsWithoutScopeClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MarketplaceRegistrationDefinitionsWithoutScopeClientListResponse) (MarketplaceRegistrationDefinitionsWithoutScopeClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return MarketplaceRegistrationDefinitionsWithoutScopeClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return MarketplaceRegistrationDefinitionsWithoutScopeClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return MarketplaceRegistrationDefinitionsWithoutScopeClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *MarketplaceRegistrationDefinitionsWithoutScopeClient) listCreateRequest(ctx context.Context, options *MarketplaceRegistrationDefinitionsWithoutScopeClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MarketplaceRegistrationDefinitionsWithoutScopeClient) listHandleResponse(resp *http.Response) (MarketplaceRegistrationDefinitionsWithoutScopeClientListResponse, error) {
	result := MarketplaceRegistrationDefinitionsWithoutScopeClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MarketplaceRegistrationDefinitionList); err != nil {
		return MarketplaceRegistrationDefinitionsWithoutScopeClientListResponse{}, err
	}
	return result, nil
}
