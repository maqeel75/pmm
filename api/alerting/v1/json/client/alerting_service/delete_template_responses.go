// Code generated by go-swagger; DO NOT EDIT.

package alerting_service

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

// DeleteTemplateReader is a Reader for the DeleteTemplate structure.
type DeleteTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTemplateOK creates a DeleteTemplateOK with default headers values
func NewDeleteTemplateOK() *DeleteTemplateOK {
	return &DeleteTemplateOK{}
}

/*
DeleteTemplateOK describes a response with status code 200, with default header values.

A successful response.
*/
type DeleteTemplateOK struct {
	Payload interface{}
}

func (o *DeleteTemplateOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/alerting/templates/{name}][%d] deleteTemplateOk  %+v", 200, o.Payload)
}

func (o *DeleteTemplateOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTemplateDefault creates a DeleteTemplateDefault with default headers values
func NewDeleteTemplateDefault(code int) *DeleteTemplateDefault {
	return &DeleteTemplateDefault{
		_statusCode: code,
	}
}

/*
DeleteTemplateDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type DeleteTemplateDefault struct {
	_statusCode int

	Payload *DeleteTemplateDefaultBody
}

// Code gets the status code for the delete template default response
func (o *DeleteTemplateDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTemplateDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/alerting/templates/{name}][%d] DeleteTemplate default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTemplateDefault) GetPayload() *DeleteTemplateDefaultBody {
	return o.Payload
}

func (o *DeleteTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(DeleteTemplateDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DeleteTemplateDefaultBody delete template default body
swagger:model DeleteTemplateDefaultBody
*/
type DeleteTemplateDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DeleteTemplateDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this delete template default body
func (o *DeleteTemplateDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteTemplateDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("DeleteTemplate default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DeleteTemplate default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this delete template default body based on the context it is used
func (o *DeleteTemplateDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeleteTemplateDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DeleteTemplate default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DeleteTemplate default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeleteTemplateDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteTemplateDefaultBody) UnmarshalBinary(b []byte) error {
	var res DeleteTemplateDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
DeleteTemplateDefaultBodyDetailsItems0 delete template default body details items0
swagger:model DeleteTemplateDefaultBodyDetailsItems0
*/
type DeleteTemplateDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this delete template default body details items0
func (o *DeleteTemplateDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delete template default body details items0 based on context it is used
func (o *DeleteTemplateDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DeleteTemplateDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeleteTemplateDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res DeleteTemplateDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
