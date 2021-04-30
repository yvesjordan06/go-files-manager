package middleware

import (
	"github.com/kataras/iris/v12"
	"log"
)

func CorsMiddleware() iris.Handler {
	return func(ctx iris.Context) {
		log.Println("Cors middleware")
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "*")

		ctx.Next()
	}
}
