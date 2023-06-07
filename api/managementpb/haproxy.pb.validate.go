// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: managementpb/haproxy.proto

package managementpb

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

// Validate checks the field values on AddHAProxyRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AddHAProxyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddHAProxyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddHAProxyRequestMultiError, or nil if none found.
func (m *AddHAProxyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AddHAProxyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NodeId

	// no validation rules for NodeName

	if all {
		switch v := interface{}(m.GetAddNode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddHAProxyRequestValidationError{
					field:  "AddNode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddHAProxyRequestValidationError{
					field:  "AddNode",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAddNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddHAProxyRequestValidationError{
				field:  "AddNode",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Address

	if utf8.RuneCountInString(m.GetServiceName()) < 1 {
		err := AddHAProxyRequestValidationError{
			field:  "ServiceName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for Scheme

	// no validation rules for MetricsPath

	if val := m.GetListenPort(); val <= 0 || val >= 65536 {
		err := AddHAProxyRequestValidationError{
			field:  "ListenPort",
			reason: "value must be inside range (0, 65536)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Environment

	// no validation rules for Cluster

	// no validation rules for ReplicationSet

	// no validation rules for CustomLabels

	// no validation rules for MetricsMode

	// no validation rules for SkipConnectionCheck

	if len(errors) > 0 {
		return AddHAProxyRequestMultiError(errors)
	}

	return nil
}

// AddHAProxyRequestMultiError is an error wrapping multiple validation errors
// returned by AddHAProxyRequest.ValidateAll() if the designated constraints
// aren't met.
type AddHAProxyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddHAProxyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddHAProxyRequestMultiError) AllErrors() []error { return m }

// AddHAProxyRequestValidationError is the validation error returned by
// AddHAProxyRequest.Validate if the designated constraints aren't met.
type AddHAProxyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddHAProxyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddHAProxyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddHAProxyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddHAProxyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddHAProxyRequestValidationError) ErrorName() string {
	return "AddHAProxyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AddHAProxyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddHAProxyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddHAProxyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddHAProxyRequestValidationError{}

// Validate checks the field values on AddHAProxyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AddHAProxyResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddHAProxyResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AddHAProxyResponseMultiError, or nil if none found.
func (m *AddHAProxyResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AddHAProxyResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetService()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddHAProxyResponseValidationError{
					field:  "Service",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddHAProxyResponseValidationError{
					field:  "Service",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddHAProxyResponseValidationError{
				field:  "Service",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetExternalExporter()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AddHAProxyResponseValidationError{
					field:  "ExternalExporter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AddHAProxyResponseValidationError{
					field:  "ExternalExporter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExternalExporter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AddHAProxyResponseValidationError{
				field:  "ExternalExporter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AddHAProxyResponseMultiError(errors)
	}

	return nil
}

// AddHAProxyResponseMultiError is an error wrapping multiple validation errors
// returned by AddHAProxyResponse.ValidateAll() if the designated constraints
// aren't met.
type AddHAProxyResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddHAProxyResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddHAProxyResponseMultiError) AllErrors() []error { return m }

// AddHAProxyResponseValidationError is the validation error returned by
// AddHAProxyResponse.Validate if the designated constraints aren't met.
type AddHAProxyResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddHAProxyResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddHAProxyResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddHAProxyResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddHAProxyResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddHAProxyResponseValidationError) ErrorName() string {
	return "AddHAProxyResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AddHAProxyResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddHAProxyResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddHAProxyResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddHAProxyResponseValidationError{}
