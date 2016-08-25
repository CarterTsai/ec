package main

import (
	"ec/config"
	"ec/lib/template"
	"ec/routers"

	"github.com/kataras/iris"
)

func main() {
	// config
	server := iris.New(config.Init())
	// Favicon
	iris.Favicon("./public/favicon.ico")
	// template with django
	server.UseTemplate(jet.New()).Directory("./views", ".html")
	// router
	router.Init(server)
	server.Listen(":8888")
	//server.ListenTLSAuto("127.0.0.1:443")
}
