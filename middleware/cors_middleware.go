package middleware

import (
	"github.com/kataras/iris/v12"
	"log"
)

func CorsMiddleware() iris.Handler {
	return func(ctx iris.Context) {
		log.Println("Cors middleware")
		ctx.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		ctx.Header("Access-Control-Allow-Methods", "*")
		ctx.Header("Access-Control-Allow-Headers", "*")

		ctx.Next()
	}
}
