package main

import (
	"log"
	"net/http"
)

const baseURL = "0.0.0.0:8080"

func main() {

	http.HandleFunc("/hello", PrintHello)

	//serve http server
	log.Println("Listening to url: " + baseURL)
	_ = http.ListenAndServe(baseURL, nil)

}

func PrintHello(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello World"))
	default:
		w.Write([]byte("hai"))
	}
}

//log fatal akan stop program ketika error
