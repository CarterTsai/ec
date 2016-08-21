package main

import (
	"ec/config"
	"ec/routers"

	"github.com/iris-contrib/template/django"
	"github.com/kataras/iris"
)

func main() {
	// config
	server := iris.New(config.Config())
	server.UseTemplate(django.New()).Directory("./views", ".html")

	server.Get("/", router.Home)
	server.Listen(":8080")
}
