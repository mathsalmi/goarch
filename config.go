package main

import "./util"

var (
	port = util.Env("PORT", "8090")

	dbtype = "mysql"
	dburl  = "root:root@/express?charset=utf8"
)
