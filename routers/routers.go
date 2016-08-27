package router

import (
	"fmt"

	"github.com/kataras/iris"
)

//Init Init Router
func Init(server *iris.Framework) {
	server.Get("/", Home)
	server.Get("/user", User)
}

// Home home route
func Home(ctx *iris.Context) {
	ctx.Session().Set("name", "CaterTsai")
	fmt.Println(ctx.Session())
	ctx.Render("index.jet", map[string]interface{}{"username": "iris", "is_admin": true})
}

// User user router
func User(ctx *iris.Context) {
	name := ctx.Session().GetString("name")
	fmt.Println(ctx.Session())
	ctx.Render("user.html", map[string]interface{}{"username": name})
}
