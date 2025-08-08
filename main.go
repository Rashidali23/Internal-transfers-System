package main

import (
	"log"
	"net/http"

	"Internal-transfers-System/db"
	"Internal-transfers-System/handler"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("main started")
	db.InitDB()
	r := mux.NewRouter()

	r.HandleFunc("/accounts", handler.CreateAccount).Methods("POST")
	r.HandleFunc("/accounts/{id}", handler.GetAccount).Methods("GET")
	r.HandleFunc("/transactions", handler.CreateTransaction).Methods("POST")

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", r)
}
