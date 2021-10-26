package repository

import (
	"github.com/rastasi/learn-golang/crud-learn-2/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type PersonRepositoryInterface interface {
	All() []model.Person
	FindById(personId int) model.Person
	Create(person *model.Person)
	Update(person *model.Person)
	Destroy(personId int)
}

type PersonRepository struct {
}

func ConnectDb() {
	var err error
	dsn := "root:toor@tcp(localhost:3306)/crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("MySQL connection error")
	}
	db.AutoMigrate(model.Person{})
}

func (r PersonRepository) All() []model.Person {
	var people []model.Person
	query := db.Select("people.*")

	if err := query.Find(&people).Error; err != nil {
		panic("Insert error")
	}
	return people
}

func (r PersonRepository) FindById(personId int) model.Person {
	var person model.Person
	db.Select(("people.*")).First(&person, personId)
	return person
}

func (r PersonRepository) Create(person *model.Person) {
	db.Create(&person)
}

func (r PersonRepository) Update(person *model.Person) {
	db.Updates(&person)
}

func (r PersonRepository) Destroy(personId int) {
	db.Delete(&model.Person{}, personId)
}
