package repository

import (
	"github.com/rastasi/learn-golang/crud/domain/model"
	"github.com/rastasi/learn-golang/crud/lib/database"
)

type AlbumRepositoryInterface interface {
	All() []model.AlbumModel
	Create(album model.AlbumModel)
}

type AlbumRepository struct{}

func (r AlbumRepository) All() []model.AlbumModel {
	var albums []model.AlbumModel

	query := database.DB.Select("album_models.*").Group("album_models.id")

	if err := query.Find(&albums).Error; err != nil {
		panic(err)
	}

	return albums
}

func (r AlbumRepository) Create(album model.AlbumModel) {
	database.DB.Create(album)
}
