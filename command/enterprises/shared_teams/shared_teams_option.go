package shared_teams

import (
	"errors"
	"fmt"
	"github.com/jmatsu/dpg/command/constant"
	"gopkg.in/urfave/cli.v2"
)

type option int

const (
	SharedTeamName option = iota
)

func (o option) name() string {
	switch o {
	case SharedTeamName:
		return constant.SharedTeamName
	}

	panic("Option name mapping is not found")
}

func (o option) Flag() cli.Flag {
	switch o {
	case SharedTeamName:
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "The name of the shared team",
		}
	}

	panic("Option name mapping is not found")
}

func GetSharedTeamName(c *cli.Context) string {
	return c.String(SharedTeamName.name())
}

func RequireSharedTeamName(name string) error {
	if name != "" {
		return nil
	}

	return errors.New(fmt.Sprintf("--%s must not be empty", SharedTeamName.name()))
}
