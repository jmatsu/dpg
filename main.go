package main

import (
	"github.com/sirupsen/logrus"
	"os"

	"github.com/jmatsu/dpg/command/apps/members"
	"github.com/jmatsu/dpg/command/apps/shared_teams"
	"github.com/jmatsu/dpg/command/apps/teams/add"
	"github.com/jmatsu/dpg/command/apps/teams/list"
	"github.com/jmatsu/dpg/command/apps/teams/remove"
	"github.com/jmatsu/dpg/command/apps/upload"
	"github.com/jmatsu/dpg/command/distributions"
	"github.com/jmatsu/dpg/command/organizations"
	"github.com/urfave/cli"
	"strconv"
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
						members.AddCommand(),
						members.ListCommand(),
						members.RemoveCommand(),
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
						shared_teams.AddCommand(),
						shared_teams.RemoveCommand(),
						shared_teams.ListCommand(),
					},
				},
			},
		},
		{
			Name:  "distribution",
			Usage: "Distribution-based Operation API",
			Subcommands: []cli.Command{
				distributions.DestroyCommand(),
			},
		},
		{
			Name:  "organization",
			Usage: "Organization-based Operation API",
			Subcommands: []cli.Command{
				organizations.CreateCommand(),
				organizations.DestroyCommand(),
				organizations.ListCommand(),
				organizations.ShowCommand(),
				organizations.UpdateCommand(),
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		logrus.Errorln(err.Error())
	}
}
