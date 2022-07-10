package gateway

import (
	"net"
	"net/http"
)

// Gateway passes on the incomming requests to the pods using
// registered handlers. It has only one default implementation.
type Gateway interface {
	Handle(string, Handler) error
	Serve(net.Listener) error
}

// Handler is used by the gateway to pass on the incomming
// requests to the pod using the same protocol the handler uses.
type Handler interface {
	http.Handler
	String() string
}

// here is the only gateway implementation needed
type gateway struct {
	options Options
}

func (g *gateway) Handle(path string, handler Handler) error {
	return nil
}

func (g *gateway) Serve(l net.Listener) error {
	return nil
}

func New(opts ...Option) Gateway {
	var options Options
	for _, o := range opts {
		o(&options)
	}
	return &gateway{options: options}
}
