//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotcentral

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

// AppsClient contains the methods for the Apps group.
// Don't use this type directly, use NewAppsClient() instead.
type AppsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAppsClient creates a new instance of AppsClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAppsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AppsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AppsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CheckNameAvailability - Check if an IoT Central application name is available.
// If the operation fails it returns an *azcore.ResponseError type.
// operationInputs - Set the name parameter in the OperationInputs structure to the name of the IoT Central application to
// check.
// options - AppsClientCheckNameAvailabilityOptions contains the optional parameters for the AppsClient.CheckNameAvailability
// method.
func (client *AppsClient) CheckNameAvailability(ctx context.Context, operationInputs OperationInputs, options *AppsClientCheckNameAvailabilityOptions) (AppsClientCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, operationInputs, options)
	if err != nil {
		return AppsClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AppsClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AppsClientCheckNameAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *AppsClient) checkNameAvailabilityCreateRequest(ctx context.Context, operationInputs OperationInputs, options *AppsClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTCentral/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, operationInputs)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *AppsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (AppsClientCheckNameAvailabilityResponse, error) {
	result := AppsClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppAvailabilityInfo); err != nil {
		return AppsClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// CheckSubdomainAvailability - Check if an IoT Central application subdomain is available.
// If the operation fails it returns an *azcore.ResponseError type.
// operationInputs - Set the name parameter in the OperationInputs structure to the subdomain of the IoT Central application
// to check.
// options - AppsClientCheckSubdomainAvailabilityOptions contains the optional parameters for the AppsClient.CheckSubdomainAvailability
// method.
func (client *AppsClient) CheckSubdomainAvailability(ctx context.Context, operationInputs OperationInputs, options *AppsClientCheckSubdomainAvailabilityOptions) (AppsClientCheckSubdomainAvailabilityResponse, error) {
	req, err := client.checkSubdomainAvailabilityCreateRequest(ctx, operationInputs, options)
	if err != nil {
		return AppsClientCheckSubdomainAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AppsClientCheckSubdomainAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AppsClientCheckSubdomainAvailabilityResponse{}, runtime.NewResponseError(resp)
	}
	return client.checkSubdomainAvailabilityHandleResponse(resp)
}

// checkSubdomainAvailabilityCreateRequest creates the CheckSubdomainAvailability request.
func (client *AppsClient) checkSubdomainAvailabilityCreateRequest(ctx context.Context, operationInputs OperationInputs, options *AppsClientCheckSubdomainAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTCentral/checkSubdomainAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, operationInputs)
}

// checkSubdomainAvailabilityHandleResponse handles the CheckSubdomainAvailability response.
func (client *AppsClient) checkSubdomainAvailabilityHandleResponse(resp *http.Response) (AppsClientCheckSubdomainAvailabilityResponse, error) {
	result := AppsClientCheckSubdomainAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppAvailabilityInfo); err != nil {
		return AppsClientCheckSubdomainAvailabilityResponse{}, err
	}
	return result, nil
}

// BeginCreateOrUpdate - Create or update the metadata of an IoT Central application. The usual pattern to modify a property
// is to retrieve the IoT Central application metadata and security metadata, and then combine them
// with the modified values in a new body to update the IoT Central application.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the IoT Central application.
// resourceName - The ARM resource name of the IoT Central application.
// app - The IoT Central application metadata and security metadata.
// options - AppsClientBeginCreateOrUpdateOptions contains the optional parameters for the AppsClient.BeginCreateOrUpdate
// method.
func (client *AppsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, app App, options *AppsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[AppsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, resourceName, app, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[AppsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[AppsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update the metadata of an IoT Central application. The usual pattern to modify a property is
// to retrieve the IoT Central application metadata and security metadata, and then combine them
// with the modified values in a new body to update the IoT Central application.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *AppsClient) createOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, app App, options *AppsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, app, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AppsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, app App, options *AppsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/iotApps/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, app)
}

// BeginDelete - Delete an IoT Central application.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the IoT Central application.
// resourceName - The ARM resource name of the IoT Central application.
// options - AppsClientBeginDeleteOptions contains the optional parameters for the AppsClient.BeginDelete method.
func (client *AppsClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, options *AppsClientBeginDeleteOptions) (*armruntime.Poller[AppsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[AppsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[AppsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete an IoT Central application.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *AppsClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, options *AppsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AppsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *AppsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/iotApps/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get the metadata of an IoT Central application.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the IoT Central application.
// resourceName - The ARM resource name of the IoT Central application.
// options - AppsClientGetOptions contains the optional parameters for the AppsClient.Get method.
func (client *AppsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, options *AppsClientGetOptions) (AppsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return AppsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AppsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AppsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AppsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *AppsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/iotApps/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AppsClient) getHandleResponse(resp *http.Response) (AppsClientGetResponse, error) {
	result := AppsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.App); err != nil {
		return AppsClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Get all the IoT Central Applications in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the IoT Central application.
// options - AppsClientListByResourceGroupOptions contains the optional parameters for the AppsClient.ListByResourceGroup
// method.
func (client *AppsClient) ListByResourceGroup(resourceGroupName string, options *AppsClientListByResourceGroupOptions) *runtime.Pager[AppsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PageProcessor[AppsClientListByResourceGroupResponse]{
		More: func(page AppsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AppsClientListByResourceGroupResponse) (AppsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AppsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AppsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AppsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *AppsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *AppsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/iotApps"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *AppsClient) listByResourceGroupHandleResponse(resp *http.Response) (AppsClientListByResourceGroupResponse, error) {
	result := AppsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppListResult); err != nil {
		return AppsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Get all IoT Central Applications in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - AppsClientListBySubscriptionOptions contains the optional parameters for the AppsClient.ListBySubscription method.
func (client *AppsClient) ListBySubscription(options *AppsClientListBySubscriptionOptions) *runtime.Pager[AppsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PageProcessor[AppsClientListBySubscriptionResponse]{
		More: func(page AppsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AppsClientListBySubscriptionResponse) (AppsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AppsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AppsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AppsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *AppsClient) listBySubscriptionCreateRequest(ctx context.Context, options *AppsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTCentral/iotApps"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *AppsClient) listBySubscriptionHandleResponse(resp *http.Response) (AppsClientListBySubscriptionResponse, error) {
	result := AppsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppListResult); err != nil {
		return AppsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// ListTemplates - Get all available application templates.
// If the operation fails it returns an *azcore.ResponseError type.
// options - AppsClientListTemplatesOptions contains the optional parameters for the AppsClient.ListTemplates method.
func (client *AppsClient) ListTemplates(options *AppsClientListTemplatesOptions) *runtime.Pager[AppsClientListTemplatesResponse] {
	return runtime.NewPager(runtime.PageProcessor[AppsClientListTemplatesResponse]{
		More: func(page AppsClientListTemplatesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AppsClientListTemplatesResponse) (AppsClientListTemplatesResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listTemplatesCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AppsClientListTemplatesResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AppsClientListTemplatesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AppsClientListTemplatesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listTemplatesHandleResponse(resp)
		},
	})
}

// listTemplatesCreateRequest creates the ListTemplates request.
func (client *AppsClient) listTemplatesCreateRequest(ctx context.Context, options *AppsClientListTemplatesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTCentral/appTemplates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listTemplatesHandleResponse handles the ListTemplates response.
func (client *AppsClient) listTemplatesHandleResponse(resp *http.Response) (AppsClientListTemplatesResponse, error) {
	result := AppsClientListTemplatesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppTemplatesResult); err != nil {
		return AppsClientListTemplatesResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update the metadata of an IoT Central application.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the IoT Central application.
// resourceName - The ARM resource name of the IoT Central application.
// appPatch - The IoT Central application metadata and security metadata.
// options - AppsClientBeginUpdateOptions contains the optional parameters for the AppsClient.BeginUpdate method.
func (client *AppsClient) BeginUpdate(ctx context.Context, resourceGroupName string, resourceName string, appPatch AppPatch, options *AppsClientBeginUpdateOptions) (*armruntime.Poller[AppsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, resourceName, appPatch, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[AppsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[AppsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - Update the metadata of an IoT Central application.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *AppsClient) update(ctx context.Context, resourceGroupName string, resourceName string, appPatch AppPatch, options *AppsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, appPatch, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *AppsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, appPatch AppPatch, options *AppsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.IoTCentral/iotApps/{resourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, appPatch)
}
