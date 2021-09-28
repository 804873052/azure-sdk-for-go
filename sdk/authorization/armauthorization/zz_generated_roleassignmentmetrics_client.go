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
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// RoleAssignmentMetricsClient contains the methods for the RoleAssignmentMetrics group.
// Don't use this type directly, use NewRoleAssignmentMetricsClient() instead.
type RoleAssignmentMetricsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewRoleAssignmentMetricsClient creates a new instance of RoleAssignmentMetricsClient with the specified values.
func NewRoleAssignmentMetricsClient(con *arm.Connection, subscriptionID string) *RoleAssignmentMetricsClient {
	return &RoleAssignmentMetricsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// GetMetricsForSubscription - Get role assignment usage metrics for a subscription
// If the operation fails it returns the *ErrorResponse error type.
func (client *RoleAssignmentMetricsClient) GetMetricsForSubscription(ctx context.Context, options *RoleAssignmentMetricsGetMetricsForSubscriptionOptions) (RoleAssignmentMetricsGetMetricsForSubscriptionResponse, error) {
	req, err := client.getMetricsForSubscriptionCreateRequest(ctx, options)
	if err != nil {
		return RoleAssignmentMetricsGetMetricsForSubscriptionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RoleAssignmentMetricsGetMetricsForSubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RoleAssignmentMetricsGetMetricsForSubscriptionResponse{}, client.getMetricsForSubscriptionHandleError(resp)
	}
	return client.getMetricsForSubscriptionHandleResponse(resp)
}

// getMetricsForSubscriptionCreateRequest creates the GetMetricsForSubscription request.
func (client *RoleAssignmentMetricsClient) getMetricsForSubscriptionCreateRequest(ctx context.Context, options *RoleAssignmentMetricsGetMetricsForSubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/roleAssignmentsUsageMetrics"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getMetricsForSubscriptionHandleResponse handles the GetMetricsForSubscription response.
func (client *RoleAssignmentMetricsClient) getMetricsForSubscriptionHandleResponse(resp *http.Response) (RoleAssignmentMetricsGetMetricsForSubscriptionResponse, error) {
	result := RoleAssignmentMetricsGetMetricsForSubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RoleAssignmentMetricsResult); err != nil {
		return RoleAssignmentMetricsGetMetricsForSubscriptionResponse{}, err
	}
	return result, nil
}

// getMetricsForSubscriptionHandleError handles the GetMetricsForSubscription error response.
func (client *RoleAssignmentMetricsClient) getMetricsForSubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
