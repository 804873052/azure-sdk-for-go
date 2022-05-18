//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagedservices

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// OperationsWithScopeClient contains the methods for the OperationsWithScope group.
// Don't use this type directly, use NewOperationsWithScopeClient() instead.
type OperationsWithScopeClient struct {
	host string
	pl   runtime.Pipeline
}

// NewOperationsWithScopeClient creates a new instance of OperationsWithScopeClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewOperationsWithScopeClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationsWithScopeClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &OperationsWithScopeClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// List - Gets a list of the operations with the scope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// scope - The scope of the resource.
// options - OperationsWithScopeClientListOptions contains the optional parameters for the OperationsWithScopeClient.List
// method.
func (client *OperationsWithScopeClient) List(ctx context.Context, scope string, options *OperationsWithScopeClientListOptions) (OperationsWithScopeClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, scope, options)
	if err != nil {
		return OperationsWithScopeClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OperationsWithScopeClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OperationsWithScopeClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *OperationsWithScopeClient) listCreateRequest(ctx context.Context, scope string, options *OperationsWithScopeClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.ManagedServices/operations"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OperationsWithScopeClient) listHandleResponse(resp *http.Response) (OperationsWithScopeClientListResponse, error) {
	result := OperationsWithScopeClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationList); err != nil {
		return OperationsWithScopeClientListResponse{}, err
	}
	return result, nil
}
