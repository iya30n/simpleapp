package router

import (
	"net/http"
)

func Get(route string, controllerMethod http.HandlerFunc) *RouteType {
	routeHandler := RouteType{
		Url:        route,
		HttpMethod: "GET",
		Callable:   controllerMethod,
	}

	routes = append(routes, routeHandler)

	return &routeHandler
}

func Post(route string, controllerMethod http.HandlerFunc) *RouteType {
	routeHandler := RouteType{
		Url:        route,
		HttpMethod: "POST",
		Callable:   controllerMethod,
	}

	routes = append(routes, routeHandler)

	return &routeHandler
}
