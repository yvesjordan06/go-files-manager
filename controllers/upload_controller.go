package controllers

import (
	"files_manager/configs"
	"github.com/kataras/iris/v12"
	"net/http"
	"path/filepath"
)

func UploadController(ctx iris.Context) {
	_, fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": err.Error()})
		return
	}

	dest := filepath.Join(configs.Application.UploadDir, fileHeader.Filename)
	_, err = ctx.SaveFormFile(fileHeader, dest)

	if err != nil {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": err.Error()})
		return
	}

	ctx.Writef("File: %s uploaded", fileHeader.Filename)
}
