package http

import (
	"encoding/json"
	"net/http"

	"github.com/moobu/moo/builder"
)

func Build(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	source := &builder.Source{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(source); err != nil {
		WriteJSON(w, nil, err)
		return
	}
	defer r.Body.Close()

	bundle, err := builder.Build(source)
	if err != nil {
		WriteJSON(w, bundle, err)
		return
	}
}

func Release(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	bundle := builder.Bundle{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&bundle); err != nil {
		WriteJSON(w, nil, err)
		return
	}
	defer r.Body.Close()

	err := builder.Release(&bundle)
	if err != nil {
		WriteJSON(w, nil, err)
		return
	}
}

func Clean(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	bundle := builder.Bundle{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&bundle); err != nil {
		WriteJSON(w, nil, err)
		return
	}
	defer r.Body.Close()

	err := builder.Clean(&bundle)
	if err != nil {
		WriteJSON(w, nil, err)
		return
	}
}
