package main

import (
	"github.com/sirupsen/logrus"
	"os"

	"github.com/urfave/cli"
	"github.com/jmatsu/dpg/command/apps/upload"
	"github.com/jmatsu/dpg/command/apps/members/list"
	"github.com/jmatsu/dpg/command/apps/members/remove"
	"github.com/jmatsu/dpg/command/apps/teams/list"
	"github.com/jmatsu/dpg/command/apps/teams/add"
	"github.com/jmatsu/dpg/command/apps/members/add"
	"github.com/jmatsu/dpg/command/apps/teams/remove"
	"strconv"
	"github.com/jmatsu/dpg/command/distributions/remove"
)

func main() {
	if b, err := strconv.ParseBool(os.Getenv("DPG_DEBUG")); err != nil && b {
		logrus.SetLevel(logrus.DebugLevel)
	}

	app := cli.NewApp()
	app.Name = "dpg"
	app.Usage = "DeployGate API client CLI"
	app.Commands = []cli.Command{
		{
			Name:  "app",
			Usage: "Application-based Operation API",
			Subcommands: []cli.Command{
				apps_upload.Command(),
				{
					Name:  "member",
					Usage: "Application-based Member API",
					Subcommands: []cli.Command{
						apps_members_add.Command(),
						apps_members_list.Command(),
						apps_members_remove.Command(),
					},
				},
				{
					Name:  "team",
					Usage: "Application-based Team API",
					Subcommands: []cli.Command{
						apps_teams_list.Command(),
						apps_teams_add.Command(),
						apps_teams_remove.Command(),
					},
				},
				{
					Name:  "shared-team",
					Usage: "Application-based Shared Team API",
					Subcommands: []cli.Command{
						apps_teams_list.Command(),
						apps_teams_add.Command(),
						apps_teams_remove.Command(),
					},
				},
			},
		},
		{
			Name:  "distribution",
			Usage: "Application-based Operation API",
			Subcommands: []cli.Command{
				distributions_remove.Command(),
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		logrus.Errorln(err.Error())
	}
}
