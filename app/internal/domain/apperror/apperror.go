package apperror

import (
	"fmt"
	"net/http"
)

type Code int

// Application error codes.
const (
	NONE Code = 0

	BadRequest          Code = 1400
	Unauthorized        Code = 1401
	Forbidden           Code = 1403
	NotFound            Code = 1404
	InvalidID           Code = 1405
	InternalServerError Code = 1500
	ServiceUnavailable  Code = 1503

	SystemNotFound      Code = 2001
	UserNotFound        Code = 2002
	DocumentNotFound    Code = 2003
	FileNotFound        Code = 2004
	AuthorityNotFound   Code = 2005
	CertificateNotFound Code = 2006
	OperationNotFound   Code = 2007
)

var httpcodes = map[Code]int{
	BadRequest:          http.StatusBadRequest,
	InvalidID:           http.StatusBadRequest,
	Unauthorized:        http.StatusUnauthorized,
	Forbidden:           http.StatusForbidden,
	NotFound:            http.StatusNotFound,
	InternalServerError: http.StatusInternalServerError,
	ServiceUnavailable:  http.StatusServiceUnavailable,

	SystemNotFound:      http.StatusNotFound,
	UserNotFound:        http.StatusNotFound,
	DocumentNotFound:    http.StatusNotFound,
	FileNotFound:        http.StatusNotFound,
	AuthorityNotFound:   http.StatusForbidden,
	CertificateNotFound: http.StatusNotFound,
	OperationNotFound:   http.StatusNotFound,
}

// Общие ошибки модулей
const (
	ErrGetExtendedTransporter = "Ошибка получения расширенных данных перевозчика"
	ErrNoToken                = "Не указан токен перевозчика"
	ErrBadAPIResponse         = "Ошибка отправки данных в API курьера"
	ErrBadCoords              = "Указаны некорректные координаты поездки"
	ErrBadTransporterType     = "Ошибка типа перевозчика"
	ErrCallTaxi               = "Ошибка вызова такси"
	ErrGetTransporters        = "Ошибка получения списка перевозчиков"
	ErrBadDBTripsUpdate       = "Ошибка обновления поездок в БД"
	ErrGetTripEstimate        = "Ошибка получения стоимости поездок"
	ErrGetTripStatus          = "Ошибка обновления статуса поездок"

	ErrorInvalidID = "Невалидный ID"
)

// ErrorStatusCode returns the associated HTTP status code for a WTF error code.
func HttpStatusCode(code Code) int {
	if v, ok := httpcodes[code]; ok {
		return v
	}

	return http.StatusInternalServerError
}

type Error interface {
	error
	StatusCode() int
	HTTPCode() int
	Message() string
	//Error() string
}
type AppError struct {
	code Code
	msg  string

	cause error
}

//type ValidationError struct {
//	GenericError
//}

func New(code Code, message string) error {
	return &AppError{
		code: code,
		msg:  message,
	}
}

func NewWithCause(code Code, message string, err error) error {
	return &AppError{
		code:  code,
		msg:   message,
		cause: err,
	}
}

func (e *AppError) Unwrap() error {
	return e.cause
}

func (e *AppError) Error() string {
	return fmt.Sprintf("ERROR: Code - %d, Message - %s", e.code, e.Message)
}

func (e *AppError) StatusCode() int {
	return int(e.code)
}

func (e *AppError) HTTPCode() int {
	return HttpStatusCode(e.code)
}

func (e *AppError) Message() string {
	return e.msg
}

//func NewValidationError(code int, msg string) *ValidationError {
//	return &ValidationError{GenericError{
//		Code:    code,
//		Message: msg,
//	}}
//}
