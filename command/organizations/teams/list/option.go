package organizations_teams_list

import (
	"github.com/urfave/cli"
)

type optionName string

const (
	apiToken         optionName = "token"
	organizationName optionName = "organization-name"
	appId            optionName = "apps-id"
	appPlatform      optionName = "apps-platform"
)

func allFlags() []cli.Flag {
	return []cli.Flag{
		apiToken.Flag(),
		organizationName.Flag(),
		appId.Flag(),
		appPlatform.Flag(),
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
	case appPlatform:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "[Required] Either of android or iOS (case insensitive)",
		}
	}

	panic("Option name mapping is not found")
}

func (name optionName) Value(c *cli.Context) interface{} {
	switch name {
	case
		apiToken,
		organizationName,
		appId,
		appPlatform:
		return c.String(name.String())
	}

	panic("Option name mapping is not found")
}
