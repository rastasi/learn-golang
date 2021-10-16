package controller

import (
	"encoding/json"
	"net/http"

	"github.com/rastasi/learn-golang/crud/domain/service"
	"github.com/rastasi/learn-golang/crud/lib/utils"
)

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	albums := service.GetAlbums()
	utils.RespondWithJSON(w, 200, albums)
}

func PostAlbum(w http.ResponseWriter, r *http.Request) {
	var body service.PostAlbumServiceParams

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	album := service.PostAlbum(body)

	utils.RespondWithJSON(w, http.StatusCreated, album)
}
