package app

import (
	"fmt"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world 123")
}
