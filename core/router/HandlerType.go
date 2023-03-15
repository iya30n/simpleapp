package router

import "net/http"

type RouteType struct {
	Url             string
	HttpMethod      string
	Callable        http.HandlerFunc
	MiddlewaresList []string
}

func (rh *RouteType) Middlewares(middlewares ...string) {
	rh.MiddlewaresList = middlewares
}
