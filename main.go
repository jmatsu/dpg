package main

import (
	"github.com/sirupsen/logrus"
	"os"

	"fmt"
	"github.com/jmatsu/dpg/command/apps"
	appDistributions "github.com/jmatsu/dpg/command/apps/distributions"
	"github.com/jmatsu/dpg/command/apps/members"
	"github.com/jmatsu/dpg/command/apps/shared_teams"
	"github.com/jmatsu/dpg/command/apps/teams"
	"github.com/jmatsu/dpg/command/distributions"
	enterpriseMembers "github.com/jmatsu/dpg/command/enterprises/members"
	enterpriseOrgMembers "github.com/jmatsu/dpg/command/enterprises/organizations/members"
	enterpriseSharedTeams "github.com/jmatsu/dpg/command/enterprises/shared_teams"
	"github.com/jmatsu/dpg/command/organizations"
	orgMembers "github.com/jmatsu/dpg/command/organizations/members"
	orgTeamMembers "github.com/jmatsu/dpg/command/organizations/teams/members"
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
				{
					Name:  "distributions",
					Usage: "Application-based Distribution API",
					Subcommands: []cli.Command{
						appDistributions.DestroyCommand(),
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
					Name:  "member",
					Usage: "Organization-based Member API",
					Subcommands: []cli.Command{
						orgMembers.AddCommand(),
						orgMembers.RemoveCommand(),
						orgMembers.ListCommand(),
					},
				},
				{
					Name:  "team",
					Usage: "Organization-based Team API",
					Subcommands: []cli.Command{
						{
							Name:  "member",
							Usage: "Organization-based Team Member API",
							Subcommands: []cli.Command{
								orgTeamMembers.AddCommand(),
								orgTeamMembers.RemoveCommand(),
								orgTeamMembers.ListCommand(),
							},
						},
					},
				},
			},
		},
		{
			Name:  "enterprise",
			Usage: "Enterprise-based Operation API",
			Subcommands: []cli.Command{
				{
					Name:  "member",
					Usage: "Enterprise-based Member API",
					Subcommands: []cli.Command{
						enterpriseMembers.AddCommand(),
						enterpriseMembers.RemoveCommand(),
						enterpriseMembers.ListCommand(),
					},
				},
				{
					Name:  "organization",
					Usage: "Enterprise-based Organization API",
					Subcommands: []cli.Command{
						{
							Name:  "members",
							Usage: "Enterprise-based Organization Member API",
							Subcommands: []cli.Command{
								enterpriseOrgMembers.AddCommand(),
								enterpriseOrgMembers.RemoveCommand(),
								enterpriseOrgMembers.ListCommand(),
							},
						},
					},
				},
				{
					Name:  "shared-team",
					Usage: "Enterprise-based Shared Team API",
					Subcommands: []cli.Command{
						enterpriseSharedTeams.AddCommand(),
						enterpriseSharedTeams.RemoveCommand(),
						enterpriseSharedTeams.ListCommand(),
					},
				},
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
