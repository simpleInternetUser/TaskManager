package handlers

import (
	"encoding/json"
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/simpleInternetUser/TaskManager/tasks"
)

func Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println(err)
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

func IndexJson(data tasks.AllTasks, r *http.Request) (int, error) {
	strId := r.FormValue("id")
	if strId == "" {
		return 0, errors.New("Empty string")
	}
	iddel, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Println(err)
	}
	i := 0
	for ; i < len(data.TasksA); i++ {
		if data.TasksA[i].Id == iddel {
			break
		}
	}
	return i, nil
}

func EditTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else {
		tmpl, err := template.ParseFiles("templates/edittask.html")
		if err != nil {
			log.Println(err)
		}

		listT := tasks.GetAllTasks()
		i, err := IndexJson(listT, r)
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}
		tmpl.ExecuteTemplate(w, "edit", listT.TasksA[i])
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func SaveTask(w http.ResponseWriter, r *http.Request) {
	newT := tasks.Task{}
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/edittask.html")
		t.Execute(w, nil)
	} else {

		PageData(&newT, r)

		newUT := tasks.GetAllTasks()
		newUT.TasksA = append(newUT.TasksA, newT)
		newData, err := json.MarshalIndent(&newUT.TasksA, "", " ")
		if err != nil {
			log.Println(err)
		}
		ioutil.WriteFile("tasks.json", newData, 0666)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}
func DelTask(w http.ResponseWriter, r *http.Request) {

	listT := tasks.GetAllTasks()
	i, err := IndexJson(listT, r)
	if err != nil {
		log.Println(err)
	}

	listT.TasksA = append(listT.TasksA[:i], listT.TasksA[i+1:]...)
	newData, err := json.MarshalIndent(&listT.TasksA, "", " ")
	if err != nil {
		log.Println(err)
	}
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
	return time.Now(), int(t.Unix())
}
