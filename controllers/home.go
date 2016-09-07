package controller

import "github.com/kataras/iris"

// Home home route
func Home(ctx *iris.Context) {
	ctx.Session().Set("name", "CaterTsai")
	ctx.Render("index.jet", map[string]interface{}{"username": "iris", "is_admin": true})
}
