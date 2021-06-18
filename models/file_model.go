package models

import (
	"files_manager/application"
	"files_manager/models/base"
	"log"

	"gorm.io/gorm"
)

//User of our application
type File struct {
	base.BaseUUID
	Name   string `json:"name,omitempty" validate:"required"`
	UserID uint   `json:"-"`
	User   *User  `json:"user,omitempty"`
	Size   int64  `json:"size"`
}

//Get a file
func (u *File) Get(conditions ...interface{}) (*gorm.DB, error) {
	query := application.DB.First(u, conditions...)
	return ParseTransactionWithError(query)
}

//Creates a file
func (u *File) Create() (*gorm.DB, error) {
	query := application.DB.Create(u)
	return ParseTransactionWithError(query)
}

//Saves a file
func (u *File) Save() (*gorm.DB, error) {
	query := application.DB.Save(u)
	return ParseTransactionWithError(query)
}

// List of files
type Files []File

// Get all files in a list
func (u *Files) All() (*gorm.DB, error) {
	query := application.DB.Find(u)
	return ParseTransactionWithError(query)
}

// Get a list of files where `conditions`
func (u *Files) Where(conditions ...interface{}) (*gorm.DB, error) {
	query := application.DB.Preload("User").Find(u, conditions...)
	return ParseTransactionWithError(query)
}

// Init function
func init() {
	log.Println("Initializing model/file")
	application.DB.AutoMigrate(&File{})

}
