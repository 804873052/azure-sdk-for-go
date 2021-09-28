//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// RestorableSQLDatabasesClient contains the methods for the RestorableSQLDatabases group.
// Don't use this type directly, use NewRestorableSQLDatabasesClient() instead.
type RestorableSQLDatabasesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewRestorableSQLDatabasesClient creates a new instance of RestorableSQLDatabasesClient with the specified values.
func NewRestorableSQLDatabasesClient(con *arm.Connection, subscriptionID string) *RestorableSQLDatabasesClient {
	return &RestorableSQLDatabasesClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// List - Show the event feed of all mutations done on all the Azure Cosmos DB SQL databases under the restorable account. This helps in scenario where
// database was accidentally deleted to get the deletion
// time. This API requires 'Microsoft.DocumentDB/locations/restorableDatabaseAccounts/…/read' permission
// If the operation fails it returns the *CloudError error type.
func (client *RestorableSQLDatabasesClient) List(ctx context.Context, location string, instanceID string, options *RestorableSQLDatabasesListOptions) (RestorableSQLDatabasesListResponse, error) {
	req, err := client.listCreateRequest(ctx, location, instanceID, options)
	if err != nil {
		return RestorableSQLDatabasesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return RestorableSQLDatabasesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RestorableSQLDatabasesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *RestorableSQLDatabasesClient) listCreateRequest(ctx context.Context, location string, instanceID string, options *RestorableSQLDatabasesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/locations/{location}/restorableDatabaseAccounts/{instanceId}/restorableSqlDatabases"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if instanceID == "" {
		return nil, errors.New("parameter instanceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instanceId}", url.PathEscape(instanceID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RestorableSQLDatabasesClient) listHandleResponse(resp *http.Response) (RestorableSQLDatabasesListResponse, error) {
	result := RestorableSQLDatabasesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableSQLDatabasesListResult); err != nil {
		return RestorableSQLDatabasesListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *RestorableSQLDatabasesClient) listHandleError(resp *http.Response) error {
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
