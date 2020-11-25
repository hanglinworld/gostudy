package main

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"

	"fmt"
	"study/favor/bootstrap"
)

func main() {
	app := iris.New()
	bootstrap.Load(app)
	app.Run(iris.Addr(fmt.Sprintf(":%d", viper.GetInt("app.port"))))
}
