package errors

import (
	"github.com/labstack/echo"
)

func New(code int, msg ...string) error {
	return echo.NewHTTPError(code, msg...)
}
