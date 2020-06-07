package distributions

import (
	"github.com/jmatsu/dpg/command"
	"gopkg.in/urfave/cli.v2"
)

func destroyFlags() []cli.Flag {
	options := append(
		command.UserAppOptions(),
		command.ApiToken,
		command.DistributionName,
	)

	return command.ToFlags(options)
}
