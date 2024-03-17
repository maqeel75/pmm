// Code generated by go-swagger; DO NOT EDIT.

package server_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new server service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for server service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AWSInstanceCheck(params *AWSInstanceCheckParams, opts ...ClientOption) (*AWSInstanceCheckOK, error)

	ChangeSettings(params *ChangeSettingsParams, opts ...ClientOption) (*ChangeSettingsOK, error)

	CheckUpdates(params *CheckUpdatesParams, opts ...ClientOption) (*CheckUpdatesOK, error)

	GetSettings(params *GetSettingsParams, opts ...ClientOption) (*GetSettingsOK, error)

	LeaderHealthCheck(params *LeaderHealthCheckParams, opts ...ClientOption) (*LeaderHealthCheckOK, error)

	Logs(params *LogsParams, writer io.Writer, opts ...ClientOption) (*LogsOK, error)

	Readiness(params *ReadinessParams, opts ...ClientOption) (*ReadinessOK, error)

	StartUpdate(params *StartUpdateParams, opts ...ClientOption) (*StartUpdateOK, error)

	UpdateStatus(params *UpdateStatusParams, opts ...ClientOption) (*UpdateStatusOK, error)

	Version(params *VersionParams, opts ...ClientOption) (*VersionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AWSInstanceCheck AWSs instance check

Checks AWS EC2 instance ID.
*/
func (a *Client) AWSInstanceCheck(params *AWSInstanceCheckParams, opts ...ClientOption) (*AWSInstanceCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAWSInstanceCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AWSInstanceCheck",
		Method:             "GET",
		PathPattern:        "/v1/server/AWSInstance",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AWSInstanceCheckReader{formats: a.formats},
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
	success, ok := result.(*AWSInstanceCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AWSInstanceCheckDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ChangeSettings changes settings

Changes PMM Server settings.
*/
func (a *Client) ChangeSettings(params *ChangeSettingsParams, opts ...ClientOption) (*ChangeSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ChangeSettings",
		Method:             "PUT",
		PathPattern:        "/v1/server/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeSettingsReader{formats: a.formats},
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
	success, ok := result.(*ChangeSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ChangeSettingsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CheckUpdates checks updates

Checks for available PMM Server updates.
*/
func (a *Client) CheckUpdates(params *CheckUpdatesParams, opts ...ClientOption) (*CheckUpdatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckUpdatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CheckUpdates",
		Method:             "POST",
		PathPattern:        "/v1/updates/Check",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CheckUpdatesReader{formats: a.formats},
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
	success, ok := result.(*CheckUpdatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CheckUpdatesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetSettings gets settings

Returns current PMM Server settings.
*/
func (a *Client) GetSettings(params *GetSettingsParams, opts ...ClientOption) (*GetSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSettingsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSettings",
		Method:             "GET",
		PathPattern:        "/v1/server/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSettingsReader{formats: a.formats},
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
	success, ok := result.(*GetSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSettingsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
LeaderHealthCheck checks leadership

Checks if the instance is the leader in a cluster. Returns an error if the instance isn't the leader.
*/
func (a *Client) LeaderHealthCheck(params *LeaderHealthCheckParams, opts ...ClientOption) (*LeaderHealthCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLeaderHealthCheckParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "LeaderHealthCheck",
		Method:             "POST",
		PathPattern:        "/v1/leaderHealthCheck",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LeaderHealthCheckReader{formats: a.formats},
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
	success, ok := result.(*LeaderHealthCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*LeaderHealthCheckDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
Logs logs

Returns the PMM Server logs.
*/
func (a *Client) Logs(params *LogsParams, writer io.Writer, opts ...ClientOption) (*LogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Logs",
		Method:             "GET",
		PathPattern:        "/logs.zip",
		ProducesMediaTypes: []string{"application/zip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &LogsReader{formats: a.formats, writer: writer},
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
	success, ok := result.(*LogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*LogsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
Readiness checks server readiness

Returns an error when Server components being restarted are not ready yet. Use this API for checking the health of Docker containers and for probing Kubernetes readiness.
*/
func (a *Client) Readiness(params *ReadinessParams, opts ...ClientOption) (*ReadinessOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadinessParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Readiness",
		Method:             "GET",
		PathPattern:        "/v1/server/readyz",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReadinessReader{formats: a.formats},
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
	success, ok := result.(*ReadinessOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadinessDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
StartUpdate starts update

Starts PMM Server update.
*/
func (a *Client) StartUpdate(params *StartUpdateParams, opts ...ClientOption) (*StartUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StartUpdate",
		Method:             "POST",
		PathPattern:        "/v1/updates/Start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartUpdateReader{formats: a.formats},
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
	success, ok := result.(*StartUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartUpdateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateStatus updates status

Returns PMM Server update status.
*/
func (a *Client) UpdateStatus(params *UpdateStatusParams, opts ...ClientOption) (*UpdateStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateStatus",
		Method:             "POST",
		PathPattern:        "/v1/updates/Status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateStatusReader{formats: a.formats},
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
	success, ok := result.(*UpdateStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateStatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
Version versions

Returns PMM Server versions.
*/
func (a *Client) Version(params *VersionParams, opts ...ClientOption) (*VersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Version",
		Method:             "GET",
		PathPattern:        "/v1/server/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VersionReader{formats: a.formats},
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
	success, ok := result.(*VersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*VersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}