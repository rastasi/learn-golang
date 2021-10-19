package repository_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rastasi/learn-golang/crud/domain/repository"
	"github.com/rastasi/learn-golang/crud/lib/test"
)

func TestAlbumRepositoryCreate(t *testing.T) {
	suite := test.NewDBMockSuite()
	existsRows := sqlmock.NewRows([]string{"exists"}).AddRow(true)
	suite.Mock.ExpectQuery("SELECT album_models.* FROM `album_models` WHERE `album_models`.`deleted_at` IS NULL").
		WillReturnRows(existsRows)
	repo := repository.AlbumRepository{DB: suite.DB}
	repo.All()
}
