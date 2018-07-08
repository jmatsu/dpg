package apps_invite

import (
	"github.com/urfave/cli"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
)

type option int

const (
	invitees      option = iota
	developerRole
)

func flags() []cli.Flag {
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

func (o option) name() string {
	switch o {
	case invitees:
		return "invitees"
	case developerRole:
		return "role"
	}

	panic("Option name mapping is not found")
}

func (o option) flag() cli.Flag {
	switch o {
	case invitees:
		return cli.StringSliceFlag{
			Name:  o.name(),
			Usage: "[Required] Comma separated names or e-mails of those who you want to invite",
		}
	case developerRole:
		return cli.BoolFlag{
			Name:   o.name(),
			Usage:  "[Old Free/Lite/Pro/Biz plans only] Specify this if invitee(s) should be the developer role, otherwise they would be the tester role. tester will be selected by default",
			Hidden: true,
		}
	}

	panic("Option name mapping is not found")
}

func getInvitees(c *cli.Context) []string {
	return c.StringSlice(invitees.name())
}

func isDeveloperRole(c *cli.Context) bool {
	return c.Bool(developerRole.name())
}
