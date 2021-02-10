package base

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (b Base) getID() uint {
	return b.ID
}
