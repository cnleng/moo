package http

import (
	std "net/http"

	"github.com/moobu/moo/client"
)

const contentType = "application/json"

type http struct {
	options client.Options
	client  *std.Client
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
