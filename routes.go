package main

import (
	"github.com/go-xorm/xorm"
	"github.com/mathsalmi/goarch/controllers"
	iris "gopkg.in/kataras/iris.v6"
)

// SetRoutes sets the routes for the application
func SetRoutes(app *iris.Framework, orm *xorm.Engine) {
	app.Get("/", controllers.Setdb(orm), controllers.Index)
}
