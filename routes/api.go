package routes

import (
	//"favor/app/http/controllers/api"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"study/favor/app/http/middleware"
	"study/favor/app/http/controller"
)

func RegisterApiRoutes(app *mvc.Application) {
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Api before: Path = %s \n\n", ctx.Path())
		ctx.Next()
	})

	app.Router.Done(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Api after: %s \n\n", ctx.Path())
		ctx.Next()
	})

	app.Handle(new(controller.LoginController))

	app.Party("/user", middleware.AuthHandler).Handle(new(controller.UserController))
}
