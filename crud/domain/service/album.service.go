package service

import (
	"github.com/rastasi/learn-golang/crud/domain/model"
	"github.com/rastasi/learn-golang/crud/domain/repository"
	"github.com/rastasi/learn-golang/crud/domain/serializer"
)

type AlbumServiceCreateParams struct {
	ID     uint
	Title  string
	Artist string
	Price  float64
}

type AlbumService struct{}

func (s AlbumService) Index() []serializer.Album {
	albums := repository.AlbumRepository{}.All()
	return serializer.AlbumSerializer{}.SerializeAlbums(albums)
}

func (s AlbumService) Create(body AlbumServiceCreateParams) serializer.Album {

	var album model.AlbumModel = model.AlbumModel{
		Title:  body.Title,
		Artist: body.Artist,
		Price:  body.Price,
	}

	repository.AlbumRepository{}.Create(album)

	return serializer.AlbumSerializer{}.SerializeAlbum(album)
}
