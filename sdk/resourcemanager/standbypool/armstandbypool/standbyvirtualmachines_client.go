// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armstandbypool

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

// StandbyVirtualMachinesClient contains the methods for the Microsoft.StandbyPool namespace.
// Don't use this type directly, use NewStandbyVirtualMachinesClient() instead.
type StandbyVirtualMachinesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewStandbyVirtualMachinesClient creates a new instance of StandbyVirtualMachinesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewStandbyVirtualMachinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*StandbyVirtualMachinesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &StandbyVirtualMachinesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a StandbyVirtualMachineResource
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - standbyVirtualMachinePoolName - Name of the standby virtual machine pool
//   - standbyVirtualMachineName - Name of the standby virtual machine
//   - options - StandbyVirtualMachinesClientGetOptions contains the optional parameters for the StandbyVirtualMachinesClient.Get
//     method.
func (client *StandbyVirtualMachinesClient) Get(ctx context.Context, resourceGroupName string, standbyVirtualMachinePoolName string, standbyVirtualMachineName string, options *StandbyVirtualMachinesClientGetOptions) (StandbyVirtualMachinesClientGetResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "StandbyVirtualMachinesClient.Get")
	req, err := client.getCreateRequest(ctx, resourceGroupName, standbyVirtualMachinePoolName, standbyVirtualMachineName, options)
	if err != nil {
		return StandbyVirtualMachinesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StandbyVirtualMachinesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StandbyVirtualMachinesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *StandbyVirtualMachinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, standbyVirtualMachinePoolName string, standbyVirtualMachineName string, options *StandbyVirtualMachinesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/{standbyVirtualMachinePoolName}/standbyVirtualMachines/{standbyVirtualMachineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if standbyVirtualMachinePoolName == "" {
		return nil, errors.New("parameter standbyVirtualMachinePoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{standbyVirtualMachinePoolName}", url.PathEscape(standbyVirtualMachinePoolName))
	if standbyVirtualMachineName == "" {
		return nil, errors.New("parameter standbyVirtualMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{standbyVirtualMachineName}", url.PathEscape(standbyVirtualMachineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StandbyVirtualMachinesClient) getHandleResponse(resp *http.Response) (StandbyVirtualMachinesClientGetResponse, error) {
	result := StandbyVirtualMachinesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StandbyVirtualMachineResource); err != nil {
		return StandbyVirtualMachinesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByStandbyVirtualMachinePoolResourcePager - List StandbyVirtualMachineResource resources by StandbyVirtualMachinePoolResource
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - standbyVirtualMachinePoolName - Name of the standby virtual machine pool
//   - options - StandbyVirtualMachinesClientListByStandbyVirtualMachinePoolResourceOptions contains the optional parameters for
//     the StandbyVirtualMachinesClient.NewListByStandbyVirtualMachinePoolResourcePager method.
func (client *StandbyVirtualMachinesClient) NewListByStandbyVirtualMachinePoolResourcePager(resourceGroupName string, standbyVirtualMachinePoolName string, options *StandbyVirtualMachinesClientListByStandbyVirtualMachinePoolResourceOptions) *runtime.Pager[StandbyVirtualMachinesClientListByStandbyVirtualMachinePoolResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[StandbyVirtualMachinesClientListByStandbyVirtualMachinePoolResourceResponse]{
		More: func(page StandbyVirtualMachinesClientListByStandbyVirtualMachinePoolResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StandbyVirtualMachinesClientListByStandbyVirtualMachinePoolResourceResponse) (StandbyVirtualMachinesClientListByStandbyVirtualMachinePoolResourceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "StandbyVirtualMachinesClient.NewListByStandbyVirtualMachinePoolResourcePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByStandbyVirtualMachinePoolResourceCreateRequest(ctx, resourceGroupName, standbyVirtualMachinePoolName, options)
			}, nil)
			if err != nil {
				return StandbyVirtualMachinesClientListByStandbyVirtualMachinePoolResourceResponse{}, err
			}
			return client.listByStandbyVirtualMachinePoolResourceHandleResponse(resp)
		},
	})
}

// listByStandbyVirtualMachinePoolResourceCreateRequest creates the ListByStandbyVirtualMachinePoolResource request.
func (client *StandbyVirtualMachinesClient) listByStandbyVirtualMachinePoolResourceCreateRequest(ctx context.Context, resourceGroupName string, standbyVirtualMachinePoolName string, options *StandbyVirtualMachinesClientListByStandbyVirtualMachinePoolResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StandbyPool/standbyVirtualMachinePools/{standbyVirtualMachinePoolName}/standbyVirtualMachines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if standbyVirtualMachinePoolName == "" {
		return nil, errors.New("parameter standbyVirtualMachinePoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{standbyVirtualMachinePoolName}", url.PathEscape(standbyVirtualMachinePoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByStandbyVirtualMachinePoolResourceHandleResponse handles the ListByStandbyVirtualMachinePoolResource response.
func (client *StandbyVirtualMachinesClient) listByStandbyVirtualMachinePoolResourceHandleResponse(resp *http.Response) (StandbyVirtualMachinesClientListByStandbyVirtualMachinePoolResourceResponse, error) {
	result := StandbyVirtualMachinesClientListByStandbyVirtualMachinePoolResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StandbyVirtualMachineResourceListResult); err != nil {
		return StandbyVirtualMachinesClientListByStandbyVirtualMachinePoolResourceResponse{}, err
	}
	return result, nil
}
