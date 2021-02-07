package application

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	/// DB application Gorm Database
	DB  = new(gorm.DB)
	err error
)

func init() {
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

}
