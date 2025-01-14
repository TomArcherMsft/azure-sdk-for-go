// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservice

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"reflect"
)

type AgentPoolsListPager interface {
	azcore.Pager
	// PageResponse returns the current AgentPoolsListResponse.
	PageResponse() AgentPoolsListResponse
}

type agentPoolsListPager struct {
	client    *AgentPoolsClient
	current   AgentPoolsListResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, AgentPoolsListResponse) (*azcore.Request, error)
}

func (p *agentPoolsListPager) Err() error {
	return p.err
}

func (p *agentPoolsListPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AgentPoolListResult.NextLink == nil || len(*p.current.AgentPoolListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *agentPoolsListPager) PageResponse() AgentPoolsListResponse {
	return p.current
}

type MaintenanceConfigurationsListByManagedClusterPager interface {
	azcore.Pager
	// PageResponse returns the current MaintenanceConfigurationsListByManagedClusterResponse.
	PageResponse() MaintenanceConfigurationsListByManagedClusterResponse
}

type maintenanceConfigurationsListByManagedClusterPager struct {
	client    *MaintenanceConfigurationsClient
	current   MaintenanceConfigurationsListByManagedClusterResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, MaintenanceConfigurationsListByManagedClusterResponse) (*azcore.Request, error)
}

func (p *maintenanceConfigurationsListByManagedClusterPager) Err() error {
	return p.err
}

func (p *maintenanceConfigurationsListByManagedClusterPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.MaintenanceConfigurationListResult.NextLink == nil || len(*p.current.MaintenanceConfigurationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listByManagedClusterHandleError(resp)
		return false
	}
	result, err := p.client.listByManagedClusterHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *maintenanceConfigurationsListByManagedClusterPager) PageResponse() MaintenanceConfigurationsListByManagedClusterResponse {
	return p.current
}

type ManagedClustersListByResourceGroupPager interface {
	azcore.Pager
	// PageResponse returns the current ManagedClustersListByResourceGroupResponse.
	PageResponse() ManagedClustersListByResourceGroupResponse
}

type managedClustersListByResourceGroupPager struct {
	client    *ManagedClustersClient
	current   ManagedClustersListByResourceGroupResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, ManagedClustersListByResourceGroupResponse) (*azcore.Request, error)
}

func (p *managedClustersListByResourceGroupPager) Err() error {
	return p.err
}

func (p *managedClustersListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ManagedClusterListResult.NextLink == nil || len(*p.current.ManagedClusterListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listByResourceGroupHandleError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *managedClustersListByResourceGroupPager) PageResponse() ManagedClustersListByResourceGroupResponse {
	return p.current
}

type ManagedClustersListOutboundNetworkDependenciesEndpointsPager interface {
	azcore.Pager
	// PageResponse returns the current ManagedClustersListOutboundNetworkDependenciesEndpointsResponse.
	PageResponse() ManagedClustersListOutboundNetworkDependenciesEndpointsResponse
}

type managedClustersListOutboundNetworkDependenciesEndpointsPager struct {
	client    *ManagedClustersClient
	current   ManagedClustersListOutboundNetworkDependenciesEndpointsResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, ManagedClustersListOutboundNetworkDependenciesEndpointsResponse) (*azcore.Request, error)
}

func (p *managedClustersListOutboundNetworkDependenciesEndpointsPager) Err() error {
	return p.err
}

func (p *managedClustersListOutboundNetworkDependenciesEndpointsPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OutboundEnvironmentEndpointCollection.NextLink == nil || len(*p.current.OutboundEnvironmentEndpointCollection.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listOutboundNetworkDependenciesEndpointsHandleError(resp)
		return false
	}
	result, err := p.client.listOutboundNetworkDependenciesEndpointsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *managedClustersListOutboundNetworkDependenciesEndpointsPager) PageResponse() ManagedClustersListOutboundNetworkDependenciesEndpointsResponse {
	return p.current
}

type ManagedClustersListPager interface {
	azcore.Pager
	// PageResponse returns the current ManagedClustersListResponse.
	PageResponse() ManagedClustersListResponse
}

type managedClustersListPager struct {
	client    *ManagedClustersClient
	current   ManagedClustersListResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, ManagedClustersListResponse) (*azcore.Request, error)
}

func (p *managedClustersListPager) Err() error {
	return p.err
}

func (p *managedClustersListPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ManagedClusterListResult.NextLink == nil || len(*p.current.ManagedClusterListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *managedClustersListPager) PageResponse() ManagedClustersListResponse {
	return p.current
}
