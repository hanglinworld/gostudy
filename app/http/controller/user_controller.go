package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
	Controller
}

func (c *UserController) BeforeActivation(b mvc.BeforeActivation) {
	fmt.Println("BeforeActivation")
	b.Handle("GET", "/permissions", "GetPermissions")
}

func (c *UserController) GetPermissions(ctx iris.Context) string {
	fmt.Println("GetPermissions")
	ctx.Next()
	return "getPermissions"
}

func (c *UserController) Get(ctx iris.Context) string {
	ctx.Next()
	return "get"
}