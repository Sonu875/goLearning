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
	router.HandleFunc("/api/customers", ch.getAllCustomer).Methods("Get")
	router.HandleFunc("/api/customer/{customer_id:[0-9]+}", ch.getCustomerByID).Methods("Get")

	router.HandleFunc("/api/time", currentTime)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
