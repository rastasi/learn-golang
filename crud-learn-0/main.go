package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func connectDb() {
	var err error
	dsn := "root:toor@tcp(localhost:3306)/crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("MySQL connection error")
	}
	db.AutoMigrate(Person{})
}

type Person struct {
	ID        int64 `gorm:"primaryKey"`
	Name      string
	Age       int8
	Gender    int8
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type PersonSerializer struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Age    int8   `json:"age"`
	Gender string `json:"gender"`
}

func serialize(person Person) PersonSerializer {
	var gender string

	if person.Gender == 1 {
		gender = "male"
	} else {
		gender = "female"
	}

	return PersonSerializer{
		ID:     person.ID,
		Name:   person.Name,
		Age:    person.Age,
		Gender: gender,
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var people []Person

	query := db.Select("people.*")

	if err := query.Find(&people).Error; err != nil {
		panic("Insert error")
	}

	var serialized []PersonSerializer

	for _, person := range people {
		serialized = append(serialized, serialize(person))
	}

	respondWithJSON(w, http.StatusOK, serialized)
}

func showHandler(w http.ResponseWriter, r *http.Request) {
	var person Person
	params := mux.Vars(r)
	db.Select(("people.*")).First(&person, params["personId"])
	respondWithJSON(w, http.StatusOK, serialize(person))
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	var person Person

	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Create(&person)
	respondWithJSON(w, http.StatusCreated, serialize(person))
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
}

func destroyHandler(w http.ResponseWriter, r *http.Request) {
}

func muxServer() {
	router := mux.NewRouter()

	router.HandleFunc("/people/", indexHandler).Methods("GET")
	router.HandleFunc("/people/", createHandler).Methods("POST")
	router.HandleFunc("/people/{personId}", showHandler).Methods("GET")
	router.HandleFunc("/people/{personId}", updateHandler).Methods("PUT")
	router.HandleFunc("/people/{personId}", destroyHandler).Methods("DELETE")

	http.ListenAndServe(":8090", router)

}

// CREATE - CREANTE
// READ   - INDEX
// READ   - SHOW
// UPDATE - UPDATE
// DELETE - DESTROY

func main() {
	connectDb()
	muxServer()
}
