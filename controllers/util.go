package controllers

import (
	"github.com/go-xorm/xorm"
	"gopkg.in/kataras/iris.v6"
)

func Setdb(orm *xorm.Engine) iris.HandlerFunc {
	return func(ctx *iris.Context) {
		ctx.Set("orm", orm)
		ctx.Next()
	}
}

func Getdb(ctx *iris.Context) *xorm.Engine { // TODO add an error
	return ctx.Get("orm").(*xorm.Engine)
}
