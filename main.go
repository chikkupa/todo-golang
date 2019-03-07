package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-application/controller"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/todo", controller.CreateTodo).Methods("POST")
	router.HandleFunc("/todo", controller.GetTodos).Methods("GET")
	router.HandleFunc("/todo/{id}", controller.GetTodo).Methods("GET")
	router.HandleFunc("/todo/{id}", controller.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todo/{id}", controller.DeleteTodo).Methods("DELETE")

	fmt.Println("Server listening @ 8080")
	log.Fatal(http.ListenAndServe(":8000", router))
}
