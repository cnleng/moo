package runtime

import (
	"io"

	"github.com/mooctl/moo/builder"
)

type Options struct{}

type Option func(*Options)

type CreateOptions struct {
	Namespace string
	Bundle    *builder.Bundle
	Image     string
	Args      []string
	Env       []string
	Output    io.Writer
}

type CreateOption func(*CreateOptions)
