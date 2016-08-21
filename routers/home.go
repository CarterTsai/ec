package router

import "github.com/kataras/iris"

// Home home route
func Home(ctx *iris.Context) {
	ctx.Render("index.html", map[string]interface{}{"username": "iris", "is_admin": true}, iris.RenderOptions{"gzip": true})
}
