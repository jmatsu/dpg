package members

import (
	"github.com/jmatsu/dpg/command"
	"gopkg.in/urfave/cli.v2"
)

func addFlags() []cli.Flag {
	options := append(
		[]command.Option{},
		command.ApiToken,
		command.EnterpriseName,
		command.UserName,
	)

	return command.ToFlags(options)
}

func listFlags() []cli.Flag {
	options := append(
		[]command.Option{},
		command.ApiToken,
		command.EnterpriseName,
	)

	return command.ToFlags(options)
}

func removeFlags() []cli.Flag {
	options := append(
		[]command.Option{},
		command.ApiToken,
		command.EnterpriseName,
		command.UserName,
	)

	return command.ToFlags(options)
}
