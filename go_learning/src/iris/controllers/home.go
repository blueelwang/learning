package controllers

import (
	// "github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
)

type HomeController struct {}

func (c *HomeController) Get() mvc.Result {
	return mvc.Response {
		ContentType: "text/html",
		Text: "<h1>Home</h1>",
	}
}

func (c *HomeController) GetPing() interface{} {
	return map[string]interface{} {
		"code": 0,
		"msg": "ok",
	}
}