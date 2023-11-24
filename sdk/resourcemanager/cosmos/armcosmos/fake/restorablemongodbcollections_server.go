//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cosmos/armcosmos/v2"
	"net/http"
	"net/url"
	"regexp"
)

// RestorableMongodbCollectionsServer is a fake server for instances of the armcosmos.RestorableMongodbCollectionsClient type.
type RestorableMongodbCollectionsServer struct {
	// NewListPager is the fake for method RestorableMongodbCollectionsClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(location string, instanceID string, options *armcosmos.RestorableMongodbCollectionsClientListOptions) (resp azfake.PagerResponder[armcosmos.RestorableMongodbCollectionsClientListResponse])
}

// NewRestorableMongodbCollectionsServerTransport creates a new instance of RestorableMongodbCollectionsServerTransport with the provided implementation.
// The returned RestorableMongodbCollectionsServerTransport instance is connected to an instance of armcosmos.RestorableMongodbCollectionsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewRestorableMongodbCollectionsServerTransport(srv *RestorableMongodbCollectionsServer) *RestorableMongodbCollectionsServerTransport {
	return &RestorableMongodbCollectionsServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armcosmos.RestorableMongodbCollectionsClientListResponse]](),
	}
}

// RestorableMongodbCollectionsServerTransport connects instances of armcosmos.RestorableMongodbCollectionsClient to instances of RestorableMongodbCollectionsServer.
// Don't use this type directly, use NewRestorableMongodbCollectionsServerTransport instead.
type RestorableMongodbCollectionsServerTransport struct {
	srv          *RestorableMongodbCollectionsServer
	newListPager *tracker[azfake.PagerResponder[armcosmos.RestorableMongodbCollectionsClientListResponse]]
}

// Do implements the policy.Transporter interface for RestorableMongodbCollectionsServerTransport.
func (r *RestorableMongodbCollectionsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "RestorableMongodbCollectionsClient.NewListPager":
		resp, err = r.dispatchNewListPager(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *RestorableMongodbCollectionsServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if r.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := r.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DocumentDB/locations/(?P<location>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restorableDatabaseAccounts/(?P<instanceId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/restorableMongodbCollections`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		locationParam, err := url.PathUnescape(matches[regex.SubexpIndex("location")])
		if err != nil {
			return nil, err
		}
		instanceIDParam, err := url.PathUnescape(matches[regex.SubexpIndex("instanceId")])
		if err != nil {
			return nil, err
		}
		restorableMongodbDatabaseRidUnescaped, err := url.QueryUnescape(qp.Get("restorableMongodbDatabaseRid"))
		if err != nil {
			return nil, err
		}
		restorableMongodbDatabaseRidParam := getOptional(restorableMongodbDatabaseRidUnescaped)
		startTimeUnescaped, err := url.QueryUnescape(qp.Get("startTime"))
		if err != nil {
			return nil, err
		}
		startTimeParam := getOptional(startTimeUnescaped)
		endTimeUnescaped, err := url.QueryUnescape(qp.Get("endTime"))
		if err != nil {
			return nil, err
		}
		endTimeParam := getOptional(endTimeUnescaped)
		var options *armcosmos.RestorableMongodbCollectionsClientListOptions
		if restorableMongodbDatabaseRidParam != nil || startTimeParam != nil || endTimeParam != nil {
			options = &armcosmos.RestorableMongodbCollectionsClientListOptions{
				RestorableMongodbDatabaseRid: restorableMongodbDatabaseRidParam,
				StartTime:                    startTimeParam,
				EndTime:                      endTimeParam,
			}
		}
		resp := r.srv.NewListPager(locationParam, instanceIDParam, options)
		newListPager = &resp
		r.newListPager.add(req, newListPager)
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		r.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		r.newListPager.remove(req)
	}
	return resp, nil
}
