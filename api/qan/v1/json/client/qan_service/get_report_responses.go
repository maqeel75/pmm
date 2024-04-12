// Code generated by go-swagger; DO NOT EDIT.

package qan_service

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

// GetReportReader is a Reader for the GetReport structure.
type GetReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetReportDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetReportOK creates a GetReportOK with default headers values
func NewGetReportOK() *GetReportOK {
	return &GetReportOK{}
}

/*
GetReportOK describes a response with status code 200, with default header values.

A successful response.
*/
type GetReportOK struct {
	Payload *GetReportOKBody
}

func (o *GetReportOK) Error() string {
	return fmt.Sprintf("[POST /v1/qan/metrics:getReport][%d] getReportOk  %+v", 200, o.Payload)
}

func (o *GetReportOK) GetPayload() *GetReportOKBody {
	return o.Payload
}

func (o *GetReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetReportOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReportDefault creates a GetReportDefault with default headers values
func NewGetReportDefault(code int) *GetReportDefault {
	return &GetReportDefault{
		_statusCode: code,
	}
}

/*
GetReportDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type GetReportDefault struct {
	_statusCode int

	Payload *GetReportDefaultBody
}

// Code gets the status code for the get report default response
func (o *GetReportDefault) Code() int {
	return o._statusCode
}

func (o *GetReportDefault) Error() string {
	return fmt.Sprintf("[POST /v1/qan/metrics:getReport][%d] GetReport default  %+v", o._statusCode, o.Payload)
}

func (o *GetReportDefault) GetPayload() *GetReportDefaultBody {
	return o.Payload
}

func (o *GetReportDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(GetReportDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetReportBody ReportRequest defines filtering of metrics report for db server or other dimentions.
swagger:model GetReportBody
*/
type GetReportBody struct {
	// period start from
	// Format: date-time
	PeriodStartFrom strfmt.DateTime `json:"period_start_from,omitempty"`

	// period start to
	// Format: date-time
	PeriodStartTo strfmt.DateTime `json:"period_start_to,omitempty"`

	// group by
	GroupBy string `json:"group_by,omitempty"`

	// labels
	Labels []*GetReportParamsBodyLabelsItems0 `json:"labels"`

	// columns
	Columns []string `json:"columns"`

	// order by
	OrderBy string `json:"order_by,omitempty"`

	// offset
	Offset int64 `json:"offset,omitempty"`

	// limit
	Limit int64 `json:"limit,omitempty"`

	// main metric
	MainMetric string `json:"main_metric,omitempty"`

	// search
	Search string `json:"search,omitempty"`
}

// Validate validates this get report body
func (o *GetReportBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePeriodStartFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePeriodStartTo(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetReportBody) validatePeriodStartFrom(formats strfmt.Registry) error {
	if swag.IsZero(o.PeriodStartFrom) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"period_start_from", "body", "date-time", o.PeriodStartFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetReportBody) validatePeriodStartTo(formats strfmt.Registry) error {
	if swag.IsZero(o.PeriodStartTo) { // not required
		return nil
	}

	if err := validate.FormatOf("body"+"."+"period_start_to", "body", "date-time", o.PeriodStartTo.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetReportBody) validateLabels(formats strfmt.Registry) error {
	if swag.IsZero(o.Labels) { // not required
		return nil
	}

	for i := 0; i < len(o.Labels); i++ {
		if swag.IsZero(o.Labels[i]) { // not required
			continue
		}

		if o.Labels[i] != nil {
			if err := o.Labels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get report body based on the context it is used
func (o *GetReportBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLabels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetReportBody) contextValidateLabels(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Labels); i++ {
		if o.Labels[i] != nil {
			if err := o.Labels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "labels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("body" + "." + "labels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetReportBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetReportBody) UnmarshalBinary(b []byte) error {
	var res GetReportBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetReportDefaultBody get report default body
swagger:model GetReportDefaultBody
*/
type GetReportDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*GetReportDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this get report default body
func (o *GetReportDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetReportDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("GetReport default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetReport default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get report default body based on the context it is used
func (o *GetReportDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetReportDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("GetReport default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("GetReport default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetReportDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetReportDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetReportDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetReportDefaultBodyDetailsItems0 get report default body details items0
swagger:model GetReportDefaultBodyDetailsItems0
*/
type GetReportDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this get report default body details items0
func (o *GetReportDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get report default body details items0 based on context it is used
func (o *GetReportDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetReportDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetReportDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res GetReportDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetReportOKBody ReportReply is list of reports per quieryids, hosts etc.
swagger:model GetReportOKBody
*/
type GetReportOKBody struct {
	// total rows
	TotalRows int64 `json:"total_rows,omitempty"`

	// offset
	Offset int64 `json:"offset,omitempty"`

	// limit
	Limit int64 `json:"limit,omitempty"`

	// rows
	Rows []*GetReportOKBodyRowsItems0 `json:"rows"`
}

// Validate validates this get report OK body
func (o *GetReportOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRows(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetReportOKBody) validateRows(formats strfmt.Registry) error {
	if swag.IsZero(o.Rows) { // not required
		return nil
	}

	for i := 0; i < len(o.Rows); i++ {
		if swag.IsZero(o.Rows[i]) { // not required
			continue
		}

		if o.Rows[i] != nil {
			if err := o.Rows[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getReportOk" + "." + "rows" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getReportOk" + "." + "rows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get report OK body based on the context it is used
func (o *GetReportOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRows(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetReportOKBody) contextValidateRows(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Rows); i++ {
		if o.Rows[i] != nil {
			if err := o.Rows[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getReportOk" + "." + "rows" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getReportOk" + "." + "rows" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetReportOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetReportOKBody) UnmarshalBinary(b []byte) error {
	var res GetReportOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetReportOKBodyRowsItems0 Row define metrics for selected dimention.
swagger:model GetReportOKBodyRowsItems0
*/
type GetReportOKBodyRowsItems0 struct {
	// rank
	Rank int64 `json:"rank,omitempty"`

	// dimension
	Dimension string `json:"dimension,omitempty"`

	// database
	Database string `json:"database,omitempty"`

	// metrics
	Metrics map[string]GetReportOKBodyRowsItems0MetricsAnon `json:"metrics,omitempty"`

	// sparkline
	Sparkline []*GetReportOKBodyRowsItems0SparklineItems0 `json:"sparkline"`

	// fingerprint
	Fingerprint string `json:"fingerprint,omitempty"`

	// num queries
	NumQueries int64 `json:"num_queries,omitempty"`

	// qps
	QPS float32 `json:"qps,omitempty"`

	// load
	Load float32 `json:"load,omitempty"`
}

// Validate validates this get report OK body rows items0
func (o *GetReportOKBodyRowsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSparkline(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetReportOKBodyRowsItems0) validateMetrics(formats strfmt.Registry) error {
	if swag.IsZero(o.Metrics) { // not required
		return nil
	}

	for k := range o.Metrics {

		if swag.IsZero(o.Metrics[k]) { // not required
			continue
		}
		if val, ok := o.Metrics[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metrics" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("metrics" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetReportOKBodyRowsItems0) validateSparkline(formats strfmt.Registry) error {
	if swag.IsZero(o.Sparkline) { // not required
		return nil
	}

	for i := 0; i < len(o.Sparkline); i++ {
		if swag.IsZero(o.Sparkline[i]) { // not required
			continue
		}

		if o.Sparkline[i] != nil {
			if err := o.Sparkline[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sparkline" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sparkline" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get report OK body rows items0 based on the context it is used
func (o *GetReportOKBodyRowsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMetrics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSparkline(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetReportOKBodyRowsItems0) contextValidateMetrics(ctx context.Context, formats strfmt.Registry) error {
	for k := range o.Metrics {
		if val, ok := o.Metrics[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}
	}

	return nil
}

func (o *GetReportOKBodyRowsItems0) contextValidateSparkline(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Sparkline); i++ {
		if o.Sparkline[i] != nil {
			if err := o.Sparkline[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sparkline" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sparkline" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetReportOKBodyRowsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetReportOKBodyRowsItems0) UnmarshalBinary(b []byte) error {
	var res GetReportOKBodyRowsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetReportOKBodyRowsItems0MetricsAnon Metric cell.
swagger:model GetReportOKBodyRowsItems0MetricsAnon
*/
type GetReportOKBodyRowsItems0MetricsAnon struct {
	// stats
	Stats *GetReportOKBodyRowsItems0MetricsAnonStats `json:"stats,omitempty"`
}

// Validate validates this get report OK body rows items0 metrics anon
func (o *GetReportOKBodyRowsItems0MetricsAnon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetReportOKBodyRowsItems0MetricsAnon) validateStats(formats strfmt.Registry) error {
	if swag.IsZero(o.Stats) { // not required
		return nil
	}

	if o.Stats != nil {
		if err := o.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get report OK body rows items0 metrics anon based on the context it is used
func (o *GetReportOKBodyRowsItems0MetricsAnon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetReportOKBodyRowsItems0MetricsAnon) contextValidateStats(ctx context.Context, formats strfmt.Registry) error {
	if o.Stats != nil {
		if err := o.Stats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetReportOKBodyRowsItems0MetricsAnon) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetReportOKBodyRowsItems0MetricsAnon) UnmarshalBinary(b []byte) error {
	var res GetReportOKBodyRowsItems0MetricsAnon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetReportOKBodyRowsItems0MetricsAnonStats Stat is statistics of specific metric.
swagger:model GetReportOKBodyRowsItems0MetricsAnonStats
*/
type GetReportOKBodyRowsItems0MetricsAnonStats struct {
	// rate
	Rate float32 `json:"rate,omitempty"`

	// cnt
	Cnt float32 `json:"cnt,omitempty"`

	// sum
	Sum float32 `json:"sum,omitempty"`

	// min
	Min float32 `json:"min,omitempty"`

	// max
	Max float32 `json:"max,omitempty"`

	// p99
	P99 float32 `json:"p99,omitempty"`

	// avg
	Avg float32 `json:"avg,omitempty"`

	// sum per sec
	SumPerSec float32 `json:"sum_per_sec,omitempty"`
}

// Validate validates this get report OK body rows items0 metrics anon stats
func (o *GetReportOKBodyRowsItems0MetricsAnonStats) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get report OK body rows items0 metrics anon stats based on context it is used
func (o *GetReportOKBodyRowsItems0MetricsAnonStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetReportOKBodyRowsItems0MetricsAnonStats) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetReportOKBodyRowsItems0MetricsAnonStats) UnmarshalBinary(b []byte) error {
	var res GetReportOKBodyRowsItems0MetricsAnonStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetReportOKBodyRowsItems0SparklineItems0 Point contains values that represents abscissa (time) and ordinate (volume etc.)
// of every point in a coordinate system of Sparklines.
swagger:model GetReportOKBodyRowsItems0SparklineItems0
*/
type GetReportOKBodyRowsItems0SparklineItems0 struct {
	// The serial number of the chart point from the largest time in the time interval to the lowest time in the time range.
	Point int64 `json:"point,omitempty"`

	// Duration beetween two points.
	TimeFrame int64 `json:"time_frame,omitempty"`

	// Time of point in format RFC3339.
	Timestamp string `json:"timestamp,omitempty"`

	// load is query_time / time_range.
	Load float32 `json:"load,omitempty"`

	// number of queries in bucket.
	NumQueriesPerSec float32 `json:"num_queries_per_sec,omitempty"`

	// number of queries with errors.
	NumQueriesWithErrorsPerSec float32 `json:"num_queries_with_errors_per_sec,omitempty"`

	// number of queries with warnings.
	NumQueriesWithWarningsPerSec float32 `json:"num_queries_with_warnings_per_sec,omitempty"`

	// The statement execution time in seconds.
	MQueryTimeSumPerSec float32 `json:"m_query_time_sum_per_sec,omitempty"`

	// The time to acquire locks in seconds.
	MLockTimeSumPerSec float32 `json:"m_lock_time_sum_per_sec,omitempty"`

	// The number of rows sent to the client.
	MRowsSentSumPerSec float32 `json:"m_rows_sent_sum_per_sec,omitempty"`

	// Number of rows scanned - SELECT.
	MRowsExaminedSumPerSec float32 `json:"m_rows_examined_sum_per_sec,omitempty"`

	// Number of rows changed - UPDATE, DELETE, INSERT.
	MRowsAffectedSumPerSec float32 `json:"m_rows_affected_sum_per_sec,omitempty"`

	// The number of rows read from tables.
	MRowsReadSumPerSec float32 `json:"m_rows_read_sum_per_sec,omitempty"`

	// The number of merge passes that the sort algorithm has had to do.
	MMergePassesSumPerSec float32 `json:"m_merge_passes_sum_per_sec,omitempty"`

	// Counts the number of page read operations scheduled.
	MInnodbIorOpsSumPerSec float32 `json:"m_innodb_io_r_ops_sum_per_sec,omitempty"`

	// Similar to innodb_IO_r_ops, but the unit is bytes.
	MInnodbIorBytesSumPerSec float32 `json:"m_innodb_io_r_bytes_sum_per_sec,omitempty"`

	// Shows how long (in seconds) it took InnoDB to actually read the data from storage.
	MInnodbIorWaitSumPerSec float32 `json:"m_innodb_io_r_wait_sum_per_sec,omitempty"`

	// Shows how long (in seconds) the query waited for row locks.
	MInnodbRecLockWaitSumPerSec float32 `json:"m_innodb_rec_lock_wait_sum_per_sec,omitempty"`

	// Shows how long (in seconds) the query spent either waiting to enter the InnoDB queue or inside that queue waiting for execution.
	MInnodbQueueWaitSumPerSec float32 `json:"m_innodb_queue_wait_sum_per_sec,omitempty"`

	// Counts approximately the number of unique pages the query accessed.
	MInnodbPagesDistinctSumPerSec float32 `json:"m_innodb_pages_distinct_sum_per_sec,omitempty"`

	// Shows how long the query is.
	MQueryLengthSumPerSec float32 `json:"m_query_length_sum_per_sec,omitempty"`

	// The number of bytes sent to all clients.
	MBytesSentSumPerSec float32 `json:"m_bytes_sent_sum_per_sec,omitempty"`

	// Number of temporary tables created on memory for the query.
	MTmpTablesSumPerSec float32 `json:"m_tmp_tables_sum_per_sec,omitempty"`

	// Number of temporary tables created on disk for the query.
	MTmpDiskTablesSumPerSec float32 `json:"m_tmp_disk_tables_sum_per_sec,omitempty"`

	// Total Size in bytes for all temporary tables used in the query.
	MTmpTableSizesSumPerSec float32 `json:"m_tmp_table_sizes_sum_per_sec,omitempty"`

	// Query Cache hits.
	MQcHitSumPerSec float32 `json:"m_qc_hit_sum_per_sec,omitempty"`

	// The query performed a full table scan.
	MFullScanSumPerSec float32 `json:"m_full_scan_sum_per_sec,omitempty"`

	// The query performed a full join (a join without indexes).
	MFullJoinSumPerSec float32 `json:"m_full_join_sum_per_sec,omitempty"`

	// The query created an implicit internal temporary table.
	MTmpTableSumPerSec float32 `json:"m_tmp_table_sum_per_sec,omitempty"`

	// The querys temporary table was stored on disk.
	MTmpTableOnDiskSumPerSec float32 `json:"m_tmp_table_on_disk_sum_per_sec,omitempty"`

	// The query used a filesort.
	MFilesortSumPerSec float32 `json:"m_filesort_sum_per_sec,omitempty"`

	// The filesort was performed on disk.
	MFilesortOnDiskSumPerSec float32 `json:"m_filesort_on_disk_sum_per_sec,omitempty"`

	// The number of joins that used a range search on a reference table.
	MSelectFullRangeJoinSumPerSec float32 `json:"m_select_full_range_join_sum_per_sec,omitempty"`

	// The number of joins that used ranges on the first table.
	MSelectRangeSumPerSec float32 `json:"m_select_range_sum_per_sec,omitempty"`

	// The number of joins without keys that check for key usage after each row.
	MSelectRangeCheckSumPerSec float32 `json:"m_select_range_check_sum_per_sec,omitempty"`

	// The number of sorts that were done using ranges.
	MSortRangeSumPerSec float32 `json:"m_sort_range_sum_per_sec,omitempty"`

	// The number of sorted rows.
	MSortRowsSumPerSec float32 `json:"m_sort_rows_sum_per_sec,omitempty"`

	// The number of sorts that were done by scanning the table.
	MSortScanSumPerSec float32 `json:"m_sort_scan_sum_per_sec,omitempty"`

	// The number of queries without index.
	MNoIndexUsedSumPerSec float32 `json:"m_no_index_used_sum_per_sec,omitempty"`

	// The number of queries without good index.
	MNoGoodIndexUsedSumPerSec float32 `json:"m_no_good_index_used_sum_per_sec,omitempty"`

	// The number of returned documents.
	MDocsReturnedSumPerSec float32 `json:"m_docs_returned_sum_per_sec,omitempty"`

	// The response length of the query result in bytes.
	MResponseLengthSumPerSec float32 `json:"m_response_length_sum_per_sec,omitempty"`

	// The number of scanned documents.
	MDocsScannedSumPerSec float32 `json:"m_docs_scanned_sum_per_sec,omitempty"`

	// Total number of shared block cache hits by the statement.
	MSharedBlksHitSumPerSec float32 `json:"m_shared_blks_hit_sum_per_sec,omitempty"`

	// Total number of shared blocks read by the statement.
	MSharedBlksReadSumPerSec float32 `json:"m_shared_blks_read_sum_per_sec,omitempty"`

	// Total number of shared blocks dirtied by the statement.
	MSharedBlksDirtiedSumPerSec float32 `json:"m_shared_blks_dirtied_sum_per_sec,omitempty"`

	// Total number of shared blocks written by the statement.
	MSharedBlksWrittenSumPerSec float32 `json:"m_shared_blks_written_sum_per_sec,omitempty"`

	// Total number of local block cache hits by the statement.
	MLocalBlksHitSumPerSec float32 `json:"m_local_blks_hit_sum_per_sec,omitempty"`

	// Total number of local blocks read by the statement.
	MLocalBlksReadSumPerSec float32 `json:"m_local_blks_read_sum_per_sec,omitempty"`

	// Total number of local blocks dirtied by the statement.
	MLocalBlksDirtiedSumPerSec float32 `json:"m_local_blks_dirtied_sum_per_sec,omitempty"`

	// Total number of local blocks written by the statement.
	MLocalBlksWrittenSumPerSec float32 `json:"m_local_blks_written_sum_per_sec,omitempty"`

	// Total number of temp blocks read by the statement.
	MTempBlksReadSumPerSec float32 `json:"m_temp_blks_read_sum_per_sec,omitempty"`

	// Total number of temp blocks written by the statement.
	MTempBlksWrittenSumPerSec float32 `json:"m_temp_blks_written_sum_per_sec,omitempty"`

	// Total time the statement spent reading blocks, in milliseconds (if track_io_timing is enabled, otherwise zero).
	MBlkReadTimeSumPerSec float32 `json:"m_blk_read_time_sum_per_sec,omitempty"`

	// Total time the statement spent writing blocks, in milliseconds (if track_io_timing is enabled, otherwise zero).
	MBlkWriteTimeSumPerSec float32 `json:"m_blk_write_time_sum_per_sec,omitempty"`

	// Total time user spent in query.
	MCPUUserTimeSumPerSec float32 `json:"m_cpu_user_time_sum_per_sec,omitempty"`

	// Total time system spent in query.
	MCPUSysTimeSumPerSec float32 `json:"m_cpu_sys_time_sum_per_sec,omitempty"`

	// Total number of planned calls.
	MPlansCallsSumPerSec float32 `json:"m_plans_calls_sum_per_sec,omitempty"`

	// Total number of WAL (Write-ahead logging) records.
	MWalRecordsSumPerSec float32 `json:"m_wal_records_sum_per_sec,omitempty"`

	// Total number of FPI (full page images) in WAL (Write-ahead logging) records.
	MWalFpiSumPerSec float32 `json:"m_wal_fpi_sum_per_sec,omitempty"`

	// Total bytes of WAL (Write-ahead logging) records.
	MWalBytesSumPerSec float32 `json:"m_wal_bytes_sum_per_sec,omitempty"`

	// Plan time in per seconds.
	MPlanTimeSumPerSec float32 `json:"m_plan_time_sum_per_sec,omitempty"`
}

// Validate validates this get report OK body rows items0 sparkline items0
func (o *GetReportOKBodyRowsItems0SparklineItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get report OK body rows items0 sparkline items0 based on context it is used
func (o *GetReportOKBodyRowsItems0SparklineItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetReportOKBodyRowsItems0SparklineItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetReportOKBodyRowsItems0SparklineItems0) UnmarshalBinary(b []byte) error {
	var res GetReportOKBodyRowsItems0SparklineItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetReportParamsBodyLabelsItems0 ReportMapFieldEntry allows to pass labels/dimentions in form like {"server": ["db1", "db2"...]}.
swagger:model GetReportParamsBodyLabelsItems0
*/
type GetReportParamsBodyLabelsItems0 struct {
	// key
	Key string `json:"key,omitempty"`

	// value
	Value []string `json:"value"`
}

// Validate validates this get report params body labels items0
func (o *GetReportParamsBodyLabelsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get report params body labels items0 based on context it is used
func (o *GetReportParamsBodyLabelsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetReportParamsBodyLabelsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetReportParamsBodyLabelsItems0) UnmarshalBinary(b []byte) error {
	var res GetReportParamsBodyLabelsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
