package main

import (
	"net/http"
	"simpleapp/controllers/admin/AdminController"
	"simpleapp/controllers/admin/AuthController"
	"simpleapp/modules/jwtHandler"
	responsehandler "simpleapp/modules/responseHandler"
)

func authMiddleware(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if len(token) < 1 || !jwtHandler.Verify(token) {
			responsehandler.Json(w, map[string]string{"message": "UnAuthenticated!",}, http.StatusBadRequest)
			
			return
		}

		fn(w, r)
	}
}

func main() {
	http.HandleFunc("/admin/admins", authMiddleware(AdminController.List))

	http.HandleFunc("/admin/login", AuthController.Login)
	http.HandleFunc("/admin/register", authMiddleware(AuthController.Register))
	http.HandleFunc("/admin/refresh-token", AuthController.Refresh)

	http.ListenAndServe(":9090", nil)
}
