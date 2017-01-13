package errors

import (
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
var debug = false

func SetDebug(isDebug bool) {
	debug = isDebug
}

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
