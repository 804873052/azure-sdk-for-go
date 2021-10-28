//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatalakeanalytics

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

// StorageAccountsClient contains the methods for the StorageAccounts group.
// Don't use this type directly, use NewStorageAccountsClient() instead.
type StorageAccountsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewStorageAccountsClient creates a new instance of StorageAccountsClient with the specified values.
func NewStorageAccountsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *StorageAccountsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &StorageAccountsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// Add - Updates the specified Data Lake Analytics account to add an Azure Storage account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *StorageAccountsClient) Add(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, parameters AddStorageAccountParameters, options *StorageAccountsAddOptions) (StorageAccountsAddResponse, error) {
	req, err := client.addCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, parameters, options)
	if err != nil {
		return StorageAccountsAddResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageAccountsAddResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageAccountsAddResponse{}, client.addHandleError(resp)
	}
	return StorageAccountsAddResponse{RawResponse: resp}, nil
}

// addCreateRequest creates the Add request.
func (client *StorageAccountsClient) addCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, parameters AddStorageAccountParameters, options *StorageAccountsAddOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// addHandleError handles the Add error response.
func (client *StorageAccountsClient) addHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Updates the specified Data Lake Analytics account to remove an Azure Storage account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *StorageAccountsClient) Delete(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsDeleteOptions) (StorageAccountsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, options)
	if err != nil {
		return StorageAccountsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageAccountsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageAccountsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return StorageAccountsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *StorageAccountsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *StorageAccountsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets the specified Azure Storage account linked to the given Data Lake Analytics account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *StorageAccountsClient) Get(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsGetOptions) (StorageAccountsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, options)
	if err != nil {
		return StorageAccountsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageAccountsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageAccountsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *StorageAccountsClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StorageAccountsClient) getHandleResponse(resp *http.Response) (StorageAccountsGetResponse, error) {
	result := StorageAccountsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageAccountInformation); err != nil {
		return StorageAccountsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *StorageAccountsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetStorageContainer - Gets the specified Azure Storage container associated with the given Data Lake Analytics and Azure Storage accounts.
// If the operation fails it returns the *ErrorResponse error type.
func (client *StorageAccountsClient) GetStorageContainer(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string, options *StorageAccountsGetStorageContainerOptions) (StorageAccountsGetStorageContainerResponse, error) {
	req, err := client.getStorageContainerCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, containerName, options)
	if err != nil {
		return StorageAccountsGetStorageContainerResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageAccountsGetStorageContainerResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageAccountsGetStorageContainerResponse{}, client.getStorageContainerHandleError(resp)
	}
	return client.getStorageContainerHandleResponse(resp)
}

// getStorageContainerCreateRequest creates the GetStorageContainer request.
func (client *StorageAccountsClient) getStorageContainerCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string, options *StorageAccountsGetStorageContainerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}/containers/{containerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getStorageContainerHandleResponse handles the GetStorageContainer response.
func (client *StorageAccountsClient) getStorageContainerHandleResponse(resp *http.Response) (StorageAccountsGetStorageContainerResponse, error) {
	result := StorageAccountsGetStorageContainerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageContainer); err != nil {
		return StorageAccountsGetStorageContainerResponse{}, err
	}
	return result, nil
}

// getStorageContainerHandleError handles the GetStorageContainer error response.
func (client *StorageAccountsClient) getStorageContainerHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByAccount - Gets the first page of Azure Storage accounts, if any, linked to the specified Data Lake Analytics account. The response includes a link
// to the next page, if any.
// If the operation fails it returns the *ErrorResponse error type.
func (client *StorageAccountsClient) ListByAccount(resourceGroupName string, accountName string, options *StorageAccountsListByAccountOptions) *StorageAccountsListByAccountPager {
	return &StorageAccountsListByAccountPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByAccountCreateRequest(ctx, resourceGroupName, accountName, options)
		},
		advancer: func(ctx context.Context, resp StorageAccountsListByAccountResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.StorageAccountInformationListResult.NextLink)
		},
	}
}

// listByAccountCreateRequest creates the ListByAccount request.
func (client *StorageAccountsClient) listByAccountCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *StorageAccountsListByAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Count != nil {
		reqQP.Set("$count", strconv.FormatBool(*options.Count))
	}
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAccountHandleResponse handles the ListByAccount response.
func (client *StorageAccountsClient) listByAccountHandleResponse(resp *http.Response) (StorageAccountsListByAccountResponse, error) {
	result := StorageAccountsListByAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageAccountInformationListResult); err != nil {
		return StorageAccountsListByAccountResponse{}, err
	}
	return result, nil
}

// listByAccountHandleError handles the ListByAccount error response.
func (client *StorageAccountsClient) listByAccountHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListSasTokens - Gets the SAS token associated with the specified Data Lake Analytics and Azure Storage account and container combination.
// If the operation fails it returns the *ErrorResponse error type.
func (client *StorageAccountsClient) ListSasTokens(resourceGroupName string, accountName string, storageAccountName string, containerName string, options *StorageAccountsListSasTokensOptions) *StorageAccountsListSasTokensPager {
	return &StorageAccountsListSasTokensPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listSasTokensCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, containerName, options)
		},
		advancer: func(ctx context.Context, resp StorageAccountsListSasTokensResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.SasTokenInformationListResult.NextLink)
		},
	}
}

// listSasTokensCreateRequest creates the ListSasTokens request.
func (client *StorageAccountsClient) listSasTokensCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, containerName string, options *StorageAccountsListSasTokensOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}/containers/{containerName}/listSasTokens"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listSasTokensHandleResponse handles the ListSasTokens response.
func (client *StorageAccountsClient) listSasTokensHandleResponse(resp *http.Response) (StorageAccountsListSasTokensResponse, error) {
	result := StorageAccountsListSasTokensResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SasTokenInformationListResult); err != nil {
		return StorageAccountsListSasTokensResponse{}, err
	}
	return result, nil
}

// listSasTokensHandleError handles the ListSasTokens error response.
func (client *StorageAccountsClient) listSasTokensHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListStorageContainers - Lists the Azure Storage containers, if any, associated with the specified Data Lake Analytics and Azure Storage account combination.
// The response includes a link to the next page of results, if any.
// If the operation fails it returns the *ErrorResponse error type.
func (client *StorageAccountsClient) ListStorageContainers(resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsListStorageContainersOptions) *StorageAccountsListStorageContainersPager {
	return &StorageAccountsListStorageContainersPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listStorageContainersCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, options)
		},
		advancer: func(ctx context.Context, resp StorageAccountsListStorageContainersResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.StorageContainerListResult.NextLink)
		},
	}
}

// listStorageContainersCreateRequest creates the ListStorageContainers request.
func (client *StorageAccountsClient) listStorageContainersCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsListStorageContainersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}/containers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listStorageContainersHandleResponse handles the ListStorageContainers response.
func (client *StorageAccountsClient) listStorageContainersHandleResponse(resp *http.Response) (StorageAccountsListStorageContainersResponse, error) {
	result := StorageAccountsListStorageContainersResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageContainerListResult); err != nil {
		return StorageAccountsListStorageContainersResponse{}, err
	}
	return result, nil
}

// listStorageContainersHandleError handles the ListStorageContainers error response.
func (client *StorageAccountsClient) listStorageContainersHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Update - Updates the Data Lake Analytics account to replace Azure Storage blob account details, such as the access key and/or suffix.
// If the operation fails it returns the *ErrorResponse error type.
func (client *StorageAccountsClient) Update(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsUpdateOptions) (StorageAccountsUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, storageAccountName, options)
	if err != nil {
		return StorageAccountsUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageAccountsUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageAccountsUpdateResponse{}, client.updateHandleError(resp)
	}
	return StorageAccountsUpdateResponse{RawResponse: resp}, nil
}

// updateCreateRequest creates the Update request.
func (client *StorageAccountsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, storageAccountName string, options *StorageAccountsUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/storageAccounts/{storageAccountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if storageAccountName == "" {
		return nil, errors.New("parameter storageAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageAccountName}", url.PathEscape(storageAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, *options.Parameters)
	}
	return req, nil
}

// updateHandleError handles the Update error response.
func (client *StorageAccountsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
