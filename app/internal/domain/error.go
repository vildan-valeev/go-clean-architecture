package domain

import "fmt"

type DomainError interface {
	StatusCode() int
	ErrorMessage() string
	Error() string
}
type GenericError struct {
	Code    int
	Message string
}

type ValidationError struct {
	GenericError
}

func (e *GenericError) Error() string {
	return fmt.Sprintf("ERROR: Code - %d, Message - %s", e.Code, e.Message)
}

func (e *GenericError) StatusCode() int {
	return e.Code
}

func (e *GenericError) ErrorMessage() string {
	return e.Message
}

func NewValidationError(code int, msg string) *ValidationError {
	return &ValidationError{GenericError{
		Code:    code,
		Message: msg,
	}}
}

var ServiceError = NewValidationError(1555, "something should have a specific value and cannot be empty")
