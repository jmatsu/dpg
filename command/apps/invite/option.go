package apps_invite

import (
	"github.com/urfave/cli"
)

type optionName string

const (
	apiToken      optionName = "token"
	appOwnerName  optionName = "app-owner"
	appId         optionName = "app-id"
	android       optionName = "android"
	ios           optionName = "ios"
	invitees      optionName = "invitees"
	developerRole optionName = "role"
)

func allFlags() []cli.Flag {
	return []cli.Flag{
		apiToken.Flag(),
		appOwnerName.Flag(),
		appId.Flag(),
		android.Flag(),
		ios.Flag(),
		invitees.Flag(),
		developerRole.Flag(),
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
	case invitees:
		return cli.StringSliceFlag{
			Name:  name.String(),
			Usage: "[Required] Comma separated names or e-mails of those who you want to invite",
		}
	case developerRole:
		return cli.BoolFlag{
			Name:   name.String(),
			Usage:  "[Old Free/Lite/Pro/Biz plans only] Specify this if invitee(s) should be the developer role, otherwise they would be the tester role. tester will be selected by default",
			Hidden: true,
		}
	}

	panic("Option name mapping is not found")
}

func (name optionName) Value(c *cli.Context) interface{} {
	switch name {
	case
		apiToken,
		appOwnerName,
		appId:
		return c.String(name.String())
	case invitees:
		return c.StringSlice(name.String())
	case
		android,
		ios,
		developerRole:
		return c.Bool(name.String())
	}

	panic("Option name mapping is not found")
}
