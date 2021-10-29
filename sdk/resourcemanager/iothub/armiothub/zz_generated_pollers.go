//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiothub

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// IotHubManualFailoverPoller provides polling facilities until the operation reaches a terminal state.
type IotHubManualFailoverPoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *IotHubManualFailoverPoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *IotHubManualFailoverPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final IotHubManualFailoverResponse will be returned.
func (p *IotHubManualFailoverPoller) FinalResponse(ctx context.Context) (IotHubManualFailoverResponse, error) {
	respType := IotHubManualFailoverResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return IotHubManualFailoverResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *IotHubManualFailoverPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// IotHubResourceCreateOrUpdatePoller provides polling facilities until the operation reaches a terminal state.
type IotHubResourceCreateOrUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *IotHubResourceCreateOrUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *IotHubResourceCreateOrUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final IotHubResourceCreateOrUpdateResponse will be returned.
func (p *IotHubResourceCreateOrUpdatePoller) FinalResponse(ctx context.Context) (IotHubResourceCreateOrUpdateResponse, error) {
	respType := IotHubResourceCreateOrUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.IotHubDescription)
	if err != nil {
		return IotHubResourceCreateOrUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *IotHubResourceCreateOrUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// IotHubResourceDeletePoller provides polling facilities until the operation reaches a terminal state.
type IotHubResourceDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *IotHubResourceDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *IotHubResourceDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final IotHubResourceDeleteResponse will be returned.
func (p *IotHubResourceDeletePoller) FinalResponse(ctx context.Context) (IotHubResourceDeleteResponse, error) {
	respType := IotHubResourceDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.IotHubDescription)
	if err != nil {
		return IotHubResourceDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *IotHubResourceDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// IotHubResourceUpdatePoller provides polling facilities until the operation reaches a terminal state.
type IotHubResourceUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *IotHubResourceUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *IotHubResourceUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final IotHubResourceUpdateResponse will be returned.
func (p *IotHubResourceUpdatePoller) FinalResponse(ctx context.Context) (IotHubResourceUpdateResponse, error) {
	respType := IotHubResourceUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.IotHubDescription)
	if err != nil {
		return IotHubResourceUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *IotHubResourceUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// PrivateEndpointConnectionsDeletePoller provides polling facilities until the operation reaches a terminal state.
type PrivateEndpointConnectionsDeletePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *PrivateEndpointConnectionsDeletePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *PrivateEndpointConnectionsDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final PrivateEndpointConnectionsDeleteResponse will be returned.
func (p *PrivateEndpointConnectionsDeletePoller) FinalResponse(ctx context.Context) (PrivateEndpointConnectionsDeleteResponse, error) {
	respType := PrivateEndpointConnectionsDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.PrivateEndpointConnection)
	if err != nil {
		return PrivateEndpointConnectionsDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *PrivateEndpointConnectionsDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

// PrivateEndpointConnectionsUpdatePoller provides polling facilities until the operation reaches a terminal state.
type PrivateEndpointConnectionsUpdatePoller struct {
	pt *azcore.Poller
}

// Done returns true if the LRO has reached a terminal state.
func (p *PrivateEndpointConnectionsUpdatePoller) Done() bool {
	return p.pt.Done()
}

// Poll fetches the latest state of the LRO.  It returns an HTTP response or error.
// If the LRO has completed successfully, the poller's state is updated and the HTTP
// response is returned.
// If the LRO has completed with failure or was cancelled, the poller's state is
// updated and the error is returned.
// If the LRO has not reached a terminal state, the poller's state is updated and
// the latest HTTP response is returned.
// If Poll fails, the poller's state is unmodified and the error is returned.
// Calling Poll on an LRO that has reached a terminal state will return the final
// HTTP response or error.
func (p *PrivateEndpointConnectionsUpdatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

// FinalResponse performs a final GET to the service and returns the final response
// for the polling operation. If there is an error performing the final GET then an error is returned.
// If the final GET succeeded then the final PrivateEndpointConnectionsUpdateResponse will be returned.
func (p *PrivateEndpointConnectionsUpdatePoller) FinalResponse(ctx context.Context) (PrivateEndpointConnectionsUpdateResponse, error) {
	respType := PrivateEndpointConnectionsUpdateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.PrivateEndpointConnection)
	if err != nil {
		return PrivateEndpointConnectionsUpdateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// ResumeToken returns a value representing the poller that can be used to resume
// the LRO at a later time. ResumeTokens are unique per service operation.
func (p *PrivateEndpointConnectionsUpdatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}
