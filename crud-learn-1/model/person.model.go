package model

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	ID        int64 `gorm:"primaryKey"`
	Name      string
	Age       int8
	Gender    int8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
