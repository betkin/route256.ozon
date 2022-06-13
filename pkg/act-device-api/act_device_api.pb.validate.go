// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ozonmp/act_device_api/v1/act_device_api.proto

package act_device_api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Empty) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// EmptyValidationError is the validation error returned by Empty.Validate if
// the designated constraints aren't met.
type EmptyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyValidationError) ErrorName() string { return "EmptyValidationError" }

// Error satisfies the builtin error interface
func (e EmptyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmpty.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyValidationError{}

// Validate checks the field values on Device with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Device) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Platform

	// no validation rules for UserId

	if v, ok := interface{}(m.GetEnteredAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeviceValidationError{
				field:  "EnteredAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DeviceValidationError is the validation error returned by Device.Validate if
// the designated constraints aren't met.
type DeviceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeviceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeviceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeviceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeviceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeviceValidationError) ErrorName() string { return "DeviceValidationError" }

// Error satisfies the builtin error interface
func (e DeviceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDevice.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeviceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeviceValidationError{}

// Validate checks the field values on CreateDeviceV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateDeviceV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetPlatform()) < 1 {
		return CreateDeviceV1RequestValidationError{
			field:  "Platform",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetUserId() <= 0 {
		return CreateDeviceV1RequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// CreateDeviceV1RequestValidationError is the validation error returned by
// CreateDeviceV1Request.Validate if the designated constraints aren't met.
type CreateDeviceV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateDeviceV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateDeviceV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateDeviceV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateDeviceV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateDeviceV1RequestValidationError) ErrorName() string {
	return "CreateDeviceV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateDeviceV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateDeviceV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateDeviceV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateDeviceV1RequestValidationError{}

// Validate checks the field values on CreateDeviceV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateDeviceV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for DeviceId

	return nil
}

// CreateDeviceV1ResponseValidationError is the validation error returned by
// CreateDeviceV1Response.Validate if the designated constraints aren't met.
type CreateDeviceV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateDeviceV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateDeviceV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateDeviceV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateDeviceV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateDeviceV1ResponseValidationError) ErrorName() string {
	return "CreateDeviceV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateDeviceV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateDeviceV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateDeviceV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateDeviceV1ResponseValidationError{}

// Validate checks the field values on DescribeDeviceV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeDeviceV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetDeviceId() <= 0 {
		return DescribeDeviceV1RequestValidationError{
			field:  "DeviceId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribeDeviceV1RequestValidationError is the validation error returned by
// DescribeDeviceV1Request.Validate if the designated constraints aren't met.
type DescribeDeviceV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeDeviceV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeDeviceV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeDeviceV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeDeviceV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeDeviceV1RequestValidationError) ErrorName() string {
	return "DescribeDeviceV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeDeviceV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeDeviceV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeDeviceV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeDeviceV1RequestValidationError{}

// Validate checks the field values on DescribeDeviceV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeDeviceV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeDeviceV1ResponseValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeDeviceV1ResponseValidationError is the validation error returned by
// DescribeDeviceV1Response.Validate if the designated constraints aren't met.
type DescribeDeviceV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeDeviceV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeDeviceV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeDeviceV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeDeviceV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeDeviceV1ResponseValidationError) ErrorName() string {
	return "DescribeDeviceV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeDeviceV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeDeviceV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeDeviceV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeDeviceV1ResponseValidationError{}

// Validate checks the field values on LogDeviceV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LogDeviceV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetDeviceId() <= 0 {
		return LogDeviceV1RequestValidationError{
			field:  "DeviceId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// LogDeviceV1RequestValidationError is the validation error returned by
// LogDeviceV1Request.Validate if the designated constraints aren't met.
type LogDeviceV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LogDeviceV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LogDeviceV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LogDeviceV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LogDeviceV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LogDeviceV1RequestValidationError) ErrorName() string {
	return "LogDeviceV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e LogDeviceV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogDeviceV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LogDeviceV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LogDeviceV1RequestValidationError{}

// Validate checks the field values on LogDeviceV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LogDeviceV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LogDeviceV1ResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// LogDeviceV1ResponseValidationError is the validation error returned by
// LogDeviceV1Response.Validate if the designated constraints aren't met.
type LogDeviceV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LogDeviceV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LogDeviceV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LogDeviceV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LogDeviceV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LogDeviceV1ResponseValidationError) ErrorName() string {
	return "LogDeviceV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e LogDeviceV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogDeviceV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LogDeviceV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LogDeviceV1ResponseValidationError{}

// Validate checks the field values on ListDevicesV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListDevicesV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Page

	// no validation rules for PerPage

	return nil
}

// ListDevicesV1RequestValidationError is the validation error returned by
// ListDevicesV1Request.Validate if the designated constraints aren't met.
type ListDevicesV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListDevicesV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListDevicesV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListDevicesV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListDevicesV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListDevicesV1RequestValidationError) ErrorName() string {
	return "ListDevicesV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListDevicesV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListDevicesV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListDevicesV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListDevicesV1RequestValidationError{}

// Validate checks the field values on ListDevicesV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListDevicesV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListDevicesV1ResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListDevicesV1ResponseValidationError is the validation error returned by
// ListDevicesV1Response.Validate if the designated constraints aren't met.
type ListDevicesV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListDevicesV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListDevicesV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListDevicesV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListDevicesV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListDevicesV1ResponseValidationError) ErrorName() string {
	return "ListDevicesV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListDevicesV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListDevicesV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListDevicesV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListDevicesV1ResponseValidationError{}

// Validate checks the field values on UpdateDeviceV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateDeviceV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetDeviceId() <= 0 {
		return UpdateDeviceV1RequestValidationError{
			field:  "DeviceId",
			reason: "value must be greater than 0",
		}
	}

	if utf8.RuneCountInString(m.GetPlatform()) < 1 {
		return UpdateDeviceV1RequestValidationError{
			field:  "Platform",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetUserId() <= 0 {
		return UpdateDeviceV1RequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// UpdateDeviceV1RequestValidationError is the validation error returned by
// UpdateDeviceV1Request.Validate if the designated constraints aren't met.
type UpdateDeviceV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateDeviceV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateDeviceV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateDeviceV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateDeviceV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateDeviceV1RequestValidationError) ErrorName() string {
	return "UpdateDeviceV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateDeviceV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateDeviceV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateDeviceV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateDeviceV1RequestValidationError{}

// Validate checks the field values on UpdateLastDeviceV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateLastDeviceV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetPlatform()) < 1 {
		return UpdateLastDeviceV1RequestValidationError{
			field:  "Platform",
			reason: "value length must be at least 1 runes",
		}
	}

	if m.GetUserId() <= 0 {
		return UpdateLastDeviceV1RequestValidationError{
			field:  "UserId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// UpdateLastDeviceV1RequestValidationError is the validation error returned by
// UpdateLastDeviceV1Request.Validate if the designated constraints aren't met.
type UpdateLastDeviceV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateLastDeviceV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateLastDeviceV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateLastDeviceV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateLastDeviceV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateLastDeviceV1RequestValidationError) ErrorName() string {
	return "UpdateLastDeviceV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateLastDeviceV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateLastDeviceV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateLastDeviceV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateLastDeviceV1RequestValidationError{}

// Validate checks the field values on UpdateDeviceV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateDeviceV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Success

	return nil
}

// UpdateDeviceV1ResponseValidationError is the validation error returned by
// UpdateDeviceV1Response.Validate if the designated constraints aren't met.
type UpdateDeviceV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateDeviceV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateDeviceV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateDeviceV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateDeviceV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateDeviceV1ResponseValidationError) ErrorName() string {
	return "UpdateDeviceV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateDeviceV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateDeviceV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateDeviceV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateDeviceV1ResponseValidationError{}

// Validate checks the field values on RemoveDeviceV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveDeviceV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetDeviceId() <= 0 {
		return RemoveDeviceV1RequestValidationError{
			field:  "DeviceId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// RemoveDeviceV1RequestValidationError is the validation error returned by
// RemoveDeviceV1Request.Validate if the designated constraints aren't met.
type RemoveDeviceV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveDeviceV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveDeviceV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveDeviceV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveDeviceV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveDeviceV1RequestValidationError) ErrorName() string {
	return "RemoveDeviceV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveDeviceV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveDeviceV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveDeviceV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveDeviceV1RequestValidationError{}

// Validate checks the field values on RemoveDeviceV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveDeviceV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Found

	return nil
}

// RemoveDeviceV1ResponseValidationError is the validation error returned by
// RemoveDeviceV1Response.Validate if the designated constraints aren't met.
type RemoveDeviceV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveDeviceV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveDeviceV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveDeviceV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveDeviceV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveDeviceV1ResponseValidationError) ErrorName() string {
	return "RemoveDeviceV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveDeviceV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveDeviceV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveDeviceV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveDeviceV1ResponseValidationError{}

// Validate checks the field values on DeviceEvent with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DeviceEvent) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 0 {
		return DeviceEventValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
	}

	if m.GetDeviceId() <= 0 {
		return DeviceEventValidationError{
			field:  "DeviceId",
			reason: "value must be greater than 0",
		}
	}

	if m.GetType() <= 0 {
		return DeviceEventValidationError{
			field:  "Type",
			reason: "value must be greater than 0",
		}
	}

	if m.GetStatus() <= 0 {
		return DeviceEventValidationError{
			field:  "Status",
			reason: "value must be greater than 0",
		}
	}

	if v, ok := interface{}(m.GetPayload()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeviceEventValidationError{
				field:  "Payload",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DeviceEventValidationError is the validation error returned by
// DeviceEvent.Validate if the designated constraints aren't met.
type DeviceEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeviceEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeviceEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeviceEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeviceEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeviceEventValidationError) ErrorName() string { return "DeviceEventValidationError" }

// Error satisfies the builtin error interface
func (e DeviceEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeviceEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeviceEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeviceEventValidationError{}

// Validate checks the field values on DeviceLog with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *DeviceLog) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 0 {
		return DeviceLogValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
	}

	if m.GetType() <= 0 {
		return DeviceLogValidationError{
			field:  "Type",
			reason: "value must be greater than 0",
		}
	}

	if m.GetStatus() <= 0 {
		return DeviceLogValidationError{
			field:  "Status",
			reason: "value must be greater than 0",
		}
	}

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeviceLogValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeviceLogValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DeviceLogValidationError is the validation error returned by
// DeviceLog.Validate if the designated constraints aren't met.
type DeviceLogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeviceLogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeviceLogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeviceLogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeviceLogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeviceLogValidationError) ErrorName() string { return "DeviceLogValidationError" }

// Error satisfies the builtin error interface
func (e DeviceLogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeviceLog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeviceLogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeviceLogValidationError{}
