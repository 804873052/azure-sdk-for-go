//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmanagementgroups

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
)

// EntitiesClient contains the methods for the Entities group.
// Don't use this type directly, use NewEntitiesClient() instead.
type EntitiesClient struct {
	internal *arm.Client
}

// NewEntitiesClient creates a new instance of EntitiesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEntitiesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*EntitiesClient, error) {
	cl, err := arm.NewClient(moduleName+".EntitiesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EntitiesClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - List all entities (Management Groups, Subscriptions, etc.) for the authenticated user.
//
// Generated from API version 2021-04-01
//   - options - EntitiesClientListOptions contains the optional parameters for the EntitiesClient.NewListPager method.
func (client *EntitiesClient) NewListPager(options *EntitiesClientListOptions) *runtime.Pager[EntitiesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[EntitiesClientListResponse]{
		More: func(page EntitiesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EntitiesClientListResponse) (EntitiesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EntitiesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return EntitiesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EntitiesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *EntitiesClient) listCreateRequest(ctx context.Context, options *EntitiesClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Management/getEntities"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-04-01")
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Select != nil {
		reqQP.Set("$select", *options.Select)
	}
	if options != nil && options.Search != nil {
		reqQP.Set("$search", string(*options.Search))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.View != nil {
		reqQP.Set("$view", string(*options.View))
	}
	if options != nil && options.GroupName != nil {
		reqQP.Set("groupName", *options.GroupName)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.CacheControl != nil {
		req.Raw().Header["Cache-Control"] = []string{*options.CacheControl}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EntitiesClient) listHandleResponse(resp *http.Response) (EntitiesClientListResponse, error) {
	result := EntitiesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EntityListResult); err != nil {
		return EntitiesClientListResponse{}, err
	}
	return result, nil
}
