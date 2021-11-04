//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

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

// EdgeNodesClient contains the methods for the EdgeNodes group.
// Don't use this type directly, use NewEdgeNodesClient() instead.
type EdgeNodesClient struct {
	ep string
	pl runtime.Pipeline
}

// NewEdgeNodesClient creates a new instance of EdgeNodesClient with the specified values.
func NewEdgeNodesClient(credential azcore.TokenCredential, options *arm.ClientOptions) *EdgeNodesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &EdgeNodesClient{ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// List - Edgenodes are the global Point of Presence (POP) locations used to deliver CDN content to end users.
// If the operation fails it returns the *ErrorResponse error type.
func (client *EdgeNodesClient) List(options *EdgeNodesListOptions) *EdgeNodesListPager {
	return &EdgeNodesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp EdgeNodesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.EdgenodeResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *EdgeNodesClient) listCreateRequest(ctx context.Context, options *EdgeNodesListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Cdn/edgenodes"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EdgeNodesClient) listHandleResponse(resp *http.Response) (EdgeNodesListResponse, error) {
	result := EdgeNodesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.EdgenodeResult); err != nil {
		return EdgeNodesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *EdgeNodesClient) listHandleError(resp *http.Response) error {
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
