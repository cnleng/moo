package http

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/moobu/moo/server"
)

type httpServer struct {
	options server.Options
}

func (s *httpServer) Serve(l net.Listener) error {
	mux := http.NewServeMux()
	// runtime
	mux.HandleFunc("/runtime/create", Create)
	mux.HandleFunc("/runtime/delete", Delete)
	mux.HandleFunc("/runtime/list", List)
	// routes
	mux.HandleFunc("/router/register", Register)
	mux.HandleFunc("/router/deregister", Deregister)
	mux.HandleFunc("/router/lookup", Lookup)
	// builder
	mux.HandleFunc("/router/build", Build)
	mux.HandleFunc("/router/release", Release)
	mux.HandleFunc("/router/clean", Clean)
	return http.Serve(l, mux)
}

func New(opts ...server.Option) server.Server {
	var options server.Options
	for _, o := range opts {
		o(&options)
	}
	return &httpServer{options: options}
}

func WriteJSON(w http.ResponseWriter, v any, err error) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(args{err.Error(), v})
}

type args struct {
	Error   string
	Content any
}
