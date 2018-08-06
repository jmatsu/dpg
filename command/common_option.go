package command

import (
	"github.com/jmatsu/dpg/command/constant"
	"gopkg.in/urfave/cli.v2"
)

type option int

const (
	ApiToken option = iota
)

func (o option) Flag() cli.Flag {
	switch o {
	case ApiToken:
		return &cli.StringFlag{
			Name:    o.name(),
			Usage:   "[Required] API token",
			EnvVars: []string{constant.ApiTokenEnv, constant.DeployGateApiTokenEnv},
		}
	}

	panic("Option name mapping is not found")
}

func (o option) name() string {
	switch o {
	case
		ApiToken:
		return constant.ApiToken
	}

	panic("Option name mapping is not found")
}

func GetApiToken(c *cli.Context) string {
	if c.IsSet(ApiToken.name()) {
		return c.String(ApiToken.name())
	} else {
		return ""
	}
}
