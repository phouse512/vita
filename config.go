package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os/user"
)

type Config struct {
	DefaultNotebook string `json:"defaultNotebook"`
	VitaDir         string `json:"vitaDir"`
}

func getHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func LoadConfiguration() (Config, error) {
	defaultPath := fmt.Sprintf("%s/.vitarc", getHomeDir())
	raw, err := ioutil.ReadFile(defaultPath)
	if err != nil {
		return Config{}, errors.New(fmt.Sprintf("No config file found at %s", defaultPath))
	}

	var conf Config
	err = json.Unmarshal(raw, &conf)
	if err != nil {
		return Config{}, errors.New("Invalid .vitarc file.")
	}
	return conf, nil
}
