package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type Notebook struct {
	Name      string `json:"name"`
	CreatedAt int    `json:"createdAt"`
}

func loadNotebook(path string) (Notebook, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return Notebook{}, errors.New("No meta.json file found.")
	}

	var notebook Notebook
	err = json.Unmarshal(raw, &notebook)
	if err != nil {
		return Notebook{}, errors.New("Invalid notebook json file.")
	}

	return notebook, nil
}

func GetNotebooks(vitaDir string) {
	files, err := ioutil.ReadDir(vitaDir)
	if err != nil {
		log.Fatal(err)
	}

	var nbs []Notebook
	for _, f := range files {
		if f.IsDir() {
			metaPath := fmt.Sprintf("%s/%s/meta.json", vitaDir, f.Name())
			nb, err := loadNotebook(metaPath)
			if err != nil {
				log.Printf("Received invalid notebook directory: %v", err)
				continue
			}
			nbs = append(nbs, nb)
		}
	}

	log.Printf("Found %d notebooks", len(nbs))
}

func GetTodayFile(vitaDir string, nbName string) {
	// get today's date
	t := time.Now()
	todayPath := fmt.Sprintf("%02d-%02d-%d.md\n", t.Month(), t.Day(), t.Year())
	fullPath := fmt.Sprintf("%s/%s/%s", vitaDir, nbName, todayPath)

	// check if file exists

	// if exists return filepath

	// if doesn't exist, create file with template, then return
}
