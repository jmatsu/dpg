package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/jmatsu/dpg/command/apps/upload"
	"github.com/jmatsu/dpg/command/apps/invite"
	"github.com/jmatsu/dpg/command/apps/users"
)

func main() {
	app := cli.NewApp()
	app.Name = "dpg"
	app.Usage = "DeployGate API client CLI"

	app.Flags = []cli.Flag{
		cli.BoolTFlag{
			Name:  "verbose",
			Usage: "print responses to stdout. (true by default)",
		},
	}

	app.Commands = []cli.Command{
		upload.Command(),
		invite.Command(),
		users.Command(),
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
