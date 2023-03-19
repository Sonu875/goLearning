package app

import (
	"log"
	"net/http"

	"github.com/Sonu875/goLearning/domain"
	"github.com/Sonu875/goLearning/service"
	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepoDb())}
	router.HandleFunc("/api/customers", ch.getAllCustomer)
	router.HandleFunc("/api/time", currentTime)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
