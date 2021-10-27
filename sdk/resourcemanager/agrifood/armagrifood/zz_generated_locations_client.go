//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armagrifood

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// LocationsClient contains the methods for the Locations group.
// Don't use this type directly, use NewLocationsClient() instead.
type LocationsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewLocationsClient creates a new instance of LocationsClient with the specified values.
func NewLocationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *LocationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &LocationsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CheckNameAvailability - Checks the name availability of the resource with requested resource name.
// If the operation fails it returns the *ErrorResponse error type.
func (client *LocationsClient) CheckNameAvailability(ctx context.Context, body CheckNameAvailabilityRequest, options *LocationsCheckNameAvailabilityOptions) (LocationsCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, body, options)
	if err != nil {
		return LocationsCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return LocationsCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return LocationsCheckNameAvailabilityResponse{}, client.checkNameAvailabilityHandleError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *LocationsClient) checkNameAvailabilityCreateRequest(ctx context.Context, body CheckNameAvailabilityRequest, options *LocationsCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AgFoodPlatform/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-05-12-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, body)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *LocationsClient) checkNameAvailabilityHandleResponse(resp *http.Response) (LocationsCheckNameAvailabilityResponse, error) {
	result := LocationsCheckNameAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResponse); err != nil {
		return LocationsCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}

// checkNameAvailabilityHandleError handles the CheckNameAvailability error response.
func (client *LocationsClient) checkNameAvailabilityHandleError(resp *http.Response) error {
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
