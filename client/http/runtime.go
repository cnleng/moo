package http

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/moobu/moo/internal/pool/buffer"
	"github.com/moobu/moo/runtime"
	server "github.com/moobu/moo/server/http"
)

func (h *http) Create(pod *runtime.Pod, opts ...runtime.CreateOption) error {
	var options runtime.CreateOptions
	for _, o := range opts {
		o(&options)
	}

	url := fmt.Sprintf("http://%s/create", h.options.Server)
	reader := buffer.Get()
	defer buffer.Put(reader)
	encoder := json.NewEncoder(reader)
	args := &server.CreateArgs{}
	if err := encoder.Encode(&args); err != nil {
		return err
	}
	_, err := h.client.Post(url, contentType, reader)
	return err
}

func (h *http) Delete(pod *runtime.Pod, opts ...runtime.DeleteOption) error {
	var options runtime.DeleteOptions
	for _, o := range opts {
		o(&options)
	}

	url := fmt.Sprintf("http://%s/delete", h.options.Server)
	reader := buffer.Get()
	defer buffer.Put(reader)
	encoder := json.NewEncoder(reader)
	args := &server.DeleteArgs{}
	if err := encoder.Encode(&args); err != nil {
		return err
	}
	_, err := h.client.Post(url, contentType, reader)
	return err
}

func (h *http) List(opts ...runtime.ListOption) ([]*runtime.Pod, error) {
	var options runtime.ListOptions
	for _, o := range opts {
		o(&options)
	}

	url := fmt.Sprintf("http://%s/list", h.options.Server)
	reader := buffer.Get()
	defer buffer.Put(reader)
	encoder := json.NewEncoder(reader)
	args := &runtime.ListOptions{}
	if err := encoder.Encode(&args); err != nil {
		return nil, err
	}
	res, err := h.client.Post(url, contentType, reader)
	if err != nil {
		return nil, err
	}

	retval := &realListResponse{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(retval); err != nil {
		return nil, err
	}
	return retval.Content, errors.New(retval.Error)
}

type realListResponse struct {
	Error   string
	Content []*runtime.Pod
}

func (h *http) Start() error {
	return nil
}

func (h *http) Stop() error {
	return nil
}
