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
	api.Get("/files/{uuid}", controllers.SingleFileController)

	{
		api.Get("/documents/users", middleware.AuthRequired(), controllers.OtherUsersController)
		api.Post("/document", middleware.AuthRequired(), controllers.NewDocumentController)
		api.Post("/document/share", middleware.AuthRequired(), controllers.NewDocumentController)
		api.Get("/documents/", middleware.AuthRequired(), controllers.MyDocumentsController)
		api.Put("/documents/{id}/complete", middleware.AuthRequired(), controllers.NewDocumentController)
		api.Put("/documents/{id}/cancel", middleware.AuthRequired(), controllers.NewDocumentController)
		api.Delete("/documents/{id}", middleware.AuthRequired(), controllers.NewDocumentController)
		api.Get("/documents/{id}", middleware.AuthRequired(), controllers.SingleDocumentController)
		api.Post("/documents/{id}/comment", middleware.AuthRequired(), controllers.NewComment)
		api.Get("/documents/{id}/comments", middleware.AuthRequired(), controllers.GetComments)
	}
}
