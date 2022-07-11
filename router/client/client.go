package client

import "github.com/moobu/moo/router"

type client struct {
	options router.Options
}

func (c *client) Register(r *router.Route) error {

	return nil
}

func (c *client) Deregister(r *router.Route) error {

	return nil
}

func (c *client) Lookup(pod string) ([]*router.Route, error) {

	return nil, router.ErrNotFound
}

func New(opts ...router.Option) router.Router {
	var options router.Options
	for _, o := range opts {
		o(&options)
	}
	return &client{options: options}
}
