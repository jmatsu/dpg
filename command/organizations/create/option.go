package organizations_create

import (
	"github.com/urfave/cli"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/organizations"
	"gopkg.in/guregu/null.v3"
)

type option int

const (
	description option = iota
)

func flags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		organizations.OrganizationName.Flag(),
		description.flag(),
	}
}

func (o option) name() string {
	switch o {
	case description:
		return "description"
	}

	panic("Option name mapping is not found")
}

func (o option) flag() cli.Flag {
	switch o {
	case description:
		return cli.StringFlag{
			Name:  o.name(),
			Usage: "A description for an organization",
		}
	}

	panic("Option name mapping is not found")
}

func getDescription(c *cli.Context) null.String {
	if x := c.String(description.name()); c.IsSet(description.name()) {
		return null.StringFrom(x)
	} else {
		return null.StringFromPtr(nil)
	}
}
