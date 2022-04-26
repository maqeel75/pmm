// Code generated by go-swagger; DO NOT EDIT.

package security_checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetSecurityCheckResultsReader is a Reader for the GetSecurityCheckResults structure.
type GetSecurityCheckResultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecurityCheckResultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSecurityCheckResultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSecurityCheckResultsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSecurityCheckResultsOK creates a GetSecurityCheckResultsOK with default headers values
func NewGetSecurityCheckResultsOK() *GetSecurityCheckResultsOK {
	return &GetSecurityCheckResultsOK{}
}

/* GetSecurityCheckResultsOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetSecurityCheckResultsOK struct {
	Payload *GetSecurityCheckResultsOKBody
}

func (o *GetSecurityCheckResultsOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/SecurityChecks/GetCheckResults][%d] getSecurityCheckResultsOk  %+v", 200, o.Payload)
}
func (o *GetSecurityCheckResultsOK) GetPayload() *GetSecurityCheckResultsOKBody {
	return o.Payload
}

func (o *GetSecurityCheckResultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetSecurityCheckResultsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecurityCheckResultsDefault creates a GetSecurityCheckResultsDefault with default headers values
func NewGetSecurityCheckResultsDefault(code int) *GetSecurityCheckResultsDefault {
	return &GetSecurityCheckResultsDefault{
		_statusCode: code,
	}
}

/* GetSecurityCheckResultsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetSecurityCheckResultsDefault struct {
	_statusCode int

	Payload *GetSecurityCheckResultsDefaultBody
}

// Code gets the status code for the get security check results default response
func (o *GetSecurityCheckResultsDefault) Code() int {
	return o._statusCode
}

func (o *GetSecurityCheckResultsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/SecurityChecks/GetCheckResults][%d] GetSecurityCheckResults default  %+v", o._statusCode, o.Payload)
}
func (o *GetSecurityCheckResultsDefault) GetPayload() *GetSecurityCheckResultsDefaultBody {
	return o.Payload
}

func (o *GetSecurityCheckResultsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetSecurityCheckResultsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetSecurityCheckResultsDefaultBody get security check results default body
swagger:model GetSecurityCheckResultsDefaultBody
*/
type GetSecurityCheckResultsDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*GetSecurityCheckResultsDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this get security check results default body
func (o *GetSecurityCheckResultsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSecurityCheckResultsDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("GetSecurityCheckResults default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetSecurityCheckResults default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get security check results default body based on the context it is used
func (o *GetSecurityCheckResultsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSecurityCheckResultsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetSecurityCheckResults default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetSecurityCheckResults default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSecurityCheckResultsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSecurityCheckResultsDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetSecurityCheckResultsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSecurityCheckResultsDefaultBodyDetailsItems0 get security check results default body details items0
swagger:model GetSecurityCheckResultsDefaultBodyDetailsItems0
*/
type GetSecurityCheckResultsDefaultBodyDetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this get security check results default body details items0
func (o *GetSecurityCheckResultsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get security check results default body details items0 based on context it is used
func (o *GetSecurityCheckResultsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSecurityCheckResultsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSecurityCheckResultsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetSecurityCheckResultsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSecurityCheckResultsOKBody get security check results OK body
swagger:model GetSecurityCheckResultsOKBody
*/
type GetSecurityCheckResultsOKBody struct {

	// results
	Results []*GetSecurityCheckResultsOKBodyResultsItems0 `json:"results"`
}

// Validate validates this get security check results OK body
func (o *GetSecurityCheckResultsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSecurityCheckResultsOKBody) validateResults(formats strfmt.Registry) error {
	if swag.IsZero(o.Results) { // not required
		return nil
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getSecurityCheckResultsOk" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getSecurityCheckResultsOk" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get security check results OK body based on the context it is used
func (o *GetSecurityCheckResultsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetSecurityCheckResultsOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getSecurityCheckResultsOk" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getSecurityCheckResultsOk" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetSecurityCheckResultsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSecurityCheckResultsOKBody) UnmarshalBinary(b []byte) error {
	var res GetSecurityCheckResultsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetSecurityCheckResultsOKBodyResultsItems0 SecurityCheckResult represents the check result returned from pmm-managed after running the check.
swagger:model GetSecurityCheckResultsOKBodyResultsItems0
*/
type GetSecurityCheckResultsOKBodyResultsItems0 struct {

	// summary
	Summary string `json:"summary,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// Severity represents severity level of the check result or alert.
	// Enum: [SEVERITY_INVALID SEVERITY_EMERGENCY SEVERITY_ALERT SEVERITY_CRITICAL SEVERITY_ERROR SEVERITY_WARNING SEVERITY_NOTICE SEVERITY_INFO SEVERITY_DEBUG]
	Severity *string `json:"severity,omitempty"`

	// labels
	Labels map[string]string `json:"labels,omitempty"`

	// URL containing information on how to resolve an issue detected by a check.
	ReadMoreURL string `json:"read_more_url,omitempty"`

	// Name of the monitored service on which the check ran.
	ServiceName string `json:"service_name,omitempty"`

	// Category to which the check belongs.
	Category string `json:"category,omitempty"`
}

// Validate validates this get security check results OK body results items0
func (o *GetSecurityCheckResultsOKBodyResultsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getSecurityCheckResultsOkBodyResultsItems0TypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SEVERITY_INVALID","SEVERITY_EMERGENCY","SEVERITY_ALERT","SEVERITY_CRITICAL","SEVERITY_ERROR","SEVERITY_WARNING","SEVERITY_NOTICE","SEVERITY_INFO","SEVERITY_DEBUG"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getSecurityCheckResultsOkBodyResultsItems0TypeSeverityPropEnum = append(getSecurityCheckResultsOkBodyResultsItems0TypeSeverityPropEnum, v)
	}
}

const (

	// GetSecurityCheckResultsOKBodyResultsItems0SeveritySEVERITYINVALID captures enum value "SEVERITY_INVALID"
	GetSecurityCheckResultsOKBodyResultsItems0SeveritySEVERITYINVALID string = "SEVERITY_INVALID"

	// GetSecurityCheckResultsOKBodyResultsItems0SeveritySEVERITYEMERGENCY captures enum value "SEVERITY_EMERGENCY"
	GetSecurityCheckResultsOKBodyResultsItems0SeveritySEVERITYEMERGENCY string = "SEVERITY_EMERGENCY"

	// GetSecurityCheckResultsOKBodyResultsItems0SeveritySEVERITYALERT captures enum value "SEVERITY_ALERT"
	GetSecurityCheckResultsOKBodyResultsItems0SeveritySEVERITYALERT string = "SEVERITY_ALERT"

	// GetSecurityCheckResultsOKBodyResultsItems0SeveritySEVERITYCRITICAL captures enum value "SEVERITY_CRITICAL"
	GetSecurityCheckResultsOKBodyResultsItems0SeveritySEVERITYCRITICAL string = "SEVERITY_CRITICAL"

	// GetSecurityCheckResultsOKBodyResultsItems0SeveritySEVERITYERROR captures enum value "SEVERITY_ERROR"
	GetSecurityCheckResultsOKBodyResultsItems0SeveritySEVERITYERROR string = "SEVERITY_ERROR"

	// GetSecurityCheckResultsOKBodyResultsItems0SeveritySEVERITYWARNING captures enum value "SEVERITY_WARNING"
	GetSecurityCheckResultsOKBodyResultsItems0SeveritySEVERITYWARNING string = "SEVERITY_WARNING"

	// GetSecurityCheckResultsOKBodyResultsItems0SeveritySEVERITYNOTICE captures enum value "SEVERITY_NOTICE"
	GetSecurityCheckResultsOKBodyResultsItems0SeveritySEVERITYNOTICE string = "SEVERITY_NOTICE"

	// GetSecurityCheckResultsOKBodyResultsItems0SeveritySEVERITYINFO captures enum value "SEVERITY_INFO"
	GetSecurityCheckResultsOKBodyResultsItems0SeveritySEVERITYINFO string = "SEVERITY_INFO"

	// GetSecurityCheckResultsOKBodyResultsItems0SeveritySEVERITYDEBUG captures enum value "SEVERITY_DEBUG"
	GetSecurityCheckResultsOKBodyResultsItems0SeveritySEVERITYDEBUG string = "SEVERITY_DEBUG"
)

// prop value enum
func (o *GetSecurityCheckResultsOKBodyResultsItems0) validateSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getSecurityCheckResultsOkBodyResultsItems0TypeSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetSecurityCheckResultsOKBodyResultsItems0) validateSeverity(formats strfmt.Registry) error {
	if swag.IsZero(o.Severity) { // not required
		return nil
	}

	// value enum
	if err := o.validateSeverityEnum("severity", "body", *o.Severity); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get security check results OK body results items0 based on context it is used
func (o *GetSecurityCheckResultsOKBodyResultsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetSecurityCheckResultsOKBodyResultsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetSecurityCheckResultsOKBodyResultsItems0) UnmarshalBinary(b []byte) error {
	var res GetSecurityCheckResultsOKBodyResultsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
