package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Notebook struct {
	Name      string `json:"name"`
	CreatedAt int    `json:"createdAt"`
}

type Entry struct {
	Name     string
	Date     time.Time
	Notebook string
}

func (e Entry) GetEntryName() string {
	return fmt.Sprintf("%02d-%02d-%02d", e.Date.Month(), e.Date.Day(), e.Date.Year())
}

func (e Entry) GetFilePath(vitaDir string) string {
	return fmt.Sprintf("%s/%s/%s.txt", vitaDir, e.Notebook, e.GetEntryName())
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

func GetTodayFile(vitaDir string, nbName string) string {
	// get today's date
	t := time.Now()

	entry := Entry{"default", t, nbName}

	fullPath := entry.GetFilePath(vitaDir)

	// check if file exists
	if _, err := os.Stat(fullPath); err == nil {
		// if exists return filepath
		return fullPath
	}

	// if doesn't exist, create file with template, then return
	temp := template.Must(template.New("entry.tmpl").ParseFiles("entry.tmpl"))

	f, err := os.Create(fullPath)
	if err != nil {
		color.Red("No notebook found with name: %s\n", nbName)
		os.Exit(1)
	}

	err = temp.Execute(f, entry)
	if err != nil {
		os.Exit(1)
	}

	f.Close()
	return fullPath
}
