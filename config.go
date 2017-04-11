package main

import "github.com/mathsalmi/goarch/util"

var (
	port = util.Env("PORT", "8090")

	dbtype = "mysql"
	dburl  = "root:root@/express?charset=utf8"
)
