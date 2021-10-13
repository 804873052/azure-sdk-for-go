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
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// StandardsClient contains the methods for the Standards group.
// Don't use this type directly, use NewStandardsClient() instead.
type StandardsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewStandardsClient creates a new instance of StandardsClient with the specified values.
func NewStandardsClient(con *arm.Connection, subscriptionID string) *StandardsClient {
	return &StandardsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// CreateOrUpdate - Create a security standard on the given scope. Available only for custom standards. Will create/update the required standard definitions.
// If the operation fails it returns the *CloudError error type.
func (client *StandardsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, standardID string, standard Standard, options *StandardsCreateOrUpdateOptions) (StandardsCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, standardID, standard, options)
	if err != nil {
		return StandardsCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StandardsCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return StandardsCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *StandardsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, standardID string, standard Standard, options *StandardsCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/standards/{standardId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if standardID == "" {
		return nil, errors.New("parameter standardID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{standardId}", url.PathEscape(standardID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, standard)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *StandardsClient) createOrUpdateHandleResponse(resp *http.Response) (StandardsCreateOrUpdateResponse, error) {
	result := StandardsCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Standard); err != nil {
		return StandardsCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *StandardsClient) createOrUpdateHandleError(resp *http.Response) error {
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

// Delete - Delete a security standard on a scope.
// If the operation fails it returns the *CloudError error type.
func (client *StandardsClient) Delete(ctx context.Context, resourceGroupName string, standardID string, options *StandardsDeleteOptions) (StandardsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, standardID, options)
	if err != nil {
		return StandardsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StandardsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return StandardsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return StandardsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *StandardsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, standardID string, options *StandardsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/standards/{standardId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if standardID == "" {
		return nil, errors.New("parameter standardID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{standardId}", url.PathEscape(standardID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *StandardsClient) deleteHandleError(resp *http.Response) error {
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

// Get - Get a specific security standard for the requested scope
// If the operation fails it returns the *CloudError error type.
func (client *StandardsClient) Get(ctx context.Context, resourceGroupName string, standardID string, options *StandardsGetOptions) (StandardsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, standardID, options)
	if err != nil {
		return StandardsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StandardsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StandardsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *StandardsClient) getCreateRequest(ctx context.Context, resourceGroupName string, standardID string, options *StandardsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/standards/{standardId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if standardID == "" {
		return nil, errors.New("parameter standardID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{standardId}", url.PathEscape(standardID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StandardsClient) getHandleResponse(resp *http.Response) (StandardsGetResponse, error) {
	result := StandardsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Standard); err != nil {
		return StandardsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *StandardsClient) getHandleError(resp *http.Response) error {
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

// List - Get security standards on all your resources inside a scope
// If the operation fails it returns the *CloudError error type.
func (client *StandardsClient) List(resourceGroupName string, options *StandardsListOptions) *StandardsListPager {
	return &StandardsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp StandardsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.StandardList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *StandardsClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *StandardsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/standards"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *StandardsClient) listHandleResponse(resp *http.Response) (StandardsListResponse, error) {
	result := StandardsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.StandardList); err != nil {
		return StandardsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *StandardsClient) listHandleError(resp *http.Response) error {
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

// ListBySubscription - Get a list of all relevant security standards over a subscription level scope.
// If the operation fails it returns the *CloudError error type.
func (client *StandardsClient) ListBySubscription(options *StandardsListBySubscriptionOptions) *StandardsListBySubscriptionPager {
	return &StandardsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp StandardsListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.StandardList.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *StandardsClient) listBySubscriptionCreateRequest(ctx context.Context, options *StandardsListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/standards"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *StandardsClient) listBySubscriptionHandleResponse(resp *http.Response) (StandardsListBySubscriptionResponse, error) {
	result := StandardsListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.StandardList); err != nil {
		return StandardsListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *StandardsClient) listBySubscriptionHandleError(resp *http.Response) error {
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
