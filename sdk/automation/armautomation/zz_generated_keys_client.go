//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// KeysClient contains the methods for the Keys group.
// Don't use this type directly, use NewKeysClient() instead.
type KeysClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewKeysClient creates a new instance of KeysClient with the specified values.
func NewKeysClient(con *arm.Connection, subscriptionID string) *KeysClient {
	return &KeysClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// ListByAutomationAccount - Retrieve the automation keys for an account.
// If the operation fails it returns the *ErrorResponse error type.
func (client *KeysClient) ListByAutomationAccount(ctx context.Context, resourceGroupName string, automationAccountName string, options *KeysListByAutomationAccountOptions) (KeysListByAutomationAccountResponse, error) {
	req, err := client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
	if err != nil {
		return KeysListByAutomationAccountResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return KeysListByAutomationAccountResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return KeysListByAutomationAccountResponse{}, client.listByAutomationAccountHandleError(resp)
	}
	return client.listByAutomationAccountHandleResponse(resp)
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *KeysClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *KeysListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/listKeys"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *KeysClient) listByAutomationAccountHandleResponse(resp *http.Response) (KeysListByAutomationAccountResponse, error) {
	result := KeysListByAutomationAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.KeyListResult); err != nil {
		return KeysListByAutomationAccountResponse{}, err
	}
	return result, nil
}

// listByAutomationAccountHandleError handles the ListByAutomationAccount error response.
func (client *KeysClient) listByAutomationAccountHandleError(resp *http.Response) error {
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
