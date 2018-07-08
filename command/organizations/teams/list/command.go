package organizations_teams_list

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/urfave/cli"
	"strings"
	"github.com/jmatsu/dpg/api/response"
	"encoding/json"
	"github.com/jmatsu/dpg/api/request/organizations/teams/list"
)

func Command() cli.Command {
	return cli.Command{
		Name:    "list-teams",
		Aliases: []string{"i"},
		Usage:   "Show teams which belong to the specified organization",
		Action:  action,
		Flags:   allFlags(),
	}
}

func action(c *cli.Context) error {
	endpoint, authority, requestParams, err := buildResource(c)

	if err != nil {
		return err
	}

	_, err = listTeams(
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

func buildResource(c *cli.Context) (*api.OrganizationTeamsEndpoint, *api.Authority, *list.Request, error) {
	authority := api.Authority{
		Token: apiToken.Value(c).(string),
	}

	endpoint := api.OrganizationTeamsEndpoint{
		BaseURL:          "https://deploygate.com",
		OrganizationName: organizationName.Value(c).(string),
		AppId:            appId.Value(c).(string),
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

func verifyInput(e api.OrganizationTeamsEndpoint, authority api.Authority, _ list.Request) error {
	if authority.Token == "" {
		return errors.New("api token must be specified")
	}

	if e.OrganizationName == "" {
		return errors.New("organization name must be specified")
	}

	if e.AppId == "" {
		return errors.New("application id must be specified")
	}

	if !strings.EqualFold(e.AppPlatform, "android") && !strings.EqualFold(e.AppPlatform, "ios") {
		return errors.New("A platform must be either of `android` or `ios`")
	}

	return nil
}

func listTeams(e api.OrganizationTeamsEndpoint, authority api.Authority, requestParam list.Request, verbose bool) (response.OrganizationTeamsListResponse, error) {
	var r response.OrganizationTeamsListResponse

	if err := verifyInput(e, authority, requestParam); err != nil {
		return r, err
	}

	if bytes, err := e.GetQueryRequest(authority, requestParam, verbose); err != nil {
		return r, err
	} else if err := json.Unmarshal(bytes, &r); err != nil {
		return r, err
	} else {
		return r, nil
	}
}
