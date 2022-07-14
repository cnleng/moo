package auto

import (
	"errors"

	"github.com/moobu/moo/builder"
)

// auto is a high-level implementation of the builder interface
type auto struct {
	options builder.Options
	next    builder.Builder
}

func (a *auto) Build(s *builder.Source, opts ...builder.BuildOption) (*builder.Bundle, error) {
	if len(s.Local) == 0 && len(s.Remote) == 0 {
		return nil, errors.New("either local or remote source is needed")
	}
	if len(s.Local) > 0 {
		return a.next.Build(s, opts...)
	}
	// we retrieve the source using the given retriever
	// if no local repository is provided.
	src, err := a.options.Retriever.Retrieve(s.Remote)
	if err != nil {
		return nil, err
	}
	return a.next.Build(src, opts...)
}

func (a auto) Clean(bun *builder.Bundle, opts ...builder.CleanOption) error {
	return a.next.Clean(bun, opts...)
}

func (a auto) String() string {
	return "auto"
}

// New creates an auto builder whose Build method retrieves the source
// code from the given remote repository if no local one is given and
// then uses the next builder to build the source code.
func New(next builder.Builder, opts ...builder.Option) builder.Builder {
	var options builder.Options
	for _, o := range opts {
		o(&options)
	}
	return &auto{
		options: options,
		next:    next,
	}
}
