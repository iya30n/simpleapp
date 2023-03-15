package router

import (
	"net/http"
	responsehandler "simpleapp/core/responseHandler"
)

func Serve(port string) {
	for _, route := range routes {
		http.HandleFunc(route.Url, checkHttpMethod(route))
	}

	http.ListenAndServe(":"+port, nil)
}

func checkHttpMethod(route RouteType) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != route.HttpMethod {
			responsehandler.Json(w, "Route not found!", http.StatusNotFound)
			return
		}

		route.Callable(w, r)
	}
}
