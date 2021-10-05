package model

import (
	"time"

	"gorm.io/gorm"
)

type AlbumModel struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Artist    string
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
