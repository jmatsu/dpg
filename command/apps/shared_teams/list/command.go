package apps_shared_teams_list

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/urfave/cli"
	"strings"
	"github.com/jmatsu/dpg/api/response"
	"encoding/json"
	"github.com/jmatsu/dpg/api/request/apps/shared_teams/list"
	"github.com/jmatsu/dpg/command"
	"github.com/jmatsu/dpg/command/apps"
)

func Command() cli.Command {
	return cli.Command{
		Name:   "list-teams",
		Usage:  "Show teams which belong to the specified application",
		Action: action,
		Flags:  flags(),
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
	)

	if err != nil {
		return err
	}

	return nil
}

func buildResource(c *cli.Context) (*api.OrganizationAppSharedTeamsEndpoint, *api.Authority, *list.Request, error) {
	authority := api.Authority{
		Token: command.GetApiToken(c),
	}

	platform, err := apps.GetAppPlatform(c)

	if err != nil {
		return nil, nil, nil, err
	}

	endpoint := api.OrganizationAppSharedTeamsEndpoint{
		BaseURL:          "https://deploygate.com",
		OrganizationName: apps.GetAppOwnerName(c),
		AppId:            apps.GetAppId(c),
		AppPlatform:      platform,
	}

	requestParams := list.Request{}

	if err := verifyInput(endpoint, authority, requestParams); err != nil {
		return nil, nil, nil, err
	}

	return &endpoint, &authority, &requestParams, nil
}

func verifyInput(e api.OrganizationAppSharedTeamsEndpoint, authority api.Authority, _ list.Request) error {
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

func listTeams(e api.OrganizationAppSharedTeamsEndpoint, authority api.Authority, requestParams list.Request) (response.AppsSharedTeamsListResponse, error) {
	var r response.AppsSharedTeamsListResponse

	if err := verifyInput(e, authority, requestParams); err != nil {
		return r, err
	}

	if bytes, err := e.GetQueryRequest(authority, requestParams); err != nil {
		return r, err
	} else if err := json.Unmarshal(bytes, &r); err != nil {
		return r, err
	} else {
		return r, nil
	}
}
