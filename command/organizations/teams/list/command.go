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
	authority := api.Authority{
		Token: apiToken.Value(c).(string),
	}

	endpoint := api.OrganizationTeamsEndpoint{
		BaseURL:          "https://deploygate.com",
		OrganizationName: organizationName.Value(c).(string),
		AppId:            appId.Value(c).(string),
		AppPlatform:      appPlatform.Value(c).(string),
	}

	_, err := listTeams(
		endpoint,
		authority,
		list.Request{
		},
		c.GlobalBoolT("verbose"),
	)

	if err != nil {
		return err
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
