//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurearcdata

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// DataControllersDeleteDataControllerPollerResponse contains the response from method DataControllers.DeleteDataController.
type DataControllersDeleteDataControllerPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *DataControllersDeleteDataControllerPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l DataControllersDeleteDataControllerPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (DataControllersDeleteDataControllerResponse, error) {
	respType := DataControllersDeleteDataControllerResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a DataControllersDeleteDataControllerPollerResponse from the provided client and resume token.
func (l *DataControllersDeleteDataControllerPollerResponse) Resume(ctx context.Context, client *DataControllersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("DataControllersClient.DeleteDataController", token, client.pl, client.deleteDataControllerHandleError)
	if err != nil {
		return err
	}
	poller := &DataControllersDeleteDataControllerPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// DataControllersDeleteDataControllerResponse contains the response from method DataControllers.DeleteDataController.
type DataControllersDeleteDataControllerResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DataControllersGetDataControllerResponse contains the response from method DataControllers.GetDataController.
type DataControllersGetDataControllerResponse struct {
	DataControllersGetDataControllerResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DataControllersGetDataControllerResult contains the result from method DataControllers.GetDataController.
type DataControllersGetDataControllerResult struct {
	DataControllerResource
}

// DataControllersListInGroupResponse contains the response from method DataControllers.ListInGroup.
type DataControllersListInGroupResponse struct {
	DataControllersListInGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DataControllersListInGroupResult contains the result from method DataControllers.ListInGroup.
type DataControllersListInGroupResult struct {
	PageOfDataControllerResource
}

// DataControllersListInSubscriptionResponse contains the response from method DataControllers.ListInSubscription.
type DataControllersListInSubscriptionResponse struct {
	DataControllersListInSubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DataControllersListInSubscriptionResult contains the result from method DataControllers.ListInSubscription.
type DataControllersListInSubscriptionResult struct {
	PageOfDataControllerResource
}

// DataControllersPatchDataControllerResponse contains the response from method DataControllers.PatchDataController.
type DataControllersPatchDataControllerResponse struct {
	DataControllersPatchDataControllerResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DataControllersPatchDataControllerResult contains the result from method DataControllers.PatchDataController.
type DataControllersPatchDataControllerResult struct {
	DataControllerResource
}

// DataControllersPutDataControllerPollerResponse contains the response from method DataControllers.PutDataController.
type DataControllersPutDataControllerPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *DataControllersPutDataControllerPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l DataControllersPutDataControllerPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (DataControllersPutDataControllerResponse, error) {
	respType := DataControllersPutDataControllerResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.DataControllerResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a DataControllersPutDataControllerPollerResponse from the provided client and resume token.
func (l *DataControllersPutDataControllerPollerResponse) Resume(ctx context.Context, client *DataControllersClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("DataControllersClient.PutDataController", token, client.pl, client.putDataControllerHandleError)
	if err != nil {
		return err
	}
	poller := &DataControllersPutDataControllerPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// DataControllersPutDataControllerResponse contains the response from method DataControllers.PutDataController.
type DataControllersPutDataControllerResponse struct {
	DataControllersPutDataControllerResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// DataControllersPutDataControllerResult contains the result from method DataControllers.PutDataController.
type DataControllersPutDataControllerResult struct {
	DataControllerResource
}

// OperationsListResponse contains the response from method Operations.List.
type OperationsListResponse struct {
	OperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsListResult contains the result from method Operations.List.
type OperationsListResult struct {
	OperationListResult
}

// SQLManagedInstancesCreatePollerResponse contains the response from method SQLManagedInstances.Create.
type SQLManagedInstancesCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SQLManagedInstancesCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SQLManagedInstancesCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SQLManagedInstancesCreateResponse, error) {
	respType := SQLManagedInstancesCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.SQLManagedInstance)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SQLManagedInstancesCreatePollerResponse from the provided client and resume token.
func (l *SQLManagedInstancesCreatePollerResponse) Resume(ctx context.Context, client *SQLManagedInstancesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SQLManagedInstancesClient.Create", token, client.pl, client.createHandleError)
	if err != nil {
		return err
	}
	poller := &SQLManagedInstancesCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SQLManagedInstancesCreateResponse contains the response from method SQLManagedInstances.Create.
type SQLManagedInstancesCreateResponse struct {
	SQLManagedInstancesCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLManagedInstancesCreateResult contains the result from method SQLManagedInstances.Create.
type SQLManagedInstancesCreateResult struct {
	SQLManagedInstance
}

// SQLManagedInstancesDeletePollerResponse contains the response from method SQLManagedInstances.Delete.
type SQLManagedInstancesDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SQLManagedInstancesDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SQLManagedInstancesDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SQLManagedInstancesDeleteResponse, error) {
	respType := SQLManagedInstancesDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SQLManagedInstancesDeletePollerResponse from the provided client and resume token.
func (l *SQLManagedInstancesDeletePollerResponse) Resume(ctx context.Context, client *SQLManagedInstancesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SQLManagedInstancesClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &SQLManagedInstancesDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SQLManagedInstancesDeleteResponse contains the response from method SQLManagedInstances.Delete.
type SQLManagedInstancesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLManagedInstancesGetResponse contains the response from method SQLManagedInstances.Get.
type SQLManagedInstancesGetResponse struct {
	SQLManagedInstancesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLManagedInstancesGetResult contains the result from method SQLManagedInstances.Get.
type SQLManagedInstancesGetResult struct {
	SQLManagedInstance
}

// SQLManagedInstancesListByResourceGroupResponse contains the response from method SQLManagedInstances.ListByResourceGroup.
type SQLManagedInstancesListByResourceGroupResponse struct {
	SQLManagedInstancesListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLManagedInstancesListByResourceGroupResult contains the result from method SQLManagedInstances.ListByResourceGroup.
type SQLManagedInstancesListByResourceGroupResult struct {
	SQLManagedInstanceListResult
}

// SQLManagedInstancesListResponse contains the response from method SQLManagedInstances.List.
type SQLManagedInstancesListResponse struct {
	SQLManagedInstancesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLManagedInstancesListResult contains the result from method SQLManagedInstances.List.
type SQLManagedInstancesListResult struct {
	SQLManagedInstanceListResult
}

// SQLManagedInstancesUpdateResponse contains the response from method SQLManagedInstances.Update.
type SQLManagedInstancesUpdateResponse struct {
	SQLManagedInstancesUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLManagedInstancesUpdateResult contains the result from method SQLManagedInstances.Update.
type SQLManagedInstancesUpdateResult struct {
	SQLManagedInstance
}

// SQLServerInstancesCreatePollerResponse contains the response from method SQLServerInstances.Create.
type SQLServerInstancesCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SQLServerInstancesCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SQLServerInstancesCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SQLServerInstancesCreateResponse, error) {
	respType := SQLServerInstancesCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.SQLServerInstance)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SQLServerInstancesCreatePollerResponse from the provided client and resume token.
func (l *SQLServerInstancesCreatePollerResponse) Resume(ctx context.Context, client *SQLServerInstancesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SQLServerInstancesClient.Create", token, client.pl, client.createHandleError)
	if err != nil {
		return err
	}
	poller := &SQLServerInstancesCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SQLServerInstancesCreateResponse contains the response from method SQLServerInstances.Create.
type SQLServerInstancesCreateResponse struct {
	SQLServerInstancesCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLServerInstancesCreateResult contains the result from method SQLServerInstances.Create.
type SQLServerInstancesCreateResult struct {
	SQLServerInstance
}

// SQLServerInstancesDeletePollerResponse contains the response from method SQLServerInstances.Delete.
type SQLServerInstancesDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SQLServerInstancesDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SQLServerInstancesDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SQLServerInstancesDeleteResponse, error) {
	respType := SQLServerInstancesDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SQLServerInstancesDeletePollerResponse from the provided client and resume token.
func (l *SQLServerInstancesDeletePollerResponse) Resume(ctx context.Context, client *SQLServerInstancesClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SQLServerInstancesClient.Delete", token, client.pl, client.deleteHandleError)
	if err != nil {
		return err
	}
	poller := &SQLServerInstancesDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SQLServerInstancesDeleteResponse contains the response from method SQLServerInstances.Delete.
type SQLServerInstancesDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLServerInstancesGetResponse contains the response from method SQLServerInstances.Get.
type SQLServerInstancesGetResponse struct {
	SQLServerInstancesGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLServerInstancesGetResult contains the result from method SQLServerInstances.Get.
type SQLServerInstancesGetResult struct {
	SQLServerInstance
}

// SQLServerInstancesListByResourceGroupResponse contains the response from method SQLServerInstances.ListByResourceGroup.
type SQLServerInstancesListByResourceGroupResponse struct {
	SQLServerInstancesListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLServerInstancesListByResourceGroupResult contains the result from method SQLServerInstances.ListByResourceGroup.
type SQLServerInstancesListByResourceGroupResult struct {
	SQLServerInstanceListResult
}

// SQLServerInstancesListResponse contains the response from method SQLServerInstances.List.
type SQLServerInstancesListResponse struct {
	SQLServerInstancesListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLServerInstancesListResult contains the result from method SQLServerInstances.List.
type SQLServerInstancesListResult struct {
	SQLServerInstanceListResult
}

// SQLServerInstancesUpdateResponse contains the response from method SQLServerInstances.Update.
type SQLServerInstancesUpdateResponse struct {
	SQLServerInstancesUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SQLServerInstancesUpdateResult contains the result from method SQLServerInstances.Update.
type SQLServerInstancesUpdateResult struct {
	SQLServerInstance
}
