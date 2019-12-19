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
	"github.com/jmatsu/dpg/command/procedure/app_manage"
	"github.com/jmatsu/dpg/version"
	"gopkg.in/urfave/cli.v2"
	"strconv"
)

func main() {
	if b, err := strconv.ParseBool(os.Getenv("DPG_DEBUG")); err == nil && b {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.ErrorLevel)
	}

	cli.AppHelpTemplate = fmt.Sprintf(`%s
COMPLETION:
	dpg --init-completion <bash|zsh>

WEBSITE:
	https://github.com/jmatsu/dpg

SUPPORT:
	https://github.com/jmatsu/dpg/issues
`, cli.AppHelpTemplate)

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(version.Template())
	}

	app := &cli.App{}
	app.Name = "dpg"
	app.Usage = "Unofficial DeployGate API Client CLI"
	app.Description = "dpg is an unofficial command line tool to access DeployGate API."
	app.Version = version.Version
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		{
			Name:  "app",
			Usage: "Application-based Operation API",
			Subcommands: []*cli.Command{
				apps.UploadCommand(),
				{
					Name:  "member",
					Usage: "Application-based Member API",
					Subcommands: []*cli.Command{
						members.AddCommand(),
						members.ListCommand(),
						members.RemoveCommand(),
					},
				},
				{
					Name:  "team",
					Usage: "Application-based Team API",
					Subcommands: []*cli.Command{
						teams.AddCommand(),
						teams.RemoveCommand(),
						teams.ListCommand(),
					},
				},
				{
					Name:  "shared-team",
					Usage: "Application-based Shared Team API",
					Subcommands: []*cli.Command{
						shared_teams.AddCommand(),
						shared_teams.RemoveCommand(),
						shared_teams.ListCommand(),
					},
				},
				{
					Name:  "distributions",
					Usage: "Application-based Distribution API",
					Subcommands: []*cli.Command{
						appDistributions.DestroyCommand(),
					},
				},
			},
		},
		{
			Name:  "distribution",
			Usage: "Distribution-based Operation API",
			Subcommands: []*cli.Command{
				distributions.DestroyCommand(),
			},
		},
		{
			Name:    "organization",
			Usage:   "Organization-based Operation API",
			Aliases: []string{"group"},
			Subcommands: []*cli.Command{
				organizations.CreateCommand(),
				organizations.DestroyCommand(),
				organizations.ListCommand(),
				organizations.ShowCommand(),
				organizations.UpdateCommand(),
				{
					Name:  "member",
					Usage: "Organization-based Member API",
					Subcommands: []*cli.Command{
						organizationMembers.AddCommand(),
						organizationMembers.RemoveCommand(),
						organizationMembers.ListCommand(),
					},
				},
				{
					Name:  "team",
					Usage: "Organization-based Team API",
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
			Usage: "Enterprise-based Operation API",
			Subcommands: []*cli.Command{
				{
					Name:  "member",
					Usage: "Enterprise-based Member API",
					Subcommands: []*cli.Command{
						enterpriseMembers.AddCommand(),
						enterpriseMembers.RemoveCommand(),
						enterpriseMembers.ListCommand(),
					},
				},
				{
					Name:  "organization",
					Usage: "Enterprise-based Organization API",
					Subcommands: []*cli.Command{
						{
							Name:  "members",
							Usage: "Enterprise-based Organization Member API",
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
					Usage: "Enterprise-based Shared Team API",
					Subcommands: []*cli.Command{
						enterpriseSharedTeams.AddCommand(),
						enterpriseSharedTeams.RemoveCommand(),
						enterpriseSharedTeams.ListCommand(),
					},
				},
			},
		},
		{
			Name:  "procedure",
			Usage: "Procedures based on combined API calls",
			Subcommands: []*cli.Command{
				{
					Name:  "app-manage",
					Usage: "Procedures to manage applications especially on CI (experimental)",
					Subcommands: []*cli.Command{
						app_manage.OnExposeCommand(),
						app_manage.OnFeatureBranchCommand(),
						app_manage.OnDeployBranchCommand(),
					},
				},
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		logrus.Warnln(err.Error())
		os.Exit(1)
	}
}
