package documentdb

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

// RestorableMongodbDatabasesClient is the client for the RestorableMongodbDatabases methods of the Documentdb service.
type RestorableMongodbDatabasesClient struct {
	BaseClient
}

// NewRestorableMongodbDatabasesClient creates an instance of the RestorableMongodbDatabasesClient client.
func NewRestorableMongodbDatabasesClient(subscriptionID string) RestorableMongodbDatabasesClient {
	return NewRestorableMongodbDatabasesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRestorableMongodbDatabasesClientWithBaseURI creates an instance of the RestorableMongodbDatabasesClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewRestorableMongodbDatabasesClientWithBaseURI(baseURI string, subscriptionID string) RestorableMongodbDatabasesClient {
	return RestorableMongodbDatabasesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List show the event feed of all mutations done on all the Azure Cosmos DB MongoDB databases under the restorable
// account.  This helps in scenario where database was accidentally deleted to get the deletion time.  This API
// requires  'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/.../read' permission
// Parameters:
// location - cosmos DB region, with spaces between words and each word capitalized.
// instanceID - the instanceId GUID of a restorable database account.
func (client RestorableMongodbDatabasesClient) List(ctx context.Context, location string, instanceID string) (result RestorableMongodbDatabasesListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RestorableMongodbDatabasesClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("documentdb.RestorableMongodbDatabasesClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, location, instanceID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.RestorableMongodbDatabasesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "documentdb.RestorableMongodbDatabasesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "documentdb.RestorableMongodbDatabasesClient", "List", resp, "Failure responding to request")
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client RestorableMongodbDatabasesClient) ListPreparer(ctx context.Context, location string, instanceID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"instanceId":     autorest.Encode("path", instanceID),
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2022-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableMongodbDatabases", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RestorableMongodbDatabasesClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RestorableMongodbDatabasesClient) ListResponder(resp *http.Response) (result RestorableMongodbDatabasesListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
