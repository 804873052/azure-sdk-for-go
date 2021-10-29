//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservices

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ReplicationUsagesClient contains the methods for the ReplicationUsages group.
// Don't use this type directly, use NewReplicationUsagesClient() instead.
type ReplicationUsagesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewReplicationUsagesClient creates a new instance of ReplicationUsagesClient with the specified values.
func NewReplicationUsagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ReplicationUsagesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &ReplicationUsagesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// List - Fetches the replication usages of the vault.
// If the operation fails it returns a generic error.
func (client *ReplicationUsagesClient) List(ctx context.Context, resourceGroupName string, vaultName string, options *ReplicationUsagesListOptions) (ReplicationUsagesListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceGroupName, vaultName, options)
	if err != nil {
		return ReplicationUsagesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReplicationUsagesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationUsagesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ReplicationUsagesClient) listCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *ReplicationUsagesListOptions) (*policy.Request, error) {
	urlPath := "/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/replicationUsages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationUsagesClient) listHandleResponse(resp *http.Response) (ReplicationUsagesListResponse, error) {
	result := ReplicationUsagesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReplicationUsageList); err != nil {
		return ReplicationUsagesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ReplicationUsagesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	if len(body) == 0 {
		return runtime.NewResponseError(errors.New(resp.Status), resp)
	}
	return runtime.NewResponseError(errors.New(string(body)), resp)
}
