package controllers

import (
	"errors"
	"github.com/kataras/iris/mvc"
)

//hello controller
type HelloController struct{}

var helloView = mvc.View{
	Name: "hello/index.html",
	Data: map[string]interface{}{
		"Title":     "Hello Page",
		"MyMessage": "Welcome to my awesome website",
	},
}

func (c *HelloController) Get() mvc.Result {
	return helloView
}

var errBadName = errors.New("bad name")

var badName = mvc.Response{Err: errBadName, Code: 400}

func (c *HelloController) GetBy(name string) mvc.Result {
	if name != "iris" {
		return badName
	}

	return mvc.View{
		Name: "hello/name.html",
		Data: name,
	}
}
