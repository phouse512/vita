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
			Subcommands: []cli.Command{},
		},
		{
			Name:    "today",
			Aliases: []string{},
			Usage:   "Get the filepath of today's notebook entry.",
			Action: func(c *cli.Context) error {
				// get default notebook
				// generate today's date, check if file exists or not
				if len(c.Args().First()) > 0 {
					fmt.Printf("%s\n", GetTodayFile(conf.VitaDir, c.Args().First()))
					return nil
				}
				fmt.Printf("%s\n", GetTodayFile(conf.VitaDir, conf.DefaultNotebook))
				return nil
			},
		},
	}
	app.Run(os.Args)
}
