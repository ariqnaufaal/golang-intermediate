package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "GLIM_Hacktiv8/golang-intermediate/assigment-1/docs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Todo struct {
	ID   string `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Todo
}

var Todos []*Todo

const baseURL = "0.0.0.0:8080"

func Connect() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "Windows_mySQL1"
	dbName := "todos_db"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// @title Todo Application
// @version 1.0
// @description This is a todo list test management application
// @contact.name Ariq Naufal
// @contact.email ariq.naufal29@gmail.com
// @host localhost:8080
// @BasePath /
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todos", Create).Methods("POST")
	r.HandleFunc("/todos", Get).Methods("GET")
	r.HandleFunc("/todos/{id}", GetByID).Methods(http.MethodGet)
	r.HandleFunc("/todos/{id}", Put).Methods(http.MethodPut)
	r.HandleFunc("/todos/{id}", Delete).Methods(http.MethodDelete)

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	// http.HandleFunc("/hello", PrintHello)

	//serve http server
	log.Println("Listening to url: " + baseURL)

	//log fatal akan stop program ketika error
	log.Fatal(http.ListenAndServe(baseURL, r))
	// log.Fatal(http.ListenAndServe(baseURL, nil))
}

/*
func PrintHello(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("hello"))
	default:
		w.Write([]byte("hai"))
	}
}
*/

///*
// Create is a handler for create todos API
// @Summary Create new todos
// @Description Create new todos with parameter id and name
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /todos [post]
func Create(w http.ResponseWriter, r *http.Request) {
	/*
		var t Todo
		decoder := json.NewDecoder(r.Body)
		_ = decoder.Decode(&t)
		Todos = append(Todos, &t)
		w.Write([]byte("Success add todo " + t.Name))
	*/

	db := Connect()
	if r.Method == "POST" {
		name := r.FormValue("name")
		insForm, err := db.Prepare("INSERT INTO todos_table(name) VALUES(?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name)
		log.Println("INSERT: Name: " + name)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)

	// var response Response

	// db := Connect()
	// defer db.Close()

	// err := r.ParseMultipartForm(4096)
	// if err != nil {
	// 	panic(err)
	// }
	// name := r.FormValue("name")

	// _, err = db.Exec("INSERT INTO todos_table(name) VALUES(?, ?)", name)

	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }
	// response.Status = 200
	// response.Message = "Insert data successfully"
	// fmt.Print("Insert data to database")

	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// json.NewEncoder(w).Encode(response)
}

// Get is a handler for get todos API
// @Summary Get new todos
// @Description get all todos list
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /todos [get]
func Get(w http.ResponseWriter, r *http.Request) {

	/*
		todosRes, _ := json.Marshal(Todos)
		w.Header().Set("Content-Type", "application/json")
		w.Write(todosRes)
	*/

	var todos Todo
	var response Response
	var arrTodo []Todo

	db := Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, name FROM todos_table")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		err = rows.Scan(&todos.ID, &todos.Name)
		if err != nil {
			log.Fatal(err.Error())
		} else {
			arrTodo = append(arrTodo, todos)
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrTodo

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

// GetByID is a handler for get todos API
// @Summary GetByID new todos
// @Description get todos by ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /todos/{id} [get]
func GetByID(w http.ResponseWriter, r *http.Request) {

	/*
		id := mux.Vars(r)["id"]
		for i := 0; i < len(Todos); i++ {
			if Todos[i].ID == id {
				todosRes, _ := json.Marshal(Todos[i])
				w.Header().Set("Content-Type", "application/json")
				w.Write(todosRes)
			}
		}
	*/

}

// Put is a handler for create todos API
// @Summary Update new todos
// @Description Update todos by ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /todos/{id} [put]
func Put(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for i := 0; i < len(Todos); i++ {
		if Todos[i].ID == id {
			var t Todo
			decoder := json.NewDecoder(r.Body)
			_ = decoder.Decode(&t)
			Todos[i] = &t
			w.Write([]byte("Success update todo " + t.ID))
			return
		}
	}
}

// Delete is a handler for create todos API
// @Summary Delete new todos
// @Description Delete string by ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {array} string
// @Router /todos/{id} [delete]
func Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for i := 0; i < len(Todos); i++ {
		if Todos[i].ID == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			w.Write([]byte("Success delete todo " + id))
			return
		}
	}
}

//*/
