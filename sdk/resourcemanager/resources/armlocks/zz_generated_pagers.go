//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlocks

import (
	"context"
	"net/http"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// AuthorizationOperationsListPager provides operations for iterating over paged responses.
type AuthorizationOperationsListPager struct {
	client    *AuthorizationOperationsClient
	current   AuthorizationOperationsListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AuthorizationOperationsListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AuthorizationOperationsListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AuthorizationOperationsListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationListResult.NextLink == nil || len(*p.current.OperationListResult.NextLink) == 0 {
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
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AuthorizationOperationsListResponse page.
func (p *AuthorizationOperationsListPager) PageResponse() AuthorizationOperationsListResponse {
	return p.current
}

// ManagementLocksListAtResourceGroupLevelPager provides operations for iterating over paged responses.
type ManagementLocksListAtResourceGroupLevelPager struct {
	client    *ManagementLocksClient
	current   ManagementLocksListAtResourceGroupLevelResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagementLocksListAtResourceGroupLevelResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagementLocksListAtResourceGroupLevelPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagementLocksListAtResourceGroupLevelPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ManagementLockListResult.NextLink == nil || len(*p.current.ManagementLockListResult.NextLink) == 0 {
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
		p.err = p.client.listAtResourceGroupLevelHandleError(resp)
		return false
	}
	result, err := p.client.listAtResourceGroupLevelHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ManagementLocksListAtResourceGroupLevelResponse page.
func (p *ManagementLocksListAtResourceGroupLevelPager) PageResponse() ManagementLocksListAtResourceGroupLevelResponse {
	return p.current
}

// ManagementLocksListAtResourceLevelPager provides operations for iterating over paged responses.
type ManagementLocksListAtResourceLevelPager struct {
	client    *ManagementLocksClient
	current   ManagementLocksListAtResourceLevelResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagementLocksListAtResourceLevelResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagementLocksListAtResourceLevelPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagementLocksListAtResourceLevelPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ManagementLockListResult.NextLink == nil || len(*p.current.ManagementLockListResult.NextLink) == 0 {
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
		p.err = p.client.listAtResourceLevelHandleError(resp)
		return false
	}
	result, err := p.client.listAtResourceLevelHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ManagementLocksListAtResourceLevelResponse page.
func (p *ManagementLocksListAtResourceLevelPager) PageResponse() ManagementLocksListAtResourceLevelResponse {
	return p.current
}

// ManagementLocksListAtSubscriptionLevelPager provides operations for iterating over paged responses.
type ManagementLocksListAtSubscriptionLevelPager struct {
	client    *ManagementLocksClient
	current   ManagementLocksListAtSubscriptionLevelResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagementLocksListAtSubscriptionLevelResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagementLocksListAtSubscriptionLevelPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagementLocksListAtSubscriptionLevelPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ManagementLockListResult.NextLink == nil || len(*p.current.ManagementLockListResult.NextLink) == 0 {
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
		p.err = p.client.listAtSubscriptionLevelHandleError(resp)
		return false
	}
	result, err := p.client.listAtSubscriptionLevelHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ManagementLocksListAtSubscriptionLevelResponse page.
func (p *ManagementLocksListAtSubscriptionLevelPager) PageResponse() ManagementLocksListAtSubscriptionLevelResponse {
	return p.current
}

// ManagementLocksListByScopePager provides operations for iterating over paged responses.
type ManagementLocksListByScopePager struct {
	client    *ManagementLocksClient
	current   ManagementLocksListByScopeResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ManagementLocksListByScopeResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ManagementLocksListByScopePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ManagementLocksListByScopePager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ManagementLockListResult.NextLink == nil || len(*p.current.ManagementLockListResult.NextLink) == 0 {
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
		p.err = p.client.listByScopeHandleError(resp)
		return false
	}
	result, err := p.client.listByScopeHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ManagementLocksListByScopeResponse page.
func (p *ManagementLocksListByScopePager) PageResponse() ManagementLocksListByScopeResponse {
	return p.current
}
