package router

import (
	"flag"
	"net/http"

	"github.com/gorilla/mux"
)

var Router *mux.Router

func init() {
	Router = mux.NewRouter()

	// Site Handler
	Router.HandleFunc("/", homeHandler).Name("home")
	Router.HandleFunc("/about", aboutHandler).Name("about")
	Router.HandleFunc("/help", helpHandler).Name("help")
	Router.HandleFunc("/privacy", privacyHandler).Name("privacy")
	Router.HandleFunc("/terms", termsHandler).Name("terms")
	Router.HandleFunc("/message-board", messageBoardHandler).Name("message-board")
	Router.HandleFunc("/networking/conversion-table", conversionTableHandler)
	Router.HandleFunc("/{category}", categoriesHandler)
	Router.HandleFunc("/{category}/{calculation}", calculationsHandler)

	// API Handler
	Router.HandleFunc("/api/{category}/{calculation}", calculationsAPIHandler)

	// Not Found Handler
	Router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	// Static Files Handler in Dev mode
	boolPtr := flag.Bool("dev", false, "Use in development")
	flag.Parse()
	if *boolPtr {
		Router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	}
}
