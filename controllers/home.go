package controller

import (
	"fmt"

	"github.com/kataras/iris"
)

// Home home route
func Home(ctx *iris.Context) {
	ctx.Session().Set("name", "CaterTsai")
	fmt.Println(ctx.Session())
	ctx.Render("index.jet", map[string]interface{}{"username": "iris", "is_admin": true})
}
