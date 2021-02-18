package routes

import (
	"files_manager/application"
	_ "files_manager/boostrap"
	"files_manager/configs"
	"github.com/kataras/iris/v12"
)

func init() {
	application.Server.HandleDir("/", iris.Dir("./public/file-manager/build/"), iris.DirOptions{
		Compress:  true,
		ShowList:  false,
		SPA:       true,
		IndexName: "index.html",
	})
	application.Server.HandleDir("uploads/", iris.Dir(configs.Application.UploadDir), iris.DirOptions{
		Compress: true,
		ShowList: false,
		//SPA: true,
	})
}
