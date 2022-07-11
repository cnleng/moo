package cmd

import "github.com/moobu/moo/internal/cli"

func init() {
	cmd.Register(&cli.Cmd{
		Name: "gateway",
		Help: "Starts Moo API gateway",
		Run:  Gateway,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "port",
				Usage: "Port the gateway listens on",
				Value: 80, // default gateway port
			},
			&cli.BoolFlag{
				Name:  "secure",
				Usage: "Use TLS",
				Value: false,
			},
			&cli.StringFlag{
				Name:  "cert",
				Usage: "TLS cert file",
			},
			&cli.StringFlag{
				Name:  "key",
				Usage: "TLS key file",
			},
		},
	})
}

func Gateway(c cli.Ctx) error {
	return nil
}
