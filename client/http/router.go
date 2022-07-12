package http

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/moobu/moo/internal/pool/buffer"
	"github.com/moobu/moo/router"
)

func (h *http) Register(route *router.Route) error {
	url := fmt.Sprintf("http://%s/register", h.options.Server)

	reader := buffer.Get()
	defer buffer.Put(reader)

	encoder := json.NewEncoder(reader)
	if err := encoder.Encode(route); err != nil {
		return err
	}
	_, err := h.client.Post(url, contentType, reader)
	return err
}

func (h *http) Deregister(route *router.Route) error {
	url := fmt.Sprintf("http://%s/deregister", h.options.Server)

	reader := buffer.Get()
	defer buffer.Put(reader)

	encoder := json.NewEncoder(reader)
	if err := encoder.Encode(route); err != nil {
		return err
	}
	_, err := h.client.Post(url, contentType, reader)
	return err
}

func (h *http) Lookup(pod string) ([]*router.Route, error) {
	route := router.Route{Pod: pod}
	url := fmt.Sprintf("http://%s/lookup", h.options.Server)

	reader := buffer.Get()
	defer buffer.Put(reader)

	encoder := json.NewEncoder(reader)
	if err := encoder.Encode(route); err != nil {
		return nil, err
	}
	res, err := h.client.Post(url, contentType, reader)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == 404 {
		return nil, router.ErrNotFound
	}

	retval := &realLookupResponse{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(retval); err != nil {
		return nil, err
	}
	defer res.Body.Close()
	// This is when HTTP sucks!
	if len(retval.Error) > 0 {
		return nil, errors.New(retval.Error)
	}
	return retval.Content, nil
}

type realLookupResponse struct {
	Error   string
	Content []*router.Route
}
