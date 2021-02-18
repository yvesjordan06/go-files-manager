package models

import (
	"files_manager/application"
	"files_manager/models/base"
	"files_manager/utilities"
	"gorm.io/gorm"

	"log"
)

//User of our application
type User struct {
	base.Base
	Username string `validate:"required" json:"username,omitempty" gorm:"unique"`
	Name     string `json:"name,omitempty" validate:"required"`
	Password string `json:"-"`
}

func (u *User) Get(conditions ...interface{}) (*gorm.DB, error) {
	query := application.DB.First(u, conditions...)
	return ParseTransactionWithError(query)
}

func (u *User) Create() (*gorm.DB, error) {
	u.Password = utilities.HashPassword(u.Password)
	query := application.DB.Create(u)
	return ParseTransactionWithError(query)
}

func (u *User) Save() (*gorm.DB, error) {
	query := application.DB.Save(u)
	return ParseTransactionWithError(query)
}

func (u *User) SetPassword(password string) (*gorm.DB, error) {
	u.Password = utilities.HashPassword(password)
	query := application.DB.Save(u)
	return ParseTransactionWithError(query)
}

type Users []User

func (u *Users) All() (*gorm.DB, error) {
	query := application.DB.Find(u)
	return ParseTransactionWithError(query)
}

func (u Users) Where(conditions ...interface{}) (*gorm.DB, error) {
	query := application.DB.Find(u, conditions...)
	return ParseTransactionWithError(query)
}

func init() {
	log.Println("Initializing model/user")
	application.DB.AutoMigrate(&User{})

	user := &User{Username: "hiro", Password: "hiro"}
	user.Create()
}
