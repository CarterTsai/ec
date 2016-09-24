package controller

import (
	"github.com/CloudyKit/jet"
	"github.com/kataras/iris"
)

// Advert home route
func Advert(ctx *iris.Context) {

	vars := make(jet.VarMap)
	vars.Set("title", "Advert")
	ctx.Render("advert.jet", vars)
}
