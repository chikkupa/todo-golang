package controller

import (
	"encoding/json"
	"net/http"

	"todo-application/model"
)

// CreateTodo Controller for create todo
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo = model.CreateTodo(todo)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
