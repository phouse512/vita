package main

import (
	"io/ioutil"
	"log"
)

func getNotebooks(vitaDir string) {
	files, err := ioutil.ReadDir(vitaDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			log.Println(f.Name())
		}
	}
}
