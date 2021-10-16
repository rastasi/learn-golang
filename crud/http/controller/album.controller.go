package controller

import (
	"encoding/json"
	"net/http"

	"github.com/rastasi/learn-golang/crud/domain/service"
	"github.com/rastasi/learn-golang/crud/lib/utils"
)

type AlbumController struct{}

func (c AlbumController) Index(w http.ResponseWriter, r *http.Request) {
	albums := service.AlbumService{}.Index()
	utils.RespondWithJSON(w, 200, albums)
}

func (c AlbumController) Create(w http.ResponseWriter, r *http.Request) {
	var body service.AlbumServiceCreateParams

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	album := service.AlbumService{}.Create(body)

	utils.RespondWithJSON(w, http.StatusCreated, album)
}
