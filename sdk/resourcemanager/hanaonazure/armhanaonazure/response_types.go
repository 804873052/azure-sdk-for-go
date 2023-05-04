//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhanaonazure

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationList
}

// ProviderInstancesClientCreateResponse contains the response from method ProviderInstancesClient.BeginCreate.
type ProviderInstancesClientCreateResponse struct {
	ProviderInstance
}

// ProviderInstancesClientDeleteResponse contains the response from method ProviderInstancesClient.BeginDelete.
type ProviderInstancesClientDeleteResponse struct {
	// placeholder for future response values
}

// ProviderInstancesClientGetResponse contains the response from method ProviderInstancesClient.Get.
type ProviderInstancesClientGetResponse struct {
	ProviderInstance
}

// ProviderInstancesClientListResponse contains the response from method ProviderInstancesClient.NewListPager.
type ProviderInstancesClientListResponse struct {
	ProviderInstanceListResult
}

// SapMonitorsClientCreateResponse contains the response from method SapMonitorsClient.BeginCreate.
type SapMonitorsClientCreateResponse struct {
	SapMonitor
}

// SapMonitorsClientDeleteResponse contains the response from method SapMonitorsClient.BeginDelete.
type SapMonitorsClientDeleteResponse struct {
	// placeholder for future response values
}

// SapMonitorsClientGetResponse contains the response from method SapMonitorsClient.Get.
type SapMonitorsClientGetResponse struct {
	SapMonitor
}

// SapMonitorsClientListResponse contains the response from method SapMonitorsClient.NewListPager.
type SapMonitorsClientListResponse struct {
	SapMonitorListResult
}

// SapMonitorsClientUpdateResponse contains the response from method SapMonitorsClient.Update.
type SapMonitorsClientUpdateResponse struct {
	SapMonitor
}
