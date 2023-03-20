package domain

import (
	"database/sql"
	"fmt"
	"log"

	appError "github.com/Sonu875/goLearning/Errors"
	_ "github.com/lib/pq"
)

type CustomerRepoDb struct {
	client *sql.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "ssonu"
	password = "ssonu"
	dbname   = "banking"
)

func (d CustomerRepoDb) FindAll() ([]Customer, *appError.AppError) {

	findAllCustomer := "select * from customers "
	rows, err := d.client.Query(findAllCustomer)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, appError.NewNotFoundError("No Customers found")

		} else {
			log.Println("Error while querying cutomers" + err.Error())
			return nil, appError.NewInternalServerError("Database related issue")
		}

	}
	var customers = make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.DateOfBirth, &c.Zipcode, &c.Status)
		if err != nil {
			log.Println("Error while looping cutomers" + err.Error())
			return nil, appError.NewInternalServerError("Something went wrong")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepoDb) GetCustomerByID(id string) (*Customer, *appError.AppError) {

	findCustomerByID := "select * from customers where customer_id=$1"
	var c Customer
	rows := d.client.QueryRow(findCustomerByID, id)
	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.DateOfBirth, &c.Zipcode, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, appError.NewNotFoundError("No Customers found")
		} else {
			log.Println("Error while looping cutomers" + err.Error())
			return nil, appError.NewInternalServerError("Database related issue")
		}

	}
	return &c, nil
}

func NewCustomerRepoDb() CustomerRepoDb {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	client, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Print(err.Error())
		panic(err)
	}
	return CustomerRepoDb{client}
}
