// Package router initialises a mux.Router instance and registers site/api routes including
// a file server for development use.
package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/robertjeffs/mathfever-go/app/controllers"
)

// Load returns a router instance with site/api routes initialised and a public file handler if dev flag
// is used.
func Load() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	// controllers
	sc := controllers.NewSiteController()
	mc := controllers.NewMathAPIController()

	fs := http.FileServer(http.Dir("./public"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	// site routes
	r.HandleFunc("/", sc.GetHomePage).Methods("GET")
	r.HandleFunc("/about", sc.GetAboutPage).Methods("GET")
	r.HandleFunc("/help", sc.GetHelpPage).Methods("GET")
	r.HandleFunc("/privacy", sc.GetPrivacyPage).Methods("GET")
	r.HandleFunc("/terms", sc.GetTermsPage).Methods("GET")
	r.HandleFunc("/message-board", sc.GetMessageBoardPage).Methods("GET")
	r.HandleFunc("/category/networking/conversion-table", sc.GetConversionTablePage).Methods("GET")
	r.HandleFunc("/category/{category}", sc.GetCategoryPage).Methods("GET")
	r.HandleFunc("/category/{category}/{calculation}", sc.GetCalculationPage).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(sc.NotFoundPage)

	// api routes
	apiRouter := r.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/category/{category}/{calculation}", mc.ProcessCalculation).Methods("POST")
	apiRouter.NotFoundHandler = http.HandlerFunc(mc.APINotFound)

	return r
}
