// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/partyaffairs/resource/partyaffairs_resource_classify.proto

package v1

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

// Validate checks the field values on ListResourceClassifyRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListResourceClassifyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListResourceClassifyRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListResourceClassifyRequestMultiError, or nil if none found.
func (m *ListResourceClassifyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListResourceClassifyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListResourceClassifyRequestMultiError(errors)
	}

	return nil
}

// ListResourceClassifyRequestMultiError is an error wrapping multiple
// validation errors returned by ListResourceClassifyRequest.ValidateAll() if
// the designated constraints aren't met.
type ListResourceClassifyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListResourceClassifyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListResourceClassifyRequestMultiError) AllErrors() []error { return m }

// ListResourceClassifyRequestValidationError is the validation error returned
// by ListResourceClassifyRequest.Validate if the designated constraints
// aren't met.
type ListResourceClassifyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResourceClassifyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResourceClassifyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResourceClassifyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResourceClassifyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResourceClassifyRequestValidationError) ErrorName() string {
	return "ListResourceClassifyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListResourceClassifyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResourceClassifyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResourceClassifyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResourceClassifyRequestValidationError{}

// Validate checks the field values on ListResourceClassifyReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListResourceClassifyReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListResourceClassifyReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListResourceClassifyReplyMultiError, or nil if none found.
func (m *ListResourceClassifyReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListResourceClassifyReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListResourceClassifyReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListResourceClassifyReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListResourceClassifyReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListResourceClassifyReplyMultiError(errors)
	}

	return nil
}

// ListResourceClassifyReplyMultiError is an error wrapping multiple validation
// errors returned by ListResourceClassifyReply.ValidateAll() if the
// designated constraints aren't met.
type ListResourceClassifyReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListResourceClassifyReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListResourceClassifyReplyMultiError) AllErrors() []error { return m }

// ListResourceClassifyReplyValidationError is the validation error returned by
// ListResourceClassifyReply.Validate if the designated constraints aren't met.
type ListResourceClassifyReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResourceClassifyReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResourceClassifyReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResourceClassifyReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResourceClassifyReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResourceClassifyReplyValidationError) ErrorName() string {
	return "ListResourceClassifyReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListResourceClassifyReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResourceClassifyReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResourceClassifyReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResourceClassifyReplyValidationError{}

// Validate checks the field values on CreateResourceClassifyRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateResourceClassifyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateResourceClassifyRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CreateResourceClassifyRequestMultiError, or nil if none found.
func (m *CreateResourceClassifyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateResourceClassifyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := CreateResourceClassifyRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Weight

	if len(errors) > 0 {
		return CreateResourceClassifyRequestMultiError(errors)
	}

	return nil
}

// CreateResourceClassifyRequestMultiError is an error wrapping multiple
// validation errors returned by CreateResourceClassifyRequest.ValidateAll()
// if the designated constraints aren't met.
type CreateResourceClassifyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateResourceClassifyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateResourceClassifyRequestMultiError) AllErrors() []error { return m }

// CreateResourceClassifyRequestValidationError is the validation error
// returned by CreateResourceClassifyRequest.Validate if the designated
// constraints aren't met.
type CreateResourceClassifyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateResourceClassifyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateResourceClassifyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateResourceClassifyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateResourceClassifyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateResourceClassifyRequestValidationError) ErrorName() string {
	return "CreateResourceClassifyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateResourceClassifyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateResourceClassifyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateResourceClassifyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateResourceClassifyRequestValidationError{}

// Validate checks the field values on CreateResourceClassifyReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateResourceClassifyReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateResourceClassifyReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateResourceClassifyReplyMultiError, or nil if none found.
func (m *CreateResourceClassifyReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateResourceClassifyReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateResourceClassifyReplyMultiError(errors)
	}

	return nil
}

// CreateResourceClassifyReplyMultiError is an error wrapping multiple
// validation errors returned by CreateResourceClassifyReply.ValidateAll() if
// the designated constraints aren't met.
type CreateResourceClassifyReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateResourceClassifyReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateResourceClassifyReplyMultiError) AllErrors() []error { return m }

// CreateResourceClassifyReplyValidationError is the validation error returned
// by CreateResourceClassifyReply.Validate if the designated constraints
// aren't met.
type CreateResourceClassifyReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateResourceClassifyReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateResourceClassifyReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateResourceClassifyReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateResourceClassifyReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateResourceClassifyReplyValidationError) ErrorName() string {
	return "CreateResourceClassifyReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CreateResourceClassifyReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateResourceClassifyReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateResourceClassifyReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateResourceClassifyReplyValidationError{}

// Validate checks the field values on UpdateResourceClassifyRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateResourceClassifyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateResourceClassifyRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// UpdateResourceClassifyRequestMultiError, or nil if none found.
func (m *UpdateResourceClassifyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateResourceClassifyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 0 {
		err := UpdateResourceClassifyRequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := UpdateResourceClassifyRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Weight

	if len(errors) > 0 {
		return UpdateResourceClassifyRequestMultiError(errors)
	}

	return nil
}

// UpdateResourceClassifyRequestMultiError is an error wrapping multiple
// validation errors returned by UpdateResourceClassifyRequest.ValidateAll()
// if the designated constraints aren't met.
type UpdateResourceClassifyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateResourceClassifyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateResourceClassifyRequestMultiError) AllErrors() []error { return m }

// UpdateResourceClassifyRequestValidationError is the validation error
// returned by UpdateResourceClassifyRequest.Validate if the designated
// constraints aren't met.
type UpdateResourceClassifyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateResourceClassifyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateResourceClassifyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateResourceClassifyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateResourceClassifyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateResourceClassifyRequestValidationError) ErrorName() string {
	return "UpdateResourceClassifyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateResourceClassifyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateResourceClassifyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateResourceClassifyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateResourceClassifyRequestValidationError{}

// Validate checks the field values on UpdateResourceClassifyReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateResourceClassifyReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateResourceClassifyReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateResourceClassifyReplyMultiError, or nil if none found.
func (m *UpdateResourceClassifyReply) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateResourceClassifyReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateResourceClassifyReplyMultiError(errors)
	}

	return nil
}

// UpdateResourceClassifyReplyMultiError is an error wrapping multiple
// validation errors returned by UpdateResourceClassifyReply.ValidateAll() if
// the designated constraints aren't met.
type UpdateResourceClassifyReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateResourceClassifyReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateResourceClassifyReplyMultiError) AllErrors() []error { return m }

// UpdateResourceClassifyReplyValidationError is the validation error returned
// by UpdateResourceClassifyReply.Validate if the designated constraints
// aren't met.
type UpdateResourceClassifyReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateResourceClassifyReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateResourceClassifyReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateResourceClassifyReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateResourceClassifyReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateResourceClassifyReplyValidationError) ErrorName() string {
	return "UpdateResourceClassifyReplyValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateResourceClassifyReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateResourceClassifyReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateResourceClassifyReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateResourceClassifyReplyValidationError{}

// Validate checks the field values on DeleteResourceClassifyRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteResourceClassifyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteResourceClassifyRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// DeleteResourceClassifyRequestMultiError, or nil if none found.
func (m *DeleteResourceClassifyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteResourceClassifyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() < 0 {
		err := DeleteResourceClassifyRequestValidationError{
			field:  "Id",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteResourceClassifyRequestMultiError(errors)
	}

	return nil
}

// DeleteResourceClassifyRequestMultiError is an error wrapping multiple
// validation errors returned by DeleteResourceClassifyRequest.ValidateAll()
// if the designated constraints aren't met.
type DeleteResourceClassifyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteResourceClassifyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteResourceClassifyRequestMultiError) AllErrors() []error { return m }

// DeleteResourceClassifyRequestValidationError is the validation error
// returned by DeleteResourceClassifyRequest.Validate if the designated
// constraints aren't met.
type DeleteResourceClassifyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteResourceClassifyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteResourceClassifyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteResourceClassifyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteResourceClassifyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteResourceClassifyRequestValidationError) ErrorName() string {
	return "DeleteResourceClassifyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteResourceClassifyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteResourceClassifyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteResourceClassifyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteResourceClassifyRequestValidationError{}

// Validate checks the field values on DeleteResourceClassifyReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteResourceClassifyReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteResourceClassifyReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteResourceClassifyReplyMultiError, or nil if none found.
func (m *DeleteResourceClassifyReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteResourceClassifyReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteResourceClassifyReplyMultiError(errors)
	}

	return nil
}

// DeleteResourceClassifyReplyMultiError is an error wrapping multiple
// validation errors returned by DeleteResourceClassifyReply.ValidateAll() if
// the designated constraints aren't met.
type DeleteResourceClassifyReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteResourceClassifyReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteResourceClassifyReplyMultiError) AllErrors() []error { return m }

// DeleteResourceClassifyReplyValidationError is the validation error returned
// by DeleteResourceClassifyReply.Validate if the designated constraints
// aren't met.
type DeleteResourceClassifyReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteResourceClassifyReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteResourceClassifyReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteResourceClassifyReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteResourceClassifyReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteResourceClassifyReplyValidationError) ErrorName() string {
	return "DeleteResourceClassifyReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteResourceClassifyReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteResourceClassifyReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteResourceClassifyReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteResourceClassifyReplyValidationError{}

// Validate checks the field values on
// ListResourceClassifyReply_ResourceClassify with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListResourceClassifyReply_ResourceClassify) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// ListResourceClassifyReply_ResourceClassify with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// ListResourceClassifyReply_ResourceClassifyMultiError, or nil if none found.
func (m *ListResourceClassifyReply_ResourceClassify) ValidateAll() error {
	return m.validate(true)
}

func (m *ListResourceClassifyReply_ResourceClassify) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Weight

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	if len(errors) > 0 {
		return ListResourceClassifyReply_ResourceClassifyMultiError(errors)
	}

	return nil
}

// ListResourceClassifyReply_ResourceClassifyMultiError is an error wrapping
// multiple validation errors returned by
// ListResourceClassifyReply_ResourceClassify.ValidateAll() if the designated
// constraints aren't met.
type ListResourceClassifyReply_ResourceClassifyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListResourceClassifyReply_ResourceClassifyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListResourceClassifyReply_ResourceClassifyMultiError) AllErrors() []error { return m }

// ListResourceClassifyReply_ResourceClassifyValidationError is the validation
// error returned by ListResourceClassifyReply_ResourceClassify.Validate if
// the designated constraints aren't met.
type ListResourceClassifyReply_ResourceClassifyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResourceClassifyReply_ResourceClassifyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResourceClassifyReply_ResourceClassifyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResourceClassifyReply_ResourceClassifyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResourceClassifyReply_ResourceClassifyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResourceClassifyReply_ResourceClassifyValidationError) ErrorName() string {
	return "ListResourceClassifyReply_ResourceClassifyValidationError"
}

// Error satisfies the builtin error interface
func (e ListResourceClassifyReply_ResourceClassifyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResourceClassifyReply_ResourceClassify.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResourceClassifyReply_ResourceClassifyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResourceClassifyReply_ResourceClassifyValidationError{}
