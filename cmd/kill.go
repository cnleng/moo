package cmd

import "github.com/moobu/moo/internal/cli"

func init() {
	cmd.Register(&cli.Cmd{
		Name: "kill",
		Help: "Kill one or more pods",
		Run:  List,
	})
}

func Kill(c cli.Ctx) error {
	return nil
}
