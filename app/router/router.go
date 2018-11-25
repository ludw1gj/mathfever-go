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

	siteTmpls := templates.CreateSiteTemplates()

	handlerWrapper := func(siteHandler controllers.SiteHandler) func(http.ResponseWriter,
		*http.Request) {
		return func(w http.ResponseWriter, r *http.Request) {
			siteHandler(w, r, siteTmpls)
		}
	}

	// site routes
	r.HandleFunc("/", handlerWrapper(controllers.HomePageHandler)).Methods("GET")
	r.HandleFunc("/about", handlerWrapper(controllers.AboutPageHandler)).Methods("GET")
	r.HandleFunc("/help", handlerWrapper(controllers.HelpPageHandler)).Methods("GET")
	r.HandleFunc("/privacy", handlerWrapper(controllers.PrivacyPageHandler)).Methods("GET")
	r.HandleFunc("/terms", handlerWrapper(controllers.TermsPageHandler)).Methods("GET")
	r.HandleFunc("/message-board", handlerWrapper(controllers.MessageBoardPageHandler)).Methods("GET")
	r.HandleFunc("/category/networking/conversion-table", handlerWrapper(controllers.ConversionTablePageHandler)).Methods("GET")
	r.HandleFunc("/category/{category}", handlerWrapper(controllers.CategoryPageHandler)).Methods("GET")
	r.HandleFunc("/category/{category}/{calculation}", handlerWrapper(controllers.CalculationPageHandler)).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(handlerWrapper(controllers.NotFoundPageHandler))

	// api routes
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/calculation", controllers.ProcessCalculationHandler).Methods("POST")
	apiRouter.NotFoundHandler = http.HandlerFunc(controllers.NotFoundHandler)

	return r
}
