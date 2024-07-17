package main

import (
	"ATM-service/internal/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	go handlers.StartLogger()

	router := mux.NewRouter()
	router.HandleFunc("/accounts", handlers.CreateAccount).Methods("POST")
	router.HandleFunc("/accounts/{id}/deposit", handlers.Deposit).Methods("POST")
	router.HandleFunc("/accounts/{id}/withdraw", handlers.Withdraw).Methods("POST")
	router.HandleFunc("/accounts/{id}/balance", handlers.GetBalance).Methods("GET")
	log.Printf("starting server...\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
