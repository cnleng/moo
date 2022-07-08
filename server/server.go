package server

import "net"

type Server interface {
	Serve(net.Listener) error
}

var DefaultServer Server

func Serve(l net.Listener) error {
	return DefaultServer.Serve(l)
}
