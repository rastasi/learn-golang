package test

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rastasi/learn-golang/crud/domain/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBMockSuite struct {
	DB   *gorm.DB
	Mock sqlmock.Sqlmock
}

func NewDBMockSuite() *DBMockSuite {
	var (
		db     *sql.DB
		gormdb *gorm.DB
		mock   sqlmock.Sqlmock
		err    error
	)

	db, mock, err = sqlmock.New()

	if err != nil {
		panic("SQL Mock error")
	}

	gormdb, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 db,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("GORM open error")
	}

	gormdb.AutoMigrate(
		model.AlbumModel{},
		model.ArtistModel{},
	)

	return &DBMockSuite{
		DB:   gormdb,
		Mock: mock,
	}
}
