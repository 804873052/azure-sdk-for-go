//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmonitor

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// MetricsClient contains the methods for the Metrics group.
// Don't use this type directly, use NewMetricsClient() instead.
type MetricsClient struct {
	ep string
	pl runtime.Pipeline
}

// NewMetricsClient creates a new instance of MetricsClient with the specified values.
func NewMetricsClient(con *arm.Connection) *MetricsClient {
	return &MetricsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version)}
}

// List - Lists the metric values for a resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MetricsClient) List(ctx context.Context, resourceURI string, options *MetricsListOptions) (MetricsListResponse, error) {
	req, err := client.listCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return MetricsListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MetricsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MetricsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *MetricsClient) listCreateRequest(ctx context.Context, resourceURI string, options *MetricsListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.Insights/metrics"
	if resourceURI == "" {
		return nil, errors.New("parameter resourceURI cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Timespan != nil {
		reqQP.Set("timespan", *options.Timespan)
	}
	if options != nil && options.Interval != nil {
		reqQP.Set("interval", *options.Interval)
	}
	if options != nil && options.Metricnames != nil {
		reqQP.Set("metricnames", *options.Metricnames)
	}
	if options != nil && options.Aggregation != nil {
		reqQP.Set("aggregation", *options.Aggregation)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("orderby", *options.Orderby)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.ResultType != nil {
		reqQP.Set("resultType", string(*options.ResultType))
	}
	reqQP.Set("api-version", "2018-01-01")
	if options != nil && options.Metricnamespace != nil {
		reqQP.Set("metricnamespace", *options.Metricnamespace)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MetricsClient) listHandleResponse(resp *http.Response) (MetricsListResponse, error) {
	result := MetricsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Response); err != nil {
		return MetricsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *MetricsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
