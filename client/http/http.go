package http

import (
	std "net/http"

	"github.com/moobu/moo/builder"
	"github.com/moobu/moo/client"
	"github.com/moobu/moo/runtime"
)

const contentType = "application/json"

type http struct {
	options client.Options
	client  *std.Client
	// Temporarily trick the go compiler
	// TODO: implement them two later
	runtime.Runtime
	builder.Builder
}

func New(opts ...client.Option) client.Client {
	var options client.Options
	for _, o := range opts {
		o(&options)
	}
	return &http{
		options: options,
		client:  &std.Client{},
	}
}
