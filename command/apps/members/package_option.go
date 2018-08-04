package members

import (
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/jmatsu/dpg/command/constant"
	"gopkg.in/urfave/cli.v2"
)

type packageOption int

const (
	invitees packageOption = iota
	developerRole
	removees
)

func (o packageOption) name() string {
	switch o {
	case invitees:
		return constant.Invitees
	case developerRole:
		return constant.DeveloperRole
	case removees:
		return constant.Removees
	}

	panic("Option name mapping is not found")
}

func (o packageOption) flag() cli.Flag {
	switch o {
	case invitees:
		return &cli.StringSliceFlag{
			Name:  o.name(),
			Usage: "[Required] Comma-separated names or e-mails of those whom you want to invite",
		}
	case developerRole:
		return &cli.BoolFlag{
			Name:   o.name(),
			Usage:  "[Old Free/Lite/Pro/Biz plans only] Specify this if invitee(s) should be Developer Role, otherwise they would have Tester Role. Tester Role will be selected by default",
			Hidden: true,
		}
	case removees:
		return &cli.StringSliceFlag{
			Name:  o.name(),
			Usage: "[Required] Comma-separated names or e-mails of those who you want to remove",
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
		invitees.flag(),
		developerRole.flag(),
	}
}

func getInvitees(c *cli.Context) []string {
	return c.StringSlice(invitees.name())
}

func isDeveloperRole(c *cli.Context) bool {
	return c.Bool(developerRole.name())
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
		removees.flag(),
	}
}

func getRemovees(c *cli.Context) []string {
	return c.StringSlice(removees.name())
}
