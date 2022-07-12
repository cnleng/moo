package cmd

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/moobu/moo/builder"
	"github.com/moobu/moo/internal/cli"
	"github.com/moobu/moo/preset"
	"github.com/moobu/moo/runtime"
	"github.com/moobu/moo/server"
)

func init() {
	cmd.Register(&cli.Cmd{
		Name: "server",
		Help: "Starts Moo server",
		Run:  Server,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "port",
				Usage: "Port the server listens on",
				Value: 11451, // default port
			},
			&cli.BoolFlag{
				Name:  "uds",
				Usage: "Use unix domain socket address",
				Value: false,
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
			&cli.BoolFlag{
				Name:  "gateway",
				Usage: "Starts the API gateway",
				Value: false,
			},
		},
	})
}

func Server(c cli.Ctx) error {
	set := c.String("preset")
	if err := preset.Use(c, set); err != nil {
		return err
	}
	log.Printf("[INFO] using preset: %s", set)

	uds := c.Bool("uds")
	ln, err := listen(c, uds)
	if err != nil {
		return err
	}
	addr := ln.Addr().String()

	errCh := make(chan error, 1)
	sigCh := make(chan os.Signal, 1)

	// starts the server and listens for termination
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	go func(l net.Listener) {
		err := server.Serve(l)
		errCh <- err
	}(ln)

	// see if we need to initiate the API geteway
	if c.Bool("gateway") {
		bin := os.Args[0]
		err := runtime.Default.Create(&runtime.Pod{Name: "gateway"},
			runtime.CreateWithNamespace("moo"),
			runtime.Bundle(&builder.Bundle{Binary: bin}),
			runtime.Args("gateway", "--server", addr),
			runtime.Output(os.Stdout))
		if err != nil {
			return err
		}
	}
	// start the runtime
	if err := runtime.Default.Start(); err != nil {
		return err
	}

	log.Printf("[INFO] server started at %s", addr)
	select {
	case err := <-errCh:
		return err
	case <-c.Done():
		return c.Err()
	case <-sigCh:
		log.Print("[INFO] stopping server")
		runtime.Default.Stop()
		return ln.Close()
	}
}
