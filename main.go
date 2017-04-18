package main

import (
	"log"
	"strings"

	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/cors"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"

	"./util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func init() {
	// start database
	orm, err := xorm.NewEngine(dbtype, dburl)
	if err != nil {
		log.Fatalln("Cannot connect to database: %v", err)
	}
	util.SetDb(orm)
}

func main() {
	app := iris.New()

	app.Adapt(
		// enable debug logger
		iris.DevLogger(),

		// enable router
		httprouter.New(),

		// enable cors
		cors.Default(),
	)

	// add routes
	SetRoutes(app)

	// start the app
	app.Listen(fmtPort(port))
}

// fmtPort adds : as prefix to port if it has none
func fmtPort(port string) string {
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	return port
}
