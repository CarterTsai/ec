package controller

import (
	"github.com/CloudyKit/jet"
	"github.com/kataras/iris"
)

// Login login router
func Login(ctx *iris.Context) {

	vars := make(jet.VarMap)
	vars.Set("title", "Ec")

	ctx.Render("login.jet", nil)
}
