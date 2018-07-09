package distributions

import (
	"github.com/urfave/cli"
)

type option int

const (
	DistributionKey option = iota
)

func (o option) name() string {
	switch o {
	case DistributionKey:
		return "distribution-key"
	}

	panic("Option name mapping is not found")
}

func (o option) Flag() cli.Flag {
	switch o {
	case DistributionKey:
		return cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] The key of the target distribution.",
		}
	}

	panic("Option name mapping is not found")
}

func GetDistributionKey(c *cli.Context) string {
	return c.String(DistributionKey.name())
}
