//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armguestconfiguration

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

// AssignmentsVMSSClient contains the methods for the GuestConfigurationAssignmentsVMSS group.
// Don't use this type directly, use NewAssignmentsVMSSClient() instead.
type AssignmentsVMSSClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAssignmentsVMSSClient creates a new instance of AssignmentsVMSSClient with the specified values.
//   - subscriptionID - Subscription ID which uniquely identify Microsoft Azure subscription. The subscription ID forms part of
//     the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAssignmentsVMSSClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AssignmentsVMSSClient, error) {
	cl, err := arm.NewClient(moduleName+".AssignmentsVMSSClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AssignmentsVMSSClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Delete - Delete a guest configuration assignment for VMSS
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-25
//   - resourceGroupName - The resource group name.
//   - vmssName - The name of the virtual machine scale set.
//   - name - The guest configuration assignment name.
//   - options - AssignmentsVMSSClientDeleteOptions contains the optional parameters for the AssignmentsVMSSClient.Delete method.
func (client *AssignmentsVMSSClient) Delete(ctx context.Context, resourceGroupName string, vmssName string, name string, options *AssignmentsVMSSClientDeleteOptions) (AssignmentsVMSSClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vmssName, name, options)
	if err != nil {
		return AssignmentsVMSSClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssignmentsVMSSClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AssignmentsVMSSClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *AssignmentsVMSSClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vmssName string, name string, options *AssignmentsVMSSClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmssName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmssName == "" {
		return nil, errors.New("parameter vmssName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmssName}", url.PathEscape(vmssName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *AssignmentsVMSSClient) deleteHandleResponse(resp *http.Response) (AssignmentsVMSSClientDeleteResponse, error) {
	result := AssignmentsVMSSClientDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Assignment); err != nil {
		return AssignmentsVMSSClientDeleteResponse{}, err
	}
	return result, nil
}

// Get - Get information about a guest configuration assignment for VMSS
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-01-25
//   - resourceGroupName - The resource group name.
//   - vmssName - The name of the virtual machine scale set.
//   - name - The guest configuration assignment name.
//   - options - AssignmentsVMSSClientGetOptions contains the optional parameters for the AssignmentsVMSSClient.Get method.
func (client *AssignmentsVMSSClient) Get(ctx context.Context, resourceGroupName string, vmssName string, name string, options *AssignmentsVMSSClientGetOptions) (AssignmentsVMSSClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vmssName, name, options)
	if err != nil {
		return AssignmentsVMSSClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssignmentsVMSSClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssignmentsVMSSClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AssignmentsVMSSClient) getCreateRequest(ctx context.Context, resourceGroupName string, vmssName string, name string, options *AssignmentsVMSSClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmssName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vmssName == "" {
		return nil, errors.New("parameter vmssName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmssName}", url.PathEscape(vmssName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AssignmentsVMSSClient) getHandleResponse(resp *http.Response) (AssignmentsVMSSClientGetResponse, error) {
	result := AssignmentsVMSSClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Assignment); err != nil {
		return AssignmentsVMSSClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List all guest configuration assignments for VMSS.
//
// Generated from API version 2022-01-25
//   - resourceGroupName - The resource group name.
//   - vmssName - The name of the virtual machine scale set.
//   - options - AssignmentsVMSSClientListOptions contains the optional parameters for the AssignmentsVMSSClient.NewListPager
//     method.
func (client *AssignmentsVMSSClient) NewListPager(resourceGroupName string, vmssName string, options *AssignmentsVMSSClientListOptions) *runtime.Pager[AssignmentsVMSSClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AssignmentsVMSSClientListResponse]{
		More: func(page AssignmentsVMSSClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *AssignmentsVMSSClientListResponse) (AssignmentsVMSSClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, resourceGroupName, vmssName, options)
			if err != nil {
				return AssignmentsVMSSClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AssignmentsVMSSClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AssignmentsVMSSClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AssignmentsVMSSClient) listCreateRequest(ctx context.Context, resourceGroupName string, vmssName string, options *AssignmentsVMSSClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmssName}/providers/Microsoft.GuestConfiguration/guestConfigurationAssignments"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if vmssName == "" {
		return nil, errors.New("parameter vmssName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vmssName}", url.PathEscape(vmssName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-25")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AssignmentsVMSSClient) listHandleResponse(resp *http.Response) (AssignmentsVMSSClientListResponse, error) {
	result := AssignmentsVMSSClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssignmentList); err != nil {
		return AssignmentsVMSSClientListResponse{}, err
	}
	return result, nil
}
