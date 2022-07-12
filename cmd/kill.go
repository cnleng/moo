package cmd

import "github.com/moobu/moo/internal/cli"

func init() {
	cmd.Register(&cli.Cmd{
		Name: "kill",
		Help: "Kill pods",
		Run:  List,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "ns",
				Usage: "Kill pods in a namespace",
				Value: "moo",
			},
			&cli.BoolFlag{
				Name:  "all",
				Usage: "Kill pods in all namespaces",
				Value: false,
			},
		},
	})
}

func Kill(c cli.Ctx) error {
	return nil
}
