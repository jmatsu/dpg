package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/jmatsu/dpg/command/apps/teams/list"
	"github.com/jmatsu/dpg/command/apps/invite"
	"github.com/jmatsu/dpg/command/apps/upload"
	"github.com/jmatsu/dpg/command/apps/members/list"
	"github.com/jmatsu/dpg/command/apps/members/remove"
	"github.com/jmatsu/dpg/command/apps/teams/add"
)

func main() {
	app := cli.NewApp()
	app.Name = "dpg"
	app.Usage = "DeployGate API client CLI"

	app.Commands = []cli.Command{
		apps_upload.Command(),
		apps_invite.Command(),
		apps_members_list.Command(),
		apps_members_remove.Command(),
		apps_teams_list.Command(),
		apps_teams_add.Command(),
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
