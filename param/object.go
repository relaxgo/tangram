package parameters

import (
	"github.com/labstack/echo"
)

func Bind(c echo.Context, obj interface{}) error {
	if err := c.Bind(obj); err != nil {
		return err
	}
	return nil
}
