//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armiotsecurity

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

// SensorsClient contains the methods for the Sensors group.
// Don't use this type directly, use NewSensorsClient() instead.
type SensorsClient struct {
	internal *arm.Client
}

// NewSensorsClient creates a new instance of SensorsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSensorsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*SensorsClient, error) {
	cl, err := arm.NewClient(moduleName+".SensorsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SensorsClient{
		internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update IoT sensor
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01-preview
//   - scope - Scope of the query (IoT Hub, /providers/Microsoft.Devices/iotHubs/myHub)
//   - sensorName - Name of the IoT sensor
//   - sensorModel - The IoT sensor model
//   - options - SensorsClientCreateOrUpdateOptions contains the optional parameters for the SensorsClient.CreateOrUpdate method.
func (client *SensorsClient) CreateOrUpdate(ctx context.Context, scope string, sensorName string, sensorModel SensorModel, options *SensorsClientCreateOrUpdateOptions) (SensorsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, scope, sensorName, sensorModel, options)
	if err != nil {
		return SensorsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SensorsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SensorsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SensorsClient) createOrUpdateCreateRequest(ctx context.Context, scope string, sensorName string, sensorModel SensorModel, options *SensorsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.IoTSecurity/sensors/{sensorName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if sensorName == "" {
		return nil, errors.New("parameter sensorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sensorName}", url.PathEscape(sensorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, sensorModel)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SensorsClient) createOrUpdateHandleResponse(resp *http.Response) (SensorsClientCreateOrUpdateResponse, error) {
	result := SensorsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SensorModel); err != nil {
		return SensorsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete IoT sensor
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01-preview
//   - scope - Scope of the query (IoT Hub, /providers/Microsoft.Devices/iotHubs/myHub)
//   - sensorName - Name of the IoT sensor
//   - options - SensorsClientDeleteOptions contains the optional parameters for the SensorsClient.Delete method.
func (client *SensorsClient) Delete(ctx context.Context, scope string, sensorName string, options *SensorsClientDeleteOptions) (SensorsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, scope, sensorName, options)
	if err != nil {
		return SensorsClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SensorsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return SensorsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return SensorsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SensorsClient) deleteCreateRequest(ctx context.Context, scope string, sensorName string, options *SensorsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.IoTSecurity/sensors/{sensorName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if sensorName == "" {
		return nil, errors.New("parameter sensorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sensorName}", url.PathEscape(sensorName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DownloadActivation - Download sensor activation file
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01-preview
//   - scope - Scope of the query (IoT Hub, /providers/Microsoft.Devices/iotHubs/myHub)
//   - sensorName - Name of the IoT sensor
//   - options - SensorsClientDownloadActivationOptions contains the optional parameters for the SensorsClient.DownloadActivation
//     method.
func (client *SensorsClient) DownloadActivation(ctx context.Context, scope string, sensorName string, options *SensorsClientDownloadActivationOptions) (SensorsClientDownloadActivationResponse, error) {
	req, err := client.downloadActivationCreateRequest(ctx, scope, sensorName, options)
	if err != nil {
		return SensorsClientDownloadActivationResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SensorsClientDownloadActivationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SensorsClientDownloadActivationResponse{}, runtime.NewResponseError(resp)
	}
	return SensorsClientDownloadActivationResponse{Body: resp.Body}, nil
}

// downloadActivationCreateRequest creates the DownloadActivation request.
func (client *SensorsClient) downloadActivationCreateRequest(ctx context.Context, scope string, sensorName string, options *SensorsClientDownloadActivationOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.IoTSecurity/sensors/{sensorName}/downloadActivation"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if sensorName == "" {
		return nil, errors.New("parameter sensorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sensorName}", url.PathEscape(sensorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	runtime.SkipBodyDownload(req)
	req.Raw().Header["Accept"] = []string{"application/zip"}
	return req, nil
}

// DownloadResetPassword - Download file for reset password of the sensor
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01-preview
//   - scope - Scope of the query (IoT Hub, /providers/Microsoft.Devices/iotHubs/myHub)
//   - sensorName - Name of the IoT sensor
//   - body - The reset password input.
//   - options - SensorsClientDownloadResetPasswordOptions contains the optional parameters for the SensorsClient.DownloadResetPassword
//     method.
func (client *SensorsClient) DownloadResetPassword(ctx context.Context, scope string, sensorName string, body ResetPasswordInput, options *SensorsClientDownloadResetPasswordOptions) (SensorsClientDownloadResetPasswordResponse, error) {
	req, err := client.downloadResetPasswordCreateRequest(ctx, scope, sensorName, body, options)
	if err != nil {
		return SensorsClientDownloadResetPasswordResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SensorsClientDownloadResetPasswordResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SensorsClientDownloadResetPasswordResponse{}, runtime.NewResponseError(resp)
	}
	return SensorsClientDownloadResetPasswordResponse{Body: resp.Body}, nil
}

// downloadResetPasswordCreateRequest creates the DownloadResetPassword request.
func (client *SensorsClient) downloadResetPasswordCreateRequest(ctx context.Context, scope string, sensorName string, body ResetPasswordInput, options *SensorsClientDownloadResetPasswordOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.IoTSecurity/sensors/{sensorName}/downloadResetPassword"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if sensorName == "" {
		return nil, errors.New("parameter sensorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sensorName}", url.PathEscape(sensorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	runtime.SkipBodyDownload(req)
	req.Raw().Header["Accept"] = []string{"application/zip"}
	return req, runtime.MarshalAsJSON(req, body)
}

// Get - Get IoT sensor
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01-preview
//   - scope - Scope of the query (IoT Hub, /providers/Microsoft.Devices/iotHubs/myHub)
//   - sensorName - Name of the IoT sensor
//   - options - SensorsClientGetOptions contains the optional parameters for the SensorsClient.Get method.
func (client *SensorsClient) Get(ctx context.Context, scope string, sensorName string, options *SensorsClientGetOptions) (SensorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, sensorName, options)
	if err != nil {
		return SensorsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SensorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SensorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SensorsClient) getCreateRequest(ctx context.Context, scope string, sensorName string, options *SensorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.IoTSecurity/sensors/{sensorName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if sensorName == "" {
		return nil, errors.New("parameter sensorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sensorName}", url.PathEscape(sensorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SensorsClient) getHandleResponse(resp *http.Response) (SensorsClientGetResponse, error) {
	result := SensorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SensorModel); err != nil {
		return SensorsClientGetResponse{}, err
	}
	return result, nil
}

// List - List IoT sensors
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01-preview
//   - scope - Scope of the query (IoT Hub, /providers/Microsoft.Devices/iotHubs/myHub)
//   - options - SensorsClientListOptions contains the optional parameters for the SensorsClient.List method.
func (client *SensorsClient) List(ctx context.Context, scope string, options *SensorsClientListOptions) (SensorsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, scope, options)
	if err != nil {
		return SensorsClientListResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SensorsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SensorsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *SensorsClient) listCreateRequest(ctx context.Context, scope string, options *SensorsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.IoTSecurity/sensors"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SensorsClient) listHandleResponse(resp *http.Response) (SensorsClientListResponse, error) {
	result := SensorsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SensorsList); err != nil {
		return SensorsClientListResponse{}, err
	}
	return result, nil
}

// TriggerTiPackageUpdate - Trigger threat intelligence package update
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01-preview
//   - scope - Scope of the query (IoT Hub, /providers/Microsoft.Devices/iotHubs/myHub)
//   - sensorName - Name of the IoT sensor
//   - options - SensorsClientTriggerTiPackageUpdateOptions contains the optional parameters for the SensorsClient.TriggerTiPackageUpdate
//     method.
func (client *SensorsClient) TriggerTiPackageUpdate(ctx context.Context, scope string, sensorName string, options *SensorsClientTriggerTiPackageUpdateOptions) (SensorsClientTriggerTiPackageUpdateResponse, error) {
	req, err := client.triggerTiPackageUpdateCreateRequest(ctx, scope, sensorName, options)
	if err != nil {
		return SensorsClientTriggerTiPackageUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SensorsClientTriggerTiPackageUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SensorsClientTriggerTiPackageUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return SensorsClientTriggerTiPackageUpdateResponse{}, nil
}

// triggerTiPackageUpdateCreateRequest creates the TriggerTiPackageUpdate request.
func (client *SensorsClient) triggerTiPackageUpdateCreateRequest(ctx context.Context, scope string, sensorName string, options *SensorsClientTriggerTiPackageUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.IoTSecurity/sensors/{sensorName}/triggerTiPackageUpdate"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if sensorName == "" {
		return nil, errors.New("parameter sensorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sensorName}", url.PathEscape(sensorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
