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

// IPGroupsServer is a fake server for instances of the armnetwork.IPGroupsClient type.
type IPGroupsServer struct {
	// BeginCreateOrUpdate is the fake for method IPGroupsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, ipGroupsName string, parameters armnetwork.IPGroup, options *armnetwork.IPGroupsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.IPGroupsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method IPGroupsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, ipGroupsName string, options *armnetwork.IPGroupsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.IPGroupsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method IPGroupsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, ipGroupsName string, options *armnetwork.IPGroupsClientGetOptions) (resp azfake.Responder[armnetwork.IPGroupsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method IPGroupsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(options *armnetwork.IPGroupsClientListOptions) (resp azfake.PagerResponder[armnetwork.IPGroupsClientListResponse])

	// NewListByResourceGroupPager is the fake for method IPGroupsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armnetwork.IPGroupsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armnetwork.IPGroupsClientListByResourceGroupResponse])

	// UpdateGroups is the fake for method IPGroupsClient.UpdateGroups
	// HTTP status codes to indicate success: http.StatusOK
	UpdateGroups func(ctx context.Context, resourceGroupName string, ipGroupsName string, parameters armnetwork.TagsObject, options *armnetwork.IPGroupsClientUpdateGroupsOptions) (resp azfake.Responder[armnetwork.IPGroupsClientUpdateGroupsResponse], errResp azfake.ErrorResponder)
}

// NewIPGroupsServerTransport creates a new instance of IPGroupsServerTransport with the provided implementation.
// The returned IPGroupsServerTransport instance is connected to an instance of armnetwork.IPGroupsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewIPGroupsServerTransport(srv *IPGroupsServer) *IPGroupsServerTransport {
	return &IPGroupsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armnetwork.IPGroupsClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armnetwork.IPGroupsClientDeleteResponse]](),
		newListPager:                newTracker[azfake.PagerResponder[armnetwork.IPGroupsClientListResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armnetwork.IPGroupsClientListByResourceGroupResponse]](),
	}
}

// IPGroupsServerTransport connects instances of armnetwork.IPGroupsClient to instances of IPGroupsServer.
// Don't use this type directly, use NewIPGroupsServerTransport instead.
type IPGroupsServerTransport struct {
	srv                         *IPGroupsServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armnetwork.IPGroupsClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armnetwork.IPGroupsClientDeleteResponse]]
	newListPager                *tracker[azfake.PagerResponder[armnetwork.IPGroupsClientListResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armnetwork.IPGroupsClientListByResourceGroupResponse]]
}

// Do implements the policy.Transporter interface for IPGroupsServerTransport.
func (i *IPGroupsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "IPGroupsClient.BeginCreateOrUpdate":
		resp, err = i.dispatchBeginCreateOrUpdate(req)
	case "IPGroupsClient.BeginDelete":
		resp, err = i.dispatchBeginDelete(req)
	case "IPGroupsClient.Get":
		resp, err = i.dispatchGet(req)
	case "IPGroupsClient.NewListPager":
		resp, err = i.dispatchNewListPager(req)
	case "IPGroupsClient.NewListByResourceGroupPager":
		resp, err = i.dispatchNewListByResourceGroupPager(req)
	case "IPGroupsClient.UpdateGroups":
		resp, err = i.dispatchUpdateGroups(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *IPGroupsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := i.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/ipGroups/(?P<ipGroupsName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.IPGroup](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		ipGroupsNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("ipGroupsName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, ipGroupsNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		i.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		i.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		i.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (i *IPGroupsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if i.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := i.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/ipGroups/(?P<ipGroupsName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		ipGroupsNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("ipGroupsName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, ipGroupsNameUnescaped, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		i.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		i.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		i.beginDelete.remove(req)
	}

	return resp, nil
}

func (i *IPGroupsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/ipGroups/(?P<ipGroupsName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
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
	ipGroupsNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("ipGroupsName")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
	var options *armnetwork.IPGroupsClientGetOptions
	if expandParam != nil {
		options = &armnetwork.IPGroupsClientGetOptions{
			Expand: expandParam,
		}
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameUnescaped, ipGroupsNameUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IPGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *IPGroupsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := i.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/ipGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resp := i.srv.NewListPager(nil)
		newListPager = &resp
		i.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.IPGroupsClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		i.newListPager.remove(req)
	}
	return resp, nil
}

func (i *IPGroupsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := i.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/ipGroups`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListByResourceGroupPager(resourceGroupNameUnescaped, nil)
		newListByResourceGroupPager = &resp
		i.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armnetwork.IPGroupsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		i.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		i.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (i *IPGroupsServerTransport) dispatchUpdateGroups(req *http.Request) (*http.Response, error) {
	if i.srv.UpdateGroups == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateGroups not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/ipGroups/(?P<ipGroupsName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetwork.TagsObject](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	ipGroupsNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("ipGroupsName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.UpdateGroups(req.Context(), resourceGroupNameUnescaped, ipGroupsNameUnescaped, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).IPGroup, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
