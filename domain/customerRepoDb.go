package domain

import (
	"database/sql"
	"fmt"
	"log"

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

func (d CustomerRepoDb) FindAll() ([]Customer, error) {

	defer d.client.Close()

	findAllCustomer := "select * from customers "
	rows, err := d.client.Query(findAllCustomer)
	if err != nil {
		log.Println("Error while querying cutomers" + err.Error())
		return nil, err
	}
	var customers = make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.DateOfBirth, &c.Zipcode, &c.Status)
		if err != nil {
			log.Println("Error while looping cutomers" + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
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
