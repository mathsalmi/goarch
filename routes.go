package main

import (
	"github.com/mathsalmi/goarch/controllers"
	iris "gopkg.in/kataras/iris.v6"
)

// SetRoutes sets the routes for the application
func SetRoutes(app *iris.Framework) {
	app.Get("/users", controllers.ListUser)
	app.Post("/users", controllers.CreateUser)
	app.Put("/users/:id", controllers.EditUser)
	app.Delete("/users/:id", controllers.DeleteUser)
}
