package http

import (
	"encoding/json"
	"net/http"

	"github.com/moobu/moo/router"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	route := &router.Route{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(route); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err := router.Register(route)
	if err != nil {
		WriteJSON(w, nil, err)
		return
	}
}

func Deregister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	route := &router.Route{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(route); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err := router.Deregister(route)
	if err != nil {
		WriteJSON(w, nil, err)
		return
	}
}

func Lookup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	args := lookupArgs{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&args); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	routes, err := router.Lookup(args.Pod)
	if err != nil {
		WriteJSON(w, routes, err)
		return
	}
}

type lookupArgs struct {
	Pod string
}
