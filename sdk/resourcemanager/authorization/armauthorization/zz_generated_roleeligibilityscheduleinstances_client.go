//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

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

// RoleEligibilityScheduleInstancesClient contains the methods for the RoleEligibilityScheduleInstances group.
// Don't use this type directly, use NewRoleEligibilityScheduleInstancesClient() instead.
type RoleEligibilityScheduleInstancesClient struct {
	ep string
	pl runtime.Pipeline
}

// NewRoleEligibilityScheduleInstancesClient creates a new instance of RoleEligibilityScheduleInstancesClient with the specified values.
func NewRoleEligibilityScheduleInstancesClient(credential azcore.TokenCredential, options *arm.ClientOptions) *RoleEligibilityScheduleInstancesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &RoleEligibilityScheduleInstancesClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Gets the specified role eligibility schedule instance.
// If the operation fails it returns the *CloudError error type.
func (client *RoleEligibilityScheduleInstancesClient) Get(ctx context.Context, scope string, roleEligibilityScheduleInstanceName string, options *RoleEligibilityScheduleInstancesGetOptions) (RoleEligibilityScheduleInstancesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, roleEligibilityScheduleInstanceName, options)
	if err != nil {
		return RoleEligibilityScheduleInstancesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleEligibilityScheduleInstancesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleEligibilityScheduleInstancesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RoleEligibilityScheduleInstancesClient) getCreateRequest(ctx context.Context, scope string, roleEligibilityScheduleInstanceName string, options *RoleEligibilityScheduleInstancesGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleEligibilityScheduleInstances/{roleEligibilityScheduleInstanceName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if roleEligibilityScheduleInstanceName == "" {
		return nil, errors.New("parameter roleEligibilityScheduleInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleEligibilityScheduleInstanceName}", url.PathEscape(roleEligibilityScheduleInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RoleEligibilityScheduleInstancesClient) getHandleResponse(resp *http.Response) (RoleEligibilityScheduleInstancesGetResponse, error) {
	result := RoleEligibilityScheduleInstancesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleEligibilityScheduleInstance); err != nil {
		return RoleEligibilityScheduleInstancesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *RoleEligibilityScheduleInstancesClient) getHandleError(resp *http.Response) error {
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

// ListForScope - Gets role eligibility schedule instances of a role eligibility schedule.
// If the operation fails it returns the *CloudError error type.
func (client *RoleEligibilityScheduleInstancesClient) ListForScope(scope string, options *RoleEligibilityScheduleInstancesListForScopeOptions) *RoleEligibilityScheduleInstancesListForScopePager {
	return &RoleEligibilityScheduleInstancesListForScopePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listForScopeCreateRequest(ctx, scope, options)
		},
		advancer: func(ctx context.Context, resp RoleEligibilityScheduleInstancesListForScopeResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.RoleEligibilityScheduleInstanceListResult.NextLink)
		},
	}
}

// listForScopeCreateRequest creates the ListForScope request.
func (client *RoleEligibilityScheduleInstancesClient) listForScopeCreateRequest(ctx context.Context, scope string, options *RoleEligibilityScheduleInstancesListForScopeOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Authorization/roleEligibilityScheduleInstances"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listForScopeHandleResponse handles the ListForScope response.
func (client *RoleEligibilityScheduleInstancesClient) listForScopeHandleResponse(resp *http.Response) (RoleEligibilityScheduleInstancesListForScopeResponse, error) {
	result := RoleEligibilityScheduleInstancesListForScopeResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleEligibilityScheduleInstanceListResult); err != nil {
		return RoleEligibilityScheduleInstancesListForScopeResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listForScopeHandleError handles the ListForScope error response.
func (client *RoleEligibilityScheduleInstancesClient) listForScopeHandleError(resp *http.Response) error {
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
