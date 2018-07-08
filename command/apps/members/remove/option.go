package apps_members_remove

import (
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/urfave/cli"
)

type option int

const (
	removees option = iota
)

func flags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		apps.AppOwnerName.Flag(),
		apps.AppId.Flag(),
		apps.Android.Flag(),
		apps.IOS.Flag(),
		removees.flag(),
	}
}

func (o option) name() string {
	switch o {
	case removees:
		return "removees"
	}

	panic("Option name mapping is not found")
}

func (o option) flag() cli.Flag {
	switch o {
	case removees:
		return cli.StringSliceFlag{
			Name:  o.name(),
			Usage: "[Required] Comma separated names or e-mails of those who you want to remove",
		}
	}

	panic("Option name mapping is not found")
}

func getRemovees(c *cli.Context) []string {
	return c.StringSlice(removees.name())
}
