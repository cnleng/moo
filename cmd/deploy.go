package cmd

import (
	"fmt"

	"github.com/moobu/moo/internal/cli"
)

func init() {
	cmd.Register(&cli.Cmd{
		Name: "deploy",
		Help: "Deploy a source",
		Args: []string{"source"},
		Run:  Deploy,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "replicas",
				Usage: "Replicas to deploy",
				Value: 1,
			},
			&cli.BoolFlag{
				Name:  "gpu",
				Usage: "Enable GPU support",
				Value: false,
			},
			&cli.BoolFlag{
				Name:  "output",
				Usage: "Enable output to stdout",
				Value: true,
			},
			&cli.StringSliceFlag{
				Name:  "env",
				Usage: "Env variables passing in",
			},
			&cli.StringSliceFlag{
				Name:  "arg",
				Usage: "Arguments passing in",
			},
		},
	})
}

func Deploy(c cli.Ctx) error {
	args := c.Args()
	fmt.Println(args)
	return nil
}
