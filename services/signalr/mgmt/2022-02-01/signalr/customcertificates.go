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

// CustomCertificatesClient is the REST API for Azure SignalR Service
type CustomCertificatesClient struct {
	BaseClient
}

// NewCustomCertificatesClient creates an instance of the CustomCertificatesClient client.
func NewCustomCertificatesClient(subscriptionID string) CustomCertificatesClient {
	return NewCustomCertificatesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCustomCertificatesClientWithBaseURI creates an instance of the CustomCertificatesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewCustomCertificatesClientWithBaseURI(baseURI string, subscriptionID string) CustomCertificatesClient {
	return CustomCertificatesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a custom certificate.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// resourceName - the name of the resource.
// certificateName - custom certificate name
func (client CustomCertificatesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, certificateName string, parameters CustomCertificate) (result CustomCertificatesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomCertificatesClient.CreateOrUpdate")
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
			Constraints: []validation.Constraint{{Target: "parameters.CustomCertificateProperties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.CustomCertificateProperties.KeyVaultBaseURI", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.CustomCertificateProperties.KeyVaultSecretName", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("signalr.CustomCertificatesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, resourceName, certificateName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.CustomCertificatesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.CustomCertificatesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client CustomCertificatesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, certificateName string, parameters CustomCertificate) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":   autorest.Encode("path", certificateName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}/customCertificates/{certificateName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client CustomCertificatesClient) CreateOrUpdateSender(req *http.Request) (future CustomCertificatesCreateOrUpdateFuture, err error) {
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
func (client CustomCertificatesClient) CreateOrUpdateResponder(resp *http.Response) (result CustomCertificate, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a custom certificate.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// resourceName - the name of the resource.
// certificateName - custom certificate name
func (client CustomCertificatesClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, certificateName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomCertificatesClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName, certificateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.CustomCertificatesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "signalr.CustomCertificatesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.CustomCertificatesClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client CustomCertificatesClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string, certificateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":   autorest.Encode("path", certificateName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-02-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}/customCertificates/{certificateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client CustomCertificatesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client CustomCertificatesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a custom certificate.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// resourceName - the name of the resource.
// certificateName - custom certificate name
func (client CustomCertificatesClient) Get(ctx context.Context, resourceGroupName string, resourceName string, certificateName string) (result CustomCertificate, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomCertificatesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, resourceName, certificateName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.CustomCertificatesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "signalr.CustomCertificatesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.CustomCertificatesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client CustomCertificatesClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceName string, certificateName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"certificateName":   autorest.Encode("path", certificateName),
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}/customCertificates/{certificateName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CustomCertificatesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CustomCertificatesClient) GetResponder(resp *http.Response) (result CustomCertificate, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list all custom certificates.
// Parameters:
// resourceGroupName - the name of the resource group that contains the resource. You can obtain this value
// from the Azure Resource Manager API or the portal.
// resourceName - the name of the resource.
func (client CustomCertificatesClient) List(ctx context.Context, resourceGroupName string, resourceName string) (result CustomCertificateListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomCertificatesClient.List")
		defer func() {
			sc := -1
			if result.ccl.Response.Response != nil {
				sc = result.ccl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.CustomCertificatesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.ccl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "signalr.CustomCertificatesClient", "List", resp, "Failure sending request")
		return
	}

	result.ccl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.CustomCertificatesClient", "List", resp, "Failure responding to request")
		return
	}
	if result.ccl.hasNextLink() && result.ccl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client CustomCertificatesClient) ListPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.SignalRService/signalR/{resourceName}/customCertificates", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client CustomCertificatesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client CustomCertificatesClient) ListResponder(resp *http.Response) (result CustomCertificateList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client CustomCertificatesClient) listNextResults(ctx context.Context, lastResults CustomCertificateList) (result CustomCertificateList, err error) {
	req, err := lastResults.customCertificateListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "signalr.CustomCertificatesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "signalr.CustomCertificatesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "signalr.CustomCertificatesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client CustomCertificatesClient) ListComplete(ctx context.Context, resourceGroupName string, resourceName string) (result CustomCertificateListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomCertificatesClient.List")
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
