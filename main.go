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
	// template with enginer
	server.UseTemplate(jet.New()).Directory("./views", ".jet")
	// router
	router.Init(server)
	// run Server
	server.Listen(":8888")
	//server.ListenTLSAuto("127.0.0.1:443")
}
