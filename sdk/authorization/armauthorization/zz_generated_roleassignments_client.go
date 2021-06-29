// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// RoleAssignmentsClient contains the methods for the RoleAssignments group.
// Don't use this type directly, use NewRoleAssignmentsClient() instead.
type RoleAssignmentsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewRoleAssignmentsClient creates a new instance of RoleAssignmentsClient with the specified values.
func NewRoleAssignmentsClient(con *armcore.Connection, subscriptionID string) *RoleAssignmentsClient {
	return &RoleAssignmentsClient{con: con, subscriptionID: subscriptionID}
}

// Create - Create or update a role assignment by scope and name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) Create(ctx context.Context, scope string, roleAssignmentName string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsCreateOptions) (RoleAssignmentResponse, error) {
	req, err := client.createCreateRequest(ctx, scope, roleAssignmentName, parameters, options)
	if err != nil {
		return RoleAssignmentResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return RoleAssignmentResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *RoleAssignmentsClient) createCreateRequest(ctx context.Context, scope string, roleAssignmentName string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsCreateOptions) (*azcore.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentName == "" {
		return nil, errors.New("parameter roleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentName}", roleAssignmentName)
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createHandleResponse handles the Create response.
func (client *RoleAssignmentsClient) createHandleResponse(resp *azcore.Response) (RoleAssignmentResponse, error) {
	var val *RoleAssignment
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RoleAssignmentResponse{}, err
	}
	return RoleAssignmentResponse{RawResponse: resp.Response, RoleAssignment: val}, nil
}

// createHandleError handles the Create error response.
func (client *RoleAssignmentsClient) createHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// CreateByID - Create or update a role assignment by ID.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) CreateByID(ctx context.Context, roleAssignmentID string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsCreateByIDOptions) (RoleAssignmentResponse, error) {
	req, err := client.createByIDCreateRequest(ctx, roleAssignmentID, parameters, options)
	if err != nil {
		return RoleAssignmentResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return RoleAssignmentResponse{}, client.createByIDHandleError(resp)
	}
	return client.createByIDHandleResponse(resp)
}

// createByIDCreateRequest creates the CreateByID request.
func (client *RoleAssignmentsClient) createByIDCreateRequest(ctx context.Context, roleAssignmentID string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsCreateByIDOptions) (*azcore.Request, error) {
	urlPath := "/{roleAssignmentId}"
	if roleAssignmentID == "" {
		return nil, errors.New("parameter roleAssignmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentId}", roleAssignmentID)
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createByIDHandleResponse handles the CreateByID response.
func (client *RoleAssignmentsClient) createByIDHandleResponse(resp *azcore.Response) (RoleAssignmentResponse, error) {
	var val *RoleAssignment
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RoleAssignmentResponse{}, err
	}
	return RoleAssignmentResponse{RawResponse: resp.Response, RoleAssignment: val}, nil
}

// createByIDHandleError handles the CreateByID error response.
func (client *RoleAssignmentsClient) createByIDHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Delete - Delete a role assignment by scope and name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) Delete(ctx context.Context, scope string, roleAssignmentName string, options *RoleAssignmentsDeleteOptions) (RoleAssignmentResponse, error) {
	req, err := client.deleteCreateRequest(ctx, scope, roleAssignmentName, options)
	if err != nil {
		return RoleAssignmentResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return RoleAssignmentResponse{}, client.deleteHandleError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *RoleAssignmentsClient) deleteCreateRequest(ctx context.Context, scope string, roleAssignmentName string, options *RoleAssignmentsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentName == "" {
		return nil, errors.New("parameter roleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentName}", roleAssignmentName)
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	if options != nil && options.TenantID != nil {
		reqQP.Set("tenantId", *options.TenantID)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *RoleAssignmentsClient) deleteHandleResponse(resp *azcore.Response) (RoleAssignmentResponse, error) {
	var val *RoleAssignment
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RoleAssignmentResponse{}, err
	}
	return RoleAssignmentResponse{RawResponse: resp.Response, RoleAssignment: val}, nil
}

// deleteHandleError handles the Delete error response.
func (client *RoleAssignmentsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// DeleteByID - Delete a role assignment by ID.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) DeleteByID(ctx context.Context, roleAssignmentID string, options *RoleAssignmentsDeleteByIDOptions) (RoleAssignmentResponse, error) {
	req, err := client.deleteByIDCreateRequest(ctx, roleAssignmentID, options)
	if err != nil {
		return RoleAssignmentResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return RoleAssignmentResponse{}, client.deleteByIDHandleError(resp)
	}
	return client.deleteByIDHandleResponse(resp)
}

// deleteByIDCreateRequest creates the DeleteByID request.
func (client *RoleAssignmentsClient) deleteByIDCreateRequest(ctx context.Context, roleAssignmentID string, options *RoleAssignmentsDeleteByIDOptions) (*azcore.Request, error) {
	urlPath := "/{roleAssignmentId}"
	if roleAssignmentID == "" {
		return nil, errors.New("parameter roleAssignmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentId}", roleAssignmentID)
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	if options != nil && options.TenantID != nil {
		reqQP.Set("tenantId", *options.TenantID)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteByIDHandleResponse handles the DeleteByID response.
func (client *RoleAssignmentsClient) deleteByIDHandleResponse(resp *azcore.Response) (RoleAssignmentResponse, error) {
	var val *RoleAssignment
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RoleAssignmentResponse{}, err
	}
	return RoleAssignmentResponse{RawResponse: resp.Response, RoleAssignment: val}, nil
}

// deleteByIDHandleError handles the DeleteByID error response.
func (client *RoleAssignmentsClient) deleteByIDHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Get a role assignment by scope and name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) Get(ctx context.Context, scope string, roleAssignmentName string, options *RoleAssignmentsGetOptions) (RoleAssignmentResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, roleAssignmentName, options)
	if err != nil {
		return RoleAssignmentResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RoleAssignmentResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoleAssignmentsClient) getCreateRequest(ctx context.Context, scope string, roleAssignmentName string, options *RoleAssignmentsGetOptions) (*azcore.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentName == "" {
		return nil, errors.New("parameter roleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentName}", roleAssignmentName)
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	if options != nil && options.TenantID != nil {
		reqQP.Set("tenantId", *options.TenantID)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleAssignmentsClient) getHandleResponse(resp *azcore.Response) (RoleAssignmentResponse, error) {
	var val *RoleAssignment
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RoleAssignmentResponse{}, err
	}
	return RoleAssignmentResponse{RawResponse: resp.Response, RoleAssignment: val}, nil
}

// getHandleError handles the Get error response.
func (client *RoleAssignmentsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetByID - Get a role assignment by ID.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) GetByID(ctx context.Context, roleAssignmentID string, options *RoleAssignmentsGetByIDOptions) (RoleAssignmentResponse, error) {
	req, err := client.getByIDCreateRequest(ctx, roleAssignmentID, options)
	if err != nil {
		return RoleAssignmentResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return RoleAssignmentResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return RoleAssignmentResponse{}, client.getByIDHandleError(resp)
	}
	return client.getByIDHandleResponse(resp)
}

// getByIDCreateRequest creates the GetByID request.
func (client *RoleAssignmentsClient) getByIDCreateRequest(ctx context.Context, roleAssignmentID string, options *RoleAssignmentsGetByIDOptions) (*azcore.Request, error) {
	urlPath := "/{roleAssignmentId}"
	if roleAssignmentID == "" {
		return nil, errors.New("parameter roleAssignmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentId}", roleAssignmentID)
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	if options != nil && options.TenantID != nil {
		reqQP.Set("tenantId", *options.TenantID)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getByIDHandleResponse handles the GetByID response.
func (client *RoleAssignmentsClient) getByIDHandleResponse(resp *azcore.Response) (RoleAssignmentResponse, error) {
	var val *RoleAssignment
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RoleAssignmentResponse{}, err
	}
	return RoleAssignmentResponse{RawResponse: resp.Response, RoleAssignment: val}, nil
}

// getByIDHandleError handles the GetByID error response.
func (client *RoleAssignmentsClient) getByIDHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListForResource - List all role assignments that apply to a resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) ListForResource(resourceGroupName string, resourceProviderNamespace string, resourceType string, resourceName string, options *RoleAssignmentsListForResourceOptions) RoleAssignmentListResultPager {
	return &roleAssignmentListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listForResourceCreateRequest(ctx, resourceGroupName, resourceProviderNamespace, resourceType, resourceName, options)
		},
		responder: client.listForResourceHandleResponse,
		errorer:   client.listForResourceHandleError,
		advancer: func(ctx context.Context, resp RoleAssignmentListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RoleAssignmentListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listForResourceCreateRequest creates the ListForResource request.
func (client *RoleAssignmentsClient) listForResourceCreateRequest(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, resourceType string, resourceName string, options *RoleAssignmentsListForResourceOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/roleAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceProviderNamespace == "" {
		return nil, errors.New("parameter resourceProviderNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceProviderNamespace}", resourceProviderNamespace)
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", resourceType)
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", resourceName)
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	if options != nil && options.TenantID != nil {
		reqQP.Set("tenantId", *options.TenantID)
	}
	req.URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listForResourceHandleResponse handles the ListForResource response.
func (client *RoleAssignmentsClient) listForResourceHandleResponse(resp *azcore.Response) (RoleAssignmentListResultResponse, error) {
	var val *RoleAssignmentListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RoleAssignmentListResultResponse{}, err
	}
	return RoleAssignmentListResultResponse{RawResponse: resp.Response, RoleAssignmentListResult: val}, nil
}

// listForResourceHandleError handles the ListForResource error response.
func (client *RoleAssignmentsClient) listForResourceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListForResourceGroup - List all role assignments that apply to a resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) ListForResourceGroup(resourceGroupName string, options *RoleAssignmentsListForResourceGroupOptions) RoleAssignmentListResultPager {
	return &roleAssignmentListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listForResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listForResourceGroupHandleResponse,
		errorer:   client.listForResourceGroupHandleError,
		advancer: func(ctx context.Context, resp RoleAssignmentListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RoleAssignmentListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listForResourceGroupCreateRequest creates the ListForResourceGroup request.
func (client *RoleAssignmentsClient) listForResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *RoleAssignmentsListForResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/roleAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	if options != nil && options.TenantID != nil {
		reqQP.Set("tenantId", *options.TenantID)
	}
	req.URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listForResourceGroupHandleResponse handles the ListForResourceGroup response.
func (client *RoleAssignmentsClient) listForResourceGroupHandleResponse(resp *azcore.Response) (RoleAssignmentListResultResponse, error) {
	var val *RoleAssignmentListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RoleAssignmentListResultResponse{}, err
	}
	return RoleAssignmentListResultResponse{RawResponse: resp.Response, RoleAssignmentListResult: val}, nil
}

// listForResourceGroupHandleError handles the ListForResourceGroup error response.
func (client *RoleAssignmentsClient) listForResourceGroupHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListForScope - List all role assignments that apply to a scope.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) ListForScope(scope string, options *RoleAssignmentsListForScopeOptions) RoleAssignmentListResultPager {
	return &roleAssignmentListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listForScopeCreateRequest(ctx, scope, options)
		},
		responder: client.listForScopeHandleResponse,
		errorer:   client.listForScopeHandleError,
		advancer: func(ctx context.Context, resp RoleAssignmentListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RoleAssignmentListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleAssignmentsClient) listForScopeCreateRequest(ctx context.Context, scope string, options *RoleAssignmentsListForScopeOptions) (*azcore.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	if options != nil && options.TenantID != nil {
		reqQP.Set("tenantId", *options.TenantID)
	}
	req.URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleAssignmentsClient) listForScopeHandleResponse(resp *azcore.Response) (RoleAssignmentListResultResponse, error) {
	var val *RoleAssignmentListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RoleAssignmentListResultResponse{}, err
	}
	return RoleAssignmentListResultResponse{RawResponse: resp.Response, RoleAssignmentListResult: val}, nil
}

// listForScopeHandleError handles the ListForScope error response.
func (client *RoleAssignmentsClient) listForScopeHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListForSubscription - List all role assignments that apply to a subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) ListForSubscription(options *RoleAssignmentsListForSubscriptionOptions) RoleAssignmentListResultPager {
	return &roleAssignmentListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listForSubscriptionCreateRequest(ctx, options)
		},
		responder: client.listForSubscriptionHandleResponse,
		errorer:   client.listForSubscriptionHandleError,
		advancer: func(ctx context.Context, resp RoleAssignmentListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.RoleAssignmentListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listForSubscriptionCreateRequest creates the ListForSubscription request.
func (client *RoleAssignmentsClient) listForSubscriptionCreateRequest(ctx context.Context, options *RoleAssignmentsListForSubscriptionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/roleAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	if options != nil && options.TenantID != nil {
		reqQP.Set("tenantId", *options.TenantID)
	}
	req.URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listForSubscriptionHandleResponse handles the ListForSubscription response.
func (client *RoleAssignmentsClient) listForSubscriptionHandleResponse(resp *azcore.Response) (RoleAssignmentListResultResponse, error) {
	var val *RoleAssignmentListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return RoleAssignmentListResultResponse{}, err
	}
	return RoleAssignmentListResultResponse{RawResponse: resp.Response, RoleAssignmentListResult: val}, nil
}

// listForSubscriptionHandleError handles the ListForSubscription error response.
func (client *RoleAssignmentsClient) listForSubscriptionHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Validate - Validate a role assignment create or update operation by scope and name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) Validate(ctx context.Context, scope string, roleAssignmentName string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsValidateOptions) (ValidationResponseResponse, error) {
	req, err := client.validateCreateRequest(ctx, scope, roleAssignmentName, parameters, options)
	if err != nil {
		return ValidationResponseResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ValidationResponseResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ValidationResponseResponse{}, client.validateHandleError(resp)
	}
	return client.validateHandleResponse(resp)
}

// validateCreateRequest creates the Validate request.
func (client *RoleAssignmentsClient) validateCreateRequest(ctx context.Context, scope string, roleAssignmentName string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsValidateOptions) (*azcore.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}/validate"
	if scope == "" {
		return nil, errors.New("parameter scope cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleAssignmentName == "" {
		return nil, errors.New("parameter roleAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentName}", roleAssignmentName)
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// validateHandleResponse handles the Validate response.
func (client *RoleAssignmentsClient) validateHandleResponse(resp *azcore.Response) (ValidationResponseResponse, error) {
	var val *ValidationResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ValidationResponseResponse{}, err
	}
	return ValidationResponseResponse{RawResponse: resp.Response, ValidationResponse: val}, nil
}

// validateHandleError handles the Validate error response.
func (client *RoleAssignmentsClient) validateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ValidateByID - Validate a role assignment create or update operation by ID.
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentsClient) ValidateByID(ctx context.Context, roleAssignmentID string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsValidateByIDOptions) (ValidationResponseResponse, error) {
	req, err := client.validateByIDCreateRequest(ctx, roleAssignmentID, parameters, options)
	if err != nil {
		return ValidationResponseResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ValidationResponseResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ValidationResponseResponse{}, client.validateByIDHandleError(resp)
	}
	return client.validateByIDHandleResponse(resp)
}

// validateByIDCreateRequest creates the ValidateByID request.
func (client *RoleAssignmentsClient) validateByIDCreateRequest(ctx context.Context, roleAssignmentID string, parameters RoleAssignmentCreateParameters, options *RoleAssignmentsValidateByIDOptions) (*azcore.Request, error) {
	urlPath := "/{roleAssignmentId}/validate"
	if roleAssignmentID == "" {
		return nil, errors.New("parameter roleAssignmentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleAssignmentId}", roleAssignmentID)
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// validateByIDHandleResponse handles the ValidateByID response.
func (client *RoleAssignmentsClient) validateByIDHandleResponse(resp *azcore.Response) (ValidationResponseResponse, error) {
	var val *ValidationResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ValidationResponseResponse{}, err
	}
	return ValidationResponseResponse{RawResponse: resp.Response, ValidationResponse: val}, nil
}

// validateByIDHandleError handles the ValidateByID error response.
func (client *RoleAssignmentsClient) validateByIDHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
