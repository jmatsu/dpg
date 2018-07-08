package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/jmatsu/dpg/command/organizations/teams/list"
	"github.com/jmatsu/dpg/command/apps/invite"
	"github.com/jmatsu/dpg/command/apps/upload"
	"github.com/jmatsu/dpg/command/apps/members/list"
	"github.com/jmatsu/dpg/command/apps/members/remove"
	"github.com/jmatsu/dpg/command/organizations/teams/add"
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
		apps_upload.Command(),
		apps_invite.Command(),
		apps_members_list.Command(),
		apps_members_remove.Command(),
		organizations_teams_list.Command(),
		organizations_teams_add.Command(),
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
