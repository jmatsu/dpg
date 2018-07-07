package upload

import (
	"github.com/urfave/cli"
	"gopkg.in/guregu/null.v3"
)

type OptionName string

const (
	AppFilePath        OptionName = "app"
	ApiToken           OptionName = "token"
	AppOwnerName       OptionName = "app-owner"
	IsPublic           OptionName = "visible"
	EnableNotification OptionName = "enableNotification"
	ShortMessage       OptionName = "message"
	DistributionKey    OptionName = "distributionKey"
	DistributionName   OptionName = "distributionName"
	ReleaseNote        OptionName = "releaseNote"
)

func (name OptionName) String() string {
	return string(name)
}

func (name OptionName) Flag() cli.Flag {
	switch name {
	case AppFilePath:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "[Required] A path of an application file to be uploaded",
		}
	case ApiToken:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "[Required] API token",
		}
	case AppOwnerName:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "[Required] An owner of applications",
		}
	case IsPublic:
		return cli.BoolFlag{
			Name:  name.String(),
			Usage: "Specify true if an application to be uploaded should be public",
		}
	case EnableNotification:
		return cli.BoolFlag{
			Name:  name.String(),
			Usage: "[iOS only] Specify true if iOS's notifications should be enabled",
		}
	case ShortMessage:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "A short message to explain this update",
		}
	case DistributionKey:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "A key of a distribution which an application will be uploaded to",
		}
	case DistributionName:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "A name of a distribution which an application will be uploaded to",
		}
	case ReleaseNote:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "A release note for this revision",
		}
	}

	panic("Option name mapping is not found")
}

func (name OptionName) Value(c *cli.Context) interface{} {
	switch name {
	case
		AppFilePath,
		ApiToken,
		AppOwnerName:
		return c.String(name.String())
	case IsPublic:
		return c.Bool(name.String())
	case EnableNotification:
		return c.Bool(string(name))
	case
		ShortMessage,
		DistributionKey,
		DistributionName,
		ReleaseNote:
		if x := c.String(name.String()); c.IsSet(name.String()) {
			return null.StringFrom(x)
		} else {
			return null.StringFromPtr(nil)
		}
	}

	panic("Option name mapping is not found")
}
