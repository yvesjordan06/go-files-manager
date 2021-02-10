package models

import (
	"files_manager/application"
	"files_manager/models/base"
	"gorm.io/gorm"

	"log"
)

//User of our application
type File struct {
	base.BaseUUID
	Name   string `json:"name,omitempty" validate:"required"`
	UserID uint   `json:"-"`
	User   *User  `json:"user,omitempty"`
	Size   int64  `json:"size"`
}

func (u *File) Get(conditions ...interface{}) (*gorm.DB, error) {
	query := application.DB.First(u, conditions...)
	return ParseTransactionWithError(query)
}

func (u *File) Create() (*gorm.DB, error) {
	query := application.DB.Create(u)
	return ParseTransactionWithError(query)
}

func (u *File) Save() (*gorm.DB, error) {
	query := application.DB.Save(u)
	return ParseTransactionWithError(query)
}

type Files []File

func (u *Files) All() (*gorm.DB, error) {
	query := application.DB.Find(u)
	return ParseTransactionWithError(query)
}

func (u *Files) Where(conditions ...interface{}) (*gorm.DB, error) {
	query := application.DB.Preload("User").Find(u, conditions...)
	return ParseTransactionWithError(query)
}

func init() {
	log.Println("Initializing model/file")
	application.DB.AutoMigrate(&File{})

}
