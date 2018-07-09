package apps

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/apps/upload"
	"github.com/jmatsu/dpg/command"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func UploadCommand() cli.Command {
	return cli.Command{
		Name:   "upload",
		Usage:  "Upload either android application or iOS application to the specified owner space",
		Action: command.CommandAction(newUploadCommand),
		Flags:  uploadFlags(),
	}
}

type uploadCommand struct {
	endpoint    *api.AppsEndpoint
	authority   *api.Authority
	requestBody *upload.Request
}

func newUploadCommand(c *cli.Context) (command.Command, error) {
	appFilePath := getAppFilePath(c)

	if platform, err := GetAppPlatform(c); err != nil {
		return nil, err
	} else if platform == "android" {
		if !strings.HasSuffix(appFilePath, ".apk") {
			return nil, errors.New("A file path must be an apk file")
		}
	} else {
		if !strings.HasSuffix(appFilePath, ".ipa") {
			return nil, errors.New("A file path must be an ipa file")
		}
	}

	cmd := uploadCommand{
		authority: &api.Authority{
			Token: command.GetApiToken(c),
		},
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

func (cmd uploadCommand) VerifyInput() error {
	if cmd.authority.Token == "" {
		return errors.New("api token must be specified")
	}

	if cmd.endpoint.AppOwnerName == "" {
		return errors.New("application owner must be specified")
	}

	if cmd.requestBody.AppFilePath == "" {
		return errors.New("application file path must be specified")
	}

	if _, err := os.Stat(cmd.requestBody.AppFilePath); os.IsNotExist(err) {
		return err
	}

	if cmd.requestBody.DistributionKey.Valid && cmd.requestBody.DistributionKey.String == "" {
		return errors.New("empty distribution key is not allowed")
	}

	if cmd.requestBody.DistributionName.Valid && cmd.requestBody.DistributionName.String == "" {
		return errors.New("empty distribution name is not allowed")
	}

	return nil
}

func (cmd uploadCommand) Run() (string, error) {
	if bytes, err := cmd.endpoint.MultiPartFormRequest(*cmd.authority, *cmd.requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
