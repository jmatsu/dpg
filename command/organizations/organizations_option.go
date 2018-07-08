package organizations

import (
	"github.com/urfave/cli"
)

type option int

const (
	OrganizationName option = iota
)

func (o option) name() string {
	switch o {
	case OrganizationName:
		return "organization-name"
	}

	panic("Option name mapping is not found")
}

func (o option) Flag() cli.Flag {
	switch o {
	case OrganizationName:
		return cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] A name of a distribution to be operated.",
		}
	}

	panic("Option name mapping is not found")
}

func GetOrganizationName(c *cli.Context) string {
	return c.String(OrganizationName.name())
}
