package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
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
				fmt.Println("Here we go")
				// get default notebook
				// generate today's date, check if file exists or not
				return nil
			},
		},
	}
	app.Run(os.Args)
}
