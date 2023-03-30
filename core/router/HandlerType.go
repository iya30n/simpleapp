package router

import (
	"net/http"
	"simpleapp/core/middleware"
)

type RouteType struct {
	Url             string
	HttpMethod      string
	Callable        http.HandlerFunc
	MiddlewaresList []middleware.MiddlewareContract
}

func (rt *RouteType) Middlewares(middlewares ...middleware.MiddlewareContract) {
	rt.MiddlewaresList = append(rt.MiddlewaresList, middlewares...)
}
