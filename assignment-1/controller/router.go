package controller

import (
	"log"
	"net/http"

	"GLIM_Hacktiv8/golang-intermediate/assignment-1/service"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

const baseURL = "0.0.0.0:8080"

func Routing() {

	r := mux.NewRouter()
	r.HandleFunc("/todos", service.Get).Methods("GET")
	r.HandleFunc("/todos", service.Create).Methods("POST")
	r.HandleFunc("/todos/{id}", service.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/todos/{id}", service.Put).Methods(http.MethodPut)
	r.HandleFunc("/todos/{id}", service.Delete).Methods(http.MethodDelete)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	// http.HandleFunc("/hello", PrintHello)

	log.Println("Listening in url " + baseURL)
	log.Fatal(http.ListenAndServe(baseURL, r))
}
