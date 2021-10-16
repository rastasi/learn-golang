package service

import (
	"github.com/rastasi/learn-golang/crud/domain/model"
	"github.com/rastasi/learn-golang/crud/domain/repository"
	"github.com/rastasi/learn-golang/crud/domain/serializer"
)

func GetAlbums() []serializer.AlbumSerializer {
	albums := repository.GetAlbums()
	return serializer.SerializeAlbums(albums)
}

type PostAlbumServiceParams struct {
	ID     uint
	Title  string
	Artist string
	Price  float64
}

func PostAlbum(body PostAlbumServiceParams) serializer.AlbumSerializer {

	var album model.AlbumModel = model.AlbumModel{
		Title:  body.Title,
		Artist: body.Artist,
		Price:  body.Price,
	}

	repository.PostAlbum(album)

	return serializer.SerializeAlbum(album)
}
