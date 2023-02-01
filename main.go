package main

import (
	"net/http"
	"simpleapp/controllers/admin/AdminController"
)

func main() {
	http.HandleFunc("/admin/admins", AdminController.List)
	http.ListenAndServe(":9090", nil)
}
