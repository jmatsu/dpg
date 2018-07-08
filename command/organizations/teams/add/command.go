package organizations_teams_add

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/urfave/cli"
	"strings"
	"github.com/jmatsu/dpg/api/response"
	"encoding/json"
	"github.com/jmatsu/dpg/api/request/organizations/teams/add"
)

func Command() cli.Command {
	return cli.Command{
		Name:    "add-teams",
		Aliases: []string{"i"},
		Usage:   "Add a team to the specified organization",
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

	_, err := addToTeam(
		endpoint,
		authority,
		add.Request{
			TeamName: teamName.Value(c).(string),
		},
		c.GlobalBoolT("verbose"),
	)

	if err != nil {
		return err
	}

	return nil
}

func addToTeam(e api.OrganizationTeamsEndpoint, authority api.Authority, requestBody add.Request, verbose bool) (response.OrganizationTeamsListResponse, error) {
	var r response.OrganizationTeamsListResponse

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

func verifyInput(e api.OrganizationTeamsEndpoint, authority api.Authority, request add.Request) error {
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

	if request.TeamName == "" {
		return errors.New("team name must be specified")
	}

	return nil
}
