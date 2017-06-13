package router

import (
	"flag"
	"net/http"

	"github.com/FriedPigeon/mathfever-go/handlers"
	"github.com/gorilla/mux"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter().StrictSlash(true)

	// API Router
	apiRouter := Router.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/", handlers.GetCategories).Methods("GET")
	apiRouter.HandleFunc("/{category}", handlers.GetCategory).Methods("GET")
	apiRouter.HandleFunc("/{category}/{calculation}", handlers.GetCalculation).Methods("GET")
	apiRouter.HandleFunc("/{category}/{calculation}", handlers.DoCalculation).Methods("POST")

	apiRouter.NotFoundHandler = http.HandlerFunc(handlers.NotFoundAPI)

	// Site Router
	Router.HandleFunc("/", handlers.Home).Methods("GET")
	Router.HandleFunc("/about", handlers.About).Methods("GET")
	Router.HandleFunc("/help", handlers.Help).Methods("GET")
	Router.HandleFunc("/privacy", handlers.Privacy).Methods("GET")
	Router.HandleFunc("/terms", handlers.Terms).Methods("GET")
	Router.HandleFunc("/message-board", handlers.MessageBoard).Methods("GET")
	Router.HandleFunc("/networking/conversion-table", handlers.ConversionTable).Methods("GET")
	Router.HandleFunc("/{category}", handlers.Category).Methods("GET")
	Router.HandleFunc("/{category}/{calculation}", handlers.Calculation).Methods("GET")

	Router.NotFoundHandler = http.HandlerFunc(handlers.NotFound)

	// Static Files Handler in Dev mode
	boolPtr := flag.Bool("dev", false, "Use in development")
	flag.Parse()
	if *boolPtr {
		Router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	}
}
