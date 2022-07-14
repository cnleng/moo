package mixed

import (
	"fmt"

	"github.com/moobu/moo/builder"
)

type mixed struct {
	options builder.Options
	langs   map[string]builder.Builder
}

func (m mixed) Build(s *builder.Source, opts ...builder.BuildOption) (*builder.Bundle, error) {
	if lang, ok := m.langs[s.Type]; ok {
		return lang.Build(s, opts...)
	}
	return nil, fmt.Errorf("no builder implemented for %s", s.Type)
}

func (m mixed) Clean(b *builder.Bundle, opts ...builder.CleanOption) error {
	if lang, ok := m.langs[b.Source.Type]; ok {
		return lang.Clean(b, opts...)
	}
	return fmt.Errorf("no builder implemented for %s", b.Source.Type)
}

func (mixed) String() string {
	return "mixed"
}

// New returns a mixed builder that specifies which named builder use
// according to the type of the source.
func New(bs []builder.Builder, opts ...builder.Option) builder.Builder {
	var options builder.Options
	for _, o := range opts {
		o(&options)
	}

	langs := make(map[string]builder.Builder, len(bs))
	for _, b := range bs {
		langs[b.String()] = b
	}
	return &mixed{
		options: options,
		langs:   langs,
	}
}
