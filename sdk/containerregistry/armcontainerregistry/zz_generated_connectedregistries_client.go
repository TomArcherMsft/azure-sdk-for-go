// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerregistry

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ConnectedRegistriesClient contains the methods for the ConnectedRegistries group.
// Don't use this type directly, use NewConnectedRegistriesClient() instead.
type ConnectedRegistriesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewConnectedRegistriesClient creates a new instance of ConnectedRegistriesClient with the specified values.
func NewConnectedRegistriesClient(con *armcore.Connection, subscriptionID string) *ConnectedRegistriesClient {
	return &ConnectedRegistriesClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreate - Creates a connected registry for a container registry with the specified parameters.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectedRegistriesClient) BeginCreate(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryCreateParameters ConnectedRegistry, options *ConnectedRegistriesBeginCreateOptions) (ConnectedRegistryPollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, registryName, connectedRegistryName, connectedRegistryCreateParameters, options)
	if err != nil {
		return ConnectedRegistryPollerResponse{}, err
	}
	result := ConnectedRegistryPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ConnectedRegistriesClient.Create", "", resp, client.con.Pipeline(), client.createHandleError)
	if err != nil {
		return ConnectedRegistryPollerResponse{}, err
	}
	poller := &connectedRegistryPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ConnectedRegistryResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreate creates a new ConnectedRegistryPoller from the specified resume token.
// token - The value must come from a previous call to ConnectedRegistryPoller.ResumeToken().
func (client *ConnectedRegistriesClient) ResumeCreate(ctx context.Context, token string) (ConnectedRegistryPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ConnectedRegistriesClient.Create", token, client.con.Pipeline(), client.createHandleError)
	if err != nil {
		return ConnectedRegistryPollerResponse{}, err
	}
	poller := &connectedRegistryPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ConnectedRegistryPollerResponse{}, err
	}
	result := ConnectedRegistryPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ConnectedRegistryResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Create - Creates a connected registry for a container registry with the specified parameters.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectedRegistriesClient) create(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryCreateParameters ConnectedRegistry, options *ConnectedRegistriesBeginCreateOptions) (*azcore.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, registryName, connectedRegistryName, connectedRegistryCreateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ConnectedRegistriesClient) createCreateRequest(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryCreateParameters ConnectedRegistry, options *ConnectedRegistriesBeginCreateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if connectedRegistryName == "" {
		return nil, errors.New("parameter connectedRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedRegistryName}", url.PathEscape(connectedRegistryName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(connectedRegistryCreateParameters)
}

// createHandleError handles the Create error response.
func (client *ConnectedRegistriesClient) createHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDeactivate - Deactivates the connected registry instance.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectedRegistriesClient) BeginDeactivate(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesBeginDeactivateOptions) (HTTPPollerResponse, error) {
	resp, err := client.deactivate(ctx, resourceGroupName, registryName, connectedRegistryName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ConnectedRegistriesClient.Deactivate", "", resp, client.con.Pipeline(), client.deactivateHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDeactivate creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *ConnectedRegistriesClient) ResumeDeactivate(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ConnectedRegistriesClient.Deactivate", token, client.con.Pipeline(), client.deactivateHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Deactivate - Deactivates the connected registry instance.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectedRegistriesClient) deactivate(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesBeginDeactivateOptions) (*azcore.Response, error) {
	req, err := client.deactivateCreateRequest(ctx, resourceGroupName, registryName, connectedRegistryName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.deactivateHandleError(resp)
	}
	return resp, nil
}

// deactivateCreateRequest creates the Deactivate request.
func (client *ConnectedRegistriesClient) deactivateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesBeginDeactivateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}/deactivate"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if connectedRegistryName == "" {
		return nil, errors.New("parameter connectedRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedRegistryName}", url.PathEscape(connectedRegistryName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deactivateHandleError handles the Deactivate error response.
func (client *ConnectedRegistriesClient) deactivateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDelete - Deletes a connected registry from a container registry.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectedRegistriesClient) BeginDelete(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesBeginDeleteOptions) (HTTPPollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, registryName, connectedRegistryName, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ConnectedRegistriesClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDelete creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client *ConnectedRegistriesClient) ResumeDelete(ctx context.Context, token string) (HTTPPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ConnectedRegistriesClient.Delete", token, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Delete - Deletes a connected registry from a container registry.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectedRegistriesClient) deleteOperation(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, registryName, connectedRegistryName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConnectedRegistriesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if connectedRegistryName == "" {
		return nil, errors.New("parameter connectedRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedRegistryName}", url.PathEscape(connectedRegistryName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ConnectedRegistriesClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Gets the properties of the connected registry.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectedRegistriesClient) Get(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesGetOptions) (ConnectedRegistryResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, registryName, connectedRegistryName, options)
	if err != nil {
		return ConnectedRegistryResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ConnectedRegistryResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ConnectedRegistryResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConnectedRegistriesClient) getCreateRequest(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, options *ConnectedRegistriesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if connectedRegistryName == "" {
		return nil, errors.New("parameter connectedRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedRegistryName}", url.PathEscape(connectedRegistryName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConnectedRegistriesClient) getHandleResponse(resp *azcore.Response) (ConnectedRegistryResponse, error) {
	var val *ConnectedRegistry
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ConnectedRegistryResponse{}, err
	}
	return ConnectedRegistryResponse{RawResponse: resp.Response, ConnectedRegistry: val}, nil
}

// getHandleError handles the Get error response.
func (client *ConnectedRegistriesClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - Lists all connected registries for the specified container registry.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectedRegistriesClient) List(resourceGroupName string, registryName string, options *ConnectedRegistriesListOptions) ConnectedRegistryListResultPager {
	return &connectedRegistryListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, registryName, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ConnectedRegistryListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ConnectedRegistryListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client *ConnectedRegistriesClient) listCreateRequest(ctx context.Context, resourceGroupName string, registryName string, options *ConnectedRegistriesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ConnectedRegistriesClient) listHandleResponse(resp *azcore.Response) (ConnectedRegistryListResultResponse, error) {
	var val *ConnectedRegistryListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ConnectedRegistryListResultResponse{}, err
	}
	return ConnectedRegistryListResultResponse{RawResponse: resp.Response, ConnectedRegistryListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *ConnectedRegistriesClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginUpdate - Updates a connected registry with the specified parameters.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectedRegistriesClient) BeginUpdate(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryUpdateParameters ConnectedRegistryUpdateParameters, options *ConnectedRegistriesBeginUpdateOptions) (ConnectedRegistryPollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, registryName, connectedRegistryName, connectedRegistryUpdateParameters, options)
	if err != nil {
		return ConnectedRegistryPollerResponse{}, err
	}
	result := ConnectedRegistryPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ConnectedRegistriesClient.Update", "", resp, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return ConnectedRegistryPollerResponse{}, err
	}
	poller := &connectedRegistryPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ConnectedRegistryResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdate creates a new ConnectedRegistryPoller from the specified resume token.
// token - The value must come from a previous call to ConnectedRegistryPoller.ResumeToken().
func (client *ConnectedRegistriesClient) ResumeUpdate(ctx context.Context, token string) (ConnectedRegistryPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("ConnectedRegistriesClient.Update", token, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return ConnectedRegistryPollerResponse{}, err
	}
	poller := &connectedRegistryPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return ConnectedRegistryPollerResponse{}, err
	}
	result := ConnectedRegistryPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (ConnectedRegistryResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// Update - Updates a connected registry with the specified parameters.
// If the operation fails it returns the *ErrorResponse error type.
func (client *ConnectedRegistriesClient) update(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryUpdateParameters ConnectedRegistryUpdateParameters, options *ConnectedRegistriesBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, registryName, connectedRegistryName, connectedRegistryUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ConnectedRegistriesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, registryName string, connectedRegistryName string, connectedRegistryUpdateParameters ConnectedRegistryUpdateParameters, options *ConnectedRegistriesBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerRegistry/registries/{registryName}/connectedRegistries/{connectedRegistryName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if registryName == "" {
		return nil, errors.New("parameter registryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{registryName}", url.PathEscape(registryName))
	if connectedRegistryName == "" {
		return nil, errors.New("parameter connectedRegistryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectedRegistryName}", url.PathEscape(connectedRegistryName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(connectedRegistryUpdateParameters)
}

// updateHandleError handles the Update error response.
func (client *ConnectedRegistriesClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}