package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/simpleInternetUser/TaskManager/tasks"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, tasks.GetAllTasks())
}

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8080", nil)
}
