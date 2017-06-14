package main

import (
	"log"
	"net/http"

	"github.com/FriedPigeon/mathfever-go/handler"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", handler.Router))
}
