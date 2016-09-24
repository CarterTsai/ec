package controller

import (
	"github.com/CloudyKit/jet"
	"github.com/kataras/iris"
)

// Home home route
func Home(ctx *iris.Context) {
	//ctx.Session().Set("name", "CaterTsai")

	vars := make(jet.VarMap)
	vars.Set("title", "Ec")
	ctx.Render("index.jet", vars)
}
