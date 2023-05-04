//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcdn

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

// CustomDomainsClient contains the methods for the CustomDomains group.
// Don't use this type directly, use NewCustomDomainsClient() instead.
type CustomDomainsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCustomDomainsClient creates a new instance of CustomDomainsClient with the specified values.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCustomDomainsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CustomDomainsClient, error) {
	cl, err := arm.NewClient(moduleName+".CustomDomainsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CustomDomainsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a new custom domain within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the CDN profile which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - customDomainName - Name of the custom domain within an endpoint.
//   - customDomainProperties - Properties required to create a new custom domain.
//   - options - CustomDomainsClientBeginCreateOptions contains the optional parameters for the CustomDomainsClient.BeginCreate
//     method.
func (client *CustomDomainsClient) BeginCreate(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, customDomainProperties CustomDomainParameters, options *CustomDomainsClientBeginCreateOptions) (*runtime.Poller[CustomDomainsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, profileName, endpointName, customDomainName, customDomainProperties, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[CustomDomainsClientCreateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[CustomDomainsClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates a new custom domain within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
func (client *CustomDomainsClient) create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, customDomainProperties CustomDomainParameters, options *CustomDomainsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, profileName, endpointName, customDomainName, customDomainProperties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *CustomDomainsClient) createCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, customDomainProperties CustomDomainParameters, options *CustomDomainsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if customDomainName == "" {
		return nil, errors.New("parameter customDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customDomainName}", url.PathEscape(customDomainName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, customDomainProperties)
}

// BeginDelete - Deletes an existing custom domain within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the CDN profile which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - customDomainName - Name of the custom domain within an endpoint.
//   - options - CustomDomainsClientBeginDeleteOptions contains the optional parameters for the CustomDomainsClient.BeginDelete
//     method.
func (client *CustomDomainsClient) BeginDelete(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, options *CustomDomainsClientBeginDeleteOptions) (*runtime.Poller[CustomDomainsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, profileName, endpointName, customDomainName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[CustomDomainsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[CustomDomainsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes an existing custom domain within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
func (client *CustomDomainsClient) deleteOperation(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, options *CustomDomainsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, profileName, endpointName, customDomainName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CustomDomainsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, options *CustomDomainsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if customDomainName == "" {
		return nil, errors.New("parameter customDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customDomainName}", url.PathEscape(customDomainName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DisableCustomHTTPS - Disable https delivery of the custom domain.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the CDN profile which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - customDomainName - Name of the custom domain within an endpoint.
//   - options - CustomDomainsClientDisableCustomHTTPSOptions contains the optional parameters for the CustomDomainsClient.DisableCustomHTTPS
//     method.
func (client *CustomDomainsClient) DisableCustomHTTPS(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, options *CustomDomainsClientDisableCustomHTTPSOptions) (CustomDomainsClientDisableCustomHTTPSResponse, error) {
	req, err := client.disableCustomHTTPSCreateRequest(ctx, resourceGroupName, profileName, endpointName, customDomainName, options)
	if err != nil {
		return CustomDomainsClientDisableCustomHTTPSResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomDomainsClientDisableCustomHTTPSResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return CustomDomainsClientDisableCustomHTTPSResponse{}, runtime.NewResponseError(resp)
	}
	return client.disableCustomHTTPSHandleResponse(resp)
}

// disableCustomHTTPSCreateRequest creates the DisableCustomHTTPS request.
func (client *CustomDomainsClient) disableCustomHTTPSCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, options *CustomDomainsClientDisableCustomHTTPSOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}/disableCustomHttps"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if customDomainName == "" {
		return nil, errors.New("parameter customDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customDomainName}", url.PathEscape(customDomainName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// disableCustomHTTPSHandleResponse handles the DisableCustomHTTPS response.
func (client *CustomDomainsClient) disableCustomHTTPSHandleResponse(resp *http.Response) (CustomDomainsClientDisableCustomHTTPSResponse, error) {
	result := CustomDomainsClientDisableCustomHTTPSResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomDomain); err != nil {
		return CustomDomainsClientDisableCustomHTTPSResponse{}, err
	}
	return result, nil
}

// EnableCustomHTTPS - Enable https delivery of the custom domain.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the CDN profile which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - customDomainName - Name of the custom domain within an endpoint.
//   - options - CustomDomainsClientEnableCustomHTTPSOptions contains the optional parameters for the CustomDomainsClient.EnableCustomHTTPS
//     method.
func (client *CustomDomainsClient) EnableCustomHTTPS(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, options *CustomDomainsClientEnableCustomHTTPSOptions) (CustomDomainsClientEnableCustomHTTPSResponse, error) {
	req, err := client.enableCustomHTTPSCreateRequest(ctx, resourceGroupName, profileName, endpointName, customDomainName, options)
	if err != nil {
		return CustomDomainsClientEnableCustomHTTPSResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomDomainsClientEnableCustomHTTPSResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return CustomDomainsClientEnableCustomHTTPSResponse{}, runtime.NewResponseError(resp)
	}
	return client.enableCustomHTTPSHandleResponse(resp)
}

// enableCustomHTTPSCreateRequest creates the EnableCustomHTTPS request.
func (client *CustomDomainsClient) enableCustomHTTPSCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, options *CustomDomainsClientEnableCustomHTTPSOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}/enableCustomHttps"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if customDomainName == "" {
		return nil, errors.New("parameter customDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customDomainName}", url.PathEscape(customDomainName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.CustomDomainHTTPSParameters != nil {
		return req, runtime.MarshalAsJSON(req, options.CustomDomainHTTPSParameters)
	}
	return req, nil
}

// enableCustomHTTPSHandleResponse handles the EnableCustomHTTPS response.
func (client *CustomDomainsClient) enableCustomHTTPSHandleResponse(resp *http.Response) (CustomDomainsClientEnableCustomHTTPSResponse, error) {
	result := CustomDomainsClientEnableCustomHTTPSResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomDomain); err != nil {
		return CustomDomainsClientEnableCustomHTTPSResponse{}, err
	}
	return result, nil
}

// Get - Gets an existing custom domain within an endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the CDN profile which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - customDomainName - Name of the custom domain within an endpoint.
//   - options - CustomDomainsClientGetOptions contains the optional parameters for the CustomDomainsClient.Get method.
func (client *CustomDomainsClient) Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, options *CustomDomainsClientGetOptions) (CustomDomainsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, endpointName, customDomainName, options)
	if err != nil {
		return CustomDomainsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CustomDomainsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomDomainsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CustomDomainsClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, options *CustomDomainsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if customDomainName == "" {
		return nil, errors.New("parameter customDomainName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customDomainName}", url.PathEscape(customDomainName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CustomDomainsClient) getHandleResponse(resp *http.Response) (CustomDomainsClientGetResponse, error) {
	result := CustomDomainsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomDomain); err != nil {
		return CustomDomainsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByEndpointPager - Lists all of the existing custom domains within an endpoint.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - Name of the Resource group within the Azure subscription.
//   - profileName - Name of the CDN profile which is unique within the resource group.
//   - endpointName - Name of the endpoint under the profile which is unique globally.
//   - options - CustomDomainsClientListByEndpointOptions contains the optional parameters for the CustomDomainsClient.NewListByEndpointPager
//     method.
func (client *CustomDomainsClient) NewListByEndpointPager(resourceGroupName string, profileName string, endpointName string, options *CustomDomainsClientListByEndpointOptions) *runtime.Pager[CustomDomainsClientListByEndpointResponse] {
	return runtime.NewPager(runtime.PagingHandler[CustomDomainsClientListByEndpointResponse]{
		More: func(page CustomDomainsClientListByEndpointResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CustomDomainsClientListByEndpointResponse) (CustomDomainsClientListByEndpointResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByEndpointCreateRequest(ctx, resourceGroupName, profileName, endpointName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CustomDomainsClientListByEndpointResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return CustomDomainsClientListByEndpointResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CustomDomainsClientListByEndpointResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByEndpointHandleResponse(resp)
		},
	})
}

// listByEndpointCreateRequest creates the ListByEndpoint request.
func (client *CustomDomainsClient) listByEndpointCreateRequest(ctx context.Context, resourceGroupName string, profileName string, endpointName string, options *CustomDomainsClientListByEndpointOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	if endpointName == "" {
		return nil, errors.New("parameter endpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", url.PathEscape(endpointName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByEndpointHandleResponse handles the ListByEndpoint response.
func (client *CustomDomainsClient) listByEndpointHandleResponse(resp *http.Response) (CustomDomainsClientListByEndpointResponse, error) {
	result := CustomDomainsClientListByEndpointResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomDomainListResult); err != nil {
		return CustomDomainsClientListByEndpointResponse{}, err
	}
	return result, nil
}
