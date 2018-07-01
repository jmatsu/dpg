package main

import (
	"github.com/jmatsu/dpg/command/upload"
	"github.com/urfave/cli"
)

func Upload() cli.Command {
	return cli.Command{
		Name:    "upload",
		Aliases: []string{"u"},
		Usage:   "Upload either android application or iOS application to the specified owner space",
		Action:  upload.App,
		Flags: []cli.Flag{
			upload.AppFilePath.Flag(),
			upload.ApiToken.Flag(),
			upload.AppOwnerName.Flag(),
		},
		Subcommands: []cli.Command{
			{
				Name:    "android",
				Aliases: []string{"a"},
				Usage:   "Upload an android application to the specified owner space",
				Action:  upload.AndroidApp,
				Flags: []cli.Flag{
					upload.AppFilePath.Flag(),
					upload.ApiToken.Flag(),
					upload.AppOwnerName.Flag(),
				},
			},
			{
				Name:    "ios",
				Aliases: []string{"i"},
				Usage:   "Upload an iOS application to the specified owner space",
				Action:  upload.IOSApp,
				Flags: []cli.Flag{
					upload.AppFilePath.Flag(),
					upload.ApiToken.Flag(),
					upload.AppOwnerName.Flag(),
					upload.SuppressNotification.Flag(),
				},
			},
		},
	}
}
