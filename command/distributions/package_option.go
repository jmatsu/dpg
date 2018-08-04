package distributions

import (
	"github.com/jmatsu/dpg/command"
	"gopkg.in/urfave/cli.v2"
)

func removeFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		DistributionKey.Flag(),
	}
}
