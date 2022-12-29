package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"goapp/database"
	"goapp/entity"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var Db *sql.DB
var err error

// Create
func createTodo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createTodo")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Access-Control-Allow-Origin")
	if r.Method == "OPTIONS" {
		return
	}
	var t entity.Todo
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		fmt.Printf("Error while decoding request: %v", err)
	}
	// t.Id = strconv.Itoa(len(todos) + 1) // メモリで実施する場合
	// todos = append(todos, todo) // メモリで実施する場合
	cmd := `INSERT INTO todo (Todo) VALUES (?)`
	_, err = Db.Exec(cmd, t.Todo)
	if err != nil {
		fmt.Printf("Error while Inserting into DB: %v", err)
	}
	todoList := fetchTodoList()
	json.NewEncoder(w).Encode(todoList)
}

func fetchTodoList() []entity.Todo {
	cmd := "select * from todo"
	rows, err := Db.Query(cmd)
	if err != nil {
		fmt.Printf("Error while fetching data: %v", err)
	}
	var todoList []entity.Todo
	for rows.Next() {
		var t entity.Todo
		err := rows.Scan(&t.Id, &t.Todo)
		if err != nil {
			fmt.Printf("Error while scanning data: %v", err)
		}
		todoList = append(todoList, t)
	}
	rows.Close()
	return todoList
}

// Read
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	todoList := fetchTodoList()
	json.NewEncoder(w).Encode(todoList)
}

// Update
func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Access-Control-Allow-Origin")
	if r.Method == "OPTIONS" {
		return
	}
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var t entity.Todo
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		fmt.Printf("Error while decoding data: %v", err)
	}
	// todos[id].Todo = todo.Todo メモリで実施する場合
	v := t.Todo
	cmd := `UPDATE todo SET Todo=? WHERE Id=?`
	_, err = Db.Exec(cmd, v, id)
	if err != nil {
		fmt.Printf("Error while updating data: %v", err)
	}
}

// Delete
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	if r.Method == "OPTIONS" {
		return
	}
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Printf("Error while converting data: %v", err)
	}
	cmd := `DELETE FROM todo WHERE Id=?`
	Db.Exec(cmd, id)
	// for index, item := range todos { // メモリで実行する場合
	// 	if item.Id == params["id"] {
	// 		todos = append(todos[:index], todos[index+1:]...)
	// 		break
	// 	}
	// }
	todoList := fetchTodoList()
	json.NewEncoder(w).Encode(todoList)
}

func main() {
	Db = database.Connect()

	r := mux.NewRouter()
	r.HandleFunc("/todos", getTodos).Methods("GET")
	r.HandleFunc("/todos", createTodo).Methods("POST", "OPTIONS")
	r.HandleFunc("/todos/update/{id}", updateTodo).Methods("PUT", "OPTIONS")
	r.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE", "OPTIONS")

	fmt.Println("Start listening on port 8000")
	err := http.ListenAndServe(":8000", r)
	if err != nil {
		fmt.Printf("Error while starting server: %v", err)
	}
}
