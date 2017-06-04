package router

import (
	"flag"
	"net/http"

	"github.com/FriedPigeon/mathfever-go/controller"
	"github.com/gorilla/mux"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter()

	// Site Handler
	sc := controller.NewSiteController()
	Router.HandleFunc("/", sc.HomeHandler).Methods("GET")
	Router.HandleFunc("/about", sc.AboutHandler).Methods("GET")
	Router.HandleFunc("/help", sc.HelpHandler).Methods("GET")
	Router.HandleFunc("/privacy", sc.PrivacyHandler).Methods("GET")
	Router.HandleFunc("/terms", sc.TermsHandler).Methods("GET")
	Router.HandleFunc("/message-board", sc.MessageBoardHandler).Methods("GET")
	Router.HandleFunc("/networking/conversion-table", sc.ConversionTableHandler).Methods("GET")
	Router.HandleFunc("/{category}", sc.CategoryHandler).Methods("GET")
	Router.HandleFunc("/{category}/{calculation}", sc.CalculationHandler).Methods("GET")

	Router.NotFoundHandler = http.HandlerFunc(sc.NotFoundHandler)

	// API Handler
	ac := controller.NewApiController()
	apiRoute := Router.PathPrefix("/api").Subrouter()
	apiRoute.HandleFunc("/{category}/{calculation}", ac.DoCalculation).Methods("POST")

	apiRoute.NotFoundHandler = http.HandlerFunc(ac.NotFoundHandler)

	// Static Files Handler in Dev mode
	boolPtr := flag.Bool("dev", false, "Use in development")
	flag.Parse()
	if *boolPtr {
		Router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	}
}
