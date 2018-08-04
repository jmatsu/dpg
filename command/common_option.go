package command

import (
	"github.com/sirupsen/logrus"
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
		return "token"
	}

	panic("Option name mapping is not found")
}

func GetApiToken(c *cli.Context) string {
	token := c.String(ApiToken.name())

	logrus.Debugf("Token %s\n", token)

	return token
}
