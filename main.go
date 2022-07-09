package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"bookstore/handler"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	dbhost     = "localhost"
	dbport     = "5432"
	dbuser     = "postgres"
	dbpassword = "12345"
	dbname     = "postgres"
)

func main() {

	fmt.Println("Starting the server")

	connectionstring := fmt.Sprintf("user=%s dbname=%s port=%s password=%s sslmode=disable", dbuser, dbname, dbport, dbpassword)
	db, err := sql.Open("postgres", connectionstring)

	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	hr := handler.New(db, router)

	router.HandleFunc("/book", hr.GetBook).Methods("GET")

	http.ListenAndServe(":8001", router)

}
