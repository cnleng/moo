package cmd

import "github.com/moobu/moo/internal/cli"

func init() {
	cmd.Register(&cli.Cmd{
		Name: "list",
		Help: "List pods",
		Run:  List,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "ns",
				Usage: "List pods in a namespace",
				Value: "moo",
			},
			&cli.BoolFlag{
				Name:  "all",
				Usage: "List pods in all namespaces",
				Value: false,
			},
		},
	})
}

func List(c cli.Ctx) error {
	return nil
}
