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

	"github.com/simpleInternetUser/TaskManager/config"
	"github.com/simpleInternetUser/TaskManager/tasks"
)

var CONF, _ = config.ReadConfig("config/config.json")

func Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(CONF.PathTempl + "index.html")
	printErr(err)
	fd := tasks.AllTasks{}
	dat, err := fd.List()
	printErr(err)
	tmpl.Execute(w, dat)
}

func AddNewTask(w http.ResponseWriter, r *http.Request) {

	newT := tasks.Task{}
	if r.Method == "GET" {
		t, _ := template.ParseFiles(CONF.PathTempl + "addtask.html")
		t.Execute(w, nil)
	} else {

		PageData(&newT, r)
		fd := tasks.AllTasks{}
		newUT, err := fd.List()
		if err != nil {
			log.Println(err)
		}
		newUT.TasksA = append(newUT.TasksA, newT)
		newData, _ := json.MarshalIndent(&newUT.TasksA, "", " ")
		ioutil.WriteFile(CONF.PathTasks, newData, 0666)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func IndexJson(data tasks.AllTasks, r *http.Request) (int, error) {

	strId := r.FormValue("id")
	if strId == "" {
		return 0, errors.New("bject id not passed, empty string")
	}

	iddel, err := strconv.Atoi(r.FormValue("id"))
	printErr(err)

	i := 0
	for ; i < len(data.TasksA); i++ {
		if data.TasksA[i].Id == iddel {
			break
		}
	}
	return i, nil
}

func EditTask(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(CONF.PathTempl + "edittask.html")
	printErr(err)

	fd := tasks.AllTasks{}
	listT, err := fd.List()
	printErr(err)

	i, err := IndexJson(listT, r)
	printErr(err)

	if r.Method == "POST" {
		PageData(&listT.TasksA[i], r)
		newData, err := json.MarshalIndent(&listT.TasksA, "", " ")
		printErr(err)
		ioutil.WriteFile(CONF.PathTasks, newData, 0666)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else {
		tmpl.ExecuteTemplate(w, "edit", listT.TasksA[i])
	}
}

func DelTask(w http.ResponseWriter, r *http.Request) {

	fd := tasks.AllTasks{}
	listT, err := fd.List()
	printErr(err)
	i, err := IndexJson(listT, r)
	printErr(err)

	listT.TasksA = append(listT.TasksA[:i], listT.TasksA[i+1:]...)
	newData, err := json.MarshalIndent(&listT.TasksA, "", " ")
	printErr(err)
	ioutil.WriteFile(CONF.PathTasks, newData, 0666)

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

func printErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func getId(r *http.Request) (int, error) {
	strId := r.FormValue("id")
	if strId == "" {
		return 0, errors.New("bject id not passed, empty string")
	}

	iddel, err := strconv.Atoi(r.FormValue("id"))
	return iddel, err
}
