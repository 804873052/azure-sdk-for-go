//go:build go1.16
// +build go1.16

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.4.3, generator: @autorest/go@4.0.0-preview.27)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysql

import (
	"context"
	"net/http"
	"reflect"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// AdvisorsListByServerPager provides operations for iterating over paged responses.
type AdvisorsListByServerPager struct {
	client    *AdvisorsClient
	current   AdvisorsListByServerResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, AdvisorsListByServerResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *AdvisorsListByServerPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *AdvisorsListByServerPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AdvisorsResultList.NextLink == nil || len(*p.current.AdvisorsResultList.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByServerHandleError(resp)
		return false
	}
	result, err := p.client.listByServerHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current AdvisorsListByServerResponse page.
func (p *AdvisorsListByServerPager) PageResponse() AdvisorsListByServerResponse {
	return p.current
}

// LocationBasedRecommendedActionSessionsResultListPager provides operations for iterating over paged responses.
type LocationBasedRecommendedActionSessionsResultListPager struct {
	client    *LocationBasedRecommendedActionSessionsResultClient
	current   LocationBasedRecommendedActionSessionsResultListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, LocationBasedRecommendedActionSessionsResultListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *LocationBasedRecommendedActionSessionsResultListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *LocationBasedRecommendedActionSessionsResultListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RecommendationActionsResultList.NextLink == nil || len(*p.current.RecommendationActionsResultList.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
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

// PageResponse returns the current LocationBasedRecommendedActionSessionsResultListResponse page.
func (p *LocationBasedRecommendedActionSessionsResultListPager) PageResponse() LocationBasedRecommendedActionSessionsResultListResponse {
	return p.current
}

// PrivateEndpointConnectionsListByServerPager provides operations for iterating over paged responses.
type PrivateEndpointConnectionsListByServerPager struct {
	client    *PrivateEndpointConnectionsClient
	current   PrivateEndpointConnectionsListByServerResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateEndpointConnectionsListByServerResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateEndpointConnectionsListByServerPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateEndpointConnectionsListByServerPager) NextPage(ctx context.Context) bool {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByServerHandleError(resp)
		return false
	}
	result, err := p.client.listByServerHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PrivateEndpointConnectionsListByServerResponse page.
func (p *PrivateEndpointConnectionsListByServerPager) PageResponse() PrivateEndpointConnectionsListByServerResponse {
	return p.current
}

// PrivateLinkResourcesListByServerPager provides operations for iterating over paged responses.
type PrivateLinkResourcesListByServerPager struct {
	client    *PrivateLinkResourcesClient
	current   PrivateLinkResourcesListByServerResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, PrivateLinkResourcesListByServerResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PrivateLinkResourcesListByServerPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PrivateLinkResourcesListByServerPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.PrivateLinkResourceListResult.NextLink == nil || len(*p.current.PrivateLinkResourceListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByServerHandleError(resp)
		return false
	}
	result, err := p.client.listByServerHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PrivateLinkResourcesListByServerResponse page.
func (p *PrivateLinkResourcesListByServerPager) PageResponse() PrivateLinkResourcesListByServerResponse {
	return p.current
}

// QueryTextsListByServerPager provides operations for iterating over paged responses.
type QueryTextsListByServerPager struct {
	client    *QueryTextsClient
	current   QueryTextsListByServerResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, QueryTextsListByServerResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *QueryTextsListByServerPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *QueryTextsListByServerPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.QueryTextsResultList.NextLink == nil || len(*p.current.QueryTextsResultList.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByServerHandleError(resp)
		return false
	}
	result, err := p.client.listByServerHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current QueryTextsListByServerResponse page.
func (p *QueryTextsListByServerPager) PageResponse() QueryTextsListByServerResponse {
	return p.current
}

// RecommendedActionsListByServerPager provides operations for iterating over paged responses.
type RecommendedActionsListByServerPager struct {
	client    *RecommendedActionsClient
	current   RecommendedActionsListByServerResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, RecommendedActionsListByServerResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *RecommendedActionsListByServerPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *RecommendedActionsListByServerPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RecommendationActionsResultList.NextLink == nil || len(*p.current.RecommendationActionsResultList.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByServerHandleError(resp)
		return false
	}
	result, err := p.client.listByServerHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current RecommendedActionsListByServerResponse page.
func (p *RecommendedActionsListByServerPager) PageResponse() RecommendedActionsListByServerResponse {
	return p.current
}

// ServerKeysListPager provides operations for iterating over paged responses.
type ServerKeysListPager struct {
	client    *ServerKeysClient
	current   ServerKeysListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ServerKeysListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ServerKeysListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ServerKeysListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ServerKeyListResult.NextLink == nil || len(*p.current.ServerKeyListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
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

// PageResponse returns the current ServerKeysListResponse page.
func (p *ServerKeysListPager) PageResponse() ServerKeysListResponse {
	return p.current
}

// ServerSecurityAlertPoliciesListByServerPager provides operations for iterating over paged responses.
type ServerSecurityAlertPoliciesListByServerPager struct {
	client    *ServerSecurityAlertPoliciesClient
	current   ServerSecurityAlertPoliciesListByServerResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, ServerSecurityAlertPoliciesListByServerResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *ServerSecurityAlertPoliciesListByServerPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *ServerSecurityAlertPoliciesListByServerPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ServerSecurityAlertPolicyListResult.NextLink == nil || len(*p.current.ServerSecurityAlertPolicyListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByServerHandleError(resp)
		return false
	}
	result, err := p.client.listByServerHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current ServerSecurityAlertPoliciesListByServerResponse page.
func (p *ServerSecurityAlertPoliciesListByServerPager) PageResponse() ServerSecurityAlertPoliciesListByServerResponse {
	return p.current
}

// TopQueryStatisticsListByServerPager provides operations for iterating over paged responses.
type TopQueryStatisticsListByServerPager struct {
	client    *TopQueryStatisticsClient
	current   TopQueryStatisticsListByServerResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, TopQueryStatisticsListByServerResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *TopQueryStatisticsListByServerPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *TopQueryStatisticsListByServerPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.TopQueryStatisticsResultList.NextLink == nil || len(*p.current.TopQueryStatisticsResultList.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByServerHandleError(resp)
		return false
	}
	result, err := p.client.listByServerHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current TopQueryStatisticsListByServerResponse page.
func (p *TopQueryStatisticsListByServerPager) PageResponse() TopQueryStatisticsListByServerResponse {
	return p.current
}

// VirtualNetworkRulesListByServerPager provides operations for iterating over paged responses.
type VirtualNetworkRulesListByServerPager struct {
	client    *VirtualNetworkRulesClient
	current   VirtualNetworkRulesListByServerResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, VirtualNetworkRulesListByServerResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *VirtualNetworkRulesListByServerPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *VirtualNetworkRulesListByServerPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.VirtualNetworkRuleListResult.NextLink == nil || len(*p.current.VirtualNetworkRuleListResult.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByServerHandleError(resp)
		return false
	}
	result, err := p.client.listByServerHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current VirtualNetworkRulesListByServerResponse page.
func (p *VirtualNetworkRulesListByServerPager) PageResponse() VirtualNetworkRulesListByServerResponse {
	return p.current
}

// WaitStatisticsListByServerPager provides operations for iterating over paged responses.
type WaitStatisticsListByServerPager struct {
	client    *WaitStatisticsClient
	current   WaitStatisticsListByServerResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, WaitStatisticsListByServerResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *WaitStatisticsListByServerPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *WaitStatisticsListByServerPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.WaitStatisticsResultList.NextLink == nil || len(*p.current.WaitStatisticsResultList.NextLink) == 0 {
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
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = p.client.listByServerHandleError(resp)
		return false
	}
	result, err := p.client.listByServerHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current WaitStatisticsListByServerResponse page.
func (p *WaitStatisticsListByServerPager) PageResponse() WaitStatisticsListByServerResponse {
	return p.current
}
