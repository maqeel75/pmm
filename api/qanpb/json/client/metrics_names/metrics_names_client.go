// Code generated by go-swagger; DO NOT EDIT.

package metrics_names

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new metrics names API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for metrics names API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetMetricsNames(params *GetMetricsNamesParams, opts ...ClientOption) (*GetMetricsNamesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetMetricsNames gets metrics names gets map of metrics names
*/
func (a *Client) GetMetricsNames(params *GetMetricsNamesParams, opts ...ClientOption) (*GetMetricsNamesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMetricsNamesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetMetricsNames",
		Method:             "POST",
		PathPattern:        "/v0/qan/GetMetricsNames",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetMetricsNamesReader{formats: a.formats},
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
	success, ok := result.(*GetMetricsNamesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetMetricsNamesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
