//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridconnectivity

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// ServiceConfigurationsClient contains the methods for the ServiceConfigurations group.
// Don't use this type directly, use NewServiceConfigurationsClient() instead.
type ServiceConfigurationsClient struct {
	internal *arm.Client
}

// NewServiceConfigurationsClient creates a new instance of ServiceConfigurationsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServiceConfigurationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ServiceConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServiceConfigurationsClient{
		internal: cl,
	}
	return client, nil
}

// CreateOrupdate - Create or update a service in serviceConfiguration for the endpoint resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
//   - endpointName - The endpoint name.
//   - serviceConfigurationName - The service name.
//   - serviceConfigurationResource - Service details
//   - options - ServiceConfigurationsClientCreateOrupdateOptions contains the optional parameters for the ServiceConfigurationsClient.CreateOrupdate
//     method.
func (client *ServiceConfigurationsClient) CreateOrupdate(ctx context.Context, resourceURI string, endpointName string, serviceConfigurationName string, serviceConfigurationResource ServiceConfigurationResource, options *ServiceConfigurationsClientCreateOrupdateOptions) (ServiceConfigurationsClientCreateOrupdateResponse, error) {
	var err error
	const operationName = "ServiceConfigurationsClient.CreateOrupdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrupdateCreateRequest(ctx, resourceURI, endpointName, serviceConfigurationName, serviceConfigurationResource, options)
	if err != nil {
		return ServiceConfigurationsClientCreateOrupdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceConfigurationsClientCreateOrupdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ServiceConfigurationsClientCreateOrupdateResponse{}, err
	}
	resp, err := client.createOrupdateHandleResponse(httpResp)
	return resp, err
}

// createOrupdateCreateRequest creates the CreateOrupdate request.
func (client *ServiceConfigurationsClient) createOrupdateCreateRequest(ctx context.Context, resourceURI string, endpointName string, serviceConfigurationName string, serviceConfigurationResource ServiceConfigurationResource, options *ServiceConfigurationsClientCreateOrupdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.HybridConnectivity/endpoints/{endpointName}/serviceConfigurations/{serviceConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", endpointName)
	urlPath = strings.ReplaceAll(urlPath, "{serviceConfigurationName}", serviceConfigurationName)
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, serviceConfigurationResource); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrupdateHandleResponse handles the CreateOrupdate response.
func (client *ServiceConfigurationsClient) createOrupdateHandleResponse(resp *http.Response) (ServiceConfigurationsClientCreateOrupdateResponse, error) {
	result := ServiceConfigurationsClientCreateOrupdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceConfigurationResource); err != nil {
		return ServiceConfigurationsClientCreateOrupdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the service details to the target resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
//   - endpointName - The endpoint name.
//   - serviceConfigurationName - The service name.
//   - options - ServiceConfigurationsClientDeleteOptions contains the optional parameters for the ServiceConfigurationsClient.Delete
//     method.
func (client *ServiceConfigurationsClient) Delete(ctx context.Context, resourceURI string, endpointName string, serviceConfigurationName string, options *ServiceConfigurationsClientDeleteOptions) (ServiceConfigurationsClientDeleteResponse, error) {
	var err error
	const operationName = "ServiceConfigurationsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceURI, endpointName, serviceConfigurationName, options)
	if err != nil {
		return ServiceConfigurationsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceConfigurationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ServiceConfigurationsClientDeleteResponse{}, err
	}
	return ServiceConfigurationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ServiceConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceURI string, endpointName string, serviceConfigurationName string, options *ServiceConfigurationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.HybridConnectivity/endpoints/{endpointName}/serviceConfigurations/{serviceConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", endpointName)
	urlPath = strings.ReplaceAll(urlPath, "{serviceConfigurationName}", serviceConfigurationName)
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the details about the service to the resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
//   - endpointName - The endpoint name.
//   - serviceConfigurationName - The service name.
//   - options - ServiceConfigurationsClientGetOptions contains the optional parameters for the ServiceConfigurationsClient.Get
//     method.
func (client *ServiceConfigurationsClient) Get(ctx context.Context, resourceURI string, endpointName string, serviceConfigurationName string, options *ServiceConfigurationsClientGetOptions) (ServiceConfigurationsClientGetResponse, error) {
	var err error
	const operationName = "ServiceConfigurationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceURI, endpointName, serviceConfigurationName, options)
	if err != nil {
		return ServiceConfigurationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServiceConfigurationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServiceConfigurationsClient) getCreateRequest(ctx context.Context, resourceURI string, endpointName string, serviceConfigurationName string, options *ServiceConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.HybridConnectivity/endpoints/{endpointName}/serviceConfigurations/{serviceConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", endpointName)
	urlPath = strings.ReplaceAll(urlPath, "{serviceConfigurationName}", serviceConfigurationName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServiceConfigurationsClient) getHandleResponse(resp *http.Response) (ServiceConfigurationsClientGetResponse, error) {
	result := ServiceConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceConfigurationResource); err != nil {
		return ServiceConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByEndpointResourcePager - API to enumerate registered services in service configurations under a Endpoint Resource
//
// Generated from API version 2023-03-15
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
//   - endpointName - The endpoint name.
//   - options - ServiceConfigurationsClientListByEndpointResourceOptions contains the optional parameters for the ServiceConfigurationsClient.NewListByEndpointResourcePager
//     method.
func (client *ServiceConfigurationsClient) NewListByEndpointResourcePager(resourceURI string, endpointName string, options *ServiceConfigurationsClientListByEndpointResourceOptions) *runtime.Pager[ServiceConfigurationsClientListByEndpointResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServiceConfigurationsClientListByEndpointResourceResponse]{
		More: func(page ServiceConfigurationsClientListByEndpointResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServiceConfigurationsClientListByEndpointResourceResponse) (ServiceConfigurationsClientListByEndpointResourceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServiceConfigurationsClient.NewListByEndpointResourcePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByEndpointResourceCreateRequest(ctx, resourceURI, endpointName, options)
			}, nil)
			if err != nil {
				return ServiceConfigurationsClientListByEndpointResourceResponse{}, err
			}
			return client.listByEndpointResourceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByEndpointResourceCreateRequest creates the ListByEndpointResource request.
func (client *ServiceConfigurationsClient) listByEndpointResourceCreateRequest(ctx context.Context, resourceURI string, endpointName string, options *ServiceConfigurationsClientListByEndpointResourceOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.HybridConnectivity/endpoints/{endpointName}/serviceConfigurations"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", endpointName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByEndpointResourceHandleResponse handles the ListByEndpointResource response.
func (client *ServiceConfigurationsClient) listByEndpointResourceHandleResponse(resp *http.Response) (ServiceConfigurationsClientListByEndpointResourceResponse, error) {
	result := ServiceConfigurationsClientListByEndpointResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceConfigurationList); err != nil {
		return ServiceConfigurationsClientListByEndpointResourceResponse{}, err
	}
	return result, nil
}

// Update - Update the service details in the service configurations of the target resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15
//   - resourceURI - The fully qualified Azure Resource manager identifier of the resource to be connected.
//   - endpointName - The endpoint name.
//   - serviceConfigurationName - The service name.
//   - serviceConfigurationResource - Service details
//   - options - ServiceConfigurationsClientUpdateOptions contains the optional parameters for the ServiceConfigurationsClient.Update
//     method.
func (client *ServiceConfigurationsClient) Update(ctx context.Context, resourceURI string, endpointName string, serviceConfigurationName string, serviceConfigurationResource ServiceConfigurationResourcePatch, options *ServiceConfigurationsClientUpdateOptions) (ServiceConfigurationsClientUpdateResponse, error) {
	var err error
	const operationName = "ServiceConfigurationsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceURI, endpointName, serviceConfigurationName, serviceConfigurationResource, options)
	if err != nil {
		return ServiceConfigurationsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceConfigurationsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServiceConfigurationsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ServiceConfigurationsClient) updateCreateRequest(ctx context.Context, resourceURI string, endpointName string, serviceConfigurationName string, serviceConfigurationResource ServiceConfigurationResourcePatch, options *ServiceConfigurationsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.HybridConnectivity/endpoints/{endpointName}/serviceConfigurations/{serviceConfigurationName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	urlPath = strings.ReplaceAll(urlPath, "{endpointName}", endpointName)
	urlPath = strings.ReplaceAll(urlPath, "{serviceConfigurationName}", serviceConfigurationName)
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, serviceConfigurationResource); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ServiceConfigurationsClient) updateHandleResponse(resp *http.Response) (ServiceConfigurationsClientUpdateResponse, error) {
	result := ServiceConfigurationsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceConfigurationResource); err != nil {
		return ServiceConfigurationsClientUpdateResponse{}, err
	}
	return result, nil
}
