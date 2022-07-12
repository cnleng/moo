package cmd

import "github.com/moobu/moo/internal/cli"

func init() {
	cmd.Register(&cli.Cmd{
		Name: "kill",
		Help: "kill one or more pods",
		Run:  Kill,
	})
}

func Kill(c cli.Ctx) error {
	return nil
}
