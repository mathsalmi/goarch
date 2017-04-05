package main

import (
	"log"
	"os"
	"strings"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/cors"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func main() {
	app := iris.New()

	// start database
	orm, err := xorm.NewEngine(dbtype, dburl)
	if err != nil {
		log.Fatalln("Cannot connect to database: %v", err)
	}

	app.Adapt(
		// enable debug logger
		iris.DevLogger(),

		// enable router
		httprouter.New(),

		// enable cors
		cors.Default(),
	)

	// add routes
	SetRoutes(app, orm)

	// start the app
	app.Listen(port)
}

// env returns a OS env var or default value
func env(key string, defvalue string) string {
	env := os.Getenv(key)

	if env == "" {
		env = defvalue
	}

	if !strings.HasPrefix(env, ":") {
		env = ":" + env
	}

	return env
}
