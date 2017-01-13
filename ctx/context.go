package context

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine"
)

type Context interface {
	EchoContent() echo.Context
	QueryParam(string) string
	Param(string) string
	Get(string) interface{}
	Set(string, interface{})

	Cookie(string) (engine.Cookie, error)
	SetCookie(engine.Cookie)
}

type bundleContext struct {
	echo.Context
}

func (context bundleContext) EchoContent() echo.Context {
	return context.Context
}

func New(c echo.Context) Context {
	return bundleContext{c}
}
