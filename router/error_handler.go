package router

import (
	"log"
	"net/http"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	err := notFoundTpl.ExecuteTemplate(w, "base.gohtml", nil)
	if err != nil {
		log.Println(err)
		errorHandler(w, r)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	err := errorTpl.ExecuteTemplate(w, "base.gohtml", nil)
	if err != nil {
		log.Println(err)
	}
}
