package organizations

import (
	"errors"
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/api/request/organizations/update"
	"github.com/jmatsu/dpg/command"
	"github.com/urfave/cli"
)

func UpdateCommand() cli.Command {
	return cli.Command{
		Name:   "update",
		Usage:  "Update an organization",
		Action: command.AuthorizedCommandAction(newUpdateCommand),
		Flags:  updateFlags(),
	}
}

type updateCommand struct {
	endpoint    *api.OrganizationsEndpoint
	requestBody *update.Request
}

func newUpdateCommand(c *cli.Context) (command.Command, error) {
	description := getUpdateDescription(c)

	if !description.Valid {
		return nil, errors.New("description must be specified")
	}

	cmd := updateCommand{
		endpoint: &api.OrganizationsEndpoint{
			BaseURL:          api.EndpointURL,
			OrganizationName: GetOrganizationName(c),
		},
		requestBody: &update.Request{
			Description: description.String,
		},
	}

	if err := cmd.VerifyInput(); err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cmd updateCommand) VerifyInput() error {
	if cmd.endpoint.OrganizationName == "" {
		return errors.New("organization name must be specified")
	}

	if cmd.requestBody.Description == "" {
		return errors.New("organization name must be specified")
	}

	return nil
}

func (cmd updateCommand) Run(authorization *api.Authorization) (string, error) {
	if bytes, err := cmd.endpoint.PatchRequest(*authorization, *cmd.requestBody); err != nil {
		return "", err
	} else {
		return string(bytes), nil
	}
}
