package router

import (
	controller "ec/controllers"

	"github.com/kataras/iris"
)

//Init Init Router
func Init(server *iris.Framework) {
	// Favicon
	server.Favicon("./public/favicon.ico")
	// Static File
	server.StaticServe("./public", "/public")
	// All Router
	server.Get("/", controller.Home)
	server.Get("/user", controller.User)
	server.Get("/login", controller.Login)
	// Error Router
	server.OnError(iris.StatusNotFound, controller.Error404)
}
