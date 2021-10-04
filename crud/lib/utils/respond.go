package utils

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func AddRouter(r *mux.Router, prefix string, subRouter mux.Router) {
	r.PathPrefix(prefix).Handler(http.StripPrefix(prefix, &subRouter))
}
