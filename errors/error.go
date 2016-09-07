package errors

import (
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

type HTTPError struct {
	HTTPCode int `json:"-"`
	Code     string
	Message  string
}

const defaultCode = "UNKOWN_ERROR"

var ErrNotFound = New(http.StatusNotFound, "NOT_FOUND", "未找到该内容")

func New(httpCode int, msgList ...string) error {
	len := len(msgList)
	code := ""
	msg := ""

	switch len {
	case 0:
		code = defaultCode
	case 1:
		code = msgList[0]
	default:
		code = msgList[0]
		msg = strings.Join(msgList[1:], ";")
	}

	return &HTTPError{httpCode, code, msg}
}

func (err *HTTPError) Error() string {
	return err.Message
}

func HandleError(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var httpErr *HTTPError
	ok := false

	if httpErr, ok = err.(*HTTPError); ok {
		code = httpErr.HTTPCode
	} else if err, ok := err.(*echo.HTTPError); ok {
		httpErr = &HTTPError{
			HTTPCode: err.Code,
			Code:     defaultCode,
			Message:  err.Message,
		}
	}

	if !c.Response().Committed() {
		if c.Request().Method() == echo.HEAD { // Issue #608
			c.NoContent(code)
		} else {
			c.JSON(code, httpErr)
		}
	}
}
