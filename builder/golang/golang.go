package golang

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/moobu/moo/builder"
)

type golang struct {
	options builder.Options
}

func (g *golang) Build(s *builder.Source, opts ...builder.BuildOption) (*builder.Bundle, error) {
	var options builder.BuildOptions
	for _, o := range opts {
		o(&options)
	}

	bin := s.Name
	cmd := exec.Command("go", "build", "-ldflags=\"-s -w\"", "-o", bin)
	cmd.Stdout = g.options.Output
	cmd.Dir = s.Local

	return &builder.Bundle{
		Ref:    options.Ref,
		Entry:  []string{filepath.Join(s.Local, bin)},
		Source: s,
	}, nil
}

func (g *golang) Clean(b *builder.Bundle, opts ...builder.CleanOption) error {
	// we clean up just by removing the binary
	return os.Remove(filepath.Join(b.Source.Local, b.Source.Name))
}

func (g golang) String() string {
	return "go"
}

// New returns a python builder that uses conda for dependencies isolation.
func New(opts ...builder.Option) builder.Builder {
	var options builder.Options
	for _, o := range opts {
		o(&options)
	}
	return &golang{options: options}
}
