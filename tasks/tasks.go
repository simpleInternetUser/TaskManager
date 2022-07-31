package tasks

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Task struct {
	Id          int
	Status      string
	Title       string
	Description string
	Date        string
}

type AllTasks struct {
	TasksA []*Task
}

func GetAllTasks() *AllTasks {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	datfile, err := ioutil.ReadAll(file)
	var aT AllTasks
	json.Unmarshal(datfile, &aT.TasksA)
	if err != nil {
		log.Fatal(err)
	}
	return &aT
}
