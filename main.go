package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/simpleInternetUser/TaskManager/handlers"
)

func main() {
	r := mux.NewRouter()
	//mux := http.NewServeMux()
	r.HandleFunc("/", handlers.Index)
	r.HandleFunc("/addtask/", handlers.AddNewTask)
	r.HandleFunc("/edittask/", handlers.EditTask)
	r.HandleFunc("/deltask/", handlers.DelTask)

	log.Println("Starting a web server on localhost:8080")
	err := http.ListenAndServe(handlers.CONF.ServerPort, r)
	if err != nil {
		log.Println(err)
	}
}
