//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ServerCommunicationLinksClient contains the methods for the ServerCommunicationLinks group.
// Don't use this type directly, use NewServerCommunicationLinksClient() instead.
type ServerCommunicationLinksClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewServerCommunicationLinksClient creates a new instance of ServerCommunicationLinksClient with the specified values.
func NewServerCommunicationLinksClient(con *arm.Connection, subscriptionID string) *ServerCommunicationLinksClient {
	return &ServerCommunicationLinksClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates a server communication link.
// If the operation fails it returns a generic error.
func (client *ServerCommunicationLinksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, communicationLinkName string, parameters ServerCommunicationLink, options *ServerCommunicationLinksBeginCreateOrUpdateOptions) (ServerCommunicationLinksCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, communicationLinkName, parameters, options)
	if err != nil {
		return ServerCommunicationLinksCreateOrUpdatePollerResponse{}, err
	}
	result := ServerCommunicationLinksCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServerCommunicationLinksClient.CreateOrUpdate", "", resp, client.pl, client.createOrUpdateHandleError)
	if err != nil {
		return ServerCommunicationLinksCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ServerCommunicationLinksCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a server communication link.
// If the operation fails it returns a generic error.
func (client *ServerCommunicationLinksClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, communicationLinkName string, parameters ServerCommunicationLink, options *ServerCommunicationLinksBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, communicationLinkName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerCommunicationLinksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, communicationLinkName string, parameters ServerCommunicationLink, options *ServerCommunicationLinksBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/communicationLinks/{communicationLinkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if communicationLinkName == "" {
		return nil, errors.New("parameter communicationLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communicationLinkName}", url.PathEscape(communicationLinkName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ServerCommunicationLinksClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Delete - Deletes a server communication link.
// If the operation fails it returns a generic error.
func (client *ServerCommunicationLinksClient) Delete(ctx context.Context, resourceGroupName string, serverName string, communicationLinkName string, options *ServerCommunicationLinksDeleteOptions) (ServerCommunicationLinksDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, communicationLinkName, options)
	if err != nil {
		return ServerCommunicationLinksDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerCommunicationLinksDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerCommunicationLinksDeleteResponse{}, client.deleteHandleError(resp)
	}
	return ServerCommunicationLinksDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServerCommunicationLinksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, communicationLinkName string, options *ServerCommunicationLinksDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/communicationLinks/{communicationLinkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if communicationLinkName == "" {
		return nil, errors.New("parameter communicationLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communicationLinkName}", url.PathEscape(communicationLinkName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ServerCommunicationLinksClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// Get - Returns a server communication link.
// If the operation fails it returns a generic error.
func (client *ServerCommunicationLinksClient) Get(ctx context.Context, resourceGroupName string, serverName string, communicationLinkName string, options *ServerCommunicationLinksGetOptions) (ServerCommunicationLinksGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, communicationLinkName, options)
	if err != nil {
		return ServerCommunicationLinksGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerCommunicationLinksGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerCommunicationLinksGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerCommunicationLinksClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, communicationLinkName string, options *ServerCommunicationLinksGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/communicationLinks/{communicationLinkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if communicationLinkName == "" {
		return nil, errors.New("parameter communicationLinkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{communicationLinkName}", url.PathEscape(communicationLinkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerCommunicationLinksClient) getHandleResponse(resp *http.Response) (ServerCommunicationLinksGetResponse, error) {
	result := ServerCommunicationLinksGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerCommunicationLink); err != nil {
		return ServerCommunicationLinksGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ServerCommunicationLinksClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}

// ListByServer - Gets a list of server communication links.
// If the operation fails it returns a generic error.
func (client *ServerCommunicationLinksClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string, options *ServerCommunicationLinksListByServerOptions) (ServerCommunicationLinksListByServerResponse, error) {
	req, err := client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServerCommunicationLinksListByServerResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerCommunicationLinksListByServerResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerCommunicationLinksListByServerResponse{}, client.listByServerHandleError(resp)
	}
	return client.listByServerHandleResponse(resp)
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerCommunicationLinksClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerCommunicationLinksListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/communicationLinks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2014-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerCommunicationLinksClient) listByServerHandleResponse(resp *http.Response) (ServerCommunicationLinksListByServerResponse, error) {
	result := ServerCommunicationLinksListByServerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerCommunicationLinkListResult); err != nil {
		return ServerCommunicationLinksListByServerResponse{}, err
	}
	return result, nil
}

// listByServerHandleError handles the ListByServer error response.
func (client *ServerCommunicationLinksClient) listByServerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
