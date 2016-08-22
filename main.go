package main

import (
	"ec/config"
	"ec/routers"

	"github.com/iris-contrib/template/django"
	"github.com/kataras/iris"
)

func main() {
	// config
	server := iris.New(config.Init())
	// template with django
	server.UseTemplate(django.New()).Directory("./views", ".html")
	// router
	router.Init(server)
	server.Listen(":8080")
	//server.ListenTLSAuto("127.0.0.1:443")
}
