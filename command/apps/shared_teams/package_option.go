package shared_teams

import (
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/jmatsu/dpg/command/apps/teams"
	"github.com/urfave/cli"
)

type packageOption int

const (
	appOwnerName packageOption = iota
	teamName
)

func (o packageOption) name() string {
	switch o {
	case appOwnerName:
		return apps.AppOwnerName.Name()
	case teamName:
		return "team-name"
	}

	panic("Option name mapping is not found")
}

func (o packageOption) flag() cli.Flag {
	switch o {
	case appOwnerName:
		return cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] An owner of application(s). Only group is allowed.",
		}
	case teamName:
		return cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] A team name to be operated",
		}
	}

	panic("Option name mapping is not found")
}

func getAppOwnerName(c *cli.Context) string {
	return c.String(appOwnerName.name())
}

func getTeamName(c *cli.Context) string {
	return c.String(teamName.name())
}

func addFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		teams.AppOwnerName.Flag(),
		apps.AppId.Flag(),
		apps.Android.Flag(),
		apps.IOS.Flag(),
		teamName.flag(),
	}
}

func listFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		teams.AppOwnerName.Flag(),
		apps.AppId.Flag(),
		apps.Android.Flag(),
		apps.IOS.Flag(),
	}
}

func removeFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		teams.AppOwnerName.Flag(),
		apps.AppId.Flag(),
		apps.Android.Flag(),
		apps.IOS.Flag(),
		teamName.flag(),
	}
}
