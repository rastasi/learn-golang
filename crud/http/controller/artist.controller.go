package controller

import (
	"encoding/json"
	"net/http"

	"github.com/rastasi/learn-golang/crud/domain/service"
	"github.com/rastasi/learn-golang/crud/lib/utils"
)

type ArtistControllerInterface interface {
	Index(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type ArtistController struct {
	ArtistService service.ArtistServiceInterface
}

func (c ArtistController) Index(w http.ResponseWriter, r *http.Request) {
	artists := c.ArtistService.Index()
	utils.RespondWithJSON(w, 200, artists)
}

func (c ArtistController) Create(w http.ResponseWriter, r *http.Request) {
	var body service.ArtistServiceCreateParams

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	artist := c.ArtistService.Create(body)

	utils.RespondWithJSON(w, http.StatusCreated, artist)
}
