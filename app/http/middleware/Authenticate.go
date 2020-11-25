package middleware

import (
	"github.com/kataras/iris/v12"
)

func AuthHandler(ctx iris.Context) {
	ctx.Next()
}