package service

import (
	"github.com/rastasi/learn-golang/crud/domain/model"
	"github.com/rastasi/learn-golang/crud/domain/repository"
	"github.com/rastasi/learn-golang/crud/domain/serializer"
)

type ArtistServiceCreateParams struct {
	ID     uint
	Name   string
	Active int8
}

type ArtistServiceInterface interface {
	Index() []serializer.Artist
	Create(body ArtistServiceCreateParams) serializer.Artist
}

type ArtistService struct {
	ArtistRepository repository.ArtistRepositoryInterface
}

func (s ArtistService) Index() []serializer.Artist {
	artists := s.ArtistRepository.All()
	return serializer.ArtistSerializer{}.SerializeArtists(artists)
}

func (s ArtistService) Create(body ArtistServiceCreateParams) serializer.Artist {

	var artist model.ArtistModel = model.ArtistModel{
		ID:     3244,
		Name:   body.Name,
		Active: body.Active,
	}

	s.ArtistRepository.Create(&artist)

	return serializer.ArtistSerializer{}.SerializeArtist(artist)
}
