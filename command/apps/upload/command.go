package apps_upload

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/apps/upload"
	"github.com/urfave/cli"
	"gopkg.in/guregu/null.v3"
	"github.com/jmatsu/dpg/api/response"
	"encoding/json"
	"os"
	"strings"
)

func Command() cli.Command {
	return cli.Command{
		Name:   "upload",
		Usage:  "Upload either android application or iOS application to the specified owner space",
		Action: action,
		Flags:  allFlags(),
	}
}

func action(c *cli.Context) error {
	endpoint, authority, requestBody, err := buildResource(c)

	_, err = uploadApp(
		*endpoint,
		*authority,
		*requestBody,
		c.GlobalBoolT("verbose"),
	)

	if err != nil {
		return err
	}

	return nil
}

func buildResource(c *cli.Context) (*api.AppUploadEndpoint, *api.Authority, *upload.Request, error) {
	authority := api.Authority{
		Token: apiToken.Value(c).(string),
	}

	endpoint := api.AppUploadEndpoint{
		BaseURL:      "https://deploygate.com",
		AppOwnerName: appOwnerName.Value(c).(string),
	}

	requestBody := upload.Request{
		AppFilePath:        appFilePath.Value(c).(string),
		AppVisible:         isPublic.Value(c).(bool),
		EnableNotification: enableNotification.Value(c).(bool),
		ShortMessage:       shortMessage.Value(c).(null.String),
		DistributionKey:    distributionKey.Value(c).(null.String),
		DistributionName:   distributionName.Value(c).(null.String),
		ReleaseNote:        releaseNote.Value(c).(null.String),
	}

	isAndroid := android.Value(c).(bool)
	isIOS := ios.Value(c).(bool)

	if isAndroid && isIOS {
		return nil, nil, nil, errors.New("only one option of android or ios is allowed")
	}

	if !isAndroid && !isIOS {
		return nil, nil, nil, errors.New("either of android or ios must be specified")
	}

	if isAndroid {
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

func uploadApp(e api.AppUploadEndpoint, authority api.Authority, requestBody upload.Request, verbose bool) (response.AppUploadResponse, error) {
	var r response.AppUploadResponse

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
