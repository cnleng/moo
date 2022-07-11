package http

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/moobu/moo/gateway"
)

type proxy struct {
	options gateway.ProxyOptions
}

func (p proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// for now, we only allow POST requests
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// strip out the protocol name and find out
	// all the routes to the pod
	pod := r.URL.Path[5:]
	routes, err := p.options.Router.Lookup(pod)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	// select a random route
	route := routes[rand.Int()%len(routes)]
	rawURL := fmt.Sprintf("http://%s", route.Address)
	target, err := url.Parse(rawURL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// reverse proxy to the selected route
	httputil.NewSingleHostReverseProxy(target).ServeHTTP(w, r)
}

func (p proxy) String() string {
	return "http"
}

func New(opts ...gateway.ProxyOption) gateway.Proxy {
	var options gateway.ProxyOptions
	for _, o := range opts {
		o(&options)
	}
	return &proxy{options: options}
}
