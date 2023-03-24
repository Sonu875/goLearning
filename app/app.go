package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Sonu875/goLearning/domain"
	"github.com/Sonu875/goLearning/logger"
	"github.com/Sonu875/goLearning/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func santityCheck() {
	switch {
	case os.Getenv("APP_HOST") == "":
		logger.Fatal("Application host is not set as evn")
	case os.Getenv("APP_PORT") == "":
		logger.Fatal("Application port is not set as evn")
	case os.Getenv("DB_HOST") == "":
		logger.Fatal("Application host is not set as evn")
	case os.Getenv("DB_PORT") == "":
		logger.Fatal("Application host is not set as evn")
	case os.Getenv("DB_PASSWORD") == "":
		logger.Fatal("Application host is not set as evn")
	case os.Getenv("DB_NAME") == "":
		logger.Fatal("Application host is not set as evn")
	case os.Getenv("DB_USER") == "":
		logger.Fatal("Application host is not set as evn")
	}
}

func Start() {
	santityCheck()
	router := mux.NewRouter()
	dbClient := getDbClient()
	customerRepoDB := domain.NewCustomerRepoDb(dbClient)
	accountRepoDB := domain.NewAcountRepoDb(dbClient)
	ch := CustomerHandler{service.NewCustomerService(customerRepoDB)}
	ac := AccountHandler{service.NewAccountService(accountRepoDB)}
	router.HandleFunc("/api/customers", ch.getAllCustomer).Methods("Get")
	router.HandleFunc("/api/customer/{customer_id:[0-9]+}", ch.getCustomerByID).Methods("Get")
	router.HandleFunc("/api/customer/{customer_id:[0-9]+}/account", ac.NewAccount).Methods(http.MethodPost)

	router.HandleFunc("/api/time", currentTime)

	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router))
}

func getDbClient() *sqlx.DB {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	client, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Print(err.Error())
		panic(err)
	}
	return client
}
