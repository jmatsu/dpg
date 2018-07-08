package apps_members_list

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/urfave/cli"
	"strings"
	"github.com/jmatsu/dpg/api/response"
	"encoding/json"
	"github.com/jmatsu/dpg/api/request/apps/members/list"
)

func Command() cli.Command {
	return cli.Command{
		Name:    "list-users",
		Aliases: []string{"i"},
		Usage:   "Show users who have joined to the specified application (expect the apps owner)",
		Action:  action,
		Flags:   allFlags(),
	}
}

func action(c *cli.Context) error {
	endpoint, authority, requestParams, err := buildResource(c)

	if err != nil {
		return err
	}

	_, err = listUsers(
		*endpoint,
		*authority,
		*requestParams,
		c.GlobalBoolT("verbose"),
	)

	if err != nil {
		return err
	}

	return nil
}

func buildResource(c *cli.Context) (*api.AppMemberEndpoint, *api.Authority, *list.Request, error) {
	authority := api.Authority{
		Token: apiToken.Value(c).(string),
	}

	endpoint := api.AppMemberEndpoint{
		BaseURL:      "https://deploygate.com",
		AppOwnerName: appOwnerName.Value(c).(string),
		AppId:        appId.Value(c).(string),
	}

	requestParams := list.Request{}

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

	if err := verifyInput(endpoint, authority, requestParams); err != nil {
		return nil, nil, nil, err
	}

	return &endpoint, &authority, &requestParams, nil
}

func verifyInput(e api.AppMemberEndpoint, authority api.Authority, _ list.Request) error {
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

	return nil
}

func listUsers(e api.AppMemberEndpoint, authority api.Authority, requestParam list.Request, verbose bool) (response.AppUsersResponse, error) {
	var r response.AppUsersResponse

	if bytes, err := e.GetQueryRequest(authority, requestParam, verbose); err != nil {
		return r, err
	} else if err := json.Unmarshal(bytes, &r); err != nil {
		return r, err
	} else {
		return r, nil
	}
}
