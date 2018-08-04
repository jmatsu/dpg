package apps

import (
	"errors"
	"fmt"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/apps/upload"
	"github.com/jmatsu/dpg/command"
	"github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
	"os"
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
	endpoint    *api.AppsEndpoint
	requestBody *upload.Request
}

func NewUploadCommand(c *cli.Context) (command.Command, error) {
	appFilePath := getAppFilePath(c)

	if platform, err := GetAppPlatform(c); err != nil {
		return nil, err
	} else if platform == "android" {
		if !strings.HasSuffix(appFilePath, ".apk") {
			return nil, errors.New("an application file must be an apk file")
		}
	} else {
		if !strings.HasSuffix(appFilePath, ".ipa") {
			return nil, errors.New("an application file must be an ipa file")
		}
	}

	cmd := uploadCommand{
		endpoint: &api.AppsEndpoint{
			BaseURL:      api.EndpointURL,
			AppOwnerName: GetAppOwnerName(c),
		},
		requestBody: &upload.Request{
			AppFilePath:        appFilePath,
			AppVisible:         isPublc(c),
			EnableNotification: isEnabledNotification(c),
			ShortMessage:       getShortMessage(c),
			DistributionKey:    getDistributionKey(c),
			DistributionName:   getDistributionName(c),
			ReleaseNote:        getReleaseNote(c),
		},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

/*
Endpoint:
	app owner's name is required
Parameters:
	the specified file must exist
	the specified distribution key must not be empty
	the specified distribution name is not used if a distribution key is also specified
*/
func (cmd uploadCommand) VerifyInput() error {
	if err := RequireAppOwnerName(cmd.endpoint.AppOwnerName); err != nil {
		return err
	}

	if cmd.requestBody.AppFilePath == "" {
		return errors.New(fmt.Sprintf("--%s must not be empty", appFilePath.name()))
	}

	if f, err := os.Stat(cmd.requestBody.AppFilePath); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("%s is not found", cmd.requestBody.AppFilePath))
	} else if err != nil {
		return errors.New(fmt.Sprintf("seriously wrong when trying to open %s", cmd.requestBody.AppFilePath))
	} else if f != nil && f.Size() == 0 {
		return errors.New("an application file must not be an empty file")
	}

	if cmd.requestBody.DistributionKey.Valid && cmd.requestBody.DistributionKey.String == "" {
		return errors.New(fmt.Sprintf("--%s must not be empty if specified", distributionKey.name()))
	}

	if cmd.requestBody.DistributionKey.Valid && cmd.requestBody.DistributionName.String != "" {
		logrus.Warnf("--%s was specified so --%s wouldn't be used", distributionKey.name(), distributionName.name())
	}

	return nil
}

func (cmd uploadCommand) Run(authorization *api.Authorization) (string, error) {
	if bytes, err := cmd.endpoint.MultiPartFormRequest(*authorization, *cmd.requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
