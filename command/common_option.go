package command

import (
	"github.com/urfave/cli"
)

type option int

const (
	ApiToken option = iota
	//OrganizationName CommonOptionName = "organization-name"
	//TeamName CommonOptionName = "team-name"
)

func (o option) Flag() cli.Flag {
	switch o {
	case ApiToken:
		return cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] API token",
		}
		//case OrganizationName:
		//	return cli.StringFlag{
		//		name:  name.String(),
		//		Usage: "[Required] A name of an organization",
		//	}
	}

	panic("Option name mapping is not found")
}

func (o option) name() string {
	switch o {
	case
		ApiToken:
		return "token"
	}

	panic("Option name mapping is not found")
}

func GetApiToken(c *cli.Context) string {
	return c.String(ApiToken.name())
}
