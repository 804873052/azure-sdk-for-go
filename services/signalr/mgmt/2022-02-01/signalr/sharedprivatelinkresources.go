package signalr

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SharedPrivateLinkResourcesClient is the REST API for Azure SignalR Service
type SharedPrivateLinkResourcesClient struct {
	BaseClient
}

// NewSharedPrivateLinkResourcesClient creates an instance of the SharedPrivateLinkResourcesClient client.
func NewSharedPrivateLinkResourcesClient(subscriptionID string) SharedPrivateLinkResourcesClient {
	return NewSharedPrivateLinkResourcesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSharedPrivateLinkResourcesClientWithBaseURI creates an instance of the SharedPrivateLinkResourcesClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewSharedPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) SharedPrivateLinkResourcesClient {
	return SharedPrivateLinkResourcesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a shared private link resource
// Parameters:
// sharedPrivateLinkResourceName - the name of the shared private link resource
// parameters - the shared private link resource
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// resourceName - the name of the resource.
func (client SharedPrivateLinkResourcesClient) CreateOrUpdate(ctx context.Context, sharedPrivateLinkResourceName string, parameters SharedPrivateLinkResource, resourceGroupName string, resourceName string) (result SharedPrivateLinkResourcesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharedPrivateLinkResourcesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.SharedPrivateLinkResourceProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.SharedPrivateLinkResourceProperties.GroupID", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.SharedPrivateLinkResourceProperties.PrivateLinkResourceID", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("signalr.SharedPrivateLinkResourcesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, sharedPrivateLinkResourceName, parameters, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.SharedPrivateLinkResourcesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.SharedPrivateLinkResourcesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SharedPrivateLinkResourcesClient) CreateOrUpdatePreparer(ctx context.Context, sharedPrivateLinkResourceName string, parameters SharedPrivateLinkResource, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"resourceName":                  autorest.Encode("path", resourceName),
		"sharedPrivateLinkResourceName": autorest.Encode("path", sharedPrivateLinkResourceName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}/sharedPrivateLinkResources/{sharedPrivateLinkResourceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SharedPrivateLinkResourcesClient) CreateOrUpdateSender(req *http.Request) (future SharedPrivateLinkResourcesCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client SharedPrivateLinkResourcesClient) CreateOrUpdateResponder(resp *http.Response) (result SharedPrivateLinkResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete the specified shared private link resource
// Parameters:
// sharedPrivateLinkResourceName - the name of the shared private link resource
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// resourceName - the name of the resource.
func (client SharedPrivateLinkResourcesClient) Delete(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string) (result SharedPrivateLinkResourcesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharedPrivateLinkResourcesClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, sharedPrivateLinkResourceName, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.SharedPrivateLinkResourcesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.SharedPrivateLinkResourcesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SharedPrivateLinkResourcesClient) DeletePreparer(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"resourceName":                  autorest.Encode("path", resourceName),
		"sharedPrivateLinkResourceName": autorest.Encode("path", sharedPrivateLinkResourceName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}/sharedPrivateLinkResources/{sharedPrivateLinkResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SharedPrivateLinkResourcesClient) DeleteSender(req *http.Request) (future SharedPrivateLinkResourcesDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SharedPrivateLinkResourcesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get the specified shared private link resource
// Parameters:
// sharedPrivateLinkResourceName - the name of the shared private link resource
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// resourceName - the name of the resource.
func (client SharedPrivateLinkResourcesClient) Get(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string) (result SharedPrivateLinkResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharedPrivateLinkResourcesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, sharedPrivateLinkResourceName, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.SharedPrivateLinkResourcesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "signalr.SharedPrivateLinkResourcesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.SharedPrivateLinkResourcesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SharedPrivateLinkResourcesClient) GetPreparer(ctx context.Context, sharedPrivateLinkResourceName string, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"resourceName":                  autorest.Encode("path", resourceName),
		"sharedPrivateLinkResourceName": autorest.Encode("path", sharedPrivateLinkResourceName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}/sharedPrivateLinkResources/{sharedPrivateLinkResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SharedPrivateLinkResourcesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SharedPrivateLinkResourcesClient) GetResponder(resp *http.Response) (result SharedPrivateLinkResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list shared private link resources
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// resourceName - the name of the resource.
func (client SharedPrivateLinkResourcesClient) List(ctx context.Context, resourceGroupName string, resourceName string) (result SharedPrivateLinkResourceListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharedPrivateLinkResourcesClient.List")
		defer func() {
			sc := -1
			if result.splrl.Response.Response != nil {
				sc = result.splrl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.SharedPrivateLinkResourcesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.splrl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "signalr.SharedPrivateLinkResourcesClient", "List", resp, "Failure sending request")
		return
	}

	result.splrl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.SharedPrivateLinkResourcesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.splrl.hasNextLink() && result.splrl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client SharedPrivateLinkResourcesClient) ListPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}/sharedPrivateLinkResources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client SharedPrivateLinkResourcesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client SharedPrivateLinkResourcesClient) ListResponder(resp *http.Response) (result SharedPrivateLinkResourceList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client SharedPrivateLinkResourcesClient) listNextResults(ctx context.Context, lastResults SharedPrivateLinkResourceList) (result SharedPrivateLinkResourceList, err error) {
	req, err := lastResults.sharedPrivateLinkResourceListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "signalr.SharedPrivateLinkResourcesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "signalr.SharedPrivateLinkResourcesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.SharedPrivateLinkResourcesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client SharedPrivateLinkResourcesClient) ListComplete(ctx context.Context, resourceGroupName string, resourceName string) (result SharedPrivateLinkResourceListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharedPrivateLinkResourcesClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, resourceGroupName, resourceName)
	return
}
