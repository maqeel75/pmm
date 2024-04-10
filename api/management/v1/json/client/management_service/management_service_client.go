// Code generated by go-swagger; DO NOT EDIT.

package management_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new management service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for management service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddAnnotation(params *AddAnnotationParams, opts ...ClientOption) (*AddAnnotationOK, error)

	AddAzureDatabase(params *AddAzureDatabaseParams, opts ...ClientOption) (*AddAzureDatabaseOK, error)

	AddService(params *AddServiceParams, opts ...ClientOption) (*AddServiceOK, error)

	DiscoverAzureDatabase(params *DiscoverAzureDatabaseParams, opts ...ClientOption) (*DiscoverAzureDatabaseOK, error)

	DiscoverRDS(params *DiscoverRDSParams, opts ...ClientOption) (*DiscoverRDSOK, error)

	GetNode(params *GetNodeParams, opts ...ClientOption) (*GetNodeOK, error)

	ListAgents(params *ListAgentsParams, opts ...ClientOption) (*ListAgentsOK, error)

	ListNodes(params *ListNodesParams, opts ...ClientOption) (*ListNodesOK, error)

	ListServices(params *ListServicesParams, opts ...ClientOption) (*ListServicesOK, error)

	RegisterNode(params *RegisterNodeParams, opts ...ClientOption) (*RegisterNodeOK, error)

	RemoveService(params *RemoveServiceParams, opts ...ClientOption) (*RemoveServiceOK, error)

	UnregisterNode(params *UnregisterNodeParams, opts ...ClientOption) (*UnregisterNodeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddAnnotation adds an annotation

Adds an annotation.
*/
func (a *Client) AddAnnotation(params *AddAnnotationParams, opts ...ClientOption) (*AddAnnotationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddAnnotationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddAnnotation",
		Method:             "POST",
		PathPattern:        "/v1/management/annotations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddAnnotationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddAnnotationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddAnnotationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AddAzureDatabase adds azure database

Adds an Azure Database instance.
*/
func (a *Client) AddAzureDatabase(params *AddAzureDatabaseParams, opts ...ClientOption) (*AddAzureDatabaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddAzureDatabaseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddAzureDatabase",
		Method:             "POST",
		PathPattern:        "/v1/management/services/azure",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddAzureDatabaseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddAzureDatabaseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddAzureDatabaseDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AddService adds a service

Adds a service and starts several agents.
*/
func (a *Client) AddService(params *AddServiceParams, opts ...ClientOption) (*AddServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddService",
		Method:             "POST",
		PathPattern:        "/v1/management/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddServiceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DiscoverAzureDatabase discovers azure database

Discovers Azure Database for MySQL, MariaDB and PostgreSQL Server instances.
*/
func (a *Client) DiscoverAzureDatabase(params *DiscoverAzureDatabaseParams, opts ...ClientOption) (*DiscoverAzureDatabaseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDiscoverAzureDatabaseParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DiscoverAzureDatabase",
		Method:             "POST",
		PathPattern:        "/v1/management/services:discoverAzure",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DiscoverAzureDatabaseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DiscoverAzureDatabaseOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DiscoverAzureDatabaseDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DiscoverRDS discovers RDS

Discovers RDS instances.
*/
func (a *Client) DiscoverRDS(params *DiscoverRDSParams, opts ...ClientOption) (*DiscoverRDSOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDiscoverRDSParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DiscoverRDS",
		Method:             "POST",
		PathPattern:        "/v1/management/services:discoverRDS",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DiscoverRDSReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DiscoverRDSOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DiscoverRDSDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetNode gets node

Gets a single Node by ID.
*/
func (a *Client) GetNode(params *GetNodeParams, opts ...ClientOption) (*GetNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetNode",
		Method:             "GET",
		PathPattern:        "/v1/management/nodes/{node_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetNodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListAgents lists agents

Lists Agents with filter.
*/
func (a *Client) ListAgents(params *ListAgentsParams, opts ...ClientOption) (*ListAgentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAgentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListAgents",
		Method:             "GET",
		PathPattern:        "/v1/management/agents",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListAgentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAgentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListAgentsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListNodes lists nodes

Lists Nodes with filter.
*/
func (a *Client) ListNodes(params *ListNodesParams, opts ...ClientOption) (*ListNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNodesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListNodes",
		Method:             "GET",
		PathPattern:        "/v1/management/nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListNodesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListNodesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListNodesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListServices lists services

Returns a filtered list of Services.
*/
func (a *Client) ListServices(params *ListServicesParams, opts ...ClientOption) (*ListServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListServicesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListServices",
		Method:             "GET",
		PathPattern:        "/v1/management/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListServicesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListServicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListServicesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RegisterNode registers a node

Registers a new Node and a pmm-agent.
*/
func (a *Client) RegisterNode(params *RegisterNodeParams, opts ...ClientOption) (*RegisterNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterNodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RegisterNode",
		Method:             "POST",
		PathPattern:        "/v1/management/nodes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RegisterNodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RegisterNodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RegisterNodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RemoveService removes a service

Removes a Service along with its Agents.
*/
func (a *Client) RemoveService(params *RemoveServiceParams, opts ...ClientOption) (*RemoveServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RemoveService",
		Method:             "DELETE",
		PathPattern:        "/v1/management/services/{service_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RemoveServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemoveServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RemoveServiceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UnregisterNode unregisters a node

Unregisters a Node and pmm-agent
*/
func (a *Client) UnregisterNode(params *UnregisterNodeParams, opts ...ClientOption) (*UnregisterNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUnregisterNodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UnregisterNode",
		Method:             "DELETE",
		PathPattern:        "/v1/management/nodes/{node_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UnregisterNodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UnregisterNodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UnregisterNodeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
