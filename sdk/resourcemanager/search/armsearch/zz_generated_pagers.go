//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsearch

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// PrivateEndpointConnectionsListByServicePager provides operations for iterating over paged responses.
type PrivateEndpointConnectionsListByServicePager struct {
	client    *PrivateEndpointConnectionsClient
	current   PrivateEndpointConnectionsListByServiceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateEndpointConnectionsListByServiceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateEndpointConnectionsListByServicePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateEndpointConnectionsListByServicePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateEndpointConnectionListResult.NextLink == nil || len(*p.current.PrivateEndpointConnectionListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByServiceHandleError(resp)
		return false
	}
	result, err := p.client.listByServiceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PrivateEndpointConnectionsListByServiceResponse page.
func (p *PrivateEndpointConnectionsListByServicePager) PageResponse() PrivateEndpointConnectionsListByServiceResponse {
	return p.current
}

// QueryKeysListBySearchServicePager provides operations for iterating over paged responses.
type QueryKeysListBySearchServicePager struct {
	client    *QueryKeysClient
	current   QueryKeysListBySearchServiceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, QueryKeysListBySearchServiceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *QueryKeysListBySearchServicePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *QueryKeysListBySearchServicePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ListQueryKeysResult.NextLink == nil || len(*p.current.ListQueryKeysResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listBySearchServiceHandleError(resp)
		return false
	}
	result, err := p.client.listBySearchServiceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current QueryKeysListBySearchServiceResponse page.
func (p *QueryKeysListBySearchServicePager) PageResponse() QueryKeysListBySearchServiceResponse {
	return p.current
}

// ServicesListByResourceGroupPager provides operations for iterating over paged responses.
type ServicesListByResourceGroupPager struct {
	client    *ServicesClient
	current   ServicesListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ServicesListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ServicesListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ServicesListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SearchServiceListResult.NextLink == nil || len(*p.current.SearchServiceListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByResourceGroupHandleError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ServicesListByResourceGroupResponse page.
func (p *ServicesListByResourceGroupPager) PageResponse() ServicesListByResourceGroupResponse {
	return p.current
}

// ServicesListBySubscriptionPager provides operations for iterating over paged responses.
type ServicesListBySubscriptionPager struct {
	client    *ServicesClient
	current   ServicesListBySubscriptionResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ServicesListBySubscriptionResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ServicesListBySubscriptionPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ServicesListBySubscriptionPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SearchServiceListResult.NextLink == nil || len(*p.current.SearchServiceListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listBySubscriptionHandleError(resp)
		return false
	}
	result, err := p.client.listBySubscriptionHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ServicesListBySubscriptionResponse page.
func (p *ServicesListBySubscriptionPager) PageResponse() ServicesListBySubscriptionResponse {
	return p.current
}

// SharedPrivateLinkResourcesListByServicePager provides operations for iterating over paged responses.
type SharedPrivateLinkResourcesListByServicePager struct {
	client    *SharedPrivateLinkResourcesClient
	current   SharedPrivateLinkResourcesListByServiceResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SharedPrivateLinkResourcesListByServiceResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SharedPrivateLinkResourcesListByServicePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SharedPrivateLinkResourcesListByServicePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SharedPrivateLinkResourceListResult.NextLink == nil || len(*p.current.SharedPrivateLinkResourceListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByServiceHandleError(resp)
		return false
	}
	result, err := p.client.listByServiceHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SharedPrivateLinkResourcesListByServiceResponse page.
func (p *SharedPrivateLinkResourcesListByServicePager) PageResponse() SharedPrivateLinkResourcesListByServiceResponse {
	return p.current
}
