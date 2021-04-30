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
		//Create a new document
		api.Post("/document", middleware.AuthRequired(), controllers.NewDocumentController)

		//Get Current user created Documents
		api.Get("/documents/", middleware.AuthRequired(), controllers.MyDocumentsController)

		//Get Current user created Documents
		api.Get("/documents/{id}", middleware.AuthRequired(), controllers.SingleDocumentController)

		//Share a document by ID
		api.Post("/documents/{id}/share", middleware.AuthRequired(), controllers.ShareDocument)

		//Get a document shares
		api.Get("/documents/{id}/shares", middleware.AuthRequired(), controllers.GetDocumentShares)

		//Get all shared document I received or send
		api.Get("/shares", middleware.AuthRequired(), controllers.GetShared)

		//Get to open a share i received or sent
		api.Get("/shares/{id}", middleware.AuthRequired(), controllers.OpenShare)

		//To forward a share
		api.Post("/shares/{id}/forward", middleware.AuthRequired(), controllers.ForwardShare)
		//To complete a share
		api.Patch("/shares/{id}/complete", middleware.AuthRequired(), controllers.CompleteShareStatus)

		//To cancel a share
		api.Patch("/shares/{id}/cancel", middleware.AuthRequired(), controllers.CancelShareStatus)

		//To delete a share
		api.Delete("/shares/{id}", middleware.AuthRequired(), controllers.DeleteShare)

		// Comments section on a share document

		api.Post("/documents/{id}/comment", middleware.AuthRequired(), controllers.NewComment)
		api.Get("/documents/{id}/comments", middleware.AuthRequired(), controllers.GetComments)
	}
}
