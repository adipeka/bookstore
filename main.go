package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/adipeka/bookstore/handler"
)

const (
	dbhost     = "localhost"
	dbport     = "5432"
	dbuser     = "postgres"
	dbpassword = "12345"
	dbname     = "postgres"
)

func main() {

	connectionstring := fmt.Sprintf("user=%s dbname=%s port=%s password=%s sslmode=disable", dbuser, dbname, dbport, dbpassword)
	db, err := sql.Open("postgres", connectionstring)

	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	hr := handler.New(db,router)

}
