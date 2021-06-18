package models

import (
	"files_manager/application"
	"files_manager/models/base"
	"files_manager/utilities"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"log"
)

//Token for user login
type Token struct {
	base.Base
	Token    string `json:"token"`
	Disabled bool   `json:"-"`
	UserID   uint   `json:"-"`
	User     User   `json:"user"`
}

func (u *Token) Get(conditions ...interface{}) (*gorm.DB, error) {
	query := application.DB.Preload(clause.Associations).First(u, conditions...)
	return ParseTransactionWithError(query)
}

func (u *Token) Create() (*gorm.DB, error) {
	var user User
	_, err := user.Get(u.UserID)
	if err != nil {
		return nil, err
	}

	u.Token = utilities.HashPassword(time.Now().String())
	query := application.DB.Create(u)
	return ParseTransactionWithError(query)
}

func (u *Token) Save() (*gorm.DB, error) {
	query := application.DB.Model(u).Select("disabled").Updates(u)
	return ParseTransactionWithError(query)
}

type Tokens []Token

func (u *Tokens) All() (*gorm.DB, error) {
	query := application.DB.Find(u)
	return ParseTransactionWithError(query)
}

func (u *Tokens) Where(conditions ...interface{}) (*gorm.DB, error) {
	query := application.DB.Find(u, conditions...)
	return ParseTransactionWithError(query)
}

func init() {
	log.Println("Initializing model/token")
	application.DB.AutoMigrate(&Token{})
}
