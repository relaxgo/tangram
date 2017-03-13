package errors

import "strings"

type HTTPError interface {
	error
	HTTPStatus() int
	Code() string
	Message() string
}

type Error struct {
	status  int `json:"-"`
	code    string
	message string
	detail  string
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
