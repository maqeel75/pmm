// Code generated by go-swagger; DO NOT EDIT.

package restore_service

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

// ListRestoresReader is a Reader for the ListRestores structure.
type ListRestoresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRestoresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRestoresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListRestoresDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListRestoresOK creates a ListRestoresOK with default headers values
func NewListRestoresOK() *ListRestoresOK {
	return &ListRestoresOK{}
}

/*
ListRestoresOK describes a response with status code 200, with default header values.

A successful response.
*/
type ListRestoresOK struct {
	Payload *ListRestoresOKBody
}

func (o *ListRestoresOK) Error() string {
	return fmt.Sprintf("[GET /v1/backups/restores][%d] listRestoresOk  %+v", 200, o.Payload)
}

func (o *ListRestoresOK) GetPayload() *ListRestoresOKBody {
	return o.Payload
}

func (o *ListRestoresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListRestoresOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRestoresDefault creates a ListRestoresDefault with default headers values
func NewListRestoresDefault(code int) *ListRestoresDefault {
	return &ListRestoresDefault{
		_statusCode: code,
	}
}

/*
ListRestoresDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ListRestoresDefault struct {
	_statusCode int

	Payload *ListRestoresDefaultBody
}

// Code gets the status code for the list restores default response
func (o *ListRestoresDefault) Code() int {
	return o._statusCode
}

func (o *ListRestoresDefault) Error() string {
	return fmt.Sprintf("[GET /v1/backups/restores][%d] ListRestores default  %+v", o._statusCode, o.Payload)
}

func (o *ListRestoresDefault) GetPayload() *ListRestoresDefaultBody {
	return o.Payload
}

func (o *ListRestoresDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(ListRestoresDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ListRestoresDefaultBody list restores default body
swagger:model ListRestoresDefaultBody
*/
type ListRestoresDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*ListRestoresDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this list restores default body
func (o *ListRestoresDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListRestoresDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ListRestores default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListRestores default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list restores default body based on the context it is used
func (o *ListRestoresDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListRestoresDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListRestores default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ListRestores default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListRestoresDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListRestoresDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListRestoresDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListRestoresDefaultBodyDetailsItems0 list restores default body details items0
swagger:model ListRestoresDefaultBodyDetailsItems0
*/
type ListRestoresDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this list restores default body details items0
func (o *ListRestoresDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this list restores default body details items0 based on context it is used
func (o *ListRestoresDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListRestoresDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListRestoresDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ListRestoresDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListRestoresOKBody list restores OK body
swagger:model ListRestoresOKBody
*/
type ListRestoresOKBody struct {
	// items
	Items []*ListRestoresOKBodyItemsItems0 `json:"items"`
}

// Validate validates this list restores OK body
func (o *ListRestoresOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListRestoresOKBody) validateItems(formats strfmt.Registry) error {
	if swag.IsZero(o.Items) { // not required
		return nil
	}

	for i := 0; i < len(o.Items); i++ {
		if swag.IsZero(o.Items[i]) { // not required
			continue
		}

		if o.Items[i] != nil {
			if err := o.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listRestoresOk" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listRestoresOk" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list restores OK body based on the context it is used
func (o *ListRestoresOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListRestoresOKBody) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Items); i++ {
		if o.Items[i] != nil {
			if err := o.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listRestoresOk" + "." + "items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("listRestoresOk" + "." + "items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListRestoresOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListRestoresOKBody) UnmarshalBinary(b []byte) error {
	var res ListRestoresOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
ListRestoresOKBodyItemsItems0 RestoreHistoryItem represents single backup restore item.
swagger:model ListRestoresOKBodyItemsItems0
*/
type ListRestoresOKBodyItemsItems0 struct {
	// Machine-readable restore id.
	RestoreID string `json:"restore_id,omitempty"`

	// ID of the artifact used for restore.
	ArtifactID string `json:"artifact_id,omitempty"`

	// Artifact name used for restore.
	Name string `json:"name,omitempty"`

	// Database vendor e.g. PostgreSQL, MongoDB, MySQL.
	Vendor string `json:"vendor,omitempty"`

	// Machine-readable location ID.
	LocationID string `json:"location_id,omitempty"`

	// Location name.
	LocationName string `json:"location_name,omitempty"`

	// Machine-readable service ID.
	ServiceID string `json:"service_id,omitempty"`

	// Service name.
	ServiceName string `json:"service_name,omitempty"`

	// DataModel is a model used for performing a backup.
	// Enum: [DATA_MODEL_UNSPECIFIED DATA_MODEL_PHYSICAL DATA_MODEL_LOGICAL]
	DataModel *string `json:"data_model,omitempty"`

	// RestoreStatus shows the current status of execution of restore.
	// Enum: [RESTORE_STATUS_UNSPECIFIED RESTORE_STATUS_IN_PROGRESS RESTORE_STATUS_SUCCESS RESTORE_STATUS_ERROR]
	Status *string `json:"status,omitempty"`

	// Restore start time.
	// Format: date-time
	StartedAt strfmt.DateTime `json:"started_at,omitempty"`

	// Restore finish time.
	// Format: date-time
	FinishedAt strfmt.DateTime `json:"finished_at,omitempty"`

	// PITR timestamp is filled for PITR restores, empty otherwise.
	// Format: date-time
	PitrTimestamp strfmt.DateTime `json:"pitr_timestamp,omitempty"`
}

// Validate validates this list restores OK body items items0
func (o *ListRestoresOKBodyItemsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDataModel(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFinishedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePitrTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var listRestoresOkBodyItemsItems0TypeDataModelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DATA_MODEL_UNSPECIFIED","DATA_MODEL_PHYSICAL","DATA_MODEL_LOGICAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listRestoresOkBodyItemsItems0TypeDataModelPropEnum = append(listRestoresOkBodyItemsItems0TypeDataModelPropEnum, v)
	}
}

const (

	// ListRestoresOKBodyItemsItems0DataModelDATAMODELUNSPECIFIED captures enum value "DATA_MODEL_UNSPECIFIED"
	ListRestoresOKBodyItemsItems0DataModelDATAMODELUNSPECIFIED string = "DATA_MODEL_UNSPECIFIED"

	// ListRestoresOKBodyItemsItems0DataModelDATAMODELPHYSICAL captures enum value "DATA_MODEL_PHYSICAL"
	ListRestoresOKBodyItemsItems0DataModelDATAMODELPHYSICAL string = "DATA_MODEL_PHYSICAL"

	// ListRestoresOKBodyItemsItems0DataModelDATAMODELLOGICAL captures enum value "DATA_MODEL_LOGICAL"
	ListRestoresOKBodyItemsItems0DataModelDATAMODELLOGICAL string = "DATA_MODEL_LOGICAL"
)

// prop value enum
func (o *ListRestoresOKBodyItemsItems0) validateDataModelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listRestoresOkBodyItemsItems0TypeDataModelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListRestoresOKBodyItemsItems0) validateDataModel(formats strfmt.Registry) error {
	if swag.IsZero(o.DataModel) { // not required
		return nil
	}

	// value enum
	if err := o.validateDataModelEnum("data_model", "body", *o.DataModel); err != nil {
		return err
	}

	return nil
}

var listRestoresOkBodyItemsItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RESTORE_STATUS_UNSPECIFIED","RESTORE_STATUS_IN_PROGRESS","RESTORE_STATUS_SUCCESS","RESTORE_STATUS_ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		listRestoresOkBodyItemsItems0TypeStatusPropEnum = append(listRestoresOkBodyItemsItems0TypeStatusPropEnum, v)
	}
}

const (

	// ListRestoresOKBodyItemsItems0StatusRESTORESTATUSUNSPECIFIED captures enum value "RESTORE_STATUS_UNSPECIFIED"
	ListRestoresOKBodyItemsItems0StatusRESTORESTATUSUNSPECIFIED string = "RESTORE_STATUS_UNSPECIFIED"

	// ListRestoresOKBodyItemsItems0StatusRESTORESTATUSINPROGRESS captures enum value "RESTORE_STATUS_IN_PROGRESS"
	ListRestoresOKBodyItemsItems0StatusRESTORESTATUSINPROGRESS string = "RESTORE_STATUS_IN_PROGRESS"

	// ListRestoresOKBodyItemsItems0StatusRESTORESTATUSSUCCESS captures enum value "RESTORE_STATUS_SUCCESS"
	ListRestoresOKBodyItemsItems0StatusRESTORESTATUSSUCCESS string = "RESTORE_STATUS_SUCCESS"

	// ListRestoresOKBodyItemsItems0StatusRESTORESTATUSERROR captures enum value "RESTORE_STATUS_ERROR"
	ListRestoresOKBodyItemsItems0StatusRESTORESTATUSERROR string = "RESTORE_STATUS_ERROR"
)

// prop value enum
func (o *ListRestoresOKBodyItemsItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, listRestoresOkBodyItemsItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ListRestoresOKBodyItemsItems0) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

func (o *ListRestoresOKBodyItemsItems0) validateStartedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at", "body", "date-time", o.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ListRestoresOKBodyItemsItems0) validateFinishedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.FinishedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("finished_at", "body", "date-time", o.FinishedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *ListRestoresOKBodyItemsItems0) validatePitrTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(o.PitrTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("pitr_timestamp", "body", "date-time", o.PitrTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this list restores OK body items items0 based on context it is used
func (o *ListRestoresOKBodyItemsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ListRestoresOKBodyItemsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListRestoresOKBodyItemsItems0) UnmarshalBinary(b []byte) error {
	var res ListRestoresOKBodyItemsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
