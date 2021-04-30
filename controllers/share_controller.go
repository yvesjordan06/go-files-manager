package controllers

import (
	"files_manager/models"
	"files_manager/models/base"
	"github.com/kataras/iris/v12"
	"net/http"
)

/// ForwardShare gets a share from the given id in order to share the document
/// Makes sure the share exists
/// Looks for other share of the document where the senders and receiver are same
/// If finds it, reset the delete and status of the share
/// Else creates a new share from the old one, then set the users
func ForwardShare(ctx iris.Context) {

}

// GetShared gets the list of all shares where the current user is Receiver or Sender
// Makes use of Query Parameters
// sender=1 => Where i am the sender
// receiver=1 => Where i am the receiver
// status={string} => Where status is {string}
func GetShared(ctx iris.Context) {

}

// OpenShare get a single share by ID
// Change the status of the share if the user is receiver
// Make sur the user is either sender or receiver else return not found
func OpenShare(ctx iris.Context) {

}

// GetDocumentShares get the list of all share for a specific document
// Only the owner of the document can get this resource
func GetDocumentShares(ctx iris.Context) {
	shares := new(models.Shares)
	documentID, _ := ctx.Params().GetInt("id")
	_, _ = shares.Where(&models.Share{DocumentID: uint(documentID)})

	ctx.StopWithJSON(http.StatusOK, shares)
}

// GetDocumentShares get the list of all share for a specific document
// Only the owner of the document can get this resource
func ShareDocument(ctx iris.Context) {
	share := new(models.Share)
	user := ctx.Values().Get("user").(*models.User)
	documentID, _ := ctx.Params().GetInt("id")

	data := new(struct {
		To int `json:"to" validate:"required"`
	})

	err := ctx.ReadJSON(data)
	if err != nil {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": err.Error(), "suggestion": "Verify your input and try again."})
		return
	}

	receiver := new(models.User)
	_, err = receiver.Get(&models.User{
		Base: base.Base{
			ID: uint(data.To),
		},
	})

	if err != nil {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": err.Error(), "suggestion": "The receiver does not exist"})
		return
	}

	duplicate := new(models.Share)
	query, _ := duplicate.Get(&models.Share{SenderID: user.ID, ReceiverID: receiver.ID, DocumentID: uint(documentID)})

	if query.RowsAffected > 0 {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": "Duplicate", "suggestion": "This document has already been shared"})
		return
	}

	share.DocumentID = uint(documentID)
	share.SenderID = user.ID
	share.Sender = user
	share.ReceiverID = receiver.ID
	share.Receiver = receiver
	share.Status = models.DocumentStatusPending
	_, err = share.Create()
	if err != nil {
		ctx.StopWithJSON(http.StatusInternalServerError, iris.Map{"error": err.Error(), "suggestion": "Couldn't share the document"})
		return
	}
	ctx.StopWithJSON(http.StatusOK, share)
}

// CompleteShareStatus get a share by id and set it's status to complete
// Only the owner of the receiver of the share can do this action
func CompleteShareStatus(ctx iris.Context) {

}

func CancelShareStatus(ctx iris.Context) {

}

// DeleteShare get a share by id and set it's status to delete and soft delete it
// Only the owner or the receiver of the share can do this action
func DeleteShare(ctx iris.Context) {
	share := new(models.Share)
	shareID, _ := ctx.Params().GetInt("id")
	user := ctx.Values().Get("user").(*models.User)

	_, err := share.Get(&models.Share{
		Base: base.Base{
			ID: uint(shareID),
		},
	})

	if err != nil {
		ctx.StopWithJSON(http.StatusBadRequest, iris.Map{"error": err.Error(), "suggestion": "The document doesn't exist"})
		return
	}

	if share.SenderID == user.ID {
		share.SenderDeleted = true
	}

	if share.ReceiverID == user.ID {
		share.ReceiverDeleted = true
		share.Status = models.DocumentStatusDeleted
	}

	_, _ = share.Save()

	ctx.StopWithStatus(http.StatusOK)

}
