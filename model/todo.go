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

// GetTodo to get a todo
func GetTodo(id int) Todo {
	var todo Todo

	for i := 0; i < len(todos); i++ {
		if todos[i].ID == id {
			return todos[i]
		}
	}

	return todo
}

// UpdateTodo Update todo
func UpdateTodo(id int, todo Todo) Todo {
	for i := 0; i < len(todos); i++ {
		if todos[i].ID == id {
			todos[i].Title = todo.Title
			todos[i].Status = todo.Status
		}
	}

	return todo
}

// DeleteTodo Delete a todo from list
func DeleteTodo(id int) Todo {
	var todo Todo

	for i := 0; i < len(todos); i++ {
		if todos[i].ID == id {
			todo = todos[i]
			todos = append(todos[:i], todos[i+1:]...)
		}
	}

	return todo
}

// To generate next index value
func getNextIndex() int {
	index++
	return index
}
