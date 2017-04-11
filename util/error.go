package util

import (
	"bytes"
	"encoding/json"
)

// AppError represents an error
type AppError struct {
	errors []error
}

// NewError returns a pointer of error
func NewError(errors ...error) *AppError {
	return &AppError{errors: errors}
}

// Append appends a new error to the list of errors
//
// This method is chainable which means it returns
// the instance of AppError
func (e *AppError) Append(err error) *AppError {
	e.errors = append(e.errors, err)
	return e
}

// Has returns true if err has been appended
func (e *AppError) Has(err error) bool {
	for _, item := range e.errors {
		if item == err {
			return true
		}
	}

	return false
}

// IsEmpty returns true if no error have been appended
func (e *AppError) IsEmpty() bool {
	return len(e.errors) == 0
}

// Value returns either an instance of AppError or nil
// if IsEmpty() is true
func (e *AppError) Value() *AppError {
	if e.IsEmpty() {
		return nil
	}
	return e
}

// Error returns all errors in a string
func (e *AppError) Error() string {
	b := bytes.Buffer{}

	for _, err := range e.errors {
		b.WriteString(err.Error())
	}

	return b.String()
}

// MarshalJSON transformers the internal []error in
// []string and then converters it to JSON
func (e *AppError) MarshalJSON() ([]byte, error) {
	var slice []string
	for _, item := range e.errors {
		slice = append(slice, item.Error())
	}

	s := struct {
		Errors []string `json:"errors"` // must be exported
	}{
		Errors: slice,
	}

	return json.Marshal(s)
}
