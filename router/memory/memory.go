package memory

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
	m.Lock()
	defer m.Unlock()

	pod := r.Pod
	if _, ok := m.routes[r.Pod]; !ok {
		m.routes[pod] = make(map[uint32]*router.Route)
	}
	sum := r.Sum()
	if _, ok := m.routes[pod][sum]; ok {
		return router.ErrDuplicated
	}
	m.routes[pod][sum] = r
	return nil
}

func (m *memory) Deregister(r *router.Route) error {
	m.Lock()
	defer m.Unlock()

	pod := r.Pod
	if _, ok := m.routes[pod]; !ok {
		return nil
	}
	sum := r.Sum()
	if _, ok := m.routes[pod][sum]; !ok {
		return nil
	}
	delete(m.routes[pod], sum)
	return nil
}

func (m *memory) Lookup(pod string) ([]*router.Route, error) {
	m.RLock()
	defer m.RUnlock()

	routes, ok := m.routes[pod]
	if !ok {
		return nil, router.ErrNotFound
	}

	// clone in case that routes change
	clone := make([]*router.Route, 0, len(routes))
	for _, route := range routes {
		clone = append(clone, route)
	}
	return clone, nil
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
