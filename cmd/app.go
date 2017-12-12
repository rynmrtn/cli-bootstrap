package cmd

import (
	"github.com/urfave/cli"
)

func App() *cli.App {
	app = cli.NewApp()
	app.Name = ""
	app.Usage = ""
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "sample-flag",
			Usage:  "",
			EnvVar: "SAMPLE_VAR",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "subcmd",
			Usage:  "",
			Action: RunSubcmd,
		},
	}

	return app
}
