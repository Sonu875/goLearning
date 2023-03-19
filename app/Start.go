package app

import (
	"fmt"
	"log"
	"net/http"
)

func Start() {
	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
