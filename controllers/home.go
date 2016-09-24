package controller

import (
	"github.com/CloudyKit/jet"
	"github.com/kataras/iris"
)

// Home home route
func Home(ctx *iris.Context) {
	ctx.Session().Set("name", "CaterTsai")
	ctx.Render("index.jet", jet.VarMap{"username": "iris", "is_admin": true})
}
