package cmd

import (
	"fmt"
	"os"

	"github.com/moobu/moo/builder"
	"github.com/moobu/moo/client"
	"github.com/moobu/moo/client/http"
	"github.com/moobu/moo/internal/cli"
)

func init() {
	cmd.Register(&cli.Cmd{
		Name: "build",
		Help: "build a bundle",
		Pos:  []string{"path"},
		Run:  List,
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
	name := c.String("name")
	deps := c.StringSlice("dep")

	dir := c.Pos()[0]
	if dir == "." {
		if dir, err = os.Getwd(); err != nil {
			return
		}
	}

	source := &builder.Source{
		Name: name,
		Dir:  dir,
	}

	cli := http.New(client.Server(c.String("server")))
	bundle, err := cli.Build(source, builder.Deps(deps...))
	if err != nil {
		return
	}

	fmt.Printf("%s\n", bundle.Sum)
	return
}
