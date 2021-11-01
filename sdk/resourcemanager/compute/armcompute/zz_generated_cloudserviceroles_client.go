//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// CloudServiceRolesClient contains the methods for the CloudServiceRoles group.
// Don't use this type directly, use NewCloudServiceRolesClient() instead.
type CloudServiceRolesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewCloudServiceRolesClient creates a new instance of CloudServiceRolesClient with the specified values.
func NewCloudServiceRolesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *CloudServiceRolesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &CloudServiceRolesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Gets a role from a cloud service.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServiceRolesClient) Get(ctx context.Context, roleName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRolesGetOptions) (CloudServiceRolesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, roleName, resourceGroupName, cloudServiceName, options)
	if err != nil {
		return CloudServiceRolesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CloudServiceRolesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CloudServiceRolesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CloudServiceRolesClient) getCreateRequest(ctx context.Context, roleName string, resourceGroupName string, cloudServiceName string, options *CloudServiceRolesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roles/{roleName}"
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CloudServiceRolesClient) getHandleResponse(resp *http.Response) (CloudServiceRolesGetResponse, error) {
	result := CloudServiceRolesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CloudServiceRole); err != nil {
		return CloudServiceRolesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *CloudServiceRolesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Gets a list of all roles in a cloud service. Use nextLink property in the response to get the next page of roles. Do this till nextLink is null
// to fetch all the roles.
// If the operation fails it returns the *CloudError error type.
func (client *CloudServiceRolesClient) List(resourceGroupName string, cloudServiceName string, options *CloudServiceRolesListOptions) *CloudServiceRolesListPager {
	return &CloudServiceRolesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, cloudServiceName, options)
		},
		advancer: func(ctx context.Context, resp CloudServiceRolesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CloudServiceRoleListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *CloudServiceRolesClient) listCreateRequest(ctx context.Context, resourceGroupName string, cloudServiceName string, options *CloudServiceRolesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/cloudServices/{cloudServiceName}/roles"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudServiceName == "" {
		return nil, errors.New("parameter cloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudServiceName}", url.PathEscape(cloudServiceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CloudServiceRolesClient) listHandleResponse(resp *http.Response) (CloudServiceRolesListResponse, error) {
	result := CloudServiceRolesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CloudServiceRoleListResult); err != nil {
		return CloudServiceRolesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *CloudServiceRolesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
