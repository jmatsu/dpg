package organizations

import (
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/constant"
	"gopkg.in/guregu/null.v3"
	"gopkg.in/urfave/cli.v2"
)

type packageOption int

const (
	description packageOption = iota
)

func (o packageOption) name() string {
	switch o {
	case description:
		return constant.Description
	}

	panic("Option name mapping is not found")
}

func (o packageOption) flag() cli.Flag {
	switch o {
	case description:
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "A description for the organization",
		}
	}

	panic("Option name mapping is not found")
}

func createFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		OrganizationName.Flag(),
		description.flag(),
	}
}

func getCreateDescription(c *cli.Context) null.String {
	return null.StringFrom(c.String(description.name()))
}

func destroyFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		OrganizationName.Flag(),
	}
}

func listFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
	}
}

func showFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		OrganizationName.Flag(),
	}
}

func updateFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		OrganizationName.Flag(),
		description.flag(),
	}
}

func getUpdateDescription(c *cli.Context) null.String {
	if x := c.String(description.name()); c.IsSet(description.name()) {
		return null.StringFrom(x)
	} else {
		return null.StringFromPtr(nil)
	}
}
