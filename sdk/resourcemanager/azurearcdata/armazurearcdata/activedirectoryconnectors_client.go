//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armazurearcdata

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

// ActiveDirectoryConnectorsClient contains the methods for the ActiveDirectoryConnectors group.
// Don't use this type directly, use NewActiveDirectoryConnectorsClient() instead.
type ActiveDirectoryConnectorsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewActiveDirectoryConnectorsClient creates a new instance of ActiveDirectoryConnectorsClient with the specified values.
//   - subscriptionID - The ID of the Azure subscription
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewActiveDirectoryConnectorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ActiveDirectoryConnectorsClient, error) {
	cl, err := arm.NewClient(moduleName+".ActiveDirectoryConnectorsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ActiveDirectoryConnectorsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates or replaces an Active Directory connector resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01-preview
//   - resourceGroupName - The name of the Azure resource group
//   - dataControllerName - The name of the data controller
//   - activeDirectoryConnectorName - The name of the Active Directory connector instance
//   - activeDirectoryConnectorResource - desc
//   - options - ActiveDirectoryConnectorsClientBeginCreateOptions contains the optional parameters for the ActiveDirectoryConnectorsClient.BeginCreate
//     method.
func (client *ActiveDirectoryConnectorsClient) BeginCreate(ctx context.Context, resourceGroupName string, dataControllerName string, activeDirectoryConnectorName string, activeDirectoryConnectorResource ActiveDirectoryConnectorResource, options *ActiveDirectoryConnectorsClientBeginCreateOptions) (*runtime.Poller[ActiveDirectoryConnectorsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, dataControllerName, activeDirectoryConnectorName, activeDirectoryConnectorResource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ActiveDirectoryConnectorsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ActiveDirectoryConnectorsClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates or replaces an Active Directory connector resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01-preview
func (client *ActiveDirectoryConnectorsClient) create(ctx context.Context, resourceGroupName string, dataControllerName string, activeDirectoryConnectorName string, activeDirectoryConnectorResource ActiveDirectoryConnectorResource, options *ActiveDirectoryConnectorsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, dataControllerName, activeDirectoryConnectorName, activeDirectoryConnectorResource, options)
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

// createCreateRequest creates the Create request.
func (client *ActiveDirectoryConnectorsClient) createCreateRequest(ctx context.Context, resourceGroupName string, dataControllerName string, activeDirectoryConnectorName string, activeDirectoryConnectorResource ActiveDirectoryConnectorResource, options *ActiveDirectoryConnectorsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/dataControllers/{dataControllerName}/activeDirectoryConnectors/{activeDirectoryConnectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataControllerName == "" {
		return nil, errors.New("parameter dataControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataControllerName}", url.PathEscape(dataControllerName))
	if activeDirectoryConnectorName == "" {
		return nil, errors.New("parameter activeDirectoryConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{activeDirectoryConnectorName}", url.PathEscape(activeDirectoryConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, activeDirectoryConnectorResource)
}

// BeginDelete - Deletes an Active Directory connector resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01-preview
//   - resourceGroupName - The name of the Azure resource group
//   - dataControllerName - The name of the data controller
//   - activeDirectoryConnectorName - The name of the Active Directory connector instance
//   - options - ActiveDirectoryConnectorsClientBeginDeleteOptions contains the optional parameters for the ActiveDirectoryConnectorsClient.BeginDelete
//     method.
func (client *ActiveDirectoryConnectorsClient) BeginDelete(ctx context.Context, resourceGroupName string, dataControllerName string, activeDirectoryConnectorName string, options *ActiveDirectoryConnectorsClientBeginDeleteOptions) (*runtime.Poller[ActiveDirectoryConnectorsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, dataControllerName, activeDirectoryConnectorName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ActiveDirectoryConnectorsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ActiveDirectoryConnectorsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes an Active Directory connector resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01-preview
func (client *ActiveDirectoryConnectorsClient) deleteOperation(ctx context.Context, resourceGroupName string, dataControllerName string, activeDirectoryConnectorName string, options *ActiveDirectoryConnectorsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, dataControllerName, activeDirectoryConnectorName, options)
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
func (client *ActiveDirectoryConnectorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, dataControllerName string, activeDirectoryConnectorName string, options *ActiveDirectoryConnectorsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/dataControllers/{dataControllerName}/activeDirectoryConnectors/{activeDirectoryConnectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataControllerName == "" {
		return nil, errors.New("parameter dataControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataControllerName}", url.PathEscape(dataControllerName))
	if activeDirectoryConnectorName == "" {
		return nil, errors.New("parameter activeDirectoryConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{activeDirectoryConnectorName}", url.PathEscape(activeDirectoryConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves an Active Directory connector resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-03-01-preview
//   - resourceGroupName - The name of the Azure resource group
//   - dataControllerName - The name of the data controller
//   - activeDirectoryConnectorName - The name of the Active Directory connector instance
//   - options - ActiveDirectoryConnectorsClientGetOptions contains the optional parameters for the ActiveDirectoryConnectorsClient.Get
//     method.
func (client *ActiveDirectoryConnectorsClient) Get(ctx context.Context, resourceGroupName string, dataControllerName string, activeDirectoryConnectorName string, options *ActiveDirectoryConnectorsClientGetOptions) (ActiveDirectoryConnectorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, dataControllerName, activeDirectoryConnectorName, options)
	if err != nil {
		return ActiveDirectoryConnectorsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ActiveDirectoryConnectorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ActiveDirectoryConnectorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ActiveDirectoryConnectorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, dataControllerName string, activeDirectoryConnectorName string, options *ActiveDirectoryConnectorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/dataControllers/{dataControllerName}/activeDirectoryConnectors/{activeDirectoryConnectorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataControllerName == "" {
		return nil, errors.New("parameter dataControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataControllerName}", url.PathEscape(dataControllerName))
	if activeDirectoryConnectorName == "" {
		return nil, errors.New("parameter activeDirectoryConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{activeDirectoryConnectorName}", url.PathEscape(activeDirectoryConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ActiveDirectoryConnectorsClient) getHandleResponse(resp *http.Response) (ActiveDirectoryConnectorsClientGetResponse, error) {
	result := ActiveDirectoryConnectorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActiveDirectoryConnectorResource); err != nil {
		return ActiveDirectoryConnectorsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List the active directory connectors associated with the given data controller.
//
// Generated from API version 2022-03-01-preview
//   - resourceGroupName - The name of the Azure resource group
//   - dataControllerName - The name of the data controller
//   - options - ActiveDirectoryConnectorsClientListOptions contains the optional parameters for the ActiveDirectoryConnectorsClient.NewListPager
//     method.
func (client *ActiveDirectoryConnectorsClient) NewListPager(resourceGroupName string, dataControllerName string, options *ActiveDirectoryConnectorsClientListOptions) *runtime.Pager[ActiveDirectoryConnectorsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ActiveDirectoryConnectorsClientListResponse]{
		More: func(page ActiveDirectoryConnectorsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ActiveDirectoryConnectorsClientListResponse) (ActiveDirectoryConnectorsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, dataControllerName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ActiveDirectoryConnectorsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ActiveDirectoryConnectorsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ActiveDirectoryConnectorsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ActiveDirectoryConnectorsClient) listCreateRequest(ctx context.Context, resourceGroupName string, dataControllerName string, options *ActiveDirectoryConnectorsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/dataControllers/{dataControllerName}/activeDirectoryConnectors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dataControllerName == "" {
		return nil, errors.New("parameter dataControllerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dataControllerName}", url.PathEscape(dataControllerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ActiveDirectoryConnectorsClient) listHandleResponse(resp *http.Response) (ActiveDirectoryConnectorsClientListResponse, error) {
	result := ActiveDirectoryConnectorsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActiveDirectoryConnectorListResult); err != nil {
		return ActiveDirectoryConnectorsClientListResponse{}, err
	}
	return result, nil
}
