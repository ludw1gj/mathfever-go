package router

import (
	"net/http"
	"github.com/gorilla/mux"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter()

	// Static Files Handler
	Router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	// Site Handler
	Router.HandleFunc("/", IndexHandler)
	Router.HandleFunc("/about", AboutHandler)
	Router.HandleFunc("/help", HelpHandler)
	Router.HandleFunc("/privacy", PrivacyHandler)
	Router.HandleFunc("/terms", TermsHandler)
	Router.HandleFunc("/message-board", MessageBoardHandler)
	Router.HandleFunc("/networking/conversion-table", ConversionTableHandler)
	Router.HandleFunc("/{category}", CategoriesHandler)
	Router.HandleFunc("/{category}/{calculation}", CalculationsHandler)

	// API Handler
	Router.HandleFunc("/api/{category}/{calculation}", CalculationsAPIHandler).Methods("POST")

	// Not Found Handler
	Router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
}