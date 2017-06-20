// Package router initialises a mux.Router instance and registers site/api routes including
// a file server for development use.
package route

import (
	"net/http"

	"flag"

	"github.com/FriedPigeon/mathfever-go/handler/api"
	"github.com/FriedPigeon/mathfever-go/handler/site"
	"github.com/gorilla/mux"
)

// Load returns a router instance with site/api routes initialised and a static file handler if dev flag
// is used.
func Load() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	// site routes
	r.HandleFunc("/", site.Home).Methods("GET")
	r.HandleFunc("/about", site.About).Methods("GET")
	r.HandleFunc("/help", site.Help).Methods("GET")
	r.HandleFunc("/privacy", site.Privacy).Methods("GET")
	r.HandleFunc("/terms", site.Terms).Methods("GET")
	r.HandleFunc("/message-board", site.MessageBoard).Methods("GET")
	r.HandleFunc("/category/networking/conversion-table", site.ConversionTable).Methods("GET")
	r.HandleFunc("/category/{category}", site.CategoryPage).Methods("GET")
	r.HandleFunc("/category/{category}/{calculation}", site.CalculationPage).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(site.NotFound)

	// static files handler in dev mode
	boolPtr := flag.Bool("dev", false, "Use in development")
	flag.Parse()
	if *boolPtr {
		fs := http.FileServer(http.Dir("./static"))
		r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	}

	// api routes
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/category/{category}/{calculation}", api.DoCalculation).Methods("POST")
	apiRouter.NotFoundHandler = http.HandlerFunc(api.NotFoundAPI)

	return r
}
