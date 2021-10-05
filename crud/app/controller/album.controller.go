package controller

import (
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

func PostAlbums(w http.ResponseWriter, r *http.Request) {
	var album model.AlbumModel

	utils.RespondWithJSON(w, 201, serializer.SerializeAlbum(album))
}
