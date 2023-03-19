package app

import (
	"log"
	"net/http"
)

func Start() {

	http.HandleFunc("/time", greet)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
