// Code generated by go-swagger; DO NOT EDIT.

package role_service

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

// UpdateRoleReader is a Reader for the UpdateRole structure.
type UpdateRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateRoleOK creates a UpdateRoleOK with default headers values
func NewUpdateRoleOK() *UpdateRoleOK {
	return &UpdateRoleOK{}
}

/*
UpdateRoleOK describes a response with status code 200, with default header values.

A successful response.
*/
type UpdateRoleOK struct {
	Payload interface{}
}

func (o *UpdateRoleOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/Role/Update][%d] updateRoleOk  %+v", 200, o.Payload)
}

func (o *UpdateRoleOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRoleDefault creates a UpdateRoleDefault with default headers values
func NewUpdateRoleDefault(code int) *UpdateRoleDefault {
	return &UpdateRoleDefault{
		_statusCode: code,
	}
}

/*
UpdateRoleDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UpdateRoleDefault struct {
	_statusCode int

	Payload *UpdateRoleDefaultBody
}

// Code gets the status code for the update role default response
func (o *UpdateRoleDefault) Code() int {
	return o._statusCode
}

func (o *UpdateRoleDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/Role/Update][%d] UpdateRole default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateRoleDefault) GetPayload() *UpdateRoleDefaultBody {
	return o.Payload
}

func (o *UpdateRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(UpdateRoleDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateRoleBody update role body
swagger:model UpdateRoleBody
*/
type UpdateRoleBody struct {
	// role id
	RoleID int64 `json:"role_id,omitempty"`

	// title
	Title *string `json:"title,omitempty"`

	// filter
	Filter *string `json:"filter,omitempty"`

	// description
	Description *string `json:"description,omitempty"`
}

// Validate validates this update role body
func (o *UpdateRoleBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update role body based on context it is used
func (o *UpdateRoleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateRoleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateRoleBody) UnmarshalBinary(b []byte) error {
	var res UpdateRoleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateRoleDefaultBody update role default body
swagger:model UpdateRoleDefaultBody
*/
type UpdateRoleDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*UpdateRoleDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this update role default body
func (o *UpdateRoleDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateRoleDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("UpdateRole default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UpdateRole default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update role default body based on the context it is used
func (o *UpdateRoleDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateRoleDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("UpdateRole default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("UpdateRole default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateRoleDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateRoleDefaultBody) UnmarshalBinary(b []byte) error {
	var res UpdateRoleDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateRoleDefaultBodyDetailsItems0 update role default body details items0
swagger:model UpdateRoleDefaultBodyDetailsItems0
*/
type UpdateRoleDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this update role default body details items0
func (o *UpdateRoleDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update role default body details items0 based on context it is used
func (o *UpdateRoleDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateRoleDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateRoleDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateRoleDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
