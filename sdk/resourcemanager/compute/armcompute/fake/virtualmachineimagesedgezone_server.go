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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute/v5"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// VirtualMachineImagesEdgeZoneServer is a fake server for instances of the armcompute.VirtualMachineImagesEdgeZoneClient type.
type VirtualMachineImagesEdgeZoneServer struct {
	// Get is the fake for method VirtualMachineImagesEdgeZoneClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, location string, edgeZone string, publisherName string, offer string, skus string, version string, options *armcompute.VirtualMachineImagesEdgeZoneClientGetOptions) (resp azfake.Responder[armcompute.VirtualMachineImagesEdgeZoneClientGetResponse], errResp azfake.ErrorResponder)

	// List is the fake for method VirtualMachineImagesEdgeZoneClient.List
	// HTTP status codes to indicate success: http.StatusOK
	List func(ctx context.Context, location string, edgeZone string, publisherName string, offer string, skus string, options *armcompute.VirtualMachineImagesEdgeZoneClientListOptions) (resp azfake.Responder[armcompute.VirtualMachineImagesEdgeZoneClientListResponse], errResp azfake.ErrorResponder)

	// ListOffers is the fake for method VirtualMachineImagesEdgeZoneClient.ListOffers
	// HTTP status codes to indicate success: http.StatusOK
	ListOffers func(ctx context.Context, location string, edgeZone string, publisherName string, options *armcompute.VirtualMachineImagesEdgeZoneClientListOffersOptions) (resp azfake.Responder[armcompute.VirtualMachineImagesEdgeZoneClientListOffersResponse], errResp azfake.ErrorResponder)

	// ListPublishers is the fake for method VirtualMachineImagesEdgeZoneClient.ListPublishers
	// HTTP status codes to indicate success: http.StatusOK
	ListPublishers func(ctx context.Context, location string, edgeZone string, options *armcompute.VirtualMachineImagesEdgeZoneClientListPublishersOptions) (resp azfake.Responder[armcompute.VirtualMachineImagesEdgeZoneClientListPublishersResponse], errResp azfake.ErrorResponder)

	// ListSKUs is the fake for method VirtualMachineImagesEdgeZoneClient.ListSKUs
	// HTTP status codes to indicate success: http.StatusOK
	ListSKUs func(ctx context.Context, location string, edgeZone string, publisherName string, offer string, options *armcompute.VirtualMachineImagesEdgeZoneClientListSKUsOptions) (resp azfake.Responder[armcompute.VirtualMachineImagesEdgeZoneClientListSKUsResponse], errResp azfake.ErrorResponder)
}

// NewVirtualMachineImagesEdgeZoneServerTransport creates a new instance of VirtualMachineImagesEdgeZoneServerTransport with the provided implementation.
// The returned VirtualMachineImagesEdgeZoneServerTransport instance is connected to an instance of armcompute.VirtualMachineImagesEdgeZoneClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewVirtualMachineImagesEdgeZoneServerTransport(srv *VirtualMachineImagesEdgeZoneServer) *VirtualMachineImagesEdgeZoneServerTransport {
	return &VirtualMachineImagesEdgeZoneServerTransport{srv: srv}
}

// VirtualMachineImagesEdgeZoneServerTransport connects instances of armcompute.VirtualMachineImagesEdgeZoneClient to instances of VirtualMachineImagesEdgeZoneServer.
// Don't use this type directly, use NewVirtualMachineImagesEdgeZoneServerTransport instead.
type VirtualMachineImagesEdgeZoneServerTransport struct {
	srv *VirtualMachineImagesEdgeZoneServer
}

// Do implements the policy.Transporter interface for VirtualMachineImagesEdgeZoneServerTransport.
func (v *VirtualMachineImagesEdgeZoneServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "VirtualMachineImagesEdgeZoneClient.Get":
		resp, err = v.dispatchGet(req)
	case "VirtualMachineImagesEdgeZoneClient.List":
		resp, err = v.dispatchList(req)
	case "VirtualMachineImagesEdgeZoneClient.ListOffers":
		resp, err = v.dispatchListOffers(req)
	case "VirtualMachineImagesEdgeZoneClient.ListPublishers":
		resp, err = v.dispatchListPublishers(req)
	case "VirtualMachineImagesEdgeZoneClient.ListSKUs":
		resp, err = v.dispatchListSKUs(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (v *VirtualMachineImagesEdgeZoneServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if v.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/edgeZones/(?P<edgeZone>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifacttypes/vmimage/offers/(?P<offer>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/skus/(?P<skus>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions/(?P<version>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 7 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	edgeZoneUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("edgeZone")])
	if err != nil {
		return nil, err
	}
	publisherNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
	if err != nil {
		return nil, err
	}
	offerUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("offer")])
	if err != nil {
		return nil, err
	}
	skusUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("skus")])
	if err != nil {
		return nil, err
	}
	versionUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("version")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.Get(req.Context(), locationUnescaped, edgeZoneUnescaped, publisherNameUnescaped, offerUnescaped, skusUnescaped, versionUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineImage, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineImagesEdgeZoneServerTransport) dispatchList(req *http.Request) (*http.Response, error) {
	if v.srv.List == nil {
		return nil, &nonRetriableError{errors.New("fake for method List not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/edgeZones/(?P<edgeZone>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifacttypes/vmimage/offers/(?P<offer>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/skus/(?P<skus>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/versions`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 6 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	qp := req.URL.Query()
	locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	edgeZoneUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("edgeZone")])
	if err != nil {
		return nil, err
	}
	publisherNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
	if err != nil {
		return nil, err
	}
	offerUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("offer")])
	if err != nil {
		return nil, err
	}
	skusUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("skus")])
	if err != nil {
		return nil, err
	}
	expandUnescaped, err := url.QueryUnescape(qp.Get("$expand"))
	if err != nil {
		return nil, err
	}
	expandParam := getOptional(expandUnescaped)
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
	orderbyUnescaped, err := url.QueryUnescape(qp.Get("$orderby"))
	if err != nil {
		return nil, err
	}
	orderbyParam := getOptional(orderbyUnescaped)
	var options *armcompute.VirtualMachineImagesEdgeZoneClientListOptions
	if expandParam != nil || topParam != nil || orderbyParam != nil {
		options = &armcompute.VirtualMachineImagesEdgeZoneClientListOptions{
			Expand:  expandParam,
			Top:     topParam,
			Orderby: orderbyParam,
		}
	}
	respr, errRespr := v.srv.List(req.Context(), locationUnescaped, edgeZoneUnescaped, publisherNameUnescaped, offerUnescaped, skusUnescaped, options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineImageResourceArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineImagesEdgeZoneServerTransport) dispatchListOffers(req *http.Request) (*http.Response, error) {
	if v.srv.ListOffers == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListOffers not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/edgeZones/(?P<edgeZone>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifacttypes/vmimage/offers`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 4 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	edgeZoneUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("edgeZone")])
	if err != nil {
		return nil, err
	}
	publisherNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.ListOffers(req.Context(), locationUnescaped, edgeZoneUnescaped, publisherNameUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineImageResourceArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineImagesEdgeZoneServerTransport) dispatchListPublishers(req *http.Request) (*http.Response, error) {
	if v.srv.ListPublishers == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListPublishers not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/edgeZones/(?P<edgeZone>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publishers`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	edgeZoneUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("edgeZone")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.ListPublishers(req.Context(), locationUnescaped, edgeZoneUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineImageResourceArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (v *VirtualMachineImagesEdgeZoneServerTransport) dispatchListSKUs(req *http.Request) (*http.Response, error) {
	if v.srv.ListSKUs == nil {
		return nil, &nonRetriableError{errors.New("fake for method ListSKUs not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft.Compute/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/edgeZones/(?P<edgeZone>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publishers/(?P<publisherName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/artifacttypes/vmimage/offers/(?P<offer>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/skus`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 5 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	locationUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
	if err != nil {
		return nil, err
	}
	edgeZoneUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("edgeZone")])
	if err != nil {
		return nil, err
	}
	publisherNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("publisherName")])
	if err != nil {
		return nil, err
	}
	offerUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("offer")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := v.srv.ListSKUs(req.Context(), locationUnescaped, edgeZoneUnescaped, publisherNameUnescaped, offerUnescaped, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).VirtualMachineImageResourceArray, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
