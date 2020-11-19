package main 

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"go_learning/thirdparty/iris/controllers"
)

func main() {
	app := iris.New()
	
	// app.OnErrorCode(iris.StatusNotFound, errorHandler)
	// app.OnErrorCode(iris.StatusInternalServerError, errorHandler)
    // 为所有的大于等于400的状态码注册一个处理器：
    app.OnAnyErrorCode(errorHandler)

	app.Handle("GET", "/name/{name}", func(ctx iris.Context) {
		params := ctx.Params()
		name := params.Get("name")
		ctx.Writef("<h1>Hello %s!</h1>", name)
	})

	mvc.New(app.Party("/admin")).Handle(new(controllers.AdminController))
	mvc.New(app.Party("/home")).Handle(new(controllers.HomeController))
	
	go app.Run(iris.Addr(":8090"), iris.WithConfiguration(iris.Configuration{
		DisableStartupLog: true,
	}))
	app.Run(iris.Addr(":8091"))
}
func errorHandler(ctx iris.Context) {
	ctx.Writef("<b>Internal server error! Code:%d</b><br/>", ctx.ResponseWriter().StatusCode())
}
