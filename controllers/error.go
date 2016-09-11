package controller

import (
	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
)

var errorLogger = logger.New(nil)

// Error404 home route
func Error404(ctx *iris.Context) {
	errorLogger.Serve(ctx)
	ctx.Session().Set("name", "CaterTsai")
	ctx.Write("My Custom 404 error page ")
}
