//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	"net/http"
	"net/url"
	"regexp"
)

// DedicatedHostGroupsServer is a fake server for instances of the armcompute.DedicatedHostGroupsClient type.
type DedicatedHostGroupsServer struct {
	// CreateOrUpdate is the fake for method DedicatedHostGroupsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, hostGroupName string, parameters armcompute.DedicatedHostGroup, options *armcompute.DedicatedHostGroupsClientCreateOrUpdateOptions) (resp azfake.Responder[armcompute.DedicatedHostGroupsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method DedicatedHostGroupsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, hostGroupName string, options *armcompute.DedicatedHostGroupsClientDeleteOptions) (resp azfake.Responder[armcompute.DedicatedHostGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method DedicatedHostGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, hostGroupName string, options *armcompute.DedicatedHostGroupsClientGetOptions) (resp azfake.Responder[armcompute.DedicatedHostGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method DedicatedHostGroupsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armcompute.DedicatedHostGroupsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armcompute.DedicatedHostGroupsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method DedicatedHostGroupsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armcompute.DedicatedHostGroupsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armcompute.DedicatedHostGroupsClientListBySubscriptionResponse])

	// Update is the fake for method DedicatedHostGroupsClient.Update
	// HTTP status codes to indicate success: http.StatusOK
	Update func(ctx context.Context, resourceGroupName string, hostGroupName string, parameters armcompute.DedicatedHostGroupUpdate, options *armcompute.DedicatedHostGroupsClientUpdateOptions) (resp azfake.Responder[armcompute.DedicatedHostGroupsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewDedicatedHostGroupsServerTransport creates a new instance of DedicatedHostGroupsServerTransport with the provided implementation.
// The returned DedicatedHostGroupsServerTransport instance is connected to an instance of armcompute.DedicatedHostGroupsClient by way of the
// undefined.Transporter field.
func NewDedicatedHostGroupsServerTransport(srv *DedicatedHostGroupsServer) *DedicatedHostGroupsServerTransport {
	return &DedicatedHostGroupsServerTransport{srv: srv}
}

// DedicatedHostGroupsServerTransport connects instances of armcompute.DedicatedHostGroupsClient to instances of DedicatedHostGroupsServer.
// Don't use this type directly, use NewDedicatedHostGroupsServerTransport instead.
type DedicatedHostGroupsServerTransport struct {
	srv                         *DedicatedHostGroupsServer
	newListByResourceGroupPager *azfake.PagerResponder[armcompute.DedicatedHostGroupsClientListByResourceGroupResponse]
	newListBySubscriptionPager  *azfake.PagerResponder[armcompute.DedicatedHostGroupsClientListBySubscriptionResponse]
}

// Do implements the policy.Transporter interface for DedicatedHostGroupsServerTransport.
func (d *DedicatedHostGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "DedicatedHostGroupsClient.CreateOrUpdate":
		resp, err = d.dispatchCreateOrUpdate(req)
	case "DedicatedHostGroupsClient.Delete":
		resp, err = d.dispatchDelete(req)
	case "DedicatedHostGroupsClient.Get":
		resp, err = d.dispatchGet(req)
	case "DedicatedHostGroupsClient.NewListByResourceGroupPager":
		resp, err = d.dispatchNewListByResourceGroupPager(req)
	case "DedicatedHostGroupsClient.NewListBySubscriptionPager":
		resp, err = d.dispatchNewListBySubscriptionPager(req)
	case "DedicatedHostGroupsClient.Update":
		resp, err = d.dispatchUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (d *DedicatedHostGroupsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/hostGroups/(?P<hostGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcompute.DedicatedHostGroup](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hostGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("hostGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.CreateOrUpdate(req.Context(), resourceGroupNameUnescaped, hostGroupNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DedicatedHostGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DedicatedHostGroupsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if d.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/hostGroups/(?P<hostGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hostGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("hostGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Delete(req.Context(), resourceGroupNameUnescaped, hostGroupNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DedicatedHostGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if d.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/hostGroups/(?P<hostGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hostGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("hostGroupName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(armcompute.InstanceViewTypes(expandUnescaped))
	var options *armcompute.DedicatedHostGroupsClientGetOptions
	if expandParam != nil {
		options = &armcompute.DedicatedHostGroupsClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := d.srv.Get(req.Context(), resourceGroupNameUnescaped, hostGroupNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DedicatedHostGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (d *DedicatedHostGroupsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	if d.newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/hostGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := d.srv.NewListByResourceGroupPager(resourceGroupNameUnescaped, nil)
		d.newListByResourceGroupPager = &resp
		server.PagerResponderInjectNextLinks(d.newListByResourceGroupPager, req, func(page *armcompute.DedicatedHostGroupsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(d.newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(d.newListByResourceGroupPager) {
		d.newListByResourceGroupPager = nil
	}
	return resp, nil
}

func (d *DedicatedHostGroupsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if d.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	if d.newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/hostGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := d.srv.NewListBySubscriptionPager(nil)
		d.newListBySubscriptionPager = &resp
		server.PagerResponderInjectNextLinks(d.newListBySubscriptionPager, req, func(page *armcompute.DedicatedHostGroupsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(d.newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(d.newListBySubscriptionPager) {
		d.newListBySubscriptionPager = nil
	}
	return resp, nil
}

func (d *DedicatedHostGroupsServerTransport) dispatchUpdate(req *http.Request) (*http.Response, error) {
	if d.srv.Update == nil {
		return nil, &nonRetriableError{errors.New("fake for method Update not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/hostGroups/(?P<hostGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armcompute.DedicatedHostGroupUpdate](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	hostGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("hostGroupName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := d.srv.Update(req.Context(), resourceGroupNameUnescaped, hostGroupNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DedicatedHostGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
