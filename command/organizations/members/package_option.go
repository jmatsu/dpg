package members

import (
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/urfave/cli"
	"gopkg.in/guregu/null.v3"
)

type packageOption int

const (
	userName packageOption = iota
	userEmail
)

func (o packageOption) name() string {
	switch o {
	case userName:
		return "username"
	case userEmail:
		return "email"
	}

	panic("Option name mapping is not found")
}

func (o packageOption) flag() cli.Flag {
	switch o {
	case userName:
		return cli.StringSliceFlag{
			Name:  o.name(),
			Usage: "[Either of this or email is required] A name of a user. ",
		}
	case userEmail:
		return cli.BoolFlag{
			Name:   o.name(),
			Usage:  "[Either of this or username is required] An email of a user",
			Hidden: true,
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
		userName.flag(),
		userEmail.flag(),
	}
}

func getUserName(c *cli.Context) null.String {
	return null.NewString(c.String(userName.name()), c.IsSet(userName.name()))
}

func getUserEmail(c *cli.Context) null.String {
	return null.NewString(c.String(userEmail.name()), c.IsSet(userEmail.name()))
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
		userName.flag(),
		userEmail.flag(),
	}
}

func getUserNameOrEmail(c *cli.Context) null.String {
	if c.IsSet(userName.name()) {
		return null.StringFrom(c.String(userName.name()))
	} else if c.IsSet(userEmail.name()) {
		return null.StringFrom(c.String(userEmail.name()))
	} else {
		return null.StringFromPtr(nil)
	}
}