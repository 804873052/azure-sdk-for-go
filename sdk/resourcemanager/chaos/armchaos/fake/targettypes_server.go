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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/chaos/armchaos"
	"net/http"
	"net/url"
	"regexp"
)

// TargetTypesServer is a fake server for instances of the armchaos.TargetTypesClient type.
type TargetTypesServer struct {
	// Get is the fake for method TargetTypesClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, locationName string, targetTypeName string, options *armchaos.TargetTypesClientGetOptions) (resp azfake.Responder[armchaos.TargetTypesClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method TargetTypesClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(locationName string, options *armchaos.TargetTypesClientListOptions) (resp azfake.PagerResponder[armchaos.TargetTypesClientListResponse])
}

// NewTargetTypesServerTransport creates a new instance of TargetTypesServerTransport with the provided implementation.
// The returned TargetTypesServerTransport instance is connected to an instance of armchaos.TargetTypesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewTargetTypesServerTransport(srv *TargetTypesServer) *TargetTypesServerTransport {
	return &TargetTypesServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armchaos.TargetTypesClientListResponse]](),
	}
}

// TargetTypesServerTransport connects instances of armchaos.TargetTypesClient to instances of TargetTypesServer.
// Don't use this type directly, use NewTargetTypesServerTransport instead.
type TargetTypesServerTransport struct {
	srv          *TargetTypesServer
	newListPager *tracker[azfake.PagerResponder[armchaos.TargetTypesClientListResponse]]
}

// Do implements the policy.Transporter interface for TargetTypesServerTransport.
func (t *TargetTypesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "TargetTypesClient.Get":
		resp, err = t.dispatchGet(req)
	case "TargetTypesClient.NewListPager":
		resp, err = t.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TargetTypesServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if t.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Chaos/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/targetTypes/(?P<targetTypeName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
	if err != nil {
		return nil, err
	}
	targetTypeNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("targetTypeName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := t.srv.Get(req.Context(), locationNameParam, targetTypeNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).TargetType, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *TargetTypesServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if t.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := t.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.Chaos/locations/(?P<locationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/targetTypes`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		locationNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("locationName")])
		if err != nil {
			return nil, err
		}
		continuationTokenUnescaped, err := url.QueryUnescape(qp.Get("continuationToken"))
		if err != nil {
			return nil, err
		}
		continuationTokenParam := getOptional(continuationTokenUnescaped)
		var options *armchaos.TargetTypesClientListOptions
		if continuationTokenParam != nil {
			options = &armchaos.TargetTypesClientListOptions{
				ContinuationToken: continuationTokenParam,
			}
		}
		resp := t.srv.NewListPager(locationNameParam, options)
		newListPager = &resp
		t.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armchaos.TargetTypesClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		t.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		t.newListPager.remove(req)
	}
	return resp, nil
}