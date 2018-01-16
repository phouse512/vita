package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

type Configuration struct {
}

func main() {
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
				fmt.Println("Listing all existing notebooks: ")
				getNotebooks("/Users/philiphouse/os/go/src/github.com/phouse512/vita")
				return nil
			},
		},
	}
	app.Run(os.Args)
}
