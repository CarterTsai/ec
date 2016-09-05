package controller

import "github.com/kataras/iris"

// Login login router
func Login(ctx *iris.Context) {
	ctx.Render("login.jet", nil)
}
