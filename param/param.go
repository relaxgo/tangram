package param

import (
	"strconv"
	"strings"
)

type PostValues interface {
	PostFormValue(string) string
}

type Values interface {
	Param(string) string
	QueryParam(string) string
	PostValues
}

func String(vs Values, key string) string {
	v := vs.Param(key)
	if v != "" {
		return v
	}
	v = vs.QueryParam(key)
	if v != "" {
		return v
	}
	v = vs.PostFormValue(key)
	if v != "" {
		return v
	}
	return ""
}

func Int(vs Values, key string) int {
	v, _ := strconv.Atoi(String(vs, key))
	return v
}

func Bool(vs Values, key string) bool {
	v := String(vs, key)
	return strings.ToLower(v) == "true"
}

func Object(p PostValues, v interface{}) {
}
