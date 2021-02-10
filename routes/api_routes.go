package routes

import (
	"files_manager/application"
	"files_manager/controllers"
	"files_manager/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"net/http"
)

var api = application.Server.Party("api/")

func goNext(c iris.Context) {
	c.Next()
}

func init() {
	api.Options("*", func(context *context.Context) {
		context.StopWithStatus(http.StatusOK)
	})
	api.Post("/files/upload", middleware.AuthRequired(), controllers.UploadController)
	api.Get("/files/me", middleware.AuthRequired(), controllers.MyFilesController)
	api.Get("/files/{uuid}", middleware.AuthRequired(), controllers.SingleFileController)
}
