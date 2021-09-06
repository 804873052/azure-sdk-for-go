package dataprotection

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RecoveryPointsClient is the open API 2.0 Specs for Azure Data Protection service
type RecoveryPointsClient struct {
	BaseClient
}

// NewRecoveryPointsClient creates an instance of the RecoveryPointsClient client.
func NewRecoveryPointsClient(subscriptionID string) RecoveryPointsClient {
	return NewRecoveryPointsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewRecoveryPointsClientWithBaseURI creates an instance of the RecoveryPointsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewRecoveryPointsClientWithBaseURI(baseURI string, subscriptionID string) RecoveryPointsClient {
	return RecoveryPointsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a Recovery Point using recoveryPointId for a Datasource.
// Parameters:
// vaultName - the name of the backup vault.
// resourceGroupName - the name of the resource group where the backup vault is present.
// backupInstanceName - the name of the backup instance
func (client RecoveryPointsClient) Get(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, recoveryPointID string) (result AzureBackupRecoveryPointResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecoveryPointsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, vaultName, resourceGroupName, backupInstanceName, recoveryPointID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.RecoveryPointsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dataprotection.RecoveryPointsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.RecoveryPointsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client RecoveryPointsClient) GetPreparer(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, recoveryPointID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupInstanceName": autorest.Encode("path", backupInstanceName),
		"recoveryPointId":    autorest.Encode("path", recoveryPointID),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"vaultName":          autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupInstances/{backupInstanceName}/recoveryPoints/{recoveryPointId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client RecoveryPointsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client RecoveryPointsClient) GetResponder(resp *http.Response) (result AzureBackupRecoveryPointResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List returns a list of Recovery Points for a DataSource in a vault.
// Parameters:
// vaultName - the name of the backup vault.
// resourceGroupName - the name of the resource group where the backup vault is present.
// backupInstanceName - the name of the backup instance
// filter - oData filter options.
// skipToken - skipToken Filter.
func (client RecoveryPointsClient) List(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, filter string, skipToken string) (result AzureBackupRecoveryPointResourceListPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecoveryPointsClient.List")
		defer func() {
			sc := -1
			if result.abrprl.Response.Response != nil {
				sc = result.abrprl.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, vaultName, resourceGroupName, backupInstanceName, filter, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.RecoveryPointsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.abrprl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "dataprotection.RecoveryPointsClient", "List", resp, "Failure sending request")
		return
	}

	result.abrprl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.RecoveryPointsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.abrprl.hasNextLink() && result.abrprl.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client RecoveryPointsClient) ListPreparer(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, filter string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"backupInstanceName": autorest.Encode("path", backupInstanceName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
		"vaultName":          autorest.Encode("path", vaultName),
	}

	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupInstances/{backupInstanceName}/recoveryPoints", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client RecoveryPointsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client RecoveryPointsClient) ListResponder(resp *http.Response) (result AzureBackupRecoveryPointResourceList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client RecoveryPointsClient) listNextResults(ctx context.Context, lastResults AzureBackupRecoveryPointResourceList) (result AzureBackupRecoveryPointResourceList, err error) {
	req, err := lastResults.azureBackupRecoveryPointResourceListPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "dataprotection.RecoveryPointsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "dataprotection.RecoveryPointsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataprotection.RecoveryPointsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client RecoveryPointsClient) ListComplete(ctx context.Context, vaultName string, resourceGroupName string, backupInstanceName string, filter string, skipToken string) (result AzureBackupRecoveryPointResourceListIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RecoveryPointsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, vaultName, resourceGroupName, backupInstanceName, filter, skipToken)
	return
}
