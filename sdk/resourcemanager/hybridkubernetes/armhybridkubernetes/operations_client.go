//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridkubernetes

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// OperationsClient contains the methods for the Operations group.
// Don't use this type directly, use NewOperationsClient() instead.
type OperationsClient struct {
	internal *arm.Client
}

// NewOperationsClient creates a new instance of OperationsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOperationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationsClient, error) {
	cl, err := arm.NewClient(moduleName+".OperationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OperationsClient{
		internal: cl,
	}
	return client, nil
}

// NewGetPager - Lists all of the available API operations for Connected Cluster resource.
//
// Generated from API version 2021-10-01
//   - options - OperationsClientGetOptions contains the optional parameters for the OperationsClient.NewGetPager method.
func (client *OperationsClient) NewGetPager(options *OperationsClientGetOptions) *runtime.Pager[OperationsClientGetResponse] {
	return runtime.NewPager(runtime.PagingHandler[OperationsClientGetResponse]{
		More: func(page OperationsClientGetResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OperationsClientGetResponse) (OperationsClientGetResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OperationsClientGetResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return OperationsClientGetResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OperationsClientGetResponse{}, runtime.NewResponseError(resp)
			}
			return client.getHandleResponse(resp)
		},
	})
}

// getCreateRequest creates the Get request.
func (client *OperationsClient) getCreateRequest(ctx context.Context, options *OperationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Kubernetes/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OperationsClient) getHandleResponse(resp *http.Response) (OperationsClientGetResponse, error) {
	result := OperationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationList); err != nil {
		return OperationsClientGetResponse{}, err
	}
	return result, nil
}
