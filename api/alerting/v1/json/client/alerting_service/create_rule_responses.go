// Code generated by go-swagger; DO NOT EDIT.

package alerting_service

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

// CreateRuleReader is a Reader for the CreateRule structure.
type CreateRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCreateRuleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateRuleOK creates a CreateRuleOK with default headers values
func NewCreateRuleOK() *CreateRuleOK {
	return &CreateRuleOK{}
}

/*
CreateRuleOK describes a response with status code 200, with default header values.

A successful response.
*/
type CreateRuleOK struct {
	Payload interface{}
}

func (o *CreateRuleOK) Error() string {
	return fmt.Sprintf("[POST /v1/alerting/rules][%d] createRuleOk  %+v", 200, o.Payload)
}

func (o *CreateRuleOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRuleDefault creates a CreateRuleDefault with default headers values
func NewCreateRuleDefault(code int) *CreateRuleDefault {
	return &CreateRuleDefault{
		_statusCode: code,
	}
}

/*
CreateRuleDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CreateRuleDefault struct {
	_statusCode int

	Payload *CreateRuleDefaultBody
}

// Code gets the status code for the create rule default response
func (o *CreateRuleDefault) Code() int {
	return o._statusCode
}

func (o *CreateRuleDefault) Error() string {
	return fmt.Sprintf("[POST /v1/alerting/rules][%d] CreateRule default  %+v", o._statusCode, o.Payload)
}

func (o *CreateRuleDefault) GetPayload() *CreateRuleDefaultBody {
	return o.Payload
}

func (o *CreateRuleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(CreateRuleDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateRuleBody create rule body
swagger:model CreateRuleBody
*/
type CreateRuleBody struct {
	// Template name.
	TemplateName string `json:"template_name,omitempty"`

	// Rule name.
	Name string `json:"name,omitempty"`

	// Rule group name.
	Group string `json:"group,omitempty"`

	// Folder UID.
	FolderUID string `json:"folder_uid,omitempty"`

	// Rule parameters. All template parameters should be set.
	Params []*CreateRuleParamsBodyParamsItems0 `json:"params"`

	// Rule duration. Should be set.
	For string `json:"for,omitempty"`

	// Severity represents severity level of the check result or alert.
	// Enum: [SEVERITY_UNSPECIFIED SEVERITY_EMERGENCY SEVERITY_ALERT SEVERITY_CRITICAL SEVERITY_ERROR SEVERITY_WARNING SEVERITY_NOTICE SEVERITY_INFO SEVERITY_DEBUG]
	Severity *string `json:"severity,omitempty"`

	// All custom labels to add or remove (with empty values) to default labels from template.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Filters.
	Filters []*CreateRuleParamsBodyFiltersItems0 `json:"filters"`

	// Evaluation Interval
	Interval string `json:"interval,omitempty"`
}

// Validate validates this create rule body
func (o *CreateRuleBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateRuleBody) validateParams(formats strfmt.Registry) error {
	if swag.IsZero(o.Params) { // not required
		return nil
	}

	for i := 0; i < len(o.Params); i++ {
		if swag.IsZero(o.Params[i]) { // not required
			continue
		}

		if o.Params[i] != nil {
			if err := o.Params[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var createRuleBodyTypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SEVERITY_UNSPECIFIED","SEVERITY_EMERGENCY","SEVERITY_ALERT","SEVERITY_CRITICAL","SEVERITY_ERROR","SEVERITY_WARNING","SEVERITY_NOTICE","SEVERITY_INFO","SEVERITY_DEBUG"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createRuleBodyTypeSeverityPropEnum = append(createRuleBodyTypeSeverityPropEnum, v)
	}
}

const (

	// CreateRuleBodySeveritySEVERITYUNSPECIFIED captures enum value "SEVERITY_UNSPECIFIED"
	CreateRuleBodySeveritySEVERITYUNSPECIFIED string = "SEVERITY_UNSPECIFIED"

	// CreateRuleBodySeveritySEVERITYEMERGENCY captures enum value "SEVERITY_EMERGENCY"
	CreateRuleBodySeveritySEVERITYEMERGENCY string = "SEVERITY_EMERGENCY"

	// CreateRuleBodySeveritySEVERITYALERT captures enum value "SEVERITY_ALERT"
	CreateRuleBodySeveritySEVERITYALERT string = "SEVERITY_ALERT"

	// CreateRuleBodySeveritySEVERITYCRITICAL captures enum value "SEVERITY_CRITICAL"
	CreateRuleBodySeveritySEVERITYCRITICAL string = "SEVERITY_CRITICAL"

	// CreateRuleBodySeveritySEVERITYERROR captures enum value "SEVERITY_ERROR"
	CreateRuleBodySeveritySEVERITYERROR string = "SEVERITY_ERROR"

	// CreateRuleBodySeveritySEVERITYWARNING captures enum value "SEVERITY_WARNING"
	CreateRuleBodySeveritySEVERITYWARNING string = "SEVERITY_WARNING"

	// CreateRuleBodySeveritySEVERITYNOTICE captures enum value "SEVERITY_NOTICE"
	CreateRuleBodySeveritySEVERITYNOTICE string = "SEVERITY_NOTICE"

	// CreateRuleBodySeveritySEVERITYINFO captures enum value "SEVERITY_INFO"
	CreateRuleBodySeveritySEVERITYINFO string = "SEVERITY_INFO"

	// CreateRuleBodySeveritySEVERITYDEBUG captures enum value "SEVERITY_DEBUG"
	CreateRuleBodySeveritySEVERITYDEBUG string = "SEVERITY_DEBUG"
)

// prop value enum
func (o *CreateRuleBody) validateSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createRuleBodyTypeSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateRuleBody) validateSeverity(formats strfmt.Registry) error {
	if swag.IsZero(o.Severity) { // not required
		return nil
	}

	// value enum
	if err := o.validateSeverityEnum("body"+"."+"severity", "body", *o.Severity); err != nil {
		return err
	}

	return nil
}

func (o *CreateRuleBody) validateFilters(formats strfmt.Registry) error {
	if swag.IsZero(o.Filters) { // not required
		return nil
	}

	for i := 0; i < len(o.Filters); i++ {
		if swag.IsZero(o.Filters[i]) { // not required
			continue
		}

		if o.Filters[i] != nil {
			if err := o.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create rule body based on the context it is used
func (o *CreateRuleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateRuleBody) contextValidateParams(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Params); i++ {
		if o.Params[i] != nil {
			if err := o.Params[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "params" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

func (o *CreateRuleBody) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Filters); i++ {
		if o.Filters[i] != nil {
			if err := o.Filters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateRuleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateRuleBody) UnmarshalBinary(b []byte) error {
	var res CreateRuleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateRuleDefaultBody create rule default body
swagger:model CreateRuleDefaultBody
*/
type CreateRuleDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*CreateRuleDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this create rule default body
func (o *CreateRuleDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateRuleDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("CreateRule default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("CreateRule default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create rule default body based on the context it is used
func (o *CreateRuleDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateRuleDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CreateRule default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("CreateRule default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateRuleDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateRuleDefaultBody) UnmarshalBinary(b []byte) error {
	var res CreateRuleDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateRuleDefaultBodyDetailsItems0 create rule default body details items0
swagger:model CreateRuleDefaultBodyDetailsItems0
*/
type CreateRuleDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this create rule default body details items0
func (o *CreateRuleDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create rule default body details items0 based on context it is used
func (o *CreateRuleDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateRuleDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateRuleDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res CreateRuleDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateRuleParamsBodyFiltersItems0 Filter represents a single filter condition.
swagger:model CreateRuleParamsBodyFiltersItems0
*/
type CreateRuleParamsBodyFiltersItems0 struct {
	// FilterType represents filter matching type.
	// Enum: [FILTER_TYPE_UNSPECIFIED FILTER_TYPE_MATCH FILTER_TYPE_MISMATCH]
	Type *string `json:"type,omitempty"`

	// label
	Label string `json:"label,omitempty"`

	// regexp
	Regexp string `json:"regexp,omitempty"`
}

// Validate validates this create rule params body filters items0
func (o *CreateRuleParamsBodyFiltersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createRuleParamsBodyFiltersItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FILTER_TYPE_UNSPECIFIED","FILTER_TYPE_MATCH","FILTER_TYPE_MISMATCH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createRuleParamsBodyFiltersItems0TypeTypePropEnum = append(createRuleParamsBodyFiltersItems0TypeTypePropEnum, v)
	}
}

const (

	// CreateRuleParamsBodyFiltersItems0TypeFILTERTYPEUNSPECIFIED captures enum value "FILTER_TYPE_UNSPECIFIED"
	CreateRuleParamsBodyFiltersItems0TypeFILTERTYPEUNSPECIFIED string = "FILTER_TYPE_UNSPECIFIED"

	// CreateRuleParamsBodyFiltersItems0TypeFILTERTYPEMATCH captures enum value "FILTER_TYPE_MATCH"
	CreateRuleParamsBodyFiltersItems0TypeFILTERTYPEMATCH string = "FILTER_TYPE_MATCH"

	// CreateRuleParamsBodyFiltersItems0TypeFILTERTYPEMISMATCH captures enum value "FILTER_TYPE_MISMATCH"
	CreateRuleParamsBodyFiltersItems0TypeFILTERTYPEMISMATCH string = "FILTER_TYPE_MISMATCH"
)

// prop value enum
func (o *CreateRuleParamsBodyFiltersItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createRuleParamsBodyFiltersItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateRuleParamsBodyFiltersItems0) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create rule params body filters items0 based on context it is used
func (o *CreateRuleParamsBodyFiltersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateRuleParamsBodyFiltersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateRuleParamsBodyFiltersItems0) UnmarshalBinary(b []byte) error {
	var res CreateRuleParamsBodyFiltersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateRuleParamsBodyParamsItems0 ParamValue represents a single rule parameter value.
swagger:model CreateRuleParamsBodyParamsItems0
*/
type CreateRuleParamsBodyParamsItems0 struct {
	// Machine-readable name (ID) that is used in expression.
	Name string `json:"name,omitempty"`

	// ParamType represents template parameter type.
	// Enum: [PARAM_TYPE_UNSPECIFIED PARAM_TYPE_BOOL PARAM_TYPE_FLOAT PARAM_TYPE_STRING]
	Type *string `json:"type,omitempty"`

	// Bool value.
	Bool bool `json:"bool,omitempty"`

	// Float value.
	Float float64 `json:"float,omitempty"`

	// String value.
	String string `json:"string,omitempty"`
}

// Validate validates this create rule params body params items0
func (o *CreateRuleParamsBodyParamsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createRuleParamsBodyParamsItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PARAM_TYPE_UNSPECIFIED","PARAM_TYPE_BOOL","PARAM_TYPE_FLOAT","PARAM_TYPE_STRING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createRuleParamsBodyParamsItems0TypeTypePropEnum = append(createRuleParamsBodyParamsItems0TypeTypePropEnum, v)
	}
}

const (

	// CreateRuleParamsBodyParamsItems0TypePARAMTYPEUNSPECIFIED captures enum value "PARAM_TYPE_UNSPECIFIED"
	CreateRuleParamsBodyParamsItems0TypePARAMTYPEUNSPECIFIED string = "PARAM_TYPE_UNSPECIFIED"

	// CreateRuleParamsBodyParamsItems0TypePARAMTYPEBOOL captures enum value "PARAM_TYPE_BOOL"
	CreateRuleParamsBodyParamsItems0TypePARAMTYPEBOOL string = "PARAM_TYPE_BOOL"

	// CreateRuleParamsBodyParamsItems0TypePARAMTYPEFLOAT captures enum value "PARAM_TYPE_FLOAT"
	CreateRuleParamsBodyParamsItems0TypePARAMTYPEFLOAT string = "PARAM_TYPE_FLOAT"

	// CreateRuleParamsBodyParamsItems0TypePARAMTYPESTRING captures enum value "PARAM_TYPE_STRING"
	CreateRuleParamsBodyParamsItems0TypePARAMTYPESTRING string = "PARAM_TYPE_STRING"
)

// prop value enum
func (o *CreateRuleParamsBodyParamsItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createRuleParamsBodyParamsItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateRuleParamsBodyParamsItems0) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create rule params body params items0 based on context it is used
func (o *CreateRuleParamsBodyParamsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateRuleParamsBodyParamsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateRuleParamsBodyParamsItems0) UnmarshalBinary(b []byte) error {
	var res CreateRuleParamsBodyParamsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
