package routes

import (
	"files_manager/application"
	"files_manager/controllers"
	"files_manager/middleware"
	"net/http"

	"github.com/kataras/iris/v12/context"
)

var api = application.Server.Party("api/")

/* func goNext(c iris.Context) {
	c.Next()
} */

func init() {
	api.Options("*", func(context *context.Context) {
		context.StopWithStatus(http.StatusOK)
	})
	api.Post("/files/upload", middleware.AuthRequired(), controllers.UploadController)
	//api.Get("/files/me", middleware.AuthRequired(), controllers.MyFilesController)
	api.Get("/files/{uuid}", controllers.SingleFileController)

	{

		//Post a new document
		api.Post("/document", middleware.AuthRequired(), controllers.NewDocumentController)

		//Share a Document
		api.Post("/documents/forward", middleware.AuthRequired(), controllers.ForwardShareController)

		//Get all created document
		api.Get("/documents/", middleware.AuthRequired(), controllers.MyDcoumentsController)

		//Get all received shared document
		api.Get("/documents/received", middleware.AuthRequired(), controllers.MyReceivedSharesController)

		//Get all fowarded shared document
		api.Get("/documents/forwarded", middleware.AuthRequired(), controllers.MyForwardedSharesController)

		api.Put("/documents/{id}/{status}", middleware.AuthRequired(), controllers.SetStatusToDocumentAs)

		api.Delete("/documents/{id}", middleware.AuthRequired(), controllers.NewDocumentController)
		api.Get("/documents/{id}", middleware.AuthRequired(), controllers.GetDocument)
	}

	{
		//Get all user
		api.Get("/users", middleware.AuthRequired(), controllers.GetAllUsers)
	}
}
