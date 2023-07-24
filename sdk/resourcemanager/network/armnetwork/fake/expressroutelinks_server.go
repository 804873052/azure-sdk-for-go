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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/network/armnetwork/v3"
	"net/http"
	"net/url"
	"regexp"
)

// ExpressRouteLinksServer is a fake server for instances of the armnetwork.ExpressRouteLinksClient type.
type ExpressRouteLinksServer struct {
	// Get is the fake for method ExpressRouteLinksClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, expressRoutePortName string, linkName string, options *armnetwork.ExpressRouteLinksClientGetOptions) (resp azfake.Responder[armnetwork.ExpressRouteLinksClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ExpressRouteLinksClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, expressRoutePortName string, options *armnetwork.ExpressRouteLinksClientListOptions) (resp azfake.PagerResponder[armnetwork.ExpressRouteLinksClientListResponse])
}

// NewExpressRouteLinksServerTransport creates a new instance of ExpressRouteLinksServerTransport with the provided implementation.
// The returned ExpressRouteLinksServerTransport instance is connected to an instance of armnetwork.ExpressRouteLinksClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewExpressRouteLinksServerTransport(srv *ExpressRouteLinksServer) *ExpressRouteLinksServerTransport {
	return &ExpressRouteLinksServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armnetwork.ExpressRouteLinksClientListResponse]](),
	}
}

// ExpressRouteLinksServerTransport connects instances of armnetwork.ExpressRouteLinksClient to instances of ExpressRouteLinksServer.
// Don't use this type directly, use NewExpressRouteLinksServerTransport instead.
type ExpressRouteLinksServerTransport struct {
	srv          *ExpressRouteLinksServer
	newListPager *tracker[azfake.PagerResponder[armnetwork.ExpressRouteLinksClientListResponse]]
}

// Do implements the policy.Transporter interface for ExpressRouteLinksServerTransport.
func (e *ExpressRouteLinksServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ExpressRouteLinksClient.Get":
		resp, err = e.dispatchGet(req)
	case "ExpressRouteLinksClient.NewListPager":
		resp, err = e.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (e *ExpressRouteLinksServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if e.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/ExpressRoutePorts/(?P<expressRoutePortName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/links/(?P<linkName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	expressRoutePortNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("expressRoutePortName")])
	if err != nil {
		return nil, err
	}
	linkNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("linkName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := e.srv.Get(req.Context(), resourceGroupNameUnescaped, expressRoutePortNameUnescaped, linkNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ExpressRouteLink, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (e *ExpressRouteLinksServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if e.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := e.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/ExpressRoutePorts/(?P<expressRoutePortName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/links`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		expressRoutePortNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("expressRoutePortName")])
		if err != nil {
			return nil, err
		}
		resp := e.srv.NewListPager(resourceGroupNameUnescaped, expressRoutePortNameUnescaped, nil)
		newListPager = &resp
		e.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.ExpressRouteLinksClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		e.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		e.newListPager.remove(req)
	}
	return resp, nil
}
