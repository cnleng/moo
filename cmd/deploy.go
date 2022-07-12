package cmd

import (
	"fmt"

	"github.com/moobu/moo/client"
	"github.com/moobu/moo/client/http"
	"github.com/moobu/moo/internal/cli"
	"github.com/moobu/moo/runtime"
)

func init() {
	cmd.Register(&cli.Cmd{
		Name: "deploy",
		Help: "Deploy a pod",
		Pos:  []string{"source"},
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
			&cli.StringFlag{
				Name:  "server",
				Usage: "Address of the server",
				Value: defaultServerAddr,
			},
		},
	})
}

func Deploy(c cli.Ctx) error {
	source := c.Pos()[0]
	fmt.Println(source)
	cli := http.New(client.Server(c.String("server")))
	cli.Create(&runtime.Pod{})
	return nil
}
