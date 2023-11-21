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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/delegatednetwork/armdelegatednetwork"
	"net/http"
	"net/url"
	"regexp"
)

// ControllerServer is a fake server for instances of the armdelegatednetwork.ControllerClient type.
type ControllerServer struct {
	// BeginCreate is the fake for method ControllerClient.BeginCreate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreate func(ctx context.Context, resourceGroupName string, resourceName string, parameters armdelegatednetwork.DelegatedController, options *armdelegatednetwork.ControllerClientBeginCreateOptions) (resp azfake.PollerResponder[armdelegatednetwork.ControllerClientCreateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method ControllerClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, resourceName string, options *armdelegatednetwork.ControllerClientBeginDeleteOptions) (resp azfake.PollerResponder[armdelegatednetwork.ControllerClientDeleteResponse], errResp azfake.ErrorResponder)

	// GetDetails is the fake for method ControllerClient.GetDetails
	// HTTP status codes to indicate success: http.StatusOK
	GetDetails func(ctx context.Context, resourceGroupName string, resourceName string, options *armdelegatednetwork.ControllerClientGetDetailsOptions) (resp azfake.Responder[armdelegatednetwork.ControllerClientGetDetailsResponse], errResp azfake.ErrorResponder)

	// Patch is the fake for method ControllerClient.Patch
	// HTTP status codes to indicate success: http.StatusOK
	Patch func(ctx context.Context, resourceGroupName string, resourceName string, parameters armdelegatednetwork.ControllerResourceUpdateParameters, options *armdelegatednetwork.ControllerClientPatchOptions) (resp azfake.Responder[armdelegatednetwork.ControllerClientPatchResponse], errResp azfake.ErrorResponder)
}

// NewControllerServerTransport creates a new instance of ControllerServerTransport with the provided implementation.
// The returned ControllerServerTransport instance is connected to an instance of armdelegatednetwork.ControllerClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewControllerServerTransport(srv *ControllerServer) *ControllerServerTransport {
	return &ControllerServerTransport{
		srv:         srv,
		beginCreate: newTracker[azfake.PollerResponder[armdelegatednetwork.ControllerClientCreateResponse]](),
		beginDelete: newTracker[azfake.PollerResponder[armdelegatednetwork.ControllerClientDeleteResponse]](),
	}
}

// ControllerServerTransport connects instances of armdelegatednetwork.ControllerClient to instances of ControllerServer.
// Don't use this type directly, use NewControllerServerTransport instead.
type ControllerServerTransport struct {
	srv         *ControllerServer
	beginCreate *tracker[azfake.PollerResponder[armdelegatednetwork.ControllerClientCreateResponse]]
	beginDelete *tracker[azfake.PollerResponder[armdelegatednetwork.ControllerClientDeleteResponse]]
}

// Do implements the policy.Transporter interface for ControllerServerTransport.
func (c *ControllerServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "ControllerClient.BeginCreate":
		resp, err = c.dispatchBeginCreate(req)
	case "ControllerClient.BeginDelete":
		resp, err = c.dispatchBeginDelete(req)
	case "ControllerClient.GetDetails":
		resp, err = c.dispatchGetDetails(req)
	case "ControllerClient.Patch":
		resp, err = c.dispatchPatch(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *ControllerServerTransport) dispatchBeginCreate(req *http.Request) (*http.Response, error) {
	if c.srv.BeginCreate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreate not implemented")}
	}
	beginCreate := c.beginCreate.get(req)
	if beginCreate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DelegatedNetwork/controller/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdelegatednetwork.DelegatedController](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginCreate(req.Context(), resourceGroupNameParam, resourceNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreate = &respr
		c.beginCreate.add(req, beginCreate)
	}

	resp, err := server.PollerResponderNext(beginCreate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		c.beginCreate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreate) {
		c.beginCreate.remove(req)
	}

	return resp, nil
}

func (c *ControllerServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if c.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := c.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DelegatedNetwork/controller/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := c.srv.BeginDelete(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		c.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		c.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		c.beginDelete.remove(req)
	}

	return resp, nil
}

func (c *ControllerServerTransport) dispatchGetDetails(req *http.Request) (*http.Response, error) {
	if c.srv.GetDetails == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetDetails not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DelegatedNetwork/controller/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.GetDetails(req.Context(), resourceGroupNameParam, resourceNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DelegatedController, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ControllerServerTransport) dispatchPatch(req *http.Request) (*http.Response, error) {
	if c.srv.Patch == nil {
		return nil, &nonRetriableError{errors.New("fake for method Patch not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DelegatedNetwork/controller/(?P<resourceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armdelegatednetwork.ControllerResourceUpdateParameters](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	resourceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := c.srv.Patch(req.Context(), resourceGroupNameParam, resourceNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).DelegatedController, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
