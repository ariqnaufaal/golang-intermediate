package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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

	w.Header().Set("Content-Type", "application/json")
	var todos []Todo
	DB.Find(&todos)
	json.NewEncoder(w).Encode(todos)

	/*
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
	*/
}

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

	w.Header().Set("Content-Type", "application/json")
	var todo Todo

	json.NewDecoder(r.Body).Decode(&todo)
	fmt.Println(r.Body)
	DB.Create(&todo)
	json.NewEncoder(w).Encode(todo)
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

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var todo Todo
	DB.First(&todo, params["id"])
	json.NewEncoder(w).Encode(todo)

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
	/*
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
	*/
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var todo Todo
	DB.First(&todo, params["id"])
	json.NewDecoder(r.Body).Decode(&todo)
	DB.Save(&todo)
	json.NewEncoder(w).Encode(todo)
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
	/*
		id := mux.Vars(r)["id"]
		for i := 0; i < len(Todos); i++ {
			if Todos[i].ID == id {
				Todos = append(Todos[:i], Todos[i+1:]...)
				w.Write([]byte("Success delete todo " + id))
				return
			}
		}
	*/
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var todo Todo
	DB.Delete(&todo, params["id"])
	json.NewEncoder(w).Encode(todo)
}

//*/
