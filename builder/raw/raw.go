package raw

import (
	"errors"

	"github.com/moobu/moo/builder"
)

type raw struct {
	options builder.Options
}

func (r *raw) Build(s *builder.Source, opts ...builder.BuildOption) (*builder.Bundle, error) {
	var options builder.BuildOptions
	for _, o := range opts {
		o(&options)
	}
	if len(s.Local) == 0 && len(s.Remote) == 0 {
		return nil, errors.New("either local or remote source is needed")
	}
	if len(s.Local) > 0 {
		return r.build(s, &options)
	}
	// we retrieve the source using the given retriever
	// if no local repository is provided.
	src, err := r.options.Retriever.Retrieve(s.Remote)
	if err != nil {
		return nil, err
	}
	return r.build(src, &options)
}

func (r *raw) build(s *builder.Source, opts *builder.BuildOptions) (*builder.Bundle, error) {
	return nil, nil
}

func (l *raw) Clean(b *builder.Bundle, opts ...builder.CleanOption) error {
	return nil
}

func New(opts ...builder.Option) builder.Builder {
	var options builder.Options
	for _, o := range opts {
		o(&options)
	}
	return &raw{
		options: options,
	}
}
