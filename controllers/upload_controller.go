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
