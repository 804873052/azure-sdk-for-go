//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmarketplaceordering

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// MarketplaceAgreementsClient contains the methods for the MarketplaceAgreements group.
// Don't use this type directly, use NewMarketplaceAgreementsClient() instead.
type MarketplaceAgreementsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMarketplaceAgreementsClient creates a new instance of MarketplaceAgreementsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMarketplaceAgreementsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MarketplaceAgreementsClient, error) {
	cl, err := arm.NewClient(moduleName+".MarketplaceAgreementsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MarketplaceAgreementsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Cancel - Cancel marketplace terms.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-01-01
//   - publisherID - Publisher identifier string of image being deployed.
//   - offerID - Offer identifier string of image being deployed.
//   - planID - Plan identifier string of image being deployed.
//   - options - MarketplaceAgreementsClientCancelOptions contains the optional parameters for the MarketplaceAgreementsClient.Cancel
//     method.
func (client *MarketplaceAgreementsClient) Cancel(ctx context.Context, publisherID string, offerID string, planID string, options *MarketplaceAgreementsClientCancelOptions) (MarketplaceAgreementsClientCancelResponse, error) {
	req, err := client.cancelCreateRequest(ctx, publisherID, offerID, planID, options)
	if err != nil {
		return MarketplaceAgreementsClientCancelResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MarketplaceAgreementsClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MarketplaceAgreementsClientCancelResponse{}, runtime.NewResponseError(resp)
	}
	return client.cancelHandleResponse(resp)
}

// cancelCreateRequest creates the Cancel request.
func (client *MarketplaceAgreementsClient) cancelCreateRequest(ctx context.Context, publisherID string, offerID string, planID string, options *MarketplaceAgreementsClientCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/agreements/{publisherId}/offers/{offerId}/plans/{planId}/cancel"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if publisherID == "" {
		return nil, errors.New("parameter publisherID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherId}", url.PathEscape(publisherID))
	if offerID == "" {
		return nil, errors.New("parameter offerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offerId}", url.PathEscape(offerID))
	if planID == "" {
		return nil, errors.New("parameter planID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planId}", url.PathEscape(planID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// cancelHandleResponse handles the Cancel response.
func (client *MarketplaceAgreementsClient) cancelHandleResponse(resp *http.Response) (MarketplaceAgreementsClientCancelResponse, error) {
	result := MarketplaceAgreementsClientCancelResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgreementTerms); err != nil {
		return MarketplaceAgreementsClientCancelResponse{}, err
	}
	return result, nil
}

// Create - Save marketplace terms.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-01-01
//   - offerType - Offer Type, currently only virtualmachine type is supported.
//   - publisherID - Publisher identifier string of image being deployed.
//   - offerID - Offer identifier string of image being deployed.
//   - planID - Plan identifier string of image being deployed.
//   - parameters - Parameters supplied to the Create Marketplace Terms operation.
//   - options - MarketplaceAgreementsClientCreateOptions contains the optional parameters for the MarketplaceAgreementsClient.Create
//     method.
func (client *MarketplaceAgreementsClient) Create(ctx context.Context, offerType OfferType, publisherID string, offerID string, planID string, parameters AgreementTerms, options *MarketplaceAgreementsClientCreateOptions) (MarketplaceAgreementsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, offerType, publisherID, offerID, planID, parameters, options)
	if err != nil {
		return MarketplaceAgreementsClientCreateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MarketplaceAgreementsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MarketplaceAgreementsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *MarketplaceAgreementsClient) createCreateRequest(ctx context.Context, offerType OfferType, publisherID string, offerID string, planID string, parameters AgreementTerms, options *MarketplaceAgreementsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/offerTypes/{offerType}/publishers/{publisherId}/offers/{offerId}/plans/{planId}/agreements/current"
	if offerType == "" {
		return nil, errors.New("parameter offerType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offerType}", url.PathEscape(string(offerType)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if publisherID == "" {
		return nil, errors.New("parameter publisherID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherId}", url.PathEscape(publisherID))
	if offerID == "" {
		return nil, errors.New("parameter offerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offerId}", url.PathEscape(offerID))
	if planID == "" {
		return nil, errors.New("parameter planID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planId}", url.PathEscape(planID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createHandleResponse handles the Create response.
func (client *MarketplaceAgreementsClient) createHandleResponse(resp *http.Response) (MarketplaceAgreementsClientCreateResponse, error) {
	result := MarketplaceAgreementsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgreementTerms); err != nil {
		return MarketplaceAgreementsClientCreateResponse{}, err
	}
	return result, nil
}

// Get - Get marketplace terms.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-01-01
//   - offerType - Offer Type, currently only virtualmachine type is supported.
//   - publisherID - Publisher identifier string of image being deployed.
//   - offerID - Offer identifier string of image being deployed.
//   - planID - Plan identifier string of image being deployed.
//   - options - MarketplaceAgreementsClientGetOptions contains the optional parameters for the MarketplaceAgreementsClient.Get
//     method.
func (client *MarketplaceAgreementsClient) Get(ctx context.Context, offerType OfferType, publisherID string, offerID string, planID string, options *MarketplaceAgreementsClientGetOptions) (MarketplaceAgreementsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, offerType, publisherID, offerID, planID, options)
	if err != nil {
		return MarketplaceAgreementsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MarketplaceAgreementsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MarketplaceAgreementsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MarketplaceAgreementsClient) getCreateRequest(ctx context.Context, offerType OfferType, publisherID string, offerID string, planID string, options *MarketplaceAgreementsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/offerTypes/{offerType}/publishers/{publisherId}/offers/{offerId}/plans/{planId}/agreements/current"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if offerType == "" {
		return nil, errors.New("parameter offerType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offerType}", url.PathEscape(string(offerType)))
	if publisherID == "" {
		return nil, errors.New("parameter publisherID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherId}", url.PathEscape(publisherID))
	if offerID == "" {
		return nil, errors.New("parameter offerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offerId}", url.PathEscape(offerID))
	if planID == "" {
		return nil, errors.New("parameter planID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planId}", url.PathEscape(planID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MarketplaceAgreementsClient) getHandleResponse(resp *http.Response) (MarketplaceAgreementsClientGetResponse, error) {
	result := MarketplaceAgreementsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgreementTerms); err != nil {
		return MarketplaceAgreementsClientGetResponse{}, err
	}
	return result, nil
}

// GetAgreement - Get marketplace agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-01-01
//   - publisherID - Publisher identifier string of image being deployed.
//   - offerID - Offer identifier string of image being deployed.
//   - planID - Plan identifier string of image being deployed.
//   - options - MarketplaceAgreementsClientGetAgreementOptions contains the optional parameters for the MarketplaceAgreementsClient.GetAgreement
//     method.
func (client *MarketplaceAgreementsClient) GetAgreement(ctx context.Context, publisherID string, offerID string, planID string, options *MarketplaceAgreementsClientGetAgreementOptions) (MarketplaceAgreementsClientGetAgreementResponse, error) {
	req, err := client.getAgreementCreateRequest(ctx, publisherID, offerID, planID, options)
	if err != nil {
		return MarketplaceAgreementsClientGetAgreementResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MarketplaceAgreementsClientGetAgreementResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MarketplaceAgreementsClientGetAgreementResponse{}, runtime.NewResponseError(resp)
	}
	return client.getAgreementHandleResponse(resp)
}

// getAgreementCreateRequest creates the GetAgreement request.
func (client *MarketplaceAgreementsClient) getAgreementCreateRequest(ctx context.Context, publisherID string, offerID string, planID string, options *MarketplaceAgreementsClientGetAgreementOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/agreements/{publisherId}/offers/{offerId}/plans/{planId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if publisherID == "" {
		return nil, errors.New("parameter publisherID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherId}", url.PathEscape(publisherID))
	if offerID == "" {
		return nil, errors.New("parameter offerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offerId}", url.PathEscape(offerID))
	if planID == "" {
		return nil, errors.New("parameter planID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planId}", url.PathEscape(planID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getAgreementHandleResponse handles the GetAgreement response.
func (client *MarketplaceAgreementsClient) getAgreementHandleResponse(resp *http.Response) (MarketplaceAgreementsClientGetAgreementResponse, error) {
	result := MarketplaceAgreementsClientGetAgreementResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgreementTerms); err != nil {
		return MarketplaceAgreementsClientGetAgreementResponse{}, err
	}
	return result, nil
}

// List - List marketplace agreements in the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-01-01
//   - options - MarketplaceAgreementsClientListOptions contains the optional parameters for the MarketplaceAgreementsClient.List
//     method.
func (client *MarketplaceAgreementsClient) List(ctx context.Context, options *MarketplaceAgreementsClientListOptions) (MarketplaceAgreementsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return MarketplaceAgreementsClientListResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MarketplaceAgreementsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MarketplaceAgreementsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *MarketplaceAgreementsClient) listCreateRequest(ctx context.Context, options *MarketplaceAgreementsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/agreements"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MarketplaceAgreementsClient) listHandleResponse(resp *http.Response) (MarketplaceAgreementsClientListResponse, error) {
	result := MarketplaceAgreementsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgreementTermsArray); err != nil {
		return MarketplaceAgreementsClientListResponse{}, err
	}
	return result, nil
}

// Sign - Sign marketplace terms.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-01-01
//   - publisherID - Publisher identifier string of image being deployed.
//   - offerID - Offer identifier string of image being deployed.
//   - planID - Plan identifier string of image being deployed.
//   - options - MarketplaceAgreementsClientSignOptions contains the optional parameters for the MarketplaceAgreementsClient.Sign
//     method.
func (client *MarketplaceAgreementsClient) Sign(ctx context.Context, publisherID string, offerID string, planID string, options *MarketplaceAgreementsClientSignOptions) (MarketplaceAgreementsClientSignResponse, error) {
	req, err := client.signCreateRequest(ctx, publisherID, offerID, planID, options)
	if err != nil {
		return MarketplaceAgreementsClientSignResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MarketplaceAgreementsClientSignResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MarketplaceAgreementsClientSignResponse{}, runtime.NewResponseError(resp)
	}
	return client.signHandleResponse(resp)
}

// signCreateRequest creates the Sign request.
func (client *MarketplaceAgreementsClient) signCreateRequest(ctx context.Context, publisherID string, offerID string, planID string, options *MarketplaceAgreementsClientSignOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.MarketplaceOrdering/agreements/{publisherId}/offers/{offerId}/plans/{planId}/sign"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if publisherID == "" {
		return nil, errors.New("parameter publisherID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publisherId}", url.PathEscape(publisherID))
	if offerID == "" {
		return nil, errors.New("parameter offerID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{offerId}", url.PathEscape(offerID))
	if planID == "" {
		return nil, errors.New("parameter planID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planId}", url.PathEscape(planID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// signHandleResponse handles the Sign response.
func (client *MarketplaceAgreementsClient) signHandleResponse(resp *http.Response) (MarketplaceAgreementsClientSignResponse, error) {
	result := MarketplaceAgreementsClientSignResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AgreementTerms); err != nil {
		return MarketplaceAgreementsClientSignResponse{}, err
	}
	return result, nil
}
