// Code generated by go-swagger; DO NOT EDIT.

package azure_database

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

// NewAddAzureDatabaseParams creates a new AddAzureDatabaseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddAzureDatabaseParams() *AddAzureDatabaseParams {
	return &AddAzureDatabaseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddAzureDatabaseParamsWithTimeout creates a new AddAzureDatabaseParams object
// with the ability to set a timeout on a request.
func NewAddAzureDatabaseParamsWithTimeout(timeout time.Duration) *AddAzureDatabaseParams {
	return &AddAzureDatabaseParams{
		timeout: timeout,
	}
}

// NewAddAzureDatabaseParamsWithContext creates a new AddAzureDatabaseParams object
// with the ability to set a context for a request.
func NewAddAzureDatabaseParamsWithContext(ctx context.Context) *AddAzureDatabaseParams {
	return &AddAzureDatabaseParams{
		Context: ctx,
	}
}

// NewAddAzureDatabaseParamsWithHTTPClient creates a new AddAzureDatabaseParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddAzureDatabaseParamsWithHTTPClient(client *http.Client) *AddAzureDatabaseParams {
	return &AddAzureDatabaseParams{
		HTTPClient: client,
	}
}

/* AddAzureDatabaseParams contains all the parameters to send to the API endpoint
   for the add azure database operation.

   Typically these are written to a http.Request.
*/
type AddAzureDatabaseParams struct {

	// Body.
	Body AddAzureDatabaseBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add azure database params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAzureDatabaseParams) WithDefaults() *AddAzureDatabaseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add azure database params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddAzureDatabaseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add azure database params
func (o *AddAzureDatabaseParams) WithTimeout(timeout time.Duration) *AddAzureDatabaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add azure database params
func (o *AddAzureDatabaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add azure database params
func (o *AddAzureDatabaseParams) WithContext(ctx context.Context) *AddAzureDatabaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add azure database params
func (o *AddAzureDatabaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add azure database params
func (o *AddAzureDatabaseParams) WithHTTPClient(client *http.Client) *AddAzureDatabaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add azure database params
func (o *AddAzureDatabaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add azure database params
func (o *AddAzureDatabaseParams) WithBody(body AddAzureDatabaseBody) *AddAzureDatabaseParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add azure database params
func (o *AddAzureDatabaseParams) SetBody(body AddAzureDatabaseBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddAzureDatabaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
