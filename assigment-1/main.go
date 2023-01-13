package main

import (
	"log"
	"net/http"
	//"github.com/gorilla/mux"
)

type Todo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var Todos []*Todo

const baseURL = "0.0.0.0:8080"

func main() {

	/*
		r := mux.NewRouter()
		r.HandleFunc("/todos", Create).Methods(http.MethodPost)
		r.HandleFunc("todos", Get).Methods(http.MethodGet)
	*/

	http.HandleFunc("/hello", PrintHello)

	//serve http server
	log.Println("Listening to url: " + baseURL)
	log.Fatal(http.ListenAndServe(baseURL, nil))
	// log.Fatal(http.ListenAndServe(baseURL, r))

}

func PrintHello(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("hello"))
	default:
		w.Write([]byte("hai"))
	}
}

/*
func Create(w http.ResponseWriter, r *http.Request) {
	var t Todo
	decoder := json.NewDecoder(r.Body)
	_ = decoder.Decode(&t)
	Todos = append(Todos, &t)
	w.Write([]byte("Success add todo " + t.Name))
}

func Get(w http.ResponseWriter, r *http.Request) {
	todosRes, _ := json.Marshal(Todos)
	w.Header().Set("Content-Type", "application/json")
	w.Write(todosRes)
}

*/

//log fatal akan stop program ketika error
