package controller

import (
	"encoding/json"
	"net/http"

	"github.com/rastasi/learn-golang/crud/app/model"
	"github.com/rastasi/learn-golang/crud/app/repository"
	"github.com/rastasi/learn-golang/crud/app/serializer"
	"github.com/rastasi/learn-golang/crud/lib/utils"
)

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	albums := repository.GetAlbums()
	utils.RespondWithJSON(w, 200, serializer.SerializeAlbums(albums))
}

type PostAlbumsRequestPayload struct {
	ID     uint
	Title  string
	Artist string
	Price  float64
}

func PostAlbums(w http.ResponseWriter, r *http.Request) {
	var body PostAlbumsRequestPayload

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var album model.AlbumModel = model.AlbumModel{
		Title:  body.Title,
		Artist: body.Artist,
		Price:  body.Price,
	}

	repository.PostAlbums(album)

	utils.RespondWithJSON(w, http.StatusCreated, serializer.SerializeAlbum(album))
}
