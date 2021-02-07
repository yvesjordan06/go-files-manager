package controllers

import (
	"github.com/kataras/iris/v12"
	"net/http"
)

func Home(ctx iris.Context) {
	ctx.StatusCode(http.StatusOK)
	ctx.Text("Working")
}
