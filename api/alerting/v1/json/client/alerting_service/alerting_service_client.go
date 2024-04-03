// Code generated by go-swagger; DO NOT EDIT.

package alerting_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new alerting service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for alerting service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateRule(params *CreateRuleParams, opts ...ClientOption) (*CreateRuleOK, error)

	CreateTemplate(params *CreateTemplateParams, opts ...ClientOption) (*CreateTemplateOK, error)

	DeleteTemplate(params *DeleteTemplateParams, opts ...ClientOption) (*DeleteTemplateOK, error)

	ListTemplates(params *ListTemplatesParams, opts ...ClientOption) (*ListTemplatesOK, error)

	UpdateTemplate(params *UpdateTemplateParams, opts ...ClientOption) (*UpdateTemplateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateRule creates rule creates alerting rule from the given template
*/
func (a *Client) CreateRule(params *CreateRuleParams, opts ...ClientOption) (*CreateRuleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRuleParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateRule",
		Method:             "POST",
		PathPattern:        "/v1/alerting/rules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateRuleReader{formats: a.formats},
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
	success, ok := result.(*CreateRuleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateRuleDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateTemplate creates template creates a new template
*/
func (a *Client) CreateTemplate(params *CreateTemplateParams, opts ...ClientOption) (*CreateTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateTemplateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateTemplate",
		Method:             "POST",
		PathPattern:        "/v1/alerting/templates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateTemplateReader{formats: a.formats},
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
	success, ok := result.(*CreateTemplateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateTemplateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteTemplate deletes template deletes existing previously created via API
*/
func (a *Client) DeleteTemplate(params *DeleteTemplateParams, opts ...ClientOption) (*DeleteTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteTemplateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteTemplate",
		Method:             "DELETE",
		PathPattern:        "/v1/alerting/templates/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteTemplateReader{formats: a.formats},
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
	success, ok := result.(*DeleteTemplateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteTemplateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListTemplates lists templates returns a list of all collected alert rule templates
*/
func (a *Client) ListTemplates(params *ListTemplatesParams, opts ...ClientOption) (*ListTemplatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTemplatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListTemplates",
		Method:             "GET",
		PathPattern:        "/v1/alerting/templates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListTemplatesReader{formats: a.formats},
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
	success, ok := result.(*ListTemplatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListTemplatesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateTemplate updates template updates existing template previously created via API
*/
func (a *Client) UpdateTemplate(params *UpdateTemplateParams, opts ...ClientOption) (*UpdateTemplateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateTemplateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateTemplate",
		Method:             "PUT",
		PathPattern:        "/v1/alerting/templates/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateTemplateReader{formats: a.formats},
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
	success, ok := result.(*UpdateTemplateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateTemplateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
