package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

const (
	NoNextStr           = "NoNo"
	LastNoNextStr       = "NoNextNextNo"
	SecondLastNoNextStr = "NoNextNextNo"
	AllNextStr          = "AllNextNextAll"
)

/**
1. Middleware is executed by peeling onions
2. "Next" function must be added in middleware
*/
func TestNoNext(t *testing.T) {
	app := iris.New()
	NoNext(app)
	ht := httptest.New(t, app)
	want := fmt.Sprintf("%s", NoNextStr)
	ht.GET("/no/next").Expect().Status(httptest.StatusOK).Body().Equal(want)
	if !t.Failed() {
		t.Logf("✌️️️✌️️️✌️️️✌️️️✌️️️")
	} else {
		t.Logf("💔💔💔💔💔")
	}
}

func TestLastNoNext(t *testing.T) {
	app := iris.New()
	LastNoNext(app)
	ht := httptest.New(t, app)
	want := fmt.Sprintf("%s", LastNoNextStr)
	ht.GET("/no/next").Expect().Status(httptest.StatusOK).Body().Equal(want)
	if !t.Failed() {
		t.Logf("✌️️️✌️️️✌️️️✌️️️✌️️️")
	} else {
		t.Logf("💔💔💔💔💔")
	}
}

func TestSecondLastNoNext(t *testing.T) {
	app := iris.New()
	SecondLastNoNext(app)
	ht := httptest.New(t, app)
	want := fmt.Sprintf("%s", SecondLastNoNextStr)
	ht.GET("/no/next").Expect().Status(httptest.StatusOK).Body().Equal(want)
	if !t.Failed() {
		t.Logf("✌️️️✌️️️✌️️️✌️️️✌️️️")
	} else {
		t.Logf("💔💔💔💔💔")
	}
}

func TestAllNext(t *testing.T) {
	app := iris.New()
	AllNext(app)
	ht := httptest.New(t, app)
	want := fmt.Sprintf("%s", AllNextStr)
	ht.GET("/all/next").Expect().Status(httptest.StatusOK).Body().Equal(want)
	if !t.Failed() {
		t.Logf("✌️️️✌️️️✌️️️✌️️️✌️️️")
	} else {
		t.Logf("💔💔💔💔💔")
	}

}
