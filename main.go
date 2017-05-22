package main

import (
	"log"
	"net/http"

	"github.com/spottywolf/mathfever/router"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", router.Router))
}
