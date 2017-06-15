package api

import (
	"encoding/json"
	"net/http"
)

type errorJson struct {
	Error string `json:"error"`
}

func NotFoundAPI(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(errorJson{"api route not found"})
}
