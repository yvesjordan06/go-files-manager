package models

import (
	"files_manager/application"
	"files_manager/models/base"
	"gorm.io/gorm"
	"log"
)

type Comment struct {
	base.Base
	Text       string `json:"text,omitempty" validate:"required"`
	UserID     uint   `json:"-"`
	User       *User  `json:"user,omitempty"`
	DocumentID *uint  `json:"-"`
}

type Comments []Comment

func (u *Comment) Get(conditions ...interface{}) (*gorm.DB, error) {
	query := application.DB.Preload("User").First(u, conditions...)
	return ParseTransactionWithError(query)
}

func (u *Comment) Create() (*gorm.DB, error) {
	query := application.DB.Create(u)
	return ParseTransactionWithError(query)
}

func (u *Comment) Save() (*gorm.DB, error) {
	query := application.DB.Save(u)
	return ParseTransactionWithError(query)
}

func (u *Comments) All() (*gorm.DB, error) {
	query := application.DB.Find(u)
	return ParseTransactionWithError(query)
}

func (u *Comments) Where(conditions ...interface{}) (*gorm.DB, error) {
	query := application.DB.Preload("User").Order("created_at desc").Find(u, conditions...)
	return ParseTransactionWithError(query)
}

func init() {
	log.Println("Initializing model/comment")
	application.DB.AutoMigrate(&Comment{})

}
