package main

import (
	"hendlers/hendlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", hendlers.Index)
	http.HandleFunc("/addtask/", hendlers.AddNewTask)
	http.HandleFunc("/edittask/", hendlers.EditTask)
	http.HandleFunc("/savetask/", hendlers.SaveTask)
	http.HandleFunc("/deltask/", hendlers.DelTask)
	http.ListenAndServe(":8080", nil)
}
