package distributions

import (
	"github.com/urfave/cli"
)

type option int

const (
	DistributionName option = iota
)

func (o option) name() string {
	switch o {
	case DistributionName:
		return "distribution-name"
	}

	panic("Option name mapping is not found")
}

func (o option) Flag() cli.Flag {
	switch o {
	case DistributionName:
		return cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] The name of the target distribution.",
		}
	}

	panic("Option name mapping is not found")
}

func GetDistributionName(c *cli.Context) string {
	return c.String(DistributionName.name())
}
