package routes

import (
	"files_manager/application"
	"files_manager/controllers"
	"github.com/kataras/iris/v12"
)

var api = application.Server.Party("api/")

func goNext(c iris.Context) {
	c.Next()
}

func init() {
	api.Post("/upload", controllers.UploadController)
}
