package retriever

import (
	"errors"

	"github.com/moobu/moo/builder"
)

// retriever is a high-level implementation of the builder interface
type retriever struct {
	options builder.Options
	next    builder.Builder
}

func (r *retriever) Build(s *builder.Source, opts ...builder.BuildOption) (*builder.Bundle, error) {
	if len(s.Local) == 0 && len(s.Remote) == 0 {
		return nil, errors.New("either local or remote source is needed")
	}
	if len(s.Local) > 0 {
		return r.next.Build(s, opts...)
	}
	// we retrieve the source using the given retriever
	// if no local repository is provided.
	src, err := r.options.Retriever.Retrieve(s.Remote)
	if err != nil {
		return nil, err
	}
	return r.next.Build(src, opts...)
}

func (r retriever) Clean(bun *builder.Bundle, opts ...builder.CleanOption) error {
	return r.next.Clean(bun, opts...)
}

func (r retriever) String() string {
	return "retriever"
}

// New creates a retriever builder whose Build method retrieves the source
// code from the given URL of the remote repository and then uses the next
// builder to build the source code.
func New(next builder.Builder, opts ...builder.Option) builder.Builder {
	var options builder.Options
	for _, o := range opts {
		o(&options)
	}
	return &retriever{
		options: options,
		next:    next,
	}
}
