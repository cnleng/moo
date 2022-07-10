package local

import (
	"sync"

	"github.com/moobu/moo/router"
)

type memory struct {
	sync.RWMutex
	options router.Options
	routes  map[string]map[uint32]*router.Route
}

func (m *memory) Register(r *router.Route) error {
	return nil
}

func (m *memory) Deregister(r *router.Route) error {
	return nil
}

func (m *memory) Lookup(pod string) ([]*router.Route, error) {
	return nil, nil
}

func New(opts ...router.Option) router.Router {
	var options router.Options
	for _, o := range opts {
		o(&options)
	}
	return &memory{
		options: options,
		routes:  make(map[string]map[uint32]*router.Route),
	}
}
