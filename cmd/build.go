package cmd

import (
	"fmt"

	"github.com/moobu/moo/builder"
	"github.com/moobu/moo/client"
	"github.com/moobu/moo/client/http"
	"github.com/moobu/moo/internal/cli"
)

func init() {
	cmd.Register(&cli.Cmd{
		Name: "build",
		Help: "Build a bundle",
		Pos:  []string{"path"},
		Run:  List,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "server",
				Usage: "Address of the server",
				Value: defaultServerAddr,
			},
			&cli.StringFlag{
				Name:  "name",
				Usage: "Name of the bundle",
			},
		},
	})
}

func Build(c cli.Ctx) error {
	path := c.Pos()[0]
	name := c.String("name")
	cli := http.New(client.Server(c.String("server")))
	bundle, err := cli.Build(&builder.Source{Name: name, Dir: path})
	if err != nil {
		return err
	}
	fmt.Printf("type=%s\npath=%s\n", bundle.Type, bundle.Binary)
	return nil
}
