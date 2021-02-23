package models

import (
	"files_manager/application"
	"files_manager/models/base"
	"gorm.io/gorm"
	"log"
)

type Share struct {
	base.Base
	Document   *User `json:"document,omitempty"`
	DocumentID *uint `json:"-"`
	SenderID   uint  `json:"-"`
	Sender     *User `json:"sender,omitempty"`
	ReceiverID uint  `json:"receiver_id" validate:"required"`
	Receiver   *User `json:"receiver,omitempty"`
}

type Shares []Share

func (u *Share) Get(conditions ...interface{}) (*gorm.DB, error) {
	query := application.DB.Preload("Sender").Preload("Receiver").First(u, conditions...)
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

func (u *Shares) All() (*gorm.DB, error) {
	query := application.DB.Find(u)
	return ParseTransactionWithError(query)
}

func (u *Shares) Where(conditions ...interface{}) (*gorm.DB, error) {
	query := application.DB.Preload("Sender").Preload("Receiver").Order("created_at desc").Find(u, conditions...)
	return ParseTransactionWithError(query)
}

func init() {
	log.Println("Initializing model/comment")
	application.DB.AutoMigrate(&Share{})

}
