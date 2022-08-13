package main

import (
	"log"
	"net/http"

	"github.com/simpleInternetUser/TaskManager/config"
	"github.com/simpleInternetUser/TaskManager/handlers"
)

func main() {

	conf, err := config.ReadConfig("config/config.json")
	if err != nil {
		log.Println(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Index)
	mux.HandleFunc("/addtask/", handlers.AddNewTask)
	mux.HandleFunc("/edittask/", handlers.EditTask)
	mux.HandleFunc("/deltask/", handlers.DelTask)

	log.Println("Starting a web server on localhost:8080")
	err = http.ListenAndServe(conf.ServerPort, mux)
	if err != nil {
		log.Println(err)
	}
}
