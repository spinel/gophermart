package types

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

// Define the errors we want to use inside the business logic & domain.
var (
	ErrNotFound            = errors.New("resource not found")
	ErrConflict            = errors.New("datamodel conflict")
	ErrForbidden           = errors.New("forbidden access")
	ErrNeedMore            = errors.New("need more input")
	ErrBadRequest          = errors.New("bad request")
	ErrPartialOk           = errors.New("partial okay")
	ErrDuplicateEntry      = errors.New("duplicate entry")
	ErrGone                = errors.New("resource gone")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrNotAllowed          = errors.New("operation not allowed")
	ErrBusy                = errors.New("resource is busy")
	ErrUnauthorized        = errors.New("unauthorized")
)

// HTTPError is our custom HTTP error to get a proper string output.
type HTTPError struct {
	Code    int
	Message string
}

// HTTPCode returns the HTTP code of a given custom HTTP error, with 500 as default.
func HTTPCode(err error) int {
	code := 0
	he, ok := err.(*echo.HTTPError)
	if ok {
		code = he.Code
	}
	return code
}

func ErrWrap(err error) int {
	if err != nil {
		switch {

		case errors.Cause(err) == ErrUnprocessableEntity:
			return http.StatusUnprocessableEntity
		case errors.Cause(err) == ErrDuplicateEntry:
			return http.StatusBadRequest
		default:
			return http.StatusInternalServerError
		}
	}
	return 500
}
