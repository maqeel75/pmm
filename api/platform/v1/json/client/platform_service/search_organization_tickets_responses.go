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
	"github.com/go-openapi/validate"
)

// SearchOrganizationTicketsReader is a Reader for the SearchOrganizationTickets structure.
type SearchOrganizationTicketsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchOrganizationTicketsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchOrganizationTicketsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSearchOrganizationTicketsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSearchOrganizationTicketsOK creates a SearchOrganizationTicketsOK with default headers values
func NewSearchOrganizationTicketsOK() *SearchOrganizationTicketsOK {
	return &SearchOrganizationTicketsOK{}
}

/*
SearchOrganizationTicketsOK describes a response with status code 200, with default header values.

A successful response.
*/
type SearchOrganizationTicketsOK struct {
	Payload *SearchOrganizationTicketsOKBody
}

func (o *SearchOrganizationTicketsOK) Error() string {
	return fmt.Sprintf("[POST /v1/platform/SearchOrganizationTickets][%d] searchOrganizationTicketsOk  %+v", 200, o.Payload)
}

func (o *SearchOrganizationTicketsOK) GetPayload() *SearchOrganizationTicketsOKBody {
	return o.Payload
}

func (o *SearchOrganizationTicketsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(SearchOrganizationTicketsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchOrganizationTicketsDefault creates a SearchOrganizationTicketsDefault with default headers values
func NewSearchOrganizationTicketsDefault(code int) *SearchOrganizationTicketsDefault {
	return &SearchOrganizationTicketsDefault{
		_statusCode: code,
	}
}

/*
SearchOrganizationTicketsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type SearchOrganizationTicketsDefault struct {
	_statusCode int

	Payload *SearchOrganizationTicketsDefaultBody
}

// Code gets the status code for the search organization tickets default response
func (o *SearchOrganizationTicketsDefault) Code() int {
	return o._statusCode
}

func (o *SearchOrganizationTicketsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/platform/SearchOrganizationTickets][%d] SearchOrganizationTickets default  %+v", o._statusCode, o.Payload)
}

func (o *SearchOrganizationTicketsDefault) GetPayload() *SearchOrganizationTicketsDefaultBody {
	return o.Payload
}

func (o *SearchOrganizationTicketsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(SearchOrganizationTicketsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SearchOrganizationTicketsDefaultBody search organization tickets default body
swagger:model SearchOrganizationTicketsDefaultBody
*/
type SearchOrganizationTicketsDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*SearchOrganizationTicketsDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this search organization tickets default body
func (o *SearchOrganizationTicketsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchOrganizationTicketsDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("SearchOrganizationTickets default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SearchOrganizationTickets default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this search organization tickets default body based on the context it is used
func (o *SearchOrganizationTicketsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchOrganizationTicketsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SearchOrganizationTickets default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SearchOrganizationTickets default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SearchOrganizationTicketsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchOrganizationTicketsDefaultBody) UnmarshalBinary(b []byte) error {
	var res SearchOrganizationTicketsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchOrganizationTicketsDefaultBodyDetailsItems0 search organization tickets default body details items0
swagger:model SearchOrganizationTicketsDefaultBodyDetailsItems0
*/
type SearchOrganizationTicketsDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this search organization tickets default body details items0
func (o *SearchOrganizationTicketsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search organization tickets default body details items0 based on context it is used
func (o *SearchOrganizationTicketsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchOrganizationTicketsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchOrganizationTicketsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res SearchOrganizationTicketsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchOrganizationTicketsOKBody search organization tickets OK body
swagger:model SearchOrganizationTicketsOKBody
*/
type SearchOrganizationTicketsOKBody struct {
	// Support tickets belonging to the Percona Portal Organization.
	Tickets []*SearchOrganizationTicketsOKBodyTicketsItems0 `json:"tickets"`
}

// Validate validates this search organization tickets OK body
func (o *SearchOrganizationTicketsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTickets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchOrganizationTicketsOKBody) validateTickets(formats strfmt.Registry) error {
	if swag.IsZero(o.Tickets) { // not required
		return nil
	}

	for i := 0; i < len(o.Tickets); i++ {
		if swag.IsZero(o.Tickets[i]) { // not required
			continue
		}

		if o.Tickets[i] != nil {
			if err := o.Tickets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("searchOrganizationTicketsOk" + "." + "tickets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("searchOrganizationTicketsOk" + "." + "tickets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this search organization tickets OK body based on the context it is used
func (o *SearchOrganizationTicketsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateTickets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchOrganizationTicketsOKBody) contextValidateTickets(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Tickets); i++ {
		if o.Tickets[i] != nil {
			if err := o.Tickets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("searchOrganizationTicketsOk" + "." + "tickets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("searchOrganizationTicketsOk" + "." + "tickets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SearchOrganizationTicketsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchOrganizationTicketsOKBody) UnmarshalBinary(b []byte) error {
	var res SearchOrganizationTicketsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SearchOrganizationTicketsOKBodyTicketsItems0 OrganizationTicket contains information about the support ticket.
swagger:model SearchOrganizationTicketsOKBodyTicketsItems0
*/
type SearchOrganizationTicketsOKBodyTicketsItems0 struct {
	// Ticket number.
	Number string `json:"number,omitempty"`

	// Ticket short description.
	ShortDescription string `json:"short_description,omitempty"`

	// Ticket priority.
	Priority string `json:"priority,omitempty"`

	// Ticket state.
	State string `json:"state,omitempty"`

	// Ticket creation time.
	// Format: date-time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// Department.
	Department string `json:"department,omitempty"`

	// Ticket requester.
	Requester string `json:"requester,omitempty"`

	// Task type.
	TaskType string `json:"task_type,omitempty"`

	// Ticket url.
	URL string `json:"url,omitempty"`
}

// Validate validates this search organization tickets OK body tickets items0
func (o *SearchOrganizationTicketsOKBodyTicketsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SearchOrganizationTicketsOKBodyTicketsItems0) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(o.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", o.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this search organization tickets OK body tickets items0 based on context it is used
func (o *SearchOrganizationTicketsOKBodyTicketsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SearchOrganizationTicketsOKBodyTicketsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SearchOrganizationTicketsOKBodyTicketsItems0) UnmarshalBinary(b []byte) error {
	var res SearchOrganizationTicketsOKBodyTicketsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
