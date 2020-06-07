package distributions

import (
	"github.com/jmatsu/dpg/command"
	"gopkg.in/urfave/cli.v2"
)

func removeFlags() []cli.Flag {
	options := append(
		[]command.Option{},
		command.ApiToken,
		command.DistributionKey,
	)

	return command.ToFlags(options)
}
