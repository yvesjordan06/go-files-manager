package controllers

import (
	"files_manager/application"
	"files_manager/configs"
	"files_manager/models"
	"files_manager/models/base"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/kataras/iris/v12"
	uuid "github.com/satori/go.uuid"
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

	//Creating the new id
	uid := uuid.NewV4()

	//Creating the file destination and saving the file
	dest := filepath.Join(configs.Application.UploadDir, uid.String())
	_, err = ctx.SaveFormFile(fileHeader, dest)

	if err != nil {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": err.Error()})
		return
	}

	//Creating the object that links to the file and saving it to the Database
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

//MyReceivedSharesController is used to get the shares of a user
// Those document shared with the user
func MyReceivedSharesController(ctx iris.Context) {
	shares := new(models.Shares)

	//Getting the current user
	user := ctx.Values().Get("user").(*models.User)

	//Getting the shares where the user is the receiver
	_, err := shares.Where(&models.Share{ReceiverID: user.ID})

	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error()})
		return
	}

	ctx.StopWithJSON(http.StatusOK, shares)

}

//MyFowardedSharesController is used to get the shares of a user
// Those document shared with the user
func MyForwardedSharesController(ctx iris.Context) {

	shares := new(models.Shares)

	//Getting the current user
	user := ctx.Values().Get("user").(*models.User)

	//Getting the shares where the user is the receiver
	_, err := shares.Where(&models.Share{UserID: user.ID})

	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error()})
		return
	}

	ctx.StopWithJSON(http.StatusOK, shares)

}

//MyFowardedSharesController is used to foward a share from a user to another
// Those document shared with the user
func ForwardShareController(ctx iris.Context) {
	receiver := struct {
		ShareID    int `json:"share_id" validate:"required"`
		ReceiverID int `json:"receiver_id" validate:"required"`
	}{}

	err := ctx.ReadJSON(&receiver)

	if err != nil {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": err.Error(), "suggestion": "Verify your input and try again."})
		return
	}

	user := ctx.Values().Get("user").(*models.User)

	//Getting the initial share to fowared
	initialShare := new(models.Share)
	_, err = initialShare.Get(&models.Share{Base: base.Base{
		ID: uint(receiver.ShareID),
	}})

	if err != nil {
		ctx.StopWithJSON(http.StatusNotFound, iris.Map{"error": err.Error()})
		return
	}

	//Checking if the user is the receiver of the share
	if user.ID != initialShare.ReceiverID {
		ctx.StopWithJSON(http.StatusNotFound, iris.Map{"error": "You do not have access to the document you are trying to share"})
		return
	}

	share := &models.Share{
		UserID:     user.ID,
		DocumentID: initialShare.DocumentID,
		Document:   initialShare.Document,
		ReceiverID: uint(receiver.ReceiverID),
	}

	_, err = share.Create()

	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error()})
		return
	}

	// Updating the document status
	updateDocument := new(models.Document)
	_, _ = updateDocument.Get(&models.Document{Base: base.Base{ID: initialShare.DocumentID}})
	updateDocument.LastShare = &share.ID
	updateDocument.Status = models.DocumentStatusPending
	_, err = updateDocument.Save()

	if err != nil {

		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error()})
		return
	}

	ctx.StopWithJSON(http.StatusOK, share)

}

//SingleFileController is responsible for returning the specified file, But it checks if the requesting user has
//permission to view it
func SingleFileController(ctx iris.Context) {
	file := new(models.File)
	strUuid := ctx.Params().Get("uuid")
	fileUuid := uuid.FromStringOrNil(strUuid)
	_, err := file.Get(&models.File{
		BaseUUID: base.BaseUUID{
			ID: &fileUuid,
		},
	})

	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error(), "suggestion": "The file you are trying to access does not exist or you don't have permission to view its content."})
		return
	}

	ctx.StatusCode(http.StatusOK)
	_ = ctx.SendFile(filepath.Join(configs.Application.UploadDir, strUuid), file.Name)
	ctx.StopExecution()

}

func NewDocumentController(ctx iris.Context) {
	document := new(models.Document)
	documents := new(models.Documents)

	json := struct {
		Title      string `json:"title,omitempty" validate:"required"` //The tilte of the documents
		Reference  string `json:"reference,omitempty" validate:"required"`
		Object     string `json:"object,omitempty" validate:"required"`
		FileID     string `json:"file_id,omitempty" validate:"required"`
		ReceiverID int    `json:"receiver_id,omitempty" validate:"required"`
	}{}
	err := ctx.ReadJSON(&json)
	user := ctx.Values().Get("user").(*models.User)

	if err != nil {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": err.Error(), "suggestion": "Verify your input and try again."})
		return
	}

	document.Object = json.Object
	document.Reference = json.Reference
	document.Title = json.Title

	document.FileID = uuid.FromStringOrNil(json.FileID)

	//Checking if the file has already been shared or assigned

	r, _ := documents.Where(&models.Document{UserID: user.ID, FileID: document.FileID})

	if r.RowsAffected > 0 {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": "This file has already been shared", "suggestion": "Please verify the receiver or check the file"})
		return
	}

	//Checking if the receiver is the sender
	if user.ID == uint(json.ReceiverID) {
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

	//Testing if the receiver existe
	receiver := new(models.User)
	_, err = receiver.Get(&models.User{
		Base: base.Base{
			ID: uint(json.ReceiverID),
		},
	})

	if err != nil {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": err.Error(), "suggestion": "The receiver does not exist"})
		return
	}

	document.UserID = user.ID
	document.User = user
	document.Status = models.DocumentStatusIdle
	document.File = file

	_, err = document.Create()
	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error(), "suggestion": "We have issues creating your document"})
		return
	}

	share := &models.Share{
		DocumentID: document.ID,
		Document:   document,
		User:       user,
		UserID:     user.ID,
		Receiver:   receiver,
		ReceiverID: receiver.ID,
	}

	_, err = share.Create()
	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error(), "suggestion": "We have issues creating your document"})
		return
	}

	document.Status = models.DocumentStatusPending
	document.LastShare = &share.ID
	_, err = document.Save()
	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error(), "suggestion": "We have issues creating your document"})
		return
	}

	ctx.StopWithJSON(http.StatusCreated, share)

}

//MyDcoumentsController is used to get the documents of a user
// Those document created by the user
func MyDcoumentsController(ctx iris.Context) {
	documents := new(models.Documents)

	//Getting the current user
	user := ctx.Values().Get("user").(*models.User)

	//Getting the shares where the user is the receiver
	_, err := documents.Where(&models.Document{UserID: user.ID})

	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error()})
		return
	}

	ctx.StopWithJSON(http.StatusOK, documents)

}

//MyDcoumentsController is used to get the documents of a user
// Those document created by the user
func SetStatusToDocumentAs(ctx iris.Context) {
	documentStrID := ctx.Params().Get("id")

	docID, _ := strconv.Atoi(documentStrID)

	//Getting the current user
	user := ctx.Values().Get("user").(*models.User)

	//Check if the user has access to the document
	shares := new(models.Shares)
	res, _ := shares.Where(&models.Share{ReceiverID: user.ID, DocumentID: uint(docID)})
	if res.RowsAffected < 1 {

		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": "Sorry you don't have the right to change this document status"})
		return

	}

	//Getting the document and updating the document
	document := new(models.Document)
	_, err := document.Get(&models.Document{Base: base.Base{ID: uint(docID)}})

	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error()})
		return
	}

	status := ctx.Params().Get("status")
	switch status {
	case "complete":
		document.Status = models.DocumentStatusComplete
	case "cancel":
		document.Status = models.DocumentStatusCancel
	default:
		ctx.StopWithJSON(http.StatusNotFound, iris.Map{"error": "Only 'complet' and 'cancel' status are supported"})
		return
	}

	document.Save()

	ctx.StopWithJSON(http.StatusOK, document)

}

func GetAllUsers(ctx iris.Context) {
	users := new(models.Users)
	user := ctx.Values().Get("user").(*models.User)
	application.DB.Not(models.User{Base: base.Base{ID: user.ID}}).Find(users)

	ctx.StopWithJSON(200, users)

}

func GetDocument(ctx iris.Context) {
	id, _ := ctx.Params().GetInt("id")
	doc := new(models.Document)

	db := application.DB.Preload("User").Where(&models.Document{Base: base.Base{ID: uint(id)}}).First(doc)

	if db.Error != nil {
		ctx.StopWithJSON(http.StatusNotFound, iris.Map{"error": db.Error.Error()})
		return
	}
	ctx.StopWithJSON(200, doc)
}
