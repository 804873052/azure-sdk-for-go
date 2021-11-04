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
	"strconv"
	"strings"
)

// VirtualMachineImagesEdgeZoneClient contains the methods for the VirtualMachineImagesEdgeZone group.
// Don't use this type directly, use NewVirtualMachineImagesEdgeZoneClient() instead.
type VirtualMachineImagesEdgeZoneClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewVirtualMachineImagesEdgeZoneClient creates a new instance of VirtualMachineImagesEdgeZoneClient with the specified values.
func NewVirtualMachineImagesEdgeZoneClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VirtualMachineImagesEdgeZoneClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &VirtualMachineImagesEdgeZoneClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Get - Gets a virtual machine image in an edge zone.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineImagesEdgeZoneClient) Get(ctx context.Context, location string, edgeZone string, publisherName string, offer string, skus string, version string, options *VirtualMachineImagesEdgeZoneGetOptions) (VirtualMachineImagesEdgeZoneGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, edgeZone, publisherName, offer, skus, version, options)
	if err != nil {
		return VirtualMachineImagesEdgeZoneGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineImagesEdgeZoneGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineImagesEdgeZoneGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VirtualMachineImagesEdgeZoneClient) getCreateRequest(ctx context.Context, location string, edgeZone string, publisherName string, offer string, skus string, version string, options *VirtualMachineImagesEdgeZoneGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/edgeZones/{edgeZone}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions/{version}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if edgeZone == "" {
		return nil, errors.New("parameter edgeZone cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{edgeZone}", url.PathEscape(edgeZone))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if offer == "" {
		return nil, errors.New("parameter offer cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
	if skus == "" {
		return nil, errors.New("parameter skus cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skus}", url.PathEscape(skus))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *VirtualMachineImagesEdgeZoneClient) getHandleResponse(resp *http.Response) (VirtualMachineImagesEdgeZoneGetResponse, error) {
	result := VirtualMachineImagesEdgeZoneGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineImage); err != nil {
		return VirtualMachineImagesEdgeZoneGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VirtualMachineImagesEdgeZoneClient) getHandleError(resp *http.Response) error {
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

// List - Gets a list of all virtual machine image versions for the specified location, edge zone, publisher, offer, and SKU.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineImagesEdgeZoneClient) List(ctx context.Context, location string, edgeZone string, publisherName string, offer string, skus string, options *VirtualMachineImagesEdgeZoneListOptions) (VirtualMachineImagesEdgeZoneListResponse, error) {
	req, err := client.listCreateRequest(ctx, location, edgeZone, publisherName, offer, skus, options)
	if err != nil {
		return VirtualMachineImagesEdgeZoneListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineImagesEdgeZoneListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineImagesEdgeZoneListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *VirtualMachineImagesEdgeZoneClient) listCreateRequest(ctx context.Context, location string, edgeZone string, publisherName string, offer string, skus string, options *VirtualMachineImagesEdgeZoneListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/edgeZones/{edgeZone}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus/{skus}/versions"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if edgeZone == "" {
		return nil, errors.New("parameter edgeZone cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{edgeZone}", url.PathEscape(edgeZone))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if offer == "" {
		return nil, errors.New("parameter offer cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
	if skus == "" {
		return nil, errors.New("parameter skus cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skus}", url.PathEscape(skus))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2021-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineImagesEdgeZoneClient) listHandleResponse(resp *http.Response) (VirtualMachineImagesEdgeZoneListResponse, error) {
	result := VirtualMachineImagesEdgeZoneListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineImageResourceArray); err != nil {
		return VirtualMachineImagesEdgeZoneListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *VirtualMachineImagesEdgeZoneClient) listHandleError(resp *http.Response) error {
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

// ListOffers - Gets a list of virtual machine image offers for the specified location, edge zone and publisher.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineImagesEdgeZoneClient) ListOffers(ctx context.Context, location string, edgeZone string, publisherName string, options *VirtualMachineImagesEdgeZoneListOffersOptions) (VirtualMachineImagesEdgeZoneListOffersResponse, error) {
	req, err := client.listOffersCreateRequest(ctx, location, edgeZone, publisherName, options)
	if err != nil {
		return VirtualMachineImagesEdgeZoneListOffersResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineImagesEdgeZoneListOffersResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineImagesEdgeZoneListOffersResponse{}, client.listOffersHandleError(resp)
	}
	return client.listOffersHandleResponse(resp)
}

// listOffersCreateRequest creates the ListOffers request.
func (client *VirtualMachineImagesEdgeZoneClient) listOffersCreateRequest(ctx context.Context, location string, edgeZone string, publisherName string, options *VirtualMachineImagesEdgeZoneListOffersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/edgeZones/{edgeZone}/publishers/{publisherName}/artifacttypes/vmimage/offers"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if edgeZone == "" {
		return nil, errors.New("parameter edgeZone cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{edgeZone}", url.PathEscape(edgeZone))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listOffersHandleResponse handles the ListOffers response.
func (client *VirtualMachineImagesEdgeZoneClient) listOffersHandleResponse(resp *http.Response) (VirtualMachineImagesEdgeZoneListOffersResponse, error) {
	result := VirtualMachineImagesEdgeZoneListOffersResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineImageResourceArray); err != nil {
		return VirtualMachineImagesEdgeZoneListOffersResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listOffersHandleError handles the ListOffers error response.
func (client *VirtualMachineImagesEdgeZoneClient) listOffersHandleError(resp *http.Response) error {
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

// ListPublishers - Gets a list of virtual machine image publishers for the specified Azure location and edge zone.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineImagesEdgeZoneClient) ListPublishers(ctx context.Context, location string, edgeZone string, options *VirtualMachineImagesEdgeZoneListPublishersOptions) (VirtualMachineImagesEdgeZoneListPublishersResponse, error) {
	req, err := client.listPublishersCreateRequest(ctx, location, edgeZone, options)
	if err != nil {
		return VirtualMachineImagesEdgeZoneListPublishersResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineImagesEdgeZoneListPublishersResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineImagesEdgeZoneListPublishersResponse{}, client.listPublishersHandleError(resp)
	}
	return client.listPublishersHandleResponse(resp)
}

// listPublishersCreateRequest creates the ListPublishers request.
func (client *VirtualMachineImagesEdgeZoneClient) listPublishersCreateRequest(ctx context.Context, location string, edgeZone string, options *VirtualMachineImagesEdgeZoneListPublishersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/edgeZones/{edgeZone}/publishers"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if edgeZone == "" {
		return nil, errors.New("parameter edgeZone cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{edgeZone}", url.PathEscape(edgeZone))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listPublishersHandleResponse handles the ListPublishers response.
func (client *VirtualMachineImagesEdgeZoneClient) listPublishersHandleResponse(resp *http.Response) (VirtualMachineImagesEdgeZoneListPublishersResponse, error) {
	result := VirtualMachineImagesEdgeZoneListPublishersResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineImageResourceArray); err != nil {
		return VirtualMachineImagesEdgeZoneListPublishersResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listPublishersHandleError handles the ListPublishers error response.
func (client *VirtualMachineImagesEdgeZoneClient) listPublishersHandleError(resp *http.Response) error {
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

// ListSKUs - Gets a list of virtual machine image SKUs for the specified location, edge zone, publisher, and offer.
// If the operation fails it returns the *CloudError error type.
func (client *VirtualMachineImagesEdgeZoneClient) ListSKUs(ctx context.Context, location string, edgeZone string, publisherName string, offer string, options *VirtualMachineImagesEdgeZoneListSKUsOptions) (VirtualMachineImagesEdgeZoneListSKUsResponse, error) {
	req, err := client.listSKUsCreateRequest(ctx, location, edgeZone, publisherName, offer, options)
	if err != nil {
		return VirtualMachineImagesEdgeZoneListSKUsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VirtualMachineImagesEdgeZoneListSKUsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VirtualMachineImagesEdgeZoneListSKUsResponse{}, client.listSKUsHandleError(resp)
	}
	return client.listSKUsHandleResponse(resp)
}

// listSKUsCreateRequest creates the ListSKUs request.
func (client *VirtualMachineImagesEdgeZoneClient) listSKUsCreateRequest(ctx context.Context, location string, edgeZone string, publisherName string, offer string, options *VirtualMachineImagesEdgeZoneListSKUsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/edgeZones/{edgeZone}/publishers/{publisherName}/artifacttypes/vmimage/offers/{offer}/skus"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if edgeZone == "" {
		return nil, errors.New("parameter edgeZone cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{edgeZone}", url.PathEscape(edgeZone))
	if publisherName == "" {
		return nil, errors.New("parameter publisherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherName}", url.PathEscape(publisherName))
	if offer == "" {
		return nil, errors.New("parameter offer cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offer}", url.PathEscape(offer))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listSKUsHandleResponse handles the ListSKUs response.
func (client *VirtualMachineImagesEdgeZoneClient) listSKUsHandleResponse(resp *http.Response) (VirtualMachineImagesEdgeZoneListSKUsResponse, error) {
	result := VirtualMachineImagesEdgeZoneListSKUsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualMachineImageResourceArray); err != nil {
		return VirtualMachineImagesEdgeZoneListSKUsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listSKUsHandleError handles the ListSKUs error response.
func (client *VirtualMachineImagesEdgeZoneClient) listSKUsHandleError(resp *http.Response) error {
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
