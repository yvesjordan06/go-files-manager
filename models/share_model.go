package models

import (
	"files_manager/application"
	"files_manager/models/base"
	"log"

	"gorm.io/gorm"
)

//A Share permit a user to share an existing document from one user to another
type Share struct {
	base.Base
	Document   *Document `json:"document,omitempty"`
	DocumentID uint      `json:"file_id,omitempty" validate:"required"`
	UserID     uint      `json:"user_id"`
	User       *User     `json:"expeditor,omitempty"`
	ReceiverID uint      `json:"receiver_id,omitempty" validate:"required"`
	Receiver   *User     `json:"receiver,omitempty"`
}

//
func (u *Share) Get(conditions ...interface{}) (*gorm.DB, error) {
	query := application.DB.Preload("User").Preload("Receiver").First(u, conditions...)
	return ParseTransactionWithError(query)
}

func (u *Share) Create() (*gorm.DB, error) {
	query := application.DB.Create(u)
	return ParseTransactionWithError(query)
}

func (u *Share) Save() (*gorm.DB, error) {
	query := application.DB.Save(u)
	return ParseTransactionWithError(query)
}

type Shares []Share

func (u *Shares) All() (*gorm.DB, error) {
	query := application.DB.Find(u)
	return ParseTransactionWithError(query)
}

func (u *Shares) Where(conditions ...interface{}) (*gorm.DB, error) {
	query := application.DB.Preload("User").Preload("Document").Find(u, conditions...)
	return ParseTransactionWithError(query)
}

func init() {
	log.Println("Initializing model/share")
	application.DB.AutoMigrate(&Share{})

}
