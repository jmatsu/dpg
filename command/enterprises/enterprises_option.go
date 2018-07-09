package enterprises

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
)

type option int

const (
	EnterpriseName option = iota
)

func (o option) name() string {
	switch o {
	case EnterpriseName:
		return "enterprise-name"
	}

	panic("Option name mapping is not found")
}

func (o option) Flag() cli.Flag {
	switch o {
	case EnterpriseName:
		return cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] The name of the target enterprise.",
		}
	}

	panic("Option name mapping is not found")
}

func GetEnterpriseName(c *cli.Context) string {
	return c.String(EnterpriseName.name())
}

func RequireEnterpriseName(name string) error {
	if name != "" {
		return nil
	}

	return errors.New(fmt.Sprintf("--%s must not be empty", EnterpriseName.name()))
}
