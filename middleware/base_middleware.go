package middleware

import (
	"github.com/kataras/iris/v12"
	"log"
)

func BasicMiddleware() iris.Handler {
	return basicMiddlewareHandler
}

func basicMiddlewareHandler(ctx iris.Context) {
	log.Print("Basic middleware called")
	ctx.Header("Provider", "Hiro Hamada")
	ctx.Next()
}
