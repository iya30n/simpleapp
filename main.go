package main

import (
	"net/http"
	"simpleapp/controllers/admin/AdminController"
	"simpleapp/controllers/admin/AuthController"
)

func main() {
	http.HandleFunc("/admin/admins", AdminController.List)
	http.HandleFunc("/admin/login", AuthController.Login)

	http.ListenAndServe(":9090", nil)
}