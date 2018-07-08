package apps_upload

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/apps/upload"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func Command() cli.Command {
	return cli.Command{
		Name:   "upload",
		Usage:  "Upload either android application or iOS application to the specified owner space",
		Action: action,
		Flags:  flags(),
	}
}

func action(c *cli.Context) error {
	endpoint, authority, requestBody, err := buildResource(c)

	_, err = uploadApp(
		*endpoint,
		*authority,
		*requestBody,
	)

	if err != nil {
		return err
	}

	return nil
}

func buildResource(c *cli.Context) (*api.AppUploadEndpoint, *api.Authority, *upload.Request, error) {
	authority := api.Authority{
		Token: command.GetApiToken(c),
	}

	endpoint := api.AppUploadEndpoint{
		BaseURL:      api.EndpointURL,
		AppOwnerName: apps.GetAppOwnerName(c),
	}

	requestBody := upload.Request{
		AppFilePath:        getAppFilePath(c),
		AppVisible:         isPublc(c),
		EnableNotification: isEnabledNotification(c),
		ShortMessage:       getShortMessage(c),
		DistributionKey:    getDistributionKey(c),
		DistributionName:   getDistributionName(c),
		ReleaseNote:        getReleaseNote(c),
	}

	if platform, err := apps.GetAppPlatform(c); err != nil {
		return nil, nil, nil, err
	} else if platform == "android" {
		if !strings.HasSuffix(requestBody.AppFilePath, ".apk") {
			return nil, nil, nil, errors.New("A file path must be an apk file")
		}
	} else {
		if !strings.HasSuffix(requestBody.AppFilePath, ".ipa") {
			return nil, nil, nil, errors.New("A file path must be an ipa file")
		}
	}

	if err := verifyInput(endpoint, authority, requestBody); err != nil {
		return nil, nil, nil, err
	}

	return &endpoint, &authority, &requestBody, nil
}

func verifyInput(e api.AppUploadEndpoint, authority api.Authority, requestBody upload.Request) error {
	if authority.Token == "" {
		return errors.New("api token must be specified")
	}

	if e.AppOwnerName == "" {
		return errors.New("application owner must be specified")
	}

	if requestBody.AppFilePath == "" {
		return errors.New("application file path must be specified")
	}

	if _, err := os.Stat(requestBody.AppFilePath); os.IsNotExist(err) {
		return err
	}

	if requestBody.DistributionKey.Valid && requestBody.DistributionKey.String == "" {
		return errors.New("empty distribution key is not allowed")
	}

	if requestBody.DistributionName.Valid && requestBody.DistributionName.String == "" {
		return errors.New("empty distribution name is not allowed")
	}

	return nil
}

func uploadApp(e api.AppUploadEndpoint, authority api.Authority, requestBody upload.Request) (string, error) {
	if bytes, err := e.MultiPartFormRequest(authority, requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
