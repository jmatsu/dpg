package apps

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/constant"
	"github.com/jmatsu/dpg/request/apps"
	"github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v3"
	"gopkg.in/urfave/cli.v2"
	"strings"
)

func UploadCommand() *cli.Command {
	return &cli.Command{
		Name:   "upload",
		Usage:  "Upload either android application or iOS application to the specified owner space",
		Action: command.AuthorizedCommandAction(NewUploadCommand),
		Flags:  UploadFlags(),
	}
}

type uploadCommand struct {
	appOwnerName string
	requestBody  apps.UploadRequest
}

func NewUploadCommand(c *cli.Context) (command.Command, error) {
	appOwnerName, err := command.RequireAppOwnerName(c)

	if err != nil {
		return nil, err
	}

	appFilePath, err := command.RequireAppFilePath(c)

	if err != nil {
		return nil, err
	}

	if platform, err := command.GetAppPlatform(c); err != nil {
		return nil, err
	} else if platform.Valid {
		if platform.String == constant.Android {
			if !strings.HasSuffix(appFilePath, ".apk") && !strings.HasSuffix(appFilePath, ".aab") {
				return nil, errors.New("an application file must be an apk file or an aab file")
			}
		} else if platform.String == constant.IOS {
			if !strings.HasSuffix(appFilePath, ".ipa") {
				return nil, errors.New("an application file must be an ipa file")
			}
		}
	}

	distributionKey, err := command.GetDistributionKey(c)

	if err != nil {
		return nil, err
	}

	distributionName, err := command.GetDistributionName(c)

	if err != nil {
		return nil, err
	}

	shortMessage, err := command.GetShortMessage(c)

	if err != nil {
		return nil, err
	}

	releaseNote, err := command.GetReleaseNote(c)

	if err != nil {
		return nil, err
	}

	if distributionKey.Valid && distributionName.Valid {
		distributionName = null.StringFromPtr(nil)
		logrus.Warnf("distribution name was ignored because distribution key was also specified\n")
	}

	cmd := uploadCommand{
		appOwnerName: appOwnerName,
		requestBody: apps.UploadRequest{
			AppFilePath:        appFilePath,
			AppVisible:         command.IsPublic(c),
			EnableNotification: command.IsEnabledNotification(c),
			ShortMessage:       shortMessage,
			DistributionKey:    distributionKey,
			DistributionName:   distributionName,
			ReleaseNote:        releaseNote,
		},
	}

	return cmd, nil
}

func (cmd uploadCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).UploadApp(cmd.appOwnerName, cmd.requestBody)
}
