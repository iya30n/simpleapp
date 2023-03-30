package router

import (
	"net/http"
	responsehandler "simpleapp/core/responseHandler"
)

func Serve(port string) {
	for _, route := range routes {
		http.HandleFunc(route.Url, makeHandler(route))
	}

	http.ListenAndServe(":"+port, nil)
}

func checkHttpMethod(route *RouteType, w http.ResponseWriter, r *http.Request) bool {
	if r.Method != route.HttpMethod {
		responsehandler.Json(w, "Route not found!", http.StatusNotFound)
		return false
	}

	return true
}

func makeHandler(route *RouteType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !checkHttpMethod(route, w, r) {return}

		for _, m := range route.MiddlewaresList {
			if !m.Handle(w, r) {return}
		}

		route.Callable(w, r)
	}
}
