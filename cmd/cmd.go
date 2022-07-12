package cmd

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/moobu/moo/internal/cli"
)

const defaultServerAddr = "127.0.0.1:11451"

var cmd = &cli.Cmd{
	Name:     "moo",
	Help:     "A pluggable serverless engine",
	Version:  "v0.0.1",
	Wildcard: true,
}

func Run() error {
	cmd.Init()
	return cmd.RunCtx(context.Background())
}

func listen(c cli.Ctx, uds bool) (net.Listener, error) {
	// we use TCP if the flag uds is not set,
	// otherwise use the UNIX doamin socket
	network := "tcp"
	address := fmt.Sprintf(":%d", c.Int("port"))
	if uds {
		address = filepath.Join(os.TempDir(), "moo.sock")
		network = "unix"
	}

	listener, err := net.Listen(network, address)
	if err != nil {
		return nil, err
	}

	// see if we need to wrap a TLS listener
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

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
