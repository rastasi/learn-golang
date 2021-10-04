package controller

import (
	"net/http"

	"github.com/rastasi/learn-golang/crud/app/model"
	"github.com/rastasi/learn-golang/crud/lib/utils"
)

var albums = []model.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, albums)
}

func PostAlbums(w http.ResponseWriter, r *http.Request) {
	var newAlbum model.Album

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	utils.RespondWithJSON(w, 201, newAlbum)
}
