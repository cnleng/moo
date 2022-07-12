package http

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/moobu/moo/builder"
	"github.com/moobu/moo/internal/pool/buffer"
	server "github.com/moobu/moo/server/http"
)

type realBuildResponse struct {
	Error   string
	Content *builder.Bundle
}

func (h *http) Build(s *builder.Source, opts ...builder.BuildOption) (*builder.Bundle, error) {
	var options builder.BuildOptions
	for _, o := range opts {
		o(&options)
	}
	args := server.BuildArgs{Source: s, Options: &options}
	url := fmt.Sprintf("http://%s/build", h.options.Server)

	reader := buffer.Get()
	defer buffer.Put(reader)

	encoder := json.NewEncoder(reader)
	if err := encoder.Encode(args); err != nil {
		return nil, err
	}

	res, err := h.client.Post(url, contentType, reader)
	if err != nil {
		return nil, err
	}

	retval := &realBuildResponse{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(retval); err != nil {
		return nil, err
	}
	return retval.Content, errors.New(retval.Error)
}

func (h *http) Release(b *builder.Bundle, opts ...builder.ReleaseOption) error {
	var options builder.ReleaseOptions
	for _, o := range opts {
		o(&options)
	}

	args := server.ReleaseArgs{Bundle: b, Options: &options}
	url := fmt.Sprintf("http://%s/release", h.options.Server)

	reader := buffer.Get()
	defer buffer.Put(reader)

	encoder := json.NewEncoder(reader)
	if err := encoder.Encode(args); err != nil {
		return err
	}
	_, err := h.client.Post(url, contentType, reader)
	return err
}

func (h *http) Clean(b *builder.Bundle, opts ...builder.CleanOption) error {
	var options builder.CleanOptions
	for _, o := range opts {
		o(&options)
	}

	args := server.CleanArgs{Bundle: b, Options: &options}
	url := fmt.Sprintf("http://%s/clean", h.options.Server)

	reader := buffer.Get()
	defer buffer.Put(reader)

	encoder := json.NewEncoder(reader)
	if err := encoder.Encode(args); err != nil {
		return err
	}
	_, err := h.client.Post(url, contentType, reader)
	return err
}
