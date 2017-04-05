package main

var (
	port = env("PORT", "8090")

	dbtype = "mysql"
	dburl  = "root:root@/express?charset=utf8"
)
