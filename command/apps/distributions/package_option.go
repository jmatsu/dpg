package distributions

import (
	"github.com/jmatsu/dpg/command"
	"github.com/urfave/cli"
)

func removeFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		DistributionKey.Flag(),
	}
}
