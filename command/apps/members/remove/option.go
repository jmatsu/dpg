package apps_members_remove

import (
	"github.com/urfave/cli"
)

type optionName string

const (
	apiToken     optionName = "token"
	appOwnerName optionName = "app-owner"
	appId        optionName = "app-id"
	appPlatform  optionName = "app-platform"
	removees     optionName = "removees"
)

func allFlags() []cli.Flag {
	return []cli.Flag{
		apiToken.Flag(),
		appOwnerName.Flag(),
		appId.Flag(),
		appPlatform.Flag(),
		removees.Flag(),
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
	case appOwnerName:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "[Required] An owner of applications",
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
	case removees:
		return cli.StringSliceFlag{
			Name:  name.String(),
			Usage: "[Required] Comma separated names or e-mails of those who you want to remove",
		}
	}

	panic("Option name mapping is not found")
}

func (name optionName) Value(c *cli.Context) interface{} {
	switch name {
	case
		apiToken,
		appOwnerName,
		appId,
		appPlatform:
		return c.String(name.String())
	case removees:
		return c.StringSlice(name.String())
	}

	panic("Option name mapping is not found")
}
