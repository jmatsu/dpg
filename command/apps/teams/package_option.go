package teams

import (
	"github.com/jmatsu/dpg/command"
	"gopkg.in/urfave/cli.v2"
)

func addFlags() []cli.Flag {
	options := append(
		command.OrganizationAppOptions(),
		command.ApiToken,
		command.TeamName,
	)

	return command.ToFlags(options)
}

func listFlags() []cli.Flag {
	options := append(
		command.OrganizationAppOptions(),
		command.ApiToken,
	)

	return command.ToFlags(options)
}

func removeFlags() []cli.Flag {
	options := append(
		command.OrganizationAppOptions(),
		command.ApiToken,
		command.TeamName,
	)

	return command.ToFlags(options)
}
