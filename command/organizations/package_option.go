package organizations

import (
	"github.com/jmatsu/dpg/command"
	"gopkg.in/urfave/cli.v2"
)

func createFlags() []cli.Flag {
	options := append(
		[]command.Option{},
		command.ApiToken,
		command.OrganizationName,
		command.OrganizationDescription,
	)

	return command.ToFlags(options)
}

func destroyFlags() []cli.Flag {
	options := append(
		[]command.Option{},
		command.ApiToken,
		command.OrganizationName,
	)

	return command.ToFlags(options)
}

func listFlags() []cli.Flag {
	options := append(
		[]command.Option{},
		command.ApiToken,
	)

	return command.ToFlags(options)
}

func showFlags() []cli.Flag {
	options := append(
		[]command.Option{},
		command.ApiToken,
		command.OrganizationName,
	)

	return command.ToFlags(options)
}

func updateFlags() []cli.Flag {
	options := append(
		[]command.Option{},
		command.ApiToken,
		command.OrganizationName,
		command.OrganizationDescription,
	)

	return command.ToFlags(options)
}
