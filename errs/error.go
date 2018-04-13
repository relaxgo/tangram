package errs

import (
	"fmt"
	"strings"
)

type HTTPError interface {
	error
	HTTPStatus() int
	Code() string
	Message() string
}

type Error struct {
	status  int
	code    string
	message string
}

func New(httpCode int, msgList ...string) Error {
	len := len(msgList)
	code := ""
	msg := ""

	switch len {
	case 0:
		code = ""
	case 1:
		code = msgList[0]
	default:
		code = msgList[0]
		msg = strings.Join(msgList[1:], ". ")
	}

	return Error{httpCode, code, msg}
}

func Wrap(err error) Error {
	if e, ok := err.(Error); ok {
		return e
	}
	return New(500, "UNKNOWN_ERROR", err.Error())
}

func (err Error) Error() string   { return err.message }
func (err Error) HTTPStatus() int { return err.status }
func (err Error) Code() string    { return err.code }
func (err Error) Message() string { return err.message }

func (err Error) WithMessage(msg string) Error {
	return Error{
		status:  err.status,
		code:    err.code,
		message: msg,
	}
}

func (err Error) AppendMessage(msg string) Error {
	return Error{
		status:  err.status,
		code:    err.code,
		message: err.message + ". " + msg,
	}
}

func (err Error) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"code\": \"%v\",\"message\": \"%v\"}", err.code, err.message)), nil
}

func (err Error) String() string {
	return fmt.Sprintf("http_status: %v, code: %v, message: %v", err.status, err.code, err.message)
}
