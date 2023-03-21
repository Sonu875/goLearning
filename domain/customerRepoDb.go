package domain

import (
	"database/sql"
	"fmt"
	"log"

	appError "github.com/Sonu875/goLearning/Errors"
	"github.com/Sonu875/goLearning/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CustomerRepoDb struct {
	client *sqlx.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "ssonu"
	password = "ssonu"
	dbname   = "banking"
)

func (d CustomerRepoDb) FindAll(status string) ([]Customer, *appError.AppError) {
	var findAllCustomer string
	var customers = make([]Customer, 0)
	if status == "" {
		findAllCustomer = "select * from customers "
	} else if status == "inactive" {
		findAllCustomer = "select * from customers where status=false"
	} else {
		findAllCustomer = "select * from customers where status=true"
	}
	err := d.client.Select(&customers, findAllCustomer)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, appError.NewNotFoundError("No Customers found")

		} else {
			logger.Error("Error while querying cutomers" + err.Error())
			return nil, appError.NewInternalServerError("Database related issue")
		}

	}
	return customers, nil
}

func (d CustomerRepoDb) GetCustomerByID(id string) (*Customer, *appError.AppError) {

	findCustomerByID := "select * from customers where customer_id=$1"
	var c Customer
	err := d.client.Get(&c, findCustomerByID, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, appError.NewNotFoundError("No Customers found")
		} else {
			logger.Error("Error while looping cutomers" + err.Error())
			return nil, appError.NewInternalServerError("Database related issue")
		}

	}
	return &c, nil
}

func NewCustomerRepoDb() CustomerRepoDb {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	client, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Print(err.Error())
		panic(err)
	}
	return CustomerRepoDb{client}
}
