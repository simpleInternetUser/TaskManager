package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Configuration struct {
	ServerPort string
	PathTempl  string
	PathTasks  string
}

func ReadConfig(FileName string) (Configuration, error) {
	cofigF, err := ioutil.ReadFile(FileName)
	if err != nil {
		return Configuration{}, errors.New("Unable to read config file.")
	}
	var config Configuration
	err = json.Unmarshal(cofigF, &config)
	if err != nil {
		return Configuration{}, errors.New("Invalid JSON file.")
	}
	return config, nil
}
