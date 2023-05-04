//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevops

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

// PipelinesClient contains the methods for the Pipelines group.
// Don't use this type directly, use NewPipelinesClient() instead.
type PipelinesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPipelinesClient creates a new instance of PipelinesClient with the specified values.
//   - subscriptionID - Unique identifier of the Azure subscription. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPipelinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PipelinesClient, error) {
	cl, err := arm.NewClient(moduleName+".PipelinesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PipelinesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates an Azure Pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-07-01-preview
//   - resourceGroupName - Name of the resource group within the Azure subscription.
//   - pipelineName - The name of the Azure Pipeline resource in ARM.
//   - createOperationParameters - The request payload to create the Azure Pipeline.
//   - options - PipelinesClientBeginCreateOrUpdateOptions contains the optional parameters for the PipelinesClient.BeginCreateOrUpdate
//     method.
func (client *PipelinesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, pipelineName string, createOperationParameters Pipeline, options *PipelinesClientBeginCreateOrUpdateOptions) (*runtime.Poller[PipelinesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, pipelineName, createOperationParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[PipelinesClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[PipelinesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates an Azure Pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-07-01-preview
func (client *PipelinesClient) createOrUpdate(ctx context.Context, resourceGroupName string, pipelineName string, createOperationParameters Pipeline, options *PipelinesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, pipelineName, createOperationParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PipelinesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, pipelineName string, createOperationParameters Pipeline, options *PipelinesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevOps/pipelines/{pipelineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, createOperationParameters)
}

// Delete - Deletes an Azure Pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-07-01-preview
//   - resourceGroupName - Name of the resource group within the Azure subscription.
//   - pipelineName - The name of the Azure Pipeline resource.
//   - options - PipelinesClientDeleteOptions contains the optional parameters for the PipelinesClient.Delete method.
func (client *PipelinesClient) Delete(ctx context.Context, resourceGroupName string, pipelineName string, options *PipelinesClientDeleteOptions) (PipelinesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, pipelineName, options)
	if err != nil {
		return PipelinesClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelinesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return PipelinesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return PipelinesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PipelinesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, pipelineName string, options *PipelinesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevOps/pipelines/{pipelineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets an existing Azure Pipeline.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-07-01-preview
//   - resourceGroupName - Name of the resource group within the Azure subscription.
//   - pipelineName - The name of the Azure Pipeline resource in ARM.
//   - options - PipelinesClientGetOptions contains the optional parameters for the PipelinesClient.Get method.
func (client *PipelinesClient) Get(ctx context.Context, resourceGroupName string, pipelineName string, options *PipelinesClientGetOptions) (PipelinesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, pipelineName, options)
	if err != nil {
		return PipelinesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelinesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PipelinesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PipelinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, pipelineName string, options *PipelinesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevOps/pipelines/{pipelineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PipelinesClient) getHandleResponse(resp *http.Response) (PipelinesClientGetResponse, error) {
	result := PipelinesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Pipeline); err != nil {
		return PipelinesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all Azure Pipelines under the specified resource group.
//
// Generated from API version 2019-07-01-preview
//   - resourceGroupName - Name of the resource group within the Azure subscription.
//   - options - PipelinesClientListByResourceGroupOptions contains the optional parameters for the PipelinesClient.NewListByResourceGroupPager
//     method.
func (client *PipelinesClient) NewListByResourceGroupPager(resourceGroupName string, options *PipelinesClientListByResourceGroupOptions) *runtime.Pager[PipelinesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[PipelinesClientListByResourceGroupResponse]{
		More: func(page PipelinesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PipelinesClientListByResourceGroupResponse) (PipelinesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PipelinesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PipelinesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PipelinesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *PipelinesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *PipelinesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevOps/pipelines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
	reqQP.Set("api-version", "2019-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *PipelinesClient) listByResourceGroupHandleResponse(resp *http.Response) (PipelinesClientListByResourceGroupResponse, error) {
	result := PipelinesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineListResult); err != nil {
		return PipelinesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all Azure Pipelines under the specified subscription.
//
// Generated from API version 2019-07-01-preview
//   - options - PipelinesClientListBySubscriptionOptions contains the optional parameters for the PipelinesClient.NewListBySubscriptionPager
//     method.
func (client *PipelinesClient) NewListBySubscriptionPager(options *PipelinesClientListBySubscriptionOptions) *runtime.Pager[PipelinesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[PipelinesClientListBySubscriptionResponse]{
		More: func(page PipelinesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PipelinesClientListBySubscriptionResponse) (PipelinesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PipelinesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PipelinesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PipelinesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *PipelinesClient) listBySubscriptionCreateRequest(ctx context.Context, options *PipelinesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DevOps/pipelines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *PipelinesClient) listBySubscriptionHandleResponse(resp *http.Response) (PipelinesClientListBySubscriptionResponse, error) {
	result := PipelinesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineListResult); err != nil {
		return PipelinesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Updates the properties of an Azure Pipeline. Currently, only tags can be updated.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-07-01-preview
//   - resourceGroupName - Name of the resource group within the Azure subscription.
//   - pipelineName - The name of the Azure Pipeline resource.
//   - updateOperationParameters - The request payload containing the properties to update in the Azure Pipeline.
//   - options - PipelinesClientUpdateOptions contains the optional parameters for the PipelinesClient.Update method.
func (client *PipelinesClient) Update(ctx context.Context, resourceGroupName string, pipelineName string, updateOperationParameters PipelineUpdateParameters, options *PipelinesClientUpdateOptions) (PipelinesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, pipelineName, updateOperationParameters, options)
	if err != nil {
		return PipelinesClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PipelinesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PipelinesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *PipelinesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, pipelineName string, updateOperationParameters PipelineUpdateParameters, options *PipelinesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevOps/pipelines/{pipelineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, updateOperationParameters)
}

// updateHandleResponse handles the Update response.
func (client *PipelinesClient) updateHandleResponse(resp *http.Response) (PipelinesClientUpdateResponse, error) {
	result := PipelinesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Pipeline); err != nil {
		return PipelinesClientUpdateResponse{}, err
	}
	return result, nil
}
