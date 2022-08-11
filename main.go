package main

import (
	"log"
	"net/http"

	"github.com/simpleInternetUser/TaskManager/handlers"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/addtask/", handlers.AddNewTask)
	mux.HandleFunc("/edittask/", handlers.EditTask)
	mux.HandleFunc("/deltask/", handlers.DelTask)

	log.Println("Starting a web server on localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println(err)
	}
}
