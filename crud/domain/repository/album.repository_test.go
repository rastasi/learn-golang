package repository_test

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rastasi/learn-golang/crud/domain/repository"
)

func TestAlbumRepositoryCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("Mock error")
	}
	repo := repository.AlbumRepository{
		DB: db,
	}
	fmt.Println(repo)
}
