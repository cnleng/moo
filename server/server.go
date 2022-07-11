package server

import "net"

type Server interface {
	Serve(net.Listener) error
}

var Default Server

func Serve(l net.Listener) error {
	return Default.Serve(l)
}
