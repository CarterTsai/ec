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
	// Static File Path
	iris.StaticServe("./public")
	// Favicon
	iris.Favicon("./public/favicon.ico")
	// template with django
	server.UseTemplate(jet.New()).Directory("./views", ".jet")
	// router
	router.Init(server)
	server.Listen(":8888")
	//server.ListenTLSAuto("127.0.0.1:443")
}
