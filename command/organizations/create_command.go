package organizations

import (
	"github.com/jmatsu/dpg/api"
	"github.com/jmatsu/dpg/command"
	"gopkg.in/guregu/null.v3"
	"gopkg.in/urfave/cli.v2"
)

func CreateCommand() *cli.Command {
	return &cli.Command{
		Name:   "create",
		Usage:  "Create an organization",
		Action: command.AuthorizedCommandAction(NewCreateCommand),
		Flags:  createFlags(),
	}
}

type createCommand struct {
	organizationName string
	description null.String
}

func NewCreateCommand(c *cli.Context) (command.Command, error) {
	organizationName, err := command.RequireOrganizationName(c)

	if err != nil {
		return nil, err
	}

	description, err := command.GetOrganizationDescription(c)

	if err != nil {
		return nil, err
	}

	cmd := createCommand{
			organizationName: organizationName,
			description: description,
	}

	return cmd, nil
}

func (cmd createCommand) Run(authorization *api.Authorization) (string, error) {
	return api.NewClient(*authorization).CreateOrganization(cmd.organizationName, cmd.description)
}
