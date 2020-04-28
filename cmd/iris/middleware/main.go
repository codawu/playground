package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

/**
1. 第一个handler执行完后没有调用"Next",会连续打印两个"No",调用结束
*/
func NoNext(app *iris.Application) {
	app.Get("/no/next", func(ctx context.Context) {
		_, _ = ctx.Write([]byte("No"))
		_, _ = ctx.Write([]byte("No"))
	}, func(ctx context.Context) {
		_, _ = ctx.Write([]byte("Next"))
	})
}

/**
	1. 第一个handler执行完后调用了"Next"，会继续执行第二个函数
	2. 第二个函数会连续打印两个"Next",然后回退到第一个函数
    3. 第一个函数继续执行"Next"函数之后的代码，打印"No"
*/
func LastNoNext(app *iris.Application) {
	app.Get("/no/next",
		func(ctx context.Context) {
			_, _ = ctx.Write([]byte("No"))
			ctx.Next()
			_, _ = ctx.Write([]byte("No"))
		}, func(ctx context.Context) {
			_, _ = ctx.Write([]byte("Next"))
			_, _ = ctx.Write([]byte("Next"))
		})
}

/**
	1. 第一个handler执行完后调用了"Next"，会继续执行第二个函数
	2. 第二个函数会连续打印两个"Next",然后回退到第一个函数
    3. 第一个函数继续执行"Next"函数之后的代码，打印"No"

	第三个函数不会被执行
*/
func SecondLastNoNext(app *iris.Application) {
	app.Get("/no/next",
		func(ctx context.Context) {
			_, _ = ctx.Write([]byte("No"))
			ctx.Next()
			_, _ = ctx.Write([]byte("No"))
		}, func(ctx context.Context) {
			_, _ = ctx.Write([]byte("Next"))
			_, _ = ctx.Write([]byte("Next"))
		},
		func(ctx context.Context) {
			_, _ = ctx.Write([]byte("Second"))
			_, _ = ctx.Write([]byte("Second"))
		})
}

/**
	1. 第一个handler执行完后调用了"Next"，会继续执行第二个函数
	2. 第二个函数会打印一个"Next",然后执行"Next"，因为第二个函数之后没有继续执行的函数，故不产生任何操作
		继续打印"Next"
    3. 第一个函数继续执行"Next"函数之后的代码，打印"All"
*/
func AllNext(app *iris.Application) {
	app.Get("/all/next",
		func(ctx context.Context) {
			_, _ = ctx.Write([]byte("All"))
			ctx.Next()
			_, _ = ctx.Write([]byte("All"))
		}, func(ctx context.Context) {
			_, _ = ctx.Write([]byte("Next"))
			ctx.Next()
			_, _ = ctx.Write([]byte("Next"))
		})
}
