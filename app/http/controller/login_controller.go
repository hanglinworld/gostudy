package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type LoginController struct {
	Controller
}

func (c *LoginController) BeforeActivation(b mvc.BeforeActivation) {
	fmt.Println("BeforeActivation")
	b.Handle("POST", "/register", "Register")
	b.Handle("POST", "/login", "Login")
}

func (c LoginController) Register(ctx iris.Context) string {
	ctx.Next()

	return "Register"
}

func (c LoginController) Login(ctx iris.Context) string {
	ctx.Next()

	return "Login"
}
