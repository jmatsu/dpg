package teams

import (
	"github.com/urfave/cli"
	"github.com/jmatsu/dpg/command/apps"
)

type option int

const (
	AppOwnerName option = iota
)

func (o option) name() string {
	switch o {
	case AppOwnerName:
		return apps.AppOwnerName.Name()
	}

	panic("Option name mapping is not found")
}

func (o option) Flag() cli.Flag {
	switch o {
	case AppOwnerName:
		return cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] An owner of application(s). Only group is allowed.",
		}
	}

	panic("Option name mapping is not found")
}

func GetAppOwnerName(c *cli.Context) string {
	return c.String(AppOwnerName.name())
}
