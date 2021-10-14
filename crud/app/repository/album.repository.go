package repository

import (
	"github.com/rastasi/learn-golang/crud/app/model"
	"github.com/rastasi/learn-golang/crud/lib/database"
)

func GetAlbums() []model.AlbumModel {
	var albums []model.AlbumModel

	query := database.DB.Select("album_models.*").Group("album_models.id")

	if err := query.Find(&albums).Error; err != nil {
		panic(err)
	}

	return albums
}

func PostAlbums(album model.AlbumModel) {
	database.DB.Create(album)
}
