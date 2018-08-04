/*
	package main is an entry point of all commands
*/
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
	enterpriseOrganizationMembers "github.com/jmatsu/dpg/command/enterprises/organizations/members"
	enterpriseSharedTeams "github.com/jmatsu/dpg/command/enterprises/shared_teams"
	"github.com/jmatsu/dpg/command/organizations"
	organizationMembers "github.com/jmatsu/dpg/command/organizations/members"
	organizationTeamMembers "github.com/jmatsu/dpg/command/organizations/teams/members"
	"gopkg.in/urfave/cli.v2"
	"strconv"
)

func main() {
	if b, err := strconv.ParseBool(os.Getenv("DPG_DEBUG")); err != nil && b {
		logrus.SetLevel(logrus.DebugLevel)
	}

	app := (&cli.App{})
	app.Name = "dpg"
	app.Usage = "DeployGate API client CLI"
	app.Version = "0.1"
	app.EnableShellCompletion = true
	app.Commands = []*cli.Command{
		{
			Name:  "app",
			Usage: "application-based Operation API",
			Subcommands: []*cli.Command{
				apps.UploadCommand(),
				{
					Name:  "member",
					Usage: "application-based Member API",
					Subcommands: []*cli.Command{
						members.AddCommand(),
						members.ListCommand(),
						members.RemoveCommand(),
					},
				},
				{
					Name:  "team",
					Usage: "application-based Team API",
					Subcommands: []*cli.Command{
						teams.AddCommand(),
						teams.RemoveCommand(),
						teams.ListCommand(),
					},
				},
				{
					Name:  "shared-team",
					Usage: "application-based Shared Team API",
					Subcommands: []*cli.Command{
						shared_teams.AddCommand(),
						shared_teams.RemoveCommand(),
						shared_teams.ListCommand(),
					},
				},
				{
					Name:  "distributions",
					Usage: "application-based Distribution API",
					Subcommands: []*cli.Command{
						appDistributions.DestroyCommand(),
					},
				},
			},
		},
		{
			Name:  "distribution",
			Usage: "distribution-based Operation API",
			Subcommands: []*cli.Command{
				distributions.DestroyCommand(),
			},
		},
		{
			Name:  "organization",
			Usage: "organization-based Operation API",
			Subcommands: []*cli.Command{
				organizations.CreateCommand(),
				organizations.DestroyCommand(),
				organizations.ListCommand(),
				organizations.ShowCommand(),
				organizations.UpdateCommand(),
				{
					Name:  "member",
					Usage: "organization-based Member API",
					Subcommands: []*cli.Command{
						organizationMembers.AddCommand(),
						organizationMembers.RemoveCommand(),
						organizationMembers.ListCommand(),
					},
				},
				{
					Name:  "team",
					Usage: "organization-based Team API",
					Subcommands: []*cli.Command{
						{
							Name:  "member",
							Usage: "Organization-based Team Member API",
							Subcommands: []*cli.Command{
								organizationTeamMembers.AddCommand(),
								organizationTeamMembers.RemoveCommand(),
								organizationTeamMembers.ListCommand(),
							},
						},
					},
				},
			},
		},
		{
			Name:  "enterprise",
			Usage: "enterprise-based Operation API",
			Subcommands: []*cli.Command{
				{
					Name:  "member",
					Usage: "enterprise-based Member API",
					Subcommands: []*cli.Command{
						enterpriseMembers.AddCommand(),
						enterpriseMembers.RemoveCommand(),
						enterpriseMembers.ListCommand(),
					},
				},
				{
					Name:  "organization",
					Usage: "enterprise-based Organization API",
					Subcommands: []*cli.Command{
						{
							Name:  "members",
							Usage: "enterprise-based Organization Member API",
							Subcommands: []*cli.Command{
								enterpriseOrganizationMembers.AddCommand(),
								enterpriseOrganizationMembers.RemoveCommand(),
								enterpriseOrganizationMembers.ListCommand(),
							},
						},
					},
				},
				{
					Name:  "shared-team",
					Usage: "enterprise-based Shared Team API",
					Subcommands: []*cli.Command{
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
