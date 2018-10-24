package main

import (
	"log"
	"net/http"

	"github.com/ludw1gj/mathfever-go/app/router"
)

func main() {
	log.Println("mathfever.xyz listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router.Load()))
}
