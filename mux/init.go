package mux

import (
	"github.com/fasthttp/session/v2"
	"github.com/valyala/fasthttp"
)

type Controller struct {
	Method  string
	Querys  map[string]string
	Handler func(*fasthttp.RequestCtx, *session.Session)
}

type Route struct {
	Ctrl Controller
	Sons map[string]Route
}
