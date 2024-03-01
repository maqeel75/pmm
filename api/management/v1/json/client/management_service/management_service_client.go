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

	AddExternal(params *AddExternalParams, opts ...ClientOption) (*AddExternalOK, error)

	AddHAProxy(params *AddHAProxyParams, opts ...ClientOption) (*AddHAProxyOK, error)

	AddMongoDB(params *AddMongoDBParams, opts ...ClientOption) (*AddMongoDBOK, error)

	AddMySQL(params *AddMySQLParams, opts ...ClientOption) (*AddMySQLOK, error)

	AddPostgreSQL(params *AddPostgreSQLParams, opts ...ClientOption) (*AddPostgreSQLOK, error)

	AddProxySQL(params *AddProxySQLParams, opts ...ClientOption) (*AddProxySQLOK, error)

	AddRDS(params *AddRDSParams, opts ...ClientOption) (*AddRDSOK, error)

	DiscoverRDS(params *DiscoverRDSParams, opts ...ClientOption) (*DiscoverRDSOK, error)

	RegisterNode(params *RegisterNodeParams, opts ...ClientOption) (*RegisterNodeOK, error)

	RemoveService(params *RemoveServiceParams, opts ...ClientOption) (*RemoveServiceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AddAnnotation adds annotation

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
		PathPattern:        "/v1/management/Annotations/Add",
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
AddExternal adds external service

Adds external service and adds external exporter. It automatically adds a service to inventory, which is running on provided "node_id", then adds an "external exporter" agent to inventory, which is running on provided "runs_on_node_id".
*/
func (a *Client) AddExternal(params *AddExternalParams, opts ...ClientOption) (*AddExternalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddExternalParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddExternal",
		Method:             "POST",
		PathPattern:        "/v1/management/External/Add",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddExternalReader{formats: a.formats},
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
	success, ok := result.(*AddExternalOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddExternalDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AddHAProxy adds HA proxy

Adds HAProxy service and external exporter. It automatically adds a service to inventory, which is running on the provided "node_id", then adds an "external exporter" agent to the inventory.
*/
func (a *Client) AddHAProxy(params *AddHAProxyParams, opts ...ClientOption) (*AddHAProxyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddHAProxyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddHAProxy",
		Method:             "POST",
		PathPattern:        "/v1/management/HAProxy/Add",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddHAProxyReader{formats: a.formats},
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
	success, ok := result.(*AddHAProxyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddHAProxyDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AddMongoDB adds mongo DB

Adds MongoDB Service and starts several Agents. It automatically adds a service to inventory, which is running on the provided "node_id", then adds "mongodb_exporter", and "qan_mongodb_profiler" agents with the provided "pmm_agent_id" and other parameters.
*/
func (a *Client) AddMongoDB(params *AddMongoDBParams, opts ...ClientOption) (*AddMongoDBOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddMongoDBParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddMongoDB",
		Method:             "POST",
		PathPattern:        "/v1/management/MongoDB/Add",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddMongoDBReader{formats: a.formats},
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
	success, ok := result.(*AddMongoDBOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddMongoDBDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AddMySQL adds my SQL

Adds MySQL Service and starts several Agents. It automatically adds a service to inventory, which is running on the provided "node_id", then adds "mysqld_exporter", and "qan_mysql_perfschema" agents with the provided "pmm_agent_id" and other parameters.
*/
func (a *Client) AddMySQL(params *AddMySQLParams, opts ...ClientOption) (*AddMySQLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddMySQLParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddMySQL",
		Method:             "POST",
		PathPattern:        "/v1/management/MySQL/Add",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddMySQLReader{formats: a.formats},
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
	success, ok := result.(*AddMySQLOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddMySQLDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AddPostgreSQL adds postgre SQL

Adds PostgreSQL Service and starts postgres exporter. It automatically adds a service to inventory, which is running on provided "node_id", then adds "postgres_exporter" with provided "pmm_agent_id" and other parameters.
*/
func (a *Client) AddPostgreSQL(params *AddPostgreSQLParams, opts ...ClientOption) (*AddPostgreSQLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddPostgreSQLParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddPostgreSQL",
		Method:             "POST",
		PathPattern:        "/v1/management/PostgreSQL/Add",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddPostgreSQLReader{formats: a.formats},
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
	success, ok := result.(*AddPostgreSQLOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddPostgreSQLDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AddProxySQL adds proxy SQL

Adds ProxySQL Service and starts several Agents. It automatically adds a service to inventory, which is running on provided "node_id", then adds "proxysql_exporter" with provided "pmm_agent_id" and other parameters.
*/
func (a *Client) AddProxySQL(params *AddProxySQLParams, opts ...ClientOption) (*AddProxySQLOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddProxySQLParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddProxySQL",
		Method:             "POST",
		PathPattern:        "/v1/management/ProxySQL/Add",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddProxySQLReader{formats: a.formats},
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
	success, ok := result.(*AddProxySQLOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddProxySQLDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AddRDS adds RDS

Adds RDS instance.
*/
func (a *Client) AddRDS(params *AddRDSParams, opts ...ClientOption) (*AddRDSOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddRDSParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AddRDS",
		Method:             "POST",
		PathPattern:        "/v1/management/RDS/Add",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddRDSReader{formats: a.formats},
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
	success, ok := result.(*AddRDSOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AddRDSDefault)
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
		PathPattern:        "/v1/management/RDS/Discover",
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
RegisterNode registers node

Registers a new Node and pmm-agent.
*/
func (a *Client) RegisterNode(params *RegisterNodeParams, opts ...ClientOption) (*RegisterNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterNodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RegisterNode",
		Method:             "POST",
		PathPattern:        "/v1/management/Node/Register",
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
RemoveService removes service

Removes a Service along with Agents.
*/
func (a *Client) RemoveService(params *RemoveServiceParams, opts ...ClientOption) (*RemoveServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemoveServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RemoveService",
		Method:             "POST",
		PathPattern:        "/v1/management/Service/Remove",
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

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
