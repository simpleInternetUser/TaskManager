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

type Tasks interface {
	Get(id int) (Task, error)
	List() ([]Task, error)
	Create(Task) (Task, error)
	Update(Task) (Task, error)
	Delete(id int) error
}

func (t Task) List() ([]Task, error) {

	conf, err := config.ReadConfig("config/config.json")
	if err != nil {
		return nil, err
	}
	datfile, err := ioutil.ReadFile(conf.PathTasks)
	if err != nil {
		return nil, err
	}
	var aT []Task
	err = json.Unmarshal(datfile, &aT)
	if err != nil {
		return nil, err
	}
	return aT, nil
}
