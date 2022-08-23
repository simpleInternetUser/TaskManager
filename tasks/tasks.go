package tasks

import (
	"encoding/json"
	"fmt"
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
	//Create(Task) (Task, error)
	Update() (Task, error)
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

func (t AllTasks) Get(id int) (*Task, error) {
	for i := 0; i < len(t.TasksA); i++ {
		if t.TasksA[i].Id == id {
			return &t.TasksA[i], nil
		}
	}
	return nil, fmt.Errorf("no record with id = %d", id)
}

//func (t AllTasks) Create(Task) (Task, error)

func (t AllTasks) Update() error {
	conf, err := config.ReadConfig("config/config.json")
	if err != nil {
		return err
	}
	newData, err := json.MarshalIndent(&t.TasksA, "", " ")
	if err != nil {
		return err
	}
	ioutil.WriteFile(conf.PathTasks, newData, 0666)
	return nil
}
func (t AllTasks) Delete(id int) (AllTasks, error) {
	for i := 0; i < len(t.TasksA); i++ {
		if t.TasksA[i].Id == id {
			t.TasksA = append(t.TasksA[:i], t.TasksA[i+1:]...)
			return t, nil
		}
	}

	return t, fmt.Errorf("no record with id = %d could not be deleted", id)
}
