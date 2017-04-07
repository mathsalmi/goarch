package util

import "bytes"

type AppError struct {
	errors []error
}

func NewError() *AppError {
	return &AppError{}
}

func (e *AppError) Append(err error) {
	e.errors = append(e.errors, err)
}

func (e *AppError) Error() string {
	b := bytes.Buffer{}

	for _, err := range e.errors {
		b.WriteString(err.Error())
	}

	return b.String()
}
