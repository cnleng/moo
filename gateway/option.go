package gateway

import "github.com/moobu/moo/router"

type Options struct {
	Router router.Router
}

type Option func(*Options)

func Router(r router.Router) Option {
	return func(o *Options) {
		o.Router = r
	}
}
