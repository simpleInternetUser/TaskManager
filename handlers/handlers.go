package handlers

import (
	"errors"
	"html/template"
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

		pageData(&newT, r)
		listT, err := tasks.AllTasks{}.List()
		printErr(err)
		listT.TasksA = append(listT.TasksA, newT)
		err = listT.Update()
		printErr(err)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func EditTask(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(CONF.PathTempl + "edittask.html")
	printErr(err)

	listT, err := tasks.AllTasks{}.List()
	printErr(err)

	id, err := getId(r)
	printErr(err)
	edt, err := listT.Get(id)
	printErr(err)

	if r.Method == "POST" {
		pageData(edt, r)
		err = listT.Update()
		printErr(err)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else {
		tmpl.ExecuteTemplate(w, "edit", edt)
	}
}

func DelTask(w http.ResponseWriter, r *http.Request) {

	listT, err := tasks.AllTasks{}.List()
	printErr(err)
	id, err := getId(r)
	printErr(err)
	listT, err = listT.Delete(id)
	printErr(err)
	err = listT.Update()
	printErr(err)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func pageData(t *tasks.Task, r *http.Request) {
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
	return t, int(t.Unix())
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
