package router

type Router interface {
	Handle(*Route) error
	Lookup(string) ([]*Route, error)
}

type Route struct {
	Service string
	Path    string
}
