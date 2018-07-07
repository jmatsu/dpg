package upload

import (
	"errors"
	"fmt"
	"github.com/jmatsu/dpg/api"
	requestAppUpload "github.com/jmatsu/dpg/api/request/app/upload"
	"github.com/urfave/cli"
	"gopkg.in/guregu/null.v3"
	"strings"
	"github.com/jmatsu/dpg/api/response"
	"encoding/json"
	"os"
)

func Command() cli.Command {
	return cli.Command{
		Name:    "upload",
		Aliases: []string{"u"},
		Usage:   "Upload either android application or iOS application to the specified owner space",
		Action:  action,
		Flags: []cli.Flag{
			AppFilePath.Flag(),
			ApiToken.Flag(),
			AppOwnerName.Flag(),
			IsPublic.Flag(),
			EnableNotification.Flag(),
			ShortMessage.Flag(),
			DistributionKey.Flag(),
			DistributionName.Flag(),
			ReleaseNote.Flag(),
		},
		Subcommands: []cli.Command{
			{
				Name:    "android",
				Aliases: []string{"a"},
				Usage:   "Upload an android application to the specified owner space",
				Action:  androidAppAction,
				Flags: []cli.Flag{
					AppFilePath.Flag(),
					ApiToken.Flag(),
					AppOwnerName.Flag(),
				},
			},
			{
				Name:    "ios",
				Aliases: []string{"i"},
				Usage:   "Upload an iOS application to the specified owner space",
				Action:  iOSAppAction,
				Flags: []cli.Flag{
					AppFilePath.Flag(),
					ApiToken.Flag(),
					AppOwnerName.Flag(),
					EnableNotification.Flag(),
				},
			},
		},
	}
}

func androidAppAction(c *cli.Context) error {
	if err := verifyAndroidApp(c); err != nil {
		return err
	}

	return action(c)
}

func iOSAppAction(c *cli.Context) error {
	if err := verifyIOSApp(c); err != nil {
		return err
	}

	return action(c)
}

func action(c *cli.Context) error {
	authority := api.Authority{
		Token: ApiToken.Value(c).(string),
	}

	endpoint := api.UploadAppEndpoint{
		BaseURL:      "https://deploygate.com",
		AppOwnerName: AppOwnerName.Value(c).(string),
	}

	resp, err := uploadApp(
		endpoint,
		authority,
		requestAppUpload.Request{
			AppFilePath:        AppFilePath.Value(c).(string),
			AppVisible:         IsPublic.Value(c).(bool),
			EnableNotification: EnableNotification.Value(c).(bool),
			ShortMessage:       ShortMessage.Value(c).(null.String),
			DistributionKey:    DistributionKey.Value(c).(null.String),
			DistributionName:   DistributionName.Value(c).(null.String),
			ReleaseNote:        ReleaseNote.Value(c).(null.String),
		},
		c.GlobalBoolT("verbose"),
	)

	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}

func uploadApp(e api.UploadAppEndpoint, authority api.Authority, requestBody requestAppUpload.Request, verbose bool) (response.UploadAppResponse, error) {
	var r response.UploadAppResponse

	if err := verifyInput(e, authority, requestBody); err != nil {
		return r, err
	}

	if bytes, err := e.MultiPartFormRequest(authority, requestBody, verbose); err != nil {
		return r, err
	} else if err := json.Unmarshal(bytes, &r); err != nil {
		return r, err
	} else {
		return r, nil
	}
}

func verifyAndroidApp(c *cli.Context) error {
	appFilePath := AppFilePath.Value(c).(string)

	if !strings.HasSuffix(appFilePath, ".apk") {
		return errors.New("A file path must be an apk file")
	}

	return nil
}

func verifyIOSApp(c *cli.Context) error {
	appFilePath := AppFilePath.Value(c).(string)

	if !strings.HasSuffix(appFilePath, ".ipa") {
		return errors.New("A file path must be an ipa file")
	}

	return nil
}

func verifyInput(e api.UploadAppEndpoint, authority api.Authority, requestBody requestAppUpload.Request) error {
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
