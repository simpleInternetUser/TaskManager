package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	ServerPort string
	PathTempl  string
	PathTasks  string
}

func ReadConfig(FileName string) (Configuration, error) {
	cofigF, err := ioutil.ReadFile(FileName)
	if err != nil {
		log.Println("Unable to read config file.")
		return Configuration{}, err
	}
	var config Configuration
	err = json.Unmarshal(cofigF, &config)
	if err != nil {
		log.Println("Invalid JSON file.")
		return Configuration{}, err
	}
	return config, nil
}
