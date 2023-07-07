//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// DO NOT EDIT.

package armazuresphere

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// DeviceGroupsClient contains the methods for the DeviceGroups group.
// Don't use this type directly, use NewDeviceGroupsClient() instead.
type DeviceGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDeviceGroupsClient creates a new instance of DeviceGroupsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDeviceGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DeviceGroupsClient, error) {
	cl, err := arm.NewClient(moduleName+".DeviceGroupsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DeviceGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginClaimDevices - Bulk claims the devices. Use '.unassigned' or '.default' for the device group and product names when
// bulk claiming devices to a catalog only.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - catalogName - Name of catalog
//   - productName - Name of product.
//   - deviceGroupName - Name of device group.
//   - claimDevicesRequest - Bulk claim devices request body.
//   - options - DeviceGroupsClientBeginClaimDevicesOptions contains the optional parameters for the DeviceGroupsClient.BeginClaimDevices
//     method.
func (client *DeviceGroupsClient) BeginClaimDevices(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, claimDevicesRequest ClaimDevicesRequest, options *DeviceGroupsClientBeginClaimDevicesOptions) (*runtime.Poller[DeviceGroupsClientClaimDevicesResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.claimDevices(ctx, resourceGroupName, catalogName, productName, deviceGroupName, claimDevicesRequest, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DeviceGroupsClientClaimDevicesResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DeviceGroupsClientClaimDevicesResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// ClaimDevices - Bulk claims the devices. Use '.unassigned' or '.default' for the device group and product names when bulk
// claiming devices to a catalog only.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
func (client *DeviceGroupsClient) claimDevices(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, claimDevicesRequest ClaimDevicesRequest, options *DeviceGroupsClientBeginClaimDevicesOptions) (*http.Response, error) {
	req, err := client.claimDevicesCreateRequest(ctx, resourceGroupName, catalogName, productName, deviceGroupName, claimDevicesRequest, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// claimDevicesCreateRequest creates the ClaimDevices request.
func (client *DeviceGroupsClient) claimDevicesCreateRequest(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, claimDevicesRequest ClaimDevicesRequest, options *DeviceGroupsClientBeginClaimDevicesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureSphere/catalogs/{catalogName}/products/{productName}/deviceGroups/{deviceGroupName}/claimDevices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, claimDevicesRequest)
}

// CountDevices - Counts devices in device group. '.default' and '.unassigned' are system defined values and cannot be used
// for product or device group name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - catalogName - Name of catalog
//   - productName - Name of product.
//   - deviceGroupName - Name of device group.
//   - options - DeviceGroupsClientCountDevicesOptions contains the optional parameters for the DeviceGroupsClient.CountDevices
//     method.
func (client *DeviceGroupsClient) CountDevices(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, options *DeviceGroupsClientCountDevicesOptions) (DeviceGroupsClientCountDevicesResponse, error) {
	req, err := client.countDevicesCreateRequest(ctx, resourceGroupName, catalogName, productName, deviceGroupName, options)
	if err != nil {
		return DeviceGroupsClientCountDevicesResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeviceGroupsClientCountDevicesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeviceGroupsClientCountDevicesResponse{}, runtime.NewResponseError(resp)
	}
	return client.countDevicesHandleResponse(resp)
}

// countDevicesCreateRequest creates the CountDevices request.
func (client *DeviceGroupsClient) countDevicesCreateRequest(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, options *DeviceGroupsClientCountDevicesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureSphere/catalogs/{catalogName}/products/{productName}/deviceGroups/{deviceGroupName}/countDevices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// countDevicesHandleResponse handles the CountDevices response.
func (client *DeviceGroupsClient) countDevicesHandleResponse(resp *http.Response) (DeviceGroupsClientCountDevicesResponse, error) {
	result := DeviceGroupsClientCountDevicesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CountDeviceResponse); err != nil {
		return DeviceGroupsClientCountDevicesResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Create a DeviceGroup. '.default' and '.unassigned' are system defined values and cannot be used for
// product or device group name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - catalogName - Name of catalog
//   - productName - Name of product.
//   - deviceGroupName - Name of device group.
//   - resource - Resource create parameters.
//   - options - DeviceGroupsClientBeginCreateOrUpdateOptions contains the optional parameters for the DeviceGroupsClient.BeginCreateOrUpdate
//     method.
func (client *DeviceGroupsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, resource DeviceGroup, options *DeviceGroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[DeviceGroupsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, catalogName, productName, deviceGroupName, resource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DeviceGroupsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DeviceGroupsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a DeviceGroup. '.default' and '.unassigned' are system defined values and cannot be used for product
// or device group name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
func (client *DeviceGroupsClient) createOrUpdate(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, resource DeviceGroup, options *DeviceGroupsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, catalogName, productName, deviceGroupName, resource, options)
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
func (client *DeviceGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, resource DeviceGroup, options *DeviceGroupsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureSphere/catalogs/{catalogName}/products/{productName}/deviceGroups/{deviceGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, resource)
}

// BeginDelete - Delete a DeviceGroup. '.default' and '.unassigned' are system defined values and cannot be used for product
// or device group name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - catalogName - Name of catalog
//   - productName - Name of product.
//   - deviceGroupName - Name of device group.
//   - options - DeviceGroupsClientBeginDeleteOptions contains the optional parameters for the DeviceGroupsClient.BeginDelete
//     method.
func (client *DeviceGroupsClient) BeginDelete(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, options *DeviceGroupsClientBeginDeleteOptions) (*runtime.Poller[DeviceGroupsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, catalogName, productName, deviceGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DeviceGroupsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DeviceGroupsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete a DeviceGroup. '.default' and '.unassigned' are system defined values and cannot be used for product or
// device group name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
func (client *DeviceGroupsClient) deleteOperation(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, options *DeviceGroupsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, catalogName, productName, deviceGroupName, options)
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
func (client *DeviceGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, options *DeviceGroupsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureSphere/catalogs/{catalogName}/products/{productName}/deviceGroups/{deviceGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a DeviceGroup. '.default' and '.unassigned' are system defined values and cannot be used for product or device
// group name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - catalogName - Name of catalog
//   - productName - Name of product.
//   - deviceGroupName - Name of device group.
//   - options - DeviceGroupsClientGetOptions contains the optional parameters for the DeviceGroupsClient.Get method.
func (client *DeviceGroupsClient) Get(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, options *DeviceGroupsClientGetOptions) (DeviceGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, catalogName, productName, deviceGroupName, options)
	if err != nil {
		return DeviceGroupsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeviceGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DeviceGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DeviceGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, options *DeviceGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureSphere/catalogs/{catalogName}/products/{productName}/deviceGroups/{deviceGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DeviceGroupsClient) getHandleResponse(resp *http.Response) (DeviceGroupsClientGetResponse, error) {
	result := DeviceGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceGroup); err != nil {
		return DeviceGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByProductPager - List DeviceGroup resources by Product. '.default' and '.unassigned' are system defined values and
// cannot be used for product name.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - catalogName - Name of catalog
//   - productName - Name of product.
//   - options - DeviceGroupsClientListByProductOptions contains the optional parameters for the DeviceGroupsClient.NewListByProductPager
//     method.
func (client *DeviceGroupsClient) NewListByProductPager(resourceGroupName string, catalogName string, productName string, options *DeviceGroupsClientListByProductOptions) *runtime.Pager[DeviceGroupsClientListByProductResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeviceGroupsClientListByProductResponse]{
		More: func(page DeviceGroupsClientListByProductResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeviceGroupsClientListByProductResponse) (DeviceGroupsClientListByProductResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByProductCreateRequest(ctx, resourceGroupName, catalogName, productName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DeviceGroupsClientListByProductResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DeviceGroupsClientListByProductResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DeviceGroupsClientListByProductResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByProductHandleResponse(resp)
		},
	})
}

// listByProductCreateRequest creates the ListByProduct request.
func (client *DeviceGroupsClient) listByProductCreateRequest(ctx context.Context, resourceGroupName string, catalogName string, productName string, options *DeviceGroupsClientListByProductOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureSphere/catalogs/{catalogName}/products/{productName}/deviceGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Maxpagesize != nil {
		reqQP.Set("$maxpagesize", strconv.FormatInt(int64(*options.Maxpagesize), 10))
	}
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProductHandleResponse handles the ListByProduct response.
func (client *DeviceGroupsClient) listByProductHandleResponse(resp *http.Response) (DeviceGroupsClientListByProductResponse, error) {
	result := DeviceGroupsClientListByProductResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceGroupListResult); err != nil {
		return DeviceGroupsClientListByProductResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a DeviceGroup. '.default' and '.unassigned' are system defined values and cannot be used for product
// or device group name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - catalogName - Name of catalog
//   - productName - Name of product.
//   - deviceGroupName - Name of device group.
//   - properties - The resource properties to be updated.
//   - options - DeviceGroupsClientBeginUpdateOptions contains the optional parameters for the DeviceGroupsClient.BeginUpdate
//     method.
func (client *DeviceGroupsClient) BeginUpdate(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, properties DeviceGroupUpdate, options *DeviceGroupsClientBeginUpdateOptions) (*runtime.Poller[DeviceGroupsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, catalogName, productName, deviceGroupName, properties, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DeviceGroupsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DeviceGroupsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update a DeviceGroup. '.default' and '.unassigned' are system defined values and cannot be used for product or
// device group name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01-preview
func (client *DeviceGroupsClient) update(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, properties DeviceGroupUpdate, options *DeviceGroupsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, catalogName, productName, deviceGroupName, properties, options)
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

// updateCreateRequest creates the Update request.
func (client *DeviceGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, catalogName string, productName string, deviceGroupName string, properties DeviceGroupUpdate, options *DeviceGroupsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureSphere/catalogs/{catalogName}/products/{productName}/deviceGroups/{deviceGroupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if productName == "" {
		return nil, errors.New("parameter productName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productName}", url.PathEscape(productName))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, properties)
}