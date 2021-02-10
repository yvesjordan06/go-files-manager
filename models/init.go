package models

// Use this files to register your packaged models
import (
	"gorm.io/gorm"
)

type SingleDBQuery interface {
	Create() (*gorm.DB, error)
	Save() (*gorm.DB, error)
	Get(conditions ...interface{}) (*gorm.DB, error)
}

type GroupDBQuery interface {
	All() (*gorm.DB, error)
	Where(conditions ...interface{}) (*gorm.DB, error)
}

func ParseTransactionWithError(db *gorm.DB) (*gorm.DB, error) {
	if db.Error != nil {
		return db, db.Error
	}
	return db, nil
}
