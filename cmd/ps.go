package cmd

import "github.com/moobu/moo/internal/cli"

func init() {
	cmd.Register(&cli.Cmd{
		Name: "ps",
		Help: "List running pods",
		Run:  List,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "ns",
				Usage: "Filter by namespace",
				Value: "moo",
			},
			&cli.BoolFlag{
				Name:  "all",
				Usage: "List all",
				Value: false,
			},
		},
	})
}

func List(c cli.Ctx) error {
	return nil
}
