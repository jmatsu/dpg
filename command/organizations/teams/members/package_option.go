package members

import (
	"errors"
	"fmt"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/constant"
	"github.com/jmatsu/dpg/command/organizations"
	"github.com/jmatsu/dpg/command/organizations/teams"
	"gopkg.in/urfave/cli.v2"
)

type packageOption int

const (
	userName packageOption = iota
)

func (o packageOption) name() string {
	switch o {
	case userName:
		return constant.UserName
	}

	panic("Option name mapping is not found")
}

func (o packageOption) flag() cli.Flag {
	switch o {
	case userName:
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "The name of the organization's user",
		}
	}

	panic("Option name mapping is not found")
}

func addFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		organizations.OrganizationName.Flag(),
		teams.TeamName.Flag(),
		userName.flag(),
	}
}

func listFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		organizations.OrganizationName.Flag(),
		teams.TeamName.Flag(),
	}
}

func removeFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		organizations.OrganizationName.Flag(),
		teams.TeamName.Flag(),
		userName.flag(),
	}
}

func getUserName(c *cli.Context) string {
	return c.String(userName.name())
}

func requireUserName(name string) error {
	if name != "" {
		return nil
	}

	return errors.New(fmt.Sprintf("--%s must not be empty", userName.name()))
}
