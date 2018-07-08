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
	endpoint, authority, requestBody, err := buildResource(c)

	if err != nil {
		return err
	}

	_, err = addToTeam(
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

func buildResource(c *cli.Context) (*api.OrganizationTeamsEndpoint, *api.Authority, *add.Request, error) {
	authority := api.Authority{
		Token: apiToken.Value(c).(string),
	}

	endpoint := api.OrganizationTeamsEndpoint{
		BaseURL:          "https://deploygate.com",
		OrganizationName: organizationName.Value(c).(string),
		AppId:            appId.Value(c).(string),
	}

	requestBody := add.Request{
		TeamName: teamName.Value(c).(string),
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
