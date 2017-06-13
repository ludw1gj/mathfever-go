package router

import (
	"flag"
	"net/http"

	"github.com/FriedPigeon/mathfever-go/handler"
	"github.com/gorilla/mux"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter().StrictSlash(true)

	// API Router
	apiRouter := Router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/", handler.GetCategories).Methods("GET")
	apiRouter.HandleFunc("/{category}", handler.GetCategory).Methods("GET")
	apiRouter.HandleFunc("/{category}/{calculation}", handler.GetCalculation).Methods("GET")
	apiRouter.HandleFunc("/{category}/{calculation}", handler.DoCalculation).Methods("POST")

	apiRouter.NotFoundHandler = http.HandlerFunc(handler.NotFoundAPI)

	// Site Router
	Router.HandleFunc("/", handler.Home).Methods("GET")
	Router.HandleFunc("/about", handler.About).Methods("GET")
	Router.HandleFunc("/help", handler.Help).Methods("GET")
	Router.HandleFunc("/privacy", handler.Privacy).Methods("GET")
	Router.HandleFunc("/terms", handler.Terms).Methods("GET")
	Router.HandleFunc("/message-board", handler.MessageBoard).Methods("GET")
	Router.HandleFunc("/networking/conversion-table", handler.ConversionTable).Methods("GET")
	Router.HandleFunc("/{category}", handler.CategoryPage).Methods("GET")
	Router.HandleFunc("/{category}/{calculation}", handler.CalculationPage).Methods("GET")

	Router.NotFoundHandler = http.HandlerFunc(handler.NotFound)

	// Static Files Handler in Dev mode
	boolPtr := flag.Bool("dev", false, "Use in development")
	flag.Parse()
	if *boolPtr {
		Router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	}
}
