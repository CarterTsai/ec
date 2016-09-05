package router

import (
	controller "ec/controllers"

	"github.com/kataras/iris"
)

//Init Init Router
func Init(server *iris.Framework) {
	server.Get("/", controller.Home)
	server.Get("/user", controller.User)
	server.Get("/login", controller.Login)
}
