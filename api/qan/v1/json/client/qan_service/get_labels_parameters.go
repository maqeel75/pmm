// Code generated by go-swagger; DO NOT EDIT.

package qan_service

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

// NewGetLabelsParams creates a new GetLabelsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLabelsParams() *GetLabelsParams {
	return &GetLabelsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLabelsParamsWithTimeout creates a new GetLabelsParams object
// with the ability to set a timeout on a request.
func NewGetLabelsParamsWithTimeout(timeout time.Duration) *GetLabelsParams {
	return &GetLabelsParams{
		timeout: timeout,
	}
}

// NewGetLabelsParamsWithContext creates a new GetLabelsParams object
// with the ability to set a context for a request.
func NewGetLabelsParamsWithContext(ctx context.Context) *GetLabelsParams {
	return &GetLabelsParams{
		Context: ctx,
	}
}

// NewGetLabelsParamsWithHTTPClient creates a new GetLabelsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLabelsParamsWithHTTPClient(client *http.Client) *GetLabelsParams {
	return &GetLabelsParams{
		HTTPClient: client,
	}
}

/*
GetLabelsParams contains all the parameters to send to the API endpoint

	for the get labels operation.

	Typically these are written to a http.Request.
*/
type GetLabelsParams struct {
	/* Body.

	     GetLabelsRequest defines filtering of object detail's labels for specific value of
	dimension (ex.: host=hostname1 or queryid=1D410B4BE5060972.
	*/
	Body GetLabelsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get labels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLabelsParams) WithDefaults() *GetLabelsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get labels params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLabelsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get labels params
func (o *GetLabelsParams) WithTimeout(timeout time.Duration) *GetLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get labels params
func (o *GetLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get labels params
func (o *GetLabelsParams) WithContext(ctx context.Context) *GetLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get labels params
func (o *GetLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get labels params
func (o *GetLabelsParams) WithHTTPClient(client *http.Client) *GetLabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get labels params
func (o *GetLabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get labels params
func (o *GetLabelsParams) WithBody(body GetLabelsBody) *GetLabelsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get labels params
func (o *GetLabelsParams) SetBody(body GetLabelsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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