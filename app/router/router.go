// Package router initialises a mux.Router instance and registers site/api routes including
// a file server for development use.
package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ludw1gj/mathfever-go/app/controllers"
	"github.com/ludw1gj/mathfever-go/app/templates"
)

// Load returns a router instance with routes and a file server.
func Load() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	fs := http.FileServer(http.Dir("./public"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	// controllers
	sc := controllers.SiteController{Tmpls: templates.CreateSiteTemplates()}

	// site routes
	r.HandleFunc("/", sc.HomePageHandler).Methods("GET")
	r.HandleFunc("/about", sc.AboutPageHandler).Methods("GET")
	r.HandleFunc("/help", sc.HelpPageHandler).Methods("GET")
	r.HandleFunc("/privacy", sc.PrivacyPageHandler).Methods("GET")
	r.HandleFunc("/terms", sc.TermsPageHandler).Methods("GET")
	r.HandleFunc("/message-board", sc.MessageBoardPageHandler).Methods("GET")
	r.HandleFunc("/category/networking/conversion-table", sc.ConversionTablePageHandler).Methods("GET")
	r.HandleFunc("/category/{category}", sc.CategoryPageHandler).Methods("GET")
	r.HandleFunc("/category/{category}/{calculation}", sc.CalculationPageHandler).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(sc.NotFoundPageHandler)

	// api routes
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/calculation", controllers.ProcessCalculationHandler).Methods("POST")
	apiRouter.NotFoundHandler = http.HandlerFunc(controllers.NotFoundHandler)

	return r
}
