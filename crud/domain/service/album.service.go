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

type AlbumServiceInterface interface {
	Index() []serializer.Album
	Create(body AlbumServiceCreateParams) serializer.Album
}

type AlbumService struct {
	AlbumRepository repository.AlbumRepositoryInterface
}

func (s AlbumService) Index() []serializer.Album {
	albums := s.AlbumRepository.All()
	return serializer.AlbumSerializer{}.SerializeAlbums(albums)
}

func (s AlbumService) Create(body AlbumServiceCreateParams) serializer.Album {

	var album model.AlbumModel = model.AlbumModel{
		Title:  body.Title,
		Artist: body.Artist,
		Price:  body.Price,
	}

	s.AlbumRepository.Create(&album)

	return serializer.AlbumSerializer{}.SerializeAlbum(album)
}
