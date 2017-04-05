package main

import (
	"github.com/go-xorm/xorm"
	"github.com/mathsalmi/goarch/controllers"
	iris "gopkg.in/kataras/iris.v6"
)

// SetRoutes sets the routes for the application
func SetRoutes(app *iris.Framework, orm *xorm.Engine) {
	app.Get("/users", controllers.Setdb(orm), controllers.ListUser)
	app.Post("/users", controllers.Setdb(orm), controllers.CreateUser)
	app.Put("/users/:id", controllers.Setdb(orm), controllers.EditUser)
	app.Delete("/users/:id", controllers.Setdb(orm), controllers.DeleteUser)
}
