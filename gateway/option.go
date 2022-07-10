package gateway

import "github.com/moobu/moo/router"

type Options struct {
	Addr   string
	Router router.Router
}

type Option func(*Options)

func Addr(addr string) Option {
	return func(o *Options) {
		o.Addr = addr
	}
}

func Router(r router.Router) Option {
	return func(o *Options) {
		o.Router = r
	}
}
