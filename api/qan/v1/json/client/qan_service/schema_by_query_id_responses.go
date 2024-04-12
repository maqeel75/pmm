// Code generated by go-swagger; DO NOT EDIT.

package qan_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SchemaByQueryIDReader is a Reader for the SchemaByQueryID structure.
type SchemaByQueryIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchemaByQueryIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSchemaByQueryIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSchemaByQueryIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSchemaByQueryIDOK creates a SchemaByQueryIDOK with default headers values
func NewSchemaByQueryIDOK() *SchemaByQueryIDOK {
	return &SchemaByQueryIDOK{}
}

/*
SchemaByQueryIDOK describes a response with status code 200, with default header values.

A successful response.
*/
type SchemaByQueryIDOK struct {
	Payload *SchemaByQueryIDOKBody
}

func (o *SchemaByQueryIDOK) Error() string {
	return fmt.Sprintf("[POST /v1/qan/query:getSchema][%d] schemaByQueryIdOk  %+v", 200, o.Payload)
}

func (o *SchemaByQueryIDOK) GetPayload() *SchemaByQueryIDOKBody {
	return o.Payload
}

func (o *SchemaByQueryIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(SchemaByQueryIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaByQueryIDDefault creates a SchemaByQueryIDDefault with default headers values
func NewSchemaByQueryIDDefault(code int) *SchemaByQueryIDDefault {
	return &SchemaByQueryIDDefault{
		_statusCode: code,
	}
}

/*
SchemaByQueryIDDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SchemaByQueryIDDefault struct {
	_statusCode int

	Payload *SchemaByQueryIDDefaultBody
}

// Code gets the status code for the schema by query ID default response
func (o *SchemaByQueryIDDefault) Code() int {
	return o._statusCode
}

func (o *SchemaByQueryIDDefault) Error() string {
	return fmt.Sprintf("[POST /v1/qan/query:getSchema][%d] SchemaByQueryID default  %+v", o._statusCode, o.Payload)
}

func (o *SchemaByQueryIDDefault) GetPayload() *SchemaByQueryIDDefaultBody {
	return o.Payload
}

func (o *SchemaByQueryIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(SchemaByQueryIDDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SchemaByQueryIDBody SchemaByQueryIDRequest returns schema for given query ID and service ID.
swagger:model SchemaByQueryIDBody
*/
type SchemaByQueryIDBody struct {
	// service id
	ServiceID string `json:"service_id,omitempty"`

	// query id
	QueryID string `json:"query_id,omitempty"`
}

// Validate validates this schema by query ID body
func (o *SchemaByQueryIDBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this schema by query ID body based on context it is used
func (o *SchemaByQueryIDBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SchemaByQueryIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SchemaByQueryIDBody) UnmarshalBinary(b []byte) error {
	var res SchemaByQueryIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SchemaByQueryIDDefaultBody schema by query ID default body
swagger:model SchemaByQueryIDDefaultBody
*/
type SchemaByQueryIDDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*SchemaByQueryIDDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this schema by query ID default body
func (o *SchemaByQueryIDDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SchemaByQueryIDDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SchemaByQueryID default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SchemaByQueryID default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this schema by query ID default body based on the context it is used
func (o *SchemaByQueryIDDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SchemaByQueryIDDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SchemaByQueryID default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SchemaByQueryID default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SchemaByQueryIDDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SchemaByQueryIDDefaultBody) UnmarshalBinary(b []byte) error {
	var res SchemaByQueryIDDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SchemaByQueryIDDefaultBodyDetailsItems0 schema by query ID default body details items0
swagger:model SchemaByQueryIDDefaultBodyDetailsItems0
*/
type SchemaByQueryIDDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this schema by query ID default body details items0
func (o *SchemaByQueryIDDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this schema by query ID default body details items0 based on context it is used
func (o *SchemaByQueryIDDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SchemaByQueryIDDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SchemaByQueryIDDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res SchemaByQueryIDDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SchemaByQueryIDOKBody SchemaByQueryIDResponse is schema for given query ID and service ID.
swagger:model SchemaByQueryIDOKBody
*/
type SchemaByQueryIDOKBody struct {
	// schema
	Schema string `json:"schema,omitempty"`
}

// Validate validates this schema by query ID OK body
func (o *SchemaByQueryIDOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this schema by query ID OK body based on context it is used
func (o *SchemaByQueryIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SchemaByQueryIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SchemaByQueryIDOKBody) UnmarshalBinary(b []byte) error {
	var res SchemaByQueryIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
