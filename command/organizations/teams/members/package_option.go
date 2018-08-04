package members

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
	userName
)

func (o packageOption) name() string {
	switch o {
	case teamName:
		return constant.TeamName
	case userName:
		return constant.UserName
	}

	panic("Option name mapping is not found")
}

func (o packageOption) flag() cli.Flag {
	switch o {
	case teamName:
		return &cli.StringSliceFlag{
			Name:  o.name(),
			Usage: "[Required] A name of a team",
		}
	case userName:
		return &cli.StringSliceFlag{
			Name:  o.name(),
			Usage: "[Required] A name of a user. ",
		}
	}

	panic("Option name mapping is not found")
}

func addFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		apps.AppOwnerName.Flag(),
		apps.AppId.Flag(),
		apps.Android.Flag(),
		apps.IOS.Flag(),
		teamName.flag(),
		userName.flag(),
	}
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

func getUserName(c *cli.Context) string {
	return c.String(userName.name())
}

func requireUserName(name string) error {
	if name != "" {
		return nil
	}

	return errors.New(fmt.Sprintf("--%s must not be empty", userName.name()))
}

func listFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		apps.AppOwnerName.Flag(),
		apps.AppId.Flag(),
		apps.Android.Flag(),
		apps.IOS.Flag(),
		teamName.flag(),
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
		userName.flag(),
	}
}
