package controllers

import (
	"files_manager/configs"
	"files_manager/models"
	"files_manager/models/base"
	"github.com/kataras/iris/v12"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"os"
	"path/filepath"
)

//UploadController is the controller for uploading files to the server
//This controller will upload then return the file id which can then be used
func UploadController(ctx iris.Context) {
	_, fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": err.Error()})
		return
	}

	if _, err := os.Stat(configs.Application.UploadDir); os.IsNotExist(err) {
		os.Mkdir(configs.Application.UploadDir, os.ModePerm)
	}
	uid := uuid.NewV4()
	dest := filepath.Join(configs.Application.UploadDir, uid.String())
	_, err = ctx.SaveFormFile(fileHeader, dest)

	if err != nil {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": err.Error()})
		return
	}

	file := &models.File{
		BaseUUID: base.BaseUUID{ID: &uid},
		Name:     fileHeader.Filename,
		UserID:   ctx.Values().Get("user").(*models.User).ID,
		Size:     fileHeader.Size,
		User:     ctx.Values().Get("user").(*models.User),
	}

	_, err = file.Create()

	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error()})
		return
	}

	ctx.StopWithJSON(http.StatusCreated, file)
}

//MyFilesController returns the list of files added by the user
func MyFilesController(ctx iris.Context) {
	files := new(models.Files)
	user := ctx.Values().Get("user").(*models.User)
	_, err := files.Where(&models.File{UserID: user.ID})

	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error()})
		return
	}

	ctx.StopWithJSON(http.StatusOK, files)

}

//SingleFileController is responsible for returning the specified file, But it checks if the requesting user has
//permission to view it
func SingleFileController(ctx iris.Context) {
	file := new(models.File)
	fileUid := ctx.Params().Get("uuid")
	fileUuid := uuid.FromStringOrNil(fileUid)
	user := ctx.Values().Get("user").(*models.User)
	_, err := file.Get(&models.File{
		BaseUUID: base.BaseUUID{
			ID: &fileUuid,
		},
		UserID: user.ID,
	})

	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error(), "suggestion": "The file you are trying to access does not exist or you don't have permission to view its content."})
		return
	}

	ctx.StatusCode(http.StatusOK)
	ctx.SendFile(filepath.Join(configs.Application.UploadDir, fileUid), file.Name)
	ctx.StopExecution()

}

func NewDocumentController(ctx iris.Context) {
	document := new(models.Document)
	err := ctx.ReadJSON(document)
	user := ctx.Values().Get("user").(*models.User)

	if err != nil {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": err.Error(), "suggestion": "Verify your input and try again."})
		return
	}

	//Check if file exist and the user has rights to send
	file := new(models.File)
	_, err = file.Get(&models.File{
		BaseUUID: base.BaseUUID{ID: &document.FileID},
	})

	if err != nil {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": err.Error(), "suggestion": "Your file does not exist"})
		return
	}

	documents := new(models.Documents)

	r, _ := documents.Where(&models.Document{FileID: document.FileID, AssignedID: &user.ID})

	if r.RowsAffected == 0 && file.UserID != user.ID {
		ctx.StopWithJSON(http.StatusUnauthorized, iris.Map{"error": "No permission", "suggestion": "The file you are trying to share is not available for you."})
		return
	}

	document.UserID = user.ID
	document.User = user
	document.Status = models.DocumentStatusPending
	document.File = file

	_, err = document.Create()
	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error(), "suggestion": "We have issues creating your document"})
		return
	}

	ctx.StopWithJSON(http.StatusCreated, document)

}
