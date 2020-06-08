package members

import (
	"github.com/jmatsu/dpg/command"
	"gopkg.in/urfave/cli.v2"
)

func addFlags() []cli.Flag {
	options := append(
		command.EnterpriseOptions(),
		command.ApiToken,
		command.OrganizationName,
		command.UserName,
	)

	return command.ToFlags(options)
}

func listFlags() []cli.Flag {
	options := append(
		command.EnterpriseOptions(),
		command.ApiToken,
		command.OrganizationName,
	)

	return command.ToFlags(options)
}

func removeFlags() []cli.Flag {
	options := append(
		command.EnterpriseOptions(),
		command.ApiToken,
		command.OrganizationName,
		command.UserName,
	)

	return command.ToFlags(options)
}
