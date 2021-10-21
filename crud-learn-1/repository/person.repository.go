package repository

import (
	"github.com/rastasi/learn-golang/crud-learn-1/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDb() {
	var err error
	dsn := "root:toor@tcp(localhost:3306)/crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("MySQL connection error")
	}
	db.AutoMigrate(model.Person{})
}

func All() []model.Person {
	var people []model.Person
	query := db.Select("people.*")

	if err := query.Find(&people).Error; err != nil {
		panic("Insert error")
	}
	return people
}

func FindById(personId int) model.Person {
	var person model.Person
	db.Select(("people.*")).First(&person, personId)
	return person
}

func Create(person *model.Person) {
	db.Create(&person)
}

func Update(person *model.Person) {
	db.Updates(&person)
}

func Destroy(personId int) {
	db.Delete(&model.Person{}, personId)
}
