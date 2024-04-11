// Code generated by go-swagger; DO NOT EDIT.

package platform_service

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

// ConnectReader is a Reader for the Connect structure.
type ConnectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConnectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConnectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConnectOK creates a ConnectOK with default headers values
func NewConnectOK() *ConnectOK {
	return &ConnectOK{}
}

/*
ConnectOK describes a response with status code 200, with default header values.

A successful response.
*/
type ConnectOK struct {
	Payload interface{}
}

func (o *ConnectOK) Error() string {
	return fmt.Sprintf("[POST /v1/platform:connect][%d] connectOk  %+v", 200, o.Payload)
}

func (o *ConnectOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ConnectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectDefault creates a ConnectDefault with default headers values
func NewConnectDefault(code int) *ConnectDefault {
	return &ConnectDefault{
		_statusCode: code,
	}
}

/*
ConnectDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ConnectDefault struct {
	_statusCode int

	Payload *ConnectDefaultBody
}

// Code gets the status code for the connect default response
func (o *ConnectDefault) Code() int {
	return o._statusCode
}

func (o *ConnectDefault) Error() string {
	return fmt.Sprintf("[POST /v1/platform:connect][%d] Connect default  %+v", o._statusCode, o.Payload)
}

func (o *ConnectDefault) GetPayload() *ConnectDefaultBody {
	return o.Payload
}

func (o *ConnectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ConnectDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ConnectBody connect body
swagger:model ConnectBody
*/
type ConnectBody struct {
	// User defined human readable PMM Server Name.
	ServerName string `json:"server_name,omitempty"`

	// Personal Access Token that the user obtains from Percona Portal.
	PersonalAccessToken string `json:"personal_access_token,omitempty"`
}

// Validate validates this connect body
func (o *ConnectBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this connect body based on context it is used
func (o *ConnectBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ConnectBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectBody) UnmarshalBinary(b []byte) error {
	var res ConnectBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ConnectDefaultBody connect default body
swagger:model ConnectDefaultBody
*/
type ConnectDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ConnectDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this connect default body
func (o *ConnectDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConnectDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("Connect default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Connect default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this connect default body based on the context it is used
func (o *ConnectDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConnectDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Connect default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Connect default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ConnectDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectDefaultBody) UnmarshalBinary(b []byte) error {
	var res ConnectDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ConnectDefaultBodyDetailsItems0 connect default body details items0
swagger:model ConnectDefaultBodyDetailsItems0
*/
type ConnectDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this connect default body details items0
func (o *ConnectDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this connect default body details items0 based on context it is used
func (o *ConnectDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ConnectDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ConnectDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
