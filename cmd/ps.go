package cmd

import "github.com/moobu/moo/internal/cli"

func init() {
	cmd.Register(&cli.Cmd{
		Name: "ps",
		Help: "list running pods",
		Run:  List,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "ns",
				Usage: "filter by namespace",
				Value: "moo",
			},
			&cli.BoolFlag{
				Name:  "all",
				Usage: "list all",
			},
		},
	})
}

func List(c cli.Ctx) error {
	return nil
}
