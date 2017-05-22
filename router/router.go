package router

import (
	"net/http"
	"github.com/gorilla/mux"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter()

	// Static Files Handler
	Router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// Site Handler
	Router.HandleFunc("/", indexHandler)
	Router.HandleFunc("/about", aboutHandler)
	Router.HandleFunc("/help", helpHandler)
	Router.HandleFunc("/privacy", privacyHandler)
	Router.HandleFunc("/terms", termsHandler)
	Router.HandleFunc("/message-board", messageBoardHandler)
	Router.HandleFunc("/networking/conversion-table", conversionTableHandler)
	Router.HandleFunc("/{category}", categoriesHandler)
	Router.HandleFunc("/{category}/{calculation}", calculationsHandler)

	// API Handler
	Router.HandleFunc("/api/{category}/{calculation}", calculationsAPIHandler).Methods("POST")

	// Not Found Handler
	Router.NotFoundHandler = http.HandlerFunc(notFoundHandler)
}