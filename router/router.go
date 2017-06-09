package router

import (
	"flag"
	"net/http"

	"github.com/FriedPigeon/mathfever-go/controller"
	"github.com/gorilla/mux"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter().StrictSlash(true)

	// API Router
	apiRouter := Router.PathPrefix("/api").Subrouter()

	apiCtrl := controller.NewApiController()
	apiRouter.HandleFunc("/", apiCtrl.GetCategories).Methods("GET")
	apiRouter.HandleFunc("/{category}", apiCtrl.GetCategory).Methods("GET")
	apiRouter.HandleFunc("/{category}/{calculation}", apiCtrl.GetCalculation).Methods("GET")
	apiRouter.HandleFunc("/{category}/{calculation}", apiCtrl.DoCalculation).Methods("POST")

	apiRouter.NotFoundHandler = http.HandlerFunc(apiCtrl.NotFoundAPI)

	// Site Router
	siteCtrl := controller.NewSiteController()
	Router.HandleFunc("/", siteCtrl.Home).Methods("GET")
	Router.HandleFunc("/about", siteCtrl.About).Methods("GET")
	Router.HandleFunc("/help", siteCtrl.Help).Methods("GET")
	Router.HandleFunc("/privacy", siteCtrl.Privacy).Methods("GET")
	Router.HandleFunc("/terms", siteCtrl.Terms).Methods("GET")
	Router.HandleFunc("/message-board", siteCtrl.MessageBoard).Methods("GET")
	Router.HandleFunc("/networking/conversion-table", siteCtrl.ConversionTable).Methods("GET")
	Router.HandleFunc("/{category}", siteCtrl.Category).Methods("GET")
	Router.HandleFunc("/{category}/{calculation}", siteCtrl.Calculation).Methods("GET")

	Router.NotFoundHandler = http.HandlerFunc(siteCtrl.NotFound)

	// Static Files Handler in Dev mode
	boolPtr := flag.Bool("dev", false, "Use in development")
	flag.Parse()
	if *boolPtr {
		Router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	}
}
