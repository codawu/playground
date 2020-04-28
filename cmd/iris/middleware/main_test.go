package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

/**
1. Middleware is executed by peeling onions
2. "Next" function must be added in middleware
*/
func TestHelloWorld(t *testing.T) {
	app := iris.New()
	InitRouter(app)
	ht := httptest.New(t, app)
	want := "BeforeHelloWorldAfter"
	ht.GET("/hello/world").Expect().Status(httptest.StatusOK).Body().Equal(want)
}
