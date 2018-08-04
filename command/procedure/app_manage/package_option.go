package app_manage

import (
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/jmatsu/dpg/command/constant"
	"gopkg.in/urfave/cli.v2"
)

type packageOption int

const (
	appFilePath packageOption = iota
	enableNotification
	shortMessage
	distributionName
	releaseNote
)

func onFeatureBranchFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		apps.AppOwnerName.Flag(),
		apps.Android.Flag(),
		apps.IOS.Flag(),
		appFilePath.flag(),
		enableNotification.flag(),
		shortMessage.flag(),
		distributionName.flag(),
		releaseNote.flag(),
	}
}

func onDeployBranchFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		apps.AppOwnerName.Flag(),
		apps.AppId.Flag(),
		apps.Android.Flag(),
		apps.IOS.Flag(),
		distributionName.flag(),
	}
}

func (o packageOption) name() string {
	switch o {
	case appFilePath:
		return constant.AppFilePath
	case enableNotification:
		return constant.EnableNotification
	case distributionName:
		return constant.DistributionName
	case shortMessage:
		return constant.ShortMessage
	case releaseNote:
		return constant.ReleaseNote
	}

	panic("Option name mapping is not found")
}

func (o packageOption) flag() cli.Flag {
	switch o {
	case appFilePath:
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "[Required] The file path of the application to be uploaded",
		}
	case enableNotification:
		return &cli.BoolFlag{
			Name:  o.name(),
			Usage: "[iOS only] Specify true if iOS's notifications should be enabled",
		}
	case shortMessage:
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "A short message to explain this update",
		}
	case distributionName:
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "A name of a distribution to be created or updated",
		}
	case releaseNote:
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "A release note for this revision",
		}
	}

	panic("Option name mapping is not found")
}
