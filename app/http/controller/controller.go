package controller

import (
	"fmt"
	"github.com/kataras/iris/v12/mvc"
)

type Controller struct {
}

func (c *Controller) AfterActivation(a mvc.AfterActivation) {
	fmt.Println("AfterActivation")
}
