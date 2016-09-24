package main

import (
	"ec/config"
	"ec/lib/templateEngine"
	"ec/routers"

	"github.com/kataras/iris"
)

func main() {
	// config
	server := iris.New(config.Init())
	// logger

	// template with enginer
	viewGlobal := jet.DefaultConfig()
	viewGlobal.Layout = ""
	viewGlobal.Vars["Version"] = "v1.0.0"
	viewGlobal.Vars["Title"] = "EC - V1"

	server.UseTemplate(jet.New(viewGlobal)).Directory("./views", ".jet")
	// router
	router.Init(server)
	// run Server
	server.Listen(":8888")
	//server.ListenTLSAuto("127.0.0.1:443")
}
