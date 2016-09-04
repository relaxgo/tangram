package errors

import (
	"github.com/labstack/echo"
)

type CommonError struct {
	*echo.HTTPError
}

func New(code int, msg ...string) error {
	return CommonError{HTTPError: echo.NewHTTPError(code, msg...)}
}
