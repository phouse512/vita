package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

func main() {
	conf, err := LoadConfiguration()
	if err != nil {
		log.Fatalf("Unable to load .vitarc, please check that file exists and is correct.")
	}

	app := cli.NewApp()
	app.Name = "vita"
	app.Usage = "Here's how you use vita"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Not sure.")
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:  "notebooks",
			Usage: "list all notebooks",
			Action: func(c *cli.Context) error {
				GetNotebooks(conf.VitaDir)
				return nil
			},
		},
		{
			Name:    "today",
			Aliases: []string{},
			Usage:   "Get the filepath of today's notebook entry.",
			Action: func(c *cli.Context) error {
				fmt.Printf("Loading default notebook, %s\n", conf.DefaultNotebook)
				// get default notebook
				// generate today's date, check if file exists or not
				t := time.Now()
				fmt.Printf("%02d-%02d-%d.md\n", t.Month(), t.Day(), t.Year())
				return nil
			},
		},
	}
	app.Run(os.Args)
}
