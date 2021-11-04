//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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

// WorkspaceSettingsClient contains the methods for the WorkspaceSettings group.
// Don't use this type directly, use NewWorkspaceSettingsClient() instead.
type WorkspaceSettingsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewWorkspaceSettingsClient creates a new instance of WorkspaceSettingsClient with the specified values.
func NewWorkspaceSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WorkspaceSettingsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &WorkspaceSettingsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Create - creating settings about where we should store your security data and logs
// If the operation fails it returns the *CloudError error type.
func (client *WorkspaceSettingsClient) Create(ctx context.Context, workspaceSettingName string, workspaceSetting WorkspaceSetting, options *WorkspaceSettingsCreateOptions) (WorkspaceSettingsCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, workspaceSettingName, workspaceSetting, options)
	if err != nil {
		return WorkspaceSettingsCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceSettingsCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceSettingsCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *WorkspaceSettingsClient) createCreateRequest(ctx context.Context, workspaceSettingName string, workspaceSetting WorkspaceSetting, options *WorkspaceSettingsCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings/{workspaceSettingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceSettingName == "" {
		return nil, errors.New("parameter workspaceSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceSettingName}", url.PathEscape(workspaceSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, workspaceSetting)
}

// createHandleResponse handles the Create response.
func (client *WorkspaceSettingsClient) createHandleResponse(resp *http.Response) (WorkspaceSettingsCreateResponse, error) {
	result := WorkspaceSettingsCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceSetting); err != nil {
		return WorkspaceSettingsCreateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *WorkspaceSettingsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes the custom workspace settings for this subscription. new VMs will report to the default workspace
// If the operation fails it returns the *CloudError error type.
func (client *WorkspaceSettingsClient) Delete(ctx context.Context, workspaceSettingName string, options *WorkspaceSettingsDeleteOptions) (WorkspaceSettingsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, workspaceSettingName, options)
	if err != nil {
		return WorkspaceSettingsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceSettingsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return WorkspaceSettingsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return WorkspaceSettingsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspaceSettingsClient) deleteCreateRequest(ctx context.Context, workspaceSettingName string, options *WorkspaceSettingsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings/{workspaceSettingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceSettingName == "" {
		return nil, errors.New("parameter workspaceSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceSettingName}", url.PathEscape(workspaceSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *WorkspaceSettingsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Settings about where we should store your security data and logs. If the result is empty, it means that no custom-workspace configuration was set
// If the operation fails it returns the *CloudError error type.
func (client *WorkspaceSettingsClient) Get(ctx context.Context, workspaceSettingName string, options *WorkspaceSettingsGetOptions) (WorkspaceSettingsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, workspaceSettingName, options)
	if err != nil {
		return WorkspaceSettingsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceSettingsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceSettingsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspaceSettingsClient) getCreateRequest(ctx context.Context, workspaceSettingName string, options *WorkspaceSettingsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings/{workspaceSettingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceSettingName == "" {
		return nil, errors.New("parameter workspaceSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceSettingName}", url.PathEscape(workspaceSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspaceSettingsClient) getHandleResponse(resp *http.Response) (WorkspaceSettingsGetResponse, error) {
	result := WorkspaceSettingsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceSetting); err != nil {
		return WorkspaceSettingsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *WorkspaceSettingsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Settings about where we should store your security data and logs. If the result is empty, it means that no custom-workspace configuration was
// set
// If the operation fails it returns the *CloudError error type.
func (client *WorkspaceSettingsClient) List(options *WorkspaceSettingsListOptions) *WorkspaceSettingsListPager {
	return &WorkspaceSettingsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp WorkspaceSettingsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WorkspaceSettingList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *WorkspaceSettingsClient) listCreateRequest(ctx context.Context, options *WorkspaceSettingsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkspaceSettingsClient) listHandleResponse(resp *http.Response) (WorkspaceSettingsListResponse, error) {
	result := WorkspaceSettingsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceSettingList); err != nil {
		return WorkspaceSettingsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *WorkspaceSettingsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Settings about where we should store your security data and logs
// If the operation fails it returns the *CloudError error type.
func (client *WorkspaceSettingsClient) Update(ctx context.Context, workspaceSettingName string, workspaceSetting WorkspaceSetting, options *WorkspaceSettingsUpdateOptions) (WorkspaceSettingsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, workspaceSettingName, workspaceSetting, options)
	if err != nil {
		return WorkspaceSettingsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspaceSettingsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceSettingsUpdateResponse{}, client.updateHandleError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *WorkspaceSettingsClient) updateCreateRequest(ctx context.Context, workspaceSettingName string, workspaceSetting WorkspaceSetting, options *WorkspaceSettingsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/workspaceSettings/{workspaceSettingName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceSettingName == "" {
		return nil, errors.New("parameter workspaceSettingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceSettingName}", url.PathEscape(workspaceSettingName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, workspaceSetting)
}

// updateHandleResponse handles the Update response.
func (client *WorkspaceSettingsClient) updateHandleResponse(resp *http.Response) (WorkspaceSettingsUpdateResponse, error) {
	result := WorkspaceSettingsUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceSetting); err != nil {
		return WorkspaceSettingsUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// updateHandleError handles the Update error response.
func (client *WorkspaceSettingsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
