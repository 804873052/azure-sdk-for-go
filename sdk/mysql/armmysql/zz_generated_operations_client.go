//go:build go1.16
// +build go1.16

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.4.3, generator: @autorest/go@4.0.0-preview.27)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysql

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// OperationsClient contains the methods for the Operations group.
// Don't use this type directly, use NewOperationsClient() instead.
type OperationsClient struct {
	con *connection
}

// NewOperationsClient creates a new instance of OperationsClient with the specified values.
func NewOperationsClient(con *connection) *OperationsClient {
	return &OperationsClient{con: con}
}

// List - Lists all of the available REST API operations.
// If the operation fails it returns the *CloudError error type.
func (client *OperationsClient) List(ctx context.Context, options *OperationsListOptions) (OperationsListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return OperationsListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return OperationsListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OperationsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *OperationsClient) listCreateRequest(ctx context.Context, options *OperationsListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.DBforMySQL/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OperationsClient) listHandleResponse(resp *http.Response) (OperationsListResponse, error) {
	result := OperationsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationListResult); err != nil {
		return OperationsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *OperationsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}
