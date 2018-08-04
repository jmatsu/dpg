package organizations

import (
	"errors"
	"fmt"
	"github.com/jmatsu/dpg/command/constant"
	"gopkg.in/urfave/cli.v2"
)

type option int

const (
	OrganizationName option = iota
)

func (o option) name() string {
	switch o {
	case OrganizationName:
		return constant.OrganizationName
	}

	panic("Option name mapping is not found")
}

func (o option) Flag() cli.Flag {
	switch o {
	case OrganizationName:
		return &cli.StringSliceFlag{
			Name:    o.name(),
			Aliases: []string{constant.OrganizationNameAlias},
			Usage:   "The name of the target enterprise's organization",
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
