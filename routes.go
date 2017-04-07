package main

import (
	"github.com/go-xorm/xorm"
	"github.com/mathsalmi/goarch/controllers"
	"github.com/mathsalmi/goarch/util"
	iris "gopkg.in/kataras/iris.v6"
)

// SetRoutes sets the routes for the application
func SetRoutes(app *iris.Framework, orm *xorm.Engine) {
	app.Get("/users", util.Setdb(orm), controllers.ListUser)
	app.Post("/users", util.Setdb(orm), controllers.CreateUser)
	app.Put("/users/:id", util.Setdb(orm), controllers.EditUser)
	app.Delete("/users/:id", util.Setdb(orm), controllers.DeleteUser)
}
