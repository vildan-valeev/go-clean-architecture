package v1

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/vildan-valeev/go-clean-architecture/internal/domain/apperror"
	"net/http"
)

type HTTPError interface {
	error
	StatusCoder
	ErrMessenger
}

type StatusCoder interface {
	StatusCode() int
	HTTPCode() int
}

type ErrMessenger interface {
	Message() string
}

type httpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func newHTTPError(code int, message string) httpError {
	return httpError{
		Code:    code,
		Message: message,
	}
}
func errorResponse(c *fiber.Ctx, err error) error {
	//var pgErr *pgconn.PgError
	//var appErr *apperror.AppError
	fmt.Println(err.Error())
	//switch {
	//case errors.As(err, &appErr):
	//	e, ok := err.(HTTPError)
	//	s := newHTTPError(err)
	//	return c.Status(e.HTTPCode()).JSON(s)
	//case errors.As(err, &pgErr):
	//	fmt.Println("pgErr", err)
	//	return internalServerError(c)
	//default:
	//	return internalServerError(c)
	//}
	var e HTTPError
	ok := errors.As(err, &e)
	if !ok {
		return internalServerError(c)
	}
	httpEr := newHTTPError(e.StatusCode(), e.Message())
	return c.Status(e.HTTPCode()).JSON(httpEr)

}
func internalServerError(c *fiber.Ctx) error {
	return c.Status(http.StatusInternalServerError).
		JSON(newHTTPError(int(apperror.InternalServerError), http.StatusText(http.StatusInternalServerError)))
}
