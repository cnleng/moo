package cmd

import "github.com/moobu/moo/internal/cli"

func init() {
	cmd.Register(&cli.Cmd{
		Name:  "logs",
		About: "output log file",
		Run:   Logs,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "watch",
				Usage: "enable live stream",
			},
		},
	})
}

func Logs(c cli.Ctx) error {
	return nil
}
