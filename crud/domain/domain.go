package domain

import (
	"github.com/rastasi/learn-golang/crud/domain/repository"
	"github.com/rastasi/learn-golang/crud/domain/service"
)

var AlbumService service.AlbumService = service.AlbumService{
	AlbumRepository: repository.AlbumRepository{},
}
