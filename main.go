package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/simpleInternetUser/TaskManager/tasks"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, tasks.GetAllTasks())
}

func GetDate() (string, int) {
	y, m, d := time.Now().Date()
	str := fmt.Sprintf("%v %v %v", d, m, y)
	t := time.Now()
	id := t.Year() + t.Day() + int(t.Month()) + t.Hour() + t.Minute() + t.Second()
	return str, id
}

func AddNewTask(w http.ResponseWriter, r *http.Request) {

	newT := &tasks.Task{}
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/addtask.html")
		t.Execute(w, nil)
	} else {

		newT.Title = r.FormValue("title")
		newT.Description = r.FormValue("description")
		newT.Status = "new"
		newT.Date, newT.Id = GetDate()

		file, _ := os.OpenFile("tasks.json", os.O_RDWR, 0644)
		defer file.Close()
		var newUT tasks.AllTasks
		bfile, _ := ioutil.ReadAll(file)
		json.Unmarshal(bfile, &newUT.TasksA)
		newUT.TasksA = append(newUT.TasksA, newT)
		newData, _ := json.MarshalIndent(&newUT.TasksA, "", " ")
		ioutil.WriteFile("tasks.json", newData, 0666)

		http.Redirect(w, r, "/", 301)
	}
}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/addtask/", AddNewTask)
	http.ListenAndServe(":8080", nil)
}
