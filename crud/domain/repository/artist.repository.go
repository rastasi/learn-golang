package repository

import (
	"github.com/rastasi/learn-golang/crud/domain/model"
	"gorm.io/gorm"
)

type ArtistRepositoryInterface interface {
	All() []model.ArtistModel
	Create(Artist *model.ArtistModel)
}

type ArtistRepository struct {
	DB *gorm.DB
}

func (r ArtistRepository) All() []model.ArtistModel {
	var artists []model.ArtistModel

	query := r.DB.Select("artist_models.*").Group("artist_models.id")

	if err := query.Find(&artists).Error; err != nil {
		panic(err)
	}

	return artists
}

func (r ArtistRepository) Create(artist *model.ArtistModel) {
	r.DB.Create(artist)
}
