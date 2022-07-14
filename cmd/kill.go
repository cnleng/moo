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
		Name:  "kill",
		About: "kill one or more pods",
		Pos:   []string{"pod"},
		Run:   Kill,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "server",
				Usage: "specify the address of Moo server",
				Value: defaultServerAddr,
			},
		},
	})
}

func Kill(c cli.Ctx) error {
	// connect to Moo server
	cli := http.New(client.Server(c.String("server")))
	rawPod := c.Pos()[0]
	pod, err := runtime.Parse(rawPod)
	if err != nil {
		return err
	}
	// TODO: specify the user's namespace
	err = cli.Delete(pod, runtime.DeleteWithNamespace("default"))
	if err != nil {
		return err
	}
	fmt.Printf("Killed %s\n", pod)
	return nil
}
