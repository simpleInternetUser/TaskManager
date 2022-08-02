package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

		newUT := tasks.GetAllTasks()
		newUT.TasksA = append(newUT.TasksA, newT)
		newData, _ := json.MarshalIndent(&newUT.TasksA, "", " ")
		ioutil.WriteFile("tasks.json", newData, 0666)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func DelTask(w http.ResponseWriter, r *http.Request) {

	iddel, _ := strconv.Atoi(r.FormValue("id"))
	listT := tasks.GetAllTasks()

	i := 0
	for ; i < len(listT.TasksA); i++ {
		if listT.TasksA[i].Id == iddel {
			break
		}
	}

	listT.TasksA = append(listT.TasksA[:i], listT.TasksA[i+1:]...)
	newData, _ := json.MarshalIndent(&listT.TasksA, "", " ")
	ioutil.WriteFile("tasks.json", newData, 0666)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func EditTask(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/edittask.html")
	if err != nil {
		log.Fatal(err)
	}

	iddel, _ := strconv.Atoi(r.FormValue("id"))
	listT := tasks.GetAllTasks()

	i := 0
	for ; i < len(listT.TasksA); i++ {
		if listT.TasksA[i].Id == iddel {
			break
		}
	}
	tmpl.Execute(w, listT.TasksA[i])

	listT.TasksA[i].Title = r.FormValue("title")
	listT.TasksA[i].Description = r.FormValue("description")

	fmt.Println(listT.TasksA[i])
	//newData, _ := json.MarshalIndent(&listT.TasksA, "", " ")
	//ioutil.WriteFile("tasks.json", newData, 0666)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
