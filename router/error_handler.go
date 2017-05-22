package router

import (
	"log"
	"net/http"
)

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	err := notFoundTpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	err := errorTpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

