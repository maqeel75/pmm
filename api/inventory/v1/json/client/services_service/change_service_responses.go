// Code generated by go-swagger; DO NOT EDIT.

package services_service

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

// ChangeServiceReader is a Reader for the ChangeService structure.
type ChangeServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChangeServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChangeServiceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeServiceOK creates a ChangeServiceOK with default headers values
func NewChangeServiceOK() *ChangeServiceOK {
	return &ChangeServiceOK{}
}

/*
ChangeServiceOK describes a response with status code 200, with default header values.

A successful response.
*/
type ChangeServiceOK struct {
	Payload interface{}
}

func (o *ChangeServiceOK) Error() string {
	return fmt.Sprintf("[PUT /v1/inventory/services/{service_id}][%d] changeServiceOk  %+v", 200, o.Payload)
}

func (o *ChangeServiceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ChangeServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeServiceDefault creates a ChangeServiceDefault with default headers values
func NewChangeServiceDefault(code int) *ChangeServiceDefault {
	return &ChangeServiceDefault{
		_statusCode: code,
	}
}

/*
ChangeServiceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ChangeServiceDefault struct {
	_statusCode int

	Payload *ChangeServiceDefaultBody
}

// Code gets the status code for the change service default response
func (o *ChangeServiceDefault) Code() int {
	return o._statusCode
}

func (o *ChangeServiceDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/inventory/services/{service_id}][%d] ChangeService default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeServiceDefault) GetPayload() *ChangeServiceDefaultBody {
	return o.Payload
}

func (o *ChangeServiceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ChangeServiceDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ChangeServiceBody change service body
swagger:model ChangeServiceBody
*/
type ChangeServiceBody struct {
	// environment
	Environment *string `json:"environment,omitempty"`

	// cluster
	Cluster *string `json:"cluster,omitempty"`

	// replication set
	ReplicationSet *string `json:"replication_set,omitempty"`

	// external group
	ExternalGroup *string `json:"external_group,omitempty"`

	// custom labels
	CustomLabels *ChangeServiceParamsBodyCustomLabels `json:"custom_labels,omitempty"`
}

// Validate validates this change service body
func (o *ChangeServiceBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCustomLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeServiceBody) validateCustomLabels(formats strfmt.Registry) error {
	if swag.IsZero(o.CustomLabels) { // not required
		return nil
	}

	if o.CustomLabels != nil {
		if err := o.CustomLabels.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "custom_labels")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "custom_labels")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this change service body based on the context it is used
func (o *ChangeServiceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCustomLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeServiceBody) contextValidateCustomLabels(ctx context.Context, formats strfmt.Registry) error {
	if o.CustomLabels != nil {
		if err := o.CustomLabels.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "custom_labels")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body" + "." + "custom_labels")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeServiceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeServiceBody) UnmarshalBinary(b []byte) error {
	var res ChangeServiceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeServiceDefaultBody change service default body
swagger:model ChangeServiceDefaultBody
*/
type ChangeServiceDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ChangeServiceDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this change service default body
func (o *ChangeServiceDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeServiceDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ChangeService default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangeService default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this change service default body based on the context it is used
func (o *ChangeServiceDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeServiceDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ChangeService default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChangeService default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeServiceDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeServiceDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeServiceDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeServiceDefaultBodyDetailsItems0 change service default body details items0
swagger:model ChangeServiceDefaultBodyDetailsItems0
*/
type ChangeServiceDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this change service default body details items0
func (o *ChangeServiceDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change service default body details items0 based on context it is used
func (o *ChangeServiceDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeServiceDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeServiceDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ChangeServiceDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ChangeServiceParamsBodyCustomLabels A wrapper for map[string]string. This type allows to distinguish between an empty map and a null value.
swagger:model ChangeServiceParamsBodyCustomLabels
*/
type ChangeServiceParamsBodyCustomLabels struct {
	// values
	Values map[string]string `json:"values,omitempty"`
}

// Validate validates this change service params body custom labels
func (o *ChangeServiceParamsBodyCustomLabels) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this change service params body custom labels based on context it is used
func (o *ChangeServiceParamsBodyCustomLabels) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeServiceParamsBodyCustomLabels) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeServiceParamsBodyCustomLabels) UnmarshalBinary(b []byte) error {
	var res ChangeServiceParamsBodyCustomLabels
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
