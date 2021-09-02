//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ExpressRouteGatewaysClient contains the methods for the ExpressRouteGateways group.
// Don't use this type directly, use NewExpressRouteGatewaysClient() instead.
type ExpressRouteGatewaysClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewExpressRouteGatewaysClient creates a new instance of ExpressRouteGatewaysClient with the specified values.
func NewExpressRouteGatewaysClient(con *arm.Connection, subscriptionID string) *ExpressRouteGatewaysClient {
	return &ExpressRouteGatewaysClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a ExpressRoute gateway in a specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteGatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, putExpressRouteGatewayParameters ExpressRouteGateway, options *ExpressRouteGatewaysBeginCreateOrUpdateOptions) (ExpressRouteGatewaysCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, expressRouteGatewayName, putExpressRouteGatewayParameters, options)
	if err != nil {
		return ExpressRouteGatewaysCreateOrUpdatePollerResponse{}, err
	}
	result := ExpressRouteGatewaysCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExpressRouteGatewaysClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return ExpressRouteGatewaysCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ExpressRouteGatewaysCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a ExpressRoute gateway in a specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteGatewaysClient) createOrUpdate(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, putExpressRouteGatewayParameters ExpressRouteGateway, options *ExpressRouteGatewaysBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, expressRouteGatewayName, putExpressRouteGatewayParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExpressRouteGatewaysClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, putExpressRouteGatewayParameters ExpressRouteGateway, options *ExpressRouteGatewaysBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRouteGatewayName == "" {
		return nil, errors.New("parameter expressRouteGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRouteGatewayName}", url.PathEscape(expressRouteGatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, putExpressRouteGatewayParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ExpressRouteGatewaysClient) createOrUpdateHandleError(resp *http.Response) error {
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

// BeginDelete - Deletes the specified ExpressRoute gateway in a resource group. An ExpressRoute gateway resource can only be deleted when there are no
// connection subresources.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteGatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, options *ExpressRouteGatewaysBeginDeleteOptions) (ExpressRouteGatewaysDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, expressRouteGatewayName, options)
	if err != nil {
		return ExpressRouteGatewaysDeletePollerResponse{}, err
	}
	result := ExpressRouteGatewaysDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExpressRouteGatewaysClient.Delete", "location", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return ExpressRouteGatewaysDeletePollerResponse{}, err
	}
	result.Poller = &ExpressRouteGatewaysDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified ExpressRoute gateway in a resource group. An ExpressRoute gateway resource can only be deleted when there are no connection
// subresources.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteGatewaysClient) deleteOperation(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, options *ExpressRouteGatewaysBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, expressRouteGatewayName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ExpressRouteGatewaysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, options *ExpressRouteGatewaysBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRouteGatewayName == "" {
		return nil, errors.New("parameter expressRouteGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRouteGatewayName}", url.PathEscape(expressRouteGatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ExpressRouteGatewaysClient) deleteHandleError(resp *http.Response) error {
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

// Get - Fetches the details of a ExpressRoute gateway in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteGatewaysClient) Get(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, options *ExpressRouteGatewaysGetOptions) (ExpressRouteGatewaysGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, expressRouteGatewayName, options)
	if err != nil {
		return ExpressRouteGatewaysGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRouteGatewaysGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRouteGatewaysGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRouteGatewaysClient) getCreateRequest(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, options *ExpressRouteGatewaysGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRouteGatewayName == "" {
		return nil, errors.New("parameter expressRouteGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRouteGatewayName}", url.PathEscape(expressRouteGatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExpressRouteGatewaysClient) getHandleResponse(resp *http.Response) (ExpressRouteGatewaysGetResponse, error) {
	result := ExpressRouteGatewaysGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteGateway); err != nil {
		return ExpressRouteGatewaysGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ExpressRouteGatewaysClient) getHandleError(resp *http.Response) error {
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

// ListByResourceGroup - Lists ExpressRoute gateways in a given resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteGatewaysClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, options *ExpressRouteGatewaysListByResourceGroupOptions) (ExpressRouteGatewaysListByResourceGroupResponse, error) {
	req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return ExpressRouteGatewaysListByResourceGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRouteGatewaysListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRouteGatewaysListByResourceGroupResponse{}, client.listByResourceGroupHandleError(resp)
	}
	return client.listByResourceGroupHandleResponse(resp)
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ExpressRouteGatewaysClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ExpressRouteGatewaysListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ExpressRouteGatewaysClient) listByResourceGroupHandleResponse(resp *http.Response) (ExpressRouteGatewaysListByResourceGroupResponse, error) {
	result := ExpressRouteGatewaysListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteGatewayList); err != nil {
		return ExpressRouteGatewaysListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ExpressRouteGatewaysClient) listByResourceGroupHandleError(resp *http.Response) error {
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

// ListBySubscription - Lists ExpressRoute gateways under a given subscription.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteGatewaysClient) ListBySubscription(ctx context.Context, options *ExpressRouteGatewaysListBySubscriptionOptions) (ExpressRouteGatewaysListBySubscriptionResponse, error) {
	req, err := client.listBySubscriptionCreateRequest(ctx, options)
	if err != nil {
		return ExpressRouteGatewaysListBySubscriptionResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRouteGatewaysListBySubscriptionResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRouteGatewaysListBySubscriptionResponse{}, client.listBySubscriptionHandleError(resp)
	}
	return client.listBySubscriptionHandleResponse(resp)
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ExpressRouteGatewaysClient) listBySubscriptionCreateRequest(ctx context.Context, options *ExpressRouteGatewaysListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteGateways"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ExpressRouteGatewaysClient) listBySubscriptionHandleResponse(resp *http.Response) (ExpressRouteGatewaysListBySubscriptionResponse, error) {
	result := ExpressRouteGatewaysListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteGatewayList); err != nil {
		return ExpressRouteGatewaysListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *ExpressRouteGatewaysClient) listBySubscriptionHandleError(resp *http.Response) error {
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

// BeginUpdateTags - Updates express route gateway tags.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteGatewaysClient) BeginUpdateTags(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, expressRouteGatewayParameters TagsObject, options *ExpressRouteGatewaysBeginUpdateTagsOptions) (ExpressRouteGatewaysUpdateTagsPollerResponse, error) {
	resp, err := client.updateTags(ctx, resourceGroupName, expressRouteGatewayName, expressRouteGatewayParameters, options)
	if err != nil {
		return ExpressRouteGatewaysUpdateTagsPollerResponse{}, err
	}
	result := ExpressRouteGatewaysUpdateTagsPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ExpressRouteGatewaysClient.UpdateTags", "azure-async-operation", resp, client.pl, client.updateTagsHandleError)
	if err != nil {
		return ExpressRouteGatewaysUpdateTagsPollerResponse{}, err
	}
	result.Poller = &ExpressRouteGatewaysUpdateTagsPoller{
		pt: pt,
	}
	return result, nil
}

// UpdateTags - Updates express route gateway tags.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRouteGatewaysClient) updateTags(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, expressRouteGatewayParameters TagsObject, options *ExpressRouteGatewaysBeginUpdateTagsOptions) (*http.Response, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, expressRouteGatewayName, expressRouteGatewayParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateTagsHandleError(resp)
	}
	return resp, nil
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ExpressRouteGatewaysClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, expressRouteGatewayName string, expressRouteGatewayParameters TagsObject, options *ExpressRouteGatewaysBeginUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/expressRouteGateways/{expressRouteGatewayName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRouteGatewayName == "" {
		return nil, errors.New("parameter expressRouteGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRouteGatewayName}", url.PathEscape(expressRouteGatewayName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, expressRouteGatewayParameters)
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *ExpressRouteGatewaysClient) updateTagsHandleError(resp *http.Response) error {
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
