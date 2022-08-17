package tasks

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"github.com/simpleInternetUser/TaskManager/config"
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

	conf, err := config.ReadConfig("config/config.json")
	if err != nil {
		log.Println(err)
	}
	datfile, err := ioutil.ReadFile(conf.PathTasks)
	if err != nil {
		log.Println(err)
	}
	var aT AllTasks
	err = json.Unmarshal(datfile, &aT.TasksA)
	if err != nil {
		log.Println(err)
	}
	return aT
}
