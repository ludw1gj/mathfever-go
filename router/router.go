// Package router initialises a mux.Router instance and registers Site/API routes including
// a file server for development use.
package router

import (
	"net/http"

	"flag"

	"github.com/FriedPigeon/mathfever-go/handler/api"
	"github.com/FriedPigeon/mathfever-go/handler/site"
	"github.com/gorilla/mux"
)

// Router to be used by http.ListenAndServe on which Site/API routes are registered to.
var Router *mux.Router

func init() {
	Router = mux.NewRouter().StrictSlash(true)

	// Site Routes
	Router.HandleFunc("/", site.Home).Methods("GET")
	Router.HandleFunc("/about", site.About).Methods("GET")
	Router.HandleFunc("/help", site.Help).Methods("GET")
	Router.HandleFunc("/privacy", site.Privacy).Methods("GET")
	Router.HandleFunc("/terms", site.Terms).Methods("GET")
	Router.HandleFunc("/message-board", site.MessageBoard).Methods("GET")
	Router.HandleFunc("/category/networking/conversion-table", site.ConversionTable).Methods("GET")
	Router.HandleFunc("/category/{category}", site.CategoryPage).Methods("GET")
	Router.HandleFunc("/category/{category}/{calculation}", site.CalculationPage).Methods("GET")
	Router.NotFoundHandler = http.HandlerFunc(site.NotFound)

	// Static Files Handler in Dev mode
	boolPtr := flag.Bool("dev", false, "Use in development")
	flag.Parse()
	if *boolPtr {
		fs := http.FileServer(http.Dir("./static"))
		Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	}

	// API Routes
	apiRouter := Router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/category/{category}/{calculation}", api.DoCalculation).Methods("POST")
	apiRouter.NotFoundHandler = http.HandlerFunc(api.NotFoundAPI)
}
