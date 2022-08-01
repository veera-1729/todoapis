package main

import (
	"net/http"
	"quickstart/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	router.HandleFunc("/Addtask", controllers.AddTask).Methods("POST")
	http.ListenAndServe(":12345", router)

}
