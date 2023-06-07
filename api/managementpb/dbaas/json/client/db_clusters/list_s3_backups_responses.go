// Code generated by go-swagger; DO NOT EDIT.

package db_clusters

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

// ListS3BackupsReader is a Reader for the ListS3Backups structure.
type ListS3BackupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListS3BackupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListS3BackupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListS3BackupsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListS3BackupsOK creates a ListS3BackupsOK with default headers values
func NewListS3BackupsOK() *ListS3BackupsOK {
	return &ListS3BackupsOK{}
}

/*
ListS3BackupsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListS3BackupsOK struct {
	Payload *ListS3BackupsOKBody
}

func (o *ListS3BackupsOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Backups/List][%d] listS3BackupsOk  %+v", 200, o.Payload)
}

func (o *ListS3BackupsOK) GetPayload() *ListS3BackupsOKBody {
	return o.Payload
}

func (o *ListS3BackupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListS3BackupsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListS3BackupsDefault creates a ListS3BackupsDefault with default headers values
func NewListS3BackupsDefault(code int) *ListS3BackupsDefault {
	return &ListS3BackupsDefault{
		_statusCode: code,
	}
}

/*
ListS3BackupsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListS3BackupsDefault struct {
	_statusCode int

	Payload *ListS3BackupsDefaultBody
}

// Code gets the status code for the list s3 backups default response
func (o *ListS3BackupsDefault) Code() int {
	return o._statusCode
}

func (o *ListS3BackupsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Backups/List][%d] ListS3Backups default  %+v", o._statusCode, o.Payload)
}

func (o *ListS3BackupsDefault) GetPayload() *ListS3BackupsDefaultBody {
	return o.Payload
}

func (o *ListS3BackupsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListS3BackupsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ListS3BackupsBody list s3 backups body
swagger:model ListS3BackupsBody
*/
type ListS3BackupsBody struct {
	// Backup Location ID to list backups from
	LocationID string `json:"location_id,omitempty"`
}

// Validate validates this list s3 backups body
func (o *ListS3BackupsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list s3 backups body based on context it is used
func (o *ListS3BackupsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListS3BackupsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListS3BackupsBody) UnmarshalBinary(b []byte) error {
	var res ListS3BackupsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListS3BackupsDefaultBody list s3 backups default body
swagger:model ListS3BackupsDefaultBody
*/
type ListS3BackupsDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ListS3BackupsDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this list s3 backups default body
func (o *ListS3BackupsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListS3BackupsDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ListS3Backups default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListS3Backups default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list s3 backups default body based on the context it is used
func (o *ListS3BackupsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListS3BackupsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListS3Backups default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListS3Backups default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListS3BackupsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListS3BackupsDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListS3BackupsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListS3BackupsDefaultBodyDetailsItems0 list s3 backups default body details items0
swagger:model ListS3BackupsDefaultBodyDetailsItems0
*/
type ListS3BackupsDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this list s3 backups default body details items0
func (o *ListS3BackupsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list s3 backups default body details items0 based on context it is used
func (o *ListS3BackupsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListS3BackupsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListS3BackupsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ListS3BackupsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListS3BackupsOKBody list s3 backups OK body
swagger:model ListS3BackupsOKBody
*/
type ListS3BackupsOKBody struct {
	// Backup list.
	Backups []*ListS3BackupsOKBodyBackupsItems0 `json:"backups"`
}

// Validate validates this list s3 backups OK body
func (o *ListS3BackupsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBackups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListS3BackupsOKBody) validateBackups(formats strfmt.Registry) error {
	if swag.IsZero(o.Backups) { // not required
		return nil
	}

	for i := 0; i < len(o.Backups); i++ {
		if swag.IsZero(o.Backups[i]) { // not required
			continue
		}

		if o.Backups[i] != nil {
			if err := o.Backups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listS3BackupsOk" + "." + "backups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listS3BackupsOk" + "." + "backups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list s3 backups OK body based on the context it is used
func (o *ListS3BackupsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBackups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListS3BackupsOKBody) contextValidateBackups(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Backups); i++ {
		if o.Backups[i] != nil {
			if err := o.Backups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listS3BackupsOk" + "." + "backups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listS3BackupsOk" + "." + "backups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListS3BackupsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListS3BackupsOKBody) UnmarshalBinary(b []byte) error {
	var res ListS3BackupsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListS3BackupsOKBodyBackupsItems0 list s3 backups OK body backups items0
swagger:model ListS3BackupsOKBodyBackupsItems0
*/
type ListS3BackupsOKBodyBackupsItems0 struct {
	// Key of a filename on s3.
	Key string `json:"key,omitempty"`
}

// Validate validates this list s3 backups OK body backups items0
func (o *ListS3BackupsOKBodyBackupsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list s3 backups OK body backups items0 based on context it is used
func (o *ListS3BackupsOKBodyBackupsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListS3BackupsOKBodyBackupsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListS3BackupsOKBodyBackupsItems0) UnmarshalBinary(b []byte) error {
	var res ListS3BackupsOKBodyBackupsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
