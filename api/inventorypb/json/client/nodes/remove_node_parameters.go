// Code generated by go-swagger; DO NOT EDIT.

package nodes

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

// NewRemoveNodeParams creates a new RemoveNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveNodeParams() *RemoveNodeParams {
	return &RemoveNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveNodeParamsWithTimeout creates a new RemoveNodeParams object
// with the ability to set a timeout on a request.
func NewRemoveNodeParamsWithTimeout(timeout time.Duration) *RemoveNodeParams {
	return &RemoveNodeParams{
		timeout: timeout,
	}
}

// NewRemoveNodeParamsWithContext creates a new RemoveNodeParams object
// with the ability to set a context for a request.
func NewRemoveNodeParamsWithContext(ctx context.Context) *RemoveNodeParams {
	return &RemoveNodeParams{
		Context: ctx,
	}
}

// NewRemoveNodeParamsWithHTTPClient creates a new RemoveNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveNodeParamsWithHTTPClient(client *http.Client) *RemoveNodeParams {
	return &RemoveNodeParams{
		HTTPClient: client,
	}
}

/* RemoveNodeParams contains all the parameters to send to the API endpoint
   for the remove node operation.

   Typically these are written to a http.Request.
*/
type RemoveNodeParams struct {

	// Body.
	Body RemoveNodeBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveNodeParams) WithDefaults() *RemoveNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveNodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the remove node params
func (o *RemoveNodeParams) WithTimeout(timeout time.Duration) *RemoveNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove node params
func (o *RemoveNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove node params
func (o *RemoveNodeParams) WithContext(ctx context.Context) *RemoveNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove node params
func (o *RemoveNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove node params
func (o *RemoveNodeParams) WithHTTPClient(client *http.Client) *RemoveNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove node params
func (o *RemoveNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the remove node params
func (o *RemoveNodeParams) WithBody(body RemoveNodeBody) *RemoveNodeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the remove node params
func (o *RemoveNodeParams) SetBody(body RemoveNodeBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
