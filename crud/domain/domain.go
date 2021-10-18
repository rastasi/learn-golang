package domain

import (
	"github.com/rastasi/learn-golang/crud/domain/repository"
	"github.com/rastasi/learn-golang/crud/domain/service"
	"github.com/rastasi/learn-golang/crud/lib/database"
)

type Domain struct {
	AlbumService service.AlbumService
}

func NewDomain() Domain {
	var AlbumService service.AlbumService = service.AlbumService{
		AlbumRepository: repository.AlbumRepository{
			DB: database.Init(),
		},
	}
	return Domain{
		AlbumService: AlbumService,
	}
}
