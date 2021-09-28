//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armweb

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// DomainRegistrationProviderClient contains the methods for the DomainRegistrationProvider group.
// Don't use this type directly, use NewDomainRegistrationProviderClient() instead.
type DomainRegistrationProviderClient struct {
	ep string
	pl runtime.Pipeline
}

// NewDomainRegistrationProviderClient creates a new instance of DomainRegistrationProviderClient with the specified values.
func NewDomainRegistrationProviderClient(con *arm.Connection) *DomainRegistrationProviderClient {
	return &DomainRegistrationProviderClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// ListOperations - Description for Implements Csm operations Api to exposes the list of available Csm Apis under the resource provider
// If the operation fails it returns the *DefaultErrorResponse error type.
func (client *DomainRegistrationProviderClient) ListOperations(options *DomainRegistrationProviderListOperationsOptions) *DomainRegistrationProviderListOperationsPager {
	return &DomainRegistrationProviderListOperationsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listOperationsCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp DomainRegistrationProviderListOperationsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CsmOperationCollection.NextLink)
		},
	}
}

// listOperationsCreateRequest creates the ListOperations request.
func (client *DomainRegistrationProviderClient) listOperationsCreateRequest(ctx context.Context, options *DomainRegistrationProviderListOperationsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.DomainRegistration/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listOperationsHandleResponse handles the ListOperations response.
func (client *DomainRegistrationProviderClient) listOperationsHandleResponse(resp *http.Response) (DomainRegistrationProviderListOperationsResponse, error) {
	result := DomainRegistrationProviderListOperationsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CsmOperationCollection); err != nil {
		return DomainRegistrationProviderListOperationsResponse{}, err
	}
	return result, nil
}

// listOperationsHandleError handles the ListOperations error response.
func (client *DomainRegistrationProviderClient) listOperationsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := DefaultErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
