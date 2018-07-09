package main

import (
	"github.com/sirupsen/logrus"
	"os"

	"fmt"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/jmatsu/dpg/command/apps/members"
	"github.com/jmatsu/dpg/command/apps/shared_teams"
	"github.com/jmatsu/dpg/command/apps/teams"
	"github.com/jmatsu/dpg/command/distributions"
	"github.com/jmatsu/dpg/command/organizations"
	members2 "github.com/jmatsu/dpg/command/organizations/members"
	members3 "github.com/jmatsu/dpg/command/organizations/teams/members"
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
				apps.UploadCommand(),
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
						teams.AddCommand(),
						teams.RemoveCommand(),
						teams.ListCommand(),
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
				{
					Name:  "members",
					Usage: "Member-based Operation API",
					Subcommands: []cli.Command{
						members2.AddCommand(),
						members2.RemoveCommand(),
						members2.ListCommand(),
					},
				},
				{
					Name:  "teams",
					Usage: "Team-based Operation API",
					Subcommands: []cli.Command{
						{
							Name:  "members",
							Usage: "Member-based Operation API",
							Subcommands: []cli.Command{
								members3.AddCommand(),
								members3.RemoveCommand(),
								members3.ListCommand(),
							},
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
