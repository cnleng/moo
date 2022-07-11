package http

import (
	"encoding/json"
	"net/http"

	"github.com/moobu/moo/runtime"
)

type creArgs struct {
	Pod     *runtime.Pod
	Options *runtime.CreateOptions
}

func Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	args := &creArgs{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(args); err != nil {
		WriteJSON(w, nil, err)
		return
	}
	defer r.Body.Close()

	options := args.Options
	opts := []runtime.CreateOption{
		runtime.Args(options.Args...),
		runtime.Env(options.Env...),
		runtime.Image(options.Image),
		runtime.Replicas(options.Replicas),
		runtime.Bundle(options.Bundle),
		runtime.CreateWithNamespace(options.Namespace),
	}
	err := runtime.Create(args.Pod, opts...)
	WriteJSON(w, nil, err)
}

type delArgs struct {
	Pod     *runtime.Pod
	Options *runtime.DeleteOptions
}

func Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	args := &delArgs{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(args); err != nil {
		WriteJSON(w, nil, err)
		return
	}
	defer r.Body.Close()

	options := args.Options
	err := runtime.Delete(args.Pod, runtime.DeleteWithNamespace(options.Namespace))
	WriteJSON(w, nil, err)
}

func List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	options := runtime.ListOptions{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&options); err != nil {
		WriteJSON(w, nil, err)
		return
	}
	defer r.Body.Close()

	opts := []runtime.ListOption{
		runtime.Name(options.Name),
		runtime.Tag(options.Tag),
		runtime.ListWithNamespace(options.Namespace),
	}
	pods, err := runtime.List(opts...)
	WriteJSON(w, pods, err)
}
