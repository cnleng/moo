package cmd

import (
	"github.com/moobu/moo/internal/cli"
)

func init() {
	cmd.Register(&cli.Cmd{
		Name:  "build",
		About: "build a bundle",
		Pos:   []string{"path"},
		Run:   List,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "server",
				Usage: "address of the server",
				Value: defaultServerAddr,
			},
			&cli.StringFlag{
				Name:  "name",
				Usage: "name of the bundle",
			},
			&cli.StringSliceFlag{
				Name:  "dep",
				Usage: "dependencies to be installed",
			},
		},
	})
}

func Build(c cli.Ctx) (err error) {

	return
}
