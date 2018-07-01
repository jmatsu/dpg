package upload

import "github.com/urfave/cli"

type OptionName string

const (
	AppFilePath          OptionName = "app"
	ApiToken             OptionName = "token"
	AppOwnerName         OptionName = "app-owner"
	IsPrivate            OptionName = "visible"
	SuppressNotification OptionName = "suppressNotification"
	ShortMessage         OptionName = "message"
	DistributionKey      OptionName = "distributionKey"
	DistributionName     OptionName = "distributionName"
	ReleaseNote          OptionName = "releaseNote"
)

func (name OptionName) String() string {
	return string(name)
}

func (name OptionName) Flag() cli.Flag {
	switch name {
	case AppFilePath:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "A path of an application file to be uploaded",
		}
	case ApiToken:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "API token",
		}
	case AppOwnerName:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "An owner of applications",
		}
	case IsPrivate:
		return cli.BoolTFlag{
			Name:  name.String(),
			Usage: "Specify false if an application to be uploaded should be public",
		}
	case SuppressNotification:
		return cli.BoolTFlag{
			Name:  name.String(),
			Usage: "Specify false if iOS's notifications should be enabled",
		}
	case ShortMessage:
		return cli.StringFlag{
			Name:  name.String(),
			Usage: "A short message",
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
