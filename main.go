package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ws"
	app.Usage = "make your workspace"

	app.Flags = []cli.Flag{
		cli.BoolTFlag{
			Name:  "verbose",
			Usage: "print responses to stdout",
		},
	}

	app.Commands = []cli.Command{
		Upload(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
