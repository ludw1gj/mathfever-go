package router

import (
	"net/http"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	renderTemplate(w, r, notFoundTpml, "base.gohtml", nil)
}

func serverErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	renderTemplate(w, r, errorTpml, "base.gohtml", nil)
}
