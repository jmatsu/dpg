package command

import (
	"github.com/jmatsu/dpg/command/constant"
	"github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
	"os"
)

type option int

const (
	ApiToken option = iota
)

func (o option) Flag() cli.Flag {
	switch o {
	case ApiToken:
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] API token",
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
	var token string

	if c.IsSet(ApiToken.name()) {
		token = c.String(ApiToken.name())
	} else {
		token = os.Getenv("DEPLOYGATE_API_TOKEN")
	}

	logrus.Debugf("Token %s\n", token)

	return token
}
