package database

import (
	"fmt"
	"os"

	"github.com/rastasi/learn-golang/crud/app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func migrate(db *gorm.DB) {
	db.AutoMigrate(&model.AlbumModel{})
}

type DbConnectData struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func getDbConnectData() DbConnectData {
	return DbConnectData{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
	}
}

func Setup() {
	c := getDbConnectData()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	migrate(db)
	DB = db
}
