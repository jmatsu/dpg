package shared_teams

import (
	"errors"
	"fmt"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
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
		return cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] A name of a target team",
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
		apps.AppOwnerName.Flag(),
		apps.AppId.Flag(),
		apps.Android.Flag(),
		apps.IOS.Flag(),
		sharedTeamName.flag(),
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
		sharedTeamName.flag(),
	}
}
