package model

// Todo Structure to save todo
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

var todos = make([]Todo, 0)
var index = 0

// CreateTodo Create a new todo
func CreateTodo(todo Todo) Todo {
	todo.ID = getNextIndex()
	todos = append(todos, todo)
	return todo
}

// GetTodos Get list of todos
func GetTodos() []Todo {
	return todos
}

// To generate next index value
func getNextIndex() int {
	index++
	return index
}
