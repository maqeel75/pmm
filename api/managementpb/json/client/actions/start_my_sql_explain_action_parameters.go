// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewStartMySQLExplainActionParams creates a new StartMySQLExplainActionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartMySQLExplainActionParams() *StartMySQLExplainActionParams {
	return &StartMySQLExplainActionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartMySQLExplainActionParamsWithTimeout creates a new StartMySQLExplainActionParams object
// with the ability to set a timeout on a request.
func NewStartMySQLExplainActionParamsWithTimeout(timeout time.Duration) *StartMySQLExplainActionParams {
	return &StartMySQLExplainActionParams{
		timeout: timeout,
	}
}

// NewStartMySQLExplainActionParamsWithContext creates a new StartMySQLExplainActionParams object
// with the ability to set a context for a request.
func NewStartMySQLExplainActionParamsWithContext(ctx context.Context) *StartMySQLExplainActionParams {
	return &StartMySQLExplainActionParams{
		Context: ctx,
	}
}

// NewStartMySQLExplainActionParamsWithHTTPClient creates a new StartMySQLExplainActionParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartMySQLExplainActionParamsWithHTTPClient(client *http.Client) *StartMySQLExplainActionParams {
	return &StartMySQLExplainActionParams{
		HTTPClient: client,
	}
}

/*
StartMySQLExplainActionParams contains all the parameters to send to the API endpoint

	for the start my SQL explain action operation.

	Typically these are written to a http.Request.
*/
type StartMySQLExplainActionParams struct {
	// Body.
	Body StartMySQLExplainActionBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the start my SQL explain action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartMySQLExplainActionParams) WithDefaults() *StartMySQLExplainActionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start my SQL explain action params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartMySQLExplainActionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the start my SQL explain action params
func (o *StartMySQLExplainActionParams) WithTimeout(timeout time.Duration) *StartMySQLExplainActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start my SQL explain action params
func (o *StartMySQLExplainActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start my SQL explain action params
func (o *StartMySQLExplainActionParams) WithContext(ctx context.Context) *StartMySQLExplainActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start my SQL explain action params
func (o *StartMySQLExplainActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start my SQL explain action params
func (o *StartMySQLExplainActionParams) WithHTTPClient(client *http.Client) *StartMySQLExplainActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start my SQL explain action params
func (o *StartMySQLExplainActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the start my SQL explain action params
func (o *StartMySQLExplainActionParams) WithBody(body StartMySQLExplainActionBody) *StartMySQLExplainActionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the start my SQL explain action params
func (o *StartMySQLExplainActionParams) SetBody(body StartMySQLExplainActionBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *StartMySQLExplainActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}