//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetworkcloud

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

// L3NetworksClient contains the methods for the L3Networks group.
// Don't use this type directly, use NewL3NetworksClient() instead.
type L3NetworksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewL3NetworksClient creates a new instance of L3NetworksClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewL3NetworksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*L3NetworksClient, error) {
	cl, err := arm.NewClient(moduleName+".L3NetworksClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &L3NetworksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a new layer 3 (L3) network or update the properties of the existing network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l3NetworkName - The name of the L3 network.
//   - l3NetworkParameters - The request body.
//   - options - L3NetworksClientBeginCreateOrUpdateOptions contains the optional parameters for the L3NetworksClient.BeginCreateOrUpdate
//     method.
func (client *L3NetworksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, l3NetworkName string, l3NetworkParameters L3Network, options *L3NetworksClientBeginCreateOrUpdateOptions) (*runtime.Poller[L3NetworksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, l3NetworkName, l3NetworkParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[L3NetworksClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[L3NetworksClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a new layer 3 (L3) network or update the properties of the existing network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *L3NetworksClient) createOrUpdate(ctx context.Context, resourceGroupName string, l3NetworkName string, l3NetworkParameters L3Network, options *L3NetworksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, l3NetworkName, l3NetworkParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *L3NetworksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, l3NetworkName string, l3NetworkParameters L3Network, options *L3NetworksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/l3Networks/{l3NetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l3NetworkName == "" {
		return nil, errors.New("parameter l3NetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l3NetworkName}", url.PathEscape(l3NetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, l3NetworkParameters)
}

// BeginDelete - Delete the provided layer 3 (L3) network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l3NetworkName - The name of the L3 network.
//   - options - L3NetworksClientBeginDeleteOptions contains the optional parameters for the L3NetworksClient.BeginDelete method.
func (client *L3NetworksClient) BeginDelete(ctx context.Context, resourceGroupName string, l3NetworkName string, options *L3NetworksClientBeginDeleteOptions) (*runtime.Poller[L3NetworksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, l3NetworkName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[L3NetworksClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[L3NetworksClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete the provided layer 3 (L3) network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *L3NetworksClient) deleteOperation(ctx context.Context, resourceGroupName string, l3NetworkName string, options *L3NetworksClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, l3NetworkName, options)
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
func (client *L3NetworksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, l3NetworkName string, options *L3NetworksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/l3Networks/{l3NetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l3NetworkName == "" {
		return nil, errors.New("parameter l3NetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l3NetworkName}", url.PathEscape(l3NetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get properties of the provided layer 3 (L3) network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l3NetworkName - The name of the L3 network.
//   - options - L3NetworksClientGetOptions contains the optional parameters for the L3NetworksClient.Get method.
func (client *L3NetworksClient) Get(ctx context.Context, resourceGroupName string, l3NetworkName string, options *L3NetworksClientGetOptions) (L3NetworksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, l3NetworkName, options)
	if err != nil {
		return L3NetworksClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return L3NetworksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return L3NetworksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *L3NetworksClient) getCreateRequest(ctx context.Context, resourceGroupName string, l3NetworkName string, options *L3NetworksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/l3Networks/{l3NetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l3NetworkName == "" {
		return nil, errors.New("parameter l3NetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l3NetworkName}", url.PathEscape(l3NetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *L3NetworksClient) getHandleResponse(resp *http.Response) (L3NetworksClientGetResponse, error) {
	result := L3NetworksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.L3Network); err != nil {
		return L3NetworksClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get a list of layer 3 (L3) networks in the provided resource group.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - L3NetworksClientListByResourceGroupOptions contains the optional parameters for the L3NetworksClient.NewListByResourceGroupPager
//     method.
func (client *L3NetworksClient) NewListByResourceGroupPager(resourceGroupName string, options *L3NetworksClientListByResourceGroupOptions) *runtime.Pager[L3NetworksClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[L3NetworksClientListByResourceGroupResponse]{
		More: func(page L3NetworksClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *L3NetworksClientListByResourceGroupResponse) (L3NetworksClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return L3NetworksClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return L3NetworksClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return L3NetworksClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *L3NetworksClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *L3NetworksClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/l3Networks"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *L3NetworksClient) listByResourceGroupHandleResponse(resp *http.Response) (L3NetworksClientListByResourceGroupResponse, error) {
	result := L3NetworksClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.L3NetworkList); err != nil {
		return L3NetworksClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Get a list of layer 3 (L3) networks in the provided subscription.
//
// Generated from API version 2023-05-01-preview
//   - options - L3NetworksClientListBySubscriptionOptions contains the optional parameters for the L3NetworksClient.NewListBySubscriptionPager
//     method.
func (client *L3NetworksClient) NewListBySubscriptionPager(options *L3NetworksClientListBySubscriptionOptions) *runtime.Pager[L3NetworksClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[L3NetworksClientListBySubscriptionResponse]{
		More: func(page L3NetworksClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *L3NetworksClientListBySubscriptionResponse) (L3NetworksClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return L3NetworksClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return L3NetworksClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return L3NetworksClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *L3NetworksClient) listBySubscriptionCreateRequest(ctx context.Context, options *L3NetworksClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetworkCloud/l3Networks"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *L3NetworksClient) listBySubscriptionHandleResponse(resp *http.Response) (L3NetworksClientListBySubscriptionResponse, error) {
	result := L3NetworksClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.L3NetworkList); err != nil {
		return L3NetworksClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update tags associated with the provided layer 3 (L3) network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l3NetworkName - The name of the L3 network.
//   - l3NetworkUpdateParameters - The request body.
//   - options - L3NetworksClientUpdateOptions contains the optional parameters for the L3NetworksClient.Update method.
func (client *L3NetworksClient) Update(ctx context.Context, resourceGroupName string, l3NetworkName string, l3NetworkUpdateParameters L3NetworkPatchParameters, options *L3NetworksClientUpdateOptions) (L3NetworksClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, l3NetworkName, l3NetworkUpdateParameters, options)
	if err != nil {
		return L3NetworksClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return L3NetworksClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return L3NetworksClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *L3NetworksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, l3NetworkName string, l3NetworkUpdateParameters L3NetworkPatchParameters, options *L3NetworksClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/l3Networks/{l3NetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l3NetworkName == "" {
		return nil, errors.New("parameter l3NetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l3NetworkName}", url.PathEscape(l3NetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, l3NetworkUpdateParameters)
}

// updateHandleResponse handles the Update response.
func (client *L3NetworksClient) updateHandleResponse(resp *http.Response) (L3NetworksClientUpdateResponse, error) {
	result := L3NetworksClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.L3Network); err != nil {
		return L3NetworksClientUpdateResponse{}, err
	}
	return result, nil
}
