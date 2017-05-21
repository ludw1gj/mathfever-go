package router

import (
	"log"
	"net/http"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	err := notFoundTpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	err := errorTpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

