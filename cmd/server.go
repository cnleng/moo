package cmd

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/moobu/moo/internal/cli"
	"github.com/moobu/moo/presets"
	"github.com/moobu/moo/server"
)

func init() {
	cmd.Register(&cli.Cmd{
		Name: "server",
		Help: "Starts Moo server",
		Run:  Server,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "presets",
				Usage: "Presets initializing the server",
				Value: "local", // default presets
			},
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
		},
	})
}

func Server(c cli.Ctx) error {
	preset := c.String("presets")
	if err := presets.Use(c, preset); err != nil {
		return err
	}
	log.Printf("[INFO] using presets: %s", preset)

	l, err := listen(c)
	if err != nil {
		return err
	}

	errCh := make(chan error)
	sigCh := make(chan os.Signal)

	go func(l net.Listener) {
		err := server.Serve(l)
		errCh <- err
	}(l)

	log.Printf("[INFO] listening on %s", l.Addr())

	select {
	case err := <-errCh:
		return err
	case <-c.Done():
		return c.Err()
	case <-sigCh:
		// TODO: do some close stuff if necessary?
		log.Printf("[INFO] server stopped")
		return nil
	}
}

func listen(c cli.Ctx) (net.Listener, error) {
	//  we use the unix domain socket if using the local
	//  presets or explicitly emitted the uds flag
	network := "tcp"
	address := fmt.Sprintf(":%d", c.Int("port"))
	if c.Bool("uds") {
		address = filepath.Join(os.TempDir(), "moo.sock")
		network = "unix"
	}

	listener, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}

	// see if we need the TLS listener
	if !c.Bool("secure") {
		return listener, nil
	}

	cert, key := c.String("cert"), c.String("key")
	if len(cert) == 0 || len(key) == 0 {
		return nil, errors.New("certificates not provided")
	}

	certificate, err := tls.LoadX509KeyPair(cert, key)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAnyClientCert,
	}
	return tls.NewListener(listener, config), nil
}
