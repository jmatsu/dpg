package members

import (
	"github.com/jmatsu/dpg/command"
	"gopkg.in/urfave/cli.v2"
)

func addFlags() []cli.Flag {
	options := append(
		command.UserAppOptions(),
		command.ApiToken,
		command.Invitees,
		command.DeveloperRole,
	)

	return command.ToFlags(options)
}

func listFlags() []cli.Flag {
	options := append(
		command.UserAppOptions(),
		command.ApiToken,
	)

	return command.ToFlags(options)
}

func removeFlags() []cli.Flag {
	options := append(
		command.UserAppOptions(),
		command.ApiToken,
		command.Removees,
	)

	return command.ToFlags(options)
}
