package cmd

import "github.com/moobu/moo/internal/cli"

func init() {
	cmd.Register(&cli.Cmd{
		Name:  "release",
		Help:  "release a bundle",
		Pos:   []string{"path"},
		Run:   List,
		Flags: []cli.Flag{},
	})
}

func Release(c cli.Ctx) error {
	return nil
}
