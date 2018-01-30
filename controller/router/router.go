// Package router initialises a mux.Router instance and registers site/api routes including
// a file server for development use.
package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Load returns a router instance with site/api routes initialised and a static file handler if dev flag
// is used.
func Load() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	// site routes
	r.HandleFunc("/", getHomePage).Methods("GET")
	r.HandleFunc("/about", getAboutPage).Methods("GET")
	r.HandleFunc("/help", getHelpPage).Methods("GET")
	r.HandleFunc("/privacy", getPrivacyPage).Methods("GET")
	r.HandleFunc("/terms", getTermsPage).Methods("GET")
	r.HandleFunc("/message-board", getMessageBoardPage).Methods("GET")
	r.HandleFunc("/category/networking/conversion-table", getConversionTablePage).Methods("GET")
	r.HandleFunc("/category/{category}", getCategoryPage).Methods("GET")
	r.HandleFunc("/category/{category}/{calculation}", getCalculationPage).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(notFoundPage)

	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	// api routes
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/category/{category}/{calculation}", processCalculation).Methods("POST")
	apiRouter.NotFoundHandler = http.HandlerFunc(apiNotFound)

	return r
}
