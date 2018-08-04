package distributions

import (
	"github.com/jmatsu/dpg/command"
	"gopkg.in/urfave/cli.v2"
)

func destroyFlags() []cli.Flag {
	return []cli.Flag{
		command.ApiToken.Flag(),
		DistributionName.Flag(),
	}
}
