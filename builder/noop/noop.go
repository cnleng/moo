package noop

import "github.com/moobu/moo/builder"

type noop struct {
	options builder.Options
}

func (n noop) Build(s *builder.Source, opts ...builder.BuildOption) (*builder.Bundle, error) {
	return &builder.Bundle{}, nil
}

func (n noop) Release(b *builder.Bundle, opts ...builder.ReleaseOption) error {
	return nil
}

func (n noop) Clean(b *builder.Bundle, opts ...builder.CleanOption) error {
	return nil
}

func New(opts ...builder.Option) builder.Builder {
	var options builder.Options
	for _, o := range opts {
		o(&options)
	}
	return &noop{
		options: options,
	}
}
