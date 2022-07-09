package runtime

import (
	"io"

	"github.com/moobu/moo/builder"
)

type Options struct{}

type Option func(*Options)

type CreateOptions struct {
	Args      []string
	Env       []string
	Bundle    *builder.Bundle
	Output    io.Writer
	Image     string
	Namespace string
	Replicas  int
}

type CreateOption func(*CreateOptions)

type ListOptions struct {
	Name      string
	Tag       string
	Namespace string
}

type ListOption func(*ListOptions)

type DeleteOptions struct {
	Namespace string
}

type DeleteOption func(*DeleteOptions)
