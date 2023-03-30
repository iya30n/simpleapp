package main

import (
	"simpleapp/app/controllers/admin/AdminController"
	"simpleapp/app/controllers/admin/AuthController"
	"simpleapp/app/middlewares"
	"simpleapp/core/router"
)

func main() {
	router.Get("/admin/admins", AdminController.List).Middlewares(middlewares.Auth{})
	router.Post("/admin/admins/edit/", AdminController.Update).Middlewares(middlewares.Auth{})

	router.Post("/admin/login", AuthController.Login)
	router.Post("/admin/register", AuthController.Register).Middlewares(middlewares.Auth{})
	router.Post("/admin/refresh-token", AuthController.Refresh)

	router.Serve("9090")
}