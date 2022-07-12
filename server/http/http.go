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
	mux.HandleFunc("/create", Create)
	mux.HandleFunc("/delete", Delete)
	mux.HandleFunc("/list", List)
	// routes
	mux.HandleFunc("/register", Register)
	mux.HandleFunc("/deregister", Deregister)
	mux.HandleFunc("/lookup", Lookup)
	// builder
	mux.HandleFunc("/build", Build)
	mux.HandleFunc("/release", Release)
	mux.HandleFunc("/clean", Clean)
	return http.Serve(l, mux)
}

func (httpServer) String() string {
	return "http"
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
	errmsg := ""
	if err != nil {
		errmsg = err.Error()
	}
	return json.NewEncoder(w).Encode(Args{errmsg, v})
}

type Args struct {
	Error   string
	Content any
}
