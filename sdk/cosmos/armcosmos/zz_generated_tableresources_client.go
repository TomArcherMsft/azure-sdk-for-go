// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

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

// TableResourcesClient contains the methods for the TableResources group.
// Don't use this type directly, use NewTableResourcesClient() instead.
type TableResourcesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewTableResourcesClient creates a new instance of TableResourcesClient with the specified values.
func NewTableResourcesClient(con *armcore.Connection, subscriptionID string) *TableResourcesClient {
	return &TableResourcesClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateUpdateTable - Create or update an Azure Cosmos DB Table
// If the operation fails it returns a generic error.
func (client *TableResourcesClient) BeginCreateUpdateTable(ctx context.Context, resourceGroupName string, accountName string, tableName string, createUpdateTableParameters TableCreateUpdateParameters, options *TableResourcesBeginCreateUpdateTableOptions) (TableResourcesCreateUpdateTablePollerResponse, error) {
	resp, err := client.createUpdateTable(ctx, resourceGroupName, accountName, tableName, createUpdateTableParameters, options)
	if err != nil {
		return TableResourcesCreateUpdateTablePollerResponse{}, err
	}
	result := TableResourcesCreateUpdateTablePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("TableResourcesClient.CreateUpdateTable", "", resp, client.con.Pipeline(), client.createUpdateTableHandleError)
	if err != nil {
		return TableResourcesCreateUpdateTablePollerResponse{}, err
	}
	poller := &tableResourcesCreateUpdateTablePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (TableResourcesCreateUpdateTableResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateUpdateTable creates a new TableResourcesCreateUpdateTablePoller from the specified resume token.
// token - The value must come from a previous call to TableResourcesCreateUpdateTablePoller.ResumeToken().
func (client *TableResourcesClient) ResumeCreateUpdateTable(ctx context.Context, token string) (TableResourcesCreateUpdateTablePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("TableResourcesClient.CreateUpdateTable", token, client.con.Pipeline(), client.createUpdateTableHandleError)
	if err != nil {
		return TableResourcesCreateUpdateTablePollerResponse{}, err
	}
	poller := &tableResourcesCreateUpdateTablePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return TableResourcesCreateUpdateTablePollerResponse{}, err
	}
	result := TableResourcesCreateUpdateTablePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (TableResourcesCreateUpdateTableResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// CreateUpdateTable - Create or update an Azure Cosmos DB Table
// If the operation fails it returns a generic error.
func (client *TableResourcesClient) createUpdateTable(ctx context.Context, resourceGroupName string, accountName string, tableName string, createUpdateTableParameters TableCreateUpdateParameters, options *TableResourcesBeginCreateUpdateTableOptions) (*azcore.Response, error) {
	req, err := client.createUpdateTableCreateRequest(ctx, resourceGroupName, accountName, tableName, createUpdateTableParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createUpdateTableHandleError(resp)
	}
	return resp, nil
}

// createUpdateTableCreateRequest creates the CreateUpdateTable request.
func (client *TableResourcesClient) createUpdateTableCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, createUpdateTableParameters TableCreateUpdateParameters, options *TableResourcesBeginCreateUpdateTableOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(createUpdateTableParameters)
}

// createUpdateTableHandleError handles the CreateUpdateTable error response.
func (client *TableResourcesClient) createUpdateTableHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginDeleteTable - Deletes an existing Azure Cosmos DB Table.
// If the operation fails it returns a generic error.
func (client *TableResourcesClient) BeginDeleteTable(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesBeginDeleteTableOptions) (TableResourcesDeleteTablePollerResponse, error) {
	resp, err := client.deleteTable(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return TableResourcesDeleteTablePollerResponse{}, err
	}
	result := TableResourcesDeleteTablePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("TableResourcesClient.DeleteTable", "", resp, client.con.Pipeline(), client.deleteTableHandleError)
	if err != nil {
		return TableResourcesDeleteTablePollerResponse{}, err
	}
	poller := &tableResourcesDeleteTablePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (TableResourcesDeleteTableResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeDeleteTable creates a new TableResourcesDeleteTablePoller from the specified resume token.
// token - The value must come from a previous call to TableResourcesDeleteTablePoller.ResumeToken().
func (client *TableResourcesClient) ResumeDeleteTable(ctx context.Context, token string) (TableResourcesDeleteTablePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("TableResourcesClient.DeleteTable", token, client.con.Pipeline(), client.deleteTableHandleError)
	if err != nil {
		return TableResourcesDeleteTablePollerResponse{}, err
	}
	poller := &tableResourcesDeleteTablePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return TableResourcesDeleteTablePollerResponse{}, err
	}
	result := TableResourcesDeleteTablePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (TableResourcesDeleteTableResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// DeleteTable - Deletes an existing Azure Cosmos DB Table.
// If the operation fails it returns a generic error.
func (client *TableResourcesClient) deleteTable(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesBeginDeleteTableOptions) (*azcore.Response, error) {
	req, err := client.deleteTableCreateRequest(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteTableHandleError(resp)
	}
	return resp, nil
}

// deleteTableCreateRequest creates the DeleteTable request.
func (client *TableResourcesClient) deleteTableCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesBeginDeleteTableOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteTableHandleError handles the DeleteTable error response.
func (client *TableResourcesClient) deleteTableHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// GetTable - Gets the Tables under an existing Azure Cosmos DB database account with the provided name.
// If the operation fails it returns a generic error.
func (client *TableResourcesClient) GetTable(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesGetTableOptions) (TableResourcesGetTableResponse, error) {
	req, err := client.getTableCreateRequest(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return TableResourcesGetTableResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TableResourcesGetTableResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TableResourcesGetTableResponse{}, client.getTableHandleError(resp)
	}
	return client.getTableHandleResponse(resp)
}

// getTableCreateRequest creates the GetTable request.
func (client *TableResourcesClient) getTableCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesGetTableOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getTableHandleResponse handles the GetTable response.
func (client *TableResourcesClient) getTableHandleResponse(resp *azcore.Response) (TableResourcesGetTableResponse, error) {
	result := TableResourcesGetTableResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.TableGetResults); err != nil {
		return TableResourcesGetTableResponse{}, err
	}
	return result, nil
}

// getTableHandleError handles the GetTable error response.
func (client *TableResourcesClient) getTableHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// GetTableThroughput - Gets the RUs per second of the Table under an existing Azure Cosmos DB database account with the provided name.
// If the operation fails it returns a generic error.
func (client *TableResourcesClient) GetTableThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesGetTableThroughputOptions) (TableResourcesGetTableThroughputResponse, error) {
	req, err := client.getTableThroughputCreateRequest(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return TableResourcesGetTableThroughputResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TableResourcesGetTableThroughputResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TableResourcesGetTableThroughputResponse{}, client.getTableThroughputHandleError(resp)
	}
	return client.getTableThroughputHandleResponse(resp)
}

// getTableThroughputCreateRequest creates the GetTableThroughput request.
func (client *TableResourcesClient) getTableThroughputCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesGetTableThroughputOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/throughputSettings/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getTableThroughputHandleResponse handles the GetTableThroughput response.
func (client *TableResourcesClient) getTableThroughputHandleResponse(resp *azcore.Response) (TableResourcesGetTableThroughputResponse, error) {
	result := TableResourcesGetTableThroughputResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ThroughputSettingsGetResults); err != nil {
		return TableResourcesGetTableThroughputResponse{}, err
	}
	return result, nil
}

// getTableThroughputHandleError handles the GetTableThroughput error response.
func (client *TableResourcesClient) getTableThroughputHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListTables - Lists the Tables under an existing Azure Cosmos DB database account.
// If the operation fails it returns a generic error.
func (client *TableResourcesClient) ListTables(ctx context.Context, resourceGroupName string, accountName string, options *TableResourcesListTablesOptions) (TableResourcesListTablesResponse, error) {
	req, err := client.listTablesCreateRequest(ctx, resourceGroupName, accountName, options)
	if err != nil {
		return TableResourcesListTablesResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return TableResourcesListTablesResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TableResourcesListTablesResponse{}, client.listTablesHandleError(resp)
	}
	return client.listTablesHandleResponse(resp)
}

// listTablesCreateRequest creates the ListTables request.
func (client *TableResourcesClient) listTablesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *TableResourcesListTablesOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listTablesHandleResponse handles the ListTables response.
func (client *TableResourcesClient) listTablesHandleResponse(resp *azcore.Response) (TableResourcesListTablesResponse, error) {
	result := TableResourcesListTablesResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.TableListResult); err != nil {
		return TableResourcesListTablesResponse{}, err
	}
	return result, nil
}

// listTablesHandleError handles the ListTables error response.
func (client *TableResourcesClient) listTablesHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginMigrateTableToAutoscale - Migrate an Azure Cosmos DB Table from manual throughput to autoscale
// If the operation fails it returns the *CloudError error type.
func (client *TableResourcesClient) BeginMigrateTableToAutoscale(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesBeginMigrateTableToAutoscaleOptions) (TableResourcesMigrateTableToAutoscalePollerResponse, error) {
	resp, err := client.migrateTableToAutoscale(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return TableResourcesMigrateTableToAutoscalePollerResponse{}, err
	}
	result := TableResourcesMigrateTableToAutoscalePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("TableResourcesClient.MigrateTableToAutoscale", "", resp, client.con.Pipeline(), client.migrateTableToAutoscaleHandleError)
	if err != nil {
		return TableResourcesMigrateTableToAutoscalePollerResponse{}, err
	}
	poller := &tableResourcesMigrateTableToAutoscalePoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (TableResourcesMigrateTableToAutoscaleResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeMigrateTableToAutoscale creates a new TableResourcesMigrateTableToAutoscalePoller from the specified resume token.
// token - The value must come from a previous call to TableResourcesMigrateTableToAutoscalePoller.ResumeToken().
func (client *TableResourcesClient) ResumeMigrateTableToAutoscale(ctx context.Context, token string) (TableResourcesMigrateTableToAutoscalePollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("TableResourcesClient.MigrateTableToAutoscale", token, client.con.Pipeline(), client.migrateTableToAutoscaleHandleError)
	if err != nil {
		return TableResourcesMigrateTableToAutoscalePollerResponse{}, err
	}
	poller := &tableResourcesMigrateTableToAutoscalePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return TableResourcesMigrateTableToAutoscalePollerResponse{}, err
	}
	result := TableResourcesMigrateTableToAutoscalePollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (TableResourcesMigrateTableToAutoscaleResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// MigrateTableToAutoscale - Migrate an Azure Cosmos DB Table from manual throughput to autoscale
// If the operation fails it returns the *CloudError error type.
func (client *TableResourcesClient) migrateTableToAutoscale(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesBeginMigrateTableToAutoscaleOptions) (*azcore.Response, error) {
	req, err := client.migrateTableToAutoscaleCreateRequest(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.migrateTableToAutoscaleHandleError(resp)
	}
	return resp, nil
}

// migrateTableToAutoscaleCreateRequest creates the MigrateTableToAutoscale request.
func (client *TableResourcesClient) migrateTableToAutoscaleCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesBeginMigrateTableToAutoscaleOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/throughputSettings/default/migrateToAutoscale"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// migrateTableToAutoscaleHandleError handles the MigrateTableToAutoscale error response.
func (client *TableResourcesClient) migrateTableToAutoscaleHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginMigrateTableToManualThroughput - Migrate an Azure Cosmos DB Table from autoscale to manual throughput
// If the operation fails it returns the *CloudError error type.
func (client *TableResourcesClient) BeginMigrateTableToManualThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesBeginMigrateTableToManualThroughputOptions) (TableResourcesMigrateTableToManualThroughputPollerResponse, error) {
	resp, err := client.migrateTableToManualThroughput(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return TableResourcesMigrateTableToManualThroughputPollerResponse{}, err
	}
	result := TableResourcesMigrateTableToManualThroughputPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("TableResourcesClient.MigrateTableToManualThroughput", "", resp, client.con.Pipeline(), client.migrateTableToManualThroughputHandleError)
	if err != nil {
		return TableResourcesMigrateTableToManualThroughputPollerResponse{}, err
	}
	poller := &tableResourcesMigrateTableToManualThroughputPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (TableResourcesMigrateTableToManualThroughputResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeMigrateTableToManualThroughput creates a new TableResourcesMigrateTableToManualThroughputPoller from the specified resume token.
// token - The value must come from a previous call to TableResourcesMigrateTableToManualThroughputPoller.ResumeToken().
func (client *TableResourcesClient) ResumeMigrateTableToManualThroughput(ctx context.Context, token string) (TableResourcesMigrateTableToManualThroughputPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("TableResourcesClient.MigrateTableToManualThroughput", token, client.con.Pipeline(), client.migrateTableToManualThroughputHandleError)
	if err != nil {
		return TableResourcesMigrateTableToManualThroughputPollerResponse{}, err
	}
	poller := &tableResourcesMigrateTableToManualThroughputPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return TableResourcesMigrateTableToManualThroughputPollerResponse{}, err
	}
	result := TableResourcesMigrateTableToManualThroughputPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (TableResourcesMigrateTableToManualThroughputResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// MigrateTableToManualThroughput - Migrate an Azure Cosmos DB Table from autoscale to manual throughput
// If the operation fails it returns the *CloudError error type.
func (client *TableResourcesClient) migrateTableToManualThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesBeginMigrateTableToManualThroughputOptions) (*azcore.Response, error) {
	req, err := client.migrateTableToManualThroughputCreateRequest(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.migrateTableToManualThroughputHandleError(resp)
	}
	return resp, nil
}

// migrateTableToManualThroughputCreateRequest creates the MigrateTableToManualThroughput request.
func (client *TableResourcesClient) migrateTableToManualThroughputCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesBeginMigrateTableToManualThroughputOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/throughputSettings/default/migrateToManualThroughput"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// migrateTableToManualThroughputHandleError handles the MigrateTableToManualThroughput error response.
func (client *TableResourcesClient) migrateTableToManualThroughputHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginUpdateTableThroughput - Update RUs per second of an Azure Cosmos DB Table
// If the operation fails it returns a generic error.
func (client *TableResourcesClient) BeginUpdateTableThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string, updateThroughputParameters ThroughputSettingsUpdateParameters, options *TableResourcesBeginUpdateTableThroughputOptions) (TableResourcesUpdateTableThroughputPollerResponse, error) {
	resp, err := client.updateTableThroughput(ctx, resourceGroupName, accountName, tableName, updateThroughputParameters, options)
	if err != nil {
		return TableResourcesUpdateTableThroughputPollerResponse{}, err
	}
	result := TableResourcesUpdateTableThroughputPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("TableResourcesClient.UpdateTableThroughput", "", resp, client.con.Pipeline(), client.updateTableThroughputHandleError)
	if err != nil {
		return TableResourcesUpdateTableThroughputPollerResponse{}, err
	}
	poller := &tableResourcesUpdateTableThroughputPoller{
		pt: pt,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (TableResourcesUpdateTableThroughputResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeUpdateTableThroughput creates a new TableResourcesUpdateTableThroughputPoller from the specified resume token.
// token - The value must come from a previous call to TableResourcesUpdateTableThroughputPoller.ResumeToken().
func (client *TableResourcesClient) ResumeUpdateTableThroughput(ctx context.Context, token string) (TableResourcesUpdateTableThroughputPollerResponse, error) {
	pt, err := armcore.NewLROPollerFromResumeToken("TableResourcesClient.UpdateTableThroughput", token, client.con.Pipeline(), client.updateTableThroughputHandleError)
	if err != nil {
		return TableResourcesUpdateTableThroughputPollerResponse{}, err
	}
	poller := &tableResourcesUpdateTableThroughputPoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return TableResourcesUpdateTableThroughputPollerResponse{}, err
	}
	result := TableResourcesUpdateTableThroughputPollerResponse{
		RawResponse: resp,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (TableResourcesUpdateTableThroughputResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// UpdateTableThroughput - Update RUs per second of an Azure Cosmos DB Table
// If the operation fails it returns a generic error.
func (client *TableResourcesClient) updateTableThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string, updateThroughputParameters ThroughputSettingsUpdateParameters, options *TableResourcesBeginUpdateTableThroughputOptions) (*azcore.Response, error) {
	req, err := client.updateTableThroughputCreateRequest(ctx, resourceGroupName, accountName, tableName, updateThroughputParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.updateTableThroughputHandleError(resp)
	}
	return resp, nil
}

// updateTableThroughputCreateRequest creates the UpdateTableThroughput request.
func (client *TableResourcesClient) updateTableThroughputCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, updateThroughputParameters ThroughputSettingsUpdateParameters, options *TableResourcesBeginUpdateTableThroughputOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/throughputSettings/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(updateThroughputParameters)
}

// updateTableThroughputHandleError handles the UpdateTableThroughput error response.
func (client *TableResourcesClient) updateTableThroughputHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
