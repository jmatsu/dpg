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

	//

	featureBranch

	prefix
)

func exposeFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		apps.AppOwnerName.Flag(),
		apps.Android.Flag(),
		apps.IOS.Flag(),
		appFilePath.flag(),
		enableNotification.flag(),
		distributionName.flag(),
		featureBranch.flag(),
		prefix.flag(),
	}
}

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
	case featureBranch:
		return constant.IsFeatureBranch
	case prefix:
		return constant.Prefix
	}

	panic("Option name mapping is not found")
}

func (o packageOption) flag() cli.Flag {
	switch o {
	case appFilePath:
		return &cli.PathFlag{
			Name:    o.name(),
			EnvVars: []string{constant.AppFilePathEnv},
			Usage:   "[Required] The file path of the application to be uploaded",
		}
	case enableNotification:
		return &cli.BoolFlag{
			Name:    o.name(),
			EnvVars: []string{constant.EnableNotificationEnv},
			Usage:   "[iOS only] Specify true if iOS's notifications should be enabled",
		}
	case shortMessage:
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "A short message to explain this update",
		}
	case distributionName:
		return &cli.StringFlag{
			Name:    o.name(),
			EnvVars: []string{constant.DistributionNameEnv},
			Usage:   "A name of a distribution to be created or updated",
		}
	case releaseNote:
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "A release note for this revision",
		}
	case featureBranch:
		return &cli.BoolFlag{
			Name:    o.name(),
			Usage:   "expose variables for feature branch if specified",
		}
	case prefix:
		return &cli.StringFlag{
			Name:  o.name(),
			Usage: "A prefix of each lines to be exported",
		}
	}

	panic("Option name mapping is not found")
}
