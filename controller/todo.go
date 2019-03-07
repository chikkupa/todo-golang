package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"todo-application/model"

	"github.com/gorilla/mux"
)

// CreateTodo Controller for create todo
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo = model.CreateTodo(todo)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// GetTodos To get the list of todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := model.GetTodos()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// GetTodo To get todo details
func GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	todo := model.GetTodo(id)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// UpdateTodo Update a todo
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var todo model.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo = model.UpdateTodo(id, todo)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

// DeleteTodo Delete a todo
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	todo := model.DeleteTodo(id)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
