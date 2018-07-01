package upload

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/response"
	"github.com/urfave/cli"
	"gopkg.in/guregu/null.v3"
	"net/http"
	"strings"
)

func AndroidApp(c *cli.Context) error {
	appFilePath := c.String(AppFilePath.String())

	if !strings.HasSuffix(appFilePath, ".apk") {
		return errors.New("A file path must be an apk file")
	}

	return App(c)
}

func IOSApp(c *cli.Context) error {
	appFilePath := c.String(AppFilePath.String())

	if !strings.HasSuffix(appFilePath, ".ipa") {
		return errors.New("A file path must be an ipa file")
	}

	return App(c)
}

func App(c *cli.Context) error {
	authority := api.Authority{
		Token: c.String(ApiToken.String()),
	}

	var shortMessage *string

	if x := c.String(ShortMessage.String()); c.IsSet(ShortMessage.String()) {
		shortMessage = &x
	}

	var distKey *string

	if x := c.String(DistributionKey.String()); c.IsSet(DistributionKey.String()) {
		distKey = &x
	}

	var distName *string

	if x := c.String(DistributionName.String()); c.IsSet(DistributionName.String()) {
		distName = &x
	}

	var releaseNote *string

	if x := c.String(ReleaseNote.String()); c.IsSet(ReleaseNote.String()) {
		releaseNote = &x
	}

	resp, err := uploadApp(
		authority,
		c.String(AppOwnerName.String()),
		Request{
			AppFilePath:          c.String(AppFilePath.String()),
			AppVisibility:        BoolToVisibility(c.BoolT(IsPrivate.String())),
			SuppressNotification: c.BoolT(string(SuppressNotification)),
			ShortMessage:         null.StringFromPtr(shortMessage),
			DistributionKey:      null.StringFromPtr(distKey),
			DistributionName:     null.StringFromPtr(distName),
			ReleaseNote:          null.StringFromPtr(releaseNote),
		},
		c.GlobalBoolT("verbose"),
	)

	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}

func uploadApp(authority api.Authority, appOwnerName string, requestBody Request, verbose bool) (response.UploadAppResponse, error) {
	var r response.UploadAppResponse

	data, contentType, err := requestBody.MultiPartForm(authority)

	if err != nil {
		return r, err
	}

	url := fmt.Sprintf("https://deploygate.com/api/users/%s/apps", appOwnerName)

	req, _ := http.NewRequest(http.MethodPost, url, &data)
	req.Header.Set("Content-Type", contentType)

	resp, err := new(http.Client).Do(req)

	if err != nil {
		return r, err
	}

	defer resp.Body.Close()

	bytes, errResp, err := api.Response(*resp, verbose)

	if err != nil {
		return r, err
	}

	if errResp != nil {
		return r, errors.New("api returned an error response")
	}

	r = response.UploadAppResponse{}

	if err := json.Unmarshal(bytes, &r); err != nil {
		return r, err
	}

	return r, nil
}
