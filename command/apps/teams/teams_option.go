package teams

import (
	"github.com/jmatsu/dpg/command/apps"
	"github.com/urfave/cli"
)

type option int

const (
	AppOwnerName option = iota
	TeamName
)

func (o option) name() string {
	switch o {
	case AppOwnerName:
		return apps.AppOwnerName.Name()
	case TeamName:
		return "team-name"
	}

	panic("Option name mapping is not found")
}

func (o option) Flag() cli.Flag {
	switch o {
	case AppOwnerName:
		return cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] An owner of application(s). Only group is allowed.",
		}
	case TeamName:
		return cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] A team name to be operated",
		}
	}

	panic("Option name mapping is not found")
}

func GetAppOwnerName(c *cli.Context) string {
	return c.String(AppOwnerName.name())
}

func GetTeamName(c *cli.Context) string {
	return c.String(TeamName.name())
}
