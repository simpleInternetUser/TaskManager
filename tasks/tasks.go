package tasks

import (
	"encoding/json"
	"io/ioutil"
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

type Tasks interface {
	Get(id int) (Task, error)
	List() (AllTasks, error)
	Create(Task) (Task, error)
	Update(Task) (Task, error)
	Delete(id int) error
}

func (aT AllTasks) List() (AllTasks, error) {

	conf, err := config.ReadConfig("config/config.json")
	if err != nil {
		return AllTasks{}, err
	}
	datfile, err := ioutil.ReadFile(conf.PathTasks)
	if err != nil {
		return AllTasks{}, err
	}

	err = json.Unmarshal(datfile, &aT.TasksA)
	if err != nil {
		return AllTasks{}, err
	}
	return aT, nil
}

//func (t AllTasks) Get(id int) (Task, error)
