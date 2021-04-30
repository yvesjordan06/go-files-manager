package models

import (
	"files_manager/application"
	"files_manager/models/base"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"log"
)

const (
	DocumentStatusPending   = "pending"
	DocumentStatusRead      = "read"
	DocumentStatusDeleted   = "deleted"
	DocumentStatusForwarded = "forwarded"
	DocumentStatusComplete  = "complete"
	DocumentStatusCancel    = "cancel"
)

//User of our application
type Document struct {
	base.Base
	Title     string `json:"title,omitempty" validate:"required"`
	Reference string `json:"reference,omitempty" validate:"required"`
	Object    string `json:"object,omitempty" validate:"required"`
	//Status    string    `json:"status,omitempty"`
	File   *File     `json:"file,omitempty"`
	FileID uuid.UUID `json:"file_id,omitempty" validate:"required"`
	//ReceivedAt      *time.Time `json:"received_at"`
	UserID uint  `json:"-"`
	User   *User `json:"user,omitempty"`
	//ReceiverID      uint       `json:"receiver_id,omitempty" validate:"required"`
	//Receiver        *User      `json:"receiver,omitempty"`
	//AssignedID      *uint      `json:"assigned_id"`
	//Assigned        *User      `json:"assigned_to,omitempty"`
	Comments *Comments `json:"comments"`
	//UserDeleted     bool       `json:"-" gorm:"default:false"`
	//ReceiverDeleted bool       `json:"-" gorm:"default:false"`
}

func (u *Document) Get(conditions ...interface{}) (*gorm.DB, error) {
	query := application.DB.Preload("User").Preload("File").First(u, conditions...)
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
	query := application.DB.Preload("User").Preload("File").Order("created_at desc").Find(u, conditions...)
	return ParseTransactionWithError(query)
}

func init() {
	log.Println("Initializing model/document")
	_ = application.DB.AutoMigrate(&Document{})

}
