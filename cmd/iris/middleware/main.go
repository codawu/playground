package main

import (
	"github.com/kataras/iris/v12"
)

func before(ctx iris.Context) {
	_, _ = ctx.Write([]byte("Before"))
	ctx.Next()
}

func after(ctx iris.Context) {
	ctx.Next()
	_, _ = ctx.Write([]byte("After"))
}

func HelloWorld(ctx iris.Context) {
	_, _ = ctx.Write([]byte("HelloWorld"))
}

func InitRouter(app *iris.Application) {
	app.Get("/hello/world", before, after, HelloWorld)
}
