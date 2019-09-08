package distributions

import (
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/jmatsu/dpg/command/constant"
	"gopkg.in/urfave/cli.v2"
)

type packageOption int

const (
	distributionName packageOption = iota
)

func (o packageOption) name() string {
	switch o {
	case distributionName:
		return constant.DistributionName
	}

	panic("Option name mapping is not found")
}

func (o packageOption) flag() cli.Flag {
	switch o {
	case distributionName:
		return &cli.StringFlag{
			Name:    o.name(),
			Usage:   "[Required] The name of the target distribution.",
			EnvVars: []string{constant.DistributionNameEnv},
		}
	}

	panic("Option name mapping is not found")
}

func destroyFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		apps.AppOwnerName.Flag(),
		apps.AppId.Flag(),
		apps.Android.Flag(),
		apps.IOS.Flag(),
		distributionName.flag(),
	}
}

func GetDistributionName(c *cli.Context) string {
	return c.String(distributionName.name())
}
