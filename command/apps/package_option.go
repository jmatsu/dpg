package apps

import (
	"github.com/jmatsu/dpg/command"
	"github.com/urfave/cli"
	"gopkg.in/guregu/null.v3"
)

type packageOption int

const (
	appFilePath        packageOption = iota
	isPublic
	enableNotification
	shortMessage
	distributionKey
	distributionName
	releaseNote
)

func uploadFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		AppOwnerName.Flag(),
		Android.Flag(),
		IOS.Flag(),
		appFilePath.flag(),
		isPublic.flag(),
		enableNotification.flag(),
		shortMessage.flag(),
		distributionKey.flag(),
		distributionName.flag(),
		releaseNote.flag(),
	}
}

func (o packageOption) name() string {
	switch o {
	case appFilePath:
		return "app"
	case isPublic:
		return "public"
	case enableNotification:
		return "enableNotification"
	case shortMessage:
		return "message"
	case distributionKey:
		return "distributionKey"
	case distributionName:
		return "distributionName"
	case releaseNote:
		return "releaseNote"
	}

	panic("Option name mapping is not found")
}

func (o packageOption) flag() cli.Flag {
	switch o {
	case appFilePath:
		return cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] A path of an application file to be uploaded",
		}
	case isPublic:
		return cli.BoolFlag{
			Name:  o.name(),
			Usage: "Specify true if an application to be uploaded should be public",
		}
	case enableNotification:
		return cli.BoolFlag{
			Name:  o.name(),
			Usage: "[iOS only] Specify true if iOS's notifications should be enabled",
		}
	case shortMessage:
		return cli.StringFlag{
			Name:  o.name(),
			Usage: "A short message to explain this update",
		}
	case distributionKey:
		return cli.StringFlag{
			Name:  o.name(),
			Usage: "A key of a distribution which an application will be uploaded to",
		}
	case distributionName:
		return cli.StringFlag{
			Name:  o.name(),
			Usage: "A name of a distribution which an application will be uploaded to",
		}
	case releaseNote:
		return cli.StringFlag{
			Name:  o.name(),
			Usage: "A release note for this revision",
		}
	}

	panic("Option name mapping is not found")
}

func getAppFilePath(c *cli.Context) string {
	return c.String(appFilePath.name())
}

func isPublc(c *cli.Context) bool {
	return c.Bool(isPublic.name())
}

func isEnabledNotification(c *cli.Context) bool {
	return c.Bool(enableNotification.name())
}

func getShortMessage(c *cli.Context) null.String {
	if x := c.String(shortMessage.name()); c.IsSet(shortMessage.name()) {
		return null.StringFrom(x)
	} else {
		return null.StringFromPtr(nil)
	}
}

func getDistributionKey(c *cli.Context) null.String {
	if x := c.String(distributionKey.name()); c.IsSet(distributionKey.name()) {
		return null.StringFrom(x)
	} else {
		return null.StringFromPtr(nil)
	}
}

func getDistributionName(c *cli.Context) null.String {
	if x := c.String(distributionName.name()); c.IsSet(distributionName.name()) {
		return null.StringFrom(x)
	} else {
		return null.StringFromPtr(nil)
	}
}

func getReleaseNote(c *cli.Context) null.String {
	if x := c.String(releaseNote.name()); c.IsSet(releaseNote.name()) {
		return null.StringFrom(x)
	} else {
		return null.StringFromPtr(nil)
	}
}