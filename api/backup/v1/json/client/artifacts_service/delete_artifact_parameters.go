// Code generated by go-swagger; DO NOT EDIT.

package artifacts_service

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

// NewDeleteArtifactParams creates a new DeleteArtifactParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteArtifactParams() *DeleteArtifactParams {
	return &DeleteArtifactParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteArtifactParamsWithTimeout creates a new DeleteArtifactParams object
// with the ability to set a timeout on a request.
func NewDeleteArtifactParamsWithTimeout(timeout time.Duration) *DeleteArtifactParams {
	return &DeleteArtifactParams{
		timeout: timeout,
	}
}

// NewDeleteArtifactParamsWithContext creates a new DeleteArtifactParams object
// with the ability to set a context for a request.
func NewDeleteArtifactParamsWithContext(ctx context.Context) *DeleteArtifactParams {
	return &DeleteArtifactParams{
		Context: ctx,
	}
}

// NewDeleteArtifactParamsWithHTTPClient creates a new DeleteArtifactParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteArtifactParamsWithHTTPClient(client *http.Client) *DeleteArtifactParams {
	return &DeleteArtifactParams{
		HTTPClient: client,
	}
}

/*
DeleteArtifactParams contains all the parameters to send to the API endpoint

	for the delete artifact operation.

	Typically these are written to a http.Request.
*/
type DeleteArtifactParams struct {
	/* ArtifactID.

	   Machine-readable artifact ID.
	*/
	ArtifactID string

	// Body.
	Body DeleteArtifactBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete artifact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteArtifactParams) WithDefaults() *DeleteArtifactParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete artifact params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteArtifactParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete artifact params
func (o *DeleteArtifactParams) WithTimeout(timeout time.Duration) *DeleteArtifactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete artifact params
func (o *DeleteArtifactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete artifact params
func (o *DeleteArtifactParams) WithContext(ctx context.Context) *DeleteArtifactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete artifact params
func (o *DeleteArtifactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete artifact params
func (o *DeleteArtifactParams) WithHTTPClient(client *http.Client) *DeleteArtifactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete artifact params
func (o *DeleteArtifactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArtifactID adds the artifactID to the delete artifact params
func (o *DeleteArtifactParams) WithArtifactID(artifactID string) *DeleteArtifactParams {
	o.SetArtifactID(artifactID)
	return o
}

// SetArtifactID adds the artifactId to the delete artifact params
func (o *DeleteArtifactParams) SetArtifactID(artifactID string) {
	o.ArtifactID = artifactID
}

// WithBody adds the body to the delete artifact params
func (o *DeleteArtifactParams) WithBody(body DeleteArtifactBody) *DeleteArtifactParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete artifact params
func (o *DeleteArtifactParams) SetBody(body DeleteArtifactBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteArtifactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param artifact_id
	if err := r.SetPathParam("artifact_id", o.ArtifactID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
