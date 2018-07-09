package shared_teams

import (
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/enterprises"
	"github.com/urfave/cli"
)

type packageOption int

const (
	sharedTeamName packageOption = iota
)

func (o packageOption) name() string {
	switch o {
	case sharedTeamName:
		return "team-name"
	}

	panic("Option name mapping is not found")
}

func (o packageOption) flag() cli.Flag {
	switch o {
	case sharedTeamName:
		return cli.StringSliceFlag{
			Name:  o.name(),
			Usage: "A name of a shared team",
		}
	}

	panic("Option name mapping is not found")
}

func getSharedTeamName(c *cli.Context) string {
	return c.String(sharedTeamName.name())
}

func addFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		enterprises.EnterpriseName.Flag(),
		sharedTeamName.flag(),
	}
}

func listFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		enterprises.EnterpriseName.Flag(),
	}
}

func removeFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		enterprises.EnterpriseName.Flag(),
		sharedTeamName.flag(),
	}
}
