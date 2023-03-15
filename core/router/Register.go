package router

import (
	"net/http"
)

func Get(route string, controllerMethod http.HandlerFunc) *RouteType {
	return register(route, "GET", controllerMethod)
}

func Post(route string, controllerMethod http.HandlerFunc) *RouteType {
	return register(route, "POST", controllerMethod)
}

func Put(route string, controllerMethod http.HandlerFunc) *RouteType {
	return register(route, "PUT", controllerMethod)
}

func Delete(route string, controllerMethod http.HandlerFunc) *RouteType {
	return register(route, "DELETE", controllerMethod)
}

func register(route string, method string, controllerMethod http.HandlerFunc) *RouteType {
	routeHandler := RouteType{
		Url:        route,
		HttpMethod: method,
		Callable:   controllerMethod,
	}

	routes = append(routes, routeHandler)

	return &routeHandler
}
