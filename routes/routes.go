package routes

import (
	"files_manager/application"
	"files_manager/configs"
	"files_manager/controllers"
	"github.com/kataras/iris/v12"
)

func init() {
	application.Server.Get("/", controllers.Home)
	application.Server.HandleDir("uploads/", iris.Dir(configs.Application.UploadDir), iris.DirOptions{
		Compress: true,
		ShowList: false,
		//SPA: true,
	})
}
