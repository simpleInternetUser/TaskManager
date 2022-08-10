package tasks

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Task struct {
	Id          int
	Title       string
	Status      string
	Description string
	Date        time.Time
}

type AllTasks struct {
	TasksA []Task
}

func GetAllTasks() AllTasks {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	datfile, err := ioutil.ReadAll(file)
	var aT AllTasks
	json.Unmarshal(datfile, &aT.TasksA)
	if err != nil {
		log.Fatal(err)
	}
	return aT
}
