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
	"strconv"
)

// ScopeConnectionsServer is a fake server for instances of the armnetwork.ScopeConnectionsClient type.
type ScopeConnectionsServer struct {
	// CreateOrUpdate is the fake for method ScopeConnectionsClient.CreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	CreateOrUpdate func(ctx context.Context, resourceGroupName string, networkManagerName string, scopeConnectionName string, parameters armnetwork.ScopeConnection, options *armnetwork.ScopeConnectionsClientCreateOrUpdateOptions) (resp azfake.Responder[armnetwork.ScopeConnectionsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// Delete is the fake for method ScopeConnectionsClient.Delete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusNoContent
	Delete func(ctx context.Context, resourceGroupName string, networkManagerName string, scopeConnectionName string, options *armnetwork.ScopeConnectionsClientDeleteOptions) (resp azfake.Responder[armnetwork.ScopeConnectionsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method ScopeConnectionsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkManagerName string, scopeConnectionName string, options *armnetwork.ScopeConnectionsClientGetOptions) (resp azfake.Responder[armnetwork.ScopeConnectionsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method ScopeConnectionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkManagerName string, options *armnetwork.ScopeConnectionsClientListOptions) (resp azfake.PagerResponder[armnetwork.ScopeConnectionsClientListResponse])
}

// NewScopeConnectionsServerTransport creates a new instance of ScopeConnectionsServerTransport with the provided implementation.
// The returned ScopeConnectionsServerTransport instance is connected to an instance of armnetwork.ScopeConnectionsClient by way of the
// undefined.Transporter field.
func NewScopeConnectionsServerTransport(srv *ScopeConnectionsServer) *ScopeConnectionsServerTransport {
	return &ScopeConnectionsServerTransport{srv: srv}
}

// ScopeConnectionsServerTransport connects instances of armnetwork.ScopeConnectionsClient to instances of ScopeConnectionsServer.
// Don't use this type directly, use NewScopeConnectionsServerTransport instead.
type ScopeConnectionsServerTransport struct {
	srv          *ScopeConnectionsServer
	newListPager *azfake.PagerResponder[armnetwork.ScopeConnectionsClientListResponse]
}

// Do implements the policy.Transporter interface for ScopeConnectionsServerTransport.
func (s *ScopeConnectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ScopeConnectionsClient.CreateOrUpdate":
		resp, err = s.dispatchCreateOrUpdate(req)
	case "ScopeConnectionsClient.Delete":
		resp, err = s.dispatchDelete(req)
	case "ScopeConnectionsClient.Get":
		resp, err = s.dispatchGet(req)
	case "ScopeConnectionsClient.NewListPager":
		resp, err = s.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *ScopeConnectionsServerTransport) dispatchCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if s.srv.CreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method CreateOrUpdate not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/scopeConnections/(?P<scopeConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.ScopeConnection](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkManagerNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
	if err != nil {
		return nil, err
	}
	scopeConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("scopeConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.CreateOrUpdate(req.Context(), resourceGroupNameUnescaped, networkManagerNameUnescaped, scopeConnectionNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK, http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ScopeConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScopeConnectionsServerTransport) dispatchDelete(req *http.Request) (*http.Response, error) {
	if s.srv.Delete == nil {
		return nil, &nonRetriableError{errors.New("fake for method Delete not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/scopeConnections/(?P<scopeConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkManagerNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
	if err != nil {
		return nil, err
	}
	scopeConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("scopeConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Delete(req.Context(), resourceGroupNameUnescaped, networkManagerNameUnescaped, scopeConnectionNameUnescaped, nil)
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

func (s *ScopeConnectionsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if s.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/scopeConnections/(?P<scopeConnectionName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkManagerNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
	if err != nil {
		return nil, err
	}
	scopeConnectionNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("scopeConnectionName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := s.srv.Get(req.Context(), resourceGroupNameUnescaped, networkManagerNameUnescaped, scopeConnectionNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).ScopeConnection, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ScopeConnectionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if s.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	if s.newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/networkManagers/(?P<networkManagerName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/scopeConnections`
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
		networkManagerNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("networkManagerName")])
		if err != nil {
			return nil, err
		}
		topUnescaped, err := url.QueryUnescape(qp.Get("$top"))
		if err != nil {
			return nil, err
		}
		topParam, err := parseOptional(topUnescaped, func(v string) (int32, error) {
			p, parseErr := strconv.ParseInt(v, 10, 32)
			if parseErr != nil {
				return 0, parseErr
			}
			return int32(p), nil
		})
		if err != nil {
			return nil, err
		}
		skipTokenUnescaped, err := url.QueryUnescape(qp.Get("$skipToken"))
		if err != nil {
			return nil, err
		}
		skipTokenParam := getOptional(skipTokenUnescaped)
		var options *armnetwork.ScopeConnectionsClientListOptions
		if topParam != nil || skipTokenParam != nil {
			options = &armnetwork.ScopeConnectionsClientListOptions{
				Top:       topParam,
				SkipToken: skipTokenParam,
			}
		}
		resp := s.srv.NewListPager(resourceGroupNameUnescaped, networkManagerNameUnescaped, options)
		s.newListPager = &resp
		server.PagerResponderInjectNextLinks(s.newListPager, req, func(page *armnetwork.ScopeConnectionsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(s.newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(s.newListPager) {
		s.newListPager = nil
	}
	return resp, nil
}
