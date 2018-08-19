package shared_teams

import (
	"errors"
	"fmt"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/constant"
	"github.com/jmatsu/dpg/command/enterprises"
	"gopkg.in/urfave/cli.v2"
)

type packageOption int

const (
	sharedTeamName packageOption = iota
	description
)

func (o packageOption) name() string {
	switch o {
	case sharedTeamName:
		return constant.SharedTeamName
	case description:
		return constant.Description
	}

	panic("Option name mapping is not found")
}

func (o packageOption) flag() cli.Flag {
	switch o {
	case sharedTeamName:
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "The name of the shared team",
		}
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "The description of the shared team",
		}
	}

	panic("Option name mapping is not found")
}

func getSharedTeamName(c *cli.Context) string {
	return c.String(sharedTeamName.name())
}

func requireSharedTeamName(name string) error {
	if name != "" {
		return nil
	}

	return errors.New(fmt.Sprintf("--%s must not be empty", sharedTeamName.name()))
}

func addFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		enterprises.EnterpriseName.Flag(),
		sharedTeamName.flag(),
		description.flag(),
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
