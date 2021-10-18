package domain

import (
	"github.com/rastasi/learn-golang/crud/domain/repository"
	"github.com/rastasi/learn-golang/crud/domain/service"
	"github.com/rastasi/learn-golang/crud/lib/database"
)

type Domain struct {
	AlbumService  service.AlbumService
	ArtistService service.ArtistService
}

func NewDomain() Domain {
	DB := database.Init()
	return Domain{
		AlbumService: service.AlbumService{
			AlbumRepository: repository.AlbumRepository{
				DB: DB,
			},
		},
		ArtistService: service.ArtistService{
			ArtistRepository: repository.ArtistRepository{
				DB: DB,
			},
		},
	}
}
