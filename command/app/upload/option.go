package upload

import (
	"github.com/urfave/cli"
	"gopkg.in/guregu/null.v3"
)

type optionName string

const (
	apiToken           optionName = "token"
	appOwnerName       optionName = "app-owner"
	appFilePath        optionName = "app"
	isPublic           optionName = "visible"
	enableNotification optionName = "enableNotification"
	shortMessage       optionName = "message"
	distributionKey    optionName = "distributionKey"
	distributionName   optionName = "distributionName"
	releaseNote        optionName = "releaseNote"
)

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
	case appFilePath:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "[Required] A path of an application file to be uploaded",
		}
	case isPublic:
		return cli.BoolFlag{
			Name:  name.String(),
			Usage: "Specify true if an application to be uploaded should be public",
		}
	case enableNotification:
		return cli.BoolFlag{
			Name:  name.String(),
			Usage: "[iOS only] Specify true if iOS's notifications should be enabled",
		}
	case shortMessage:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "A short message to explain this update",
		}
	case distributionKey:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "A key of a distribution which an application will be uploaded to",
		}
	case distributionName:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "A name of a distribution which an application will be uploaded to",
		}
	case releaseNote:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "A release note for this revision",
		}
	}

	panic("Option name mapping is not found")
}

func (name optionName) Value(c *cli.Context) interface{} {
	switch name {
	case
		apiToken,
		appOwnerName,
		appFilePath:
		return c.String(name.String())
	case isPublic:
		return c.Bool(name.String())
	case enableNotification:
		return c.Bool(string(name))
	case
		shortMessage,
		distributionKey,
		distributionName,
		releaseNote:
		if x := c.String(name.String()); c.IsSet(name.String()) {
			return null.StringFrom(x)
		} else {
			return null.StringFromPtr(nil)
		}
	}

	panic("Option name mapping is not found")
}
