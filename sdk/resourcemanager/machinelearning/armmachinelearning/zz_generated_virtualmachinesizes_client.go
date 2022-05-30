//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VirtualMachineSizesClient contains the methods for the VirtualMachineSizes group.
// Don't use this type directly, use NewVirtualMachineSizesClient() instead.
type VirtualMachineSizesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualMachineSizesClient creates a new instance of VirtualMachineSizesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualMachineSizesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualMachineSizesClient, error) {
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
	client := &VirtualMachineSizesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// List - Returns supported VM Sizes in a location
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-02-01-preview
// location - The location upon which virtual-machine-sizes is queried.
// options - VirtualMachineSizesClientListOptions contains the optional parameters for the VirtualMachineSizesClient.List
// method.
func (client *VirtualMachineSizesClient) List(ctx context.Context, location string, options *VirtualMachineSizesClientListOptions) (VirtualMachineSizesClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, location, options)
	if err != nil {
		return VirtualMachineSizesClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineSizesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineSizesClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *VirtualMachineSizesClient) listCreateRequest(ctx context.Context, location string, options *VirtualMachineSizesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearningServices/locations/{location}/vmSizes"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineSizesClient) listHandleResponse(resp *http.Response) (VirtualMachineSizesClientListResponse, error) {
	result := VirtualMachineSizesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineSizeListResult); err != nil {
		return VirtualMachineSizesClientListResponse{}, err
	}
	return result, nil
}
