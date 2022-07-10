package http

import (
	"fmt"
	"net/http"

	"github.com/moobu/moo/gateway"
)

type handler struct {
	options gateway.ProxyOptions
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// for now, we only allow POST requests
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	pod := r.URL.Path[4:] // strip the protocol name
	routes, err := h.options.Router.Lookup(pod)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// TODO: select one route from routes
	fmt.Fprint(w, pod)
}

func (h handler) String() string {
	return "http"
}

func New(opts ...gateway.ProxyOption) gateway.Proxy {
	var options gateway.ProxyOptions
	for _, o := range opts {
		o(&options)
	}
	return &handler{options: options}
}
