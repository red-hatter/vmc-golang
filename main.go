package main

import (
	"vmc-golang/datasource"
	"vmc-golang/repositories"
	"vmc-golang/services"
	"vmc-golang/web/controllers"
	"vmc-golang/web/middleware"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	_ "vmc-golang/emq"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.RegisterView(iris.HTML("./web/views", ".html"))

	mvc.New(app.Party("/hello")).Handle(new(controllers.HelloController))

	mvc.Configure(app.Party("/movies"), movies)

	//设置启动参数
	app.Run(
		// 端口
		iris.Addr("localhost:8080"),
		// 错误处理
		iris.WithoutServerError(iris.ErrServerClosed),
		// 实现
		iris.WithOptimizations,
	)
}

func movies(app *mvc.Application) {
	app.Router.Use(middleware.BasicAuth)
	repo := repositories.NewMovieRepository(datasource.Movies)
	movieService := services.NewMovieService(repo)
	app.Register(movieService)

	app.Handle(new(controllers.MovieController))
}
