package main

import (
	"database/sql"
	"github.com/pkg/errors"
)

type Customer struct {
	CustomerId string
	Name       string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "liudehan:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		panic(err)
	}
}

func QueryCustomerById(id string) (Customer, error) {
	var customer Customer
	row := Db.QueryRow("select id ,name from customer where id = ?", id)
	err := row.Scan(&customer.CustomerId, &customer.Name)
	if err != nil {
		return customer, errors.Wrap(err, "dao#QueryCustomerById err")
	}
	return customer, nil
}
