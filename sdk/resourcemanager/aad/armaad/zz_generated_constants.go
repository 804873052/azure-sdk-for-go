//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armaad

const (
	moduleName    = "armaad"
	moduleVersion = "v0.3.0"
)

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateFailed       PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateProvisioning PrivateEndpointConnectionProvisioningState = "Provisioning"
	PrivateEndpointConnectionProvisioningStateSucceeded    PrivateEndpointConnectionProvisioningState = "Succeeded"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateProvisioning,
		PrivateEndpointConnectionProvisioningStateSucceeded,
	}
}

// PrivateEndpointServiceConnectionStatus - The private endpoint connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved     PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusDisconnected PrivateEndpointServiceConnectionStatus = "Disconnected"
	PrivateEndpointServiceConnectionStatusPending      PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected     PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns the possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{
		PrivateEndpointServiceConnectionStatusApproved,
		PrivateEndpointServiceConnectionStatusDisconnected,
		PrivateEndpointServiceConnectionStatusPending,
		PrivateEndpointServiceConnectionStatusRejected,
	}
}
