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

var conf, _ = config.ReadConfig("config/config.json")

func Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(conf.PathTempl + "index.html")
	if err != nil {
		log.Println(err)
	}

	var t tasks.Tasks
	tmp, err := t.List()
	if err != nil {
		log.Println(err)
	}

	tmpl.Execute(w, tmp)
}

func AddNewTask(w http.ResponseWriter, r *http.Request) {

	newT := tasks.Task{}
	if r.Method == "GET" {
		t, _ := template.ParseFiles(conf.PathTempl + "addtask.html")
		t.Execute(w, nil)
	} else {

		PageData(&newT, r)
		var t tasks.Task
		newUT, err := t.List()
		if err != nil {
			log.Println(err)
		}
		newUT = append(newUT, newT)
		newData, err := json.MarshalIndent(&newUT, "", " ")
		if err != nil {
			log.Println(err)
		}
		ioutil.WriteFile(conf.PathTasks, newData, 0666)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

func IndexJson(data []tasks.Task, r *http.Request) (int, error) {

	strId := r.FormValue("id")
	if strId == "" {
		return 0, errors.New("object id not passed, empty string")
	}

	iddel, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Println(err)
	}

	i := 0
	for ; i < len(data); i++ {
		if data[i].Id == iddel {
			break
		}
	}
	return i, nil
}

func EditTask(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(conf.PathTempl + "edittask.html")
	if err != nil {
		log.Println(err)
	}
	var t tasks.Tasks
	listT, _ := t.List()
	i, err := IndexJson(listT, r)
	if err != nil {
		log.Println(err)
	}
	if r.Method == "POST" {
		PageData(&listT[i], r)
		newData, err := json.MarshalIndent(&listT, "", " ")
		if err != nil {
			log.Println(err)
		}
		ioutil.WriteFile(conf.PathTasks, newData, 0666)

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else {
		tmpl.ExecuteTemplate(w, "edit", listT[i])
	}
}

func DelTask(w http.ResponseWriter, r *http.Request) {

	var t tasks.Tasks
	listT, _ := t.List()
	i, err := IndexJson(listT, r)
	if err != nil {
		log.Println(err)
	}
	id := listT[i].Id
	listT = append(listT[:i], listT[i+1:]...)
	newData, err := json.MarshalIndent(&listT, "", " ")
	if err != nil {
		log.Println(err)
	}
	ioutil.WriteFile(conf.PathTasks, newData, 0666)
	log.Println("Entry with id =", id, " deleted")
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
