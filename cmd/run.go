package cmd

import (
	"github.com/moobu/moo/client"
	"github.com/moobu/moo/client/http"
	"github.com/moobu/moo/internal/cli"
	"github.com/moobu/moo/runtime"
)

func init() {
	cmd.Register(&cli.Cmd{
		Name:  "run",
		About: "run a pod",
		Pos:   []string{"source"},
		Run:   Run,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "name",
				Usage: "name of the source",
			},
			&cli.StringFlag{
				Name:  "ref",
				Usage: "reference of the source",
			},
			&cli.IntFlag{
				Name:  "replicas",
				Usage: "replicas to be deployed",
				Value: 1,
			},
			&cli.BoolFlag{
				Name:  "gpu",
				Usage: "enable GPU support",
			},
			&cli.BoolFlag{
				Name:  "output",
				Usage: "enable output to stdout",
				Value: true,
			},
			&cli.StringSliceFlag{
				Name:  "env",
				Usage: "environment variables to run with",
			},
			&cli.StringSliceFlag{
				Name:  "arg",
				Usage: "arguments to run with",
			},
			&cli.StringFlag{
				Name:  "server",
				Usage: "address of the server",
				Value: defaultServerAddr,
			},
		},
	})
}

func Run(c cli.Ctx) error {
	// source := c.Pos()[0]
	// bin := c.String("bin")

	cli := http.New(client.Server(c.String("server")))
	cli.Create(&runtime.Pod{})
	return nil
}
