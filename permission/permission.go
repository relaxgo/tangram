package permission

import (
	"log"

	"github.com/labstack/echo"
)

type Recognizer func(*echo.Context, string, string) bool

var permissions = make(map[string][]string)
var recognizers = make(map[string]Recognizer)

var isStrict = false

func Allow(role, method string) {
	recognizer := recognizers[role]
	if recognizer == nil {
		log.Fatalln("there is no recognizer for ", role)
	}

	roles := permissions[method]
	if roles == nil {
		permissions[method] = []string{role}
	}

	for _, item := range roles {
		if item == role {
			return
		}
	}

	permissions[method] = append(roles, role)
}

func IsAllow(c *echo.Context, model, method string) bool {
	roles := permissions[method]
	if roles == nil {
		return !isStrict
	}

	for _, role := range roles {
		recognizer := recognizers[role]
		if recognizer(c, model, method) {
			return true
		}
	}

	return false
}

func AddRole(role string, recognizer Recognizer) {
	rcg := recognizers[role]
	if rcg != nil {
		log.Fatalln("recognizer for ", role, " is exist")
	}

	recognizers[role] = recognizer
}