package cmd

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/moobu/moo/client"
	"github.com/moobu/moo/client/http"
	"github.com/moobu/moo/gateway"
	proxy "github.com/moobu/moo/gateway/http"
	"github.com/moobu/moo/internal/cli"
	"github.com/moobu/moo/router"
)

func init() {
	cmd.Register(&cli.Cmd{
		Name: "gateway",
		Help: "run the API gateway",
		Run:  Gateway,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "server",
				Usage: "address of the server",
				Value: defaultServerAddr,
			},
			&cli.IntFlag{
				Name:  "port",
				Usage: "port the gateway listens on",
				Value: defaultGatewayPort,
			},
			&cli.BoolFlag{
				Name:  "secure",
				Usage: "enable TLS",
			},
			&cli.StringFlag{
				Name:  "cert",
				Usage: "path to TLS certificate",
			},
			&cli.StringFlag{
				Name:  "key",
				Usage: "path to TLS public key",
			},
		},
	})
}

func Gateway(c cli.Ctx) error {
	ln, err := listen(c, false)
	if err != nil {
		return err
	}
	addr := ln.Addr().String()
	gw := gateway.New()
	errCh := make(chan error, 1)
	sigCh := make(chan os.Signal, 1)

	// starts the server and listens for termination
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	go func(l net.Listener) {
		err := gw.Serve(l)
		errCh <- err
	}(ln)

	// use the client as the router
	server := c.String("server")
	client := http.New(client.Server(server))
	gw.Handle(proxy.New(gateway.Router(client)))

	// register the gateway itself
	route := &router.Route{Pod: "/moo/gateway", Address: addr, Protocol: "http"}
	if err := client.Register(route); err != nil {
		return err
	}

	log.Printf("[INFO] gateway started at %s", addr)
	select {
	case err := <-errCh:
		return err
	case <-c.Done():
		return c.Err()
	case <-sigCh:
		log.Print("[INFO] stopping gateway")
		return ln.Close()
	}
}
