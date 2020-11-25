package bootstrap

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"study/favor/config"
	"study/favor/routes"
)

func Load(app *iris.Application) {
	config.Load(app)

	// 因为是接口路由， 这里会用到跨域
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	mvc.Configure(app.Party("/api", crs), routes.RegisterApiRoutes)
}
