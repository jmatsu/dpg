package organizations_teams_list

import (
	"github.com/urfave/cli"
)

type optionName string

const (
	apiToken         optionName = "token"
	organizationName optionName = "organization-name"
	appId            optionName = "app-id"
	android          optionName = "android"
	ios              optionName = "ios"
)

func allFlags() []cli.Flag {
	return []cli.Flag{
		apiToken.Flag(),
		organizationName.Flag(),
		appId.Flag(),
		android.Flag(),
		ios.Flag(),
	}
}

func (name optionName) String() string {
	return string(name)
}

func (name optionName) Flag() cli.Flag {
	switch name {
	case apiToken:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "[Required] API token",
		}
	case organizationName:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "[Required] A name of an organization",
		}
	case appId:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "[Required] An application id to invite users. e.g. com.deploygate",
		}
	case android:
		return cli.BoolFlag{
			Name:  name.String(),
			Usage: "[Required] Either of this or ios flag must be specified",
		}
	case ios:
		return cli.BoolFlag{
			Name:  name.String(),
			Usage: "[Required] Either of this or android flag must be specified",
		}
	}

	panic("Option name mapping is not found")
}

func (name optionName) Value(c *cli.Context) interface{} {
	switch name {
	case
		apiToken,
		organizationName,
		appId:
		return c.String(name.String())
	case
		android,
		ios:
		return c.Bool(name.String())
	}

	panic("Option name mapping is not found")
}
