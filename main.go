package main

import (
	"net/http"
	"simpleapp/app/controllers/admin/AdminController"
	"simpleapp/app/controllers/admin/AuthController"
	"simpleapp/app/modules/jwtHandler"
	responsehandler "simpleapp/app/modules/responseHandler"
)

func authMiddleware(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if len(token) < 1 || !jwtHandler.Verify(token) {
			responsehandler.Json(w, map[string]string{"message": "UnAuthenticated!"}, http.StatusBadRequest)

			return
		}

		fn(w, r)
	}
}

func main() {
	http.HandleFunc("/admin/admins", authMiddleware(AdminController.List))
	http.HandleFunc("/admin/admins/edit/", authMiddleware(AdminController.Update))

	http.HandleFunc("/admin/login", AuthController.Login)
	http.HandleFunc("/admin/register", authMiddleware(AuthController.Register))
	http.HandleFunc("/admin/refresh-token", AuthController.Refresh)

	http.ListenAndServe(":9090", nil)
}

/*
TODO:
	[] write tests
	[] validation handler system
	[] routing system
	[] middleware handler on routing system
*/
