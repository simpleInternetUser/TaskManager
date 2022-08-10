package handlers

import (
	"encoding/json"
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

func AddNewTask(w http.ResponseWriter, r *http.Request) {

	newT := tasks.Task{}
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/addtask.html")
		t.Execute(w, nil)
	} else {

		PageData(&newT, r)

		newUT := tasks.GetAllTasks()
		newUT.TasksA = append(newUT.TasksA, newT)
		newData, _ := json.MarshalIndent(&newUT.TasksA, "", " ")
		ioutil.WriteFile("tasks.json", newData, 0666)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func IndexJson(data tasks.AllTasks, r *http.Request) int {
	iddel, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Fatal(err)
	}
	i := 0
	for ; i < len(data.TasksA); i++ {
		if data.TasksA[i].Id == iddel {
			break
		}
	}
	return i
}

func EditTask(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/edittask.html")
	if err != nil {
		log.Fatal(err)
	}

	listT := tasks.GetAllTasks()
	i := IndexJson(listT, r)
	tmpl.ExecuteTemplate(w, "edit", listT.TasksA[i])
	//PageData(listT.TasksA[i], r)

	//http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
func SaveTask(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	http.Redirect(w, r, "/"+title, http.StatusFound)
}
func DelTask(w http.ResponseWriter, r *http.Request) {

	listT := tasks.GetAllTasks()
	i := IndexJson(listT, r)

	listT.TasksA = append(listT.TasksA[:i], listT.TasksA[i+1:]...)
	newData, _ := json.MarshalIndent(&listT.TasksA, "", " ")
	ioutil.WriteFile("tasks.json", newData, 0666)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func PageData(t *tasks.Task, r *http.Request) {
	t.Title = r.FormValue("title")
	t.Description = r.FormValue("description")
	if t.Id == 0 {
		t.Status = "new"
		t.Date, t.Id = DateNowAndId()
	} else {
		t.Status = r.FormValue("status")
	}
}

func DateNowAndId() (time.Time, int) {
	t := time.Now()
	id := t.Unix()
	return time.Now(), int(id)
}
