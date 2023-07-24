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

// InterfaceTapConfigurationsServer is a fake server for instances of the armnetwork.InterfaceTapConfigurationsClient type.
type InterfaceTapConfigurationsServer struct {
	// BeginCreateOrUpdate is the fake for method InterfaceTapConfigurationsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, tapConfigurationParameters armnetwork.InterfaceTapConfiguration, options *armnetwork.InterfaceTapConfigurationsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetwork.InterfaceTapConfigurationsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method InterfaceTapConfigurationsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *armnetwork.InterfaceTapConfigurationsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetwork.InterfaceTapConfigurationsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method InterfaceTapConfigurationsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, networkInterfaceName string, tapConfigurationName string, options *armnetwork.InterfaceTapConfigurationsClientGetOptions) (resp azfake.Responder[armnetwork.InterfaceTapConfigurationsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method InterfaceTapConfigurationsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(resourceGroupName string, networkInterfaceName string, options *armnetwork.InterfaceTapConfigurationsClientListOptions) (resp azfake.PagerResponder[armnetwork.InterfaceTapConfigurationsClientListResponse])
}

// NewInterfaceTapConfigurationsServerTransport creates a new instance of InterfaceTapConfigurationsServerTransport with the provided implementation.
// The returned InterfaceTapConfigurationsServerTransport instance is connected to an instance of armnetwork.InterfaceTapConfigurationsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewInterfaceTapConfigurationsServerTransport(srv *InterfaceTapConfigurationsServer) *InterfaceTapConfigurationsServerTransport {
	return &InterfaceTapConfigurationsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetwork.InterfaceTapConfigurationsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnetwork.InterfaceTapConfigurationsClientDeleteResponse]](),
		newListPager:        newTracker[azfake.PagerResponder[armnetwork.InterfaceTapConfigurationsClientListResponse]](),
	}
}

// InterfaceTapConfigurationsServerTransport connects instances of armnetwork.InterfaceTapConfigurationsClient to instances of InterfaceTapConfigurationsServer.
// Don't use this type directly, use NewInterfaceTapConfigurationsServerTransport instead.
type InterfaceTapConfigurationsServerTransport struct {
	srv                 *InterfaceTapConfigurationsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetwork.InterfaceTapConfigurationsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnetwork.InterfaceTapConfigurationsClientDeleteResponse]]
	newListPager        *tracker[azfake.PagerResponder[armnetwork.InterfaceTapConfigurationsClientListResponse]]
}

// Do implements the policy.Transporter interface for InterfaceTapConfigurationsServerTransport.
func (i *InterfaceTapConfigurationsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "InterfaceTapConfigurationsClient.BeginCreateOrUpdate":
		resp, err = i.dispatchBeginCreateOrUpdate(req)
	case "InterfaceTapConfigurationsClient.BeginDelete":
		resp, err = i.dispatchBeginDelete(req)
	case "InterfaceTapConfigurationsClient.Get":
		resp, err = i.dispatchGet(req)
	case "InterfaceTapConfigurationsClient.NewListPager":
		resp, err = i.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *InterfaceTapConfigurationsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if i.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := i.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/networkInterfaces/(?P<networkInterfaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tapConfigurations/(?P<tapConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetwork.InterfaceTapConfiguration](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkInterfaceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("networkInterfaceName")])
		if err != nil {
			return nil, err
		}
		tapConfigurationNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("tapConfigurationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameUnescaped, networkInterfaceNameUnescaped, tapConfigurationNameUnescaped, body, nil)
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

func (i *InterfaceTapConfigurationsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if i.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := i.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/networkInterfaces/(?P<networkInterfaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tapConfigurations/(?P<tapConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 4 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkInterfaceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("networkInterfaceName")])
		if err != nil {
			return nil, err
		}
		tapConfigurationNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("tapConfigurationName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := i.srv.BeginDelete(req.Context(), resourceGroupNameUnescaped, networkInterfaceNameUnescaped, tapConfigurationNameUnescaped, nil)
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

func (i *InterfaceTapConfigurationsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if i.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/networkInterfaces/(?P<networkInterfaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tapConfigurations/(?P<tapConfigurationName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	networkInterfaceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("networkInterfaceName")])
	if err != nil {
		return nil, err
	}
	tapConfigurationNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("tapConfigurationName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := i.srv.Get(req.Context(), resourceGroupNameUnescaped, networkInterfaceNameUnescaped, tapConfigurationNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).InterfaceTapConfiguration, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i *InterfaceTapConfigurationsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if i.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := i.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Network/networkInterfaces/(?P<networkInterfaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/tapConfigurations`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		networkInterfaceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("networkInterfaceName")])
		if err != nil {
			return nil, err
		}
		resp := i.srv.NewListPager(resourceGroupNameUnescaped, networkInterfaceNameUnescaped, nil)
		newListPager = &resp
		i.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armnetwork.InterfaceTapConfigurationsClientListResponse, createLink func() string) {
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
