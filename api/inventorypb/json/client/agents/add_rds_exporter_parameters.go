// Code generated by go-swagger; DO NOT EDIT.

package agents

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

// NewAddRDSExporterParams creates a new AddRDSExporterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddRDSExporterParams() *AddRDSExporterParams {
	return &AddRDSExporterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddRDSExporterParamsWithTimeout creates a new AddRDSExporterParams object
// with the ability to set a timeout on a request.
func NewAddRDSExporterParamsWithTimeout(timeout time.Duration) *AddRDSExporterParams {
	return &AddRDSExporterParams{
		timeout: timeout,
	}
}

// NewAddRDSExporterParamsWithContext creates a new AddRDSExporterParams object
// with the ability to set a context for a request.
func NewAddRDSExporterParamsWithContext(ctx context.Context) *AddRDSExporterParams {
	return &AddRDSExporterParams{
		Context: ctx,
	}
}

// NewAddRDSExporterParamsWithHTTPClient creates a new AddRDSExporterParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddRDSExporterParamsWithHTTPClient(client *http.Client) *AddRDSExporterParams {
	return &AddRDSExporterParams{
		HTTPClient: client,
	}
}

/* AddRDSExporterParams contains all the parameters to send to the API endpoint
   for the add RDS exporter operation.

   Typically these are written to a http.Request.
*/
type AddRDSExporterParams struct {

	// Body.
	Body AddRDSExporterBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add RDS exporter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddRDSExporterParams) WithDefaults() *AddRDSExporterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add RDS exporter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddRDSExporterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add RDS exporter params
func (o *AddRDSExporterParams) WithTimeout(timeout time.Duration) *AddRDSExporterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add RDS exporter params
func (o *AddRDSExporterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add RDS exporter params
func (o *AddRDSExporterParams) WithContext(ctx context.Context) *AddRDSExporterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add RDS exporter params
func (o *AddRDSExporterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add RDS exporter params
func (o *AddRDSExporterParams) WithHTTPClient(client *http.Client) *AddRDSExporterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add RDS exporter params
func (o *AddRDSExporterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add RDS exporter params
func (o *AddRDSExporterParams) WithBody(body AddRDSExporterBody) *AddRDSExporterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add RDS exporter params
func (o *AddRDSExporterParams) SetBody(body AddRDSExporterBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddRDSExporterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
