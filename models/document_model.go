package models

import (
	"files_manager/application"
	"files_manager/models/base"
	"log"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// A document is the entity that is shared!
// It contains the original file

const (
	DocumentStatusIdle     = "idl"      // When the document is idle, no share
	DocumentStatusPending  = "Pending"  // When the document has been shared and it's pending
	DocumentStatusComplete = "complete" // When the document has been completed
	DocumentStatusCancel   = "cancel"   // When the document has been cancelled
)

//User of our application
type Document struct {
	base.Base
	Title     string    `json:"title,omitempty" validate:"required"` //The tilte of the documents
	Reference string    `json:"reference,omitempty" validate:"required"`
	Object    string    `json:"object,omitempty" validate:"required"`
	Status    string    `json:"status,omitempty"` // The status of the document
	File      *File     `json:"file,omitempty"`
	FileID    uuid.UUID `json:"file_id,omitempty" validate:"required"`
	UserID    uint      `json:"user_id"`
	User      *User     `json:"expeditor,omitempty"`
	LastShare *uint     `json:"last_share"`
}

//
func (u *Document) Get(conditions ...interface{}) (*gorm.DB, error) {
	query := application.DB.First(u, conditions...)
	return ParseTransactionWithError(query)
}

func (u *Document) Create() (*gorm.DB, error) {
	query := application.DB.Create(u)
	return ParseTransactionWithError(query)
}

func (u *Document) Save() (*gorm.DB, error) {
	query := application.DB.Save(u)
	return ParseTransactionWithError(query)
}

type Documents []Document

func (u *Documents) All() (*gorm.DB, error) {
	query := application.DB.Find(u)
	return ParseTransactionWithError(query)
}

func (u *Documents) Where(conditions ...interface{}) (*gorm.DB, error) {
	query := application.DB.Preload("User").Find(u, conditions...)
	return ParseTransactionWithError(query)
}

func init() {
	log.Println("Initializing model/document")
	application.DB.AutoMigrate(&Document{})

}
