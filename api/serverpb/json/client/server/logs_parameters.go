// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewLogsParams creates a new LogsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLogsParams() *LogsParams {
	return &LogsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLogsParamsWithTimeout creates a new LogsParams object
// with the ability to set a timeout on a request.
func NewLogsParamsWithTimeout(timeout time.Duration) *LogsParams {
	return &LogsParams{
		timeout: timeout,
	}
}

// NewLogsParamsWithContext creates a new LogsParams object
// with the ability to set a context for a request.
func NewLogsParamsWithContext(ctx context.Context) *LogsParams {
	return &LogsParams{
		Context: ctx,
	}
}

// NewLogsParamsWithHTTPClient creates a new LogsParams object
// with the ability to set a custom HTTPClient for a request.
func NewLogsParamsWithHTTPClient(client *http.Client) *LogsParams {
	return &LogsParams{
		HTTPClient: client,
	}
}

/* LogsParams contains all the parameters to send to the API endpoint
   for the logs operation.

   Typically these are written to a http.Request.
*/
type LogsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogsParams) WithDefaults() *LogsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the logs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the logs params
func (o *LogsParams) WithTimeout(timeout time.Duration) *LogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the logs params
func (o *LogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the logs params
func (o *LogsParams) WithContext(ctx context.Context) *LogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the logs params
func (o *LogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the logs params
func (o *LogsParams) WithHTTPClient(client *http.Client) *LogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the logs params
func (o *LogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *LogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
