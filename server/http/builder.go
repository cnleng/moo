package http

import (
	"encoding/json"
	"net/http"

	"github.com/moobu/moo/builder"
)

type BuildArgs struct {
	Source  *builder.Source
	Options *builder.BuildOptions
}

type ReleaseArgs struct {
	Bundle  *builder.Bundle
	Options *builder.ReleaseOptions
}

type CleanArgs struct {
	Bundle  *builder.Bundle
	Options *builder.CleanOptions
}

func Build(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	args := &BuildArgs{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(args); err != nil {
		WriteJSON(w, nil, err)
		return
	}
	defer r.Body.Close()

	bundle, err := builder.Build(args.Source)
	WriteJSON(w, bundle, err)
}

func Release(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	args := ReleaseArgs{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&args); err != nil {
		WriteJSON(w, nil, err)
		return
	}
	defer r.Body.Close()

	err := builder.Release(args.Bundle)
	WriteJSON(w, nil, err)
}

func Clean(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	args := CleanArgs{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&args); err != nil {
		WriteJSON(w, nil, err)
		return
	}
	defer r.Body.Close()

	err := builder.Clean(args.Bundle)
	WriteJSON(w, nil, err)
}
