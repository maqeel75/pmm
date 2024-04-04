// Code generated by go-swagger; DO NOT EDIT.

package management_service

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

// UnregisterNodeReader is a Reader for the UnregisterNode structure.
type UnregisterNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnregisterNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnregisterNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnregisterNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnregisterNodeOK creates a UnregisterNodeOK with default headers values
func NewUnregisterNodeOK() *UnregisterNodeOK {
	return &UnregisterNodeOK{}
}

/*
UnregisterNodeOK describes a response with status code 200, with default header values.

A successful response.
*/
type UnregisterNodeOK struct {
	Payload *UnregisterNodeOKBody
}

func (o *UnregisterNodeOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Node/Unregister][%d] unregisterNodeOk  %+v", 200, o.Payload)
}

func (o *UnregisterNodeOK) GetPayload() *UnregisterNodeOKBody {
	return o.Payload
}

func (o *UnregisterNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(UnregisterNodeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUnregisterNodeDefault creates a UnregisterNodeDefault with default headers values
func NewUnregisterNodeDefault(code int) *UnregisterNodeDefault {
	return &UnregisterNodeDefault{
		_statusCode: code,
	}
}

/*
UnregisterNodeDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UnregisterNodeDefault struct {
	_statusCode int

	Payload *UnregisterNodeDefaultBody
}

// Code gets the status code for the unregister node default response
func (o *UnregisterNodeDefault) Code() int {
	return o._statusCode
}

func (o *UnregisterNodeDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Node/Unregister][%d] UnregisterNode default  %+v", o._statusCode, o.Payload)
}

func (o *UnregisterNodeDefault) GetPayload() *UnregisterNodeDefaultBody {
	return o.Payload
}

func (o *UnregisterNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(UnregisterNodeDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UnregisterNodeBody unregister node body
swagger:model UnregisterNodeBody
*/
type UnregisterNodeBody struct {
	// Node_id to be unregistered.
	NodeID string `json:"node_id,omitempty"`

	// Force delete node, related service account, even if it has more service tokens attached.
	Force bool `json:"force,omitempty"`
}

// Validate validates this unregister node body
func (o *UnregisterNodeBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this unregister node body based on context it is used
func (o *UnregisterNodeBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UnregisterNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnregisterNodeBody) UnmarshalBinary(b []byte) error {
	var res UnregisterNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UnregisterNodeDefaultBody unregister node default body
swagger:model UnregisterNodeDefaultBody
*/
type UnregisterNodeDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*UnregisterNodeDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this unregister node default body
func (o *UnregisterNodeDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnregisterNodeDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("UnregisterNode default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UnregisterNode default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this unregister node default body based on the context it is used
func (o *UnregisterNodeDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UnregisterNodeDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UnregisterNode default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UnregisterNode default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UnregisterNodeDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnregisterNodeDefaultBody) UnmarshalBinary(b []byte) error {
	var res UnregisterNodeDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UnregisterNodeDefaultBodyDetailsItems0 unregister node default body details items0
swagger:model UnregisterNodeDefaultBodyDetailsItems0
*/
type UnregisterNodeDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this unregister node default body details items0
func (o *UnregisterNodeDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this unregister node default body details items0 based on context it is used
func (o *UnregisterNodeDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UnregisterNodeDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnregisterNodeDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res UnregisterNodeDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UnregisterNodeOKBody unregister node OK body
swagger:model UnregisterNodeOKBody
*/
type UnregisterNodeOKBody struct {
	// Warning message if there are more service tokens attached to service account.
	Warning string `json:"warning,omitempty"`
}

// Validate validates this unregister node OK body
func (o *UnregisterNodeOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this unregister node OK body based on context it is used
func (o *UnregisterNodeOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UnregisterNodeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UnregisterNodeOKBody) UnmarshalBinary(b []byte) error {
	var res UnregisterNodeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
