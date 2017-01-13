package parameters

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func ParamIds(c echo.Context, params []string, isStrict bool) (ids map[string]int, err error) {
	ids = make(map[string]int)
	for _, p := range params {
		pVal := c.Param(p)

		val, err := strconv.Atoi(pVal)
		if err != nil && isStrict {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "param "+p+" is not valid")
		}
		ids[p] = val
	}
	return
}

func AtoI(s string, def int) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		return def
	}
	return val
}

func AtoIWithRange(s string, def, min, max int) int {
	v := AtoI(s, def)
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
