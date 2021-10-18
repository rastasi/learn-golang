package model

import (
	"time"

	"gorm.io/gorm"
)

type ArtistModel struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Active    int8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
