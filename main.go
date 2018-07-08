package main

import (
	"github.com/sirupsen/logrus"
	"os"

	"github.com/jmatsu/dpg/command/apps/members/add"
	"github.com/jmatsu/dpg/command/apps/members/list"
	"github.com/jmatsu/dpg/command/apps/members/remove"
	"github.com/jmatsu/dpg/command/apps/teams/add"
	"github.com/jmatsu/dpg/command/apps/teams/list"
	"github.com/jmatsu/dpg/command/apps/teams/remove"
	"github.com/jmatsu/dpg/command/apps/upload"
	"github.com/jmatsu/dpg/command/distributions/remove"
	"github.com/jmatsu/dpg/command/organizations/create"
	"github.com/jmatsu/dpg/command/organizations/destroy"
	"github.com/jmatsu/dpg/command/organizations/list"
	"github.com/jmatsu/dpg/command/organizations/show"
	"github.com/jmatsu/dpg/command/organizations/update"
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
			Usage: "Distribution-based Operation API",
			Subcommands: []cli.Command{
				distributions_destroy.Command(),
			},
		},
		{
			Name:  "organization",
			Usage: "Organization-based Operation API",
			Subcommands: []cli.Command{
				organizations_create.Command(),
				organizations_destroy.Command(),
				organizations_list.Command(),
				organizations_show.Command(),
				organizations_update.Command(),
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		logrus.Errorln(err.Error())
	}
}
