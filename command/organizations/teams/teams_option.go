package teams

import (
	"errors"
	"fmt"
	"github.com/jmatsu/dpg/command/constant"
	"gopkg.in/urfave/cli.v2"
)

type option int

const (
	TeamName option = iota
)

func (o option) name() string {
	switch o {
	case TeamName:
		return constant.TeamName
	}

	panic("Option name mapping is not found")
}

func (o option) Flag() cli.Flag {
	switch o {
	case TeamName:
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] The name of the target team",
		}
	}

	panic("Option name mapping is not found")
}

func GetTeamName(c *cli.Context) string {
	return c.String(TeamName.name())
}

func RequireTeamName(name string) error {
	if name != "" {
		return nil
	}

	return errors.New(fmt.Sprintf("--%s must not be empty", TeamName.name()))
}
