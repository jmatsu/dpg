package apps

import (
	"github.com/jmatsu/dpg/command"
	"gopkg.in/urfave/cli.v2"
)

func UploadFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		command.AppOwnerName.Flag(),
		command.Android.Flag(),
		command.IOS.Flag(),
		command.AppFilePath.Flag(),
		command.Public.Flag(),
		command.EnableNotification.Flag(),
		command.ShortMessage.Flag(),
		command.DistributionKey.Flag(),
		command.DistributionName.Flag(),
		command.ReleaseNote.Flag(),
	}
}
