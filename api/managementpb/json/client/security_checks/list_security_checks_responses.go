// Code generated by go-swagger; DO NOT EDIT.

package security_checks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
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

// ListSecurityChecksReader is a Reader for the ListSecurityChecks structure.
type ListSecurityChecksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSecurityChecksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSecurityChecksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListSecurityChecksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSecurityChecksOK creates a ListSecurityChecksOK with default headers values
func NewListSecurityChecksOK() *ListSecurityChecksOK {
	return &ListSecurityChecksOK{}
}

/*ListSecurityChecksOK handles this case with default header values.

A successful response.
*/
type ListSecurityChecksOK struct {
	Payload *ListSecurityChecksOKBody
}

func (o *ListSecurityChecksOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/SecurityChecks/List][%d] listSecurityChecksOk  %+v", 200, o.Payload)
}

func (o *ListSecurityChecksOK) GetPayload() *ListSecurityChecksOKBody {
	return o.Payload
}

func (o *ListSecurityChecksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListSecurityChecksOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSecurityChecksDefault creates a ListSecurityChecksDefault with default headers values
func NewListSecurityChecksDefault(code int) *ListSecurityChecksDefault {
	return &ListSecurityChecksDefault{
		_statusCode: code,
	}
}

/*ListSecurityChecksDefault handles this case with default header values.

An unexpected error response.
*/
type ListSecurityChecksDefault struct {
	_statusCode int

	Payload *ListSecurityChecksDefaultBody
}

// Code gets the status code for the list security checks default response
func (o *ListSecurityChecksDefault) Code() int {
	return o._statusCode
}

func (o *ListSecurityChecksDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/SecurityChecks/List][%d] ListSecurityChecks default  %+v", o._statusCode, o.Payload)
}

func (o *ListSecurityChecksDefault) GetPayload() *ListSecurityChecksDefaultBody {
	return o.Payload
}

func (o *ListSecurityChecksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListSecurityChecksDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChecksItems0 SecurityCheck contains check name and status.
swagger:model ChecksItems0
*/
type ChecksItems0 struct {

	// Machine-readable name (ID) that is used in expression.
	Name string `json:"name,omitempty"`

	// True if that check is disabled.
	Disabled bool `json:"disabled,omitempty"`

	// Long human-readable description.
	Description string `json:"description,omitempty"`

	// Short human-readable summary.
	Summary string `json:"summary,omitempty"`

	// SecurityCheckInterval represents possible execution interval values for checks.
	// Enum: [SECURITY_CHECK_INTERVAL_INVALID STANDARD FREQUENT RARE]
	Interval *string `json:"interval,omitempty"`

	// Category to which the check belongs.
	Category string `json:"category,omitempty"`
}

// Validate validates this checks items0
func (o *ChecksItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var checksItems0TypeIntervalPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SECURITY_CHECK_INTERVAL_INVALID","STANDARD","FREQUENT","RARE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		checksItems0TypeIntervalPropEnum = append(checksItems0TypeIntervalPropEnum, v)
	}
}

const (

	// ChecksItems0IntervalSECURITYCHECKINTERVALINVALID captures enum value "SECURITY_CHECK_INTERVAL_INVALID"
	ChecksItems0IntervalSECURITYCHECKINTERVALINVALID string = "SECURITY_CHECK_INTERVAL_INVALID"

	// ChecksItems0IntervalSTANDARD captures enum value "STANDARD"
	ChecksItems0IntervalSTANDARD string = "STANDARD"

	// ChecksItems0IntervalFREQUENT captures enum value "FREQUENT"
	ChecksItems0IntervalFREQUENT string = "FREQUENT"

	// ChecksItems0IntervalRARE captures enum value "RARE"
	ChecksItems0IntervalRARE string = "RARE"
)

// prop value enum
func (o *ChecksItems0) validateIntervalEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, checksItems0TypeIntervalPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ChecksItems0) validateInterval(formats strfmt.Registry) error {

	if swag.IsZero(o.Interval) { // not required
		return nil
	}

	// value enum
	if err := o.validateIntervalEnum("interval", "body", *o.Interval); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChecksItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChecksItems0) UnmarshalBinary(b []byte) error {
	var res ChecksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListSecurityChecksBody list security checks body
swagger:model ListSecurityChecksBody
*/
type ListSecurityChecksBody struct {

	// filter params
	FilterParams *ListSecurityChecksParamsBodyFilterParams `json:"filter_params,omitempty"`
}

// Validate validates this list security checks body
func (o *ListSecurityChecksBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFilterParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListSecurityChecksBody) validateFilterParams(formats strfmt.Registry) error {

	if swag.IsZero(o.FilterParams) { // not required
		return nil
	}

	if o.FilterParams != nil {
		if err := o.FilterParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body" + "." + "filter_params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListSecurityChecksBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListSecurityChecksBody) UnmarshalBinary(b []byte) error {
	var res ListSecurityChecksBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListSecurityChecksDefaultBody list security checks default body
swagger:model ListSecurityChecksDefaultBody
*/
type ListSecurityChecksDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this list security checks default body
func (o *ListSecurityChecksDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListSecurityChecksDefaultBody) validateDetails(formats strfmt.Registry) error {

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
					return ve.ValidateName("ListSecurityChecks default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListSecurityChecksDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListSecurityChecksDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListSecurityChecksDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListSecurityChecksOKBody list security checks OK body
swagger:model ListSecurityChecksOKBody
*/
type ListSecurityChecksOKBody struct {

	// checks
	Checks []*ChecksItems0 `json:"checks"`

	// filter params
	FilterParams *ListSecurityChecksOKBodyFilterParams `json:"filter_params,omitempty"`
}

// Validate validates this list security checks OK body
func (o *ListSecurityChecksOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateChecks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFilterParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListSecurityChecksOKBody) validateChecks(formats strfmt.Registry) error {

	if swag.IsZero(o.Checks) { // not required
		return nil
	}

	for i := 0; i < len(o.Checks); i++ {
		if swag.IsZero(o.Checks[i]) { // not required
			continue
		}

		if o.Checks[i] != nil {
			if err := o.Checks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listSecurityChecksOk" + "." + "checks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ListSecurityChecksOKBody) validateFilterParams(formats strfmt.Registry) error {

	if swag.IsZero(o.FilterParams) { // not required
		return nil
	}

	if o.FilterParams != nil {
		if err := o.FilterParams.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("listSecurityChecksOk" + "." + "filter_params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListSecurityChecksOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListSecurityChecksOKBody) UnmarshalBinary(b []byte) error {
	var res ListSecurityChecksOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListSecurityChecksOKBodyFilterParams FilterParams is used to filter the data returned in a response.
swagger:model ListSecurityChecksOKBodyFilterParams
*/
type ListSecurityChecksOKBodyFilterParams struct {

	// Accepts multiple filters, but if/how they are handled is up to the specific implementation.
	Filters []*ListSecurityChecksOKBodyFilterParamsFiltersItems0 `json:"filters"`
}

// Validate validates this list security checks OK body filter params
func (o *ListSecurityChecksOKBodyFilterParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListSecurityChecksOKBodyFilterParams) validateFilters(formats strfmt.Registry) error {

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
					return ve.ValidateName("listSecurityChecksOk" + "." + "filter_params" + "." + "filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListSecurityChecksOKBodyFilterParams) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListSecurityChecksOKBodyFilterParams) UnmarshalBinary(b []byte) error {
	var res ListSecurityChecksOKBodyFilterParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListSecurityChecksOKBodyFilterParamsFiltersItems0 list security checks OK body filter params filters items0
swagger:model ListSecurityChecksOKBodyFilterParamsFiltersItems0
*/
type ListSecurityChecksOKBodyFilterParamsFiltersItems0 struct {

	// Key to filter the results by.
	Key string `json:"key,omitempty"`

	//  - EQUALS: Checks for results whose key is equal to the provided scalar value.
	// Valid for (int32 | int64 | string ).
	//  - NOT_EQUALS: Checks for results whose key is not equal to the provided scalar value.
	// Valid for (int32 | int64 | string).
	//  - BETWEEN: Returns results whose key lies within the inclusive range of the provided minimum and maximum values.
	// Valid for (IntRangeValues).
	//  - IN: Checks for results whose key exists in a given array.
	// Valid for (int_values | long_values | string_values).
	//  - CONTAINS: Checks for results whose key contains value as a substring
	// Valid for (string_values).
	// Enum: [INVALID EQUALS NOT_EQUALS BETWEEN IN CONTAINS]
	Op *string `json:"op,omitempty"`

	// int value
	IntValue int32 `json:"int_value,omitempty"`

	// long value
	LongValue string `json:"long_value,omitempty"`

	// string value
	StringValue string `json:"string_value,omitempty"`

	// int range values
	IntRangeValues *ListSecurityChecksOKBodyFilterParamsFiltersItems0IntRangeValues `json:"int_range_values,omitempty"`

	// int values
	IntValues *ListSecurityChecksOKBodyFilterParamsFiltersItems0IntValues `json:"int_values,omitempty"`

	// long values
	LongValues *ListSecurityChecksOKBodyFilterParamsFiltersItems0LongValues `json:"long_values,omitempty"`

	// string values
	StringValues *ListSecurityChecksOKBodyFilterParamsFiltersItems0StringValues `json:"string_values,omitempty"`
}

// Validate validates this list security checks OK body filter params filters items0
func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIntRangeValues(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIntValues(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLongValues(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStringValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var listSecurityChecksOkBodyFilterParamsFiltersItems0TypeOpPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INVALID","EQUALS","NOT_EQUALS","BETWEEN","IN","CONTAINS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listSecurityChecksOkBodyFilterParamsFiltersItems0TypeOpPropEnum = append(listSecurityChecksOkBodyFilterParamsFiltersItems0TypeOpPropEnum, v)
	}
}

const (

	// ListSecurityChecksOKBodyFilterParamsFiltersItems0OpINVALID captures enum value "INVALID"
	ListSecurityChecksOKBodyFilterParamsFiltersItems0OpINVALID string = "INVALID"

	// ListSecurityChecksOKBodyFilterParamsFiltersItems0OpEQUALS captures enum value "EQUALS"
	ListSecurityChecksOKBodyFilterParamsFiltersItems0OpEQUALS string = "EQUALS"

	// ListSecurityChecksOKBodyFilterParamsFiltersItems0OpNOTEQUALS captures enum value "NOT_EQUALS"
	ListSecurityChecksOKBodyFilterParamsFiltersItems0OpNOTEQUALS string = "NOT_EQUALS"

	// ListSecurityChecksOKBodyFilterParamsFiltersItems0OpBETWEEN captures enum value "BETWEEN"
	ListSecurityChecksOKBodyFilterParamsFiltersItems0OpBETWEEN string = "BETWEEN"

	// ListSecurityChecksOKBodyFilterParamsFiltersItems0OpIN captures enum value "IN"
	ListSecurityChecksOKBodyFilterParamsFiltersItems0OpIN string = "IN"

	// ListSecurityChecksOKBodyFilterParamsFiltersItems0OpCONTAINS captures enum value "CONTAINS"
	ListSecurityChecksOKBodyFilterParamsFiltersItems0OpCONTAINS string = "CONTAINS"
)

// prop value enum
func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0) validateOpEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listSecurityChecksOkBodyFilterParamsFiltersItems0TypeOpPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0) validateOp(formats strfmt.Registry) error {

	if swag.IsZero(o.Op) { // not required
		return nil
	}

	// value enum
	if err := o.validateOpEnum("op", "body", *o.Op); err != nil {
		return err
	}

	return nil
}

func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0) validateIntRangeValues(formats strfmt.Registry) error {

	if swag.IsZero(o.IntRangeValues) { // not required
		return nil
	}

	if o.IntRangeValues != nil {
		if err := o.IntRangeValues.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("int_range_values")
			}
			return err
		}
	}

	return nil
}

func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0) validateIntValues(formats strfmt.Registry) error {

	if swag.IsZero(o.IntValues) { // not required
		return nil
	}

	if o.IntValues != nil {
		if err := o.IntValues.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("int_values")
			}
			return err
		}
	}

	return nil
}

func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0) validateLongValues(formats strfmt.Registry) error {

	if swag.IsZero(o.LongValues) { // not required
		return nil
	}

	if o.LongValues != nil {
		if err := o.LongValues.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("long_values")
			}
			return err
		}
	}

	return nil
}

func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0) validateStringValues(formats strfmt.Registry) error {

	if swag.IsZero(o.StringValues) { // not required
		return nil
	}

	if o.StringValues != nil {
		if err := o.StringValues.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("string_values")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0) UnmarshalBinary(b []byte) error {
	var res ListSecurityChecksOKBodyFilterParamsFiltersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListSecurityChecksOKBodyFilterParamsFiltersItems0IntRangeValues list security checks OK body filter params filters items0 int range values
swagger:model ListSecurityChecksOKBodyFilterParamsFiltersItems0IntRangeValues
*/
type ListSecurityChecksOKBodyFilterParamsFiltersItems0IntRangeValues struct {

	// Minimum value in the range.
	Minimum int32 `json:"minimum,omitempty"`

	// Maximum value in the range.
	Maximum int32 `json:"maximum,omitempty"`
}

// Validate validates this list security checks OK body filter params filters items0 int range values
func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0IntRangeValues) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0IntRangeValues) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0IntRangeValues) UnmarshalBinary(b []byte) error {
	var res ListSecurityChecksOKBodyFilterParamsFiltersItems0IntRangeValues
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListSecurityChecksOKBodyFilterParamsFiltersItems0IntValues list security checks OK body filter params filters items0 int values
swagger:model ListSecurityChecksOKBodyFilterParamsFiltersItems0IntValues
*/
type ListSecurityChecksOKBodyFilterParamsFiltersItems0IntValues struct {

	// values
	Values []int32 `json:"values"`
}

// Validate validates this list security checks OK body filter params filters items0 int values
func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0IntValues) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0IntValues) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0IntValues) UnmarshalBinary(b []byte) error {
	var res ListSecurityChecksOKBodyFilterParamsFiltersItems0IntValues
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListSecurityChecksOKBodyFilterParamsFiltersItems0LongValues list security checks OK body filter params filters items0 long values
swagger:model ListSecurityChecksOKBodyFilterParamsFiltersItems0LongValues
*/
type ListSecurityChecksOKBodyFilterParamsFiltersItems0LongValues struct {

	// values
	Values []string `json:"values"`
}

// Validate validates this list security checks OK body filter params filters items0 long values
func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0LongValues) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0LongValues) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0LongValues) UnmarshalBinary(b []byte) error {
	var res ListSecurityChecksOKBodyFilterParamsFiltersItems0LongValues
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListSecurityChecksOKBodyFilterParamsFiltersItems0StringValues list security checks OK body filter params filters items0 string values
swagger:model ListSecurityChecksOKBodyFilterParamsFiltersItems0StringValues
*/
type ListSecurityChecksOKBodyFilterParamsFiltersItems0StringValues struct {

	// values
	Values []string `json:"values"`
}

// Validate validates this list security checks OK body filter params filters items0 string values
func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0StringValues) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0StringValues) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListSecurityChecksOKBodyFilterParamsFiltersItems0StringValues) UnmarshalBinary(b []byte) error {
	var res ListSecurityChecksOKBodyFilterParamsFiltersItems0StringValues
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListSecurityChecksParamsBodyFilterParams FilterParams is used to filter the data returned in a response.
swagger:model ListSecurityChecksParamsBodyFilterParams
*/
type ListSecurityChecksParamsBodyFilterParams struct {

	// Accepts multiple filters, but if/how they are handled is up to the specific implementation.
	Filters []*ListSecurityChecksParamsBodyFilterParamsFiltersItems0 `json:"filters"`
}

// Validate validates this list security checks params body filter params
func (o *ListSecurityChecksParamsBodyFilterParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListSecurityChecksParamsBodyFilterParams) validateFilters(formats strfmt.Registry) error {

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
					return ve.ValidateName("body" + "." + "filter_params" + "." + "filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListSecurityChecksParamsBodyFilterParams) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListSecurityChecksParamsBodyFilterParams) UnmarshalBinary(b []byte) error {
	var res ListSecurityChecksParamsBodyFilterParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListSecurityChecksParamsBodyFilterParamsFiltersItems0 list security checks params body filter params filters items0
swagger:model ListSecurityChecksParamsBodyFilterParamsFiltersItems0
*/
type ListSecurityChecksParamsBodyFilterParamsFiltersItems0 struct {

	// Key to filter the results by.
	Key string `json:"key,omitempty"`

	//  - EQUALS: Checks for results whose key is equal to the provided scalar value.
	// Valid for (int32 | int64 | string ).
	//  - NOT_EQUALS: Checks for results whose key is not equal to the provided scalar value.
	// Valid for (int32 | int64 | string).
	//  - BETWEEN: Returns results whose key lies within the inclusive range of the provided minimum and maximum values.
	// Valid for (IntRangeValues).
	//  - IN: Checks for results whose key exists in a given array.
	// Valid for (int_values | long_values | string_values).
	//  - CONTAINS: Checks for results whose key contains value as a substring
	// Valid for (string_values).
	// Enum: [INVALID EQUALS NOT_EQUALS BETWEEN IN CONTAINS]
	Op *string `json:"op,omitempty"`

	// int value
	IntValue int32 `json:"int_value,omitempty"`

	// long value
	LongValue string `json:"long_value,omitempty"`

	// string value
	StringValue string `json:"string_value,omitempty"`

	// int range values
	IntRangeValues *ListSecurityChecksParamsBodyFilterParamsFiltersItems0IntRangeValues `json:"int_range_values,omitempty"`

	// int values
	IntValues *ListSecurityChecksParamsBodyFilterParamsFiltersItems0IntValues `json:"int_values,omitempty"`

	// long values
	LongValues *ListSecurityChecksParamsBodyFilterParamsFiltersItems0LongValues `json:"long_values,omitempty"`

	// string values
	StringValues *ListSecurityChecksParamsBodyFilterParamsFiltersItems0StringValues `json:"string_values,omitempty"`
}

// Validate validates this list security checks params body filter params filters items0
func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIntRangeValues(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIntValues(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLongValues(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStringValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var listSecurityChecksParamsBodyFilterParamsFiltersItems0TypeOpPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INVALID","EQUALS","NOT_EQUALS","BETWEEN","IN","CONTAINS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listSecurityChecksParamsBodyFilterParamsFiltersItems0TypeOpPropEnum = append(listSecurityChecksParamsBodyFilterParamsFiltersItems0TypeOpPropEnum, v)
	}
}

const (

	// ListSecurityChecksParamsBodyFilterParamsFiltersItems0OpINVALID captures enum value "INVALID"
	ListSecurityChecksParamsBodyFilterParamsFiltersItems0OpINVALID string = "INVALID"

	// ListSecurityChecksParamsBodyFilterParamsFiltersItems0OpEQUALS captures enum value "EQUALS"
	ListSecurityChecksParamsBodyFilterParamsFiltersItems0OpEQUALS string = "EQUALS"

	// ListSecurityChecksParamsBodyFilterParamsFiltersItems0OpNOTEQUALS captures enum value "NOT_EQUALS"
	ListSecurityChecksParamsBodyFilterParamsFiltersItems0OpNOTEQUALS string = "NOT_EQUALS"

	// ListSecurityChecksParamsBodyFilterParamsFiltersItems0OpBETWEEN captures enum value "BETWEEN"
	ListSecurityChecksParamsBodyFilterParamsFiltersItems0OpBETWEEN string = "BETWEEN"

	// ListSecurityChecksParamsBodyFilterParamsFiltersItems0OpIN captures enum value "IN"
	ListSecurityChecksParamsBodyFilterParamsFiltersItems0OpIN string = "IN"

	// ListSecurityChecksParamsBodyFilterParamsFiltersItems0OpCONTAINS captures enum value "CONTAINS"
	ListSecurityChecksParamsBodyFilterParamsFiltersItems0OpCONTAINS string = "CONTAINS"
)

// prop value enum
func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0) validateOpEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listSecurityChecksParamsBodyFilterParamsFiltersItems0TypeOpPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0) validateOp(formats strfmt.Registry) error {

	if swag.IsZero(o.Op) { // not required
		return nil
	}

	// value enum
	if err := o.validateOpEnum("op", "body", *o.Op); err != nil {
		return err
	}

	return nil
}

func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0) validateIntRangeValues(formats strfmt.Registry) error {

	if swag.IsZero(o.IntRangeValues) { // not required
		return nil
	}

	if o.IntRangeValues != nil {
		if err := o.IntRangeValues.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("int_range_values")
			}
			return err
		}
	}

	return nil
}

func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0) validateIntValues(formats strfmt.Registry) error {

	if swag.IsZero(o.IntValues) { // not required
		return nil
	}

	if o.IntValues != nil {
		if err := o.IntValues.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("int_values")
			}
			return err
		}
	}

	return nil
}

func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0) validateLongValues(formats strfmt.Registry) error {

	if swag.IsZero(o.LongValues) { // not required
		return nil
	}

	if o.LongValues != nil {
		if err := o.LongValues.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("long_values")
			}
			return err
		}
	}

	return nil
}

func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0) validateStringValues(formats strfmt.Registry) error {

	if swag.IsZero(o.StringValues) { // not required
		return nil
	}

	if o.StringValues != nil {
		if err := o.StringValues.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("string_values")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0) UnmarshalBinary(b []byte) error {
	var res ListSecurityChecksParamsBodyFilterParamsFiltersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListSecurityChecksParamsBodyFilterParamsFiltersItems0IntRangeValues list security checks params body filter params filters items0 int range values
swagger:model ListSecurityChecksParamsBodyFilterParamsFiltersItems0IntRangeValues
*/
type ListSecurityChecksParamsBodyFilterParamsFiltersItems0IntRangeValues struct {

	// Minimum value in the range.
	Minimum int32 `json:"minimum,omitempty"`

	// Maximum value in the range.
	Maximum int32 `json:"maximum,omitempty"`
}

// Validate validates this list security checks params body filter params filters items0 int range values
func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0IntRangeValues) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0IntRangeValues) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0IntRangeValues) UnmarshalBinary(b []byte) error {
	var res ListSecurityChecksParamsBodyFilterParamsFiltersItems0IntRangeValues
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListSecurityChecksParamsBodyFilterParamsFiltersItems0IntValues list security checks params body filter params filters items0 int values
swagger:model ListSecurityChecksParamsBodyFilterParamsFiltersItems0IntValues
*/
type ListSecurityChecksParamsBodyFilterParamsFiltersItems0IntValues struct {

	// values
	Values []int32 `json:"values"`
}

// Validate validates this list security checks params body filter params filters items0 int values
func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0IntValues) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0IntValues) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0IntValues) UnmarshalBinary(b []byte) error {
	var res ListSecurityChecksParamsBodyFilterParamsFiltersItems0IntValues
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListSecurityChecksParamsBodyFilterParamsFiltersItems0LongValues list security checks params body filter params filters items0 long values
swagger:model ListSecurityChecksParamsBodyFilterParamsFiltersItems0LongValues
*/
type ListSecurityChecksParamsBodyFilterParamsFiltersItems0LongValues struct {

	// values
	Values []string `json:"values"`
}

// Validate validates this list security checks params body filter params filters items0 long values
func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0LongValues) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0LongValues) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0LongValues) UnmarshalBinary(b []byte) error {
	var res ListSecurityChecksParamsBodyFilterParamsFiltersItems0LongValues
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListSecurityChecksParamsBodyFilterParamsFiltersItems0StringValues list security checks params body filter params filters items0 string values
swagger:model ListSecurityChecksParamsBodyFilterParamsFiltersItems0StringValues
*/
type ListSecurityChecksParamsBodyFilterParamsFiltersItems0StringValues struct {

	// values
	Values []string `json:"values"`
}

// Validate validates this list security checks params body filter params filters items0 string values
func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0StringValues) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0StringValues) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListSecurityChecksParamsBodyFilterParamsFiltersItems0StringValues) UnmarshalBinary(b []byte) error {
	var res ListSecurityChecksParamsBodyFilterParamsFiltersItems0StringValues
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
