// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// CollectionPartitionClient contains the methods for the CollectionPartition group.
// Don't use this type directly, use NewCollectionPartitionClient() instead.
type CollectionPartitionClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewCollectionPartitionClient creates a new instance of CollectionPartitionClient with the specified values.
func NewCollectionPartitionClient(con *armcore.Connection, subscriptionID string) *CollectionPartitionClient {
	return &CollectionPartitionClient{con: con, subscriptionID: subscriptionID}
}

// ListMetrics - Retrieves the metrics determined by the given filter for the given collection, split by partition.
// If the operation fails it returns a generic error.
func (client *CollectionPartitionClient) ListMetrics(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string, options *CollectionPartitionListMetricsOptions) (CollectionPartitionListMetricsResponse, error) {
	req, err := client.listMetricsCreateRequest(ctx, resourceGroupName, accountName, databaseRid, collectionRid, filter, options)
	if err != nil {
		return CollectionPartitionListMetricsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return CollectionPartitionListMetricsResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return CollectionPartitionListMetricsResponse{}, client.listMetricsHandleError(resp)
	}
	return client.listMetricsHandleResponse(resp)
}

// listMetricsCreateRequest creates the ListMetrics request.
func (client *CollectionPartitionClient) listMetricsCreateRequest(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, filter string, options *CollectionPartitionListMetricsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/partitions/metrics"
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
	if databaseRid == "" {
		return nil, errors.New("parameter databaseRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseRid}", url.PathEscape(databaseRid))
	if collectionRid == "" {
		return nil, errors.New("parameter collectionRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectionRid}", url.PathEscape(collectionRid))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	reqQP.Set("$filter", filter)
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listMetricsHandleResponse handles the ListMetrics response.
func (client *CollectionPartitionClient) listMetricsHandleResponse(resp *azcore.Response) (CollectionPartitionListMetricsResponse, error) {
	result := CollectionPartitionListMetricsResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.PartitionMetricListResult); err != nil {
		return CollectionPartitionListMetricsResponse{}, err
	}
	return result, nil
}

// listMetricsHandleError handles the ListMetrics error response.
func (client *CollectionPartitionClient) listMetricsHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListUsages - Retrieves the usages (most recent storage data) for the given collection, split by partition.
// If the operation fails it returns a generic error.
func (client *CollectionPartitionClient) ListUsages(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, options *CollectionPartitionListUsagesOptions) (CollectionPartitionListUsagesResponse, error) {
	req, err := client.listUsagesCreateRequest(ctx, resourceGroupName, accountName, databaseRid, collectionRid, options)
	if err != nil {
		return CollectionPartitionListUsagesResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return CollectionPartitionListUsagesResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return CollectionPartitionListUsagesResponse{}, client.listUsagesHandleError(resp)
	}
	return client.listUsagesHandleResponse(resp)
}

// listUsagesCreateRequest creates the ListUsages request.
func (client *CollectionPartitionClient) listUsagesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, databaseRid string, collectionRid string, options *CollectionPartitionListUsagesOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/databases/{databaseRid}/collections/{collectionRid}/partitions/usages"
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
	if databaseRid == "" {
		return nil, errors.New("parameter databaseRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseRid}", url.PathEscape(databaseRid))
	if collectionRid == "" {
		return nil, errors.New("parameter collectionRid cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{collectionRid}", url.PathEscape(collectionRid))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2021-06-15")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listUsagesHandleResponse handles the ListUsages response.
func (client *CollectionPartitionClient) listUsagesHandleResponse(resp *azcore.Response) (CollectionPartitionListUsagesResponse, error) {
	result := CollectionPartitionListUsagesResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.PartitionUsagesResult); err != nil {
		return CollectionPartitionListUsagesResponse{}, err
	}
	return result, nil
}

// listUsagesHandleError handles the ListUsages error response.
func (client *CollectionPartitionClient) listUsagesHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
