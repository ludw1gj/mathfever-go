package handler

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter().StrictSlash(true)

	// API Router
	apiRouter := Router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/", getCategories).Methods("GET")
	apiRouter.HandleFunc("/{category}", getCategory).Methods("GET")
	apiRouter.HandleFunc("/{category}/{calculation}", getCalculation).Methods("GET")
	apiRouter.HandleFunc("/{category}/{calculation}", doCalculation).Methods("POST")
	apiRouter.NotFoundHandler = http.HandlerFunc(notFoundAPI)

	// Site Router
	Router.HandleFunc("/", home).Methods("GET").Name("Home")
	Router.HandleFunc("/about", about).Methods("GET").Name("About")
	Router.HandleFunc("/help", help).Methods("GET").Name("Help")
	Router.HandleFunc("/privacy", privacy).Methods("GET").Name("Privacy")
	Router.HandleFunc("/terms", terms).Methods("GET").Name("Terms")
	Router.HandleFunc("/message-board", messageBoard).Methods("GET").Name("Message Board")
	Router.HandleFunc("/networking/conversion-table", conversionTable).Methods("GET")
	Router.HandleFunc("/{category}", categoryPage).Methods("GET").Name("Category")
	Router.HandleFunc("/{category}/{calculation}", calculationPage).Methods("GET").Name("Calculation")
	Router.NotFoundHandler = http.HandlerFunc(notFound)

	// Static Files Handler in Dev mode
	boolPtr := flag.Bool("dev", false, "Use in development")
	flag.Parse()
	if *boolPtr {
		fs := http.FileServer(http.Dir("./assets"))
		Router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))
	}
}
