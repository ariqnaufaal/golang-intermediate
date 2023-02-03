package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func Routing() {

	r := mux.NewRouter()
	r.HandleFunc("/todos", Get).Methods("GET")
	r.HandleFunc("/todos", Create).Methods("POST")
	r.HandleFunc("/todos/{id}", GetByID).Methods(http.MethodGet)
	r.HandleFunc("/todos/{id}", Put).Methods(http.MethodPut)
	r.HandleFunc("/todos/{id}", Delete).Methods(http.MethodDelete)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	// http.HandleFunc("/hello", PrintHello)

	log.Println("Listening in url " + baseURL)
	log.Fatal(http.ListenAndServe(baseURL, r))
}
