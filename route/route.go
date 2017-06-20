// Package route initialises a mux.Router instance and registers site/api routes including
// a file server for development use.
package route

import (
	"net/http"

	"flag"

	"github.com/gorilla/mux"
)

// Load returns a router instance with site/api routes initialised and a static file handler if dev flag
// is used.
func Load() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	// site routes
	r.HandleFunc("/", getHome).Methods("GET")
	r.HandleFunc("/about", getAbout).Methods("GET")
	r.HandleFunc("/help", getHelp).Methods("GET")
	r.HandleFunc("/privacy", getPrivacy).Methods("GET")
	r.HandleFunc("/terms", getTerms).Methods("GET")
	r.HandleFunc("/message-board", getMessageBoard).Methods("GET")
	r.HandleFunc("/category/networking/conversion-table", getConversionTable).Methods("GET")
	r.HandleFunc("/category/{category}", getCategoryPage).Methods("GET")
	r.HandleFunc("/category/{category}/{calculation}", getCalculationPage).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(notFound)

	// static files handler in dev mode
	boolPtr := flag.Bool("dev", false, "Use in development")
	flag.Parse()
	if *boolPtr {
		fs := http.FileServer(http.Dir("./static"))
		r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	}

	// api routes
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/category/{category}/{calculation}", doCalculation).Methods("POST")
	apiRouter.NotFoundHandler = http.HandlerFunc(notFoundAPI)

	return r
}
