package apps_members_remove

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/urfave/cli"
	"strings"
	"github.com/jmatsu/dpg/api/response"
	"encoding/json"
	"github.com/jmatsu/dpg/api/request/apps/members/remove"
)

func Command() cli.Command {
	return cli.Command{
		Name:   "remove",
		Usage:  "Remove users from the specified application space",
		Action: action,
		Flags:  allFlags(),
	}
}

func action(c *cli.Context) error {
	endpoint, authority, requestBody, err := buildResource(c)

	if err != nil {
		return err
	}

	_, err = removeUsers(
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

func buildResource(c *cli.Context) (*api.AppMemberEndpoint, *api.Authority, *remove.Request, error) {
	authority := api.Authority{
		Token: apiToken.Value(c).(string),
	}

	endpoint := api.AppMemberEndpoint{
		BaseURL:      "https://deploygate.com",
		AppOwnerName: appOwnerName.Value(c).(string),
		AppId:        appId.Value(c).(string),
	}

	requestBody := remove.Request{
		UserNamesOrEmails: removees.Value(c).([]string),
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
		endpoint.AppPlatform = "android"
	} else {
		endpoint.AppPlatform = "ios"
	}

	if err := verifyInput(endpoint, authority, requestBody); err != nil {
		return nil, nil, nil, err
	}

	return &endpoint, &authority, &requestBody, nil
}

func removeUsers(e api.AppMemberEndpoint, authority api.Authority, requestBody remove.Request, verbose bool) (response.AppInviteResponse, error) {
	var r response.AppInviteResponse

	if err := verifyInput(e, authority, requestBody); err != nil {
		return r, err
	}

	if bytes, err := e.DeleteRequest(authority, requestBody, verbose); err != nil {
		return r, err
	} else if err := json.Unmarshal(bytes, &r); err != nil {
		return r, err
	} else {
		return r, nil
	}
}

func verifyInput(e api.AppMemberEndpoint, authority api.Authority, requestBody remove.Request) error {
	if authority.Token == "" {
		return errors.New("api token must be specified")
	}

	if e.AppOwnerName == "" {
		return errors.New("application owner must be specified")
	}

	if e.AppId == "" {
		return errors.New("application id must be specified")
	}

	if !strings.EqualFold(e.AppPlatform, "android") && !strings.EqualFold(e.AppPlatform, "ios") {
		return errors.New("A platform must be either of `android` or `ios`")
	}

	if len(requestBody.UserNamesOrEmails) == 0 {
		return errors.New("the number of removees must be greater than 0")
	}

	return nil
}
