package main

import (
	"GLIM_Hacktiv8/golang-intermediate/assignment-1/controller"
	_ "GLIM_Hacktiv8/golang-intermediate/assignment-1/docs"
	"GLIM_Hacktiv8/golang-intermediate/assignment-1/service"
)

/*
type Todo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
*/

// var Todos []*Todo

// const baseURL = "0.0.0.0:8080"

// @title Todo Application
// @version 1.0
// @description This is a todo list test management application
// @contact.name Ariq Naufal
// @contact.email ariq.naufal29@gmail.com
// @host localhost:8080
// @BasePath /
func main() {

	service.Connecting()
	controller.Routing()
	/*
		//serve http server
		 log.Println("Listening to url: " + baseURL)

		//log fatal akan stop program ketika error
		 log.Fatal(http.ListenAndServe(baseURL, r))
		// log.Fatal(http.ListenAndServe(baseURL, nil))
	*/
}
