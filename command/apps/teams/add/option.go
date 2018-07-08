package apps_teams_add

import (
	"github.com/urfave/cli"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/jmatsu/dpg/command/apps/teams"
)

type option string

const (
	teamName option = "team-name"
)

func flags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		teams.AppOwnerName.Flag(),
		apps.AppId.Flag(),
		apps.Android.Flag(),
		apps.IOS.Flag(),
		teamName.flag(),
	}
}

func (o option) name() string {
	switch o {
	case teamName:
		return "team-name"
	}

	panic("Option name mapping is not found")
}

func (o option) flag() cli.Flag {
	switch o {
	case teamName:
		return cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] A team name to be added the application",
		}
	}

	panic("Option name mapping is not found")
}

func getTeamName(c *cli.Context) string {
	return c.String(teamName.name())
}
