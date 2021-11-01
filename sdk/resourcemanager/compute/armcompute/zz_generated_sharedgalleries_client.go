//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// SharedGalleriesClient contains the methods for the SharedGalleries group.
// Don't use this type directly, use NewSharedGalleriesClient() instead.
type SharedGalleriesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewSharedGalleriesClient creates a new instance of SharedGalleriesClient with the specified values.
func NewSharedGalleriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SharedGalleriesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &SharedGalleriesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Get a shared gallery by subscription id or tenant id.
// If the operation fails it returns the *CloudError error type.
func (client *SharedGalleriesClient) Get(ctx context.Context, location string, galleryUniqueName string, options *SharedGalleriesGetOptions) (SharedGalleriesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, galleryUniqueName, options)
	if err != nil {
		return SharedGalleriesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SharedGalleriesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SharedGalleriesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SharedGalleriesClient) getCreateRequest(ctx context.Context, location string, galleryUniqueName string, options *SharedGalleriesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/sharedGalleries/{galleryUniqueName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if galleryUniqueName == "" {
		return nil, errors.New("parameter galleryUniqueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryUniqueName}", url.PathEscape(galleryUniqueName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SharedGalleriesClient) getHandleResponse(resp *http.Response) (SharedGalleriesGetResponse, error) {
	result := SharedGalleriesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SharedGallery); err != nil {
		return SharedGalleriesGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SharedGalleriesClient) getHandleError(resp *http.Response) error {
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

// List - List shared galleries by subscription id or tenant id.
// If the operation fails it returns the *CloudError error type.
func (client *SharedGalleriesClient) List(location string, options *SharedGalleriesListOptions) *SharedGalleriesListPager {
	return &SharedGalleriesListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp SharedGalleriesListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SharedGalleryList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SharedGalleriesClient) listCreateRequest(ctx context.Context, location string, options *SharedGalleriesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/sharedGalleries"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01")
	if options != nil && options.SharedTo != nil {
		reqQP.Set("sharedTo", string(*options.SharedTo))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SharedGalleriesClient) listHandleResponse(resp *http.Response) (SharedGalleriesListResponse, error) {
	result := SharedGalleriesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SharedGalleryList); err != nil {
		return SharedGalleriesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SharedGalleriesClient) listHandleError(resp *http.Response) error {
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
