package teams

import (
	"errors"
	"fmt"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/jmatsu/dpg/command/constant"
	"gopkg.in/urfave/cli.v2"
)

type packageOption int

const (
	teamName packageOption = iota
)

func (o packageOption) name() string {
	switch o {
	case teamName:
		return constant.TeamName
	}

	panic("Option name mapping is not found")
}

func (o packageOption) flag() cli.Flag {
	switch o {
	case teamName:
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] The name of the target team",
		}
	}

	panic("Option name mapping is not found")
}

func getTeamName(c *cli.Context) string {
	return c.String(teamName.name())
}

func requireTeamName(name string) error {
	if name != "" {
		return nil
	}

	return errors.New(fmt.Sprintf("--%s must not be empty", teamName.name()))
}

func addFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		apps.AppOwnerName.Flag(),
		apps.AppId.Flag(),
		apps.Android.Flag(),
		apps.IOS.Flag(),
		teamName.flag(),
	}
}

func listFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		apps.AppOwnerName.Flag(),
		apps.AppId.Flag(),
		apps.Android.Flag(),
		apps.IOS.Flag(),
	}
}

func removeFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		apps.AppOwnerName.Flag(),
		apps.AppId.Flag(),
		apps.Android.Flag(),
		apps.IOS.Flag(),
		teamName.flag(),
	}
}
