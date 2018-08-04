package organizations

import (
	"errors"
	"fmt"
	"gopkg.in/urfave/cli.v2"
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
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] The name of the organization",
		}
	}

	panic("Option name mapping is not found")
}

func GetOrganizationName(c *cli.Context) string {
	return c.String(OrganizationName.name())
}

func RequireOrganizationName(name string) error {
	if name != "" {
		return nil
	}

	return errors.New(fmt.Sprintf("--%s must not be empty", OrganizationName.name()))
}
