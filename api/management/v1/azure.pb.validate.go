// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: management/v1/azure.proto

package managementv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on DiscoverAzureDatabaseRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DiscoverAzureDatabaseRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DiscoverAzureDatabaseRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DiscoverAzureDatabaseRequestMultiError, or nil if none found.
func (m *DiscoverAzureDatabaseRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DiscoverAzureDatabaseRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetAzureClientId()) < 1 {
		err := DiscoverAzureDatabaseRequestValidationError{
			field:  "AzureClientId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetAzureClientSecret()) < 1 {
		err := DiscoverAzureDatabaseRequestValidationError{
			field:  "AzureClientSecret",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetAzureTenantId()) < 1 {
		err := DiscoverAzureDatabaseRequestValidationError{
			field:  "AzureTenantId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetAzureSubscriptionId()) < 1 {
		err := DiscoverAzureDatabaseRequestValidationError{
			field:  "AzureSubscriptionId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DiscoverAzureDatabaseRequestMultiError(errors)
	}

	return nil
}

// DiscoverAzureDatabaseRequestMultiError is an error wrapping multiple
// validation errors returned by DiscoverAzureDatabaseRequest.ValidateAll() if
// the designated constraints aren't met.
type DiscoverAzureDatabaseRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DiscoverAzureDatabaseRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DiscoverAzureDatabaseRequestMultiError) AllErrors() []error { return m }

// DiscoverAzureDatabaseRequestValidationError is the validation error returned
// by DiscoverAzureDatabaseRequest.Validate if the designated constraints
// aren't met.
type DiscoverAzureDatabaseRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DiscoverAzureDatabaseRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DiscoverAzureDatabaseRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DiscoverAzureDatabaseRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DiscoverAzureDatabaseRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DiscoverAzureDatabaseRequestValidationError) ErrorName() string {
	return "DiscoverAzureDatabaseRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DiscoverAzureDatabaseRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDiscoverAzureDatabaseRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DiscoverAzureDatabaseRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DiscoverAzureDatabaseRequestValidationError{}

// Validate checks the field values on DiscoverAzureDatabaseInstance with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DiscoverAzureDatabaseInstance) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DiscoverAzureDatabaseInstance with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// DiscoverAzureDatabaseInstanceMultiError, or nil if none found.
func (m *DiscoverAzureDatabaseInstance) ValidateAll() error {
	return m.validate(true)
}

func (m *DiscoverAzureDatabaseInstance) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for InstanceId

	// no validation rules for Region

	// no validation rules for ServiceName

	// no validation rules for Username

	// no validation rules for Address

	// no validation rules for AzureResourceGroup

	// no validation rules for Environment

	// no validation rules for Type

	// no validation rules for Az

	// no validation rules for NodeModel

	if len(errors) > 0 {
		return DiscoverAzureDatabaseInstanceMultiError(errors)
	}

	return nil
}

// DiscoverAzureDatabaseInstanceMultiError is an error wrapping multiple
// validation errors returned by DiscoverAzureDatabaseInstance.ValidateAll()
// if the designated constraints aren't met.
type DiscoverAzureDatabaseInstanceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DiscoverAzureDatabaseInstanceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DiscoverAzureDatabaseInstanceMultiError) AllErrors() []error { return m }

// DiscoverAzureDatabaseInstanceValidationError is the validation error
// returned by DiscoverAzureDatabaseInstance.Validate if the designated
// constraints aren't met.
type DiscoverAzureDatabaseInstanceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DiscoverAzureDatabaseInstanceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DiscoverAzureDatabaseInstanceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DiscoverAzureDatabaseInstanceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DiscoverAzureDatabaseInstanceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DiscoverAzureDatabaseInstanceValidationError) ErrorName() string {
	return "DiscoverAzureDatabaseInstanceValidationError"
}

// Error satisfies the builtin error interface
func (e DiscoverAzureDatabaseInstanceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDiscoverAzureDatabaseInstance.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DiscoverAzureDatabaseInstanceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DiscoverAzureDatabaseInstanceValidationError{}

// Validate checks the field values on DiscoverAzureDatabaseResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DiscoverAzureDatabaseResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DiscoverAzureDatabaseResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// DiscoverAzureDatabaseResponseMultiError, or nil if none found.
func (m *DiscoverAzureDatabaseResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DiscoverAzureDatabaseResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetAzureDatabaseInstance() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DiscoverAzureDatabaseResponseValidationError{
						field:  fmt.Sprintf("AzureDatabaseInstance[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DiscoverAzureDatabaseResponseValidationError{
						field:  fmt.Sprintf("AzureDatabaseInstance[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DiscoverAzureDatabaseResponseValidationError{
					field:  fmt.Sprintf("AzureDatabaseInstance[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return DiscoverAzureDatabaseResponseMultiError(errors)
	}

	return nil
}

// DiscoverAzureDatabaseResponseMultiError is an error wrapping multiple
// validation errors returned by DiscoverAzureDatabaseResponse.ValidateAll()
// if the designated constraints aren't met.
type DiscoverAzureDatabaseResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DiscoverAzureDatabaseResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DiscoverAzureDatabaseResponseMultiError) AllErrors() []error { return m }

// DiscoverAzureDatabaseResponseValidationError is the validation error
// returned by DiscoverAzureDatabaseResponse.Validate if the designated
// constraints aren't met.
type DiscoverAzureDatabaseResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DiscoverAzureDatabaseResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DiscoverAzureDatabaseResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DiscoverAzureDatabaseResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DiscoverAzureDatabaseResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DiscoverAzureDatabaseResponseValidationError) ErrorName() string {
	return "DiscoverAzureDatabaseResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DiscoverAzureDatabaseResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDiscoverAzureDatabaseResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DiscoverAzureDatabaseResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DiscoverAzureDatabaseResponseValidationError{}

// Validate checks the field values on AddAzureDatabaseRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddAzureDatabaseRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddAzureDatabaseRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddAzureDatabaseRequestMultiError, or nil if none found.
func (m *AddAzureDatabaseRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddAzureDatabaseRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetRegion()) < 1 {
		err := AddAzureDatabaseRequestValidationError{
			field:  "Region",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Az

	if utf8.RuneCountInString(m.GetInstanceId()) < 1 {
		err := AddAzureDatabaseRequestValidationError{
			field:  "InstanceId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for NodeModel

	if utf8.RuneCountInString(m.GetAddress()) < 1 {
		err := AddAzureDatabaseRequestValidationError{
			field:  "Address",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPort() <= 0 {
		err := AddAzureDatabaseRequestValidationError{
			field:  "Port",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for NodeName

	// no validation rules for ServiceName

	// no validation rules for Environment

	if utf8.RuneCountInString(m.GetUsername()) < 1 {
		err := AddAzureDatabaseRequestValidationError{
			field:  "Username",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Password

	if utf8.RuneCountInString(m.GetAzureClientId()) < 1 {
		err := AddAzureDatabaseRequestValidationError{
			field:  "AzureClientId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetAzureClientSecret()) < 1 {
		err := AddAzureDatabaseRequestValidationError{
			field:  "AzureClientSecret",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetAzureTenantId()) < 1 {
		err := AddAzureDatabaseRequestValidationError{
			field:  "AzureTenantId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetAzureSubscriptionId()) < 1 {
		err := AddAzureDatabaseRequestValidationError{
			field:  "AzureSubscriptionId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetAzureResourceGroup()) < 1 {
		err := AddAzureDatabaseRequestValidationError{
			field:  "AzureResourceGroup",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for AzureDatabaseExporter

	// no validation rules for Qan

	// no validation rules for CustomLabels

	// no validation rules for SkipConnectionCheck

	// no validation rules for Tls

	// no validation rules for TlsSkipVerify

	// no validation rules for DisableQueryExamples

	// no validation rules for TablestatsGroupTableLimit

	// no validation rules for Type

	if len(errors) > 0 {
		return AddAzureDatabaseRequestMultiError(errors)
	}

	return nil
}

// AddAzureDatabaseRequestMultiError is an error wrapping multiple validation
// errors returned by AddAzureDatabaseRequest.ValidateAll() if the designated
// constraints aren't met.
type AddAzureDatabaseRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddAzureDatabaseRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddAzureDatabaseRequestMultiError) AllErrors() []error { return m }

// AddAzureDatabaseRequestValidationError is the validation error returned by
// AddAzureDatabaseRequest.Validate if the designated constraints aren't met.
type AddAzureDatabaseRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddAzureDatabaseRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddAzureDatabaseRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddAzureDatabaseRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddAzureDatabaseRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddAzureDatabaseRequestValidationError) ErrorName() string {
	return "AddAzureDatabaseRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddAzureDatabaseRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddAzureDatabaseRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddAzureDatabaseRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddAzureDatabaseRequestValidationError{}

// Validate checks the field values on AddAzureDatabaseResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddAzureDatabaseResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddAzureDatabaseResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddAzureDatabaseResponseMultiError, or nil if none found.
func (m *AddAzureDatabaseResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AddAzureDatabaseResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AddAzureDatabaseResponseMultiError(errors)
	}

	return nil
}

// AddAzureDatabaseResponseMultiError is an error wrapping multiple validation
// errors returned by AddAzureDatabaseResponse.ValidateAll() if the designated
// constraints aren't met.
type AddAzureDatabaseResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddAzureDatabaseResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddAzureDatabaseResponseMultiError) AllErrors() []error { return m }

// AddAzureDatabaseResponseValidationError is the validation error returned by
// AddAzureDatabaseResponse.Validate if the designated constraints aren't met.
type AddAzureDatabaseResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddAzureDatabaseResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddAzureDatabaseResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddAzureDatabaseResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddAzureDatabaseResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddAzureDatabaseResponseValidationError) ErrorName() string {
	return "AddAzureDatabaseResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AddAzureDatabaseResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddAzureDatabaseResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddAzureDatabaseResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddAzureDatabaseResponseValidationError{}