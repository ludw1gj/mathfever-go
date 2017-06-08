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

	// Site Router
	siteCtrl := controller.NewSiteController()
	Router.HandleFunc("/", siteCtrl.HomeHandler).Methods("GET")
	Router.HandleFunc("/about", siteCtrl.AboutHandler).Methods("GET")
	Router.HandleFunc("/help", siteCtrl.HelpHandler).Methods("GET")
	Router.HandleFunc("/privacy", siteCtrl.PrivacyHandler).Methods("GET")
	Router.HandleFunc("/terms", siteCtrl.TermsHandler).Methods("GET")
	Router.HandleFunc("/message-board", siteCtrl.MessageBoardHandler).Methods("GET")
	Router.HandleFunc("/networking/conversion-table", siteCtrl.ConversionTableHandler).Methods("GET")
	Router.HandleFunc("/{category}", siteCtrl.CategoryHandler).Methods("GET")
	Router.HandleFunc("/{category}/{calculation}", siteCtrl.CalculationHandler).Methods("GET")

	Router.NotFoundHandler = http.HandlerFunc(siteCtrl.NotFoundHandler)

	// API Router
	apiRouter := Router.PathPrefix("/api").Subrouter()

	categCtrl := controller.NewCategoryController()
	apiRouter.HandleFunc("/{category}", categCtrl.CategoryHandler).Methods("POST")

	calcCtrl := controller.NewCalculationController()
	apiRouter.HandleFunc("/{category}/{calculation}", calcCtrl.CalculationHandler).Methods("POST")

	apiRouter.NotFoundHandler = http.HandlerFunc(siteCtrl.NotFoundAPIHandler)

	// Static Files Handler in Dev mode
	boolPtr := flag.Bool("dev", false, "Use in development")
	flag.Parse()
	if *boolPtr {
		Router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	}
}
