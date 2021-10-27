//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfeatures

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// FeatureClient contains the methods for the FeatureClient group.
// Don't use this type directly, use NewFeatureClient() instead.
type FeatureClient struct {
	ep string
	pl runtime.Pipeline
}

// NewFeatureClient creates a new instance of FeatureClient with the specified values.
func NewFeatureClient(credential azcore.TokenCredential, options *arm.ClientOptions) *FeatureClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &FeatureClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// ListOperations - Lists all of the available Microsoft.Features REST API operations.
// If the operation fails it returns the *ErrorResponse error type.
func (client *FeatureClient) ListOperations(options *FeatureClientListOperationsOptions) *FeatureClientListOperationsPager {
	return &FeatureClientListOperationsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listOperationsCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp FeatureClientListOperationsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.OperationListResult.NextLink)
		},
	}
}

// listOperationsCreateRequest creates the ListOperations request.
func (client *FeatureClient) listOperationsCreateRequest(ctx context.Context, options *FeatureClientListOperationsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Features/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json, text/json")
	return req, nil
}

// listOperationsHandleResponse handles the ListOperations response.
func (client *FeatureClient) listOperationsHandleResponse(resp *http.Response) (FeatureClientListOperationsResponse, error) {
	result := FeatureClientListOperationsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationListResult); err != nil {
		return FeatureClientListOperationsResponse{}, err
	}
	return result, nil
}

// listOperationsHandleError handles the ListOperations error response.
func (client *FeatureClient) listOperationsHandleError(resp *http.Response) error {
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
