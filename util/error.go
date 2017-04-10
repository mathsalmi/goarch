package util

import "bytes"

// AppError represents an error
type AppError struct {
	errors []error
}

// NewError returns a pointer of error
func NewError() *AppError {
	return &AppError{}
}

// Append appends a new error to the list of errors
func (e *AppError) Append(err error) {
	e.errors = append(e.errors, err)
}

// Error returns all errors in string
func (e *AppError) Error() string {
	b := bytes.Buffer{}

	for _, err := range e.errors {
		b.WriteString(err.Error())
	}

	return b.String()
}
