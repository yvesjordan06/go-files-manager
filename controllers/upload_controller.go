package controllers

import (
	"errors"
	"files_manager/configs"
	"files_manager/models"
	"files_manager/models/base"
	"github.com/kataras/iris/v12"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
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
		_ = os.Mkdir(configs.Application.UploadDir, os.ModePerm)
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
	//user := ctx.Values().Get("user").(*models.User)
	_, err := file.Get(&models.File{
		BaseUUID: base.BaseUUID{
			ID: &fileUuid,
		},
		//UserID: user.ID,
	})

	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error(), "suggestion": "The file you are trying to access does not exist or you don't have permission to view its content."})
		return
	}

	ctx.StatusCode(http.StatusOK)
	_ = ctx.SendFile(filepath.Join(configs.Application.UploadDir, fileUid), file.Name)
	ctx.StopExecution()

}

/***
NewDocumentController create a new document from an existing file
*/
func NewDocumentController(ctx iris.Context) {
	document := new(models.Document)
	share := new(models.Share)
	documents := new(models.Documents)
	err := ctx.ReadJSON(document)
	err = ctx.ReadJSON(share)
	user := ctx.Values().Get("user").(*models.User)

	if err != nil {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": err.Error(), "suggestion": "Verify your input and try again."})
		return
	}

	//Checking if the file has already been shared or assigned

	// Checking if the document with the file already exist
	r, _ := documents.Where(&models.Document{FileID: document.FileID})

	if r.RowsAffected > 0 {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": "This file already has a document", "suggestion": "Please create another document"})
		return
	}

	//Checking if the receiver is the sender
	if user.ID == share.ReceiverID {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": "You can not send a document to yourself", "suggestion": "You are trying to send this document to yourself, please choose another receiver"})
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

	if file.UserID != user.ID {
		ctx.StopWithJSON(http.StatusUnauthorized, iris.Map{"error": "No permission", "suggestion": "The file you are trying to share is not available for you."})
		return
	}

	receiver := new(models.User)
	_, err = receiver.Get(&models.User{
		Base: base.Base{
			ID: share.ReceiverID,
		},
	})

	if err != nil {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": err.Error(), "suggestion": "The receiver does not exist"})
		return
	}

	document.UserID = user.ID
	document.User = user
	share.Status = models.DocumentStatusPending
	document.File = file
	share.ReceiverID = receiver.ID
	share.SenderID = user.ID
	share.Sender = user
	share.Receiver = receiver

	_, err = document.Create()
	share.DocumentID = document.ID
	share.Document = document
	share.Status = models.DocumentStatusPending
	_, err = share.Create()
	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error(), "suggestion": "We have issues creating your document"})
		return
	}

	ctx.StopWithJSON(http.StatusCreated, share)

}

func MyDocumentsController(ctx iris.Context) {

	documents := &models.Documents{}
	var err error
	user := ctx.Values().Get("user").(*models.User)

	_, err = documents.Where("user_id = ?", user.ID)

	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error()})
		return
	}

	/*documents2 := *documents
	r, err := documents2.Where("received_at IS ?", nil)
	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error()})
		return
	}
	t := time.Now()

	if r.RowsAffected > 0 {
		r.Updates(&models.Document{Status: models.DocumentStatusRead, ReceivedAt: &t})
	}*/

	ctx.StopWithJSON(http.StatusOK, documents)
}

func ReceivedDocuments(ctx iris.Context) {
	user := ctx.Values().Get("user").(*models.User)
	shares := new(models.Shares)
	_, err := shares.Where(&models.Share{ReceiverID: user.ID})

	if err != nil {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": err.Error(), "suggestion": "An unknown error occured"})
		return
	}

	ctx.StopWithJSON(http.StatusOK, shares)
}

func OtherUsersController(ctx iris.Context) {
	user := ctx.Values().Get("user").(*models.User)
	users := new(models.Users)
	_, err := users.Where("id <> ?", user.ID)
	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error(), "suggestion": "We had problems finding other users"})
		return
	}

	ctx.StopWithJSON(http.StatusOK, users)
}

func SingleDocumentController(ctx iris.Context) {
	//user := ctx.Values().Get("user").(*models.User)
	document := new(models.Document)
	documentID, _ := ctx.Params().GetInt("id")
	_, err := document.Get(&models.Document{Base: base.Base{ID: uint(documentID)}})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.StopWithJSON(http.StatusNotFound, iris.Map{"error": "Document not found"})

		} else {
			ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error()})
		}

		return
	}

	ctx.StopWithJSON(http.StatusOK, document)
}

func NewComment(ctx iris.Context) {
	user := ctx.Values().Get("user").(*models.User)
	documentID, _ := ctx.Params().GetInt("id")
	document := new(models.Document)
	comment := new(models.Comment)

	err := ctx.ReadJSON(comment)

	if err != nil {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": err.Error(), "suggestion": "Verify your input and try again."})
		return
	}

	_, err = document.Get(&models.Document{Base: base.Base{ID: uint(documentID)}})

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.StopWithJSON(http.StatusNotFound, iris.Map{"error": "Document not found"})

		} else {
			ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error()})
		}

		return
	}

	comment.UserID = user.ID
	comment.User = user
	comment.DocumentID = &document.ID
	_, err = comment.Create()

	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error()})
		return
	}

	ctx.StopWithJSON(http.StatusOK, comment)

}

func GetComments(ctx iris.Context) {
	documentID, _ := ctx.Params().GetInt("id")
	docID := uint(documentID)
	comments := new(models.Comments)

	_, err := comments.Where(&models.Comment{DocumentID: &docID})

	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error()})
		return
	}

	ctx.StopWithJSON(http.StatusOK, comments)

}
