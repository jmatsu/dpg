package invite

import (
	"github.com/urfave/cli"
	"gopkg.in/guregu/null.v3"
)

type optionName string

const (
	apiToken     optionName = "token"
	appOwnerName optionName = "app-owner"
	appId        optionName = "app-id"
	appPlatform  optionName = "app-platform"
	invitees     optionName = "invitees"
	role         optionName = "role"
)

func allFlags() []cli.Flag {
	return []cli.Flag{
		apiToken.Flag(),
		appOwnerName.Flag(),
		appId.Flag(),
		appPlatform.Flag(),
		invitees.Flag(),
		role.Flag(),
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
	case invitees:
		return cli.StringSliceFlag{
			Name:  name.String(),
			Usage: "[Required] Comma separated names or e-mails of those who you want to invite",
		}
	case role:
		return cli.BoolFlag{
			Name:   name.String(),
			Usage:  "[Old Free/Lite/Pro/Biz plans only] Specify true if invitee(s) should be the developer role, otherwise they would be the tester role. false by default",
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
		appId,
		appPlatform:
		return c.String(name.String())
	case invitees:
		return c.StringSlice(name.String())
	case
		role:
		if x := c.Bool(name.String()); c.IsSet(name.String()) {
			return null.BoolFrom(x)
		} else {
			return null.BoolFromPtr(nil)
		}
	}

	panic("Option name mapping is not found")
}
