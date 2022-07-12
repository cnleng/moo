package cmd

import "github.com/moobu/moo/internal/cli"

func init() {
	cmd.Register(&cli.Cmd{
		Name: "logs",
		Help: "Stream out Pod's output",
		Run:  Logs,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "watch",
				Usage: "Enable live stream",
				Value: false,
			},
		},
	})
}

func Logs(c cli.Ctx) error {
	return nil
}
