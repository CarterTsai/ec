package controller

import "github.com/kataras/iris"

// User user router
func User(ctx *iris.Context) {
	name := ctx.Session().GetString("name")

	ctx.Render("user.jet", map[string]interface{}{"username": name})
}
