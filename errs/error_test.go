package errs

import (
	"encoding/json"
	"testing"
)

func TestNew(t *testing.T) {
	err := New(200)
	Assert(t, err, 200, "", "")

	err = New(404, "NOT_FOUND")
	Assert(t, err, 404, "NOT_FOUND", "")

	err = New(404, "NOT_FOUND", "the resource is not found")
	Assert(t, err, 404, "NOT_FOUND", "the resource is not found")

	err = New(404, "NOT_FOUND", "the resource is not found", "please check the url")
	Assert(t, err, 404, "NOT_FOUND", "the resource is not found. please check the url")
}

func TestWithMessage(t *testing.T) {
	err := New(404, "NOT_FOUND", "the resource is not found", "please check the url")
	err2 := err.WithMessage("can't find the resource")
	Assert(t, err, 404, "NOT_FOUND", "the resource is not found. please check the url")
	Assert(t, err2, 404, "NOT_FOUND", "can't find the resource")
}

func TestAppendMessage(t *testing.T) {
	err := New(404, "NOT_FOUND", "the resource is not found")
	err2 := err.AppendMessage("please check the url")
	Assert(t, err, 404, "NOT_FOUND", "the resource is not found")
	Assert(t, err2, 404, "NOT_FOUND", "the resource is not found. please check the url")
}

func Assert(t *testing.T, err HTTPError, status int, code, msg string) {
	if err.HTTPStatus() != status {
		t.Error("err is not correct", err)
	}
	if err.Code() != code {
		t.Error("err is not correct", err)
	}
	if err.Message() != msg {
		t.Error("err is not correct", err)
	}
}

func TestJSON(t *testing.T) {
	err := New(404, "NOT_FOUND", "the resource is not found")
	v, e := json.Marshal(err)
	if e != nil {
		t.Error(e)
	} else {
		t.Log(string(v))
	}
}
