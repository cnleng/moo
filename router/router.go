package router

type Router interface {
	// Register adds a route to a pod
	Register(*Route) error
	// Deregister removes a route to a pod
	Deregister(*Route) error
	// Lookup finds all routes to the same pod
	Lookup(string) ([]*Route, error)
}

type Route struct {
	Pod      string // pod name (fmt. {namespace}:{name})
	Protocol string // the protocol by which we communicate with the pod
	Address  string // pod address (e.g. 10.0.0.1:80, /tmp/moo/xxx.sock)
}
