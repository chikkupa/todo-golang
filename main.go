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
	fmt.Println("Server listening @ 8080")
	log.Fatal(http.ListenAndServe(":8000", router))
}
