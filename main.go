package main

import (
	"net/http"

	"github.com/simpleInternetUser/TaskManager/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/addtask/", handlers.AddNewTask)
	http.HandleFunc("/edittask/", handlers.EditTask)
	http.HandleFunc("/savetask/", handlers.SaveTask)
	http.HandleFunc("/deltask/", handlers.DelTask)
	http.ListenAndServe(":8080", nil)
}
